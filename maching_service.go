package sanbod

import (
	"context"
	"net/http"
)

type MatchNationalCodeWithMobileNumberService struct {
	c            *Client
	mobileNumber string
	nationalCode string
}

func (s *MatchNationalCodeWithMobileNumberService) MobileNumber(mobileNumber string) *MatchNationalCodeWithMobileNumberService {
	s.mobileNumber = mobileNumber
	return s
}

func (s *MatchNationalCodeWithMobileNumberService) NationalCode(nationalCode string) *MatchNationalCodeWithMobileNumberService {
	s.nationalCode = nationalCode
	return s
}

func (s *MatchNationalCodeWithMobileNumberService) Do(ctx context.Context, opts ...RequestOption) (res *MatchNationalCodeWithMobileNumber, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sanboom/v1/infomatching/mobilenationalid",
		secType:  secTypeAccessToken,
	}

	r.setJsonParams(params{
		"mobileNumber": s.mobileNumber,
		"nationalId":   s.nationalCode,
	})

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res = new(MatchNationalCodeWithMobileNumber)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type MatchNationalCodeWithMobileNumber struct {
	Error   bool `json:"error"`
	Message struct {
		IsMatched bool `json:"ismatched"`
	} `json:"message"`
	ResultNumber int    `json:"result_number"`
	TraceId      string `json:"trace_id"`
}

type MatchNationalCodeWithCardNumberService struct {
	c            *Client
	mobileNumber string
	nationalCode string
	cardNumber   string
}

func (s *MatchNationalCodeWithCardNumberService) MobileNumber(mobileNumber string) *MatchNationalCodeWithCardNumberService {
	s.mobileNumber = mobileNumber
	return s
}

func (s *MatchNationalCodeWithCardNumberService) NationalCode(nationalCode string) *MatchNationalCodeWithCardNumberService {
	s.nationalCode = nationalCode
	return s
}

func (s *MatchNationalCodeWithCardNumberService) CardNumber(cardNumber string) *MatchNationalCodeWithCardNumberService {
	s.cardNumber = cardNumber
	return s
}

func (s *MatchNationalCodeWithCardNumberService) Do(ctx context.Context, opts ...RequestOption) (res *MatchNationalCodeWithCardNumber, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sanboom/v1/infomatching/cardnationalid",
		secType:  secTypeAccessToken,
	}

	r.setJsonParams(params{
		"mobileNumber": s.mobileNumber,
		"nationalId":   s.nationalCode,
		"cardNumber":   s.cardNumber,
	})

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res = new(MatchNationalCodeWithCardNumber)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type MatchNationalCodeWithCardNumber struct {
	Error   bool `json:"error"`
	Message struct {
		IsMatched bool `json:"ismatched"`
	} `json:"message"`
	ResultNumber int    `json:"result_number"`
	TraceId      string `json:"trace_id"`
}
