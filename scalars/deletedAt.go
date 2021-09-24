package scalars

import (
	"errors"
	"io"
	"strconv"
	"time"

	"github.com/arsmn/fastgql/graphql"
	"gorm.io/gorm"
)

func MarshalDeletedAt(t gorm.DeletedAt) graphql.Marshaler {
	if t.Time.IsZero() {
		return graphql.Null
	}

	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.Quote(t.Time.Format(time.RFC3339)))
	})
}

func UnmarshalDeletedAt(v interface{}) (gorm.DeletedAt, error) {
	if tmpStr, ok := v.(string); ok {
		value, err := time.Parse(time.RFC3339, tmpStr)
		if err != nil {
			return gorm.DeletedAt{}, err
		}

		return gorm.DeletedAt{Time: value, Valid: true}, nil
	}
	return gorm.DeletedAt{}, errors.New("time should be RFC3339 formatted string")
}
