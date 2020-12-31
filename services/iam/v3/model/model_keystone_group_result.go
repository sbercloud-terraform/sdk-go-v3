/*
 * IAM
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

//
type KeystoneGroupResult struct {
	// 用户组描述信息。
	Description string `json:"description"`
	// 用户组ID。
	Id string `json:"id"`
	// 用户组所属账号ID。
	DomainId string `json:"domain_id"`
	// 用户组名称。
	Name  string `json:"name"`
	Links *Links `json:"links"`
	// 用户组创建时间。
	CreateTime int64 `json:"create_time"`
}

func (o KeystoneGroupResult) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneGroupResult struct{}"
	}

	return strings.Join([]string{"KeystoneGroupResult", string(data)}, " ")
}
