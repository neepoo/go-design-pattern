package singleton

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSingleton(t *testing.T)  {
	c1 := GetInstance()
	c2 := GetInstance()

	require.Equal(t, c1, c2)
	c1.AddOne() // 1
	c1.AddOne() // 2
	require.Equal(t, c1.Get(), c2.Get())
	require.Equal(t, c1.Get(), 2)
}