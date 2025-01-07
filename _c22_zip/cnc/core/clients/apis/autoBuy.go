package apis

import (
	"Nosviak2/core/database"
	"Nosviak2/core/functions"
	"Nosviak2/core/sources/layouts/json"

	"Nosviak2/core/sources/layouts/toml"
	decoder "encoding/json"
	"fmt"
	"io"
	"net/http"

	"strings"
	"time"
)

type Sellix struct {
	Event string `json:"event"`
	Data  struct {
		ID                        int          			`json:"id"`
		Uniqid                    string       			`json:"uniqid"`
		RecurringBillingID        interface{}  			`json:"recurring_billing_id"`
		Type                      string       			`json:"type"`
		Subtype                   interface{}  			`json:"subtype"`
		Total                     int          			`json:"total"`
		TotalDisplay              int          			`json:"total_display"`
		ProductVariants           interface{}  			`json:"product_variants"`
		ExchangeRate              int          			`json:"exchange_rate"`
		CryptoExchangeRate        int          			`json:"crypto_exchange_rate"`
		Currency                  string       			`json:"currency"`
		ShopID                    int          			`json:"shop_id"`
		ShopImageName             interface{}  			`json:"shop_image_name"`
		ShopImageStorage          interface{}  			`json:"shop_image_storage"`
		CloudflareImageID         interface{}  			`json:"cloudflare_image_id"`
		Name                      interface{}  			`json:"name"`
		CustomerEmail             string       			`json:"customer_email"`
		PaypalEmailDelivery       bool         			`json:"paypal_email_delivery"`
		ProductID                 string       			`json:"product_id"`
		ProductTitle              string       			`json:"product_title"`
		ProductType               string       			`json:"product_type"`
		SubscriptionID            interface{}  			`json:"subscription_id"`
		SubscriptionTime          interface{}  			`json:"subscription_time"`
		Gateway                   interface{}  			`json:"gateway"`
		Blockchain                interface{}  			`json:"blockchain"`
		PaypalApm                 interface{}  			`json:"paypal_apm"`
		StripeApm                 interface{}  			`json:"stripe_apm"`
		PaypalEmail               interface{}  			`json:"paypal_email"`
		PaypalOrderID             interface{}  			`json:"paypal_order_id"`
		PaypalPayerEmail          interface{}  			`json:"paypal_payer_email"`
		PaypalFee                 int          			`json:"paypal_fee"`
		PaypalSubscriptionID      interface{}  			`json:"paypal_subscription_id"`
		PaypalSubscriptionLink    interface{}  			`json:"paypal_subscription_link"`
		LexOrderID                interface{}  			`json:"lex_order_id"`
		LexPaymentMethod          interface{}  			`json:"lex_payment_method"`
		PaydashPaymentID          interface{}  			`json:"paydash_paymentID"`
		StripeClientSecret        interface{}  			`json:"stripe_client_secret"`
		StripePriceID             interface{}  			`json:"stripe_price_id"`
		SkrillEmail               interface{}  			`json:"skrill_email"`
		SkrillSid                 interface{}  			`json:"skrill_sid"`
		SkrillLink                interface{}  			`json:"skrill_link"`
		PerfectmoneyID            interface{}  			`json:"perfectmoney_id"`
		BinanceInvoiceID          interface{}  			`json:"binance_invoice_id"`
		BinanceQrcode             interface{}  			`json:"binance_qrcode"`
		BinanceCheckoutURL        interface{}  			`json:"binance_checkout_url"`
		CryptoAddress             interface{}  			`json:"crypto_address"`
		CryptoAmount              int          			`json:"crypto_amount"`
		CryptoReceived            int          			`json:"crypto_received"`
		CryptoURI                 interface{}  			`json:"crypto_uri"`
		CryptoConfirmationsNeeded int          			`json:"crypto_confirmations_needed"`
		CryptoScheduledPayout     bool         			`json:"crypto_scheduled_payout"`
		CryptoPayout              int          			`json:"crypto_payout"`
		FeeBilled                 bool         			`json:"fee_billed"`
		BillInfo                  interface{}  			`json:"bill_info"`
		CashappQrcode             interface{}  			`json:"cashapp_qrcode"`
		CashappNote               interface{}  			`json:"cashapp_note"`
		CashappCashtag            interface{}  			`json:"cashapp_cashtag"`
		Country                   interface{}  			`json:"country"`
		Location                  string       			`json:"location"`
		IP                        interface{}  			`json:"ip"`
		IsVpnOrProxy              bool         			`json:"is_vpn_or_proxy"`
		UserAgent                 interface{}  			`json:"user_agent"`
		Quantity                  int          			`json:"quantity"`
		CouponID                  interface{}  			`json:"coupon_id"`
		CustomFields              map[string]interface{}   	`json:"custom_fields"`
		DeveloperInvoice          bool          		`json:"developer_invoice"`
		DeveloperTitle            interface{}   		`json:"developer_title"`
		DeveloperWebhook          interface{}   		`json:"developer_webhook"`
		DeveloperReturnURL        interface{}   		`json:"developer_return_url"`
		Status                    string        		`json:"status"`
		StatusDetails             interface{}   		`json:"status_details"`
		VoidDetails               interface{}   		`json:"void_details"`
		Discount                  int           		`json:"discount"`
		FeePercentage             int           		`json:"fee_percentage"`
		DayValue                  int           		`json:"day_value"`
		Day                       string        		`json:"day"`
		Month                     string        		`json:"month"`
		Year                      int           		`json:"year"`
		ProductAddons             interface{}   		`json:"product_addons"`
		CreatedAt                 int           		`json:"created_at"`
		UpdatedAt                 int           		`json:"updated_at"`
		UpdatedBy                 int           		`json:"updated_by"`
		IPInfo                    interface{}   		`json:"ip_info"`
		Serials                   []interface{} 		`json:"serials"`
		LockedSerials             []interface{} 		`json:"locked_serials"`
		Webhooks                  []interface{} 		`json:"webhooks"`
		PaypalDispute             interface{}   		`json:"paypal_dispute"`
		ProductDownloads          []interface{} 		`json:"product_downloads"`
		StatusHistory             []struct {
			ID        int    `json:"id"`
			InvoiceID string `json:"invoice_id"`
			Status    string `json:"status"`
			Details   string `json:"details"`
			CreatedAt int    `json:"created_at"`
		} `json:"status_history"`
		CryptoTransactions           []interface{} `json:"crypto_transactions"`
		Products                     []interface{} `json:"products"`
		GatewaysAvailable            []string      `json:"gateways_available"`
		ShopPaypalCreditCard         bool          `json:"shop_paypal_credit_card"`
		ShopForcePaypalEmailDelivery bool          `json:"shop_force_paypal_email_delivery"`
	} `json:"data"`
}

// Autobuy referenes whenever sellix forwards an autobuy request.
func Autobuy(w http.ResponseWriter, r *http.Request)  {
	if !r.URL.Query().Has("key") || r.URL.Query().Get("key") != toml.ApiToml.API.PrivateKey|| strings.Join(strings.Split(r.RemoteAddr, ":")[:strings.Count(r.RemoteAddr, ":")], ":") != "99.81.24.41" {
		EncodeAndReturn(map[string]interface{}{"error": "missing confirm token", "status": false}, w)
		return	
	}

	// reads the sellix body
	content, err := io.ReadAll(r.Body)
	if err != nil {
		EncodeAndReturn(map[string]interface{}{"error": "unable to read body", "status": false}, w)
		return	
	}

	var t *Sellix = new(Sellix)
	if err := decoder.Unmarshal(content, &t); err != nil {
		EncodeAndReturn(map[string]interface{}{"error": "invalid body", "status": false}, w)
		return	
	}

	e, ok := json.Products[t.Data.ProductID]
	if !ok || e == nil {
		EncodeAndReturn(map[string]interface{}{"error": "unknown product", "status": false}, w)
		return	
	}

	username, ok := t.Data.CustomFields["username"]
	if !ok || len(fmt.Sprint(username)) == 0 {
		EncodeAndReturn(map[string]interface{}{"error": "invalid username", "status": false}, w)
		return	
	}

	// Depends on the action, without create_user it will just feed as a modify user action.
	if e.Action != "create_user" {
		context, err := database.Conn.FindUser(fmt.Sprint(username))
		if err != nil || context == nil {
			EncodeAndReturn(map[string]interface{}{"error": "invalid user", "status": false}, w)
			return	
		}

		functions.MergeFieldsWithUser(e, context)
		if err := database.Conn.EditUser(context); err != nil {
			EncodeAndReturn(map[string]interface{}{"error": "error occurred", "status": false}, w)
			return	
		}

		w.WriteHeader(http.StatusOK)
		EncodeAndReturn(map[string]interface{}{"message": "pushed changed", "status": true}, w)
		return	
	}

	password, ok := t.Data.CustomFields["password"]
	if !ok || len(fmt.Sprint(username)) == 0 {
		EncodeAndReturn(map[string]interface{}{"error": "invalid password", "status": false}, w)
		return	
	}

	user := &database.User{
		Parent: 0,
		Username: fmt.Sprint(username),
		Password: fmt.Sprint(password),
		MaxSlaves: 0,
		NewUser: true,
		Locked: false,
		Expiry: 0,
		Updated: time.Now().Unix(),
		Created: time.Now().Unix(),
	}

	if err := functions.MergeFieldsWithUser(e, user); err != nil {
		EncodeAndReturn(map[string]interface{}{"error": "error occurred", "status": false}, w)
		return	
	}

	if err := database.Conn.MakeUser(user); err != nil {
		EncodeAndReturn(map[string]interface{}{"error": "error occurred", "status": false}, w)
		return	
	}

	w.WriteHeader(http.StatusOK)
	EncodeAndReturn(map[string]interface{}{"message": "pushed changed", "status": true}, w)	
}
