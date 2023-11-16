package funcs

const (
	safaricomBaseURL   = "https://sandbox.safaricom.co.ke"
	oauthTokenEndpoint = "/oauth/v1/generate?grant_type=client_credentials"
	stkPushEndpoint    = "/mpesa/stkpush/v1/processrequest"
	lipaNaMpesaOnline  = "CustomerBuyGoodsOnline"
	//passKey            = "bfb279f9aa9bdbcf158e97dd71a467cd2e0c893059b10f78e6b72ada1ed2c919"
	shortCode = "174379"
	//callbackURL = "" // API endpoint where Safaricom will send a callback or notification to your server after a transaction is completed or when there are updates on the transaction status.
	amount      = 10000          // Change this to the amount you want to charge
	phoneNumber = "254713194323" // Change this to the customer's phone number
)
