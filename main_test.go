package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestVersion(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := setupRouter()
	w := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "/version", nil)
	if err != nil {
		t.Fatal(err)
	}

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, w.Code)
	}

	expected := `{"version":"0.1.0"}`
	if w.Body.String() != expected {
		t.Fatalf("expected body %s, got %s", expected, w.Body.String())
	}
}
