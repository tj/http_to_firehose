// Package basicauth implements a simple basic auth middleware which
// supports a single set of credentials only.
package basicauth

import (
	"encoding/base64"
	"net/http"
	"strings"
)

// BasicAuth middleware.
type BasicAuth struct {
	Handler  http.Handler
	Username string
	Password string
}

// ServeHTTP implements http.Handler.
func (b BasicAuth) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s := r.Header["Authorization"]

	if len(s) == 0 {
		unauthorized(w)
		return
	}

	authorization := strings.TrimSpace(s[0])
	credentials := strings.Split(authorization, " ")

	if len(credentials) != 2 || credentials[0] != "Basic" {
		unauthorized(w)
		return
	}

	auth, err := base64.StdEncoding.DecodeString(credentials[1])
	if err != nil {
		unauthorized(w)
		return
	}

	userpass := strings.Split(string(auth), ":")
	if len(userpass) != 2 {
		unauthorized(w)
		return
	}

	if userpass[0] == b.Username && userpass[1] == b.Password {
		b.Handler.ServeHTTP(w, r)
		return
	}

	unauthorized(w)
}

// Unauthorized request.
func unauthorized(w http.ResponseWriter) {
	w.Header().Set("WWW-Authenticate", "Basic realm=\"user\"")
	http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
}
