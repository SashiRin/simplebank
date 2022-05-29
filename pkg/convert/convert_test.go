package convert

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConvert(t *testing.T) {
	a := StrTo("-1").MustInt64()
	b, err := StrTo("sdasfsf").Int64()

	require.Equal(t, int64(-1), a)
	require.Error(t, err)
	require.Equal(t, int64(0), b)
}
