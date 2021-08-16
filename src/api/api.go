package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/kyokomi/emoji"
	"github.com/sharran-murali/apibot/src/botfactory"
	"github.com/sharran-murali/apibot/src/config"
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
	if botfactory.IsFileExist(config.GetConfigFile()) {
		profile := config.GetProfile(profileName)
		c.restyClient.SetHostURL(profile.BaseUrl)
		c.restyClient.SetHeaders(profile.Headers)
	}
	return c.restyClient.R()
}

func (c *Client) Println(resp *resty.Response) {
	fmt.Println()
	fmt.Println("---------------------------------------------------------------------------------------------------")
	botfactory.LogInfo("Request ")
	botfactory.LogSuccess(fmt.Sprintf("[%v] ", resp.Request.Method))
	fmt.Println("=>", resp.Request.URL)
	fmt.Println("---------------------------------------------------------------------------------------------------")
	fmt.Println("Time\t\t: ", resp.Time().String())
	fmt.Printf("Trace\t\t: %+v\n", resp.Request.TraceInfo())
	fmt.Println("---------------------------------------------------------------------------------------------------")
	if len(resp.Request.Header) > 0 {
		botfactory.LogBlueln("Headers")
		fmt.Println("--------")
		fmt.Print(printableHeaders(resp.Request.Header))
	}
	if resp.Request.Body != "" {
		fmt.Println("--------")
		botfactory.LogBlueln("Body")
		fmt.Println("--------")
		fmt.Println(resp.Request.Body)
	}
	fmt.Println("---------------------------------------------------------------------------------------------------")
	botfactory.LogInfo("Response ")
	if resp.IsSuccess() {
		botfactory.LogSuccessln(emoji.Sprint(":check_mark_button:") + "[" + resp.Status() + "]")
	} else {
		botfactory.LogErrorln(emoji.Sprint(":cross_mark:") + "[" + resp.Status() + "]")
	}
	fmt.Println("---------------------------------------------------------------------------------------------------")
	fmt.Println(resp)
	fmt.Println("---------------------------------------------------------------------------------------------------")
}

func printableHeaders(header http.Header) (headerStr string) {
	for k, v := range header {
		headerStr += k + "\t\t: " + strings.Join(v, ",") + "\n"
	}
	return
}
