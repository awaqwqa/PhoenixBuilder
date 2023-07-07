package BuiltlnFn

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
)

type LoadSide struct {
	*BuiltlnFn
}

func (b *LoadSide) BuiltFunc(L *lua.LState) int {
	if L.GetTop() == 1 {
		code := L.CheckString(1)
		if _, err := L.LoadString(code); err != nil {
			fmt.Println("lua插件报错", err)
		}
	} else {
		fmt.Println("加载插件需要一个参数 代表lua代码")
	}
	return 0
}

/*
func (b *LoadSide) LoadSideComponent(L *lua.LState) int {
	if L.GetTop() == 1 {
		code := L.CheckString(1)
		if _, err := L.LoadString(code); err != nil {
			fmt.Println("lua插件报错", err)
		}
	} else {
		fmt.Println("加载插件需要一个参数 代表lua代码")
	}
	return 0
}
*/