package eval

import (
	"fmt"
	"math"
	"time"

	"github.com/feng-crazy/go-utils/clock"
	"github.com/feng-crazy/go-utils/number"
)

func DataEval(lhs, rhs interface{}, op string) interface{} {
	if lhs == nil || rhs == nil {
		switch op {
		case EQ, LTE, GTE:
			if lhs == nil && rhs == nil {
				return true
			} else {
				return false
			}
		case NEQ:
			if lhs == nil && rhs == nil {
				return false
			} else {
				return true
			}
		case LT, GT:
			return false
		default:
			return nil
		}
	}
	lhs = number.ConvertNum(lhs)
	rhs = number.ConvertNum(rhs)
	// Evaluate if both sides are simple types.
	switch lhs := lhs.(type) {
	case bool:
		rhs, ok := rhs.(bool)
		if !ok {
			return invalidOpError(lhs, op, rhs)
		}
		switch op {
		case AND:
			return lhs && rhs
		case OR:
			return lhs || rhs
		case BITWISE_AND:
			return lhs && rhs
		case BITWISE_OR:
			return lhs || rhs
		case BITWISE_XOR:
			return lhs != rhs
		case EQ:
			return lhs == rhs
		case NEQ:
			return lhs != rhs
		default:
			return invalidOpError(lhs, op, rhs)
		}
	case float64:
		// Try the rhs as a float64, int64, or uint64
		rhsf, ok := rhs.(float64)
		if !ok {
			switch val := rhs.(type) {
			case int64:
				rhsf, ok = float64(val), true
			case uint64:
				rhsf, ok = float64(val), true
			}
		}
		if !ok {
			return invalidOpError(lhs, op, rhs)
		}
		rhs := rhsf
		switch op {
		case EQ:
			return lhs == rhs
		case NEQ:
			return lhs != rhs
		case LT:
			return lhs < rhs
		case LTE:
			return lhs <= rhs
		case GT:
			return lhs > rhs
		case GTE:
			return lhs >= rhs
		case ADD:
			return lhs + rhs
		case SUB:
			return lhs - rhs
		case MUL:
			return lhs * rhs
		case DIV:
			if rhs == 0 {
				return fmt.Errorf("divided by zero")
			}
			return lhs / rhs
		case MOD:
			if rhs == 0 {
				return fmt.Errorf("divided by zero")
			}
			return math.Mod(lhs, rhs)
		default:
			return invalidOpError(lhs, op, rhs)
		}
	case int64:
		// Try as a float64 to see if a float cast is required.
		switch rhs := rhs.(type) {
		case float64:
			lhs := float64(lhs)
			switch op {
			case EQ:
				return lhs == rhs
			case NEQ:
				return lhs != rhs
			case LT:
				return lhs < rhs
			case LTE:
				return lhs <= rhs
			case GT:
				return lhs > rhs
			case GTE:
				return lhs >= rhs
			case ADD:
				return lhs + rhs
			case SUB:
				return lhs - rhs
			case MUL:
				return lhs * rhs
			case DIV:
				if rhs == 0 {
					return fmt.Errorf("divided by zero")
				}
				return lhs / rhs
			case MOD:
				if rhs == 0 {
					return fmt.Errorf("divided by zero")
				}
				return math.Mod(lhs, rhs)
			default:
				return invalidOpError(lhs, op, rhs)
			}
		case int64:
			switch op {
			case EQ:
				return lhs == rhs
			case NEQ:
				return lhs != rhs
			case LT:
				return lhs < rhs
			case LTE:
				return lhs <= rhs
			case GT:
				return lhs > rhs
			case GTE:
				return lhs >= rhs
			case ADD:
				return lhs + rhs
			case SUB:
				return lhs - rhs
			case MUL:
				return lhs * rhs
			case MOD:
				if rhs == 0 {
					return fmt.Errorf("divided by zero")
				}
				return lhs % rhs
			case BITWISE_AND:
				return lhs & rhs
			case BITWISE_OR:
				return lhs | rhs
			case BITWISE_XOR:
				return lhs ^ rhs
			default:
				return invalidOpError(lhs, op, rhs)
			}
		case uint64:
			switch op {
			case EQ:
				return uint64(lhs) == rhs
			case NEQ:
				return uint64(lhs) != rhs
			case LT:
				if lhs < 0 {
					return true
				}
				return uint64(lhs) < rhs
			case LTE:
				if lhs < 0 {
					return true
				}
				return uint64(lhs) <= rhs
			case GT:
				if lhs < 0 {
					return false
				}
				return uint64(lhs) > rhs
			case GTE:
				if lhs < 0 {
					return false
				}
				return uint64(lhs) >= rhs
			case ADD:
				return uint64(lhs) + rhs
			case SUB:
				return uint64(lhs) - rhs
			case MUL:
				return uint64(lhs) * rhs
			case DIV:
				if rhs == 0 {
					return fmt.Errorf("divided by zero")
				}
				return uint64(lhs) / rhs
			case MOD:
				if rhs == 0 {
					return fmt.Errorf("divided by zero")
				}
				return uint64(lhs) % rhs
			case BITWISE_AND:
				return uint64(lhs) & rhs
			case BITWISE_OR:
				return uint64(lhs) | rhs
			case BITWISE_XOR:
				return uint64(lhs) ^ rhs
			default:
				return invalidOpError(lhs, op, rhs)
			}
		default:
			return invalidOpError(lhs, op, rhs)
		}
	case uint64:
		// Try as a float64 to see if a float cast is required.
		switch rhs := rhs.(type) {
		case float64:
			lhs := float64(lhs)
			switch op {
			case EQ:
				return lhs == rhs
			case NEQ:
				return lhs != rhs
			case LT:
				return lhs < rhs
			case LTE:
				return lhs <= rhs
			case GT:
				return lhs > rhs
			case GTE:
				return lhs >= rhs
			case ADD:
				return lhs + rhs
			case SUB:
				return lhs - rhs
			case MUL:
				return lhs * rhs
			case DIV:
				if rhs == 0 {
					return fmt.Errorf("divided by zero")
				}
				return lhs / rhs
			case MOD:
				if rhs == 0 {
					return fmt.Errorf("divided by zero")
				}
				return math.Mod(lhs, rhs)
			default:
				return invalidOpError(lhs, op, rhs)
			}
		case int64:
			switch op {
			case EQ:
				return lhs == uint64(rhs)
			case NEQ:
				return lhs != uint64(rhs)
			case LT:
				if rhs < 0 {
					return false
				}
				return lhs < uint64(rhs)
			case LTE:
				if rhs < 0 {
					return false
				}
				return lhs <= uint64(rhs)
			case GT:
				if rhs < 0 {
					return true
				}
				return lhs > uint64(rhs)
			case GTE:
				if rhs < 0 {
					return true
				}
				return lhs >= uint64(rhs)
			case ADD:
				return lhs + uint64(rhs)
			case SUB:
				return lhs - uint64(rhs)
			case MUL:
				return lhs * uint64(rhs)
			case DIV:
				if rhs == 0 {
					return fmt.Errorf("divided by zero")
				}
				return lhs / uint64(rhs)
			case MOD:
				if rhs == 0 {
					return fmt.Errorf("divided by zero")
				}
				return lhs % uint64(rhs)
			case BITWISE_AND:
				return lhs & uint64(rhs)
			case BITWISE_OR:
				return lhs | uint64(rhs)
			case BITWISE_XOR:
				return lhs ^ uint64(rhs)
			default:
				return invalidOpError(lhs, op, rhs)
			}
		case uint64:
			switch op {
			case EQ:
				return lhs == rhs
			case NEQ:
				return lhs != rhs
			case LT:
				return lhs < rhs
			case LTE:
				return lhs <= rhs
			case GT:
				return lhs > rhs
			case GTE:
				return lhs >= rhs
			case ADD:
				return lhs + rhs
			case SUB:
				return lhs - rhs
			case MUL:
				return lhs * rhs
			case DIV:
				if rhs == 0 {
					return fmt.Errorf("divided by zero")
				}
				return lhs / rhs
			case MOD:
				if rhs == 0 {
					return fmt.Errorf("divided by zero")
				}
				return lhs % rhs
			case BITWISE_AND:
				return lhs & rhs
			case BITWISE_OR:
				return lhs | rhs
			case BITWISE_XOR:
				return lhs ^ rhs
			default:
				return invalidOpError(lhs, op, rhs)
			}
		default:
			return invalidOpError(lhs, op, rhs)
		}
	case string:
		rhss, ok := rhs.(string)
		if !ok {
			return invalidOpError(lhs, op, rhs)
		}
		switch op {
		case EQ:
			return lhs == rhss
		case NEQ:
			return lhs != rhss
		case LT:
			return lhs < rhss
		case LTE:
			return lhs <= rhss
		case GT:
			return lhs > rhss
		case GTE:
			return lhs >= rhss
		default:
			return invalidOpError(lhs, op, rhs)
		}
	case time.Time:
		rt, err := clock.InterfaceToTime(rhs, "")
		if err != nil {
			return invalidOpError(lhs, op, rhs)
		}
		switch op {
		case EQ:
			return lhs.Equal(rt)
		case NEQ:
			return !lhs.Equal(rt)
		case LT:
			return lhs.Before(rt)
		case LTE:
			return lhs.Before(rt) || lhs.Equal(rt)
		case GT:
			return lhs.After(rt)
		case GTE:
			return lhs.After(rt) || lhs.Equal(rt)
		default:
			return invalidOpError(lhs, op, rhs)
		}
	default:
		return invalidOpError(lhs, op, rhs)
	}

	return invalidOpError(lhs, op, rhs)
}

func invalidOpError(lhs interface{}, op string, rhs interface{}) error {
	return fmt.Errorf("invalid operation %[1]T(%[1]v) %s %[3]T(%[3]v)", lhs, op, rhs)
}
