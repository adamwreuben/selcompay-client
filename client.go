// Package client provides support to access the Selcom Pay API service.
package client

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"time"
)

const version = "v1"

// This provides a default client configuration and is set with reasonable
// defaults. Users can replace this client with application specific settings
// using the WithClient function at the time a Client is constructed.
var defaultClient = http.Client{
	Transport: &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 15 * time.Second,
			DualStack: true,
		}).DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	},
}

// Client represents a client that can talk to the selcompay API service.
type Client struct {
	host      string
	apiKey    string
	apiSecret string
	http      *http.Client
}

// New constructs a client that can be used to talk to the selcompay api.
func New(host string, apiKey string, apiSecret string, options ...func(cln *Client)) *Client {
	cln := Client{
		// Create a PR in golang/go to improve the TrimLeft doc.
		host:      strings.TrimLeft(host, "/"),
		apiKey:    apiKey,
		apiSecret: apiSecret,
		http:      &defaultClient,
	}

	for _, option := range options {
		option(&cln)
	}

	return &cln
}

// WithClient adds a custom client for processing requests. It's recommend
// to not use the default client and provide your own.
func WithClient(http *http.Client) func(cln *Client) {
	return func(cln *Client) {
		cln.http = http
	}
}

func (cln *Client) do(ctx context.Context, method string, url string, body any, v any) error {
	resp, err := do(ctx, cln, method, url, body)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, &v); err != nil {
		return fmt.Errorf("client: response: %s: unmarshaling: error: %w", string(data), err)
	}

	return nil
}

func do(ctx context.Context, cln *Client, method string, url string, body any) (*http.Response, error) {
	var b bytes.Buffer
	if body != nil {
		if err := json.NewEncoder(&b).Encode(body); err != nil {
			return nil, fmt.Errorf("encoding request body: error: %w", err)
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, url, &b)
	if err != nil {
		return nil, fmt.Errorf("create request error: %w", err)
	}

	// Set Auth Headers
	if err := cln.setHeaders(req, &b); err != nil {
		return nil, err
	}

	resp, err := cln.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do: error: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("readall: error: %w", err)
		}

		var trustedErr Error
		if err := json.Unmarshal(data, &trustedErr); err != nil {
			return nil, fmt.Errorf("decoding: response: %s, error: %w ", string(data), err)
		}

		trustedErr.statuscode = resp.StatusCode
		return nil, fmt.Errorf("error: response: %s", trustedErr.Error())
	}

	return resp, nil
}

func (cln *Client) setHeaders(req *http.Request, body *bytes.Buffer) error {
	// Set default header values
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	authToken := fmt.Sprintf("SELCOM %s", base64Encode([]byte(cln.apiKey)))
	digestMethod := "HS256"

	signedFields, digest, timestamp, err := constructHeaderValues(cln, body)
	if err != nil {
		return err
	}

	// Set the auth Headers
	req.Header.Set("Authorization", authToken)
	req.Header.Set("Digest-Method", digestMethod)
	req.Header.Set("Digest", digest)
	req.Header.Set("Timestamp", timestamp)
	req.Header.Set("Signed-Fields", signedFields)

	return nil
}

func base64Encode(token []byte) string {
	return base64.StdEncoding.EncodeToString(token)
}

func constructHeaderValues(cln *Client, body *bytes.Buffer) (string, string, string, error) {
	var signedFields string

	var jsonData map[string]any
	if err := json.Unmarshal(body.Bytes(), &jsonData); err != nil {
		return "", "", "", fmt.Errorf("unmarshal error: %w", err)
	}
	now := time.Now().Format(time.RFC3339)

	data := fmt.Sprintf("timestamp=%s", now)
	for k, v := range jsonData {
		data = fmt.Sprintf("%s&%s=%v", data, k, v)
		if len(signedFields) == 0 {
			signedFields = k
		} else {
			signedFields = signedFields + "," + k
		}
	}

	mac := hmac.New(sha256.New, []byte(cln.apiSecret))
	_, err := mac.Write([]byte(data))
	if err != nil {
		return "", "", "", fmt.Errorf("create signature: error: %w", err)
	}

	return signedFields, base64Encode(mac.Sum(nil)), now, nil
}
