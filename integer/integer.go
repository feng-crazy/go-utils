package integer

// IntMax returns the maximum of the params
func IntMax(a, b int) int {
	if b > a {
		return b
	}
	return a
}

// IntMin returns the minimum of the params
func IntMin(a, b int) int {
	if b < a {
		return b
	}
	return a
}

// Int32Max returns the maximum of the params
func Int32Max(a, b int32) int32 {
	if b > a {
		return b
	}
	return a
}

// Int32Min returns the minimum of the params
func Int32Min(a, b int32) int32 {
	if b < a {
		return b
	}
	return a
}

// Int64Max returns the maximum of the params
func Int64Max(a, b int64) int64 {
	if b > a {
		return b
	}
	return a
}

// Int64Min returns the minimum of the params
func Int64Min(a, b int64) int64 {
	if b < a {
		return b
	}
	return a
}

// RoundToInt32 rounds floats into integer numbers.
func RoundToInt32(a float64) int32 {
	if a < 0 {
		return int32(a - 0.5)
	}
	return int32(a + 0.5)
}
