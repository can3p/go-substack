package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"

	"github.com/can3p/go-substack/types"
)

const apiRoot = "https://%s.substack.com/api/v1/"

func endpoint(substackName string, endpoint string, args ...any) string {
	str := apiRoot + endpoint
	a := []any{substackName}
	a = append(a, args...)
	return fmt.Sprintf(str, a...)
}

type Client struct {
	httpClient   http.Client
	sessionID    string
	substackName string
}

func NewClient(rt http.RoundTripper, sessionID string, substackName string) (*Client, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	return &Client{
		httpClient: http.Client{
			Transport: rt,
			Jar:       jar,
		},
		sessionID:    sessionID,
		substackName: substackName,
	}, nil
}

func (c *Client) CreateNewDraft(ctx context.Context, draft *types.Draft) (*types.Draft, error) {
	req, err := prepareRequest(ctx, draft, http.MethodPost, c.sessionID, endpoint(c.substackName, "drafts"))

	if err != nil {
		return nil, err
	}

	res, err := c.httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Unexpected error from the server. Response code: %d, Response body: %s", res.StatusCode, string(body))
	}

	var returnDraft types.Draft

	if err := json.Unmarshal(body, &returnDraft); err != nil {
		return nil, err
	}

	return &returnDraft, nil
}

func (c *Client) UpdateDraft(ctx context.Context, draftID int, draft *types.Draft) (*types.Draft, error) {
	req, err := prepareRequest(ctx, draft, http.MethodPut, c.sessionID, endpoint(c.substackName, "drafts/%d", draftID))

	if err != nil {
		return nil, err
	}

	res, err := c.httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Unexpected error from the server. Response code: %d, Response body: %s", res.StatusCode, string(body))
	}

	var returnDraft types.Draft

	if err := json.Unmarshal(body, &returnDraft); err != nil {
		return nil, err
	}

	return &returnDraft, nil
}

func prepareRequest(ctx context.Context, payload any, httpMethod string, sessionID string, url string) (*http.Request, error) {
	jsonBody, err := json.Marshal(payload)

	if err != nil {
		return nil, err
	}

	bodyReader := bytes.NewReader(jsonBody)

	req, err := http.NewRequestWithContext(ctx, httpMethod, url, bodyReader)

	if err != nil {
		return nil, err
	}

	cookie := &http.Cookie{
		Name:   "substack.sid",
		Value:  sessionID,
		MaxAge: 300,
	}

	req.AddCookie(cookie)
	req.Header.Add("Content-Type", "application/json")

	return req, nil
}
