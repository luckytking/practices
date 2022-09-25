package main

import "fmt"

type WorldMapService struct {
}

func NewWorldMapService() WorldMapService {
	return WorldMapService{}
}
func (w WorldMapService) HandleRegSuccess(u User) {
	fmt.Println(fmt.Sprintf("MapService: 欢迎%v来到新手村", u.Name))
}

type MailService struct {
}

func NewMailService() *MailService {
	return &MailService{}
}
func (w MailService) HandleRegSuccess(u User) {
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
	obsrvers []RegObserver
}

func NewUserController(obsrvers ...RegObserver) UserController {
	return UserController{obsrvers: obsrvers}
}

func (c UserController) Do() {
	userSvc := NewUserService()
	user := userSvc.Register("好奇的小明")
	for _, ob := range c.obsrvers {
		ob.HandleRegSuccess(user)
	}
	user.Start()
}

type RegObserver interface {
	HandleRegSuccess(User)
}

func main() {
	controller := NewUserController(NewWorldMapService(), NewMailService())
	controller.Do()
}
