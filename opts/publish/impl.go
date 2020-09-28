package optvar

import (
	"strings"
	"fmt"
	"github.com/docker/docker/api/types"
)

func init() {
	mod.impl = func(container types.ContainerJSON) string {

		var ports []string


		for k, v := range container.NetworkSettings.Ports {
			for _, p := range v {
				ports = append(ports, fmt.Sprintf("-p %v:%v:%v", p.HostIP, p.HostPort, k))
			}
		}

		return strings.Join(ports, "-p")
	}
}
