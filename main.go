package main

import (
	"strings"
	"context"
	"fmt"
	"os"

	// "github.com/docker/docker/api/types"
	"github.com/tg123/docker-tocmd/opts"

	"github.com/docker/docker/client"
)



func main() {

	if len(os.Args) <= 1 {
		fmt.Printf("usage: %v <containerid>\n", os.Args[0])
		os.Exit(1)
	}

	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	container, err := cli.ContainerInspect(context.Background(), os.Args[1])
	if err != nil {
		panic(err)
	}

	var opts []string

	for _, o := range optvar.GetOptVars() {
		opt := o.Generate(container)

		if opt != "" {
			opts = append(opts, opt)
		}
	}

	
	cmd := fmt.Sprintf("docker run %v %v %v", strings.Join(opts, " "), container.Config.Image, "") // TODO

	// container.Config.Cmd
	fmt.Println(cmd)


}
