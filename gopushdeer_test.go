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
	sendErr = gpd.SendText("")
	assert.NotNil(t, sendErr)
	fmt.Println(sendErr.Error())
}

func TestServerSendText(t *testing.T) {
	gpd, err := NewGoPushDeerServer("server", "key")
	assert.Nil(t, err)
	assert.NotNil(t, gpd)
	sendErr := gpd.SendText("hello world")
	assert.Nil(t, sendErr)
}
