package rpc

import (
	"context"

	"github.com/gagliardetto/solana-go"
)

type GetSupplyResult struct {
	RPCContext
	Value *SupplyResult `json:"value"`
}

type SupplyResult struct {
	// TODO: use bin.Uint64 ???
	Total                  uint64             `json:"total"`                  // Total supply in lamports
	Circulating            uint64             `json:"circulating"`            // Circulating supply in lamports
	NonCirculating         uint64             `json:"nonCirculating"`         // Non-circulating supply in lamports
	NonCirculatingAccounts []solana.PublicKey `json:"nonCirculatingAccounts"` // an array of account addresses of non-circulating accounts, as strings
}

// GetSupply returns information about the current supply.
func (cl *Client) GetSupply(
	ctx context.Context,
	commitment CommitmentType,
) (out *GetSupplyResult, err error) {
	params := []interface{}{}
	if commitment != "" {
		params = append(params,
			M{"commitment": commitment},
		)
	}
	err = cl.rpcClient.CallFor(&out, "getSupply", params)
	return
}
