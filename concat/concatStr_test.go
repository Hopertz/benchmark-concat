package concatenation

import "testing"

var resultBuffer string
var resultString string
func BenchmarkConcatWithBuffer(b *testing.B) {
   var res string
   for i:=0; i<b.N; i++ {
     // always record the result of function call to prevent
     // the compiler eliminating the function call.
	  res = ConcatWithBuffer()
   }
    // always store the result to a package level variable
    // so the compiler cannot eliminate the Benchmark itself.
        resultBuffer = res
}

func BenchmarkConcatWithJoin(b *testing.B) {
   var res string
   for i:= 0; i<b.N; i++ {
	  res = ConcatWithJoin()
   }
   resultString = res
}