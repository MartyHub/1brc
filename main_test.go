package main

import (
	"runtime"
	"testing"
)

const (
	measurementsBench = "measurements_10_000_000.txt"
	measurementsTest  = "measurements_100_000_000.txt"
)

func Test_run(_ *testing.T) {
	run(measurementsTest, false, runtime.NumCPU()*2, 1024*512)
}

func Benchmark_run_1xCores_128KB(b *testing.B) {
	for n := 0; n < b.N; n++ {
		run(measurementsBench, false, runtime.NumCPU(), 1024*128)
	}
}

func Benchmark_run_2xCores_128KB(b *testing.B) {
	for n := 0; n < b.N; n++ {
		run(measurementsBench, false, runtime.NumCPU()*2, 1024*128)
	}
}

func Benchmark_run_4xCores_128KB(b *testing.B) {
	for n := 0; n < b.N; n++ {
		run(measurementsBench, false, runtime.NumCPU()*4, 1024*128)
	}
}

func Benchmark_run_8xCores_128KB(b *testing.B) {
	for n := 0; n < b.N; n++ {
		run(measurementsBench, false, runtime.NumCPU()*8, 1024*128)
	}
}

func Benchmark_run_1xCores_256KB(b *testing.B) {
	for n := 0; n < b.N; n++ {
		run(measurementsBench, false, runtime.NumCPU(), 1024*256)
	}
}

func Benchmark_run_2xCores_256KB(b *testing.B) {
	for n := 0; n < b.N; n++ {
		run(measurementsBench, false, runtime.NumCPU()*2, 1024*256)
	}
}

func Benchmark_run_4xCores_256KB(b *testing.B) {
	for n := 0; n < b.N; n++ {
		run(measurementsBench, false, runtime.NumCPU()*4, 1024*256)
	}
}

func Benchmark_run_8xCores_256KB(b *testing.B) {
	for n := 0; n < b.N; n++ {
		run(measurementsBench, false, runtime.NumCPU()*8, 1024*256)
	}
}

func Benchmark_run_1xCores_500KB(b *testing.B) {
	for n := 0; n < b.N; n++ {
		run(measurementsBench, false, runtime.NumCPU(), 1024*512)
	}
}

func Benchmark_run_2xCores_500KB(b *testing.B) {
	for n := 0; n < b.N; n++ {
		run(measurementsBench, false, runtime.NumCPU()*2, 1024*512)
	}
}

func Benchmark_run_4xCores_500KB(b *testing.B) {
	for n := 0; n < b.N; n++ {
		run(measurementsBench, false, runtime.NumCPU()*4, 1024*512)
	}
}

func Benchmark_run_8xCores_500KB(b *testing.B) {
	for n := 0; n < b.N; n++ {
		run(measurementsBench, false, runtime.NumCPU()*8, 1024*512)
	}
}

func Benchmark_run_1xCores_1MB(b *testing.B) {
	for n := 0; n < b.N; n++ {
		run(measurementsBench, false, runtime.NumCPU(), 1024*1024)
	}
}

func Benchmark_run_2xCores_1MB(b *testing.B) {
	for n := 0; n < b.N; n++ {
		run(measurementsBench, false, runtime.NumCPU()*2, 1024*1024)
	}
}

func Benchmark_run_4xCores_1MB(b *testing.B) {
	for n := 0; n < b.N; n++ {
		run(measurementsBench, false, runtime.NumCPU()*4, 1024*1024)
	}
}

func Benchmark_run_8xCores_1MB(b *testing.B) {
	for n := 0; n < b.N; n++ {
		run(measurementsBench, false, runtime.NumCPU()*8, 1024*1024)
	}
}
