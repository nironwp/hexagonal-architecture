package handler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeProductHandlers(t *testing.T) {
	msg := "Hello json"
	result := jsonError(msg)
	assert.Equal(t, []byte(`{"message":"Hello json"}`), result)
}
