package go_zpay

import (
	"testing"
)

func TestDeposit(t *testing.T) {
	vLog := VLog{}
	//构造client
	cli := NewClient(vLog, &ZPayInitParams{
		MerchantInfo:        MerchantInfo{MERCHANT_CODE, MERCHANT_KEY},
		DepositUrl:          DEPOSIT_URL,
		WithdrawUrl:         WITHDRAW_URL,
		DepositResponseUrl:  DEPOSIT_RESPONSE_URL,
		WithdrawResponseUrl: WITHDRAW_RESPONSE_URL,
	})

	//发请求
	resp, err := cli.Deposit(GenDepositRequestDemo())
	if err != nil {
		cli.logger.Errorf("err:%s\n", err.Error())
		return
	}
	cli.logger.Infof("resp:%+v\n", resp)
}

// MYR
// bankType=ONLINE_BANKING  （在线银行）
// INR
// bankType=UPI  (二维码)
// IDR_NATIVE
// bankType=QRIS  (二维码)
// VND
// bankType=BANK_QR  （越南银行QR码）
func GenDepositRequestDemo() ZPayDepositReq {
	currency := "MYR"

	bankType := ""
	if currency == "VND" {
		bankType = "ONANK_QR"
	} else if currency == "MYR" {
		bankType = "ONLINE_BANKING"
	} else if currency == "INR" {
		bankType = "UPI"
	} else if currency == "IDR_NATIVE" {
		bankType = "QRIS"
	}
	return ZPayDepositReq{
		Currency:                  currency,
		PaymentID:                 "2026_test_ed_9",
		Amount:                    "1000.00",
		BankType:                  bankType,
		CustomerUserId:            "46277",
		CustomerUsername:          "zs",
		CustomerEmail:             "zs@example.com",
		CustomerPhone:             "09123456789",
		CustomerBankName:          "ABBANK",
		CustomerBankHolderName:    "zs",
		CustomerBankAccountNumber: "1234567890",
	}
}
