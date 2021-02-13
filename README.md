# symlink

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
