package main

import (
	"context"
	"fmt"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
)

func execAndStreamOutput(imageID, command string) error {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}

	// // Create an exec instance
	// execConfig := types.ExecConfig{
	// 	Cmd:          []string{"/bin/sh", "-c", command},
	// 	AttachStdout: true,
	// 	AttachStderr: true,
	// }
	// execIDResp, err := cli.ContainerExecCreate(ctx, containerID, execConfig)
	// if err != nil {
	// 	return err
	// }

	// // Attach to the exec instance
	// attachResp, err := cli.ContainerExecAttach(ctx, execIDResp.ID, types.ExecStartCheck{})
	// if err != nil {
	// 	return err
	// }
	// defer attachResp.Close()

	// // Stream the output
	// stdcopy.StdCopy(os.Stdout, os.Stderr, attachResp.Reader)

	return nil
}

func main() {
	fmt.Println("Starting docker service...")
	imageID := 
}
