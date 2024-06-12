package middleware

import (
	"context"
	"github.com/puny-activity/authentication/pkg/ctxbase"
	"net/http"
)

func (m *Middleware) AcceptLanguageMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		supportedLanguages := []string{"en-US", "ru-RU"}
		defaultLanguage := "en-US"
		acceptLang := r.Header.Get("Accept-Language")

		for _, supportedLanguage := range supportedLanguages {
			if supportedLanguage == acceptLang {
				r = r.WithContext(context.WithValue(r.Context(), ctxbase.CtxLang, acceptLang))
				next.ServeHTTP(w, r)
				return
			}
		}

		r = r.WithContext(context.WithValue(r.Context(), ctxbase.CtxLang, defaultLanguage))
		next.ServeHTTP(w, r)
	})
}
