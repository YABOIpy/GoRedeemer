package src

import (
	http "github.com/Danny-Dasilva/fhttp"
	"github.com/gorilla/websocket"
	"time"
)

type Redeem struct {
	Dcfd      string
	Sdcfd     string
	StripeKey string
}

type Instance struct {
	Client *http.Client
	Worker string
	Promo  string
	Token  string
	Email  string
	Pass   string
	Vcc    string
	Cfg    Config
}

type Header struct {
}
type Config struct {
	Mode struct {
		Network struct {
			Redirect bool     `json:"Redirect"`
			TimeOut  int      `json:"TimeOut"`
			Server   []string `json:"Server"`
			Proxy    string   `json:"Proxy"`
			Agent    string   `json:"Agent"`
			Ja3      string   `json:"JA3"`
		} `json:"Net"`

		Configs struct {
			PromoType string `json:"PromoType"`
			Errormsg  bool   `json:"Errors"`
			Workers   bool   `json:"Workers"`
			Wrks      int    `json:"WorkerThreads"`
			VccUses   int    `json:"VccUses"`
		} `json:"Config"`
	} `json:"Modes"`

	Con struct {
		Solution  string
		Tokenclr  string
		ProxyMode string
		Errors    bool
		Toks      int
	}
}

type JoinResp struct {
	SiteKey string `json:"captcha_sitekey"`
	RqToken string `json:"captcha_rqtoken,omitempty"`
}

type CapmonsterResp struct {
	TaskID int `json:"taskId"`
}

type TaskResp struct {
	Status   string      `json:"status"`
	Solution interface{} `json:"solution"`
}

type BoostPayload struct {
	UserPremiumGuildSubscriptionSlotIds []string `json:"user_premium_guild_subscription_slot_ids"`
}

type Subscription struct {
	Id                       string `json:"id"`
	SubscriptionId           string `json:"subscription_id"`
	PremiumGuildSubscription *struct {
		Id      string `json:"id"`
		UserId  string `json:"user_id"`
		GuildId string `json:"guild_id"`
		Ended   bool   `json:"ended"`
	} `json:"premium_guild_subscription"`
	Canceled    bool      `json:"canceled"`
	CooldownEnd time.Time `json:"cooldown_ends_at"`
}

type BoostResp struct {
	Id             string         `json:"id"`
	Type           int            `json:"type"`
	Invalid        bool           `json:"invalid"`
	Flags          int            `json:"flags"`
	Brand          string         `json:"brand"`
	Last4          string         `json:"last_4"`
	ExpiresMonth   int            `json:"expires_month"`
	ExpiresYear    int            `json:"expires_year"`
	BillingAddress BillingAddress `json:"billing_address"`
	Country        string         `json:"country"`
	PaymentGateway int            `json:"payment_gateway"`
	Default        bool           `json:"default"`
}

type XpropData struct {
	OS                 string      `json:"os"`
	Browser            string      `json:"browser"`
	Device             string      `json:"device"`
	SystemLocale       string      `json:"system_locale"`
	BrowserUserAgent   string      `json:"browser_user_agent"`
	BrowserVersion     string      `json:"browser_version"`
	OSVersion          string      `json:"os_version"`
	Referrer           string      `json:"referrer"`
	ReferringDomain    string      `json:"referring_domain"`
	SearchEngine       string      `json:"search_engine"`
	ReferrerCurrent    string      `json:"referrer_current"`
	ReferringDomainCur string      `json:"referring_domain_current"`
	ReleaseChannel     string      `json:"release_channel"`
	ClientBuildNumber  int         `json:"client_build_number"`
	ClientEventSource  interface{} `json:"client_event_source"`
}

type TokenResp struct {
	Id                string        `json:"id"`
	Username          string        `json:"username"`
	GlobalName        string        `json:"global_name"`
	Avatar            string        `json:"avatar"`
	Discriminator     string        `json:"discriminator"`
	PublicFlags       int           `json:"public_flags"`
	Flags             int           `json:"flags"`
	PurchasedFlags    int           `json:"purchased_flags"`
	PremiumUsageFlags int           `json:"premium_usage_flags"`
	Banner            string        `json:"banner"`
	BannerColor       string        `json:"banner_color"`
	AccentColor       int           `json:"accent_color"`
	Bio               string        `json:"bio"`
	Locale            string        `json:"locale"`
	NsfwAllowed       bool          `json:"nsfw_allowed"`
	MfaEnabled        bool          `json:"mfa_enabled"`
	PremiumType       int           `json:"premium_type"`
	LinkedUsers       []interface{} `json:"linked_users"`
	AvatarDecoration  string        `json:"avatar_decoration"`
	Email             string        `json:"email"`
	Verified          bool          `json:"verified"`
	Phone             interface{}   `json:"phone"`
}

type Sock struct {
	Token     string
	Ws        *websocket.Conn
	sessionID string
}

type WsResp struct {
	Op        int    `json:"op"`
	Data      Data   `json:"d,omitempty"`
	Sequence  int    `json:"s,omitempty"`
	EventName string `json:"t,omitempty"`
}

type Confirmation struct {
	Id       string   `json:"id"`
	Object   string   `json:"object"`
	Card     CardData `json:"card"`
	ClientIp string   `json:"client_ip"`
	Created  int      `json:"created"`
	Livemode bool     `json:"livemode"`
	Type     string   `json:"type"`
	Used     bool     `json:"used"`
}

type CardData struct {
	Id                 string      `json:"id"`
	Object             string      `json:"object"`
	AddressCity        interface{} `json:"address_city"`
	AddressCountry     interface{} `json:"address_country"`
	AddressLine1       interface{} `json:"address_line1"`
	AddressLine1Check  interface{} `json:"address_line1_check"`
	AddressLine2       interface{} `json:"address_line2"`
	AddressState       interface{} `json:"address_state"`
	AddressZip         interface{} `json:"address_zip"`
	AddressZipCheck    interface{} `json:"address_zip_check"`
	Brand              string      `json:"brand"`
	Country            string      `json:"country"`
	CvcCheck           string      `json:"cvc_check"`
	DynamicLast4       interface{} `json:"dynamic_last4"`
	ExpMonth           int         `json:"exp_month"`
	ExpYear            int         `json:"exp_year"`
	Funding            string      `json:"funding"`
	Last4              string      `json:"last4"`
	Name               interface{} `json:"name"`
	TokenizationMethod interface{} `json:"tokenization_method"`
	Use                string      `json:"use"`
	Wallet             interface{} `json:"wallet"`
}

type Data struct {
	Content           string                 `json:"content,omitempty"`
	GuildID           string                 `json:"guild_id,omitempty"`
	GuildId           interface{}            `json:"guild_id,omitempty"`
	MessageId         string                 `json:"id,omitempty"`
	Flags             int                    `json:"flags,omitempty"`
	Token             string                 `json:"token,omitempty"`
	Capabilities      int                    `json:"capabilities,omitempty"`
	Compress          bool                   `json:"compress,omitempty"`
	Since             int                    `json:"since,omitempty"`
	Status            string                 `json:"status"`
	Afk               bool                   `json:"afk"`
	HeartbeatInterval int                    `json:"heartbeat_interval,omitempty"`
	SessionID         string                 `json:"session_id,omitempty"`
	Channels          map[string]interface{} `json:"channels,omitempty"`
	Typing            bool                   `json:"typing,omitempty"`
	Threads           bool                   `json:"threads,omitempty"`
	Activities        bool                   `json:"activities,omitempty"`
	ThreadMemberLists interface{}            `json:"thread_member_lists,omitempty"`
	UserID            string                 `json:"user_id,omitempty"`
	MessageID         string                 `json:"message_id,omitempty"`
}

type PaymentResp struct {
	ID      string         `json:"id"`
	Type    int            `json:"type"`
	Invalid bool           `json:"invalid"`
	Flags   int            `json:"flags"`
	Brand   string         `json:"brand"`
	Last4   string         `json:"last_4"`
	ExpireM int            `json:"expires_month"`
	ExpireY int            `json:""`
	Country string         `json:""`
	GateWay string         `json:""`
	Default bool           `json:""`
	Retry   float64        `json:"retry_after,omitempty"`
	Message string         `json:"message"`
	Code    int            `json:"code"`
	Address BillingAddress `json:"billing_address"`
}

type ValidateResp struct {
	Token    string  `json:"token"`
	Secret   string  `json:"client_secret,omitempty"`
	Global   bool    `json:"global"`
	Messsage string  `json:"message,omitempty"`
	Retry    float64 `json:"retry_after,omitempty"`
}

type BillingAddress struct {
	Name     string `json:"name"`
	Line     string `json:"line_1"`
	Line2    string `json:"line_2"`
	Email    string `json:"email"`
	City     string `json:"city"`
	State    string `json:"state"`
	Country  string `json:"country"`
	PostCode string `json:"postal_code"`
}

type BrowserData struct {
	Name           string
	OS             []string `json:"os"`
	OSver          map[string][]string
	UserAgent      map[string]string
	Versions       []string
	CipherSuites   []uint16
	CipherList     map[int][]uint16
	EllipticCurves map[int][]uint16
}
type Suites struct {
	CipherSuites []uint16
}

type StripeData struct {
	Retry        float64 `json:"retry_after,omitempty"`
	ConfirmToken string  `json:"id,omitempty"`
	Payment      PaymentID
	ID           StripeIDs
}

type PaymentID struct {
	PaymentId string `json:"payment_method,omitempty"`
	Error     struct {
		Message string `json:"message"`
	} `json:"error"`
}

type StripeIDs struct {
	Muid string `json:"muid,omitempty"`
	Guid string `json:"guid,omitempty"`
	Sid  string `json:"sid,omitempty"`
}

type SourceID struct {
	ID    string  `json:"id,omitempty"`
	Msg   string  `json:"message"`
	Retry float64 `json:"retry_after,omitempty"`
}

type RedeemResp struct {
	Message   string  `json:"message,omitempty"`
	PaymentId string  `json:"payment_id,omitempty"`
	Retry     float64 `json:"retry_after,omitempty"`
}

type StripePayments struct {
	Message                           string `json:"message,omitempty"`
	StripePaymentIntentClientSecret   string `json:"stripe_payment_intent_client_secret,omitempty"`
	StripePaymentIntentClientSecretId string `json:"stripe_payment_intent_payment_method_id,omitempty"`
}

type StripeIntents struct {
	NextAction struct {
		Type         string `json:"type,omitempty"`
		UseStripeSdk struct {
			ThreeDSecure2Source string `json:"three_d_secure_2_source,omitempty"`
			ServerTranscation   string `json:"server_transaction_id,omitempty"`
			Merchant            string `json:"merchant"`
			ThreeDsMethodURL    string `json:"three_ds_method_url"`
		} `json:"use_stripe_sdk,omitempty"`
	} `json:"next_action,omitempty"`
}

type StripeAuth struct {
	State string      `json:"state,omitempty"`
	Error interface{} `json:"error,omitempty"`
}

var (
	Bnum                                          = in.BuildInfo()
	x, r, g, bg, rb, gr, u, clr, yellow, red, prp = "\u001b[30;1m", "\033[39m", "\033[32m", "\u001b[34;1m", "\u001b[0m", "\u001b[30;1m", "\u001b[4m", "\u001b[38;5;8m", "\033[33m", "\u001B[31m", "\033[35m"
)
