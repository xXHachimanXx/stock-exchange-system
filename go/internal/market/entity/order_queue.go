package entity

type OrderQueue struct {
	Orders []*Order
}

func (oq *OrderQueue) Less(i, j int) bool {
	return oq.Orders[i].Price < oq.Orders[j].Price
}

func (oq *OrderQueue) Swap(i, j int) {
	oq.Orders[i], oq.Orders[j] = oq.Orders[j], oq.Orders[i]
}

func (oq *OrderQueue) Len() int {
	return len(oq.Orders)
}

func (oq *OrderQueue) Push(order any) {
	oq.Orders = append(oq.Orders, order.(*Order))
}

func (oq *OrderQueue) Pop() any {
	oldOrders := oq.Orders
	oldOrdersSize := len(oldOrders)
	item := oldOrders[oldOrdersSize-1]
	oq.Orders = oldOrders[0 : oldOrdersSize-1]

	return item
}

func NewOrderQueue() *OrderQueue {
	return &OrderQueue{}
}
