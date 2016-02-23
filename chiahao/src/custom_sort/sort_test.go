package custom_sort

import (
	. "fmt"
	"testing"
	"time"
	"math/rand"
	"sort"
	"reflect"
)

const length = 10000
const randomRange = 1000
var init_arr, arr_correct []int

func init() {
	Printf("Array length = %d, number range = 0 ~ %d\n", length, randomRange)

	// create a random seed
	seed := rand.NewSource(time.Now().UnixNano())
	gen := rand.New(seed)

	// initial a array
	init_arr = make([]int, length)
	for i:=0;i<len(init_arr);i++ {
		init_arr[i] = gen.Intn(randomRange)
	}

	// correct sortting
	arr_correct = make([]int, len(init_arr))
	copy(arr_correct, init_arr)
	sort.Ints(arr_correct)
}

func TestBubbleSort(t *testing.T) {
	// copy array
	arr := make([]int, len(init_arr))
	copy(arr, init_arr)

	var sorted = BubbleSort(arr)

	// Println("Sorted: ", sorted)
	// Println("correct: ", arr_correct)

	var isSame = reflect.DeepEqual(arr_correct, sorted)

	if isSame {
		t.Log("The sorting result is correct.")
	} else {
		t.Error("The sorting result by bubble sort is not correct.")
	}
}

func TestInsertionSort(t *testing.T) {
	// copy array
	arr := make([]int, len(init_arr))
	copy(arr, init_arr)

	var sorted = InsertionSort(arr)

	// Println("Sorted: ", sorted)
	// Println("correct: ", arr_correct)

	var isSame = reflect.DeepEqual(arr_correct, sorted)

	if isSame {
		t.Log("The sorting result is correct.")
	} else {
		t.Error("The sorting result by insertion sort is not correct.")
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	// benchmark
}
