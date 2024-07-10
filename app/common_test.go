package app

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEnvironment_GetPorts_One(t *testing.T) {
	env := Environment{Port: "1"}
	ports := env.GetPorts()
	assert.Equal(t, len(ports), 1)
	assert.Equal(t, ports[0], "1")
}

func TestEnvironment_GetPorts_Two(t *testing.T) {
	env := Environment{SqsPort: "1", SnsPort: "2"}
	ports := env.GetPorts()
	assert.Equal(t, len(ports), 2)
	assert.Equal(t, ports[0], "1")
	assert.Equal(t, ports[1], "2")
}

func TestEnvironment_GetPorts_Three(t *testing.T) {
	env := Environment{Port: "1", SqsPort: "2", SnsPort: "3"}
	ports := env.GetPorts()
	assert.Equal(t, len(ports), 1)
	assert.Equal(t, ports[0], "1")
}

func TestEnvironment_GetListenAddresses_UnixSocket(t *testing.T) {
	env := Environment{Host: "/path/to/socket"}
	addresses := env.GetListenAddresses()
	assert.Equal(t, len(addresses), 1)
	assert.Equal(t, addresses[0].String(), "unix:///path/to/socket")
}

func TestEnvironment_GetListenAddresses_Tcp_One(t *testing.T) {
	env := Environment{Host: "localhost", Port: "1"}
	addresses := env.GetListenAddresses()
	assert.Equal(t, len(addresses), 1)
	assert.Equal(t, addresses[0].String(), "tcp://0.0.0.0:1")
}

func TestEnvironment_GetListenAddresses_Tcp_Two(t *testing.T) {
	env := Environment{Host: "localhost", SqsPort: "1", SnsPort: "2"}
	addresses := env.GetListenAddresses()
	assert.Equal(t, len(addresses), 2)
	assert.Equal(t, addresses[0].String(), "tcp://0.0.0.0:1")
	assert.Equal(t, addresses[1].String(), "tcp://0.0.0.0:2")
}
