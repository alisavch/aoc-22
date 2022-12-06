package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Queue []int

func (q Queue) Len() int {
	return len(q)
}

func (q Queue) Less(i int, j int) bool {
	return q[i] < q[j]
}

func (q Queue) Swap(i int, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *Queue) Push(x interface{}) {
	*q = append(*q, x.(int))
}

func (q *Queue) Pop() interface{} {
	old := *q
	n := len(old)
	element := old[n-1]
	*q = old[0 : n-1]
	return element
}

func main() {
	path := os.Getenv("D1_PATH")
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("cannot read file: %v", err)
	}
	defer file.Close()

	var caloriesOfOne int

	pq := &Queue{}
	heap.Init(pq)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var row string

		for _, v := range scanner.Text() {
			i, err := strconv.Atoi(string(v))
			if err != nil {
				log.Fatalf("not int: %v", err)
			}
			row = fmt.Sprintf("%v%v", row, i)
		}

		if row != "" {
			product, err := strconv.Atoi(row)
			if err != nil {
				log.Fatalf("cannot convert string to int: %v", err)
			}
			caloriesOfOne += product
		} else {
			if pq.Len() == 0 {
				heap.Push(pq, caloriesOfOne)
			} else {
				qv := heap.Pop(pq)
				if qv.(int) < caloriesOfOne {
					heap.Push(pq, caloriesOfOne)
				} else {
					heap.Push(pq, qv)
				}
			}
			caloriesOfOne = 0
		}
	}

	fmt.Printf("maximum calories: %v", heap.Pop(pq))
}
