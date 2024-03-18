# Installation
```bash
go get github.com/ErtixNodes/go-lxc
```

## Example
```go
package main

import (
    "log"
    "github.com/ErtixNodes/go-lxc"
)

func main() {
    client, err := lxc.New("vps")
    if err != nil {
        log.Fatal(err)
    }
}
```
