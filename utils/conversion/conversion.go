package conversion

import "strconv"

// IntPtr .
func IntPtr(i int32) *int32 {
	var iPtr *int32
	iPtr = new(int32)
	*iPtr = i
	return iPtr
}

// Int64Ptr .
func Int64Ptr(i int64) *int64 {
	var iPtr *int64
	iPtr = new(int64)
	*iPtr = i
	return iPtr
}

// UintPtr .
func UintPtr(i uint) *uint {
	var iPtr *uint
	iPtr = new(uint)
	*iPtr = i
	return iPtr
}

// StringToUintPtr get the environment id and turn it into a *uint to send back to lagoon for logging
// lagoon sends this as a string :(
func StringToUintPtr(s string) *uint {
	u64, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return nil
	}
	return UintPtr(uint(u64))
}
