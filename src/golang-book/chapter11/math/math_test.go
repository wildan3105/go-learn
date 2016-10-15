package math

import "testing"

// func TestAverage(t *testing.T){
//   var v float64
//   v = Average([]float64{1,2,3,3,4,5,4,5})
//   if v != 3.375 {
//     t.Error("Expected 3.375, but you got ", v)
//   }
// }

// dozens of test
type testpair struct {
  values []float64
  average float64
}

var tests = []testpair{
  {[]float64{1,2}, 1.5},
  {[]float64{1,2,3,4}, 2.5},
  {[]float64{2,2,2,2}, 2},
}

func TestAverage(t *testing.T){
  for _,pair := range tests {
    v:= Average(pair.values)
    if v != pair.average {
      t.Error(
        "For", pair.values,
        "expected", pair.average,
        "got", v,
      )
    }
  }
}
