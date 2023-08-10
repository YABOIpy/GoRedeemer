package src

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	xhttp "net/http"
)

func (in *Instance) Connect(Token string) (*WsResp, []byte) {
	w := Sock{}
	var dailer websocket.Dialer
	ws, _, err := dailer.Dial("wss://gateway.discord.gg/?v=10&encoding=json", xhttp.Header{
		"Accept-Encoding":          []string{"gzip, deflate, br"},
		"Accept-Language":          []string{"en-US,en;q=0.9"},
		"Cache-Control":            []string{"no-cache"},
		"Pragma":                   []string{"no-cache"},
		"Sec-WebSocket-Extensions": []string{"permessage-deflate; client_max_window_bits"},
		"User-Agent":               []string{"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) discord/1.0.9006 Chrome/91.0.4472.164 Electron/13.6.6 Safari/537.36"},
	})
	if err != nil {
		log.Fatal(err)
	}
	w.Ws = ws

	err = w.Ws.WriteJSON(map[string]interface{}{
		"op": 2,
		"d": map[string]interface{}{
			"token":        Token,
			"capabilities": 125,
			"properties": map[string]interface{}{
				"os":                       "Windows",
				"browser":                  "Vivaldi",
				"system_locale":            "en-US",
				"browser_user_agent":       "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) discord/1.0.9006 Chrome/91.0.4472.164 Electron/13.6.6 Safari/537.36",
				"browser_version":          "94.0",
				"os_version":               "10",
				"referrer":                 "",
				"referring_domain":         "",
				"referrer_current":         "",
				"referring_domain_current": "",
				"release_channel":          "stable",
				"client_build_number":      103981,
			},
			"presence": map[string]interface{}{
				"status":     "online",
				"since":      0,
				"activities": []string{},
				"afk":        false,
			},
			"compress": false,
			"client_state": map[string]interface{}{
				"highest_last_message_id":     "0",
				"read_state_version":          0,
				"user_guild_settings_version": -1,
				"user_settings_version":       -1,
			},
		},
	})
	var data WsResp
	var bd []byte
	if err != nil {
		log.Println(err)
	} else {
		_, b, er := w.Ws.ReadMessage()
		json.Unmarshal(b, &data)
		if er != nil {
			log.Println(err)
		}
		bd = b
	}
	return &data, bd
}
