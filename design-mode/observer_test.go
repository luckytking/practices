package design_mode

import "testing"

func TestObserver(t *testing.T) {
	concreteSubject := NewConcreteSubject()
	concreteSubject.RegisterObserver(&ConcreteObserverOne{})
	concreteSubject.RegisterObserver(&ConcreteObserverTwo{})

	concreteSubject.NotifyObservers("hello every one")
}
