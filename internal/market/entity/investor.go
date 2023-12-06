package entity

type Investor struct {
	ID                    string
	Name                  string
	InvestorAssetPosition []*InvestorAssetPosition
}

func NewInvestor(id string) *Investor {
	return &Investor{
		ID:                    id,
		InvestorAssetPosition: []*InvestorAssetPosition{},
	}
}

func (i *Investor) AddInvestorAssetPosition(investorAssetPosition *InvestorAssetPosition) {
	i.InvestorAssetPosition = append(i.InvestorAssetPosition, investorAssetPosition)
}

func (i *Investor) UpdateAssetPosition(assetId string, numShares int) {
	investorAssetPosition := i.getInvestorAssetPosition(assetId)

	if investorAssetPosition == nil {
		i.InvestorAssetPosition = append(i.InvestorAssetPosition, NewInvestorAssetPosition(assetId, numShares))
	} else {
		investorAssetPosition.NumShares += numShares
	}
}

func (i *Investor) getInvestorAssetPosition(assetID string) *InvestorAssetPosition {
	for _, currentAssetPosition := range i.InvestorAssetPosition {
		if currentAssetPosition.AssetID == assetID {
			return currentAssetPosition
		}
	}
	return nil
}

type InvestorAssetPosition struct {
	AssetID   string
	NumShares int
}

func NewInvestorAssetPosition(assetID string, numShares int) *InvestorAssetPosition {
	return &InvestorAssetPosition{
		AssetID:   assetID,
		NumShares: numShares,
	}
}
