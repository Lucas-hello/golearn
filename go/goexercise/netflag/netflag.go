package main

import "fmt"

type Flags uint

const (
	FlagUp          Flags = 1 << iota
	FlagBroadcast         //2
	FlagLoopback          //4
	FlagPintToPoint       //8
	FlagMulticast         //16
)

func main() {
	var v Flags = FlagMulticast | FlagUp
	fmt.Printf("%d %t\n", v, IsUp(v))
	TurnDown(&v)

}
func IsUp(v Flags) bool     { return v&FlagUp == FlagUp }
func TurnDown(v *Flags)     { *v &^= FlagUp }
func SetBroadcast(v *Flags) { *v |= FlagBroadcast }
func IsCast(v Flags) bool   { return v&(FlagBroadcast|FlagMulticast) != 0 }
