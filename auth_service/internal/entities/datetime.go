package entities

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

const (
	timeFormat = time.RFC3339
)

type DateTime time.Time

// UnmarshalGQL implements the graphql.Unmarshaler interface
func (dt *DateTime) UnmarshalGQL(v interface{}) error {
	datetime, ok := v.(string)
	if !ok {
		return fmt.Errorf("DateTime must be a string")
	}

	t, err := time.Parse(timeFormat, datetime)
	if !ok {
		return err
	}

	dt = (*DateTime)(&t)

	return nil
}

// MarshalGQL implements the graphql.Marshaler interface
func (dt DateTime) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(time.Time(dt).Format(timeFormat)))
}
