package src

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	//	"encoding/json"
	http "github.com/Danny-Dasilva/fhttp"
	//"io/ioutil"
	"log"
	"time"
)

func (in *Instance) GetData(Token string, C string) (TokenResp, time.Time) {
	s := time.Now()
	req, err := http.NewRequest("GET", "https://discord.com/api/v9/users/@me", nil)
	if err != nil {
		log.Println(err)
	}
	Hd.Discord(req, map[string]string{
		"authorization":      Token,
		"content-type":       "application/json",
		"cookie":             C,
		"origin":             "https://discord.com",
		"x-debug-options":    "bugReporterEnabled",
		"x-discord-locale":   "en-US",
		"x-super-properties": in.Xprop(),
	})

	resp, err := in.Client.Do(req)
	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()
	var data TokenResp
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Println(err)
	}

	switch resp.StatusCode {
	case 200:
		return data, s
	case 400:
		in.StrLog_E("Failed to return Token Data", string(body), s)
	}
	return data, s
}

func (in *Instance) NitroToken(Token string, data TokenResp, s time.Time) {
	switch data.PremiumType {
	case 2:
		in.StrLog_V("["+bg+strings.Split(Token, ".")[0]+rb+"]", fmt.Sprintf("%s#%s Nitro: Nitro | ID: %s", data.Username, data.Discriminator, data.Id), s)
		in.WriteFile("checker/valid.txt", Token)
	case 1:
		in.StrLog_V("["+bg+strings.Split(Token, ".")[0]+rb+"]", fmt.Sprintf("%s#%s Nitro: Nitro Classic | ID: %s", data.Username, data.Discriminator, data.Id), s)
		in.WriteFile("checker/valid.txt", Token)
	case 0:
		in.StrLog_E("["+bg+strings.Split(Token, ".")[0]+rb+"]", fmt.Sprintf("%s#%s Nitro: None | ID: %s", data.Username, data.Discriminator, data.Id), s)
		in.WriteFile("checker/failed.txt", Token)
	}
}
