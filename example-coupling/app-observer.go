package main

import "fmt"

type WorldMapService struct {
}

func NewWorldMapService() WorldMapService {
	return WorldMapService{}
}
func (w WorldMapService) Join(u User) {
	fmt.Println(fmt.Sprintf("MapService: 欢迎%v来到新手村", u.Name))
}

type MailService struct {
}

func NewMailService() MailService {
	return MailService{}
}
func (w MailService) Send(u User) {
	fmt.Println(fmt.Sprintf("MailService: 恭喜勇士%v,获得金币999", u.Name))
}

type UserService struct {
}

func NewUserService() UserService {
	return UserService{}
}

type User struct {
	Name string
}

func (u UserService) Register(name string) User {
	return User{Name: name}
}
func (m User) Start() {
	fmt.Printf("UserService: 让我们开始愉快的旅程吧")
}

type UserController struct {
	User     UserService
	WorldMap WorldMapService
	Mail     MailService
}

func NewUserController() UserController {
	user := NewUserService()
	worldMap := NewWorldMapService()
	mail := NewMailService()
	return UserController{
		User:     user,
		WorldMap: worldMap,
		Mail:     mail,
	}
}

func (c UserController) Do() {
	user := c.User.Register("好奇的小明")
	c.WorldMap.Join(user)
	c.Mail.Send(user)
	user.Start()
}

func main() {
	controller := NewUserController()
	controller.Do()
}
