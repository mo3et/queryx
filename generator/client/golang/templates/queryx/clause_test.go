package queryx

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClauseAddOr(t *testing.T) {
	c1 := &Clause{
		fragment: "a = ?",
		args:     []interface{}{1},
	}
	c2 := &Clause{
		fragment: "b = ?",
		args:     []interface{}{"x"},
	}

	c3 := c1.And(c2)
	require.Equal(t, "(a = ?) AND (b = ?)", c3.fragment)
	require.Equal(t, []interface{}{1, "x"}, c3.args)

	c4 := c1.Or(c2)
	require.Equal(t, "(a = ?) OR (b = ?)", c4.fragment)
	require.Equal(t, []interface{}{1, "x"}, c4.args)

	c5 := c1.And(c1.Or(c2))
	require.Equal(t, "(a = ?) AND ((a = ?) OR (b = ?))", c5.fragment)
	require.Equal(t, []interface{}{1, 1, "x"}, c5.args)

	require.Equal(t, "a = ?", c1.fragment)
	require.Equal(t, []interface{}{1}, c1.args)

	require.Equal(t, "b = ?", c2.fragment)
	require.Equal(t, []interface{}{"x"}, c2.args)
}
