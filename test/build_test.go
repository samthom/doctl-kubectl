package test

import (
	"context"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/strslice"
	"github.com/docker/docker/client"
	"testing"
)

func TestBuild(t *testing.T) {
	cli, err := client.NewEnvClient()
	if err != nil {
		t.Error("Failed to create the client: ", err)
	}
	ctx := context.Background()
	_, err = cli.ContainerCreate(ctx,
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
		"doctl-kubectl")
	if err != nil {
		t.Error("Failed to create container: ", err)
	}
}
