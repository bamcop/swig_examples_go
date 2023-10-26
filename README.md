# README
Learn how to use swig in golang by following the official examples of swig. Out of the box, without Makefile.

- simple. A minimal example showing how SWIG can be used to wrap a C function, a global variable, and a constant.
- constants. This shows how preprocessor macros and certain C declarations are turned into constants.
- variables. An example showing how to access C global variables from Go.
- enum. Wrapping enumerations.
- class. Wrapping a simple C++ class.
- reference. C++ references.
- pointer. Simple pointer handling.
- funcptr. Pointers to functions.
- template. C++ templates.
- callback. C++ callbacks using directors.
- extend. Polymorphism using directors.
- director. Example how to utilize the director feature.
- director. Example how to use goin and godirectorin.

## Test Env
```text
# sw_vers
ProductName:		macOS
ProductVersion:		14.0
BuildVersion:		23A344
```

```text
# go version
go version go1.21.3 darwin/arm64
```

```text
# swig -version

SWIG Version 4.1.1

Compiled with clang++ [aarch64-apple-darwin22.1.0]

Configured options: +pcre

Please see https://www.swig.org for reporting bugs and further information
```
