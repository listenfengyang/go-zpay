package go_zpay

import (
	"crypto/tls"
	"fmt"

	jsoniter "github.com/json-iterator/go"
	"github.com/listenfengyang/go-zpay/utils"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cast"
)

func (cli *Client) WithdrawReq(req ZPayWithdrawReq) (*ZPayWithdrawRsp, error) {

	rawURL := cli.Params.WithdrawUrl
	// Convert struct to map for signing（跳过嵌套的 AdditionalParams，单独处理）
	params := make(map[string]string)
	mapstructure.Decode(req, &params)

	params["merchantCode"] = cast.ToString(cli.Params.MerchantCode)
	params["merchantKey"] = cast.ToString(cli.Params.MerchantKey)
	params["callbackUrl"] = cast.ToString(cli.Params.WithdrawResponseUrl)

	// additionalParams：AUD / THB 币种时需要，序列化为 JSON 字符串后作为独立字段传递
	// AUD: {"way_code":"BSB","bsb_code":"042056","pay_id":"-"}
	//   or {"way_code":"PAYID","bsb_code":"-","pay_id":"test@payid"}
	// THB: {"customer_user_id":"484799","customer_username":"user_123"}
	if req.Currency == "THB_FOREX" && req.AdditionalParams != (AdditionalParamsObj{}) {
		additionalJSON, err := jsoniter.ConfigCompatibleWithStandardLibrary.MarshalToString(req.AdditionalParams)
		if err == nil && additionalJSON != "{}" {
			params["additionalParams"] = additionalJSON
		}
	}

	// Generate signature
	signStr, _ := utils.SignWithdraw(params)
	params["signature"] = signStr

	var result ZPayWithdrawRsp
	resp2, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetFormData(params).
		SetBody(params).
		SetHeaders(getHeaders()).
		SetDebug(cli.debugMode).
		SetResult(&result).
		SetError(&result).
		Post(rawURL)

	restLog, _ := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(utils.GetRestyLog(resp2))
	cli.logger.Infof("PSPResty#zpay#withdraw->%s", string(restLog))

	if err != nil {
		return nil, err
	}

	if resp2.StatusCode() != 200 {
		return nil, fmt.Errorf("status code: %d, body:%s", resp2.StatusCode(), resp2.Body())
	}

	// 服务端始终返回 HTTP 200，通过业务状态码判断成功与否
	// SetError 绑定了同一个 result 指针，resp2.Error() 永远非 nil，不能用于错误判断
	if result.Status != 200 {
		return nil, fmt.Errorf("err:&%+v, body:%s", result, resp2.Body())
	}

	return &result, nil
}
