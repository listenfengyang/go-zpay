package go_zpay

type ZPayInitParams struct {
	MerchantInfo `yaml:",inline" mapstructure:",squash"`

	DepositUrl          string `json:"depositUrl" mapstructure:"depositUrl" config:"depositUrl"  yaml:"depositUrl"`
	DepositResponseUrl  string `json:"depositResponseUrl" mapstructure:"depositResponseUrl" config:"depositResponseUrl"  yaml:"depositResponseUrl"`
	WithdrawUrl         string `json:"withdrawUrl" mapstructure:"withdrawUrl" config:"withdrawUrl"  yaml:"withdrawUrl"`
	WithdrawResponseUrl string `json:"withdrawResponseUrl" mapstructure:"withdrawResponseUrl" config:"withdrawResponseUrl"  yaml:"withdrawResponseUrl"`
}

type MerchantInfo struct {
	MerchantCode string `json:"merchantCode" mapstructure:"merchantCode" config:"merchantCode"  yaml:"merchantCode"` // merchantCode
	MerchantKey  string `json:"merchantKey" mapstructure:"merchantKey" config:"merchantKey"  yaml:"merchantKey"`     // merchantKey
}

//============================================================

// zpay入金
// 银行类型有：
// INR(value:UPI/IMPS),
// THB(VALUE:THB_decimalDown/THB_decimalUp/THB_KYC/QR),
// MYR(value:EWALLET/ONLINE_BANKING),
// MYR_NATIVE(VALUE:FPX/EWALLET),
// IDR_NATIVE(VALUE:VA/QRIS),
// VND(VALUE:ONLINE_BANKING/OFFLINE_BANKING/BANK_QR),
// BRL,BDT,AUD(VALUE:LEAVE BLANK)
type ZPayDepositReq struct {
	Currency  string `json:"currency" mapstructure:"currency"`   // 币种
	PaymentID string `json:"paymentID" mapstructure:"paymentID"` // 唯一交易ID
	Amount    string `json:"amount" mapstructure:"amount"`       // 金额
	Remark    string `json:"remark" mapstructure:"remark"`       // 备注
	BankType  string `json:"bankType" mapstructure:"bankType"`   // 银行类型
	// 以下非必填
	// ResponseURL string `json:"responseURL" mapstructure:"responseURL"` // 回调地址
	// Signature   string `json:"signature" mapstructure:"signature"`     // 签名
	// 以下参数只在THB_KYC模式加
	// CustomerUsername          string `json:"customerUsername" mapstructure:"customerUsername"`                   // 客户用户名
	// CustomerEmail             string `json:"customerEmail" mapstructure:"customerEmail"`                         // 客户邮箱
	// CustomerPhone             string `json:"customerPhone" mapstructure:"customerPhone"`                         // 客户手机号
	// CustomerBankName          string `json:"customerBankName" mapstructure:"customerBankName"`                   // 客户银行名称
	// CustomerBankHolderName    string `json:"customerBankHolderName" mapstructure:"customerBankHolderName"`       // 客户银行账号名称
	// CustomerBankAccountNumber string `json:"customerBankAccountNumber" mapstructure:"customerBankAccountNumber"` // 客户银行账号
}

type ZPayDepositRsp struct {
	Status        int32  `json:"status" mapstructure:"status"` //请求状态：200=请求成功 400=请求失败
	Message       string `json:"message" mapstructure:"message"`
	RedirectUrl   string `json:"redirect_url" mapstructure:"redirect_url"`     //重定向URL
	QrString      string `json:"qr_string" mapstructure:"qr_string"`           //二维码字符串
	TransactionId string `json:"transaction_id" mapstructure:"transaction_id"` //交易ID
	ReceivedAt    string `json:"receive_at" mapstructure:"receive_at"`         //接收时间
}

type DepositData struct {
	Status      int32  `json:"status" mapstructure:"status"` //请求状态：200=请求成功 400=请求失败
	Message     string `json:"message" mapstructure:"message"`
	RedirectUrl string `json:"redirect_url" mapstructure:"redirect_url"` //重定向URL
	QrString    string `json:"qr_string" mapstructure:"qr_string"`       //二维码字符串
}

// 入金回调
type ZPayDepositCallbackReq struct {
	Status        string `json:"status" form:"status" mapstructure:"status"` //请求状态：200=请求成功 400=请求失败
	Message       string `json:"message" form:"message" mapstructure:"message"`
	StatusCode    string `json:"status_code" form:"status_code" mapstructure:"status_code"`          //支付状态：10001=支付成功 10002=支付失败
	Amount        string `json:"amount" form:"amount" mapstructure:"amount"`                         //入金金额
	PayableAmount string `json:"payable_amount" form:"payable_amount" mapstructure:"payable_amount"` //可入金金额
	TransactionId string `json:"transaction_id" form:"transaction_id" mapstructure:"transaction_id"` //交易ID
	Currency      string `json:"currency" form:"currency" mapstructure:"currency"`                   //币种
	ReferenceCode string `json:"reference_code" form:"reference_code" mapstructure:"reference_code"` //引用订单号
	CreatedAt     string `json:"created_at" form:"created_at" mapstructure:"created_at"`             //创建时间
	UpdatedAt     string `json:"updated_at" form:"updated_at" mapstructure:"updated_at"`             //更新时间
	Timestamp     string `json:"timestamp" form:"timestamp" mapstructure:"timestamp"`                //时间戳
	Signature     string `json:"signature" form:"signature" mapstructure:"signature"`                //签名
}

// zpay出金
type ZPayWithdrawReq struct {
	Currency      string `json:"currency" mapstructure:"currency"`           // 币种
	BankName      string `json:"bankName" mapstructure:"bankName"`           //银行名称
	IfscCode      string `json:"ifscCode" mapstructure:"ifscCode"`           //IFSC码
	AccountNumber string `json:"accountNumber" mapstructure:"accountNumber"` //银行账号
	AccountName   string `json:"accountName" mapstructure:"accountName"`     //银行账号名称
	Amount        string `json:"amount" mapstructure:"amount"`               //金额
	Description   string `json:"description" mapstructure:"description"`     //描述
	MerchantRefNo string `json:"merchantRefNo" mapstructure:"merchantRefNo"` //商户订单号
	CallbackUrl   string `json:"callbackUrl" mapstructure:"callbackUrl"`     //回调地址
	Signature     string `json:"signature" mapstructure:"signature"`         //签名
}
type ZPayWithdrawRsp struct {
	Status int    `json:"status" mapstructure:"status"` //200=成功，400=失败...
	Msg    string `json:"msg" mapstructure:"msg"`
}

// 出金回调
type ZPayWithdrawCallbackReq struct {
	Status             string `json:"status" form:"status" mapstructure:"status"` //200=成功，400=失败...
	Message            string `json:"message" form:"message" mapstructure:"message"`
	StatusCode         string `json:"status_code" form:"status_code" mapstructure:"status_code"`                            //出金状态：20001=出金成功 20002=出金失败
	Amount             string `json:"amount" form:"amount" mapstructure:"amount"`                                           //出金金额
	AfterChargesAmount string `json:"after_charges_amount" form:"after_charges_amount" mapstructure:"after_charges_amount"` //出金金额（包含手续费）
	TransactionId      string `json:"transaction_id" form:"transaction_id" mapstructure:"transaction_id"`                   //出金订单号
	Currency           string `json:"currency" form:"currency" mapstructure:"currency"`                                     //币种
	ReferenceCode      string `json:"reference_code" form:"reference_code" mapstructure:"reference_code"`                   //引用订单号
	CreatedAt          string `json:"created_at" form:"created_at" mapstructure:"created_at"`                               //创建时间
	UpdatedAt          string `json:"updated_at" form:"updated_at" mapstructure:"updated_at"`                               //更新时间
	Timestamp          string `json:"timestamp" form:"timestamp" mapstructure:"timestamp"`                                  //时间戳
	Signature          string `json:"signature" form:"signature" mapstructure:"signature"`                                  //签名
}

type ZPayWithdrawCallbackRsp struct {
	Status  int32  `json:"status" mapstructure:"status"` //请求状态：200=请求成功 400=请求失败
	Message string `json:"message" mapstructure:"message"`
}
