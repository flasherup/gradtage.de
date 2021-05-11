package utils

import (
	"errors"
	"fmt"
	"github.com/flasherup/gradtage.de/common"
 	"github.com/tgglv/wc-api-go/client"
	"github.com/tgglv/wc-api-go/options"
	"io/ioutil"
	"net/http"
	"net/url"
)

func GetWoocommerceEventType(headers http.Header) string {
	webhookEvent := headers["X-Wc-Webhook-Event"]
	if len(webhookEvent) > 0 {
		return webhookEvent[0]
	}

	return common.WCUndefinedEvent
}

func FinalizeSubscription(order, email, productId string) (apiKey string, err error) {
	factory := client.Factory{}
	c := factory.NewClient(options.Basic{
		URL:    "https://energy-data.io",
		Key:    "ck_df1c6d0cb844d174447034ae29d26091194d1893",
		Secret: "cs_28ea5af092a305d9b2a83697fddf2b0962297a35",
		Options: options.Advanced{
			WPAPI:       true,
			WPAPIPrefix: "/wp-json/",
			Version:     "wc/v3",
		},
	})

	fmt.Println("order", order, "email", email, "productId", productId)

	parameters := url.Values{
		"wc-api":[]string{"software-api"},
		"request":[]string{"request_key"},
		"secret_key":[]string{"iwgcZJ0YEU"},
		"email":[]string{email},
		"product_id":[]string{productId},
	}

	if r, err := c.Get("woocommerce", parameters); err != nil {
		return  "", err
	} else if r.StatusCode != http.StatusOK {
		return "", errors.New("unexpected statusCode:" + r.Status)
	} else {
		defer r.Body.Close()
		if bodyBytes, err := ioutil.ReadAll(r.Body); err != nil {
			return  "", err
		} else {
			fmt.Println("ibody", string(bodyBytes))
		}
	}

	//Complete produc

	data := url.Values{
		"status":[]string{"completed"},
	}

	if r, err := c.Put("order/" + order, data); err != nil {
		return  "", err
	} else if r.StatusCode != http.StatusOK {
		return "", errors.New("unexpected statusCode:" + r.Status)
	} else {
		defer r.Body.Close()
		if bodyBytes, err := ioutil.ReadAll(r.Body); err != nil {
			return  "", err
		} else {
			fmt.Println("pbody", string(bodyBytes))
		}
	}


	return "", nil
}