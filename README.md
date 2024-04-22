# gopushdeer
Pushdeer SDK for Golang

With [Go module](https://github.com/golang/go/wiki/Modules) support, simply add the following import

```
import "github.com/QuBenhao/gopushdeer"
```

to your code, and then `go [build|run|test]` will automatically fetch the necessary dependencies.

Otherwise, run the following Go command to install the `gopushdeer` package:

```sh
$ go get -u github.com/QuBenhao/gopushdeer
```

### Running GoPushDeer

First you need to import gopushdeer package for using GoPushDeer, one simplest example likes the follow:

```go
package main

import (
	"fmt"
	"github.com/QuBenhao/gopushdeer"
)

func main() {
	gpd, _ := NewGoPushDeer("key")
	sendErr := gpd.sendText("hello world")
	if sendErr != nil {
		fmt.Println(sendErr)
	}
}
```