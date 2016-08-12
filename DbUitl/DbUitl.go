// DbUitl project DbUitl.go
package DbUitl

import (
	"Server/data"
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

var m_userName string
var m_userPwd string
var m_nickName string
var m_sex string
var m_stdNum string
var m_userId int

var db *sql.DB
var err error

func ConnectDb() {
	db, err = sql.Open("mysql", "root:123456@tcp(localhost:3306)/cst_project?charset=utf8")
	checkError(err)
}

func Login(userName, pwd string) *data.User {
	fmt.Println("start to Login Dbutil....")
	sqlStr := "select * from user where userName = '" + userName + "' and userPwd = '" + pwd + "'"
	rows, err := db.Query(sqlStr)
	fmt.Println("start to Query")
	checkError(err)

	fmt.Println("start to Next()")
	for rows.Next() {
		err = rows.Scan(&m_userId, &m_userPwd, &m_nickName, &m_sex, &m_stdNum, &m_userName)
		checkError(err)
	}

	user_data := &data.User{m_userId, m_userName, m_userPwd, m_nickName, m_sex, m_stdNum}
	fmt.Println("DbUitl data:", user_data)
	return user_data
}

func GetFriends(selfId string) (int, map[int]data.User) {
	fmt.Println("start getfriendDb....")
	friendsMap := make(map[int]data.User)
	fmt.Println(selfId)
	sqlStr := "select userId, userName, nickName, sex from user, friend where userId=friendId and selfId='" + selfId + "'"
	rows, err := db.Query(sqlStr)
	checkError(err)
	var index int = 0
	for rows.Next() {
		err = rows.Scan(&m_userId, &m_userName, &m_nickName, &m_sex)
		checkError(err)
		friend_data := data.User{m_userId, m_userName, m_userPwd, m_nickName, m_sex, m_stdNum}
		friendsMap[index] = friend_data
		index++
	}
	return index, friendsMap
}

func SaveMessage(chat_msg *data.Message) {
	sqlstr := "insert into message values(" + strconv.Itoa(chat_msg.SenderID) + ", " + strconv.Itoa(chat_msg.ReciverID) + ", " + strconv.Itoa(chat_msg.DataType) + ", '" + chat_msg.Content + "', '" + chat_msg.Time + "')"
	_, err := db.Exec(sqlstr)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("DbUitl err:", err)
	}
}

//func main() {
//	db := ConnectDb()
//	sqlStr := "select * from user where userId = 2008072416 and userPwd = 123456"
//	rows, err := db.Query(sqlStr)
//	checkError(err)

//	var m_userId string
//	var m_userPwd string
//	var m_nickName string
//	var m_sex string
//	var m_stdNum string

//	for rows.Next() {
//		err = rows.Scan(&m_userId, &m_userPwd, &m_nickName, &m_sex, &m_stdNum)
//		checkError(err)
//	}
//	fmt.Println(m_userId + m_stdNum)
//}
