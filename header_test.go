package tnt

import (
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	file, err := os.Open("examples/[V] Burnt Island.tnt")
	if err != nil {
		t.Error(err)
	}
	obj, err := Parse(file)
	if err != nil {
		t.Error(err)
	}
	t.Log(obj)
}
