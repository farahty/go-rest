package scalars

import (
	"errors"
	"io"
	"strconv"

	"github.com/arsmn/fastgql/graphql"
	"github.com/google/uuid"
)

func MarshalID(value uuid.UUID) graphql.Marshaler {

	return graphql.WriterFunc(func(writer io.Writer) {

		io.WriteString(writer, strconv.Quote(value.String()))
	})

}

func UnmarshalID(value interface{}) (uuid.UUID, error) {

	if tempStr, ok := value.(string); ok {

		return uuid.Parse(tempStr)
	}

	return uuid.Nil, errors.New("invalid UUID string")
}
