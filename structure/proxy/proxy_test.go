package proxy

import (
	"fmt"
	"testing"
)

func TestProxy(t *testing.T) {
	t.Run("Proxy", func(t *testing.T) {
		nginxServer := newNginxServer()
		appStatusURL := "/app/status"
		createUserURL := "/create/user"

		httpCode, body := nginxServer.handleRequest(appStatusURL, "GET")
		fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

		httpCode, body = nginxServer.handleRequest(appStatusURL, "GET")
		fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

		httpCode, body = nginxServer.handleRequest(appStatusURL, "GET")
		fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", appStatusURL, httpCode, body)

		httpCode, body = nginxServer.handleRequest(createUserURL, "POST")
		fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", createUserURL, httpCode, body)

		httpCode, body = nginxServer.handleRequest(createUserURL, "GET")
		fmt.Printf("\nUrl: %s\nHttpCode: %d\nBody: %s\n", createUserURL, httpCode, body)
	})
}
