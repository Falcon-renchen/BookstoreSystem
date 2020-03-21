package dao

import (
	"model"
	"net/http"
	"utils"
)

func AddSession(sess *model.Session) error {
	sqlStr := "insert into sessions values(?,?,?)"
	_, err := utils.Db.Exec(sqlStr,sess.SessionID,sess.UserName,sess.UserID)
	if err!=nil {
		return err
	}
	return nil
}

func DeleteSession(sessID string) error {
	sqlStr := "delete from sessions where session_id = ?"
	_, err := utils.Db.Exec(sqlStr,sessID)
	if err!=nil {
		return err
	}
	return nil
}

func GetSession(sessId string) (*model.Session, error) {
	sqlStr := "select session_id,username,user_id from sessions where session_id=?"
	//预编译
	inStmt, err := utils.Db.Prepare(sqlStr)
	if err!=nil {
		return nil,err
	}
	//执行
	row := inStmt.QueryRow(sessId)
	//创建session
	sess := &model.Session{}

	row.Scan(&sess.SessionID,&sess.UserName,&sess.UserID)
	return sess,nil
}

//IsLogin 判断用户是否已经登录 false没有登录 true 已经登录
func IsLogin(r *http.Request) (bool,*model.Session) {
	//根据Cookie的name获取Cookie
	cookie, _ := r.Cookie("user")
	if cookie!=nil {
		//获取Cookie的value
		cookieValue := cookie.Value
		//根据cookieValue去数据库中查询与之对应的Session
		session, _ := GetSession(cookieValue)
		if session.UserID > 0 {
			//已经登录
			return true,session
		}
	}
	//没有登录
	return false,nil
}