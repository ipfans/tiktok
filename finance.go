package tiktok

import "context"

const (
	_SearchSettlementsPATH   = "/api/finance/settlements/search"
	_GetOrderSettlementsPATH = "/api/finance/order/settlements"
	_SearchTransactionsPATH  = "/api/finance/transactions/search"
)

// SearchSettlements search settlements. Default PageSize is 100.
func (c *Client) SearchSettlements(ctx context.Context, p Param, query SearchSettlementsRequest) (list SettlementsList, err error) {
	param, err := c.params(p)
	if err != nil {
		return
	}
	if err = c.validate.Struct(query); err != nil {
		return
	}
	err = c.Post(ctx, _SearchSettlementsPATH, param, query, &list)
	return
}

// GetOrderSettlements get order settlements.
func (c *Client) GetOrderSettlements(ctx context.Context, p Param, ordersID string) (list SettlementsList, err error) {
	param, err := c.params(p)
	if err != nil {
		return
	}
	param.Set("order_id", ordersID)
	err = c.Get(ctx, _GetOrderSettlementsPATH, param, &list)
	return
}

// SearchTransactions search your seller account's transactions within a certain timeframe. Default PageSize is 100.
func (c *Client) SearchTransactions(ctx context.Context, p Param, query SearchTransactionsRequest) (list TransactionsList, err error) {
	param, err := c.params(p)
	if err != nil {
		return
	}
	if err = c.validate.Struct(query); err != nil {
		return
	}
	err = c.Post(ctx, _SearchTransactionsPATH, param, query, &list)
	return
}
