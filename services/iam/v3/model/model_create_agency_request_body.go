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
type CreateAgencyRequestBody struct {
	Agency *CreateAgencyOption `json:"agency"`
}

func (o CreateAgencyRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateAgencyRequestBody struct{}"
	}

	return strings.Join([]string{"CreateAgencyRequestBody", string(data)}, " ")
}
