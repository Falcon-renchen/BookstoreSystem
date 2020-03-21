package dao

import (
	"model"
	"utils"
)

//向数据库中添加订单项
func AddOrderItem(orderItem *model.OrderItem) error {
	sqlStr := "insert into order_items(counts,amount,title,author,price,img_path,order_id) values(?,?,?,?,?,?,?)"
	_, err := utils.Db.Exec(sqlStr,orderItem.Count,orderItem.Amount,orderItem.Title,orderItem.Author,orderItem.Price,orderItem.ImgPath,orderItem.OrderID)
	if err!=nil {
		return err
	}
	return nil
}

//根据订单号查询该订单的所有信息
func GetOrderItemsByOrderID(orderID string) ([]*model.OrderItem,error) {
	sqlStr := "select id,counts,amount,title,author,price,img_path,order_id from order_items where order_id=?"
	row, err := utils.Db.Query(sqlStr,orderID)
	if err!=nil {
		return nil,err
	}
	var orderItems []*model.OrderItem
	for row.Next() {
		orderItem := &model.OrderItem{}
		row.Scan(&orderItem.OrderItemID,&orderItem.Count,&orderItem.Amount,&orderItem.Title,&orderItem.Author,&orderItem.Price,&orderItem.ImgPath,&orderItem.OrderID)
		orderItems = append(orderItems, orderItem)
	}
	return orderItems,nil
}