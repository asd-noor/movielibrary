package client

import (
	"errors"
	"net"
	"net/http"
	"time"
	"movielibrary/config"
	"movielibrary/pkg/types"

	"github.com/go-resty/resty/v2"
)

type OmdbClient struct {
	restyClient *resty.Client
	BaseURL     string
	ApiKey      string
}

var omdbc = OmdbClient{}

func NewOmdbClient(conf config.OmdbConfig) OmdbClient {
	timeout := conf.Timeout * time.Second
	if timeout == 0 {
		timeout = 5 * time.Second
	}

	if conf.MaxIdleConnPerHost == 0 {
		conf.MaxIdleConnPerHost = 1
	}

	omdbc = OmdbClient{
		BaseURL: conf.BaseUrl,
		ApiKey:  conf.ApiKey,
		restyClient: resty.NewWithClient(&http.Client{
			Timeout: timeout,
			Transport: &http.Transport{
				DialContext:         (&net.Dialer{Timeout: timeout, KeepAlive: time.Minute}).DialContext,
				TLSHandshakeTimeout: timeout,
				MaxIdleConnsPerHost: conf.MaxIdleConnPerHost,
			},
		}),
	}

	return omdbc
}

func (c OmdbClient) Request(req types.HttpRequest) (*resty.Response, error) {
	req.QueryParams["apikey"] = c.ApiKey

	r := c.restyClient.R()

	for k, v := range req.Headers {
		r.SetHeader(k, v)
	}

	if len(req.QueryParams) > 0 {
		r.SetQueryParams(req.QueryParams)
	}

	if req.Body != nil {
		r.SetBody(req.Body)
	}

	switch req.Method {
	case http.MethodGet:
		return r.Get(c.BaseURL + req.Endpoint)
	case http.MethodPost:
		return r.Post(c.BaseURL + req.Endpoint)
	case http.MethodPut:
		return r.Put(c.BaseURL + req.Endpoint)
	case http.MethodPatch:
		return r.Patch(c.BaseURL + req.Endpoint)
	case http.MethodDelete:
		return r.Delete(c.BaseURL + req.Endpoint)
	default:
		return nil, errors.New("invalid HTTP Method")
	}
}
