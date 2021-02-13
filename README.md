# symlink

[![CI Status](https://img.shields.io/github/workflow/status/avakarev/go-symlink/Go%20CI%20Workflow/master?longCache=tru&label=CI%20Status&logo=github%20actions&logoColor=fff)](https://github.com/avakarev/go-symlink/actions?query=branch%3Amaster+workflow%3A%22Go+CI+Workflow%22)
[![Docs](https://pkg.go.dev/badge/github.com/avakarev/go-symlink)](https://pkg.go.dev/github.com/avakarev/go-symlink)

> Tiny Go library to manipulate symlinks

## Usage

The `symlink` provides such operations as `Read`, `Link` and `Unlink`.

```go
package main

import (
	"fmt"

	"github.com/avakarev/go-symlink"
)

func main() {
	sym := symlink.New("/path/to/source", "/path/to/target")
	if err := sym.Link(); err != nil {
		fmt.Printf("not linked: %s\n", err)
		return
	}
	fmt.Println("linked!")
}
```

## License

`go-symlink` is licensed under MIT license. (see [LICENSE](./LICENSE))
