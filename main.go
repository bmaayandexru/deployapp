package main

import "fmt"

func main() {
	/*
		var wg sync.WaitGroup
		count := 10000
		wg.Add(count)

		for i := 0; i < count; i++ {
			go func() {
				fmt.Println(i)
				wg.Done()
			}()
		}

		wg.Wait()
	*/
	fmt.Println("main()")
}

func MaxInt(a, b int) int {
	//	if a < b { // чтоб сломать тест
	if a >= b { // это верно
		return a
	}

	return b
}
