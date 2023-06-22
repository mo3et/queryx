// Code generated by queryx, DO NOT EDIT.

package queryx

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
)

type Boolean struct {
	Val  bool
	Null bool
	Set  bool
}

func NewBoolean(v bool) Boolean {
	return Boolean{Val: v}
}

func NewNullableBoolean(v *bool) Boolean {
	if v != nil {
		return NewBoolean(*v)
	}
	return Boolean{Null: true}
}

// Scan implements the Scanner interface.
func (b *Boolean) Scan(value interface{}) error {
	n := sql.NullBool{}
	err := n.Scan(value)
	if err != nil {
		return err
	}
	b.Val, b.Null = n.Bool, !n.Valid
	return nil
}

// Value implements the driver Valuer interface.
func (b Boolean) Value() (driver.Value, error) {
	if b.Null {
		return nil, nil
	}
	return b.Val, nil
}

func (b Boolean) MarshalJSON() ([]byte, error) {
	if b.Null {
		return json.Marshal(nil)
	}
	return json.Marshal(b.Val)
}

func (b *Boolean) UnmarshalJSON(bs []byte) error {
	if string(bs) == "null" {
		b.Null = true
		return nil
	}
	if err := json.Unmarshal(bs, &b.Val); err != nil {
		return err
	}
	return nil
}
