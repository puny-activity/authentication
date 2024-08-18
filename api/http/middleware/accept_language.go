package middleware

import (
	"context"
	"github.com/puny-activity/authentication/pkg/base/ctxbase"
	"github.com/puny-activity/authentication/pkg/base/headerbase"
	"net/http"
)

func (m *Middleware) AcceptLanguageMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		acceptLang := r.Header.Get(headerbase.AcceptLanguage)

		r = r.WithContext(context.WithValue(r.Context(), ctxbase.Language, acceptLang))
		next.ServeHTTP(w, r)
	})
}
