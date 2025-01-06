package main

import (
	"fmt"
	"sync"
	"time"
)

var m = sync.RWMutex{}
var wg = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

func main() {
	t0 := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go dbCall(i)
		count()
	}
	wg.Wait()
	fmt.Printf("\nTotal execution time: %v", time.Since(t0))
	fmt.Printf("\n The results are %v", results)
}

func count() {
	var res int
	for i := 0; i < 1000000; i++ {
		res += 1
	}
	wg.Done()
}

// func dbCall(i int) {
// 	var delay float32 = 2000
// 	time.Sleep(time.Duration(delay) * time.Millisecond)
// 	fmt.Println("The result from the database is:", dbData[i])
// 	m.Lock()
// 	results = append(results, dbData[i])
// 	m.Unlock()
// 	wg.Done()
// }

func dbCall(i int) {
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	// save(dbData[i])
	// log()
	wg.Done()
}

func save(result string) {
	m.Lock()
	results = append(results, result)
}

func log() {
	m.RLock()
	fmt.Printf("\nThe current results are: %v", results)
	m.RUnlock()
}
