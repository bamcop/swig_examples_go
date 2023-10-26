package pkg_director_test

import (
	"fmt"
	"os"
	"testing"

	example "github.com/bamcop/swig_examples_go/pkg_director"
)

func Compare(name string, got string, exp string) error {
	fmt.Printf("%s; Got: '%s'; Expected: '%s'\n", name, got, exp)
	if got != exp {
		return fmt.Errorf("%s returned unexpected string! Got: '%s'; Expected: '%s'\n", name, got, exp)
	}
	return nil
}

func testFooBarCpp() error {
	fb := example.NewFooBarCpp()
	defer example.DeleteFooBarCpp(fb)
	return Compare("FooBarCpp.FooBar()", fb.FooBar(), "C++ Foo, C++ Bar")
}

func testFooBarGo() error {
	fb := example.NewFooBarGo()
	defer example.DeleteFooBarGo(fb)
	return Compare("FooBarGo.FooBar()", fb.FooBar(), "Go Foo, Go Bar")
}

func TestMain(t *testing.T) {
	fmt.Println("Test output:")
	fmt.Println("------------")
	err := testFooBarCpp()
	err = testFooBarGo()
	fmt.Println("------------")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Tests failed! Last error: %s\n", err.Error())
		os.Exit(1)
	}
}
