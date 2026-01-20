package go_zpay

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestWithdrawCallback(t *testing.T) {
	vLog := VLog{}
	//构造client
	cli := NewClient(vLog, &ZPayInitParams{MerchantInfo{MERCHANT_CODE, MERCHANT_KEY}, DEPOSIT_URL, DEPOSIT_RESPONSE_URL, WITHDRAW_URL, WITHDRAW_RESPONSE_URL})

	//1. 获取请求
	req := GenWdCallbackRequestDemo() //提现的返回
	var backReq ZPayWithdrawCallbackReq
	err := json.Unmarshal([]byte(req), &backReq)
	if err != nil {
		cli.logger.Errorf("Error:%s", err.Error())
		return
	}

	fmt.Printf("backReq: %v\n", backReq)
	//2. 处理请求
	err = cli.WithdrawCallback(backReq, func(ZPayWithdrawCallbackReq) error { return nil })
	if err != nil {
		cli.logger.Errorf("Error:%s", err.Error())
		return
	}

	cli.logger.Infof("resp:%+v\n", backReq)
}

// status=200&message=Payout+successful&amount=1000.00&after_charges_amount=1012.00&transaction_id=RN20012026ABVDMHQK&
// currency=VND&reference_code=20230824152000001&status_code=10001&created_at=2026-01-20T03%3A00%3A53Z&updated_at=2026-01-20T03%3A13%3A47Z&timestamp=1768878827&signature=2A4F2C0CF9433BF674BB449DDEB15E7DFF9C7447B88F412E2645D8CDE3703F87
func GenWdCallbackRequestDemo() string {
	return `{"status":"200","message":"Payout+successful",
	"amount":"1000.00","after_charges_amount":"1012","transaction_id":"RN20012026ABVDMHQK",
	"currency":"VND","reference_code":"20230824152000001","status_code":"10001","created_at":"2026-01-20 00:00:53",
	"updated_at":"2026-01-20 00:13:47","timestamp":"1768878827",
	"signature":"2A4F2C0CF9433BF674BB449DDEB15E7DFF9C7447B88F412E2645D8CDE3703F87"}`
}
