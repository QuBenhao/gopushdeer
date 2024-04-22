package gopushdeer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBasicSendText(t *testing.T) {
	gpd, err := NewGoPushDeer("key")
	assert.Nil(t, err)
	assert.NotNil(t, gpd)
	sendErr := gpd.sendText("hello world")
	assert.Nil(t, sendErr)
}

func TestServerSendText(t *testing.T) {
	gpd, err := NewGoPushDeerServer("server", "key")
	assert.Nil(t, err)
	assert.NotNil(t, gpd)
	sendErr := gpd.sendText("hello world")
	assert.Nil(t, sendErr)
}
