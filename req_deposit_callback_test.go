package go_zpay

import (
	"encoding/json"
	"fmt"
	"testing"
)

type VLog struct {
}

func (l VLog) Debugf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Infof(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Warnf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Errorf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}

func TestCallback(t *testing.T) {
	vLog := VLog{}
	//构造client
	cli := NewClient(vLog, &ZPayInitParams{MerchantInfo{MERCHANT_CODE, MERCHANT_KEY}, DEPOSIT_URL, DEPOSIT_RESPONSE_URL, WITHDRAW_URL, WITHDRAW_RESPONSE_URL})

	//1. 获取请求
	req := GenCallbackRequestDemo() //提现的返回
	var backReq ZPayDepositCallbackReq
	err := json.Unmarshal([]byte(req), &backReq)
	if err != nil {
		cli.logger.Errorf("Error:%s", err.Error())
		return
	}

	//2. 处理请求
	err = cli.DepositCallback(backReq, func(ZPayDepositCallbackReq) error { return nil })
	if err != nil {
		cli.logger.Errorf("Error:%s", err.Error())
		return
	}
	cli.logger.Infof("resp:%+v\n", backReq)
}

// func GenCallbackRequestDemo() string {
// bill_status 1=订单已取消 2=订单已激活
// return `{"sign":"2c89857a90e2773f27583d954c91f40c","bill_no":"2025100443562675418","bill_status":1,"sys_no":"505299"}`
// }

// status=200&message=Transaction+successful&amount=188000.00&transaction_id=T17688731411IR8&currency=VND&reference_code=22515161368&status_code=10001&created_at=2026-01-20T01%3A39%3A01Z&updated_at=2026-01-20T01%3A40%3A48Z&timestamp=1768873248&signature=8B997AD600CEC6DBDDB3E6C1513443EF8A6237A9EE5E30524CE7DFBD8887C753
func GenCallbackRequestDemo() string {
	return `{"status":"200","message":"Transaction+successful",
	"amount":"188000.00","transaction_id":"T17688731411IR8",
	"currency":"VND","reference_code":"22515161368","status_code":"10001",
	"created_at":"2025-10-04 10:24:26","updated_at":"2025-10-04 10:24:26"}`

}

// return ZPayDepositCallbackReq{
// 	Status:        "200",
// 	Message:       "Transaction+successful",
// 	Amount:        "188000.00",
// 	TransactionId: "T17688731411IR8",
// 	Currency:      "VND",
// 	ReferenceCode: "22515161368",
// 	StatusCode:    "10001",
// 	CreatedAt:     "2025-10-04 10:24:26",
// 	UpdatedAt:     "2025-10-04 10:24:26",
// 	Timestamp:     "1768873248",
// 	Signature:     "8B997AD600CEC6DBDDB3E6C1513443EF8A6237A9EE5E30524CE7DFBD8887C753",
// }
