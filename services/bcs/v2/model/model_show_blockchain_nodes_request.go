/*
 * BCS
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
type ShowBlockchainNodesRequest struct {
	BlockchainId string `json:"blockchain_id"`
}

func (o ShowBlockchainNodesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowBlockchainNodesRequest struct{}"
	}

	return strings.Join([]string{"ShowBlockchainNodesRequest", string(data)}, " ")
}
