package main

import (
	"build-job-go/docker"
	"build-job-go/git"
)

func main() {
	git.Git()
	opt := docker.NewOptions()

	opt.Build()
	opt.Push()
	opt.Remove()
}
