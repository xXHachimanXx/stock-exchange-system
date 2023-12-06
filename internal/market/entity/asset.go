package entity

type Asset struct {
	ID           string
	Name         string
	MarketVolume int
}

func NewAsst(id string, name string, marketVolume int) *Asset {
	return &Asset{
		ID:           id,
		Name:         name,
		MarketVolume: marketVolume,
	}
}
