package model

type Cart struct {
	CartID string			//购物车的id
	CartItems []*CartItem	//购物车中的所有购物项
	TotalCount int64		//购物车中所有图书的总数量
	TotalAmount float64		//购物车中所有图书的总价格
	UserID int			//当前购物车所属的用户
}

//购物车图书总数量
func (cart *Cart) GetTotalCount() int64 {
	var totalcount int64
	for _, v := range cart.CartItems {
		totalcount = totalcount + v.Count
	}
	return totalcount
}

//购物车图书总价格
func (cart *Cart) GetTotalAmount() float64 {
	var totalamount float64
	for _, v := range cart.CartItems {
		totalamount = totalamount + v.GetAmount()
	}
	return totalamount
}

