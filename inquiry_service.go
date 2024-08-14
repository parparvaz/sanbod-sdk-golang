package sanbod

import (
	"context"
	"net/http"
)

type InquiryUserProfileWithImageService struct {
	c            *Client
	nationalCode string
	birthDate    string
}

func (s *InquiryUserProfileWithImageService) NationalCode(nationalCode string) *InquiryUserProfileWithImageService {
	s.nationalCode = nationalCode
	return s
}

func (s *InquiryUserProfileWithImageService) Birthdate(birthDate string) *InquiryUserProfileWithImageService {
	s.birthDate = birthDate
	return s
}

func (s *InquiryUserProfileWithImageService) Do(ctx context.Context, opts ...RequestOption) (res *InquiryUserProfileWithImage, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sanboom/v1/infoinquiry/personalwithimage",
		secType:  secTypeAccessToken,
	}

	r.setJsonParams(params{
		"nationalId": s.nationalCode,
		"birthDate":  s.birthDate,
	})

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res = new(InquiryUserProfileWithImage)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type InquiryUserProfileWithImage struct {
	Error   bool `json:"error"`
	Message struct {
		FirstName      string `json:"firstName"`
		LastName       string `json:"lastName"`
		RegisterNo     string `json:"registerNo"`
		RegisterSeries string `json:"registerSeries"`
		RegisterSerial string `json:"registerSerial"`
		NationalId     string `json:"nationalId"`
		BirthDate      string `json:"birthDate"`
		BirthPlace     string `json:"birthPlace"`
		DeathStatus    string `json:"deathStatus"`
		Gender         string `json:"gender"`
		FatherName     string `json:"fatherName"`
		Images         []struct {
			Type  string  `json:"type"`
			Image *string `json:"image"`
		} `json:"images"`
	} `json:"message"`
	ResultNumber int    `json:"result_number"`
	TraceId      string `json:"trace_id"`
}

type InquiryUserProfileService struct {
	c            *Client
	nationalCode string
	birthDate    string
}

func (s *InquiryUserProfileService) NationalCode(nationalCode string) *InquiryUserProfileService {
	s.nationalCode = nationalCode
	return s
}

func (s *InquiryUserProfileService) Birthdate(birthDate string) *InquiryUserProfileService {
	s.birthDate = birthDate
	return s
}

func (s *InquiryUserProfileService) Do(ctx context.Context, opts ...RequestOption) (res *InquiryUserProfile, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sanboom/v1/infoinquiry/personal",
		secType:  secTypeAccessToken,
	}

	r.setJsonParams(params{
		"nationalId": s.nationalCode,
		"birthDate":  s.birthDate,
	})

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res = new(InquiryUserProfile)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type InquiryUserProfile struct {
	Error   bool `json:"error"`
	Message struct {
		FirstName      string `json:"firstName"`
		LastName       string `json:"lastName"`
		RegisterNo     string `json:"registerNo"`
		RegisterSeries string `json:"registerSeries"`
		RegisterSerial string `json:"registerSerial"`
		NationalId     string `json:"nationalId"`
		BirthDate      string `json:"birthDate"`
		BirthPlace     string `json:"birthPlace"`
		DeathStatus    string `json:"deathStatus"`
		Gender         string `json:"gender"`
		FatherName     string `json:"fatherName"`
	} `json:"message"`
	ResultNumber int    `json:"result_number"`
	TraceId      string `json:"trace_id"`
}

type IbanInquiryService struct {
	c    *Client
	iban string
}

func (s *IbanInquiryService) Iban(iban string) *IbanInquiryService {
	s.iban = iban

	return s
}

func (s *IbanInquiryService) Do(ctx context.Context, opts ...RequestOption) (res *IbanInquiry, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sanboom/v1/banksinquiry/ibaninquiry",
		secType:  secTypeAccessToken,
	}

	r.setJsonParams(params{
		"iban": s.iban,
	})

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res = new(IbanInquiry)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type IbanInquiry struct {
	Error   bool `json:"error"`
	Message struct {
		Iban               string `json:"iban"`
		BankName           string `json:"bankName"`
		DepositNumber      string `json:"depositNumber"`
		DepositStatus      string `json:"depositStatus"`
		DepositDescription string `json:"depositDescription"`
		DepositComment     string `json:"depositComment"`
		OwnersInfo         []struct {
			FirstName string `json:"firstName"`
			LastName  string `json:"lastName"`
		} `json:"ownersInfo"`
	} `json:"message"`
	ResultNumber int    `json:"result_number"`
	TraceId      string `json:"trace_id"`
}
