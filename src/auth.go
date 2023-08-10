package src

import (
	"encoding/json"
	"fmt"
	http "github.com/Danny-Dasilva/fhttp"
	"io/ioutil"
	"net/url"
	"strings"
	"time"
)

func (in *Instance) DiscordIntents(Token string, link string, ID string) StripePayments {
	s := time.Now()
	req, err := http.NewRequest("GET",
		"https://discord.com/api/v9/users/@me/billing/stripe/payment-intents/payments/"+ID,
		nil,
	)
	if err != nil {
		in.StrLog_E("Encountered an Error: ", err.Error(), s)
	}

	Hd.Discord(req, map[string]string{
		"X-Debug-Options":    "bugReporterEnabled",
		"X-Discord-Locale":   "en-US",
		"X-Super-Properties": in.Xprop(),
		"Referer":            "https://discord.com/billing/promotions/" + link,
		"Authorization":      Token,
	})

	resp, err := in.Client.Do(req)
	if err != nil {
		in.StrLog_E("Encountered an Error", err.Error(), s)
	}

	defer resp.Body.Close()
	var data StripePayments
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		in.StrLog_E("Encountered an Error: ", err.Error(), s)
	}

	if err != nil {
		in.StrLog_E("Failed To Get Discord Intents", data.Message, s)
	}

	switch resp.StatusCode {
	default:
		in.StrLog_E("Failed To Get Discord Intents", string(body), s)
	case 200:
		in.StrLog_V("Got Discord Intents", data.StripePaymentIntentClientSecret[:10]+"*** "+r+"| "+gr+data.StripePaymentIntentClientSecretId[:10]+"***", s)
	case 400:
		in.StrLog_E("Failed To Get Discord Intents", data.Message, s)
	}

	return data
}

func (in *Instance) ConfirmDiscordIntents(data StripePayments, StripeKey string) {
	s := time.Now()
	link := fmt.Sprintf("https://api.stripe.com/v1/payment_intents/%v?key=%v&is_stripe_sdk=false&client_secret=%v", strings.Split(data.StripePaymentIntentClientSecret, "_secret_")[0], StripeKey, data.StripePaymentIntentClientSecret)
	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		in.StrLog_E("Encountered an Error", err.Error(), s)
	}
	Hd.StripeHeader(req, nil)
	resp, err := in.Client.Do(req)
	if err != nil {
		in.StrLog_E("Encountered an Error", err.Error(), s)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		in.StrLog_E("Encountered an Error", err.Error(), s)
	}
	switch resp.StatusCode {
	case 200:
		in.StrLog_V("Confirmed Discord Intents", "true", s)
	default:
		in.StrLog_E("Failed To Confirm Discords Intents", string(body), s)
	}

}

func (in *Instance) ConfirmPaymentIntents(StripeKey string, PaymentClientSecret string) (data StripeIntents) {
	s := time.Now()
	payload := fmt.Sprintf(`expected_payment_method_type=card&use_stripe_sdk=true&key=%v&client_secret=%v`, StripeKey, PaymentClientSecret)
	req, err := http.NewRequest("POST",
		"https://api.stripe.com/v1/payment_intents/"+strings.Split(PaymentClientSecret, "_secret_")[0]+"/confirm",
		strings.NewReader(payload),
	)
	if err != nil {
		in.StrLog_E("Encountered an Error", err.Error(), s)
	}

	Hd.StripeHeader(req, nil)
	resp, err := in.Client.Do(req)
	if err != nil {
		in.StrLog_E("Encountered an Error", err.Error(), s)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		in.StrLog_E("Encountered an Error", err.Error(), s)
	}

	switch resp.StatusCode {
	default:
		in.StrLog_E("Failed To Confirm Payment Intents", in.Errmsg(*resp)+string(body), s)
	case 200:
		in.StrLog_V("Confirmed Payment Intents", data.NextAction.UseStripeSdk.ServerTranscation, s)
	case 400:
		in.StrLog_E("Failed To Confirm Payment Intents", string(body), s)
	}

	return data
}

func (in *Instance) Authenticate(StripeKey string, Sdata StripeIntents) {
	s := time.Now()

	Width, Height := in.BrowserScreen()

	prms := url.Values{}
	params := []struct {
		key   string
		value []string
	}{
		{"source", []string{Sdata.NextAction.UseStripeSdk.ThreeDSecure2Source}},
		{"browser", []string{fmt.Sprintf(`{"fingerprintAttempted":false,"fingerprintData":null,"challengeWindowSize":null,"threeDSCompInd":"Y","browserJavaEnabled":false,"browserJavascriptEnabled":true,"browserLanguage":"en-US","browserColorDepth":"24","browserScreenHeight":"%s","browserScreenWidth":"%s","browserTZ":"240","browserUserAgent":"%s"}`, Height, Width, in.Config().Mode.Network.Agent)}},
		{"one_click_authn_device_support[hosted]", []string{"false"}},
		{"one_click_authn_device_support[same_origin_frame]", []string{"false"}},
		{"one_click_authn_device_support[spc_eligible]", []string{"true"}},
		{"one_click_authn_device_support[webauthn_eligible]", []string{"true"}},
		{"one_click_authn_device_support[publickey_credentials_get_allowed]", []string{"true"}},
		{"key", []string{StripeKey}},
	}

	for _, param := range params {
		prms[param.key] = param.value
	}

	req, err := http.NewRequest("POST",
		"https://api.stripe.com/v1/3ds2/authenticate",
		strings.NewReader(prms.Encode()),
	)
	Hd.StripeHeader(req, nil)

	resp, err := in.Client.Do(req)
	if err != nil {
		in.StrLog_E("Encountered an Error", err.Error(), s)
	}

	defer resp.Body.Close()
	var data StripeAuth
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		in.StrLog_E("Encountered an Error", err.Error(), s)
	}

	if data.State != "succeeded" {
		in.StrLog_E("Failed To Authenticate", fmt.Sprint(data.Error), s)
	}

	switch resp.StatusCode {
	default:
		in.StrLog_E("Failed To Authenticate", data.Error.(string), s)
	case 200:
		in.StrLog_V("Authenticated", data.State, s)
	case 400:
		in.StrLog_E("Failed To  Authenticate", data.State+data.Error.(string), s)

	}
}
