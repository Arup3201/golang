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

[How to upgrade your Go module?](https://go.dev/blog/using-go-modules)

**Note about `:=` operator**

In a `:=` declaration a variable v may appear even if it has already been declared, provided:

- this declaration is in the same scope as the existing declaration of v (if v is already declared in an outer scope, the declaration will create a new variable §),
- the corresponding value in the initialization is assignable to v, and
- there is at least one other variable that is created by the declaration.

**Slices**

*Why two slices can't be compared using `==` operator?*

If we could use == for slices, then it would be confusing for go users to understand the purpose. Some might want to check which elements are present in the slice while others might think it is checking whether 2 slice has the same address. To avoid this confusion, they are not allowing == between 2 slices. 

But they allow == with nil to check whether the slice is nil-slice or not. It is important because both nil-slice and non-nil empty slice has len 0, so we can't use len for checking that. This is what I understood.

*Why we need to return a slice from a function that modified the slice?*

When we pass the slice to a function, we need to return the slice if we are modifying the slice inside the function. 

Golang pass the parameters by copying the values of the arguments. It is not passing their reference. So when we pass the slice, we are passing a copy of the slice structure. It means it's pointer to the array, cap and len. 

If we modify the slice by appending something to the slice, then at one point golang might create a new slice because the old slice can't accomodate the growing capacity which also changes it's capacity and len. If that happens, it will start pointing to another address. 

If we don't return this new slice, and use the same slice that we passed to this function, then the slice actually points to the previous slice and not the modified slice - which is not correct and will give the same result as if we never modified it.

**References**

- [How To Debug Go Code with Visual Studio Code](https://www.digitalocean.com/community/tutorials/debugging-go-code-with-visual-studio-code)
- [Graceful Shutdown in Go: Practical Patterns](https://victoriametrics.com/blog/go-graceful-shutdown/)
- [Using Go Modules](https://go.dev/blog/using-go-modules)
- [Go Modules vs Package](https://stackoverflow.com/questions/61940117/go-modules-vs-package)
- [How to fix parsing go.mod module declares its path as "x" but was required as "y"](https://stackoverflow.com/questions/61311436/how-to-fix-parsing-go-mod-module-declares-its-path-as-x-but-was-required-as-y)
- [Arrays, slices (and strings): The mechanics of 'append'](https://go.dev/blog/slices)
- [Demystifying Golang Slices](https://medium.com/@andreiboar/demystifying-golang-slices-83ffe3550db5)
- [How to Compare Slice Equality in Go](https://freshman.tech/snippets/go/compare-slices/)
- [They're called Slices because they have Sharp Edges: Even More Go Pitfalls](https://www.dolthub.com/blog/2023-10-20-golang-pitfalls-3/)
- [How UTF-8 works in Go?](https://pandulaofficial.medium.com/unicode-utf-8-explained-with-examples-using-go-5f8b7f4521d)
- [The Absolute Minimum Every Software Developer Absolutely, Positively Must Know About Unicode and Character Sets (No Excuses!)](https://www.joelonsoftware.com/2003/10/08/the-absolute-minimum-every-software-developer-absolutely-positively-must-know-about-unicode-and-character-sets-no-excuses/)
- [Strings, bytes, runes and characters in Go](https://go.dev/blog/strings)
- [Go Structs and Interfaces Made Simple](https://getstream.io/blog/go-structs-interfaces/)
- [Interface methods](https://go.dev/doc/effective_go#interface_methods)