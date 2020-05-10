package pipeline

func generator(n int) <-chan int {
	outChInt := make(chan int, 100)

	go func() {
		for index := 1; index <= n; index++ {
			outChInt <- index
		}
		close(outChInt)
	}()

	return outChInt
}

func power(in <-chan int) <-chan int {
	out := make(chan int, 100)

	go func() {
		for v := range in {
			out <- v * v
		}
		close(out)
	}()

	return out
}

func sum(in <-chan int) <-chan int {
	out := make(chan int, 100)

	go func() {
		var sum int

		for v := range in {
			sum += v
		}

		out <- sum
		close(out)
	}()

	return out
}

func LaunchPipeline(n int) int {
	firstCh := generator(n)
	secondCh := power(firstCh)
	thirdCh := sum(secondCh)

	result := <-thirdCh

	return result
}
