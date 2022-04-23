package model

import (
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

type Date time.Time

// MarshalDate converts a time.Time (date) to a string
func MarshalDate(date time.Time) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, e := io.WriteString(w, fmt.Sprintf("%s%s%s", "\"", date.Format("2006-01-02"), "\""))
		if e != nil {
			panic(e)
		}
	})
}

// UnmarshalDate unmarshalls a string to a time.Time (date)
func UnmarshalDate(v interface{}) (time.Time, error) {
	str, ok := v.(string)
	if !ok {
		return time.Time{}, fmt.Errorf("date must be strings")
	}
	withoutQuotes := strings.ReplaceAll(str, "\"", "")
	i, err := time.Parse("2006-01-02", withoutQuotes)
	if err != nil {
		i, err = time.Parse("20060102", withoutQuotes)
	}
	return i, err
}
