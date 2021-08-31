package main

import "fmt"

func main (){
	bill:=user1{"bill","bill@email.com"}
	sendNotification1(&bill)
	lisa:=user1{"lisa","lisa@email.com"}

	sendNotification1(&lisa)


}
type notifier1 interface {
	notify1()
}
type user1 struct {
	name string
	email string
}

func (u *user1)notify1()  {
	fmt.Printf("Sending user email to %s<%s>\n",u.name,u.email)

}

type admin struct{
	name string
	email string
}

func (a *admin)notify1()  {
	fmt.Printf("Sending user emeil to %s<%s>\n",a.name,a.email)

}
func sendNotification1(n notifier1){
	n.notify1()
}