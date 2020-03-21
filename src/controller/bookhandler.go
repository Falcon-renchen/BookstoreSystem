package controller

import (
	"dao"
	"html/template"
	"model"
	"net/http"
	"strconv"
)

////去首页
//func IndexHandler(w http.ResponseWriter, r *http.Request) {
//	//获取页码
//	pageNo := r.FormValue("pageNo")
//	if pageNo == "" {
//		pageNo = "1"
//	}
//	//调用bookdao中获取带分页的图书的函数
//	page, _ := dao.GetPageBooks(pageNo)
//	t := template.Must(template.ParseFiles("views/index.html"))
//	t.Execute(w, page)
//}

/*func GetBooks(w http.ResponseWriter,r *http.Request)  {
	books, _ := dao.GetBooks()

	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
	t.Execute(w,books)
}*/

//带分页的图书信息
func GetPageBooks(w http.ResponseWriter, r *http.Request) {
	//获取页码
	pageNo := r.FormValue("pageNo")
	if pageNo == "" {
		pageNo = "1"
	}
	//调用bookdao中获取带分页的图书的函数
	page, _ := dao.GetPageBooks(pageNo)
	//解析模版
	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
	t.Execute(w,page)
}

/*func AddBooks(w http.ResponseWriter, r *http.Request)  {
	title := r.PostFormValue("title")
	author := r.PostFormValue("author")
	price := r.PostFormValue("price")
	sales := r.PostFormValue("sales")
	stock := r.PostFormValue("stock")

	fPrice, _ := strconv.ParseFloat(price, 64)
	iSales, _ := strconv.ParseInt(sales,10,0)
	iStock, _ := strconv.ParseInt(stock,10,0)

	book := &model.Book{
		Title:   title,
		Author:  author,
		Price:   fPrice,
		Sales:   int(iSales),
		Stock:   int(iStock),
		ImgPath: "/static/img/default.jpg",
	}
	dao.AddBooks(book)
	GetBooks(w,r)
}*/

func DeleteBooks(w http.ResponseWriter, r *http.Request) {
	bookid := r.FormValue("bookId")
	dao.DeleteBooks(bookid)
	GetPageBooks(w,r)
}

func ToUpdateBookPage(w http.ResponseWriter, r *http.Request)  {
	bookID := r.FormValue("bookId")
	book, _ := dao.GetBooksById(bookID)
	if book.ID>0 {
		//如果有图书 更新
		t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
		t.Execute(w,book)
	} else {
		//如果没有图书，添加
		t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
		t.Execute(w,"")
	}

}

func UpdateOrAddBooks(w http.ResponseWriter, r *http.Request)  {
	bookId := r.PostFormValue("bookId")
	title := r.PostFormValue("title")
	author := r.PostFormValue("author")
	price := r.PostFormValue("price")
	sales := r.PostFormValue("sales")
	stock := r.PostFormValue("stock")

	fPrice, _ := strconv.ParseFloat(price, 64)
	iSales, _ := strconv.ParseInt(sales,10,0)
	iStock, _ := strconv.ParseInt(stock,10,0)
	iBookId, _ := strconv.ParseInt(bookId,10,0)

	book := &model.Book{
		ID:		 int(iBookId),
		Title:   title,
		Author:  author,
		Price:   fPrice,
		Sales:   int(iSales),
		Stock:   int(iStock),
		ImgPath: "/static/img/default.jpg",
	}
	if book.ID > 0 {
		dao.UpdateBooks(book)
	} else {
		dao.AddBooks(book)
	}
	GetPageBooks(w,r)
}

//根据价格查询图书
func GetPageBooksByPrice(w http.ResponseWriter, r *http.Request) {
	//获取页码
	pageNo := r.FormValue("pageNo")
	minPrice := r.FormValue("min")
	maxPrice := r.FormValue("max")
	if pageNo == "" {
		pageNo = "1"
	}
	var page *model.Page
	if minPrice=="" && maxPrice=="" {
		//调用bookdao中获取带分页的图书的函数
		page, _ = dao.GetPageBooks(pageNo)
	} else {
		//调用bookdao中获取带分页的图书的函数
		page, _ = dao.GetPageBooksByPrice(pageNo, minPrice, maxPrice)
		page.MinPrice = minPrice
		page.MaxPrice = maxPrice
	}
		//调用IsLogin判断是否已经登录
		flag, session := dao.IsLogin(r)

			if flag {
				//已经登录，设置page中的islogin和username值
				page.IsLogin = true
				page.Username = session.UserName
			}


	//解析模版
	t := template.Must(template.ParseFiles("views/index.html"))
	t.Execute(w,page)
}