package main

import (
	"fmt"
	"math/rand"
	"redeemer/src"
	"sync"
	"time"
)

var (
	c   = &src.Instance{}
	cnt int
)

func main() {
	l, _ := c.Logo()
	var choice int
	c.Cls()
	fmt.Println("\u001B[34;1m"+string(l)+"\n", "\033[39m")
	c.Menu()

	fmt.Scanln(&choice)

	switch choice {
	case 1:

		s := time.Now()
		instances, err := c.Configuration()
		if err != nil {
			c.StrLog_E("Encountered an Error", err.Error(), s)
		}
		cfg := c.Config()
		var wg sync.WaitGroup

		if cfg.Mode.Configs.Wrks >= 1 {
			wg.Add(cfg.Mode.Configs.Wrks)

			for i := 0; i < cfg.Mode.Configs.Wrks; i++ {
				go func(i int) {
					defer wg.Done()
					for j := i; j < len(instances); j += cfg.Mode.Configs.Wrks {
						Run(instances, j)
					}
				}(i)
			}
			wg.Wait()
		} else {
			wg.Add(len(instances))

			for i := 0; i < len(instances); i++ {
				go func(i int) {
					defer wg.Done()
					Run(instances, i)
				}(i)
			}
			wg.Wait()
		}

	case 2:
		s := time.Now()
		instances, err := c.Configuration()

		if err != nil {
			c.StrLog_E("Encountered an Error", err.Error(), s)
		}
		var ID, st string
		var R bool

		fmt.Print("Server ID:")
		fmt.Scanln(&ID)
		fmt.Print("Boost Twice Per Token? y/n:")
		fmt.Scanln(&st)

		if st == "y" {
			R = true
		}

		var wg sync.WaitGroup
		wg.Add(len(instances))

		for i := 0; i < len(instances); i++ {
			go func(i int) {
				defer wg.Done()
				in := instances[i]
				Token := in.Token
				C := in.GetCookie()
				in.Boost(Token, C, ID, R)
			}(i)
		}
		wg.Wait()
	case 3:
		s := time.Now()
		instances, err := c.Configuration()
		if err != nil {
			c.StrLog_E("Encountered an Error", err.Error(), s)
		}

		var invite string

		fmt.Print("	discord.gg/")
		fmt.Scanln(&invite)
		var wg sync.WaitGroup
		wg.Add(len(instances))

		for i := 0; i < len(instances); i++ {
			go func(i int) {
				defer wg.Done()
				in := instances[i]
				Token := in.Token
				C := in.GetCookie()
				//in.WebSock(Token)
				//wsd, _ := in.Connect(Token)
				in.Join(Token, invite, C, "", "", "")
			}(i)
		}
		wg.Wait()
	case 4:

		s := time.Now()
		instances, err := c.Configuration()
		if err != nil {
			c.StrLog_E("Encountered an Error", err.Error(), s)
		}

		var wg sync.WaitGroup
		wg.Add(len(instances))

		for i := 0; i < len(instances); i++ {
			go func(i int) {
				defer wg.Done()
				in := instances[i]
				Token := in.Token
				C := in.GetCookie()
				data, tim := in.GetData(Token, C)
				in.NitroToken(Token, data, tim)
			}(i)
		}
		wg.Wait()
	case 5:
	default:
		fmt.Println("	Wrong Choice")
		time.Sleep(2 * time.Second)
		main()
	}
	//add nitro checker func
}

func Wait() {
	rand.Seed(time.Now().UnixNano())
	interval := rand.Intn(101) + 500
	time.Sleep(time.Duration(interval) * time.Millisecond)
}

func Run(instances []src.Instance, i int) {
	in := instances[i]
	Token := in.Token
	link := in.Promo
	in.WebSock(Token)
	C := in.GetCookie()
	key := in.Session(in.Cfg.Mode.Configs.PromoType + link)
	data := in.StripeID()
	Confirm := in.StripeToken(in.Vcc, data, key)
	Secret := in.GetSecClient(Token, link, C)
	AddressToken, address := in.GetAuth(Token, C, link)
	auth := in.StripeConfirm(Secret, key, Confirm.Id, data, address)
	payment := in.AddPayment(Token, auth, AddressToken, C, link, address, Secret, key, Confirm.Id, data)
	tf, paymentID := in.Claim(Token, link, payment)
	if tf {
	retry:
		if cnt >= 2 {
		} else {
			stripePayments := in.DiscordIntents(Token, link, paymentID)
			Wait()
			in.ConfirmDiscordIntents(stripePayments, key)
			Wait()
			DSecData := in.ConfirmPaymentIntents(key, stripePayments.StripePaymentIntentClientSecret)
			Wait()
			in.Authenticate(key, DSecData)
			Wait()
			in.ConfirmDiscordIntents(stripePayments, key)
			Wait()
			tf, paymentID = in.Claim(Token, link, payment)
			if tf {
				cnt++
				goto retry
			}
		}
	}
}
