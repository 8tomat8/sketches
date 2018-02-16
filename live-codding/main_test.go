package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBookCreate(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BookCreate(tt.args.w, tt.args.r)
		})
	}
}

func TestHttpHandler(t *testing.T) {
	srv := httptest.NewServer(http.DefaultServeMux)
	srv.Start()

	w := httptest.NewRecorder()
}
