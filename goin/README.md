# README

## cmd
### 1、generate binding code
```sh
swig -go -cgo -c++ -intgosize 64 example.i
```

### 2、run test
```sh
go test -v .
```

### 3、clean up
```sh
rm -f example_wrap.cxx example_wrap.h goin.go
```
