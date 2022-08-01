package sandbox_test

import "testing"

func TestX(t *testing.T)   {}
func Test_x(t *testing.T)  {}
func TestX_x(t *testing.T) {}

func BenchmarkX(b *testing.B)   {}
func Benchmark_x(b *testing.B)  {}
func BenchmarkX_x(b *testing.B) {}

func ExampleX(f *testing.F)   {}
func Example_x(f *testing.F)  {}
func ExampleX_x(f *testing.F) {}

func FuzzX(f *testing.F)   {}
func Fuzz_x(f *testing.F)  {}
func FuzzX_x(f *testing.F) {}

func xTestX(t *testing.T)   {}
func xTest_x(t *testing.T)  {} // want "xTest_x contains underscore. You should use mixedCap or MixedCap."
func xTestX_x(t *testing.T) {} // want "xTestX_x contains underscore. You should use mixedCap or MixedCap."

func xBenchmarkX(b *testing.B)   {}
func xBenchmark_x(b *testing.B)  {} // want "xBenchmark_x contains underscore. You should use mixedCap or MixedCap."
func xBenchmarkX_x(b *testing.B) {} // want "xBenchmarkX_x contains underscore. You should use mixedCap or MixedCap."

func xExampleX(f *testing.F)   {}
func xExample_x(f *testing.F)  {} // want "xExample_x contains underscore. You should use mixedCap or MixedCap."
func xExampleX_x(f *testing.F) {} // want "xExampleX_x contains underscore. You should use mixedCap or MixedCap."

func xFuzzX(f *testing.F)   {}
func xFuzz_x(f *testing.F)  {} // want "xFuzz_x contains underscore. You should use mixedCap or MixedCap."
func xFuzzX_x(f *testing.F) {} // want "xFuzzX_x contains underscore. You should use mixedCap or MixedCap."
