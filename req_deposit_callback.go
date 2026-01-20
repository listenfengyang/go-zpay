package go_zpay

import (
	"encoding/json"
	"errors"

	"github.com/listenfengyang/go-zpay/utils"
	"github.com/mitchellh/mapstructure"
)

// 充值-成功回调
func (cli *Client) DepositCallback(req ZPayDepositCallbackReq, processor func(ZPayDepositCallbackReq) error) error {
	//验证签名
	var params map[string]string
	mapstructure.Decode(req, &params)
	params["merchantKey"] = cli.Params.MerchantKey

	if !utils.VerifyCallback(params) {
		//签名校验失败
		reqJson, _ := json.Marshal(req)
		cli.logger.Errorf("zPay deposit back verify fail, req: %s", string(reqJson))
		return errors.New("sign verify error")
	}

	//开始处理
	return processor(req)
}
