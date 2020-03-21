package dao

import (
	"model"
	"strconv"
	"utils"
)
//获取所有图书
func GetBooks() ([]*model.Book, error) {
	sqlStr := "select id,title,author,price,sales,stock,img_path from books"

	row, err := utils.Db.Query(sqlStr)
	if err!=nil {
		return nil,err
	}
	var books []*model.Book

	for row.Next() {
		book := &model.Book{}
		row.Scan(&book.ID,&book.Title,&book.Author,&book.Price,&book.Sales,&book.Stock,&book.ImgPath)
		books = append(books,book)
	}
	return books,nil
}

//添加图书
func AddBooks(b *model.Book) error {
	sqlStr := "insert into books(title,author,price,sales,stock,img_path) values(?,?,?,?,?,?)"

	_ , err := utils.Db.Exec(sqlStr,b.Title,b.Author,b.Price,b.Sales,b.Stock,b.ImgPath)
	if err!=nil {
		return err
	}
	return nil
}

//删除图书
func DeleteBooks(bookID string) error {
	sqlStr := "delete from books where id=?"
	_, err := utils.Db.Exec(sqlStr,bookID)
	if err!=nil {
		return err
	}
	return nil
}

//查询图书
func GetBooksById(bookId string) (*model.Book, error) {
	sqlStr := "select id,title,author,price,sales,stock,img_path from books where id=?"
	row := utils.Db.QueryRow(sqlStr, bookId)
	book := &model.Book{}
	row.Scan(&book.ID,&book.Title,&book.Author,&book.Price,&book.Sales,&book.Stock,&book.ImgPath)
	return book,nil
}

//修改图书
func UpdateBooks(book *model.Book) error {
	sqlstr := "update books set title=?,author=?,price=?,sales=?,stock=? where id=?"
	_, err := utils.Db.Exec(sqlstr,book.Title,book.Author,book.Price,book.Sales,book.Stock,book.ID)
	if err!=nil {
		return err
	}
	return nil
}

//带分页的图书
func GetPageBooks(pageNo string) (*model.Page, error) {
	//转int64
	iPageNo, _ := strconv.ParseInt(pageNo,10,64)
	//获取数据库中的总记录数
	sqlStr := "select count(*) from books"
	//设置一个变量接受总记录数
	var totalRecord int64
	//执行
	row := utils.Db.QueryRow(sqlStr)
	row.Scan(&totalRecord)
	//分页   每页4条数据
	var pageSize int64 = 4
	//设置一个变量接收总记录数
	var totalPageNo int64
	if totalRecord % pageSize == 0 {
		totalPageNo = totalRecord/pageSize
	} else {
		totalPageNo = totalRecord/pageSize+1
	}
	//分页操作
	sqlStr2 := "select id,title,author,price,sales,stock,img_path from books limit ?,?"
	//执行
	rows , err := utils.Db.Query(sqlStr2,(iPageNo-1)*pageSize,pageSize)
	if err!=nil {
		return nil,err
	}
	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		rows.Scan(&book.ID,&book.Title,&book.Author,&book.Price,&book.Sales,&book.Stock,&book.ImgPath)
		books = append(books, book)
	}
	//创建page
	page := &model.Page{
		Books:        books,
		PageNo:      iPageNo,
		PageSize:    pageSize,
		TotalPageNo: totalPageNo,
		TotalRecord: totalRecord,
	}
	return page,nil
}

//获取带价格的分页图书
func GetPageBooksByPrice(pageNo string, minPrice string, maxPrice string) (*model.Page, error) {
	//转int64
	iPageNo, _ := strconv.ParseInt(pageNo,10,64)
	//获取数据库中的总记录数
	sqlStr := "select count(*) from books where price between ? and ?"
	//设置一个变量接受总记录数
	var totalRecord int64
	//执行
	row := utils.Db.QueryRow(sqlStr, minPrice, maxPrice)
	row.Scan(&totalRecord)
	//分页   每页4条数据
	var pageSize int64 = 4
	//设置一个变量接收总记录数
	var totalPageNo int64
	if totalRecord % pageSize == 0 {
		totalPageNo = totalRecord/pageSize
	} else {
		totalPageNo = totalRecord/pageSize+1
	}
	//分页操作
	sqlStr2 := "select id,title,author,price,sales,stock,img_path from books where price between ? and ? limit ?,?"
	//执行
	rows , err := utils.Db.Query(sqlStr2,minPrice,maxPrice,(iPageNo-1)*pageSize,pageSize)
	if err!=nil {
		return nil,err
	}
	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		rows.Scan(&book.ID,&book.Title,&book.Author,&book.Price,&book.Sales,&book.Stock,&book.ImgPath)
		books = append(books, book)
	}
	//创建page
	page := &model.Page{
		Books:        books,
		PageNo:      iPageNo,
		PageSize:    pageSize,
		TotalPageNo: totalPageNo,
		TotalRecord: totalRecord,
	}
	return page,nil
}