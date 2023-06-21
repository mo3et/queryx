// Code generated by queryx, DO NOT EDIT.

package queryx

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewBigInt(t *testing.T) {
	i := NewBigInt(2)
	require.Equal(t, int64(2), i.Val)
	require.Equal(t, false, i.Null)
}

func TestNewNullableBigInt(t *testing.T) {
	i := NewNullableBigInt(nil)
	require.Equal(t, true, i.Null)
}

func TestBigIntScan(t *testing.T) {
	i := NewBigInt(2)
	err := i.Scan(3)
	require.NoError(t, err)
	require.Equal(t, int64(3), i.Val)
}

func TestBigIntValue(t *testing.T) {
	i := NewBigInt(2)
	value, err := i.Value()
	require.NoError(t, err)
	require.Equal(t, int64(2), value)
}

func TestBigIntMarshalJSON(t *testing.T) {
	i := NewBigInt(2)
	_, err := i.MarshalJSON()
	require.NoError(t, err)
}

func TestBigIntUnmarshalJSON(t *testing.T) {
	i := NewBigInt(2)
	bytes, _ := i.MarshalJSON()
	b := NewBigInt(3)
	err := b.UnmarshalJSON(bytes)
	require.NoError(t, err)
	require.Equal(t, int64(2), i.Val)
	require.Equal(t, false, b.Null)
}
