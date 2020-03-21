package dao

import (
	"fmt"
	"model"
	"testing"
)

func TestBook(t *testing.T) {
	fmt.Println("测试相关函数：：：")

	t.Run("测试更新图书",testUpdateBooks)
}


func testGetBook(t *testing.T)  {
	book, _ := GetBooksById("17")
	fmt.Println("获取的图书是:",book)
}

func testUpdateBooks(t *testing.T)  {
	book := &model.Book{
		Id:      17,
		Title:   "三国演义第二部",
		Author:  "罗贯中",
		Price:   8.88,
		Sales:   100,
		Stock:   1000,
		ImgPath: "/static/img/default.jpg",
	}
	UpdateBooks(book)
}