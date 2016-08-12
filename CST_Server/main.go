// CST_Server project main.go
package main

import (
	"DbUitl"
	"MyProbuf"
	"OptUtil"
	"bytes"
	"data"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/golang/protobuf/proto"
)

const (
	msg_length = 2 * 1024
)

var msg = &DataFrame.Msg{}
var send_msg = &DataFrame.Msg{}
var socketMap = make(map[int]net.Conn)
var char_msg = &data.Message{}
var msgIndex = 0

func main() {
	log.Println("start to listen...")
	listen, err := net.Listen("tcp", ":6666")
	if err != nil {
		fmt.Println("listen error:", err)
	}
	defer listen.Close()
	DbUitl.ConnectDb()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			continue
		}
		log.Println("connected from " + conn.RemoteAddr().String())
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	data := make([]byte, msg_length)
	for {
		n, err := conn.Read(data)
		err = proto.Unmarshal(data[0:n], msg)
		if err != nil {
			fmt.Println("unmarshaling error: ", err)
		}
		switch msg.UserOpt {
		case OptUtil.REQUEST_LOGIN:
			handleLogin(conn)
			break
		case OptUtil.REQUEST_GET_FRIENDS:
			handleGetFriend(conn)
			break
		case OptUtil.REQUEST_SEND_TXT:
			handleSendText(conn)
			break
		}
	}
	fmt.Println("msg:", msg)
	defer conn.Close()
}

func handleLogin(conn net.Conn) {
	username := msg.User.UesrName
	password := msg.User.UserPwd
	fmt.Println("username:" + username)
	fmt.Println("password:" + password)
	var userdata = &data.User{}
	if username != "" && password != "" {
		userdata = DbUitl.Login(username, password)
		fmt.Println("Server:", userdata)
	}
	if userdata.UserId != 0 {
		fmt.Println("send success...")
		send_msg = &DataFrame.Msg{
			UserOpt:       OptUtil.RESULT_LOGIN,
			OptResult:     true,
			ReceiveResult: "Success",
			User: &DataFrame.User{
				UserID:   int32(userdata.UserId),
				UesrName: userdata.UserName,
				NickName: userdata.NickeName,
			},
		}
	} else {
		fmt.Println("send faild...")
		send_msg = &DataFrame.Msg{
			UserOpt:       OptUtil.RESULT_LOGIN,
			OptResult:     false,
			ReceiveResult: "UserName or Password Eorro..",
		}
	}
	socketMap[userdata.UserId] = conn
	data, err := proto.Marshal(send_msg)
	fmt.Println("Marshal:", send_msg)
	if err != nil {
		fmt.Println("marshaling error: ", err)
	}
	conn.Write(data)
	fmt.Println("Write data....", data)
}

func handleGetFriend(conn net.Conn) {
	useId := int(msg.User.UserID)
	fmt.Println("recive userid:", useId)
	var friendMap = make(map[int]data.User)
	var index int
	if useId != 0 {
		index, friendMap = DbUitl.GetFriends(strconv.Itoa(useId))
		fmt.Println("friendMap:", friendMap)
	}
	var friendlist []*DataFrame.User = make([]*DataFrame.User, index)
	if friendMap != nil {
		var listIndex int = 0
		for _, value := range friendMap {
			fmt.Println("value:", value)
			friends := new(DataFrame.User)
			friends.UserID = int32(value.UserId)
			friends.UesrName = value.UserName
			friends.NickName = value.NickeName
			friendlist[listIndex] = friends
			listIndex++
		}
		send_msg = &DataFrame.Msg{
			UserOpt:       OptUtil.RESULT_GET_FRIEND,
			OptResult:     true,
			ReceiveResult: "Success",
			Friends:       friendlist,
		}

		data, err := proto.Marshal(send_msg)
		fmt.Println("Marshal:", send_msg)
		if err != nil {
			fmt.Println("marshaling error: ", err)
		}
		conn.Write(data)
		fmt.Println("friend send_msg:", send_msg)
	}
}

func handleSendText(conn net.Conn) {
	char_msg.ReciverID = int(msg.PersonalMsg[msgIndex].RecverID)
	char_msg.SenderID = int(msg.PersonalMsg[msgIndex].SenderID)
	char_msg.Content = msg.PersonalMsg[msgIndex].Content
	char_msg.Time = msg.PersonalMsg[msgIndex].SendTime
	char_msg.DataType = int(OptUtil.MESSAGE_TYPE_TXT)
	fmt.Println("handleSendText:", char_msg)
	_, ok := socketMap[char_msg.ReciverID]
	if ok {
		//在线
	} else {

		DbUitl.SaveMessage(char_msg)
		fmt.Println("用户" + strconv.Itoa(char_msg.ReciverID) + "不在线，先把消息暂存在服务器端")
	}
	msgIndex++
}

func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)
	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	return int(x)

}

func IntToBytes(n int) []byte {
	x := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}
