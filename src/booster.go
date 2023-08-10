package src

import (
	"encoding/json"
	http "github.com/Danny-Dasilva/fhttp"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

func (in *Instance) Boost(Token string, C string, ID string, b bool) {

	s := time.Now()
	req, err := http.NewRequest("GET",
		"https://discord.com/api/v9/users/@me/guilds/premium/subscription-slots",
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}
	Hd.Discord(req, map[string]string{
		"X-Debug-Options":    "bugReporterEnabled",
		"Content-Type":       "application/json",
		"X-Discord-Locale":   "en-US",
		"X-Super-Properties": in.Xprop(),
		"cookie":             C,
		"Referer":            "https://discord.com/channels/",
		"Authorization":      Token,
	})
	resp, err := in.Client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	var data []BoostResp
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Println(err)
	}

	tkn := strings.Split(Token, ".")[0]
	switch resp.StatusCode {
	default:
		str := red + "[" + r + tkn + red + "] " + "Failed To Boost" + r + ":"
		in.StrLog_E(str, ID, s)

	case 200:
		for _, d := range data {
			var ar, slotID []string
			slotID = append(ar, d.Id)
			re, er := http.NewRequest("PUT",
				"https://discord.com/api/v9/guilds/"+ID+"/premium/subscriptions",
				in.Marshany(
					BoostPayload{
						UserPremiumGuildSubscriptionSlotIds: slotID,
					},
				),
			)
			if er != nil {
				log.Println(er)
			}

			Hd.MainHeader(re, Token, C, "")
			res, er := in.Client.Do(re)
			if er != nil {
				log.Fatal(er)
			}
			defer res.Body.Close()

			switch res.StatusCode {
			case 201:
				str := g + "[" + r + tkn + g + "] " + "Boosted" + r
				in.StrLog_V(str, ID, s)
			default:
				str := red + "[" + r + tkn + red + "] " + "Failed To Boost" + r
				in.StrLog_E(str, ID, s)
			}
		}
	}
}
