package goluas

import (
	"strings"

	"github.com/yuin/gopher-lua"
)

var exports = map[string]lua.LGFunction{
	"contains":    apiContains,
	"startswith":  apiStartsWith,
	"endswith":    apiEndsWith,
	"trim":        apiTrim,
	"trimleft":    apiTrimLeft,
	"trimright":   apiTrimRight,
	"trimprefix":  apiTrimPrefix,
	"trimpsuffix": apiTrimSuffix,
}

func Loader(L *lua.LState) int {
	mod := L.SetFuncs(L.NewTable(), exports)
	L.Push(mod)
	return 1
}

func apiContains(L *lua.LState) int {
	str := L.CheckString(1)
	lookup := L.CheckString(2)
	L.Push(lua.LBool(strings.Contains(str, lookup)))
	return 1
}

func apiStartsWith(L *lua.LState) int {
	str := L.CheckString(1)
	prefix := L.CheckString(2)
	L.Push(lua.LBool(strings.HasPrefix(str, prefix)))
	return 1
}

func apiEndsWith(L *lua.LState) int {
	str := L.CheckString(1)
	suffix := L.CheckString(2)
	L.Push(lua.LBool(strings.HasSuffix(str, suffix)))
	return 1
}

func apiTrim(L *lua.LState) int {
	str := L.CheckString(1)
	cutset := L.CheckString(2)
	L.Push(lua.LString(strings.Trim(str, cutset)))
	return 1
}

func apiTrimLeft(L *lua.LState) int {
	str := L.CheckString(1)
	cutset := L.CheckString(2)
	L.Push(lua.LString(strings.TrimLeft(str, cutset)))
	return 1
}

func apiTrimRight(L *lua.LState) int {
	str := L.CheckString(1)
	cutset := L.CheckString(2)
	L.Push(lua.LString(strings.TrimRight(str, cutset)))
	return 1
}

func apiTrimPrefix(L *lua.LState) int {
	str := L.CheckString(1)
	prefix := L.CheckString(2)
	L.Push(lua.LString(strings.TrimPrefix(str, prefix)))
	return 1
}

func apiTrimSuffix(L *lua.LState) int {
	str := L.CheckString(1)
	suffix := L.CheckString(2)
	L.Push(lua.LString(strings.TrimSuffix(str, suffix)))
	return 1
}
