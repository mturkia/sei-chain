package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sei-protocol/sei-chain/x/dex/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetOrderById(c context.Context, req *types.QueryGetOrderByIdRequest) (*types.QueryGetOrderByIdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)
	orders := k.GetOrdersByIds(ctx, req.ContractAddr, []uint64{req.Id})
	if order, ok := orders[req.Id]; !ok {
		return &types.QueryGetOrderByIdResponse{}, status.Error(codes.NotFound, "order not found")
	} else {
		return &types.QueryGetOrderByIdResponse{
			Order: &order,
		}, nil
	}
}

func (k Keeper) GetOrders(c context.Context, req *types.QueryGetOrdersRequest) (*types.QueryGetOrdersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)
	activeOrderIds := k.GetAccountActiveOrders(ctx, req.ContractAddr, req.Account)
	orders := k.GetOrdersByIds(ctx, req.ContractAddr, activeOrderIds.Ids)
	response := &types.QueryGetOrdersResponse{Orders: []*types.Order{}}
	for _, order := range orders {
		response.Orders = append(response.Orders, &order)
	}
	return response, nil
}