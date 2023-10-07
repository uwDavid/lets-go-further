package main

import (
	"context"
	"net/http"
	"uwDavid/moviedb/internal/data"
)

type contextKey string

// convert the string "user" into a contextKey type => assign it to constant
// we will use this this const as the key for getting/setting user info in the
// request context.
const userContextKey = contextKey("user")

// contextSetUser() returns a copy of the req w/ the provided User struct
// added to context
func (app *application) contextSetUser(r *http.Request, user *data.User) *http.Request {
	ctx := context.WithValue(r.Context(), userContextKey, user)
	return r.WithContext(ctx)
}

// contextGetUser() retrieves the User struct from the req context.
// we will only use it when we expect there to be a User struct value in the ctx.
// otherwise, we panic()
// Note: context values stored in req context have the type interface{}
// this means that we need to assert it back to the original type before using it
func (app *application) contextGetUser(r *http.Request) *data.User {
	user, ok := r.Context().Value(userContextKey).(*data.User)
	if !ok {
		panic("missing user value in request context")
	}
	return user
}
