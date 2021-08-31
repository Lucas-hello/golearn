package main

import "fmt"

type Flags uint
const(
	FlagUp Flags=1<<iota //1
	FlagBroadcast //广播  //2
	FlagLoopback //回传	//4
	FlagPointToPoint    //8
	FlagMulticast//多播 //32
)

func IsUp(v Flags) bool {return v&FlagUp==FlagUp}
func TrueDown(v *Flags)  {*v &^= FlagUp}
func SetBroadcast(v *Flags)  {*v|=FlagBroadcast}
func IsCast(v Flags) bool {return v&(FlagLoopback|FlagLoopback) !=0}

func main()  {
	var v Flags = FlagMulticast|FlagUp
	fmt.Printf("%b %t\n",v,IsUp(v))
	TrueDown(&v)
	fmt.Printf("%b %t\n",v,IsUp(v))
	SetBroadcast(&v)
	fmt.Printf("%b %t\n",v,IsUp(v))
	fmt.Printf("%b %t\n",v,IsCast(v))

}