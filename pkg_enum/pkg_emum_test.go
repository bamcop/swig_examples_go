package pkg_enum_test

import (
	"fmt"
	"testing"

	. "github.com/bamcop/swig_examples_go/pkg_enum"
)

func TestMain(t *testing.T) {
	// Print out the value of some enums
	fmt.Println("*** color ***")
	fmt.Println("    RED = ", RED)
	fmt.Println("    BLUE = ", BLUE)
	fmt.Println("    GREEN = ", GREEN)

	fmt.Println("\n*** Foo::speed ***")
	fmt.Println("    Foo::IMPULSE = ", FooIMPULSE)
	fmt.Println("    Foo::WARP = ", FooWARP)
	fmt.Println("    Foo::LUDICROUS = ", FooLUDICROUS)

	fmt.Println("\nTesting use of enums with functions")

	Enum_test(RED, FooIMPULSE)
	Enum_test(BLUE, FooWARP)
	Enum_test(GREEN, FooLUDICROUS)

	fmt.Println("\nTesting use of enum with class method")
	f := NewFoo()

	f.Enum_test(FooIMPULSE)
	f.Enum_test(FooWARP)
	f.Enum_test(FooLUDICROUS)
}
