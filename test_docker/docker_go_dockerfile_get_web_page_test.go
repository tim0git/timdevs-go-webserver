package test

import (
	"github.com/gruntwork-io/terratest/modules/docker"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/shell"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReturnsWebPage(t *testing.T) {
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
		OtherOptions: []string{
			"-p", "8443:8443",
		},
	}

	id := docker.Run(t, tag, options)
	defer removeContainer(t, id)

	// http request to the container
	cmd := shell.Command{
		Command: "curl",
		Args:    []string{"-s", "http://localhost:8443"},
	}

	// run the command
	output, _ := shell.RunCommandAndGetOutputE(t, cmd)

	// assert contains the essential HTML elements
	assert.Contains(t, output, "<!DOCTYPE html>")
	assert.Contains(t, output, "<html lang=\"en\" dir=\"ltr\">")
	assert.Contains(t, output, "<title>This site is currently unavailable due to essential maintenance, please try again later.</title>")
	assert.Contains(t, output, "<link href=\"temp.css\" rel=\"stylesheet\">")
	assert.Contains(t, output, "<img src=\"gear.png\" alt=\"Gear Image\" />")
	assert.Contains(t, output, "<p>Our website is currently down for maintenance.<br><br>We are sorry for any inconvenience caused.<br><br>We value your business, please visit us later to obtain a quotation.<br><br>Thank you.</p>")
}

func TestReturnsWebPageOnAnyPath(t *testing.T) {
	t.Parallel()
	// append timestamp to container name to allow running tests in parallel
	name := "inspect-test-" + random.UniqueId()

	// running the container detached to allow inspection while it is running
	options := &docker.RunOptions{
		Detach: true,
		Name:   name,
		EnvironmentVariables: []string{
			"PORT=8444",
		},
		OtherOptions: []string{
			"-p", "8444:8444",
		},
	}

	id := docker.Run(t, tag, options)
	defer removeContainer(t, id)

	// http request to the container
	cmd := shell.Command{
		Command: "curl",
		Args:    []string{"-s", "http://localhost:8444/unknown/path"},
	}

	// run the command
	output, _ := shell.RunCommandAndGetOutputE(t, cmd)

	// assert contains the essential HTML elements
	assert.Contains(t, output, "<!DOCTYPE html>")
	assert.Contains(t, output, "<html lang=\"en\" dir=\"ltr\">")
	assert.Contains(t, output, "<title>This site is currently unavailable due to essential maintenance, please try again later.</title>")
	assert.Contains(t, output, "<link href=\"temp.css\" rel=\"stylesheet\">")
	assert.Contains(t, output, "<img src=\"gear.png\" alt=\"Gear Image\" />")
	assert.Contains(t, output, "<p>Our website is currently down for maintenance.<br><br>We are sorry for any inconvenience caused.<br><br>We value your business, please visit us later to obtain a quotation.<br><br>Thank you.</p>")
}

func TestReturnsWebPageOnAnyPathIncludingQueries(t *testing.T) {
	t.Parallel()
	// append timestamp to container name to allow running tests in parallel
	name := "inspect-test-" + random.UniqueId()

	// running the container detached to allow inspection while it is running
	options := &docker.RunOptions{
		Detach: true,
		Name:   name,
		EnvironmentVariables: []string{
			"PORT=8445",
		},
		OtherOptions: []string{
			"-p", "8445:8445",
		},
	}

	id := docker.Run(t, tag, options)
	defer removeContainer(t, id)

	// http request to the container
	cmd := shell.Command{
		Command: "curl",
		Args:    []string{"-s", "http://localhost:8445/unknown/path/?portal.Launch.rest.json"},
	}

	// run the command
	output, _ := shell.RunCommandAndGetOutputE(t, cmd)

	// assert contains the essential HTML elements
	assert.Contains(t, output, "<!DOCTYPE html>")
	assert.Contains(t, output, "<html lang=\"en\" dir=\"ltr\">")
	assert.Contains(t, output, "<title>This site is currently unavailable due to essential maintenance, please try again later.</title>")
	assert.Contains(t, output, "<link href=\"temp.css\" rel=\"stylesheet\">")
	assert.Contains(t, output, "<img src=\"gear.png\" alt=\"Gear Image\" />")
	assert.Contains(t, output, "<p>Our website is currently down for maintenance.<br><br>We are sorry for any inconvenience caused.<br><br>We value your business, please visit us later to obtain a quotation.<br><br>Thank you.</p>")
}
