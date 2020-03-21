package dao

import (
	"model"
	"utils"
)

//向购物车表中插入购物车
func AddCart(cart *model.Cart) error {
	sqlStr := "insert into carts(id,total_count,total_amount,user_id) values(?,?,?,?)"
	_, err := utils.Db.Exec(sqlStr,cart.CartID,cart.GetTotalCount(),cart.GetTotalAmount(),cart.UserID)
	if err!=nil {
		return err
	}
	//获得购物车中的所有购物项
	cartItems := cart.CartItems
	//遍历得到每一个购物项
	for _, cartItem := range cartItems {
		//将购物项添加到购物车中
		AddCartItem(cartItem)
	}
	return nil
}

func GetCartByUserID(userID int) (*model.Cart, error) {
	sqlStr := "select id,total_count,total_amount,user_id from carts where user_id=?"
	row := utils.Db.QueryRow(sqlStr,userID)
	//创建一个购物车
	cart := &model.Cart{}
	err := row.Scan(&cart.CartID,&cart.TotalCount,&cart.TotalAmount,&cart.UserID)
	if err!=nil {
		return nil,err
	}
	//获取当前所有的购物项
	cartItems,_ := GetCartItemByCartID(cart.CartID)
	//将所有的购物项设置到购物车中
	cart.CartItems = cartItems
	return cart,nil
}

//更新购物车中的图书的总数量和总金额
func UpdateCart(cart *model.Cart) error {
	sqlStr := "update carts set total_count=?, total_amount=? where id = ?"
	_, err := utils.Db.Exec(sqlStr,cart.GetTotalCount(),cart.GetTotalAmount(),cart.CartID)
	if err!=nil {
		return err
	}
	return nil
}

//根据购物车的id删除购物车
func DeleteCartByCartID(cartID string) error {
	err := DeleteCartItemsByCartID(cartID)
	if err!=nil {
		return err
	}
	//写sqlStr
	sqlStr := "delete from carts where id=?"
	_, err2 := utils.Db.Exec(sqlStr,cartID)
	if err2!=nil {
		return err2
	}
	return nil
}