package gopushdeer

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBasicSendText(t *testing.T) {
	gpd, err := NewGoPushDeer("key")
	assert.Nil(t, err)
	assert.NotNil(t, gpd)
	sendErr := gpd.SendText("hello world")
	assert.Nil(t, sendErr)
	sendErr = gpd.SendText("This is cool", "some description here")
	assert.Nil(t, sendErr)
	sendErr = gpd.SendText("")
	assert.NotNil(t, sendErr)
	fmt.Println(sendErr.Error())
}

func TestSendImage(t *testing.T) {
	gpd, _ := NewGoPushDeer("PDU25472T6XKvfE7MCSHdULY2OVpXrnuJrTIirTHS")
	sendErr := gpd.SendImage("https://img-blog.csdnimg.cn/img_convert/c7d770053270fa13f19582b9bb1b8a64.png", "golang icon")
	assert.Nil(t, sendErr)
}

func TestSendMarkdown(t *testing.T) {
	gpd, _ := NewGoPushDeer("PDU25472T6XKvfE7MCSHdULY2OVpXrnuJrTIirTHS")
	sendErr := gpd.SendMarkdown("# gopushdeer\nPushdeer SDK for Golang\n\nWith [Go module](https://github.com/golang/go/wiki/Modules) support, simply add the following import\n\n```\nimport \"github.com/QuBenhao/gopushdeer\"\n```\n\nto your code, and then `go [build|run|test]` will automatically fetch the necessary dependencies.\n\nOtherwise, run the following Go command to install the `gopushdeer` package:\n\n```sh\n$ go get -u github.com/QuBenhao/gopushdeer\n```\n\n### Running GoPushDeer\n\nFirst you need to import gopushdeer package for using GoPushDeer, one simplest example likes the follow:\n\n```go\npackage main\n\nimport (\n\t\"fmt\"\n\t\"github.com/QuBenhao/gopushdeer\"\n)\n\nfunc main() {\n\tgpd, _ := NewGoPushDeer(\"key\")\n\tsendErr := gpd.SendText(\"hello world\")\n\tif sendErr != nil {\n\t\tfmt.Println(sendErr)\n\t}\n}\n```", "README.md")
	assert.Nil(t, sendErr)
}

func TestServerSendText(t *testing.T) {
	gpd, err := NewGoPushDeerServer("server", "key")
	assert.Nil(t, err)
	assert.NotNil(t, gpd)
	sendErr := gpd.SendText("hello world")
	assert.Nil(t, sendErr)
}
