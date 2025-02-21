package releases

import (
	"net/http"
	"net/url"
	"testing"
)

func TestWithHTTPClient(t *testing.T) {
	t.Run("Overridden Client", func(t *testing.T) {
		testClient := &http.Client{}
		clientOpts, err := newClientOpts(WithHTTPClient(testClient))
		if err != nil {
			t.Fatalf("Error applying option: %v", err)
		}

		if clientOpts.httpClient != testClient {
			t.Fatal("WithHTTPClient must set HTTP Client option to supplied value")
		}
	})

	t.Run("Overridden nil HTTP Client", func(t *testing.T) {
		clientOpts, err := newClientOpts(WithHTTPClient(nil))
		if err != nil {
			t.Fatalf("Error applying option: %v", err)
		}

		if clientOpts.httpClient != http.DefaultClient {
			t.Fatal("WithHTTPClient called with nil must set HTTP Client option to http.DefaultClient")
		}
	})

	t.Run("Default HTTP Client", func(t *testing.T) {
		clientOpts, err := newClientOpts()
		if err != nil {
			t.Fatalf("Error applying option: %v", err)
		}

		if clientOpts.httpClient != http.DefaultClient {
			t.Fatal("Default HTTP Client must be http.DefaultClient")
		}
	})
}

func TestWithUserAgent(t *testing.T) {
	testUserAgent := "Test User Agent"
	clientOpts, err := newClientOpts(WithUserAgent(testUserAgent))
	if err != nil {
		t.Fatalf("Error applying option: %v", err)
	}

	if *clientOpts.userAgent != testUserAgent {
		t.Fatal("WithUserAgent must set User Agent option")
	}
}

func TestWithoutUserAgent(t *testing.T) {
	clientOpts, err := newClientOpts(WithoutUserAgent())
	if err != nil {
		t.Fatalf("Error applying option: %v", err)
	}

	if clientOpts.userAgent != nil {
		t.Fatal("WithoutUserAgent must set User Agent option to nil")
	}
}

func TestWithBaseURL(t *testing.T) {
	t.Run("Overridden Base URL", func(t *testing.T) {
		baseURL := url.URL{
			Scheme: "https",
			Host:   "api.example.com",
			Path:   "releases",
		}
		clientOpts, err := newClientOpts(WithBaseURL(baseURL.String()))
		if err != nil {
			t.Fatalf("Error applying option: %v", err)
		}

		if clientOpts.baseURL.String() != baseURL.String() {
			t.Fatalf("WithBaseURL must set Base URL option to %q, was %q", baseURL.String(), clientOpts.baseURL.String())
		}
	})

	t.Run("Default Base URL", func(t *testing.T) {
		clientOpts, err := newClientOpts()
		if err != nil {
			t.Fatalf("Error applying option: %v", err)
		}

		if clientOpts.baseURL.String() != defaultBaseURL.String() {
			t.Fatalf("Default Base URL must be %q, was %q", defaultBaseURL.String(), clientOpts.baseURL.String())
		}
	})

	t.Run("Overridden Invalid URL", func(t *testing.T) {
		_, err := newClientOpts(WithBaseURL("localhost\x00"))
		if err == nil {
			t.Fatal("Invalid base URL must product error during option application error")
		}
	})
}
