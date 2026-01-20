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

func GenDepositRequestDemo() ZPayDepositReq {
	return ZPayDepositReq{
		Currency:  "VND",
		PaymentID: "22515161369",
		Amount:    "188000.00",
		BankType:  "",
	}
}
