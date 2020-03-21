package model

type Order struct {
	OrderID string //订单号
	CreateTime string //生成订单的时间
	TotalCount int64 //订单中图书的总数量
	TotalAmount float64 //订单中图书的总金额
	State int64 //未发货0  已发货1  交易成功2
	UserID int64 //订单所属的用户
}

//未发货
func (order *Order) NoSend() bool {
	return order.State == 0
}

//已发货
func (order *Order) SendComplete() bool {
	return order.State == 1
}

//交易完成
func (order *Order) Complete() bool {
	return order.State == 2
}