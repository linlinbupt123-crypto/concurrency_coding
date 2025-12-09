package question1

import (
	"fmt"
	"sync"
)

// countWords counts occurrences of each word concurrently.
func countWords(words []string) map[string]int {
    counts := make(map[string]int)
    var wg sync.WaitGroup
    for _, w := range words {
        wg.Add(1)
        go func(word string) {
            defer wg.Done()
            counts[word]++ // supposed to safely update the map
        }(w)
    }
    wg.Wait()
    return counts
}

func Run_countWords(){
	words := []string{"go", "lang", "go", "code", "lang"}
    result := countWords(words)
    fmt.Println("Word counts:", result)
}

func countWordsWithLock(words []string) map[string]int {
        counts := make(map[string]int)
    var wg sync.WaitGroup
       var mu sync.Mutex
    for _, w := range words {
        wg.Add(1)
        go func(word string) {
            defer wg.Done()
            mu.Lock()
            counts[word]++ // supposed to safely update the map
            mu.Unlock()
        }(w)
    }
    wg.Wait()
    return counts
}

func Run_countWordsWithLock(){
	words := []string{"go", "lang", "go", "code", "lang"}
    result := countWordsWithLock(words)
    fmt.Println("Word counts:", result)
}

func PaincMapWithoutLock(){
    	m := make(map[int]int)
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			m[i] = i
		}(i)
	}

	wg.Wait()
}

/*
Finally, a more efficient and idiomatic way to implement concurrent word counting in Go is to use the MapReduce pattern:
分片（Split）
把输入大数组切成多个分片。

map workers（Map）
多个 goroutine 并行处理分片，分别统计词频。

reduce 阶段合并结果（Reduce）
把所有 worker 的 map 合并。

这样就避免了多个 goroutine 同时写同一 map 的风险，也减少锁竞争。
*/

func countChunks(words []string) map[string]int {
    m := make(map[string]int)
    for _, w := range words {
        m[w]++
    }   
    return m
}

func countWordsMapReduce(words []string, workCount int) map[string]int {
    chunkSize := (len(words) + workCount - 1) / workCount
    results := make([]map[string]int, workCount)
    var wg sync.WaitGroup
    // map
    for i:=0; i < workCount; i++ {
        start := i *chunkSize
        end := start + chunkSize
        if start >= len(words) {
            break
        }
        if end > len(words) {
            end = len(words)
        }
        wg.Add(1)
        go func(idx int, slice[]string){
            defer wg.Done()
            results[idx] = countChunks(slice)
        }(i,words[start:end] )
    }
    wg.Wait()
    // reduce
    finalResult := make(map[string]int)
    for _, partialMap := range results {
        for k, v := range partialMap {
            finalResult[k] += v
        }
    }
    return finalResult
}

func Run_countWordsMapReduce(){
    words := []string{"go", "lang", "go", "code", "lang"}
    workCount := 3
    result := countWordsMapReduce(words, workCount)
    fmt.Println("Word counts (MapReduce):", result)
}