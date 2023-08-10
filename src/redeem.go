package src

import (
	"encoding/json"
	"fmt"
	http "github.com/Danny-Dasilva/fhttp"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"
)

var (
	Hd = Header{}
)

func (in *Instance) Session(URL string) (StripeKey string) {
	s := time.Now()
	req, err := http.NewRequest("GET", URL, nil)

	if err != nil {
		in.StrLog_E("Encountered an Error: ", err.Error(), s)
	}
	Hd.Discord(req, nil)

	resp, err := in.Client.Do(req)
	if err != nil {
		in.StrLog_E("Encountered an Error: ", err.Error(), s)
	}

	if !ok(resp.StatusCode) {
		in.StrLog_E("Failed To Get Session Data", in.Errmsg(*resp), s)
	} else {

		defer resp.Body.Close()
		body, er := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(er)
		}

		switch resp.StatusCode {
		default:
			in.StrLog_E("Failed", string(body), s)
		case 200:
			key := strings.Split(string(body), `STRIPE_KEY: '`)[1]
			StripeKey = strings.Split(key, `',`)[0]
			in.StrLog_V("Got StripeKey", StripeKey, s)
		case 400:
			in.StrLog_E("Failed To Get StripeKey", string(body), s)
		case 429:
			in.StrLog_R("Ratelimited", "(StripeKey)", s)
		}
	}

	return StripeKey
}

func (in *Instance) StripeID() (Data StripeData) {
	s := time.Now()
	var data = strings.NewReader(`JTdCJTIydjIlMjIlM0ExJTJDJTIyaWQlMjIlM0ElMjJmYjRmNGIwOWEwOWY1YzJlODJiNDU5ZjQwMmMwMDFjMCUyMiUyQyUyMnQlMjIlM0E0MjEuMyUyQyUyMnRhZyUyMiUzQSUyMjQuNS40MiUyMiUyQyUyMnNyYyUyMiUzQSUyMmpzJTIyJTJDJTIyYSUyMiUzQSU3QiUyMmElMjIlM0ElN0IlMjJ2JTIyJTNBJTIyZmFsc2UlMjIlMkMlMjJ0JTIyJTNBMC4zJTdEJTJDJTIyYiUyMiUzQSU3QiUyMnYlMjIlM0ElMjJ0cnVlJTIyJTJDJTIydCUyMiUzQTAlN0QlMkMlMjJjJTIyJTNBJTdCJTIydiUyMiUzQSUyMmVuLUNBJTJDZW4tR0IlMkNlbi1VUyUyQ2VuJTIyJTJDJTIydCUyMiUzQTAlN0QlMkMlMjJkJTIyJTNBJTdCJTIydiUyMiUzQSUyMldpbjMyJTIyJTJDJTIydCUyMiUzQTAuMSU3RCUyQyUyMmUlMjIlM0ElN0IlMjJ2JTIyJTNBJTIyUERGJTIwVmlld2VyJTJDaW50ZXJuYWwtcGRmLXZpZXdlciUyQ2FwcGxpY2F0aW9uJTJGcGRmJTJDcGRmJTJCJTJCdGV4dCUyRnBkZiUyQ3BkZiUyQyUyMENocm9tZSUyMFBERiUyMFZpZXdlciUyQ2ludGVybmFsLXBkZi12aWV3ZXIlMkNhcHBsaWNhdGlvbiUyRnBkZiUyQ3BkZiUyQiUyQnRleHQlMkZwZGYlMkNwZGYlMkMlMjBDaHJvbWl1bSUyMFBERiUyMFZpZXdlciUyQ2ludGVybmFsLXBkZi12aWV3ZXIlMkNhcHBsaWNhdGlvbiUyRnBkZiUyQ3BkZiUyQiUyQnRleHQlMkZwZGYlMkNwZGYlMkMlMjBNaWNyb3NvZnQlMjBFZGdlJTIwUERGJTIwVmlld2VyJTJDaW50ZXJuYWwtcGRmLXZpZXdlciUyQ2FwcGxpY2F0aW9uJTJGcGRmJTJDcGRmJTJCJTJCdGV4dCUyRnBkZiUyQ3BkZiUyQyUyMFdlYktpdCUyMGJ1aWx0LWluJTIwUERGJTJDaW50ZXJuYWwtcGRmLXZpZXdlciUyQ2FwcGxpY2F0aW9uJTJGcGRmJTJDcGRmJTJCJTJCdGV4dCUyRnBkZiUyQ3BkZiUyMiUyQyUyMnQlMjIlM0ExLjklN0QlMkMlMjJmJTIyJTNBJTdCJTIydiUyMiUzQSUyMjE5MjB3XzEwNDBoXzI0ZF8xciUyMiUyQyUyMnQlMjIlM0EwJTdEJTJDJTIyZyUyMiUzQSU3QiUyMnYlMjIlM0ElMjItNCUyMiUyQyUyMnQlMjIlM0EwJTdEJTJDJTIyaCUyMiUzQSU3QiUyMnYlMjIlM0ElMjJmYWxzZSUyMiUyQyUyMnQlMjIlM0EwJTdEJTJDJTIyaSUyMiUzQSU3QiUyMnYlMjIlM0ElMjJzZXNzaW9uU3RvcmFnZS1kaXNhYmxlZCUyQyUyMGxvY2FsU3RvcmFnZS1kaXNhYmxlZCUyMiUyQyUyMnQlMjIlM0EwLjQlN0QlMkMlMjJqJTIyJTNBJTdCJTIydiUyMiUzQSUyMjAxMDAxMDAxMDExMTExMTExMDAxMTExMDExMTExMTExMDExMTAwMTAxMTAxMTExMTAxMTExMTElMjIlMkMlMjJ0JTIyJTNBNDE4LjMlMkMlMjJhdCUyMiUzQTIyNC42JTdEJTJDJTIyayUyMiUzQSU3QiUyMnYlMjIlM0ElMjIlMjIlMkMlMjJ0JTIyJTNBMC4xJTdEJTJDJTIybCUyMiUzQSU3QiUyMnYlMjIlM0ElMjJNb3ppbGxhJTJGNS4wJTIwKFdpbmRvd3MlMjBOVCUyMDEwLjAlM0IlMjBXT1c2NCklMjBBcHBsZVdlYktpdCUyRjUzNy4zNiUyMChLSFRNTCUyQyUyMGxpa2UlMjBHZWNrbyklMjBDaHJvbWUlMkYxMDQuMC4wLjAlMjBTYWZhcmklMkY1MzcuMzYlMjIlMkMlMjJ0JTIyJTNBMCU3RCUyQyUyMm0lMjIlM0ElN0IlMjJ2JTIyJTNBJTIyJTIyJTJDJTIydCUyMiUzQTAuMSU3RCUyQyUyMm4lMjIlM0ElN0IlMjJ2JTIyJTNBJTIyZmFsc2UlMjIlMkMlMjJ0JTIyJTNBMTQzLjIlMkMlMjJhdCUyMiUzQTEuNCU3RCUyQyUyMm8lMjIlM0ElN0IlMjJ2JTIyJTNBJTIyMGJlYTg1MGZmYjliM2FhZGMwZTM4MTFmYjk5NjI4ZjYlMjIlMkMlMjJ0JTIyJTNBNzIuMSU3RCU3RCUyQyUyMmIlMjIlM0ElN0IlMjJhJTIyJTNBJTIyJTIyJTJDJTIyYiUyMiUzQSUyMmh0dHBzJTNBJTJGJTJGR1M3aHFua1pCUnBQXzdXbjktQ0dGaHpxNGtyMlgzSkM0QTNrNkJEQnZwRS5nMnU5LWhxWnZHSXFZSmNQbFBmd0pBZi12M1JneUtfeDFOcHB6QWxBMTJNJTJGQlFkTnl6cExVNG5NNllLenpuYVAxWENEMVcwREoxejNUcG50ejJacElxcyUyRjBzTHkxUFBJSTJobTNPREhpTFp1S2M2UmR5Y0xFMVZybXJ5bnRzWFh0N28lMkZOUGNxbjRDZnRMa2kyX3lBOUVKa0Vwb21sUmxJR1NvT2xuRXRXV3hOODEwJTIyJTJDJTIyYyUyMiUzQSUyMl9JbDFfZzZUOXNyNVdxLXR5SGRlTDFlZUV0ejdPN0lETzFnckMtTlpjVWslMjIlMkMlMjJkJTIyJTNBJTIyTkElMjIlMkMlMjJlJTIyJTNBJTIyTkElMjIlMkMlMjJmJTIyJTNBZmFsc2UlMkMlMjJnJTIyJTNBdHJ1ZSUyQyUyMmglMjIlM0F0cnVlJTJDJTIyaSUyMiUzQSU1QiUyMmxvY2F0aW9uJTIyJTVEJTJDJTIyaiUyMiUzQSU1QiU1RCUyQyUyMm4lMjIlM0E0MzAuMjAwMDAwMDQ3NjgzNyUyQyUyMnUlMjIlM0ElMjJkaXNjb3JkLmNvbSUyMiU3RCUyQyUyMmglMjIlM0ElMjJhNjkzYTE1NmM5Y2I2ZTM4OWMxNyUyMiU3RA==`)
	req, err := http.NewRequest(
		"POST",
		"https://m.stripe.com/6",
		data,
	)
	if err != nil {
		in.StrLog_E("Encountered an Error: ", err.Error(), s)
	}
	Hd.StripeHeader(req, nil)

	sdata := StripeIDs{}

	resp, err := in.Client.Do(req)
	if err != nil {
		in.StrLog_E("Encountered an Error: ", err.Error(), s)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &sdata)
	if err != nil {
		in.StrLog_E("Encountered an Error: ", err.Error(), s)
	}

	switch resp.StatusCode {
	default:
		in.StrLog_E("Failed", string(body), s)
	case 200:
		DataID := StripeData{
			ID: StripeIDs{
				Muid: sdata.Muid,
				Guid: sdata.Guid,
				Sid:  sdata.Sid,
			},
		}

		Data = DataID
		str := Data.ID.Muid[:10] + "** " + r + "| " + gr + Data.ID.Guid[:10] + "** " + r + "| " + gr + Data.ID.Sid[:10] + "**"
		in.StrLog_V("Got Stripe ID's", str, s)
	case 400:
		in.StrLog_E("Failed To Get Stipe ID's", in.Errmsg(*resp), s)
	case 429:
		in.StrLog_R("Ratelimited", "(Stripe ID) "+fmt.Sprint(Data.Retry)+"ms", s)
	}

	return Data
}

func (in *Instance) StripeToken(cc string, sdata StripeData, StripeKey string) (data Confirmation) {
	s := time.Now()
	vcc := strings.Split(cc, ":")
	payload := fmt.Sprintf(`card[number]=%v&card[cvc]=%v&card[exp_month]=%v&card[exp_year]=%v&guid=%v&muid=%v&sid=%v&payment_user_agent=%v&time_on_page=%v&key=%v&pasted_fields=%s`,
		vcc[0],
		vcc[2],
		vcc[1][0:2],
		vcc[1][2:],
		sdata.ID.Guid,
		sdata.ID.Muid,
		sdata.ID.Sid,
		`stripe.js%2Ff2ecd562b%3B+stripe-js-v3%2Ff2ecd562b`,
		rand.Intn(60000)+60000,
		StripeKey,
		"number%2Cexp%2Ccvc",
	)
	req, err := http.NewRequest(
		"POST",
		"https://api.stripe.com/v1/tokens",
		strings.NewReader(payload),
	)
	if err != nil {
		in.StrLog_E("Encountered an Error: ", err.Error(), s)
	}
	Hd.StripeHeader(req, map[string]string{
		"origin":  "https://js.stripe.com",
		"referer": "https://js.stripe.com",
	})

	resp, err := in.Client.Do(req)
	if err != nil {
		in.StrLog_E("Encountered an Error: ", err.Error(), s)
	}

	if resp.StatusCode == 200 {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		err = json.Unmarshal(body, &data)
		if err != nil {
			in.StrLog_E("Encountered an Error: ", err.Error(), s)
		}
		str := `{` + r + `
` + rb + `[` + prp + `Info ` + r + `]	[` + g + `✓` + r + `] CardID: ` + gr + data.Card.Id + r + `
` + rb + `[` + prp + `Info ` + r + `]	[` + g + `✓` + r + `] ConfirmID: ` + gr + data.Id + r + `
` + rb + `[` + prp + `Info ` + r + `]	[` + g + `✓` + r + `] CardBrand: ` + gr + data.Card.Brand + r + `
` + rb + `[` + prp + `Info ` + r + `]	[` + g + `✓` + r + `] CardLast4: ` + gr + data.Card.Last4 + r + ``
		in.StrLog_V("Got Confirmation Data", str, s)
	}

	return data

}

func (in *Instance) StripeConfirm(SecretKey string, StripeKey string, ConfirmToken string, sdata StripeData, ldata BillingAddress) string {

	address := BillingAddress{
		City:     ldata.City,
		Country:  ldata.Country,
		Email:    ldata.Email,
		Name:     ldata.Name,
		Line:     ldata.Line,
		Line2:    ldata.Line2,
		PostCode: ldata.PostCode,
		State:    ldata.State,
	}
	s := time.Now()
	str := strings.Split(SecretKey, "_secret_")[0]
	data := fmt.Sprintf(`payment_method_data[type]=card&payment_method_data[card][token]=%v&payment_method_data[billing_details][address][line1]=%v&payment_method_data[billing_details][address][line2]=%v&payment_method_data[billing_details][address][city]=%v&payment_method_data[billing_details][address][state]=%v&payment_method_data[billing_details][address][postal_code]=%v&payment_method_data[billing_details][address][country]=%v&payment_method_data[billing_details][name]=%v&payment_method_data[guid]=%v&payment_method_data[muid]=%v&payment_method_data[sid]=%v&payment_method_data[payment_user_agent]=%v&payment_method_data[time_on_page]=%v&expected_payment_method_type=card&use_stripe_sdk=true&key=%v&client_secret=%v`,
		ConfirmToken,
		address.Line,
		address.Line2,
		address.City,
		address.State,
		address.PostCode,
		address.Country,
		address.Name,
		sdata.ID.Guid,
		sdata.ID.Muid,
		sdata.ID.Sid,
		`stripe.js%2Ff2ecd562b%3B+stripe-js-v3%2Ff2ecd562b`,
		rand.Intn(450000)+250000,
		StripeKey,
		SecretKey,
	)

	req, err := http.NewRequest(
		"POST",
		"https://api.stripe.com/v1/setup_intents/"+str+"/confirm",
		strings.NewReader(
			strings.ReplaceAll(data, " ", "+"),
		),
	)
	if err != nil {
		in.StrLog_E("Encountered an Error: ", err.Error(), s)
	}

	Hd.StripeHeader(req, map[string]string{
		"origin":  "https://js.stripe.com",
		"referer": "https://js.stripe.com/",
	})

	resp, err := in.Client.Do(req)
	var jdata StripeData
	if err != nil {
		in.StrLog_E("Encountered an Error: ", err.Error(), s)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &jdata.Payment)
	if err != nil {
		in.StrLog_E("Encountered an Error: ", err.Error(), s)
	}

	if !ok(resp.StatusCode) {
		in.StrLog_E("Failed To Setup Stripe Intents", jdata.Payment.Error.Message, s)
	}
	in.StrLog_V("Got Payment ID", jdata.Payment.PaymentId, s)

	return jdata.Payment.PaymentId

}

func (in *Instance) AddPayment(Token string, auth string, AddressToken, C string, link string,
	data BillingAddress, Sec string, key string, ID string, d StripeData) string {

	s := time.Now()
	payload := map[string]interface{}{
		"payment_gateway": 1,
		"token":           auth,
		"billing_address": BillingAddress{
			City:     data.City,
			Country:  data.Country,
			Email:    data.Email,
			Name:     data.Name,
			Line:     data.Line,
			Line2:    data.Line2,
			PostCode: data.PostCode,
			State:    data.State,
		},
		"billing_address_token": AddressToken,
	}
	req, err := http.NewRequest(
		"POST",
		"https://discord.com/api/v9/users/@me/billing/payment-sources",
		in.Marshintf(payload),
	)
	if err != nil {
		in.StrLog_E("Encountered an Error: ", err.Error(), s)
	}
	Hd.Discord(req, map[string]string{
		"X-Debug-Options":    "bugReporterEnabled",
		"Content-Type":       "application/json",
		"X-Discord-Locale":   "en-US",
		"X-Super-Properties": in.Xprop(),
		"Referer":            "https://discord.com/billing/promotions/" + link,
		"Authorization":      Token,
	})

	resp, err := in.Client.Do(req)
	if err != nil {
		in.StrLog_E("Encountered an Error: ", err.Error(), s)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var jsonbody SourceID
	err = json.Unmarshal(body, &jsonbody)
	if err != nil {
		in.StrLog_E("Encountered an Error: ", err.Error(), s)
	}

	switch resp.StatusCode {
	default:
		in.StrLog_E("Failed", in.Errmsg(*resp)+string(body), s)
	case 200:
		in.StrLog_V("Got Payment Source ID", jsonbody.ID, s)
	case 400:
		in.StrLog_E("Failed To Get Payment Source ID", string(body), s)
		if strings.Contains(string(body), "Billing address is invalid") || strings.Contains(string(body), "Payments unavailable in this region") {
			//"BILLING_DUPLICATE_PAYMENT_SOURCE" switch token
			in.StrLog_E("Invalid Address Creating New one", "(Unkown Address)", s)
			Tkn, Adr := in.GetAuth(Token, link, C)
			ath := in.StripeConfirm(Sec, key, ID, d, Adr)
			in.AddPayment(Token, ath, Tkn, C, link, Adr, Sec, key, ID, d)
		}
	case 429:
		in.StrLog_R("Ratelimited", "(Payment Source ID) "+fmt.Sprint(jsonbody.Retry)+"ms", s)
	}

	return jsonbody.ID
}

func (in *Instance) GetSecClient(Token string, link string, C string) (Key string) {
	s := time.Now()
	req, err := http.NewRequest("POST",
		"https://discord.com/api/v9/users/@me/billing/stripe/setup-intents",
		nil,
	)

	if err != nil {
		in.StrLog_E("Encountered an Error: ", err.Error(), s)
	}
	Hd.SecHeader(req, Token, link, C)

	resp, err := in.Client.Do(req)
	if err != nil {
		in.StrLog_E("Encountered an Error: ", err.Error(), s)
	}

	defer resp.Body.Close()
	body, err := in.ReadBody(*resp)
	if err != nil {
		in.StrLog_E("Encountered an Error: ", err.Error(), s)
	}

	var data ValidateResp

	json.Unmarshal(body, &data)
	if err != nil {
		in.StrLog_E("Encountered an Error: ", err.Error(), s)
	}

	switch resp.StatusCode {
	default:
		in.StrLog_E("Failed", in.Errmsg(*resp)+string(body), s)
	case 200:
		in.StrLog_V("Got Client Key", data.Secret[:30]+"**", s)
		Key = data.Secret
	case 400:
		in.StrLog_E("Failed To Get Client Key", string(body), s)
	case 429:
		in.StrLog_R("Ratelimited", "(Client Key) "+fmt.Sprint(data.Retry)+"ms", s)
	}

	return Key
}

func (in *Instance) GetAuth(Token string, link string, C string) (string, BillingAddress) {

	s := time.Now()
	data := in.LocationData()
	req, err := http.NewRequest("POST",
		"https://discord.com/api/v9/users/@me/billing/payment-sources/validate-billing-address",
		in.Marshany(PaymentResp{
			Address: BillingAddress{
				City:     data.City,
				Country:  data.Country,
				Email:    data.Email,
				Name:     data.Name,
				Line:     data.Line,
				Line2:    data.Line2,
				PostCode: data.PostCode,
				State:    data.State,
			},
		}),
	)

	if err != nil {
		in.StrLog_E("Encountered an Error: ", err.Error(), s)
	}
	Hd.MainHeader(req, Token, C, link)

	resp, err := in.Client.Do(req)
	if err != nil {
		in.StrLog_E("Encountered an Error: ", err.Error(), s)
	}

	defer resp.Body.Close()
	var D ValidateResp
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		in.StrLog_E("Encountered an Error: ", err.Error(), s)
	}
	json.Unmarshal(body, &D)

	switch resp.StatusCode {
	default:
		in.StrLog_E("Error", string(body), s)
	case 200:
		in.StrLog_V("Got Auth Token", D.Token, s)
	case 400:
		if strings.Contains(string(body), "Billing address is invalid") {
			in.StrLog_E("Invalid Address Creating New one", "(Unkown Address)", s)
			in.GetAuth(Token, link, C)
		} else {
			in.StrLog_E("Invalid Request", string(body), s)
		}
	case 429:
		in.StrLog_R("RateLimit", "(Auth) "+fmt.Sprint(D.Retry)+"ms", s)
	}

	return D.Token, data
}

func (in *Instance) Claim(Token string, link string, ID string) (d bool, paymentID string) {
	s := time.Now()
	req, err := http.NewRequest("POST",
		"https://discord.com/api/v9/entitlements/gift-codes/"+link+"/redeem",
		in.Marshintf(map[string]interface{}{
			"channel_id":        nil,
			"payment_source_id": ID,
		}),
	)
	if err != nil {
		in.StrLog_E("Encountered an Error: ", err.Error(), s)
	}

	Hd.Discord(req, map[string]string{
		"X-Debug-Options":    "bugReporterEnabled",
		"Content-Type":       "application/json",
		"X-Discord-Locale":   "en-US",
		"X-Super-Properties": in.Xprop(),
		"Host":               "discord.com",
		"Referer":            "https://discord.com/billing/promotions/" + link,
		"Authorization":      Token,
	})

	resp, err := in.Client.Do(req)
	if err != nil {
		in.StrLog_E("Encountered an Error: ", err.Error(), s)
	}

	defer resp.Body.Close()
	var data RedeemResp
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		in.StrLog_E("Encountered an Error: ", err.Error(), s)
	}
	d = true
	switch resp.StatusCode {
	default:
		in.StrLog_E("Failed To Redeem", in.Errmsg(*resp)+string(body), s)
	case 200:
		in.StrLog_V("["+bg+strings.Split(Token, ".")[0]+rb+"]"+" Redeemed Nitro", in.Cfg.Mode.Configs.PromoType+link, s)
		in.WriteFile("claimed.txt", fmt.Sprintf("%s:%s:%s", in.Email, in.Pass, Token))
		d = false
	case 400:
		var output string
		if strings.Contains(data.Message, "Authentication") {
			output = "Bypassing 3DSec"
		} else {
			output = string(body)[:40]
		}
		in.StrLog_E("["+bg+strings.Split(Token, ".")[0]+rb+"]"+" Failed To Redeem", output, s)
	case 403:
		in.StrLog_E("["+bg+strings.Split(Token, ".")[0]+rb+"]"+" Failed To Redeem", "Locked Token", s)
	case 429:
		in.StrLog_R("Ratelimited", "(Client Key) "+fmt.Sprint(data.Retry)+"ms", s)
	}

	return d, data.PaymentId
}
