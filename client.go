package sanbod

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"github.com/google/uuid"
	jsoniter "github.com/json-iterator/go"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const (
	baseAPIMainURL = "https://api.sanbod.co"
)

const (
	CacheAccessToken  string = "access_token"
	CacheRefreshToken string = "refresh_token"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func NewClient(username, password string) *Client {
	return &Client{
		Username:   username,
		Password:   password,
		BaseURL:    baseAPIMainURL,
		UserAgent:  "Sanbod/golang",
		HTTPClient: http.DefaultClient,
		Logger:     log.New(os.Stderr, "Sanbod-golang ", log.LstdFlags),
		cache:      newCache(),
	}
}

func NewProxyClient(username, password, proxyUrl string) *Client {
	proxy, err := url.Parse(proxyUrl)
	if err != nil {
		return nil
	}
	tr := &http.Transport{
		Proxy:           http.ProxyURL(proxy),
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return &Client{
		Username:  username,
		Password:  password,
		BaseURL:   baseAPIMainURL,
		UserAgent: "Sanbod/golang",
		HTTPClient: &http.Client{
			Transport: tr,
		},
		Logger: log.New(os.Stderr, "Sanbod-golang ", log.LstdFlags),
		cache:  newCache(),
	}
}

type doFunc func(req *http.Request) (*http.Response, error)

type Client struct {
	Username   string
	Password   string
	BaseURL    string
	UserAgent  string
	HTTPClient *http.Client
	Debug      bool
	Logger     *log.Logger
	TimeOffset int64
	do         doFunc
	cache      *cache
}

func (c *Client) debug(format string, v ...interface{}) {
	if c.Debug {
		c.Logger.Printf(format, v...)
		c.Logger.Println(strings.Repeat("-", 50))
	}
}

func (c *Client) parseRequest(r *request, opts ...RequestOption) (err error) {
	for _, opt := range opts {
		opt(r)
	}
	err = r.validate()
	if err != nil {
		return err
	}

	r.query.Set("traceid", uuid.New().String())
	fullURL := fmt.Sprintf("%s%s", c.BaseURL, r.endpoint)
	header := http.Header{}
	if r.header != nil {
		header = r.header.Clone()
	}
	body := &bytes.Buffer{}
	if r.secType == secTypeRefreshToken {
		refreshToken, _ := c.cache.get(CacheRefreshToken)
		r.setJsonParams(params{
			"grant_type":    "refresh_token",
			"refresh_token": refreshToken,
		})
	} else if r.secType == secTypeAccessToken {
		_, ok := c.cache.get(CacheAccessToken)
		if !ok {
			c.getAuth()
		}
		a, _ := c.cache.get(CacheAccessToken)
		header.Add("Authorization", "Bearer "+a)
	}

	queryString := r.query.Encode()
	bodyString := r.form.Encode()
	if bodyString != "" {
		header.Set("Content-Type", "application/x-www-form-urlencoded")
		body = bytes.NewBufferString(bodyString)
	}
	if r.json != nil {
		header.Set("Content-Type", "application/json")
		body = bytes.NewBuffer(r.json)
	}
	if queryString != "" {
		fullURL = fmt.Sprintf("%s?%s", fullURL, queryString)
	}

	c.debug("full url: %s, body: %s", fullURL, bodyString)
	r.fullURL = fullURL
	r.header = header
	r.body = body
	return nil
}

func (c *Client) getAuth() {
	newClient := NewClient(c.Username, c.Password)

	res, err := newClient.NewCCTokenService().
		Scope([]string{
			"mobilenationalid",
			"cardnationalid",
			"personalinquiry",
			"citizenshipverification",
			"cardtoiban",
			"personal",
		}).ProviderCode("999").
		Do(context.Background())
	if err != nil {
		return
	}

	c.cache.set(CacheAccessToken, res.AccessToken)
	c.cache.set(CacheRefreshToken, res.RefreshToken)
}

func (c *Client) callAPI(ctx context.Context, r *request, opts ...RequestOption) (data []byte, err error) {

	err = c.parseRequest(r, opts...)
	if err != nil {
		return []byte{}, err
	}

	req, err := http.NewRequest(r.method, r.fullURL, r.body)
	if err != nil {
		return []byte{}, err
	}

	req = req.WithContext(ctx)
	req.Header = r.header
	c.debug("request: %#v", req)

	if r.secType == secTypeBasicAuth || r.secType == secTypeRefreshToken {
		req.SetBasicAuth(c.Username, c.Password)
	}

	f := c.do
	if f == nil {
		f = c.HTTPClient.Do
	}

	res, err := f(req)
	if err != nil {
		return []byte{}, err
	}

	data, err = io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}

	defer func() {
		cerr := res.Body.Close()
		if err == nil && cerr != nil {
			err = cerr
		}
	}()

	c.debug("response: %#v", res)
	c.debug("response body: %s", string(data))
	c.debug("response status code: %d", res.StatusCode)

	if res.StatusCode >= http.StatusBadRequest {

		apiErr := new(APIError)
		e := json.Unmarshal(data, apiErr)
		if e != nil {
			c.debug("failed to unmarshal json: %s", e)
		}
		return nil, apiErr
	}

	return data, nil
}

func (c *Client) SetApiEndpoint(url string) *Client {
	c.BaseURL = url
	return c
}
func (c *Client) NewACTokenService() *ACTokenService {
	return &ACTokenService{c: c}
}
func (c *Client) NewCCTokenService() *CCTokenService {
	return &CCTokenService{c: c}
}
func (c *Client) NewRefreshTokenService() *RefreshTokenService {
	return &RefreshTokenService{c: c}
}
func (c *Client) NewMatchNationalCodeWithMobileNumberService() *MatchNationalCodeWithMobileNumberService {
	return &MatchNationalCodeWithMobileNumberService{c: c}
}
func (c *Client) NewMatchNationalCodeWithCardNumberService() *MatchNationalCodeWithCardNumberService {
	return &MatchNationalCodeWithCardNumberService{c: c}
}
func (c *Client) NewInquiryUserProfileWithImageService() *InquiryUserProfileWithImageService {
	return &InquiryUserProfileWithImageService{c: c}
}
func (c *Client) NewInquiryUserProfileService() *InquiryUserProfileService {
	return &InquiryUserProfileService{c: c}
}

const (
	CentralBankOfTheIslamicRepublicOfIran string = "MARKAZI"
	BankOfIndustryMine                    string = "SANAT_VA_MADAN"
	BankMellat                            string = "MELLAT"
	RefahKBank                            string = "REFAH"
	BankMaskan                            string = "MASKAN"
	BankSepah                             string = "SEPAH"
	BankKeshavarziIran                    string = "KESHAVARZI"
	BankMelliIran                         string = "MELLI"
	TejaratBank                           string = "TEJARAT"
	BankSaderatIran                       string = "SADERAT"
	ExportDevelopmentBankOfIran           string = "TOSEAH_SADERAT"
	PostBankIran                          string = "POST"
	ToseeTaavonBank                       string = "TOSEAH_TAAVON"
	KarafarinBank                         string = "KARAFARIN"
	ParsianBank                           string = "PARSIAN"
	EghtesadNovinBank                     string = "EGHTESAD_NOVIN"
	SamanBank                             string = "SAMAN"
	PasargadBank                          string = "PASARGAD"
	SarmayehBank                          string = "SARMAYEH"
	SinaBank                              string = "SINA"
	GharzolhasaneMehrIranBank             string = "MEHR_IRAN"
	ShahrBank                             string = "SHAHR"
	AyandehBank                           string = "AYANDEH"
	TourismBank                           string = "GARDESHGARI"
	DayBank                               string = "DAY"
	IranZaminBank                         string = "IRANZAMIN"
	ResalatGharzolhasaneBank              string = "RESALAT"
	MelalCreditInstitution                string = "MELAL"
	MiddleEastBank                        string = "KHAVARMIANEH"
	NoorCreditInstitution                 string = "NOOR"
	IranVenezuelaBiNationalBank           string = "IRAN_VENEZUELA"
	UNKNOWN                               string = "UNKNOWN"
)
