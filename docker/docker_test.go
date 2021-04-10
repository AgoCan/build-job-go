package docker

import (
	"testing"
)

func TestBuild(t *testing.T) {

	opt := NewOptions()
	opt.Workspace = "../test/build"

	opt.Build()
}

func TestPush(t *testing.T) {
	opt := NewOptions()
	opt.Push()
}

func TestRemove(t *testing.T) {
	opt := NewOptions()
	opt.Remove()
}
func TestParseInstruction(t *testing.T) {

}
func TestParseParam(t *testing.T) {

}
