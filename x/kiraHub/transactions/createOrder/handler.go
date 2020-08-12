package createOrder

import sdk "github.com/KiraCore/cosmos-sdk/types"

func HandleMessage(context sdk.Context, keeper Keeper, message Message) (*sdk.Result, error) {
	keeper.CreateOrder(context, message.OrderBookID, message.OrderType, message.Amount, message.LimitPrice, message.ExpiryTime, message.Curator)
	return &sdk.Result{}, nil
}