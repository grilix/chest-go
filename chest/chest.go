package chest

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	client    *http.Client
	BaseURL   *url.URL
	UserAgent string
	Token     string

	Authentication *AuthenticationService
	Decks          *DeckService
	Collection     *CollectionService
}

type Pagination struct {
	CurrentPage int `json:"current_page"`
	TotalPages  int `json:"total_pages"`
}

type Response struct {
	*http.Response

	Token string
}

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	c := &Client{client: httpClient}

	c.Authentication = &AuthenticationService{client: c}
	c.Decks = &DeckService{client: c}
	c.Collection = &CollectionService{client: c}
	return c
}

func (c *Client) SetURL(urlStr string) {
	baseURL, _ := url.Parse(urlStr)
	c.BaseURL = baseURL
}

func (c *Client) NewRequest(method, path string, body interface{}) (
	*http.Request, error,
) {
	rel := &url.URL{Path: path}
	u := c.BaseURL.ResolveReference(rel)
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	if c.Token != "" {
		req.Header.Set("Authorization", c.Token)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	return req, nil
}

func (c *Client) Do(req *http.Request, v interface{}) (*Response, error) {
	resp, err := c.client.Do(req)

	if err != nil {
		return nil, err
	}
	response := &Response{Response: resp}

	// TODO: Remove 'Bearer' part from token
	response.Token = resp.Header.Get("Authorization")

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(v)

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			decErr := json.NewDecoder(resp.Body).Decode(v)
			if decErr == io.EOF {
				decErr = nil // ignore EOF errors caused by empty response body
			}
			if decErr != nil {
				err = decErr
			}
		}
	}

	return response, err
}
