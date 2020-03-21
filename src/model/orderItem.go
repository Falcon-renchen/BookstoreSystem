package model

type OrderItem struct {
	 OrderItemID int64 //订单项的ID
	 Count int64 //订单项中图书的总数量
	 Amount float64 // 订单项中图书的总金额
	 Title string 	//订单项中图书的书名
	 Author string  //订单项图书的作者
	 Price float64	//订单项中图书的价格
	 ImgPath string //图书的封面
	 OrderID string //订单项所属的订单
}