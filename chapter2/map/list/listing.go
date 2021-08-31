package main

import "fmt"

func main()  {
	bill:=user{"bill","bill@email.com"}
	bill.notify()
	lisa:=&user{"lisa","lisa@email.com"}
	lisa.ChangeEmail("lasa@email.com" )
	bill.ChangeEmail("newbill@www.com")
	bill.notify()

}
type user struct {
	name string
	email string
}

func (u user)notify()  {
	 fmt.Printf("Sending User Email to %s<%s>\n",u.name,u.email)
	
}
func (u *user)ChangeEmail(email string)  {
	u.email=email

}
