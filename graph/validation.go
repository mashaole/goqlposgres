package graph

import (
	"context"
	"goqlposgress/validator"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func validation(ctx context.Context, v validator.Validation) bool {
	isValid, errors := v.Validate()
	if !isValid {
		for key, er := range errors {
			graphql.AddError(ctx, &gqlerror.Error{
				Message: er,
				Extensions: map[string]interface{}{
					"field": key,
				},
			})
		}
	}
	return isValid
}
