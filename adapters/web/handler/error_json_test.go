package handler

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHandler_jsonError(t *testing.T) {
	msg := "Hello Json"
	result := jsonError(msg)
	expected := []byte(`{"message":"Hello Json"}`)
	require.Equal(t, expected, result)
}
