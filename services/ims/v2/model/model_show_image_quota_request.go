/*
 * IMS
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowImageQuotaRequest struct {
}

func (o ShowImageQuotaRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowImageQuotaRequest struct{}"
	}

	return strings.Join([]string{"ShowImageQuotaRequest", string(data)}, " ")
}
