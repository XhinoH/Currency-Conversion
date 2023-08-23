package model

import (
	"database/sql/driver"
	"fmt"
)

// BoolBit represents a boolean value stored as a bit in the database.
type BoolBit bool

// Scan implements the sql.Scanner interface.
func (b *BoolBit) Scan(value interface{}) error {
	if value == nil {
		*b = false
		return nil
	}

	switch v := value.(type) {
	case int64:
		*b = v == 1
	case []uint8:
		*b = string(v) == "\x01"
	default:
		return fmt.Errorf("unexpected type for BoolBit: %T", value)
	}

	return nil
}

// Value implements the driver.Valuer interface.
func (b BoolBit) Value() (driver.Value, error) {
	if b {
		return 1, nil
	}
	return 0, nil
}
