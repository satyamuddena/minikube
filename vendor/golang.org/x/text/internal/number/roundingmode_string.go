// Code generated by "stringer -type RoundingMode"; DO NOT EDIT.

package number

import "strconv"

const _RoundingMode_name = "ToNearestEvenToNearestZeroToNearestAwayToPositiveInfToNegativeInfToZeroAwayFromZeronumModes"

var _RoundingMode_index = [...]uint8{0, 13, 26, 39, 52, 65, 71, 83, 91}

func (i RoundingMode) String() string {
	if i >= RoundingMode(len(_RoundingMode_index)-1) {
		return "RoundingMode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _RoundingMode_name[_RoundingMode_index[i]:_RoundingMode_index[i+1]]
}