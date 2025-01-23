package app

import (
	"go.temporal.io/sdk/workflow"
)

func GreetingSomeOne(ctx workflow.Context, name string) (string, error) {
	return "Hello " + name + "!\n", nil
}
