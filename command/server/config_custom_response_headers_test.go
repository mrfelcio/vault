package server

import (
	"github.com/go-test/deep"
	"testing"
)

var defaultCustomHeaders = map[string]string {
	"Strict-Transport-Security": "max-age=1; domains",
	"Content-Security-Policy": "default-src 'others'",
	"X-Vault-Ignored": "ignored",
	"X-Custom-Header": "Custom header value default",
	"X-Frame-Options": "Deny",
	"X-Content-Type-Options": "nosniff",
	"Content-Type": "text/plain; charset=utf-8",
	"X-XSS-Protection": "1; mode=block",
}

var customHeaders307 = map[string]string {
	"X-Custom-Header": "Custom header value 307",
}

var customHeader3xx = map[string]string {
	"X-Vault-Ignored-3xx": "Ignored 3xx",
	"X-Custom-Header": "Custom header value 3xx",
}

var customHeaders200 = map[string]string {
	"Someheader-200": "200",
	"X-Custom-Header": "Custom header value 200",
}

var customHeader2xx = map[string]string {
	"X-Custom-Header": "Custom header value 2xx",
}

var customHeader400 = map[string]string {
	"Someheader-400": "400",
}

func TestCustomResponseHeadersConfigs(t *testing.T) {
	expectedCustomResponseHeader := map[string]map[string]string {
		"default": defaultCustomHeaders,
		"307": customHeaders307,
		"3xx": customHeader3xx,
		"200": customHeaders200,
		"2xx": customHeader2xx,
		"400": customHeader400,
	}

	config, err := LoadConfigFile("./test-fixtures/config_custom_response_headers_1.hcl")

	if err != nil {
		t.Fatalf("Error encountered when loading config %+v", err)
	}
	if diff := deep.Equal(expectedCustomResponseHeader, config.Listeners[0].CustomResponseHeaders); diff != nil {
		t.Fatalf("parsed custom headers do not match the expected ones")
	}
}
