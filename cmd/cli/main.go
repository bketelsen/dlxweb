package main

import (
	"context"
	"fmt"
	"time"

	"github.com/bketelsen/dlxweb/generated/client"
)

// Not yet a real useful command, but a test of the client currently
func main() {
	cl := client.New("https://thopter.goat-snake.ts.net/oto/")
	images(cl)
	//instances(cl)
	//profiles(cl)

	//instances(cl)

}

func profiles(cl *client.Client) {
	profileService := client.NewProfileService(cl)
	resp, err := profileService.List(context.Background(), client.ProfileListRequest{})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

func images(cl *client.Client) {
	imageService := client.NewImageService(cl)
	buildRequest := client.ImageBuildRequest{}
	resp, err := imageService.Build(context.Background(), buildRequest)
	if err != nil {
		fmt.Println("remote error:", err)
		return
	}
	fmt.Println(resp)

}

func instances(cl *client.Client) {

	instanceService := client.NewInstanceService(cl)
	list(instanceService)

	/*	stopRequest := client.InstanceStopRequest{
			Name: "clitest",
		}
		stopResponse, err := instanceService.Stop(context.Background(), stopRequest)
		if err != nil {
			fmt.Println("remote error:", err)
			return
		}
		fmt.Println(stopResponse)
	*/
	cresp, err := instanceService.Create(context.Background(), client.InstanceCreateRequest{
		Name:    "testservice",
		Project: "services",
	})
	if err != nil {
		fmt.Println("remote error:", err)
		return
	}
	fmt.Println(cresp)
	time.Sleep(time.Second * 3)

	resp, err := instanceService.List(context.Background(), client.InstanceListRequest{
		Project: "services",
	})
	if err != nil {
		fmt.Println("remote error:", err)
		return
	}
	fmt.Println(resp)
	/*	list(instanceService)
		startRequest := client.InstanceStartRequest{
			Name: "clitest",
		}
		startResponse, err := instanceService.Start(context.Background(), startRequest)
		if err != nil {
			fmt.Println("remote error:", err)
			return
		}
		fmt.Println(startResponse)

		list(instanceService)
	*/
}

func list(svc *client.InstanceService) {

	resp, err := svc.List(context.Background(), client.InstanceListRequest{
		Project: "",
	})
	if err != nil {
		fmt.Println("remote error:", err)
		return
	}
	fmt.Println(resp)
}
