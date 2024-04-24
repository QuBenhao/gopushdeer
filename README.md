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
	sendErr := gpd.SendText("hello world")
	if sendErr != nil {
		fmt.Println(sendErr.Error())
	}
	err1 := gpd.SendText("Hello", "test 1")
	if err1 != nil {
		fmt.Println(err1.Error())
    }
	err2 := gpd.SendImage("https://img-blog.csdnimg.cn/img_convert/c7d770053270fa13f19582b9bb1b8a64.png", "golang icon")
	if err2 != nil {
		fmt.Println(err2.Error())
    }
	err3 := gpd.SendMarkdown("# gopushdeer\nPushdeer SDK for Golang\n\nWith [Go module](https://github.com/golang/go/wiki/Modules) support, simply add the following import\n\n```\nimport \"github.com/QuBenhao/gopushdeer\"\n```\n\nto your code, and then `go [build|run|test]` will automatically fetch the necessary dependencies.\n\nOtherwise, run the following Go command to install the `gopushdeer` package:\n\n```sh\n$ go get -u github.com/QuBenhao/gopushdeer\n```\n\n### Running GoPushDeer\n\nFirst you need to import gopushdeer package for using GoPushDeer, one simplest example likes the follow:\n\n```go\npackage main\n\nimport (\n\t\"fmt\"\n\t\"github.com/QuBenhao/gopushdeer\"\n)\n\nfunc main() {\n\tgpd, _ := NewGoPushDeer(\"key\")\n\tsendErr := gpd.SendText(\"hello world\")\n\tif sendErr != nil {\n\t\tfmt.Println(sendErr)\n\t}\n}\n```", "README.md")
	if err3 != nil {
		fmt.Println(err3.Error())
    }
}
```