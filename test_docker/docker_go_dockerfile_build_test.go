package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

import (
	"github.com/gruntwork-io/terratest/modules/docker"
)

const tag = "digital-devops-go-maintenance-page"

func TestBuildsWithoutErrors(t *testing.T) {
	buildOptions := &docker.BuildOptions{
		Tags: []string{tag},
	}

	err := docker.BuildE(t, "..", buildOptions)

	if err != nil {
		errorString := err.Error()
		t.Logf("Expected no errors but got %s", errorString)
		t.Fail()
	}

	assert.Equal(t, nil, err)
}
