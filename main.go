package main

import (
	_ "golua/conf"

	lua "github.com/yuin/gopher-lua"
)

func main() {
	L := lua.NewState()
	if err := L.DoFile("lua/test.lua"); err != nil {
		panic(err)
	}
	err := L.CallByParam(
		lua.P{
			Fn:      L.GetGlobal("main"),
			NRet:    1,
			Protect: true,
		}, lua.LString("input param"))
	if err != nil {
		panic(err)
	}
	ret := L.Get(-1)
	println(string(ret.(lua.LString)))
}
