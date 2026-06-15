package go_zpay

import (
	"crypto/tls"
	"fmt"

	jsoniter "github.com/json-iterator/go"
	"github.com/listenfengyang/go-zpay/utils"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cast"
)

// 下单
func (cli *Client) Deposit(req ZPayDepositReq) (*ZPayDepositRsp, error) {

	rawURL := cli.Params.DepositUrl

	var params map[string]string
	mapstructure.Decode(req, &params)

	//补充字段
	params["merchantCode"] = cast.ToString(cli.Params.MerchantCode)
	params["merchantKey"] = cast.ToString(cli.Params.MerchantKey)
	params["responseURL"] = cast.ToString(cli.Params.DepositResponseUrl)
	params["remark"] = "prod"

	// Generate signature
	signStr, _ := utils.Sign(params)
	params["signature"] = signStr
	// fmt.Println(params)
	var result ZPayDepositRsp

	resp2, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetBody(params).
		SetHeaders(getHeaders()).
		SetDebug(cli.debugMode).
		SetResult(&result).
		SetError(&result).
		Post(rawURL)

	restLog, _ := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(utils.GetRestyLog(resp2))
	cli.logger.Infof("PSPResty#zpay#deposit->%s", string(restLog))

	if err != nil {
		return nil, err
	}

	if resp2.StatusCode() != 200 {
		return nil, fmt.Errorf("status code: %d, body:%s", resp2.StatusCode(), resp2.Body())
	}

	// 服务端 Content-Type 可能为 text/html，Resty 不会自动反序列化，手动解析 body 兜底
	if result.Status == 0 {
		_ = jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(resp2.Body(), &result)
	}
	cli.logger.Infof("PSPResty#zpay#deposit result: %+v", result)

	return &result, nil
}
