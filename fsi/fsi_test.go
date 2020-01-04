package fsi

import (
	"testing"
	"sync"
)

func BenchmarkGenerate(b *testing.B) {

	f := New()
	for i := 0; i < b.N; i++ {

		f.Generate()

	}

}

func BenchmarkGenerateGoRoutines(b *testing.B) {

	f := New()
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		
		wg.Add(1)

		go func() {

			for i := 0; i < b.N/3; i++ {

				f.Generate()
		
			}
			wg.Done()

		}()

	}

	wg.Wait()

}
