# stackerror
[![PkgGoDev](https://pkg.go.dev/badge/github.com/bikbah/stackerror)](https://pkg.go.dev/github.com/bikbah/stackerror)

Simple golang package to handle errors with stack info

## Usage

1. Install
    ```
    go get -v github.com/bikbah/stackerror
    ```

2. Use in error returning code:

    ```go
    package main

    import (
        "errors"
        "log"

        "github.com/bikbah/stackerror"
    )

    func main() {
        if err := f(); err != nil {
            // prints error with stack
            log.Fatal(err)
        }
    }

    func f() error {
        if err := g(); err != nil {
            return stackerror.New(err)
        }

        return nil
    }

    func g() error {
        return errors.New("error in stack depth")
    }
    ```

3. Output

    ```
    2023/11/01 23:48:16 error in stack depth
    /dev/go/src/test/stackerror/main.go:18[main.f]
    /dev/go/src/test/stackerror/main.go:11[main.main]
    /usr/local/go/src/runtime/proc.go:267[runtime.main]
    /usr/local/go/src/runtime/asm_arm64.s:1197[runtime.goexit]
    exit status 1
    ```
