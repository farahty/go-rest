package scalars

import (
	"io"
	"time"

	"github.com/arsmn/fastgql/graphql"
)

type DeletedAt struct{}

func (d DeletedAt) MarshalGQL(date time.Time) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {

		w.Write([]byte(date.String()))

	})
}

func (d DeletedAt) UnmarshalGQL(value interface{}) interface{} {
	switch value := value.(type) {
	case time.Time:
		buff, err := value.MarshalText()
		if err != nil {
			return nil
		}

		return string(buff)
	case *time.Time:
		if value == nil {
			return nil
		}
		return serializeDateTime(*value)
	default:
		return nil
	}
}

func serializeDateTime(value interface{}) interface{} {
	switch value := value.(type) {
	case time.Time:
		buff, err := value.MarshalText()
		if err != nil {
			return nil
		}

		return string(buff)
	case *time.Time:
		if value == nil {
			return nil
		}
		return serializeDateTime(*value)
	default:
		return nil
	}
}
