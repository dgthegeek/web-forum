//go:build go1.11
// +build go1.11

package session

import "net/http"

// newCookieFromOptions returns an http.Cookie with the options set.
func newCookieFromOptions(name, value string, options *Options) *http.Cookie {
	return &http.Cookie{
		Name:     name,
		Value:    value,
		MaxAge:   options.MaxAge,
		Secure:   options.Secure,
		HttpOnly: options.HttpOnly,
	}

}
