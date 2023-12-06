package entity

type Investor struct {
	ID                     string
	Name                   string
	InvestorAssetPositions []*AssetPosition
}

func NewInvestor(id string) *Investor {
	return &Investor{
		ID:                     id,
		InvestorAssetPositions: []*AssetPosition{},
	}
}

func (i *Investor) AddInvestorAssetPosition(investorAssetPosition *AssetPosition) {
	i.InvestorAssetPositions = append(i.InvestorAssetPositions, investorAssetPosition)
}

func (i *Investor) UpdateInvestorAssetPosition(assetId string, numShares int) {
	investorAssetPosition := i.getInvestorAssetPosition(assetId)

	if investorAssetPosition == nil {
		i.InvestorAssetPositions = append(i.InvestorAssetPositions, NewAssetPosition(assetId, numShares))
	} else {
		investorAssetPosition.NumShares += numShares
	}
}

func (i *Investor) getInvestorAssetPosition(assetID string) *AssetPosition {
	for _, currentAssetPosition := range i.InvestorAssetPositions {
		if currentAssetPosition.AssetID == assetID {
			return currentAssetPosition
		}
	}
	return nil
}

type AssetPosition struct {
	AssetID   string
	NumShares int
}

func NewAssetPosition(assetID string, numShares int) *AssetPosition {
	return &AssetPosition{
		AssetID:   assetID,
		NumShares: numShares,
	}
}
