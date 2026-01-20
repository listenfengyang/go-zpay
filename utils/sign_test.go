package utils

import (
	"fmt"
	"testing"
)

func TestSign(t *testing.T) {

	params := map[string]string{
		"coinName": "USDT",
		"orderId":  "12345678910",
		"protocol": "ERC20",
	}

	signStr, _ := Sign(params)

	fmt.Println(signStr)

}
