package httpapi

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/nifex007/sirng/backend/internal/platform/config"
)

func TestStartupFailure(t *testing.T) {
	err := Run(context.Background(), config.Config{Port: "invalid"}, NewHandler(), slog.Default())
	if err == nil {
		t.Fatal("expected listen error")
	}
}

func TestShutdown(t *testing.T) {
	for _, force := range []bool{false, true} {
		t.Run(map[bool]string{false: "drain", true: "deadline"}[force], func(t *testing.T) {
			listener, err := net.Listen("tcp", "127.0.0.1:0")
			if err != nil {
				t.Fatal(err)
			}
			entered := make(chan struct{})
			release := make(chan struct{})
			handlerDone := make(chan struct{})
			server := &http.Server{Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				defer close(handlerDone)
				close(entered)
				<-release
				w.WriteHeader(204)
			})}
			defer server.Close()
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			done := make(chan error, 1)
			timeout := time.Second
			if force {
				timeout = 20 * time.Millisecond
			}
			go func() { done <- serve(ctx, server, listener, slog.New(slog.NewTextHandler(io.Discard, nil)), timeout) }()
			clientDone := make(chan error, 1)
			go func() {
				client := &http.Client{Timeout: 2 * time.Second}
				resp, err := client.Get("http://" + listener.Addr().String())
				if err == nil {
					resp.Body.Close()
				}
				clientDone <- err
			}()
			<-entered
			cancel()
			if !force {
				close(release)
			}
			select {
			case err := <-done:
				if force && !errors.Is(err, context.DeadlineExceeded) {
					t.Errorf("expected deadline, got %v", err)
				}
				if !force && err != nil {
					t.Error(err)
				}
			case <-time.After(3 * time.Second):
				t.Error("shutdown stuck")
			}
			if force {
				close(release)
			}
			<-handlerDone
			if err := <-clientDone; !force && err != nil {
				t.Error(err)
			}
		})
	}
}

func TestRouting(t *testing.T) {
	for _, tc := range []struct {
		method, path string
		status       int
	}{{"HEAD", "/api/health", 200}, {"GET", "/missing", 404}, {"DELETE", "/api/health", 405}} {
		w := httptest.NewRecorder()
		NewHandler().ServeHTTP(w, httptest.NewRequest(tc.method, tc.path, nil))
		if w.Code != tc.status {
			t.Errorf("%s %s: %d", tc.method, tc.path, w.Code)
		}
		if tc.method == "HEAD" && w.Body.Len() != 0 {
			t.Error("HEAD body")
		}
		if tc.status == 405 && w.Header().Get("Allow") != "GET, HEAD" {
			t.Error("missing Allow")
		}
	}
}
