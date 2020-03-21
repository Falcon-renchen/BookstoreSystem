package main

import (
	"controller"
	"net/http"
)

func main()  {
	//设置处理静态资源，如css和js文件
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static"))))
	//直接去html页面
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages"))))
	//去主页
	http.HandleFunc("/main",controller.GetPageBooksByPrice)
	//去登录页面
	http.HandleFunc("/login",controller.Login)
	//去注册
	http.HandleFunc("/register",controller.Register)
	//去注销
	http.HandleFunc("/logout",controller.Logout)
	//请求验证用户名是否可用
	http.HandleFunc("/checkUserName",controller.CheckUserName)
	//获取带分页的所有图书
	http.HandleFunc("/getPageBooks",controller.GetPageBooks)
	//根据价格查询图书
	http.HandleFunc("/getPageBooksByPrice",controller.GetPageBooksByPrice)
	//添加图书
	//http.HandleFunc("/addBooks",controller.AddBooks)
	//删除图书
	http.HandleFunc("/deleteBooks",controller.DeleteBooks)
	//根据id查询获取图书
	//http.HandleFunc("/")
	//去更新图书的页面
	http.HandleFunc("/toUpdateBookPage",controller.ToUpdateBookPage)
	//修改图书
	http.HandleFunc("/updateOraddBook",controller.UpdateOrAddBooks)
	//添加图书到图书馆
	http.HandleFunc("/addBook2Cart",controller.AddBook2Cart)
	//显示购物车信息
	http.HandleFunc("/getCartInfo",controller.GetCartInfo)
	//清空购物车
	http.HandleFunc("/deleteCart",controller.DeleteCart)
	//删除购物项
	http.HandleFunc("/deleteCartItem",controller.DeleteCartItem)
	//更新购物项
	http.HandleFunc("/updateCartItem",controller.UpdateCartItem)
	//去结账
	http.HandleFunc("/checkout",controller.Checkout)
	//获取所有订单信息
	http.HandleFunc("/getOrders",controller.GetOrders)
	//获取订单详情
	http.HandleFunc("/getOrderInfo",controller.GetOrderInfo)
	//获取我的订单
	http.HandleFunc("/getMyOrder",controller.GetMyOrder)
	//发货
	http.HandleFunc("/sendOrder",controller.SendOrder)
	//收获
	http.HandleFunc("/takeOrder",controller.TakeOrder)


	http.ListenAndServe(":8080",nil)
}
