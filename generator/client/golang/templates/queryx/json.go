// Code generated by queryx, DO NOT EDIT.

package queryx

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type JSON struct {
	Val  map[string]interface{}
	Null bool
	Set  bool
}

func NewJSON(v map[string]interface{}) JSON {
	return JSON{Val: v}
}

func NewNullableJSON(v interface{}) JSON {
	if v != nil {
		return NewJSON(v.(map[string]interface{}))
	}
	return JSON{Null: true}
}

// Scan implements the Scanner interface.
func (j *JSON) Scan(value interface{}) error {
	if value == nil {
		j.Val, j.Null = nil, true
		return nil
	}
	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("JSON scan source was not []byte")
	}
	return json.Unmarshal(b, &j.Val)
}

// Value implements the driver Valuer interface.
func (j JSON) Value() (driver.Value, error) {
	return json.Marshal(j.Val)
}

// MarshalJSON implements the json.Marshaler interface.
func (j JSON) MarshalJSON() ([]byte, error) {
	if j.Null {
		return json.Marshal(nil)
	}
	return json.Marshal(j.Val)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (j *JSON) UnmarshalJSON(data []byte) error {
	s := string(data)
	if s == "{}" || s == "null" {
		j.Null = true
		return nil
	}
	m := map[string]interface{}{}
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	j.Val = m
	return nil
}
