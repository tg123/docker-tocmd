package optvar

import (
	"github.com/docker/docker/api/types"
)

type OptVar interface {
	Generate(container types.ContainerJSON) string
}

var optVars []OptVar

func AddOptVar(v OptVar){
	optVars = append(optVars, v)
}

func GetOptVars() []OptVar {
	return optVars
}