# README

## cmd
### 1、generate binding code
```sh
swig -go -cgo -intgosize 64 example.i
```

### 2、run test
```sh
go test -v .
```

### 3、clean up
```sh
rm -f example_wrap.c multimap.go
```
