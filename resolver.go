package oauth_user_scope_validator

import (
	"context"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) ValidateUserScope(ctx context.Context, user string, scope string) (*UserScopeValidationResponse, error) {
	res := UserScopeValidationResponse{
		Valid: scope != "blah",
	}
	return &res, nil
}
