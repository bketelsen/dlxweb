package cmd

import (
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
	lxd "github.com/lxc/lxd/client"
	"github.com/lxc/lxd/shared"
	"github.com/lxc/lxd/shared/api"
)

func restConsoleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Not implemented", http.StatusNotImplemented)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Get the id argument
	containerName := r.FormValue("id")
	if containerName == "" {
		http.Error(w, "Missing session id", 400)
		return
	}

	// Get console width and height
	width := r.FormValue("width")
	height := r.FormValue("height")

	if width == "" {
		width = "150"
	}

	if height == "" {
		height = "20"
	}

	widthInt, err := strconv.Atoi(width)
	if err != nil {
		http.Error(w, "Invalid width value", 400)
	}

	heightInt, err := strconv.Atoi(height)
	if err != nil {
		http.Error(w, "Invalid width value", 400)
	}

	// Setup websocket with the client
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Internal server error", 500)
		return
	}
	defer conn.Close()

	// Connect to the container
	env := make(map[string]string)
	env["USER"] = "bjk"
	env["HOME"] = "/home/bjk"
	env["TERM"] = "xterm"

	inRead, inWrite := io.Pipe()
	outRead, outWrite := io.Pipe()

	// read handler
	go func(conn *websocket.Conn, r io.Reader) {
		in := shared.ReaderToChannel(r, -1)

		for {
			buf, ok := <-in
			if !ok {
				break
			}

			err = conn.WriteMessage(websocket.TextMessage, buf)
			if err != nil {
				break
			}
		}
	}(conn, outRead)

	// write handler
	go func(conn *websocket.Conn, w io.Writer) {
		for {
			mt, payload, err := conn.ReadMessage()
			if err != nil {
				if err != io.EOF {
					break
				}
			}

			switch mt {
			case websocket.BinaryMessage:
				continue
			case websocket.TextMessage:
				w.Write(payload)
			default:
				break
			}
		}
	}(conn, inWrite)

	// control socket handler
	handler := func(conn *websocket.Conn) {
		for {
			_, _, err = conn.ReadMessage()
			if err != nil {
				break
			}
		}
	}

	req := api.ContainerExecPost{
		Command:     []string{"bash"},
		WaitForWS:   true,
		Interactive: true,
		Environment: env,
		Width:       widthInt,
		Height:      heightInt,
	}

	execArgs := lxd.ContainerExecArgs{
		Stdin:    inRead,
		Stdout:   outWrite,
		Stderr:   outWrite,
		Control:  handler,
		DataDone: make(chan bool),
	}
	lxdDaemon, err := lxd.ConnectLXDUnix("", nil)
	if err != nil {
		http.Error(w, "Internal server error", 500)
		return
	}

	op, err := lxdDaemon.ExecContainer(containerName, req, &execArgs)
	if err != nil {
		http.Error(w, "Internal server error", 500)
		return
	}

	err = op.Wait()
	if err != nil {
		http.Error(w, "Internal server error", 500)
		return
	}

	<-execArgs.DataDone

	inWrite.Close()
	outRead.Close()

}
