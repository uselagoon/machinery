package conversion

import "strconv"

// IntPtr .
func IntPtr(i int32) *int32 {
	return &i
}

// Int64Ptr .
func Int64Ptr(i int64) *int64 {
	return &i
}

// UintPtr .
func UintPtr(i uint) *uint {
	return &i
}

func StringPtr(s string) *string {
	return &s
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
