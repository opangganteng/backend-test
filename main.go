package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"
)

func sumEven(numbers []int, workerCount int) (int64, error) {
	if workerCount <= 0 {
		return 0, fmt.Errorf("jumlah worker harus lebih besar dari 0")
	}
	if len(numbers) == 0 {
		return 0, nil
	}
	if workerCount > len(numbers) {
		workerCount = len(numbers)
	}

	results := make(chan int64, workerCount)
	var workers sync.WaitGroup

	for worker := 0; worker < workerCount; worker++ {
		start := worker * len(numbers) / workerCount
		end := (worker + 1) * len(numbers) / workerCount

		workers.Add(1)
		go func(part []int) {
			defer workers.Done()

			var partialSum int64
			for _, number := range part {
				if number%2 == 0 {
					partialSum += int64(number)
				}
			}
			results <- partialSum
		}(numbers[start:end])
	}

	go func() {
		workers.Wait()
		close(results)
	}()

	var total int64
	for partialSum := range results {
		total += partialSum
	}

	return total, nil
}

func readIntegers(input io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanWords)

	var numbers []int
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, fmt.Errorf("input %q bukan integer yang valid: %w", scanner.Text(), err)
		}
		numbers = append(numbers, number)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("gagal membaca input: %w", err)
	}

	return numbers, nil
}

func run(input io.Reader, output io.Writer, workerCount int) error {
	numbers, err := readIntegers(input)
	if err != nil {
		return err
	}

	total, err := sumEven(numbers, workerCount)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintln(output, total)
	return err
}

func main() {
	workerCount := flag.Int("workers", 4, "jumlah goroutine worker")
	flag.Parse()

	if err := run(os.Stdin, os.Stdout, *workerCount); err != nil {
		fmt.Fprintln(os.Stderr, strings.TrimSpace(err.Error()))
		os.Exit(1)
	}
}
