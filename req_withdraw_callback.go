package go_zpay

import (
	"encoding/json"
	"errors"

	"github.com/listenfengyang/go-zpay/utils"
	"github.com/mitchellh/mapstructure"
)

// 出金-成功回调
func (cli *Client) WithdrawCallback(req ZPayWithdrawCallbackReq, processor func(req ZPayWithdrawCallbackReq) error) error {
	//验证签名
	var params map[string]string
	mapstructure.Decode(req, &params)
	params["merchantKey"] = cli.Params.MerchantKey

	// Verify signature
	flag := utils.VerifyCallback(params)
	if !flag {
		//签名校验失败
		reqJson, _ := json.Marshal(req)
		cli.logger.Errorf("zPay successfull back verify fail, req: %s", string(reqJson))
		return errors.New("sign verify error")
	}

	//开始处理
	return processor(req)
}
