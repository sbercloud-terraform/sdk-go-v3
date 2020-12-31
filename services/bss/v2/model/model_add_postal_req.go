/*
 * BSS
 *
 * Business Support System API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

type AddPostalReq struct {
	// |参数名称：收件人姓名| |参数约束及描述：收件人姓名|
	Recipient string `json:"recipient"`
	// |参数名称：省/自治区/直辖市。例如：江苏，不要写成：江苏省| |参数约束及描述：省/自治区/直辖市。例如：江苏，不要写成：江苏省|
	Province string `json:"province"`
	// |参数名称：市/区。例如：南京。| |参数约束及描述：市/区。例如：南京。|
	City string `json:"city"`
	// |参数名称：区。例如：雨花。| |参数约束及描述：区。例如：雨花。|
	District string `json:"district"`
	// |参数名称：邮寄详细地址。| |参数约束及描述：邮寄详细地址。|
	Address string `json:"address"`
	// |参数名称：邮编| |参数约束及描述：邮编|
	Zipcode *string `json:"zipcode,omitempty"`
	// |参数名称：手机号码，不带国家码| |参数约束及描述：手机号码，不带国家码|
	MobilePhone string `json:"mobile_phone"`
	// |参数名称：是否默认地址| |参数约束及描述：是否默认地址，默认为0。1：默认地址0：非默认地址|
	IsDefault *int32 `json:"is_default,omitempty"`
}

func (o AddPostalReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AddPostalReq struct{}"
	}

	return strings.Join([]string{"AddPostalReq", string(data)}, " ")
}
