package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/spf13/cast"
)

func Sign(params map[string]string) (string, error) {

	// 1. 定义参与签名的key
	signKeyList := []string{"merchantCode", "merchantKey", "currency", "paymentID", "responseURL", "amount"}

	// 2. Build sign string
	var sb strings.Builder
	for _, k := range signKeyList {
		v := cast.ToString(params[k])
		// 去掉小数点
		if k == "amount" {
			v = strings.ReplaceAll(v, ".", "")
		}
		//只有非空才可以参与签名
		sb.WriteString(fmt.Sprintf("%s&", v))
	}
	signStr := sb.String()
	signStr = strings.Trim(signStr, "&")

	fmt.Printf("[rawString]%s\n", signStr)

	// 2. Generate MD5
	hash := sha256.Sum256([]byte(signStr))
	signResult := hex.EncodeToString(hash[:])

	//fmt.Printf("验签str: %s\n结果: %s\n", signStr, signResult)

	fmt.Printf("[rawUpString]%s\n", strings.ToUpper(signResult))

	return strings.ToUpper(signResult), nil
}

func Verify(params map[string]string, signKey string) (bool, error) {
	// Check if signature exists in params
	signature, exists := params["sign"]
	if !exists {
		return false, nil
	}

	// Remove signature from params for verification
	delete(params, "sign")

	// Generate current signature
	currentSignature, err := Sign(params)
	if err != nil {
		return false, fmt.Errorf("signature generation failed: %w", err)
	}

	// Compare signatures
	return signature == currentSignature, nil
}

// 入金&出金回调-成功-验签
func VerifyCallback(params map[string]string) bool {
	signature := params["signature"]

	// 1. 定义参与签名的key
	signKeyList := []string{"merchantKey", "currency", "transaction_id", "reference_code", "amount"}

	// 2. Build sign string
	var sb strings.Builder
	for _, k := range signKeyList {
		v := cast.ToString(params[k])
		// 去掉小数点
		if k == "amount" {
			v = strings.ReplaceAll(v, ".", "")
		}
		//只有非空才可以参与签名
		sb.WriteString(fmt.Sprintf("%s&", v))
	}
	signStr := sb.String()
	signStr = strings.Trim(signStr, "&")

	fmt.Printf("[rawString]%s\n", signStr)

	hash := sha256.Sum256([]byte(signStr))

	fmt.Printf("signature: %s\n", signature)
	fmt.Printf("sha256 sign: %s\n", strings.ToUpper(hex.EncodeToString(hash[:])))

	return signature == strings.ToUpper(hex.EncodeToString(hash[:]))
}

// 出金
func SignWithdraw(params map[string]string) (string, error) {
	// 1. 定义参与签名的key
	signKeyList := []string{"merchantCode", "merchantKey", "currency", "merchantRefNo", "callbackUrl", "amount"}

	// 2. Build sign string
	var sb strings.Builder
	for _, k := range signKeyList {
		v := cast.ToString(params[k])
		// 去掉小数点
		if k == "amount" {
			v = strings.ReplaceAll(v, ".", "")
		}
		//只有非空才可以参与签名
		sb.WriteString(fmt.Sprintf("%s&", v))
	}
	signStr := sb.String()
	signStr = strings.Trim(signStr, "&")

	fmt.Printf("[rawString]%s\n", signStr)

	// 2. Generate MD5
	hash := sha256.Sum256([]byte(signStr))
	signResult := hex.EncodeToString(hash[:])

	//fmt.Printf("验签str: %s\n结果: %s\n", signStr, signResult)

	fmt.Printf("[rawUpString]%s\n", strings.ToUpper(signResult))

	return strings.ToUpper(signResult), nil
}

func VerifySignWithdraw(params map[string]string) (bool, error) {
	// Check if signature exists in params
	signature, exists := params["sign"]
	if !exists {
		return false, nil
	}

	// Remove signature from params for verification
	delete(params, "sign")

	// Generate current signature
	currentSignature, err := SignWithdraw(params)
	if err != nil {
		return false, fmt.Errorf("signature generation failed: %w", err)
	}

	// Compare signatures
	return signature == currentSignature, nil
}
