package sandbox_test

import "testing"

func TestX(t *testing.T)   {}
func Test_x(t *testing.T)  {}
func TestX_x(t *testing.T) {}

func BenchmarkX(b *testing.B)   {}
func Benchmark_x(b *testing.B)  {}
func BenchmarkX_x(b *testing.B) {}

func FuzzX(f *testing.F)   {}
func Fuzz_x(f *testing.F)  {}
func FuzzX_x(f *testing.F) {}
