package model

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"io"
)

type Uuid string

func MarshalUuid(u string) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		if _, err := io.WriteString(w, fmt.Sprintf(`"%s"`, u)); err != nil {
			return
		}
	})
}

func UnmarshalUuid(v interface{}) (string, error) {
	value, err := graphql.UnmarshalString(v)
	if err != nil {
		return "", err
	}

	return value, nil
}
