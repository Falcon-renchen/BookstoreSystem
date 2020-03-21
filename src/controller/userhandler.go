package controller

import (
	"dao"
	"html/template"
	"model"
	"net/http"
	"time"
	"utils"
)

func Login(w http.ResponseWriter, r *http.Request) {
	//判断是否已经登录
	flag , _ := dao.IsLogin(r)
	if flag {
		//已经登录
		//去首页
		GetPageBooksByPrice(w,r)
	} else {
		username := r.PostFormValue("username")
		password := r.PostFormValue("password")

		user, _ := dao.CheckUsernameAndPwd(username,password)

		if user.ID > 0 {
			//生成uuid作为session的id
			uuid := utils.CreateUUID()
			//创建一个session
			sess := &model.Session{
				SessionID: uuid,
				UserName:  user.Username,
				UserID:    user.ID,
			}
			//将session保存到数据库里面
			dao.AddSession(sess)
			//创建一个cookie，与session相关联
			cookie := http.Cookie{
				Name:       "user",
				Value:      uuid,
				Path:       "",
				Domain:     "",
				Expires:    time.Time{},
				RawExpires: "",
				MaxAge:     0,
				Secure:     false,
				HttpOnly:   true,
				SameSite:   0,
				Raw:        "",
				Unparsed:   nil,
			}
			//将cookie发送给浏览器
			http.SetCookie(w, &cookie)
			//说明用户名和密码正确
			t := template.Must(template.ParseFiles("views/pages/user/login_success.html"))
			t.Execute(w, user)
		} else {
			t := template.Must(template.ParseFiles("views/pages/user/login.html"))
			t.Execute(w,"用户名和密码错误")
		}
	}


}

func Register(w http.ResponseWriter,r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	email := r.PostFormValue("email")

	user, _ := dao.CheckUsername(username)
	if user.ID>0 {
		//说明用户已经存在
		t := template.Must(template.ParseFiles("views/pages/user/register.html"))
		t.Execute(w,"用户名已存在")
	} else {
		//用户不存在
		//保存数据
		dao.SaveUser(username,password,email)
		t := template.Must(template.ParseFiles("views/pages/user/register_success.html"))
		t.Execute(w,"")
	}
}

func Logout(w http.ResponseWriter, r *http.Request)  {
	//获取cookie
	cookie,_ := r.Cookie("user")
	if cookie!=nil {
		//获取cookie的值
		cookieValue := cookie.Value
		//删除数据库中与之对应的Session
		dao.DeleteSession(cookieValue)
		//设置cookie的失效
		cookie.MaxAge = -1
		//将修改之后的cookie发送给浏览器
		http.SetCookie(w,cookie)
	}
	//去首页
	GetPageBooksByPrice(w,r)
}

//CheckUserName 通过发送Ajax验证用户名是否可用
func CheckUserName(w http.ResponseWriter, r *http.Request) {
	//获取用户输入的用户名
	username := r.PostFormValue("username")
	//调用userdao中验证用户名和密码的方法
	user, _ := dao.CheckUsername(username)
	if user.ID > 0 {
		//用户名已存在
		w.Write([]byte("用户名已存在！"))
	} else {
		//用户名可用
		w.Write([]byte("<font style='color:green'>用户名可用！</font>"))
	}
}