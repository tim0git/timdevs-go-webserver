package test

import (
	"github.com/gruntwork-io/terratest/modules/docker"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/shell"
	"github.com/stretchr/testify/assert"
	"testing"
)

func removeContainer(t *testing.T, id string) {
	cmd := shell.Command{
		Command: "docker",
		Args:    []string{"container", "rm", "--force", id},
	}

	shell.RunCommand(t, cmd)
}

func TestStartsWithoutErrors(t *testing.T) {
	t.Parallel()
	// append timestamp to container name to allow running tests in parallel
	name := "inspect-test-" + random.UniqueId()

	// running the container detached to allow inspection while it is running
	options := &docker.RunOptions{
		Detach: true,
		Name:   name,
		EnvironmentVariables: []string{
			"PORT=8443",
		},
	}

	id := docker.RunAndGetID(t, tag, options)
	defer removeContainer(t, id)

	c := docker.Inspect(t, id)

	exitCode0 := uint8(0x0)

	assert.Equal(t, exitCode0, c.ExitCode)
}
