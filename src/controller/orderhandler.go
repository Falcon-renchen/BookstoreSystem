package controller

import (
	"dao"
	"html/template"
	"model"
	"net/http"
	"time"
	"utils"
)

//去结账
func Checkout(w http.ResponseWriter, r *http.Request)  {
	_, session := dao.IsLogin(r)
	userID := session.UserID
	cart, _ := dao.GetCartByUserID(userID)
	//生成订单号
	orderID := utils.CreateUUID()
	//创建生成订单的时间
	timeStr := time.Now().Format("2020-03-31 15:04:05")
	//创建Order
	order := &model.Order{
		OrderID:     orderID,
		CreateTime:  timeStr,
		TotalCount:  cart.TotalCount,
		TotalAmount: cart.TotalAmount,
		State:       0,
		UserID:      int64(userID),
	}
	//将订单保存到数据库中
	dao.AddOrder(order)
	//保存订单项
	//获取购物车中的购物项
	cartItems := cart.CartItems
	//遍历得到每一个购物项
	for _,v := range cartItems {
		//创建OrderItem
		orderItem := &model.OrderItem{
			Count:       v.Count,
			Amount:      v.Amount,
			Title:       v.Book.Title,
			Author:      v.Book.Author,
			Price:       v.Book.Price,
			ImgPath:     v.Book.ImgPath,
			OrderID:     orderID,
		}
		//将购物项保存到数据库中
		dao.AddOrderItem(orderItem)
		//更新当前购物项中图书的库存和销量
		book := v.Book
		book.Sales = book.Sales + int(v.Count)
		book.Stock = book.Stock - int(v.Count)
		//更新图书的信息
		dao.UpdateBooks(book)
	}
	//清空购物车
	dao.DeleteCartByCartID(cart.CartID)
	//将订单号设置到session中
	session.OrderID = orderID
	//解析模板
	t := template.Must(template.ParseFiles("views/pages/cart/checkout.html"))
	//执行
	t.Execute(w, session)
}

//获取我的订单
func GetMyOrder(w http.ResponseWriter, r *http.Request)  {
	_, session := dao.IsLogin(r)
	userID := session.UserID
	orders, _ := dao.GetMyOrders(userID)
	session.Orders = orders
	t := template.Must(template.ParseFiles("views/pages/order/order.html"))
	t.Execute(w,session)
}

//获取所有订单
func GetOrders(w http.ResponseWriter,r *http.Request) {
	orders, _ := dao.GetOrders()
	t := template.Must(template.ParseFiles("views/pages/order/order_manager.html"))
	t.Execute(w,orders)
}

//获取订单对应的订单项
func GetOrderInfo(w http.ResponseWriter, r *http.Request)  {
	orderID := r.FormValue("orderId")
	orderItems,_ := dao.GetOrderItemsByOrderID(orderID)
	t := template.Must(template.ParseFiles("views/pages/order/order_info.html"))
	t.Execute(w,orderItems)
}

//发货
func SendOrder(w http.ResponseWriter, r *http.Request)  {
	orderID := r.FormValue("orderId")
	dao.UpdateOrderState(orderID,1)
	GetOrders(w,r)
}

//收货
func TakeOrder(w http.ResponseWriter, r *http.Request)  {
	orderID := r.FormValue("orderId")
	dao.UpdateOrderState(orderID,2)
	GetOrders(w,r)
}