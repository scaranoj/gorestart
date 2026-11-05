package main

import (
	"fmt"

	"github.com/scaranoj/gorestart/udemy/learngo/video_lectures/exercise_testing_and_bench_HoE2/quote"
	"github.com/scaranoj/gorestart/udemy/learngo/video_lectures/exercise_testing_and_bench_HoE2/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
