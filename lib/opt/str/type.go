package optstr

import (
	"database/sql/driver"
	"fmt"
)

// String is an implementation of a string option-type.
type String struct {
	value string
	loaded bool
}

// Nothing returns a string option-type that has no value.
func Nothing() String {
	return String{}
}

// Something return a string option-type with a value.
func Something(value string) String {
	return String{
		loaded: true,
		value: value,
	}
}

// GoString makes String fit the fmt.GoStringer interface.
func (receiver String) GoString() string {
	if Nothing() == receiver {
		return "optstr.Nothing()"
	}

	return fmt.Sprintf("optstr.Something(%q)", receiver.value)
}

// MarshalText makes String fit the encoding.TextMarshaler interface.
func (receiver String) MarshalText() (text []byte, err error) {
	if Nothing() == receiver {
		return nil, errNothing
	}

	return []byte(receiver.String()), nil
}

// Scan makes String fit the sql.Scan interfaces.
func (receiver *String) Scan(src interface{}) error {
	if nil == receiver {
		return errNilReceiver
	}

	switch casted := src.(type) {
	case String:
		*receiver = casted
		return nil
	case string:
		return receiver.Set(casted)
	case []byte:
		s := string(casted)
		return receiver.Set(s)
	default:
		return fmt.Errorf("Cannot scan something of type %T into an %T.", src, *receiver)
	}
}

// Set makes String fit the flag.Value interfaces.
func (receiver *String) Set(value string) error {
	if nil == receiver {
		return errNilReceiver
	}

	*receiver = Something(value)
	return nil
}

// String makes String fit the fmt.Stinger, flag.Value interfaces.
func (receiver String) String() string {
	if Nothing() == receiver {
		return ""
	}

	return receiver.value
}

// UnmarshalText makes String fit the encoding.TextUnmarshaler interface.
func (receiver *String) UnmarshalText(text []byte) error {
	if nil == receiver {
		return errNilReceiver
	}

	if nil == text {
		return errNilSource
	}

	receiver.Set(string(text))
	return nil
}

// Value makes String fit the database/sql/driver.Valuer interface.
func (receiver String) Value() (driver.Value, error) {
	if Nothing() == receiver {
		return nil, errNothing
	}

	return receiver.value, nil
}
