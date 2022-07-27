package msgserver_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/sei-protocol/sei-chain/testutil/keeper"
	"github.com/sei-protocol/sei-chain/x/dex/keeper/msgserver"
	"github.com/sei-protocol/sei-chain/x/dex/types"
	"github.com/stretchr/testify/require"
)

func TestRegisterContract(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	server := msgserver.NewMsgServerImpl(*keeper, nil)
	contractInfo := types.ContractInfo{
		CodeId:       1,
		ContractAddr: keepertest.TestContract,
	}
	server.RegisterContract(wctx, &types.MsgRegisterContract{
		Creator:  keepertest.TestAccount,
		Contract: &contractInfo,
	})
	storedContracts := keeper.GetAllContractInfo(ctx)
	require.Equal(t, 1, len(storedContracts))
	require.Nil(t, storedContracts[0].DependentContractAddrs)

	contractInfo.DependentContractAddrs = []string{"TEST2"}
	_, err := server.RegisterContract(wctx, &types.MsgRegisterContract{
		Creator:  keepertest.TestAccount,
		Contract: &contractInfo,
	})
	require.NotNil(t, err)
	storedContracts = keeper.GetAllContractInfo(ctx)
	require.Equal(t, 1, len(storedContracts))
}

func TestRegisterContractCircularDependency(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	server := msgserver.NewMsgServerImpl(*keeper, nil)
	contractInfo1 := types.ContractInfo{
		CodeId:       1,
		ContractAddr: "test1",
	}
	server.RegisterContract(wctx, &types.MsgRegisterContract{
		Creator:  keepertest.TestAccount,
		Contract: &contractInfo1,
	})
	storedContracts := keeper.GetAllContractInfo(ctx)
	require.Equal(t, 1, len(storedContracts))

	// This contract should fail to be registered because it causes a
	// circular dependency
	contractInfo2 := types.ContractInfo{
		CodeId:                 2,
		ContractAddr:           "test2",
		DependentContractAddrs: []string{"test1"},
	}
	server.RegisterContract(wctx, &types.MsgRegisterContract{
		Creator:  keepertest.TestAccount,
		Contract: &contractInfo2,
	})
	storedContracts = keeper.GetAllContractInfo(ctx)
	require.Equal(t, 2, len(storedContracts))

	contractInfo1.DependentContractAddrs = []string{"test2"}
	_, err := server.RegisterContract(wctx, &types.MsgRegisterContract{
		Creator:  keepertest.TestAccount,
		Contract: &contractInfo1,
	})
	require.NotNil(t, err)
}
