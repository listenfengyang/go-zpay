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
	// 2. Convert struct to map for signing
	var params map[string]string
	mapstructure.Decode(req, &params)

	// b, _ := json.Marshal(req)
	// params["data"] = cast.ToString(b)
	params["merchantCode"] = cast.ToString(cli.Params.MerchantCode)
	params["merchantKey"] = cast.ToString(cli.Params.MerchantKey)
	params["callbackUrl"] = cast.ToString(cli.Params.WithdrawResponseUrl)

	// Generate signature
	signStr, _ := utils.SignWithdraw(params)
	params["signature"] = signStr
	// fmt.Println(params)

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
		//反序列化错误会在此捕捉
		return nil, fmt.Errorf("status code: %d", resp2.StatusCode())
	}

	if resp2.Error() != nil {
		//反序列化错误会在此捕捉
		return nil, fmt.Errorf("%v, body:%s", resp2.Error(), resp2.Body())
	}

	return &result, nil
}
