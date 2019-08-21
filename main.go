package main

import (
	"fmt"

	"github.com/Azure/golua/lua"
	"github.com/Azure/golua/std"
)

func main() {
	// fmt.Printf("hello world")
	var err error

	var opts = []lua.Option{lua.WithTrace(false), lua.WithVerbose(false)}
	state := lua.NewState(opts...)
	defer state.Close()
	std.Open(state)

	var libs = []struct {
		Name string
		Open lua.Func
	}{
		{"permission", lua.Func(permissionModule)},
	}

	for _, lib := range libs {
		state.Logf("opening stdlib mode %q", lib.Name)
		state.Require(lib.Name, lib.Open, true)
		state.Pop()
	}

	err = state.ExecText(`print("Hello World")`)
	if err != nil {
		fmt.Println(err)
	}

	err = state.ExecText(`
		user_id = "alex"
		role = permission.get_user_role(user_id)
		if role == nil then
			print("can not find role for", user_id)
		else
			print("got role:", role)
		end
	`)
	if err != nil {
		fmt.Println(err)
	}

}
