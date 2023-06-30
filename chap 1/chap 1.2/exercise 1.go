package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	arr := generate_random_array(100000)
	start := time.Now()
	exercise_1_1(arr)
	fmt.Printf("exercise 1.1: %dµs elapsed\n", time.Since(start).Microseconds())
	start = time.Now()
	exercise_1_2(arr)
	fmt.Printf("exercise 1.2: %dµs elapsed\n", time.Since(start).Microseconds())
}

func exercise_1_1(arr []string) {
	// fmt.Println(strings.Join(arr, " "))
	strings.Join(arr, " ")

}

func exercise_1_2(arr []string) {
	s, sep := "", " "
	for _, arg := range arr {
		s += arg + sep
	}
	// fmt.Println(s)
}

func generate_random_array(size int) []string {
	rand.Seed(time.Now().Unix())
	arr := rand.Perm(size)
	string_arr := make([]string, size)
	for i := 0; i < len(arr); i++ {
		string_arr[i] = string(arr[i])
	}
	return string_arr
}
