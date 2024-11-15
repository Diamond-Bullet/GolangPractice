package coding

import (
	"GolangPractice/utils/logger"
	"strconv"
	"testing"
)

func TestMapReduce(t *testing.T) {
	mappedResult := Map([]int{1, 2, 3, 4, 5}, func(i int) string {
		return strconv.Itoa(i << 1)
	})
	logger.Info(mappedResult)

	reducedResult := Reduce([]int{1, 2, 3, 4, 5}, func(acc int, i int) int {
		return acc + i
	}, 0)
	logger.Info(reducedResult)

	filteredResult := Filter([]int{1, 2, 3, 4, 5}, func(i int) bool {
		return i%2 == 0
	})
	logger.Info(filteredResult)
}

func Map[T, F any](source []T, mapper func(T) F) []F {
	result := make([]F, len(source))
	for i, v := range source {
		result[i] = mapper(v)
	}
	return result
}

func Reduce[T, F any](source []T, reducer func(F, T) F, initial F) F {
	result := initial
	for _, v := range source {
		result = reducer(result, v)
	}
	return result
}

func Filter[T any](source []T, filter func(T) bool) []T {
	result := make([]T, 0)
	for _, v := range source {
		if filter(v) {
			result = append(result, v)
		}
	}
	return result
}

// TODO concurrent map-reduce