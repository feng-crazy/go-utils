package integer

import "testing"

func TestIntMax(t *testing.T) {
	tests := []struct {
		nums        []int
		expectedMax int
	}{
		{
			nums:        []int{-1, 0},
			expectedMax: 0,
		},
		{
			nums:        []int{-1, -2},
			expectedMax: -1,
		},
		{
			nums:        []int{0, 1},
			expectedMax: 1,
		},
		{
			nums:        []int{1, 2},
			expectedMax: 2,
		},
	}

	for i, test := range tests {
		t.Logf("executing scenario %d", i)
		if max := IntMax(test.nums[0], test.nums[1]); max != test.expectedMax {
			t.Errorf("expected %v,  got %v", test.expectedMax, max)
		}
	}
}

func TestIntMin(t *testing.T) {
	tests := []struct {
		nums        []int
		expectedMin int
	}{
		{
			nums:        []int{-1, 0},
			expectedMin: -1,
		},
		{
			nums:        []int{-1, -2},
			expectedMin: -2,
		},
		{
			nums:        []int{0, 1},
			expectedMin: 0,
		},
		{
			nums:        []int{1, 2},
			expectedMin: 1,
		},
	}

	for i, test := range tests {
		t.Logf("executing scenario %d", i)
		if min := IntMin(test.nums[0], test.nums[1]); min != test.expectedMin {
			t.Errorf("expected %v,  got %v", test.expectedMin, min)
		}
	}
}

func TestInt32Max(t *testing.T) {
	tests := []struct {
		nums        []int32
		expectedMax int32
	}{
		{
			nums:        []int32{-1, 0},
			expectedMax: 0,
		},
		{
			nums:        []int32{-1, -2},
			expectedMax: -1,
		},
		{
			nums:        []int32{0, 1},
			expectedMax: 1,
		},
		{
			nums:        []int32{1, 2},
			expectedMax: 2,
		},
	}

	for i, test := range tests {
		t.Logf("executing scenario %d", i)
		if max := Int32Max(test.nums[0], test.nums[1]); max != test.expectedMax {
			t.Errorf("expected %v,  got %v", test.expectedMax, max)
		}
	}
}

func TestInt32Min(t *testing.T) {
	tests := []struct {
		nums        []int32
		expectedMin int32
	}{
		{
			nums:        []int32{-1, 0},
			expectedMin: -1,
		},
		{
			nums:        []int32{-1, -2},
			expectedMin: -2,
		},
		{
			nums:        []int32{0, 1},
			expectedMin: 0,
		},
		{
			nums:        []int32{1, 2},
			expectedMin: 1,
		},
	}

	for i, test := range tests {
		t.Logf("executing scenario %d", i)
		if min := Int32Min(test.nums[0], test.nums[1]); min != test.expectedMin {
			t.Errorf("expected %v,  got %v", test.expectedMin, min)
		}
	}
}

func TestInt64Max(t *testing.T) {
	tests := []struct {
		nums        []int64
		expectedMax int64
	}{
		{
			nums:        []int64{-1, 0},
			expectedMax: 0,
		},
		{
			nums:        []int64{-1, -2},
			expectedMax: -1,
		},
		{
			nums:        []int64{0, 1},
			expectedMax: 1,
		},
		{
			nums:        []int64{1, 2},
			expectedMax: 2,
		},
	}

	for i, test := range tests {
		t.Logf("executing scenario %d", i)
		if max := Int64Max(test.nums[0], test.nums[1]); max != test.expectedMax {
			t.Errorf("expected %v,  got %v", test.expectedMax, max)
		}
	}
}

func TestInt64Min(t *testing.T) {
	tests := []struct {
		nums        []int64
		expectedMin int64
	}{
		{
			nums:        []int64{-1, 0},
			expectedMin: -1,
		},
		{
			nums:        []int64{-1, -2},
			expectedMin: -2,
		},
		{
			nums:        []int64{0, 1},
			expectedMin: 0,
		},
		{
			nums:        []int64{1, 2},
			expectedMin: 1,
		},
	}

	for i, test := range tests {
		t.Logf("executing scenario %d", i)
		if min := Int64Min(test.nums[0], test.nums[1]); min != test.expectedMin {
			t.Errorf("expected %v,  got %v", test.expectedMin, min)
		}
	}
}

func TestRoundToInt32(t *testing.T) {
	tests := []struct {
		num float64
		exp int32
	}{
		{
			num: 5.5,
			exp: 6,
		},
		{
			num: -3.7,
			exp: -4,
		},
		{
			num: 3.49,
			exp: 3,
		},
		{
			num: -7.9,
			exp: -8,
		},
		{
			num: -4.499999,
			exp: -4,
		},
		{
			num: 0,
			exp: 0,
		},
	}

	for i, test := range tests {
		t.Logf("executing scenario %d", i)
		if got := RoundToInt32(test.num); got != test.exp {
			t.Errorf("expected %d, got %d", test.exp, got)
		}
	}
}
