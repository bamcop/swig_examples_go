package goin_test

import (
	"fmt"
	"testing"

	example "github.com/bamcop/swig_examples_go/goin"
)

type mycallbacks struct {
	example.Callbacks
}

var tststrs = []string{"A", "BCD", "EFGH"}
var tstint int = 5

func (v *mycallbacks) Call1(val int, strarray []string) bool {
	var rv bool = true

	for i, s := range strarray {
		fmt.Printf("%d: %s\n", i, s)
		if s != tststrs[i] {
			fmt.Printf("  ***Mismatch, expected %s\n", tststrs[i])
			rv = false
		}
	}
	if val != tstint {
		rv = false
	}
	return rv
}

func TestMain(t *testing.T) {
	cbs := &mycallbacks{}
	cbs.Callbacks = example.NewDirectorCallbacks(cbs)
	worked := example.Check1(cbs, tstint, tststrs)
	if !worked {
		panic("Data mismatch")
	}
}
