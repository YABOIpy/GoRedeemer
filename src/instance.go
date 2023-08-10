package src

import (
	http "github.com/Danny-Dasilva/fhttp"
	goclient "redeemer/client"
	"strconv"
	"strings"
	"sync"
)

func (in *Instance) Configuration() ([]Instance, error) {
	var (
		clienterr error
		worker    int
		wrkr      string
		proxy     string
		vcc       []string
		Instances []Instance
		Client    *http.Client
		cfg       = in.Config()
	)

	proxies, _, _ := in.ReadFile("proxies.txt")
	vccs, _, _ := in.ReadFile("vcc.txt")
	link, _, _ := in.ReadFile("promos.txt")
	Tokens, _, err := in.ReadFile("tokens.txt")
	if err != nil {
		return Instances, err
	}
	if len(proxies) >= 1 || strings.Count(cfg.Mode.Network.Proxy, "") > 2 {
		cfg.Con.ProxyMode = g + "True"
	}

	var wg sync.WaitGroup
	var mutex sync.Mutex
	wg.Add(len(Tokens))

	for i := 0; i < len(Tokens); i++ {
		go func(i int) {
			defer wg.Done()
			if strings.Count(cfg.Mode.Network.Proxy, "") > 2 {
				proxy = "http://" + cfg.Mode.Network.Proxy
			} else {
				if len(proxies) >= 1 {
					proxy = "http://" + proxies[i]
				} else {
					cfg.Con.ProxyMode = red + "False"
					proxy = ""
				}
			}
			//Ja3, _ := in.CreateJa3()
			//ept := strings.Split(tknf[i], ":")
			Client, clienterr = goclient.NewClient(goclient.Browser{
				JA3:       cfg.Mode.Network.Ja3,
				UserAgent: cfg.Mode.Network.Agent,
				Cookies:   nil,
			},
				cfg.Mode.Network.TimeOut,
				cfg.Mode.Network.Redirect,
				cfg.Mode.Network.Agent,
				proxy,
			)
			if clienterr != nil {
				mutex.Lock()
				err = clienterr
				mutex.Unlock()
				return
			}

			if cfg.Mode.Configs.Workers {
				worker++
				wrkr = "[WORKER-" + clr + strconv.Itoa(worker) + r + "] "
			} else {
				wrkr = " "
			}
			if cfg.Mode.Configs.VccUses <= 1 {
				vcc = append(vcc, vccs...)
			} else {
				for j := 0; j < cfg.Mode.Configs.VccUses; j++ {
					vcc = append(vcc, vccs...)
				}
			}
			if len(link) != len(Tokens) && len(vcc) != len(Tokens) {

			}

			mutex.Lock()
			Instances = append(Instances, Instance{
				Client: Client,
				Worker: wrkr,
				Promo:  strings.Split(link[i], cfg.Mode.Configs.PromoType)[1],
				Token:  Tokens[i],
				//Email:  ept[1],
				//Pass:   ept[0],
				Cfg: cfg,
				Vcc: vcc[i],
			})
			mutex.Unlock()
		}(i)
	}

	return Instances, err
}
