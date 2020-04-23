package middleware

import (
	"log"
	"net/http"
)

// ContentTypeMiddleware defines properties
type ContentTypeMiddleware struct {
	next http.Handler
}

// NewContentTypeMiddleware define function
func NewContentTypeMiddleware(next http.Handler) *ContentTypeMiddleware {
	return &ContentTypeMiddleware{next: next}
}

func (m *ContentTypeMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("Method: %s, URI: %s\n", r.Method, r.RequestURI)
	w.Header().Set("Content-Type", "application/json")
	m.next.ServeHTTP(w, r)
}
