// Code generated by "stringer -type=operationType"; DO NOT EDIT.

package maptracker

import "strconv"

const _operationType_name = "operationUnknownoperationPinoperationUnpin"

var _operationType_index = [...]uint8{0, 16, 28, 42}

func (i operationType) String() string {
	if i < 0 || i >= operationType(len(_operationType_index)-1) {
		return "operationType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _operationType_name[_operationType_index[i]:_operationType_index[i+1]]
}
