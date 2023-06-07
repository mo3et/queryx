// Code generated by queryx, DO NOT EDIT.

package queryx

import (
	"database/sql"
	"database/sql/driver"
)

type Integer struct {
	Val  int32
	Null bool
	Set  bool
}

func NewInteger(v int32) Integer {
	return Integer{Val: v}
}

func NewNullableInteger(v *int32) Integer {
	if v != nil {
		return NewInteger(*v)
	}
	return Integer{Null: true}
}

// Scan implements the Scanner interface.
func (i *Integer) Scan(value interface{}) error {
	n := sql.NullInt64{}
	err := n.Scan(value)
	if err != nil {
		return err
	}
	i.Val, i.Null = int32(n.Int64), !n.Valid
	return nil
}

// Value implements the driver Valuer interface.
func (i Integer) Value() (driver.Value, error) {
	if i.Null {
		return nil, nil
	}
	return int64(i.Val), nil
}

func (i Integer) MarshalJSON() ([]byte, error) {
	return nil, nil
}

func (i *Integer) UnmarshalJSON(text []byte) error {
	return nil
}
