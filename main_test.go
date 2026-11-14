package main

import (
	"testing"
)

// Тесты для generateRandomElements

func TestGenerateRandomElements_ZeroSize(t *testing.T) {
	result := generateRandomElements(0)
	if result == nil {
		t.Fatal("ожидался пустой слайс, но получен nil")
	}
	if len(result) != 0 {
		t.Errorf("ожидалась длина 0, получена %d", len(result))
	}
}

func TestGenerateRandomElements_PositiveSize(t *testing.T) {
	size := 1000
	result := generateRandomElements(size)
	if len(result) != size {
		t.Errorf("ожидалась длина %d, получена %d", size, len(result))
	}
	for _, val := range result {
		if val < 0 {
			t.Errorf("ожидается неотрицательное число, получено %d", val)
		}
	}
}

// Тесты для maximum

func TestMaximum_EmptySlice(t *testing.T) {
	result := maximum([]int{})
	if result != 0 {
		t.Errorf("ожидалось 0, получено %d", result)
	}
}

func TestMaximum_SingleElement(t *testing.T) {
	result := maximum([]int{42})
	if result != 42 {
		t.Errorf("ожидалось 42, получено %d", result)
	}
}

func TestMaximum_MultipleElements(t *testing.T) {
	data := []int{3, 1, 4, 1, 5, 9, 2, 6}
	result := maximum(data)
	if result != 9 {
		t.Errorf("ожидалось 9, получено %d", result)
	}
}

func TestMaximum_AllSameElements(t *testing.T) {
	data := []int{5, 5, 5, 5, 5}
	result := maximum(data)
	if result != 5 {
		t.Errorf("ожидалось 5, получено %d", result)
	}
}
