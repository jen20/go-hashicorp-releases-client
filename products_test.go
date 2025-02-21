package releases_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	releases "github.com/jen20/go-hashicorp-releases-client"
)

var testProducts = []string{"consul", "nomad", "packer", "terraform", "waypoint"}

func TestClient_Products(t *testing.T) {
	server := httptest.NewServer(makeTestProductsHandler(t))

	client, err := releases.New(releases.WithBaseURL(server.URL))
	if err != nil {
		t.Fatalf("Unexpected error constructing releases client: %v", err)
	}

	products, err := client.Products(context.Background())
	if err != nil {
		t.Fatalf("Unexpected error retrieving products: %v", err)
	}

	if !reflect.DeepEqual(products, testProducts) {
		t.Fatalf("Got products %v, expected %v", products, testProducts)
	}
}

func makeTestProductsHandler(t *testing.T) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/v1/products":
			w.Header().Set("Content-Type", "application/vnd+hashicorp.releases-api.v1+json")
			w.WriteHeader(http.StatusOK)

			body, err := json.MarshalIndent(testProducts, "", "    ")
			if err != nil {
				t.Fatalf("Failed to marshal test body: %v", err)
			}

			if _, err := w.Write(body); err != nil {
				t.Fatalf("Failed to write response body: %v", err)
			}
		default:
			t.Fatalf("Requested unexpected path: %s", r.URL.Path)
		}
	})
}
