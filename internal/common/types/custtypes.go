// Code generated;  DO NOT EDIT.

package types

import (
	"fmt"
	"math"
	"strconv"
)

const (
	ActionNone = iota
	ValuesDeeplyEqual
	RequiresReplace
	RequiresUpdate
	ControlFlagUpdate
	PortListUpdate
)

func AnyToString(v any) (string, error) {
	switch t := v.(type) {
	case string:
		return t, nil
	case bool:
		return strconv.FormatBool(t), nil
	case int64:
		return strconv.FormatInt(t, 10), nil
	case float64:
		return strconv.FormatFloat(t, 'f', -1, 64), nil
	}
	return "", fmt.Errorf("unsupported type %T", v)
}

func AnyToBool(v any) (bool, error) {
	switch t := v.(type) {
	case bool:
		return t, nil
	case string:
		return strconv.ParseBool(t)
	case float64:
		if t == 0 || t == 1 {
			return t == 1, nil
		}
	case int64:
		if t == 0 || t == 1 {
			return t == 1, nil
		}
	}
	return false, fmt.Errorf("cannot convert %v to bool", v)
}

func AnyToInt64(v any) (int64, error) {
	switch t := v.(type) {
	case int64:
		return t, nil
	case float64:
		// The upper bound is exclusive: float64(MaxInt64) rounds up to 2^63.
		if t >= -0x1p63 && t < 0x1p63 && t == math.Trunc(t) {
			return int64(t), nil
		}
	case string:
		return strconv.ParseInt(t, 10, 64)
	case bool:
		if t {
			return 1, nil
		}
		return 0, nil
	}
	return 0, fmt.Errorf("cannot convert %v to int64", v)
}

func AnyEqual(a, b any) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}
	sa, ea := AnyToString(a)
	sb, eb := AnyToString(b)
	return ea == nil && eb == nil && sa == sb
}
