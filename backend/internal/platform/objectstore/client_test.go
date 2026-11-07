package objectstore

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/nifex007/sirng/backend/internal/platform/config"
)

type transport func(*http.Request) (*http.Response, error)

func (f transport) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

func TestCheck(t *testing.T) {
	t.Setenv("AWS_ACCESS_KEY_ID", "test")
	t.Setenv("AWS_SECRET_ACCESS_KEY", "test")
	t.Setenv("AWS_SESSION_TOKEN", "")
	for _, status := range []int{200, 403, 404} {
		client, err := New(context.Background(), config.Connections{Region: "us-east-1", Endpoint: "http://s3.invalid", PathStyle: true})
		if err != nil {
			t.Fatal(err)
		}
		// Inject a transport without opening a socket or contacting AWS.
		opts := client.Options()
		opts.HTTPClient = &http.Client{Transport: transport(func(r *http.Request) (*http.Response, error) {
			if r.Method != "HEAD" || r.URL.Host != "s3.invalid" || r.URL.Path != "/test-bucket" {
				t.Errorf("unexpected request: %s %s", r.Method, r.URL)
			}
			return &http.Response{StatusCode: status, Header: make(http.Header), Body: io.NopCloser(strings.NewReader("")), Request: r}, nil
		})}
		client = s3.New(opts)
		err = Check(context.Background(), client, "test-bucket")
		if (err == nil) != (status == 200) {
			t.Fatalf("status %d: %v", status, err)
		}
	}
}
