package common

import "testing"

func TestGetAverageFloat64(t *testing.T) {
	type test struct{
		src []float64
		res float64
	}

	values := []test {
		{
			[]float64{ 1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0 },
			5.5,
		},
	}

	for _,v := range values {
		res := GetAverageFloat64(v.src)
		if res != v.res {
			t.Errorf("getAverageFloat64(%g) = %g; want %g",v.src, res, v.res)
		}
	}
}
