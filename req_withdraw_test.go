package go_zpay

import (
	"testing"
)

func TestWithdraw(t *testing.T) {
	vLog := VLog{}
	//构造client
	cli := NewClient(vLog, &ZPayInitParams{MerchantInfo{MERCHANT_CODE, MERCHANT_KEY}, DEPOSIT_URL, DEPOSIT_RESPONSE_URL, WITHDRAW_URL, WITHDRAW_RESPONSE_URL})

	//发请求
	resp, err := cli.WithdrawReq(GenWithdrawRequestDemo())
	if err != nil {
		cli.logger.Errorf("err:%s\n", err.Error())
		return
	}
	cli.logger.Infof("resp:%+v\n", resp)
}

func GenWithdrawRequestDemo() ZPayWithdrawReq {
	return ZPayWithdrawReq{
		Currency:      "VND",
		MerchantRefNo: "20230824152000001",
		Amount:        "1000.00",
		BankName:      "ABBANK",
		IfscCode:      "IBK",
		AccountNumber: "1234567890",
		AccountName:   "jane",
		TaxNumber:     "1234567890",
	}
}
