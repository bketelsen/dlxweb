
build: clean generate 
	CGO_ENABLED=0 go build -ldflags="-extldflags=-static" -o bin/dlx github.com/bketelsen/dlxweb/cmd/dlx/

.PHONY: frontend
frontend:
	cd frontend && npm run build && cd ..

tailwind:
	cd frontend && npm run build:tailwind && cd ..

remote:
	ssh thopter sudo ~/rundlx.sh

remote-kill:
	ssh thopter pkill dlxweb

clean:
	rm -rf bin/dlxweb
	rm -rf bin/dlx

generate:
	oto -template ./templates/client.go.plush \
		-out ./generated/client/client.gen.go \
	    -ignore Ignorer \
	    -pkg client \
	    ./definitions
	oto -template ./templates/server.go.plush \
		-out ./generated/server/oto.gen.go \
	    -ignore Ignorer \
	    -pkg server \
	    ./definitions
	oto -template ./templates/client.js.plush \
		-out ./generated/javascript/oto.gen.js \
	    -ignore Ignorer \
	    ./definitions
	cp ./generated/javascript/oto.gen.js ./frontend/src/oto.js
	go fmt ./generated/client/client.gen.go
	go fmt ./generated/server/oto.gen.go

deps:
	go install github.com/pacedotdev/oto@latest
