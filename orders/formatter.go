package orders

import "time"

type OrderFormatter struct {
	ID           int             `json:"id"`
	CustomerName string          `json:"customerName"`
	OrderedAt    time.Time       `json:"orderedAt"`
	Items        []ItemFormatter `json:"items"`
}

type ItemFormatter struct {
	ID          int    `json:"id"`
	Code        string `json:"code"`
	Description string `json:"description"`
	Qty         int    `json:"qty"`
}

func FormatOrder(order Order) OrderFormatter {
	formatter := OrderFormatter{}
	formatter.ID = order.ID
	formatter.CustomerName = order.CustomerName
	formatter.OrderedAt = order.CreatedAt

	itemsFormatter := []ItemFormatter{}
	for _, item := range order.Items {
		itemFormatter := ItemFormatter{}
		itemFormatter.ID = item.ID
		itemFormatter.Code = item.Code
		itemFormatter.Description = item.Description
		itemFormatter.Qty = item.Qty
		itemsFormatter = append(itemsFormatter, itemFormatter)
	}
	formatter.Items = itemsFormatter

	return formatter
}

func FormatOrders(orders []Order) []OrderFormatter {
	if len(orders) == 0 {
		return []OrderFormatter{}
	}

	ordersFormatter := []OrderFormatter{}
	for _, order := range orders {
		ordersFormatter = append(ordersFormatter, FormatOrder(order))
	}

	return ordersFormatter
}
