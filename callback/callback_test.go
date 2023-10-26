package callback_test

import (
	"fmt"
	"testing"

	. "github.com/bamcop/swig_examples_go/callback"
)

func TestMain(t *testing.T) {
	fmt.Println("Adding and calling a normal C++ callback")
	fmt.Println("----------------------------------------")

	caller := NewCaller()
	callback := NewCallback()

	caller.SetCallback(callback)
	caller.Call()
	caller.DelCallback()

	go_callback := NewGoCallback()

	fmt.Println()
	fmt.Println("Adding and calling a Go callback")
	fmt.Println("--------------------------------")

	caller.SetCallback(go_callback)
	caller.Call()
	caller.DelCallback()

	DeleteGoCallback(go_callback)

	fmt.Println()
	fmt.Println("Go exit")
}
