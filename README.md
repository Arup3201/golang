**Hello World**

Programs:

- [hellworld](./helloworld/main.go)

**Command Line Arguments**

Programs:

- [echo1](./echo1/main.go)
- [echo2](./echo2/main.go)

**Env Variables**

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
- [goarch](./goarch/main.go)
- [goos](./goarch/main.go)

**Go Modules**

Before diving into the commands, it's important to understand what Go modules are. A Go module is a collection of related Go packages that are versioned together as a single unit. It's Go's modern approach to dependency management that replaced the GOPATH-based system.

`go mod init` - Starting a New Module

`go mod init` initializes a new Go module in the current directory. Think of it as telling Go "this directory contains a new project that I want to manage as a module." It creates the foundation for modern Go dependency management.

When you run `go mod init`, you provide a module path - this is like giving your module a unique name that other modules can use to import it. The module path serves as both an identifier and a hint about where the code might be found.

**Common patterns for module paths:**

- **For GitHub projects:** `github.com/username/projectname`
- **For company projects:** `company.com/team/projectname`
- **For personal projects:** `example.com/myproject` or just `myproject`

When you run `go mod init`, Go creates a file called `go.mod` in your current directory. This file is the heart of your module - it declares the module's identity and tracks its dependencies.

**The go.mod file initially contains:**

- The module declaration (what you specified)
- The Go version requirement (based on your current Go installation)

The module path you choose has important implications:

**If you plan to publish your module:**
Choose a path that reflects where others can find it. If it's on GitHub, use the GitHub path. This makes it easy for others to `go get` your module.

**For private or local modules:**
You can use any valid module path. Some developers use company domains, some use example domains, some just use simple names.

**Import implications:**
Other packages within your module will import each other using paths relative to the module path. If your module path is `github.com/user/myapp`, then a package in `internal/utils/` would be imported as `github.com/user/myapp/internal/utils`.

`go mod tidy` - Cleaning Up Dependencies

`go mod tidy` analyzes your module's source code and ensures that your `go.mod` file accurately reflects what your code actually needs. It's like having a smart assistant that keeps your dependencies organized and minimal.

`go mod tidy` performs two important operations:

**Adding missing dependencies:**
If your code imports packages that aren't listed in `go.mod`, it finds and adds them with appropriate versions.

**Removing unused dependencies:**
If your `go.mod` lists dependencies that your code doesn't actually use, it removes them to keep things clean.

Version Selection Logic:

When `go mod tidy` needs to add a dependency, it follows sophisticated rules:

**Latest compatible version:**
It typically selects the latest version that's compatible with your Go version and other dependencies.

**Semantic versioning awareness:**
It understands semantic versioning and prefers stable releases over pre-releases.

**Minimal version selection:**
Go uses an algorithm called "Minimal Version Selection" - it picks the minimum version that satisfies all requirements, which promotes reproducibility.

The go.sum File:

`go mod tidy` also maintains a `go.sum` file alongside your `go.mod`. This file contains cryptographic checksums of your dependencies, ensuring that you get exactly the same code every time you download dependencies.

**Security and reproducibility:**
The checksums protect against tampering and ensure that your builds are reproducible across different machines and times.

**Not meant for editing:**
Unlike `go.mod`, you should never manually edit `go.sum` - let Go manage it automatically.
