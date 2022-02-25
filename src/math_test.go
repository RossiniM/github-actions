package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSum(t *testing.T) {
	result := Sum(1, 3)
	require.Equal(t, result, 4)
}
