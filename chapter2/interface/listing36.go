package main

import "fmt"

func main()  {
	bill:=user{"bill","bill@email.com" }
	sendNotification(&bill)

	
}

type notifier interface{
	notify()
}
type user struct {
	name string
	email string

}

func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n",u.name,u.email)

}
func sendNotification(n notifier){
	n.notify()
}