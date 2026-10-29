package main

import (
	"os/exec"

	"github.com/docker/distribution/reference"
)

// pull fetches an image named in the request.
func pull(name string) ([]byte, error) {
	ref, err := reference.ParseNormalizedNamed(name)
	if err != nil {
		return nil, err
	}
	return exec.Command("sh", "-c", "docker pull "+ref.String()).Output()
}
