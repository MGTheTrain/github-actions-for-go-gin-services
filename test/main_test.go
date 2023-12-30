package hello_world_service_test

import (
	"hello_world_service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHelloWorldMessage(t *testing.T) {
	expectedMessage := "Hello World from Go"

	message := hello_world_service.GetHelloWorldMessage()

	assert.Equal(t, expectedMessage, message)
}
