// data project data.go
package data

type User struct {
	UserId    int
	UserName  string
	UserPwd   string
	NickeName string
	Sex       string
	StdNum    string
}

type Message struct {
	SenderID  int
	ReciverID int
	DataType  int
	Content   string
	Time      string
}
