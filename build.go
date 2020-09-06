package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/pkg/jsonmessage"
	"github.com/docker/docker/pkg/term"
	"github.com/jhoonb/archivex"
	"os"

	"github.com/docker/docker/client"
)

func main() {
	cli, err := client.NewEnvClient()
	if err != nil {
		fmt.Printf("%s", err.Error())
	}

	tar := new (archivex.TarFile)
	dockerfile, err := os.Open("./Dockerfile")
	defer dockerfile.Close()
	if err != nil {
		fmt.Printf("%s", err.Error())
	}

	// Had to create a tar file to pass as build context to the docker ImageBuild function.
	tar.Create("/tmp/doctl-kubectl.tar")
	tar.Add("Dockerfile", dockerfile, nil)
	tar.AddAll("bin", true)
	tar.Close()
	dockerBuildContext, err := os.Open("/tmp/doctl-kubectl.tar")
	defer dockerBuildContext.Close()
	ctx := context.Background()
	buildResponse , err := cli.ImageBuild(ctx, dockerBuildContext, types.ImageBuildOptions{

		Tags:       []string{"samthom/doctl-kubectl:1.46"},
		Dockerfile: "./Dockerfile",
	})
	if err != nil {
		fmt.Printf("%s", err.Error())
	}
	defer buildResponse.Body.Close()
	fmt.Printf("********* %s **********\n", buildResponse.OSType)
	termFd, isTerm := term.GetFdInfo(os.Stderr)

	jsonmessage.DisplayJSONMessagesStream(buildResponse.Body, os.Stderr, termFd, isTerm, nil)
}