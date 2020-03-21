package model

//单个购物项
type CartItem struct {
	CartItemID int64 //购物车中单个购物项的id
	Book *Book	//购物项中的书籍信息
	Count int64	//数量
	Amount float64 //图书金额小计
	CartID string 	//当前购物项属于哪个购物车
}

//获取当前购物车中单个图书的总价格
func (cartItem *CartItem) GetAmount() float64 {
	price := cartItem.Book.Price
	return float64(cartItem.Count) * price
}