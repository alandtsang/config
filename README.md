# Config

config is used to parse various configuration files in golang

## Quick Start

To import this package, add the following line to your code:

```
import "github.com/alandtsang/config"
```

Example:

main.go

```go
package main

import (
	"fmt"

	"github.com/alandtsang/config"
)

type Data struct {
	A string `yaml:"a"`
	B B      `yaml:"b"`
}

type B struct {
	C int   `yaml:"c"`
	D []int `yaml:"d"`
}

func main() {
	cfg, err := config.NewConfig("test.yml")
	if err != nil {
		fmt.Printf("NewConfig() failed", err)
		return
	}

	var data Data
	if err = cfg.Bind(&data); err != nil {
		fmt.Printf("Bind structure failed,", err)
		return
	}

	fmt.Printf("%+v\n", data)
}
```

go run main.go

```
{A:Easy! B:{C:2 D:[3 4]}}
```