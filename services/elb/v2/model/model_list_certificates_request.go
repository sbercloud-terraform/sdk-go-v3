/*
 * ELB
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
type ListCertificatesRequest struct {
	Limit       *int32  `json:"limit,omitempty"`
	Marker      *string `json:"marker,omitempty"`
	PageReverse *string `json:"page_reverse,omitempty"`
	Id          *string `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Type        *string `json:"type,omitempty"`
	Domain      *string `json:"domain,omitempty"`
	PrivateKey  *string `json:"private_key,omitempty"`
	Certificate *string `json:"certificate,omitempty"`
}

func (o ListCertificatesRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListCertificatesRequest", string(data)}, " ")
}
