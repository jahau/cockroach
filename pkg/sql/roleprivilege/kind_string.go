// Code generated by "stringer -type=Kind"; DO NOT EDIT.

package roleprivilege

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[CREATEROLE-1]
	_ = x[NOCREATEROLE-2]
}

const _Kind_name = "CREATEROLENOCREATEROLE"

var _Kind_index = [...]uint8{0, 10, 22}

func (i Kind) String() string {
	i -= 1
	if i >= Kind(len(_Kind_index)-1) {
		return "Kind(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Kind_name[_Kind_index[i]:_Kind_index[i+1]]
}
