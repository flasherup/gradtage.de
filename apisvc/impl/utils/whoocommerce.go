package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/flasherup/gradtage.de/apisvc/config"
	"github.com/flasherup/gradtage.de/common"
	"github.com/tgglv/wc-api-go/client"
	"github.com/tgglv/wc-api-go/options"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Woocommerce struct {
	Key      string
	Secret   string
	WHSecret string
}

func NewWoocommerce(conf config.Woocommerce) *Woocommerce {
	fmt.Println("conf", conf)
	return &Woocommerce{
		Key:      conf.Key,
		Secret:   conf.Secret,
		WHSecret: conf.WHSecret,
	}
}

func GetWoocommerceEventType(headers http.Header) string {
	webhookEvent := headers["X-Wc-Webhook-Event"]
	if len(webhookEvent) > 0 {
		return webhookEvent[0]
	}

	return common.WCUndefinedEvent
}

func GetWoocommerceSignature(headers http.Header) string {
	webhookSignature := headers["X-Wc-Webhook-Signature"]
	if len(webhookSignature) > 0 {
		return webhookSignature[0]
	}

	return ""
}

func ValidateWoocommerceRequest(signature string, body []byte, secret string) bool {
	h := genHMAC256(body, []byte(secret))
	stringHmac := base64.StdEncoding.EncodeToString(h)
	return hmac.Equal([]byte(stringHmac), []byte(signature))
}

func genHMAC256(ciphertext, key []byte) []byte {
	mac := hmac.New(sha256.New, key)
	mac.Write([]byte(ciphertext))
	hmac := mac.Sum(nil)
	return hmac
}

func (wc Woocommerce) GenerateAPIKey(order, email, productId string) (apiKey string, err error) {
	factory := client.Factory{}
	c := factory.NewClient(options.Basic{
		URL:    "https://energy-data.io",
		Key:    wc.Key,
		Secret: wc.Secret,
		Options: options.Advanced{
			WPAPI:       true,
			WPAPIPrefix: "/wp-json/",
			Version:     "wc/v3",
		},
	})

	fmt.Println("order", order, "email", email, "productId", productId)

	parameters := url.Values{}
	parameters.Add("wc-api", "software-api")
	parameters.Add("request", "generate_key")
	parameters.Add("secret_key", "123456789")
	parameters.Add("email", email)
	parameters.Add("product_id", productId)
	parameters.Add("order_id", order)


	if r, err := c.Get("woocommerce", parameters); err != nil {
		return  "", err
	} else if r.StatusCode != http.StatusOK {
		return "", errors.New("unexpected statusCode:" + r.Status)
	} else {
		defer r.Body.Close()
		if bodyBytes, err := ioutil.ReadAll(r.Body); err != nil {
			return  "", err
		} else {
			//{"key":"85c08bd3-fe7f-4174-b294-efed0f1a2e52","key_id":42} Response example
			jsonResponse := struct {
				Key string `json:"key"`
				KeyId string `json:"key"`
			}{}

			if e := json.Unmarshal(bodyBytes, &jsonResponse); e != nil {
				return "", e
			}

			return jsonResponse.Key, nil
		}
	}

	return "", nil
}