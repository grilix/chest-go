package chest

type AuthenticationService struct {
	client *Client
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	Id       int    `json:"user_id"`
	Username string `json:"username"`
}

type UserResponse struct {
	User *User `json:"user"`
}

func (s *AuthenticationService) Authenticate(credentials *Credentials) (
	*User, *Response, error,
) {
	req, err := s.client.NewRequest("POST", "/api/v1/auth", credentials)
	if err != nil {
		return nil, nil, err
	}

	f := &UserResponse{}
	resp, err := s.client.Do(req, f)

	if err != nil {
		return nil, resp, err
	}

	s.client.Token = resp.Token
	return f.User, resp, nil
}
