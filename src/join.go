package src

import (
	"encoding/json"
	"fmt"
	http "github.com/Danny-Dasilva/fhttp"
	"io/ioutil"
	"log"
	"strings"
)

func (in *Instance) Join(Token string, invite string, C string, cap string, captoken, session string) {
	var payload map[string]string
	//var capcount string
	if len(cap) > 2 {
		payload = map[string]string{
			"captcha_key":     cap,
			"captcha_rqtoken": captoken,
			"session_id":      session,
		}
	} else {
		payload = map[string]string{"session_id": session}
	}

	req, err := http.NewRequest("POST",
		"https://discord.com/api/v9/invites/"+invite+"",
		in.Marsh(
			payload,
		),
	)
	if err != nil {
		log.Println(err)
	}

	Hd.Discord(req, map[string]string{
		"X-Debug-Options":    "bugReporterEnabled",
		"Content-Type":       "application/json",
		"X-Discord-Locale":   "en-US",
		"X-Super-Properties": in.Xprop(),
		"Cookie":             C,
		"Referer":            "https://discord.com/",
		"Authorization":      Token,
	})
	resp, err := in.Client.Do(req)
	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()
	var data JoinResp
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &data)

	if err != nil {
		log.Println(err)
	}

	if resp.StatusCode == 200 {
		fmt.Println(""+g+"▏"+r+"("+g+"+"+r+") Joined "+clr+"discord.gg/"+invite, r)
	} else if resp.StatusCode == 429 {
		fmt.Println(""+red+"▏"+r+"("+red+"+"+r+") Failed To Join "+clr+"discord.gg/"+invite, yellow+" RateLimit", r)
	} else if resp.StatusCode == 403 {
		fmt.Println(""+red+"▏"+r+"("+red+"+"+r+") Failed To Join "+clr+"discord.gg/"+invite, yellow+" Locked "+clr+"("+bg+strings.Split(Token, ".")[0]+rb+clr+")", r)
	} else if strings.Contains(string(body), "captcha_sitekey") {
		//if ccount >= 1 {
		//	capcount = clr + "[" + bg + strconv.Itoa(ccount) + rb + clr + "]" + r
		//}
		//fmt.Println(""+yel+"▏"+r+"("+yel+"+"+r+") Solving Captcha "+capcount+clr+"discord.gg/"+invite, r)
		//cap := in.Captcha(data.SiteKey)
		//captoken := data.RqToken
		//ccount++
		//in.Join(Token, invite, C, cap, captoken, session, ccount)
		fmt.Println(""+yellow+"▏"+r+"("+yellow+"+"+r+") Failed To Join "+clr+"discord.gg/"+invite, yellow+" Captcha", r)
	} else {
		errmsg := in.Errmsg(*resp)
		var res string
		if len(errmsg) < 1 {
			res = "Unknown Err"
		} else {
			res = errmsg
		}
		fmt.Println(""+red+"▏"+r+"("+red+"+"+r+") Failed To Join "+clr+"discord.gg/"+invite,
			res,
		)
	}
}
