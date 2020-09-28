package optvar

import (
	"strings"
	"fmt"
	"github.com/docker/docker/api/types"
)

func init() {
	mod.impl = func(container types.ContainerJSON) string {

		var binds []string

		// container.Mounts.

		// for k, v := range container.HostConfig.Binds {
		// 	for _, p := range v {
		// 		ports = append(ports, fmt.Sprintf("-p %v:%v:%v", p.HostIP, p.HostPort, k))
		// 	}
		// }

		// if len(container.HostConfig.Binds) > 0 
		for _, b := range container.HostConfig.Binds {
			binds = append(binds, fmt.Sprintf("-v %v", b))
		}

		return strings.Join(binds, "")
	}
}
