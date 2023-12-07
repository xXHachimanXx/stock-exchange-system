package transformer

import (
	"github.com/xXHachimanXx/stock-exchange-system/go/internal/market/dto"
	"github.com/xXHachimanXx/stock-exchange-system/go/internal/market/entity"
)

func TransformInput(input dto.TradeInputDTO) *entity.Order {
	asset := entity.NewAsset(input.AssetID, input.AssetID, 1000)
	investor := entity.NewInvestor(input.InvestorID)
	order := entity.NewOrder(input.OrderID, investor, asset, input.Shares, input.Price, input.OrderType)

	if input.CurrentShares > 0 {
		assetPosition := entity.NewAssetPosition(input.AssetID, input.CurrentShares)
		investor.AddInvestorAssetPosition(assetPosition)
	}

	return order
}

func TransformOutput(order *entity.Order) *dto.OrderOutputDTO {
	output := &dto.OrderOutputDTO{
		OrderID:    order.ID,
		InvestorID: order.Investor.ID,
		AssetID:    order.Asset.ID,
		OrderType:  order.OrderType,
		Status:     order.Status,
		Partial:    order.PendingShares,
		Shares:     order.Shares,
	}

	var transactionsOutputDTO []*dto.TransactionOutputDTO
	for _, t := range order.Transactions {
		transactionOutputDTO := &dto.TransactionOutputDTO{
			TransactionID: t.ID,
			BuyerID:       t.BuyingOrder.Investor.ID,
			SellerID:      t.SellingOrder.Investor.ID,
			AssetID:       t.SellingOrder.Asset.ID,
			Price:         t.Price,
			Shares:        t.SellingOrder.Shares - t.SellingOrder.PendingShares,
		}
		transactionsOutputDTO = append(transactionsOutputDTO, transactionOutputDTO)
	}

	output.TransactionsOutput = transactionsOutputDTO

	return output
}
