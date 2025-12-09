package question1

import (
	"testing"
)

func Test_Run(t *testing.T){
	Run_countWords()
}

func Test_PaincMapWithoutLock(t *testing.T){
	PaincMapWithoutLock()}

func Test_Run_countWordsWithLock(t *testing.T){
	Run_countWordsWithLock()
}

func Benchmark_countWords(b *testing.B) {
	words := []string{"go", "lang", "go", "code", "lang"}
	for i := 0; i < b.N; i++ {
		countWordsMapReduce(words,3)
	}	
}