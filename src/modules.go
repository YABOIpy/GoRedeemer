package src

import (
	"bufio"
	"bytes"
	"compress/zlib"
	"crypto/md5"
	"crypto/tls"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	http "github.com/Danny-Dasilva/fhttp"
	"github.com/andybalholm/brotli"
	"github.com/gorilla/websocket"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	shttp "net/http"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func (in *Instance) Err(err error) {
	if err != nil {
		fmt.Errorf("", err)
	}
}

func (in *Instance) Marsh(data map[string]string) *bytes.Buffer {
	d, err := json.Marshal(data)
	in.Err(err)
	byt := bytes.NewBuffer(d)

	return byt
}

func (in *Instance) Marshany(data any) *bytes.Buffer {
	d, err := json.Marshal(data)
	in.Err(err)
	byt := bytes.NewBuffer(d)

	return byt
}

func (in *Instance) Marshintf(data map[string]interface{}) *bytes.Buffer {
	d, err := json.Marshal(data)
	in.Err(err)
	byt := bytes.NewBuffer(d)

	return byt
}

func ok(code int) bool {
	for _, v := range [3]int{200, 201, 204} {
		if code == v {
			return true
		}
	}
	return false
}

var (
	buildNumber = make(map[string]string)
)

func (in *Instance) BuildInfo() string {

	js := regexp.MustCompile(`([a-zA-z0-9]+)\.js`)
	buildInfo := regexp.MustCompile(`Build Number: "\)\.concat\("([0-9]{4,8})"`)

	client := &shttp.Client{Timeout: 10 * time.Second}

	res, err := client.Get("https://discord.com/app")
	in.Err(err)

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	in.Err(err)

	rs := js.FindAllString(string(body), -1)
	asset := rs[len(rs)-1]
	if strings.Contains(asset, "invisible") {
		asset = rs[len(rs)-2]
	}

	resp, err := client.Get("https://discord.com/assets/" + asset)
	in.Err(err)

	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	in.Err(err)

	z := buildInfo.FindAllString(string(b), -1)
	e := strings.ReplaceAll(z[0], " ", "")
	buildInfos := strings.Split(e, ",")

	buildNum := strings.Split(buildInfos[0], `("`)
	buildNumber["stable"] = strings.ReplaceAll(buildNum[len(buildNum)-1], `"`, ``)

	return buildNumber["stable"]
}

func (in *Instance) Xprop() string {
	rand.Seed(time.Now().UnixNano())
	browser := in.Browser()
	data := in.getRandomBrowser(browser)

	OS := in.getRandomKey(data.OSver)
	cbn, _ := strconv.Atoi(in.BuildInfo())

	d, err := json.Marshal(XpropData{
		OS:                 OS,
		Browser:            data.Name,
		Device:             "",
		SystemLocale:       "en-US",
		BrowserUserAgent:   in.Cfg.Mode.Network.Agent,
		BrowserVersion:     in.getRandomData(data.Versions),
		OSVersion:          in.getRandomData(data.OSver[OS]),
		Referrer:           "https://www.google.com",
		ReferringDomain:    "www.google.com",
		SearchEngine:       "google",
		ReferrerCurrent:    "",
		ReferringDomainCur: "",
		ReleaseChannel:     "stable",
		ClientBuildNumber:  cbn,
		ClientEventSource:  nil,
	})
	if err != nil {
		log.Println(err)
	}
	return base64.StdEncoding.EncodeToString(d)
}

func (in *Instance) Logo() ([]byte, error) {
	logo := `CgkgIF9fX19fX19fXyAgIF9fXyAgX19fX19fXyAgX19fX19fX19fXyAgX19fX19fX19fXyAKCSAvIF9fXy8gX18gXCAvIF8gXC8gX18vIF8gXC8gX18vIF9fLyAgfC8gIC8gX18vIF8gXAoJLyAoXyAvIC9fLyAvICAsIF8vIF8vLyAvLyAvIF8vLyBfLy8gL3xfLyAvIF8vLyAsIF8vCglcX19fL1xfX19fLyBfL3xfL19fXy9fX19fL19fXy9fX18vXy8gIC9fL19fXy9fL3xffAogICAgICAgICAgICAgICAgICAgIAkJCWdpdGh1Yi5jb20veWFib2lweSAgICAg`
	return base64.StdEncoding.DecodeString(logo)
}

func (in *Instance) Menu() {
	Vcc, _, _ := in.ReadFile("vcc.txt")
	Token, _, _ := in.ReadFile("tokens.txt")
	Promo, _, _ := in.ReadFile("promos.txt")
	fmt.Print(bg + `	(` + rb + `[` + bg + strconv.Itoa(len(Vcc)) + rb + `: Vcc] [` + bg + strconv.Itoa(len(Token)) + rb + `: Tokens] [` + bg + strconv.Itoa(len(Promo)) + rb + `: Promo]` + bg + `)` + rb + `
	` + bg + `____________________________________` + rb + `
	[` + bg + `1` + rb + `] Redeemer [` + bg + `2` + rb + `] Booster [` + bg + `3` + rb + `] Joiner

	Choice >:`)
}

func (in *Instance) ReadFile(files string) ([]string, []string, error) {

	file, err := os.Open(files)
	in.Err(err)
	defer file.Close()
	var value bool
	var lines, tokens []string
	var ept []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if strings.Contains(files, "token") {
		for i := 0; i < len(lines); i++ {
			if strings.Contains(lines[i], ":") {
				format := strings.Split(lines[i], ":")
				tokens = append(tokens, format[2])
				ept = append(ept, lines[i])
				value = true
			}
		}
		if value {
			lines = nil
			lines = tokens
		}
	}

	return lines, ept, scanner.Err()
}

func (in *Instance) cookies() (Cookies string) {

	cookie := []*shttp.Cookie{}
	req, err := shttp.NewRequest("GET", "https://discord.com", nil)
	if err != nil {
		log.Println(err)
	}

	client := shttp.DefaultClient

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()

	if resp.Cookies() == nil {
		in.cookies()
	}

	cookie = append(cookie, resp.Cookies()...)
	for i := 0; i < len(cookie); i++ {
		if i == len(cookie)-1 {
			Cookies += fmt.Sprintf(`%s=%s`,
				cookie[i].Name,
				cookie[i].Value,
			)
		} else {
			Cookies += fmt.Sprintf(`%s=%s; `,
				cookie[i].Name,
				cookie[i].Value,
			)
		}
	}
	if !strings.Contains(Cookies, "locale=en-US; ") {
		Cookies += "; locale=en-US "
	}

	return Cookies
}

func (in *Instance) GetCookie() string {
	cook := in.cookies()
	return cook
}

func (in *Instance) ReadBody(resp http.Response) ([]byte, error) {

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.Header.Get("Content-Encoding") == "gzip" {
		gzipreader, err := zlib.NewReader(bytes.NewReader(body))
		if err != nil {
			return nil, err
		}
		gzipbody, err := ioutil.ReadAll(gzipreader)
		if err != nil {
			return nil, err
		}
		return gzipbody, nil
	}

	if resp.Header.Get("Content-Encoding") == "br" {
		brreader := brotli.NewReader(bytes.NewReader(body))
		brbody, err := ioutil.ReadAll(brreader)
		if err != nil {
			fmt.Println(string(brbody))
			return nil, err
		}

		return brbody, nil
	}
	return body, nil
}

func (in *Instance) Errmsg(response http.Response) (errmsg string) {
	body, _ := in.ReadBody(response)

	if in.Config().Mode.Configs.Errormsg {
		errmsg = string(body)
	} else {
		return ""
	}

	return errmsg

}

func (in *Instance) Config() Config {
	var config Config
	conf, err := os.Open("config.json")
	defer func(conf *os.File) {
		err := conf.Close()
		if err != nil {

		}
	}(conf)
	if err != nil {
		log.Fatal(err)
	}
	xp := json.NewDecoder(conf)
	errs := xp.Decode(&config)
	if errs != nil {
		return Config{}
	}
	return config
}

func (in *Instance) CreateJa3() (string, string) {
	JA3, Md5Hash := in.JA3String(in.ClientHello())
	return JA3, Md5Hash
}

func (in *Instance) ClientHello() ([]uint16, string, []uint16, uint16) {
	rand.Seed(time.Now().UnixNano())
	Cip := in.Ciphers()
	server := in.Config().Mode.Network.Server
	Ciphers := in.getRandomMapSlice(Cip[0].CipherList)
	ECurves := in.getRandomMapSlice(Cip[0].EllipticCurves)

	CurveList := make([]tls.CurveID, len(ECurves))
	for i, v := range ECurves {
		CurveList[i] = tls.CurveID(v)
	}
	tlsConfig := &tls.Config{
		ServerName:         server[0],
		MaxVersion:         tls.VersionTLS13,
		MinVersion:         tls.VersionTLS12,
		CipherSuites:       Ciphers,
		CurvePreferences:   CurveList,
		InsecureSkipVerify: true,
	}

	conn, err := tls.Dial(
		"tcp",
		server[0]+":"+server[1],
		tlsConfig,
	)
	if err != nil {
		log.Println(err)
	}

	defer conn.Close()
	ver := conn.ConnectionState().Version
	ext := conn.ConnectionState().NegotiatedProtocol
	Suite := tlsConfig.CipherSuites
	EC := tlsConfig.CurvePreferences

	Curves := make([]uint16, len(EC))
	for i, curve := range EC {
		Curves[i] = uint16(curve)
	}

	return Suite, ext, Curves, ver
}

func (in *Instance) JA3String(cipherSuites []uint16, extensions string, EllipticCurves []uint16, ver uint16) (Ja3 string, Ja3md5hash string) {
	cipherSuiteStr := ""
	ECstr := ""
	ExtStr := ""
	for _, suite := range cipherSuites {
		cipherSuiteStr += fmt.Sprintf("%d-", suite)
	}

	for _, curve := range EllipticCurves {
		ECstr += fmt.Sprintf("%d-", curve)
	}

	cipherSuiteStr = strings.TrimSuffix(cipherSuiteStr, "-")

	if extensions != "" {
		exts := strings.Split(extensions, "-")
		for _, ext := range exts {
			ExtStr += fmt.Sprintf("%d-", ext)
		}
		ExtStr = strings.TrimSuffix(ExtStr, "-")
	}
	ECstr = strings.TrimSuffix(ECstr, "-")

	Ja3 = fmt.Sprintf("%d,%s,%s%s,0-1-2", ver, cipherSuiteStr, ECstr, ExtStr)
	hash := md5.Sum([]byte(Ja3))
	Ja3md5hash = hex.EncodeToString(hash[:])
	return Ja3, Ja3md5hash
}

func (in *Instance) BrowserScreen() (string, string) {
	resolution := []string{
		"1920x1080", "2560x1440",
		"3840x2160", "1280x800",
		"1366x768", "1680x1050",
		"1440x900", "1600x900",
		"2560x1600", "1920x1200",
		"1024x768", "1280x1024",
		"2560x1080", "3440x1440",
		"5120x1440", "1280x720",
		"2560x1280", "4096x2160",
		"3840x1080", "5120x2160",
	}
	res := strings.Split(
		resolution[rand.Intn(
			len(resolution),
		)], "x",
	)

	return res[0], res[1]
}

func (in *Instance) Browser() (browsers []BrowserData) {
	browsers = []BrowserData{
		{
			Name:     "Google Chrome",
			Versions: []string{"94", "93", "92", "91", "90"},
			OSver: map[string][]string{
				"Windows": {"10", "8.1", "8", "7"},
				"Mac":     {"11", "10.15", "10.14", "10.13"},
				"Linux":   {"Ubuntu 20", "Debian 10", "Fedora 34"},
			},
		},
		{
			Name:     "Mozilla Firefox",
			Versions: []string{"93", "92", "91", "90", "89"},
			OSver: map[string][]string{
				"Windows": {"10", "8.1", "8", "7"},
				"Mac":     {"11", "10.15", "10.14", "10.13"},
				"Linux":   {"Ubuntu 20", "Debian 10", "Fedora 34"},
			},
			UserAgent: map[string]string{
				"Windows": "Mozilla/5.0 (Windows NT %s) Gecko/20100101 Firefox/%s",
				"Mac":     "Mozilla/5.0 (Macintosh; Intel Mac OS X %s) Gecko/20100101 Firefox/%s",
				"Linux":   "Mozilla/5.0 (X11; %s; Linux x86_ %s) Gecko/20100101 Firefox/%s", //cut version string 3 short for linux
			},
		},
		{
			Name:     "Safari",
			Versions: []string{"15", "14", "13", "12", "11"},
			OSver: map[string][]string{
				"Mac": {"11", "10.15", "10.14", "10.13"},
			},
			UserAgent: map[string]string{},
		},
		{
			Name:     "Microsoft Edge",
			Versions: []string{"96", "95", "94", "93", "92"},
			OSver: map[string][]string{
				"Windows": {"10", "8.1", "8", "7"},
				"Mac":     {"11", "10.15", "10.14", "10.13"},
			},
			UserAgent: map[string]string{
				"Windows": "Mozilla/5.0 (Windows NT %s; Win64; x64) AppleWebKit/%s (KHTML, like Gecko) Chrome/%s Safari/%s Edg/%s",
				"Mac":     "Mozilla/5.0 (Macintosh; Intel Mac OS X %s) AppleWebKit/%s (KHTML, like Gecko) Chrome/%s Safari/%s Edg/%s",
			},
		},
		{
			Name:     "Opera",
			Versions: []string{"80", "79", "78", "77", "76"},
			OSver: map[string][]string{
				"Windows": {"10", "8.1", "8", "7"},
				"Mac":     {"11", "10.15", "10.14", "10.13"},
				"Linux":   {"Ubuntu 20", "Debian 10", "Fedora 34"},
			},
			UserAgent: map[string]string{
				"Windows": "Mozilla/5.0 (Windows NT %s; Win64; x64) AppleWebKit/%s (KHTML, like Gecko) Chrome/%s Safari/%s OPR/%s",
				"Mac":     "Mozilla/5.0 (Macintosh; Intel Mac OS X %s) AppleWebKit/%s (KHTML, like Gecko) Chrome/%s Safari/%s OPR/%s",
				"Linux":   "Mozilla/5.0 (X11; %s; Linux x86_64) AppleWebKit/%s (KHTML, like Gecko) Chrome/%s Safari/537.36 OPR/%s",
			},
		},
		{
			Name:     "Internet Explorer",
			Versions: []string{"11", "10", "9", "8", "7"},
			OSver: map[string][]string{
				"Windows": {"10", "8.1", "8", "7"},
			},
			UserAgent: map[string]string{},
		},
	}

	return browsers
}

func (in *Instance) getRandomMapSlice(data map[int][]uint16) []uint16 {
	dex := rand.Intn(len(data))
	i := 0
	for _, val := range data {
		if i == dex {
			return val
		}
		i++
	}
	return nil
}

func (in *Instance) getRandomBrowser(browsers []BrowserData) BrowserData {
	dex := rand.Intn(len(browsers))
	return browsers[dex]
}

func (in *Instance) getRandomData(data []string) string {
	dex := rand.Intn(len(data))
	return data[dex]
}

func (in *Instance) getRandomKey(m map[string][]string) string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	index := rand.Intn(len(keys))
	return keys[index]
}

func (in *Instance) LocationData() (data BillingAddress) {
	//data.Name = randomdata.FullName(randomdata.Male)
	//data.Line = randomdata.Street()
	//data.Line2 = ""
	//data.Email = ""
	//data.City = randomdata.City()
	//data.State = randomdata.State(randomdata.Large)
	//data.Country = randomdata.Country(randomdata.TwoCharCountry)
	//data.PostCode = randomdata.Digits(4)

	data.Name = "micheal white"
	data.Line = "300 Rolling Oaks dr"
	data.Line2 = ""
	data.Email = ""
	data.City = "Thousand Oaks"
	data.State = "CA"
	data.Country = "US"
	data.PostCode = "91361"

	return data
}

//could use format
func (in *Instance) StrLog_V(text string, data string, s time.Time) {
	e := time.Since(s)
	fmt.Println("["+bg+""+e.String()[:3]+"ms"+rb+"] ["+g+"✓"+r+"]"+in.Worker+""+text+": "+gr+data, rb+r)
}

func (in *Instance) StrLog_E(text string, data string, s time.Time) {
	e := time.Since(s)
	fmt.Println("["+bg+""+e.String()[:3]+"ms"+rb+"] ["+red+"X"+r+"]"+in.Worker+""+text+": "+gr+data, rb+r)
}

func (in *Instance) StrLog_R(text string, data string, s time.Time) {
	e := time.Since(s)
	fmt.Println("["+bg+""+e.String()[:3]+"ms"+rb+"] ["+yellow+"-"+r+"]"+in.Worker+""+text+": "+gr+data, rb+r)
}

func (in *Instance) WriteFile(files string, item string) {
	f, err := os.OpenFile(files, os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	_, ers := f.WriteString(item + "\n")
	if ers != nil {
		log.Println(err)
	}
}

func (in *Instance) Cls() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func (in *Instance) Ciphers() (cipher []BrowserData) {
	cipher = []BrowserData{
		{
			CipherList: map[int][]uint16{
				1:  {0xc031, 0xc00a, 0xc009, 0xccaa, 0xcca7, 0xc023, 0xc027, 0x13c0, 0x13c1, 0xc0ac},
				2:  {0xc0ae, 0x13c3, 0x13c4, 0xc02b, 0xcca9, 0xcca8, 0xc02e, 0xc030, 0xc02c, 0xc02d},
				3:  {0xc00b, 0xc00c, 0xc00d, 0xc00e, 0xc00f, 0xc011, 0x1301, 0x1302, 0x1303, 0x1304},
				4:  {0x1305, 0x1306, 0x1307, 0x1308, 0x1309, 0x130a, 0x130b, 0x130c, 0x130d, 0x130e},
				5:  {0x130f, 0xccaa, 0xcca7, 0xc02b, 0xc00c, 0xc00d, 0xc00e, 0xc00f, 0xc011, 0x1301},
				6:  {0x1302, 0x1303, 0x1304, 0x1305, 0x1306, 0x1307, 0x1308, 0x1309, 0x130a, 0x130b},
				7:  {0x130c, 0x130d, 0x130e, 0x130f, 0xccaa, 0xcca7, 0xcca8, 0xc02e, 0xc030, 0xc02c},
				8:  {0xc02d, 0xc00a, 0xc009, 0xccaa, 0xcca7, 0xc023, 0xc027, 0xc00b, 0xc00c, 0xc00d},
				9:  {0xc00e, 0xc00f, 0xc011, 0x1301, 0x1302, 0x1303, 0x1304, 0x1305, 0x1306, 0x1307},
				10: {0x1308, 0x1309, 0x130a, 0x130b, 0x130c, 0x130d, 0x130e, 0x130f, 0xccaa, 0xcca7},
				11: {0xc02b, 0xc00c, 0xc00d, 0xc00e, 0xc00f, 0xc011, 0x1301, 0x1302, 0x1303, 0x1304},
				12: {0x1305, 0x1306, 0x1307, 0x1308, 0x1309, 0x130a, 0x130b, 0x130c, 0x130d, 0x130e},
				13: {0x130f, 0xccaa, 0xcca7, 0xc02b, 0xc00c, 0xc00d, 0xc00e, 0xc00f, 0xc011, 0x1301},
				14: {0x1302, 0x1303, 0x1304, 0x1305, 0x1306, 0x1307, 0x1308, 0x1309, 0x130a, 0x130b},
				15: {0x130c, 0x130d, 0x130e, 0x130f, 0xccaa, 0xcca7, 0xcca8, 0xc02e, 0xc030, 0xc02c},
				16: {0xc02d, 0xc00a, 0xc009, 0xccaa, 0xcca7, 0xc023, 0xc027, 0xc00b, 0xc00c, 0xc00d},
				17: {0xc00e, 0xc00f, 0xc011, 0x1301, 0x1302, 0x1303, 0x1304, 0x1305, 0x1306, 0x1307},
				18: {0x1308, 0x1309, 0x130a, 0x130b, 0x130c, 0x130d, 0x130e, 0x130f, 0xccaa, 0xcca7},
				19: {0xc02b, 0xc00c, 0xc00d, 0xc00e, 0xc00f, 0xc011, 0x1301, 0x1302, 0x1303, 0x1304},
				20: {0x1305, 0x1306, 0x1307, 0x1308, 0x1309, 0x130a, 0x130b, 0x130c, 0x130d, 0x130e},
			},
			EllipticCurves: map[int][]uint16{
				1: {0x0017, 0x0018, 0x001D, 0x001E, 0x0020, 0x0022, 0x0023, 0x0024, 0x0025, 0x0026, 0x0027, 0x0029, 0x002A, 0x002B, 0x002D, 0x002F},
			},
		},
	}
	return cipher
}

func (in *Instance) WebSock(Token string) {

	dialer := websocket.Dialer{}
	ws, _, err := dialer.Dial("wss://gateway.discord.gg/?encoding=json&v=10&compress=zlib-stream", shttp.Header{
		"Origin":     []string{"https://discord.com"},
		"User-Agent": []string{in.Cfg.Mode.Network.Agent},
	})
	in.Err(err)
	rand.Seed(time.Now().UnixNano())
	_, _, _ = ws.ReadMessage()
	browser := in.Browser()
	BrowserClient := in.getRandomBrowser(browser)

	Payload, _ := json.Marshal(map[string]interface{}{
		"op": 2,
		"d": map[string]interface{}{
			"token":        Token,
			"capabilities": 125,
			"properties": map[string]interface{}{
				"os":                       "Windows",
				"browser":                  BrowserClient.Name,
				"system_locale":            "en-US",
				"browser_user_agent":       "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) discord/1.0.9006 Chrome/91.0.4472.164 Electron/13.6.6 Safari/537.36",
				"browser_version":          in.getRandomData(BrowserClient.Versions),
				"os_version":               "10",
				"referrer":                 "",
				"referring_domain":         "",
				"referrer_current":         "",
				"referring_domain_current": "",
				"release_channel":          "stable",
				"client_build_number":      Bnum,
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

	err = ws.WriteMessage(websocket.TextMessage, Payload)
	in.Err(err)
	_, _, _ = ws.ReadMessage()

	ws.Close()
}
