package sanbod

import (
	"context"
	"net/http"
)

type ACTokenService struct {
	c *Client
}

func (s *ACTokenService) Do(ctx context.Context, opts ...RequestOption) (res *GetACToken, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/oautht/v1/authorize",
		secType:  secTypeBasicAuth,
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res = new(GetACToken)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type GetACToken struct {
}

type CCTokenService struct {
	c            *Client
	grantType    string
	scope        []string
	providerCode string
}

func (s *CCTokenService) Scope(scope []string) *CCTokenService {
	s.scope = scope
	return s
}

func (s *CCTokenService) ProviderCode(providerCode string) *CCTokenService {
	s.providerCode = providerCode
	return s
}

func (s *CCTokenService) Do(ctx context.Context, opts ...RequestOption) (res *GetCCToken, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/oauth/v1/token",
		secType:  secTypeBasicAuth,
	}
	r.setJsonParams(params{
		"grant_type":    "client_credentials",
		"scope":         s.scope,
		"provider_code": s.providerCode,
	})

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res = new(GetCCToken)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type GetCCToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
}

type RefreshTokenService struct {
	c *Client
}

func (s *RefreshTokenService) Do(ctx context.Context, opts ...RequestOption) (res *RefreshToken, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/oauth/v1/token",
		secType:  secTypeRefreshToken,
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res = new(RefreshToken)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type RefreshToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RevokeTokenService struct {
	c *Client
}

func (s *RevokeTokenService) Do(ctx context.Context, opts ...RequestOption) (res *RevokeToken, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/oauth/v1/revoke",
		secType:  secTypeNone,
	}

	accessToken, _ := s.c.cache.get(CacheAccessToken)

	r.setFormParams(params{
		"access_token": accessToken,
	})

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res = new(RevokeToken)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type RevokeToken struct {
	Error        bool   `json:"error"`
	Message      string `json:"message"`
	ResultNumber int    `json:"result_number"`
}
