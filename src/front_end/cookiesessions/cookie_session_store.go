package cookiesessions

import (
	"github.com/NullJupiter/GoTodoApp/src/front_end/helper"
	"github.com/gorilla/sessions"
)

var cookieSessionStore *sessions.CookieStore

// InitCookieSessionStore function is used to initialize a new cookie session store.
func InitCookieSessionStore() error {
	authKey, err := helper.GenerateRandomByteSlice(64)
	if err != nil {
		return err
	}
	cookieSessionStore = sessions.NewCookieStore(authKey)

	return nil
}

// GetCookieSessionStore function is used as a getter function to get the cookie session store from another package.
func GetCookieSessionStore() *sessions.CookieStore {
	return cookieSessionStore
}
