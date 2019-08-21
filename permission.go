package main

import "github.com/Azure/golua/lua"

// Permission module
func permissionModule(state *lua.State) int {
	var helperFuncs = map[string]lua.Func{
		"get_user_role":      lua.Func(getUserRole),
		"get_curr_user_role": lua.Func(getCurrentUserRole),
	}

	state.NewTableSize(0, len(helperFuncs))
	state.SetFuncs(helperFuncs, 0)

	return 1
}

func getUserRole(state *lua.State) int {
	var userID = state.CheckString(1)

	var fakeRoleMap = map[string]string{
		"tom": "admin",
		"jim": "user",
	}

	if role, ok := fakeRoleMap[userID]; ok {
		state.Push(role)
		return 1
	}

	state.Push(nil)
	return 1
}

func getCurrentUserRole(state *lua.State) int {
	state.Push("admin")
	return 1
}
