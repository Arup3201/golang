**Hello World**

Programs: 
- [hellworld](./helloworld/main.go)

**Command Line Arguments**

Programs:
- [echo1](./echo1/main.go)
- [echo2](./echo2/main.go)

**GOROOT, GOPATH, GOOS and GOARCH**

`GOROOT` is the location where the Go programming language is installed on your system. It contains standard library location (`$GOROOT/src`), Go compiles and toolchain (`go build`, `go run`, `go fmt`, `go test`, `go mod` etc) and also platform specific binaries of the standard library files(`$GOROOT/pkg`).

Platform specific locations:

- Windows:

```
C:\Program Files\Go\      (Standard installer)
C:\Go\                    (Alternative location)
```

- Linux:

```
/usr/local/go/           (Manual installation)
/usr/bin/go              (Package manager installation)
/snap/go/current/        (Snap installation)
```

- Mac:

```
/usr/local/go/           (Standard installer)
/opt/homebrew/bin/go     (Homebrew on Apple Silicon)
/usr/local/bin/go        (Homebrew on Intel Mac)
```

Programs:
- [goroot](./goroot/main.go)
