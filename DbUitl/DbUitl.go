// DbUitl project DbUitl.go
package DbUitl

import (
	"Server/FIFOQueue"
	"Server/MyProbuf"
	"Server/data"
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func ConnectDb() {
	db, err = sql.Open("mysql", "root:123456@tcp(localhost:3306)/cst_project?charset=utf8")
	checkError(err)
}

func Login(userName, pwd string) data.User {
	fmt.Println("start to Login Dbutil....")
	sqlStr := "select * from user where userName = '" + userName + "' and userPwd = '" + pwd + "'"
	rows, err := db.Query(sqlStr)
	fmt.Println("start to Query")
	checkError(err)
	user_data := data.User{}
	fmt.Println("start to Next()")
	for rows.Next() {
		err = rows.Scan(&user_data.UserId, &user_data.UserPwd, &user_data.NickeName, &user_data.Sex, &user_data.StdNum, &user_data.UserName)
		checkError(err)
	}

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
	index := 0
	friend_data := data.User{}
	for rows.Next() {
		err = rows.Scan(&friend_data.UserId, &friend_data.UserName, &friend_data.NickeName, &friend_data.Sex)
		checkError(err)
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

func GetOffLineMsg(receverID string) *FIFOQueue.Queue {
	queue := FIFOQueue.NewQueue()
	sqlStr := "select * from message where receverId='" + receverID + "'"
	rows, err := db.Query(sqlStr)
	checkError(err)
	for rows.Next() {
		msg_data := new(DataFrame.PersonalMsg)
		err = rows.Scan(&msg_data.SenderID, &msg_data.RecverID, &msg_data.MsgType, &msg_data.Content, &msg_data.SendTime)
		checkError(err)
		fmt.Println("Db getOffMsg:", queue.Enqueue(msg_data))
	}
	return queue
}

func DeleteMessage(receverID string) {
	sql := "delete from message where receverId='" + receverID + "'"
	_, err := db.Exec(sql)
	checkError(err)
}

func Close() {
	if db != nil {
		db.Close()
	}
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
