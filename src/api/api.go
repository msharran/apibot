package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/kyokomi/emoji"
	"github.com/sharran-murali/apibot/src/config"
	"github.com/sharran-murali/apibot/src/utils"
)

type Client struct {
	restyClient *resty.Client
}

func NewClient() Client {
	client := resty.New()
	client.EnableTrace()
	return Client{
		restyClient: client,
	}
}

func (c *Client) SetHeaders(headers map[string]string) {
	c.restyClient.SetHeaders(headers)
}

func (c *Client) Request(profileName string) *resty.Request {
	if utils.IsFileExist(config.GetConfigFile()) {
		profile := config.GetProfile(profileName)
		c.restyClient.SetHostURL(profile.BaseUrl)
		c.restyClient.SetHeader("Authorization", profile.AuthorizationHeader)
	}
	return c.restyClient.R()
}

func (c *Client) Println(resp *resty.Response) {
	fmt.Println()
	fmt.Println("--------------------------------------------------------------------------")
	utils.LogInfoln("Request")
	fmt.Println("--------------------------------------------------------------------------")
	fmt.Println("URL\t\t:", resp.Request.URL)
	fmt.Println("Time\t\t: ", resp.Time().String())
	fmt.Printf("Trace\t\t: %+v\n", resp.Request.TraceInfo())
	fmt.Println("--------------------------------------------------------------------------")
	utils.LogInfoln("Request headers")
	fmt.Println("--------------------------------------------------------------------------")
	fmt.Print(printableHeaders(resp.Request.Header))
	fmt.Println("--------------------------------------------------------------------------")
	utils.LogInfo("Response ")
	if resp.IsSuccess() {
		utils.LogSuccessln(emoji.Sprint(":check_mark_button:") + "[" + resp.Status() + "]")
	} else {
		utils.LogErrorln(emoji.Sprint(":cross_mark:") + "[" + resp.Status() + "]")
	}
	fmt.Println("--------------------------------------------------------------------------")
	fmt.Println(resp)
	fmt.Println("--------------------------------------------------------------------------")
}

func printableHeaders(header http.Header) (headerStr string) {
	for k, v := range header {
		headerStr += k + "\t\t: " + strings.Join(v, ",") + "\n"
	}
	return
}
