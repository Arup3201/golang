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

Windows:

```
C:\Program Files\Go\      (Standard installer)
C:\Go\                    (Alternative location)
```

Linux:

```
/usr/local/go/           (Manual installation)
/usr/bin/go              (Package manager installation)
/snap/go/current/        (Snap installation)
```

Mac:

```
/usr/local/go/           (Standard installer)
/opt/homebrew/bin/go     (Homebrew on Apple Silicon)
/usr/local/bin/go        (Homebrew on Intel Mac)
```

`GOPATH` is Go's workspace concept - it's a directory path that tells Go where to find your source code, compiled packages, and executable binaries. Think of it as Go's "home base" for organizing all your Go development work.

It contains

- `src/Directory`: This is where all source code lives. The directory structure under `src/` mirrors import paths exactly.
- `bin/Directory`: This contains executable binaries that you've compiled and installed using `go install`. When you install a Go program, it goes here automatically. Many developers add `$GOPATH/bin` to their system PATH so they can run these programs from anywhere.
- `pkg/ Directory`: This stores compiled package objects - intermediate compiled versions of packages that speed up future builds. Go manages this directory automatically, and it's organized by operating system and architecture (like `pkg/linux_amd64/` or `pkg/darwin_arm64/`).

Modules introduced a new concept - the module cache (usually at `$GOPATH/pkg/mod` or `$HOME/go/pkg/mod`).

`GOARCH` is the systems architecture on which Go program is running. There are multiple types of architecture -

- `amd64`
- `arm64`
- `386`
- `ppc64le`
- many more...

Programs:

- [goroot](./goroot/main.go)
