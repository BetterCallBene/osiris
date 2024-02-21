package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func execAndStreamOutput(imageID, command string) error {

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}

	containerConfig := container.Config{
		Image:        imageID,
		AttachStdin:  true,
		AttachStdout: true,
		AttachStderr: true,
		Tty:          true,
		//		Cmd:          []string{"/bin/sh", "-c", command},
	}

	// getpid and create unique container name
	containerName := fmt.Sprintf("%d-%s", os.Getpid(), "we-like-to-party")
	resp, err := cli.ContainerCreate(ctx, &containerConfig, nil, nil, nil, containerName)
	if err != nil {
		return err
	}

	defer cli.ContainerRemove(ctx, resp.ID, container.RemoveOptions{})

	if err := cli.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
		return err
	}

	defer cli.ContainerStop(ctx, resp.ID, container.StopOptions{})

	optionLogs := container.LogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Follow:     true,
	}

	out, err := cli.ContainerLogs(ctx, resp.ID, optionLogs)
	if err != nil {
		return err
	}

	// Stream the output
	// stdcopy.StdCopy(os.Stdout, os.Stderr, out)
	io.Copy(os.Stdout, out)

	// Stream the outputnic
	//stdcopy.StdCopy(os.Stdout, os.Stderr, attachResp.Reader)

	// Wait for the container to finish
	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			return err
		}
	case <-statusCh:
		fmt.Println("Container finished")
	}

	return nil
}

func main() {
	fmt.Println("Starting docker service...")
	imageID := "osiris-example:latest"
	command := "echo 'Hello, World!' && tail -f /dev/null"
	err := execAndStreamOutput(imageID, command)

	if err != nil {
		fmt.Println(err)
	}
}
