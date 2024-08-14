package sanbod

import (
	"context"
	"net/http"
)

type CardToAccountNumberService struct {
	c          *Client
	cardNumber string
}

func (j *CardToAccountNumberService) CardNumber(cardNumber string) *CardToAccountNumberService {
	j.cardNumber = cardNumber

	return j
}

func (j *CardToAccountNumberService) Do(ctx context.Context, opts ...RequestOption) (res *CardToAccountNumber, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sanboom/v1/banksinquiry/cardtodeposit",
		secType:  secTypeAccessToken,
	}

	r.setJsonParams(params{
		"cardNumber": j.cardNumber,
	})
	data, err := j.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res = new(CardToAccountNumber)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type CardToAccountNumber struct {
	Error   bool `json:"error"`
	Message struct {
		CardNumber    string `json:"cardNumber"`
		DepositNumber string `json:"depositNumber"`
	} `json:"message"`
	ResultNumber int    `json:"result_number"`
	TraceId      string `json:"trace_id"`
}

type CardToIbanService struct {
	c          *Client
	cardNumber string
}

func (j *CardToIbanService) CardNumber(cardNumber string) *CardToIbanService {
	j.cardNumber = cardNumber
	return j
}

func (j *CardToIbanService) Do(ctx context.Context, opts ...RequestOption) (res *CardToIban, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sanboom/v1/banksinquiry/cardtoiban",
		secType:  secTypeAccessToken,
	}

	r.setJsonParams(params{
		"cardNumber": j.cardNumber,
	})

	data, err := j.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res = new(CardToIban)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type CardToIban struct {
	Error   bool `json:"error"`
	Message struct {
		CardNumber string `json:"cardNumber"`
		Iban       string `json:"iban"`
	} `json:"message"`
	ResultNumber int    `json:"result_number"`
	TraceId      string `json:"trace_id"`
}

type AccountNumberToIbanService struct {
	c             *Client
	provider      string
	depositNumber string
}

func (j *AccountNumberToIbanService) Provider(provider string) *AccountNumberToIbanService {
	j.provider = provider

	return j
}

func (j *AccountNumberToIbanService) DepositNumber(depositNumber string) *AccountNumberToIbanService {
	j.depositNumber = depositNumber

	return j
}

func (j *AccountNumberToIbanService) Do(ctx context.Context, opts ...RequestOption) (res *AccountNumberToIban, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sanboom/v1/banksinquiry/deposittoiban",
		secType:  secTypeAccessToken,
	}

	r.setJsonParams(params{
		"provider":      j.provider,
		"depositNumber": j.depositNumber,
	})

	data, err := j.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res = new(AccountNumberToIban)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type AccountNumberToIban struct {
	Error   bool `json:"error"`
	Message struct {
		DepositNumber string `json:"depositNumber"`
		Iban          string `json:"iban"`
	} `json:"message"`
	ResultNumber int    `json:"result_number"`
	TraceId      string `json:"trace_id"`
}

type IbanToAccountNumberService struct {
	c    *Client
	iban string
}

func (j *IbanToAccountNumberService) Iban(iban string) *IbanToAccountNumberService {
	j.iban = iban

	return j
}

func (j *IbanToAccountNumberService) Do(ctx context.Context, opts ...RequestOption) (res *IbanToAccountNumber, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/banks/v1/ibantodeposit",
		secType:  secTypeAccessToken,
	}

	r.setJsonParams(params{
		"iban": j.iban,
	})

	data, err := j.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res = new(IbanToAccountNumber)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type IbanToAccountNumber struct {
	Error   bool `json:"error"`
	Message struct {
		DepositNumber string `json:"depositNumber"`
		Iban          string `json:"iban"`
	} `json:"message"`
	ResultNumber int    `json:"result_number"`
	TraceId      string `json:"trace_id"`
}
