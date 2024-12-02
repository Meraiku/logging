package logging

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewLogger(t *testing.T) {
	l := NewLogger()

	require.NotNil(t, l)

	require.Equal(t, l, Default())
}
