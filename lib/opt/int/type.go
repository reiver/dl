package optint

import (
	"database/sql/driver"
	"fmt"
	"strconv"
)

// Int is an implementation of a string option-type.
type Int struct {
	value int
	loaded bool
}

// Nothing returns a string option-type that has no value.
func Nothing() Int {
	return Int{}
}

// Something return a string option-type with a value.
func Something(value int) Int {
	return Int{
		loaded: true,
		value: value,
	}
}

// GoString makes Int fit the fmt.GoStringer interface.
func (receiver Int) GoString() string {
	if Nothing() == receiver {
		return "optstr.Nothing()"
	}

	return fmt.Sprintf("optstr.Something(%q)", receiver.value)
}

// MarshalText makes Int fit the encoding.TextMarshaler interface.
func (receiver Int) MarshalText() (text []byte, err error) {
	if Nothing() == receiver {
		return nil, errNothing
	}

	return []byte(receiver.String()), nil
}

// Scan makes Int fit the sql.Scan interfaces.
func (receiver *Int) Scan(src interface{}) error {
	if nil == receiver {
		return errNilReceiver
	}

	switch casted := src.(type) {
	case Int:
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

// Set makes Int fit the flag.Value interfaces.
func (receiver *Int) Set(value string) error {
	if nil == receiver {
		return errNilReceiver
	}

	i64, err := strconv.ParseInt(value, 10, 0)
	if nil != err {
		return err
	}
	var i int = int(i64)

	*receiver = Something(i)
	return nil
}

// String makes Int fit the fmt.Stinger, flag.Value interfaces.
func (receiver Int) String() string {
	if Nothing() == receiver {
		return ""
	}

	return strconv.FormatInt(int64(receiver.value), 10)
}

// UnmarshalText makes Int fit the encoding.TextUnmarshaler interface.
func (receiver *Int) UnmarshalText(text []byte) error {
	if nil == receiver {
		return errNilReceiver
	}

	if nil == text {
		return errNilSource
	}

	receiver.Set(string(text))
	return nil
}

// Value makes Int fit the database/sql/driver.Valuer interface.
func (receiver Int) Value() (driver.Value, error) {
	if Nothing() == receiver {
		return nil, errNothing
	}

	return receiver.value, nil
}
