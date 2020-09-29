package common

import (
	"detection-golang/business/z"
	"log"
	"net/http"
	"time"

	tchttp "github.com/lzkdev/tencentcloud-sdk-go/tencentcloud/common/http"
	"github.com/lzkdev/tencentcloud-sdk-go/tencentcloud/common/profile"
)

type Client struct {
	region      string
	httpClient  *http.Client
	httpProfile *profile.HttpProfile
	credential  *Credential
	signMethod  string
	debug       bool
}

func (c *Client) Send(request tchttp.Request, response tchttp.Response) (err error) {
	if request.GetDomain() == "" {
		domain := c.httpProfile.Endpoint
		if domain == "" {
			domain = tchttp.GetServiceDomain(request.GetService())
		}
		request.SetDomain(domain)
	}
	err = tchttp.ConstructParams(request)
	if err != nil {
		z.Log.InfoWithKey("ConstructParams error", err)
		return
	}
	tchttp.CompleteCommonParams(request, c.GetRegion())
	err = signRequest(request, c.credential, c.signMethod)
	if err != nil {
		z.Log.InfoWithKey("signRequest error", err)
		return
	}
	httpRequest, err := http.NewRequest(request.GetHttpMethod(), request.GetUrl(), request.GetBodyReader())
	if err != nil {
		z.Log.InfoWithKey("NewRequest error", err)
		return
	}
	if request.GetHttpMethod() == "POST" {
		httpRequest.Header["Content-Type"] = []string{"application/x-www-form-urlencoded"}
	}
	log.Printf("[DEBUG] http request=%v", httpRequest.URL.RawQuery)
	httpResponse, err := c.httpClient.Do(httpRequest)
	if err != nil {
		z.Log.InfoWithKey("Do error", err)
		return err
	}
	err = tchttp.ParseFromHttpResponse(httpResponse, response)
	z.Log.InfoWithKey("ParseFromHttpResponse error", err)
	z.Log.InfoWithKey("httpRequest:", httpRequest)
	z.Log.InfoWithKey("httpResponse:", httpResponse)
	return
}

func (c *Client) GetRegion() string {
	return c.region
}

func (c *Client) Init(region string) *Client {
	c.httpClient = &http.Client{}
	c.region = region
	c.signMethod = "HmacSHA256"
	c.debug = false
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	return c
}

func (c *Client) WithSecretId(secretId, secretKey string) *Client {
	c.credential = NewCredential(secretId, secretKey)
	return c
}

func (c *Client) WithProfile(clientProfile *profile.ClientProfile) *Client {
	c.signMethod = clientProfile.SignMethod
	c.httpProfile = clientProfile.HttpProfile
	c.httpClient.Timeout = time.Duration(c.httpProfile.ReqTimeout) * time.Second
	return c
}

func (c *Client) WithSignatureMethod(method string) *Client {
	c.signMethod = method
	return c
}

func NewClientWithSecretId(secretId, secretKey, region string) (client *Client, err error) {
	client = &Client{}
	client.Init(region).WithSecretId(secretId, secretKey)
	return
}
