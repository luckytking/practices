/**
*@Author:qrTang
*@Date:2022/9/25
**/

package design_mode

import (
	"fmt"
)

type Subject interface {
	RegisterObserver(Observer)
	NotifyObservers(message interface{})
}

type Observer interface {
	Update(message interface{})
}

type ConcreteSubject struct {
	observerList []Observer
}

func NewConcreteSubject() Subject {
	return &ConcreteSubject{}
}

func (o *ConcreteSubject) RegisterObserver(observer Observer) {
	o.observerList = append(o.observerList, observer)
}

func (o *ConcreteSubject) NotifyObservers(message interface{}) {
	for _, observer := range o.observerList {
		observer.Update(message)
	}
}

type ConcreteObserverOne struct {
}

func (b *ConcreteObserverOne) Update(message interface{}) {
	fmt.Printf("ConcreteObserverOne is notified.%v \n", message)
}

type ConcreteObserverTwo struct {
}

func (b *ConcreteObserverTwo) Update(message interface{}) {
	fmt.Printf("ConcreteObserverOne is notified.%v \n", message)
}
