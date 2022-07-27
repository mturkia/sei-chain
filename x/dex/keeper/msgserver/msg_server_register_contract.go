package msgserver

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sei-protocol/sei-chain/utils"
	"github.com/sei-protocol/sei-chain/x/dex/contract"
	"github.com/sei-protocol/sei-chain/x/dex/types"
)

func (k msgServer) RegisterContract(goCtx context.Context, msg *types.MsgRegisterContract) (*types.MsgRegisterContractResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// TODO: add validation such that only the user who stored the code can register contract

	if msg.Contract.DependentContractAddrs != nil {
		for _, contractAddr := range msg.Contract.DependentContractAddrs {
			contractInfo, err := k.GetContract(ctx, contractAddr)
			if err != nil {
				// validate that all dependency contracts exist
				return &types.MsgRegisterContractResponse{}, err
			}
			// update incoming paths for dependency contracts
			contractInfo.NumIncomingPaths++
			if err := k.SetContract(ctx, &contractInfo); err != nil {
				return &types.MsgRegisterContractResponse{}, err
			}
		}
	}

	// set incoming paths for current contract
	newContract := msg.Contract
	allContractInfo := k.GetAllContractInfo(ctx)
	for _, contractInfo := range allContractInfo {
		if contractInfo.DependentContractAddrs == nil {
			continue
		}
		dependencies := utils.NewStringSet(contractInfo.DependentContractAddrs)
		if dependencies.Contains(msg.Contract.ContractAddr) {
			newContract.NumIncomingPaths++
		}
	}

	// always override contract info so that it can be updated
	if err := k.SetContract(ctx, newContract); err != nil {
		return &types.MsgRegisterContractResponse{}, err
	}
	allContractInfo = append(allContractInfo, *newContract)

	if _, err := contract.TopologicalSortContractInfo(allContractInfo); err != nil {
		return &types.MsgRegisterContractResponse{}, err
	}

	return &types.MsgRegisterContractResponse{}, nil
}
