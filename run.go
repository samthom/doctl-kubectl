package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/strslice"
	"github.com/docker/docker/client"
)
// REFERENCE CODE
func main() {
	//old := os.Stdout // For keeping the backup of the real stdout
	//r, w, _ := os.Pipe()
	//os.Stdout = w
	cli, err := client.NewEnvClient()
	if err != nil {
		fmt.Printf("%s", err.Error())
	}
	ctx := context.Background()

	resp, err := cli.ContainerCreate(ctx,
		&container.Config{
			AttachStderr:true,
			AttachStdout:true,
			AttachStdin:true,
			OpenStdin: true,
			StdinOnce: true,
			Tty: true,
			//Cmd: strslice.StrSlice{
			//	"kubectl version --short --client",
			//	"doctl version",
			//},
			Image: "samthom/doctl-kubectl:1.46",
			Entrypoint: strslice.StrSlice{
				"kubectl version --short --client",
			},
		},
		nil,
		nil,
		"doctl")


	if err != nil {
		fmt.Printf("%s", err.Error())
	}
	//outc := make(chan string)
	//go func() {
	//	var buf bytes.Buffer
	//	io.Copy(&buf, r)
	//	outc <- buf.String()
	//}()

	//w.Close()
	//os.Stdout = old
	//out := <-outc
	fmt.Println("Built ID: ", resp.Warnings)
	//fmt.Println("Result of the test: ", out)
}
