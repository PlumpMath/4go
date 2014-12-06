package easy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSayHello(t *testing.T) {
	helloMsg := SayHello("Jitendra")
	assert.Equal(t, "Hello, Jitendra", helloMsg)
}

// Test Greet Message
func TestGreetMsg(t *testing.T) {
	greetMsg := GreetMsg("Jitendra")
	assert.Equal(t, "Good Day, Jitendra", greetMsg)
}
