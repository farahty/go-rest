package directives

import (
	"context"
	"fmt"

	"github.com/arsmn/fastgql/graphql"
	"github.com/nimerfarahty/go-rest/models"
)

func HasRole(ctx context.Context, obj interface{}, next graphql.Resolver, role models.Role) (res interface{}, err error) {

	status := ctx.Value("status").(string)

	if status != "ok" {
		return nil, fmt.Errorf(status)
	}

	user := ctx.Value("user").(models.UserJWT)

	if user.Role == role {
		return next(ctx)
	}

	return nil, fmt.Errorf("access denied")
}
