package auth

import "context"

type contextKey string

var userKey = contextKey("username")

func setUserInContext(ctx context.Context, username string) context.Context {
	return context.WithValue(ctx, userKey, username)
}

func GetUserFromContext(ctx context.Context) string {
	username, _ := ctx.Value(userKey).(string)
	return username
}
