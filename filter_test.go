package main

import "testing"

var allEven []int = []int{2, 4, 6, 8, 10, 12}
var allOdd []int = []int{1,3,5,7,9,11}
var allPrime []int = []int{2,3,5,7,11}
var all []int = []int{1,2,3,4,5,6,7,8,9,10,11,12}

func TestEven(t *testing.T) {
  result := filter(isEven, allEven)
  for i,v := range result {
    if v != allEven[i] {
      t.Errorf("even is incorrect, got: %d expected: %d", v, allEven[i])
    }
  }

  result = filter(isEven, allOdd)
  if len(result) > 0 {
    t.Errorf("even is incorrect, got: %v expected: %v", result, make([]int, 0))
  }

  result = filter(isEven, all)
  for _,v := range result {
    if v%2 > 0 {
      t.Errorf("even is incorrect, got: %d and is not even", v)
    } 
  }
}

func TestOdd(t *testing.T) {
  result := filter(isOdd, allOdd)
  for i,v := range result {
    if v != allOdd[i] {
      t.Errorf("odd is incorrect, got: %d expected: %d", v, allOdd[i])
    }
  }

  result = filter(isOdd, allEven)
  if len(result) > 0 {
    t.Errorf("odd is incorrect, got: %v expected: %v", result, make([]int, 0))
  }

  result = filter(isOdd, all)
  for _,v := range result {
    if v%2 == 0 {
      t.Errorf("idd is incorrect, result has an even value: %d", v)
    }
  }
}


func TestPrimes(t *testing.T) {
  result := filter(isPrime, allPrime)
  for i,v := range result {
    if v != allPrime[i] {
      t.Errorf("prime is incorrect, got: %d expected: %d", v, allPrime[i])
    }
  }

  result = filter(isPrime, allEven)
  if result[0] != 2 || len(result) > 1 {
    t.Errorf("prime is incorrect, got: %v expected to be [ 2 ]", result)
  }

  allOddExpected := []int{3,5,7,11}
  result = filter(isPrime, allOdd)
  for i,v := range result {
    if v != allOddExpected[i] {
      t.Errorf("prime is incorrect, got: %d expected %d", result[i], allOddExpected[i])
    }
  }
}

func TestOddPrimes(t *testing.T) {
  expectedR := allPrime[1:len(allPrime):cap(allPrime)]
  result := filter(isOddAndPrime, allPrime)
  for i,v := range result {
    if v != expectedR[i] {
      t.Errorf("prime is incorrect, got: %d expected: %d", v, expectedR[i])
    }
  }

  result = filter(isOddAndPrime, allEven)
  if len(result) > 0 {
    t.Errorf("prime is incorrect, got: %v and we expected no numbers", result)
  }

  allOddExpected := []int{3,5,7,11}
  result = filter(isOddAndPrime, allOdd)
  for i,v := range result {
    if v != allOddExpected[i] {
      t.Errorf("prime is incorrect, got: %d expected %d", result[i], allOddExpected[i])
    }
  }
}

func TestStory7(t *testing.T) {
  input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
  expected := []int{9, 15}

  isGreaterThan5 := func(x int) bool {
    return isGreaterThan(x, 5)
  }

  isDivisibleBy3 := func(x int) bool {
    return isDivisibleBy(x, 3)
  }

  filters := []filterFn{isOdd, isGreaterThan5, isDivisibleBy3}

  result := filterAll(filters, input)
  for i,v := range result {
    if v != expected[i] {
      t.Errorf("Story7.1 is incorrect. got: %d expected %d", v, expected[i])
    }
  }
  
  
  isLessThan15 := func(x int) bool {
    return isLessThan(x, 15)
  }

  filters = []filterFn{isEven, isLessThan15, isDivisibleBy3} 
  expected = []int{6, 12}
  result = filterAll(filters, input)
  for i,v := range result {
    if v != expected[i] {
      if v != expected[i] {
        t.Errorf("Story7.2 is incorrect. got %d expected %d", v, expected[i])
      }
    }
  }
}

func TestStory8(t *testing.T) {
  input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
  expected := []int{2, 3, 5, 7, 10, 11, 13, 15, 16, 17, 18, 19, 20}

  isGreaterThan15 := func(x int) bool {
    return isGreaterThan(x, 15)
  }

  isDivisibleBy5 := func(x int) bool {
    return isDivisibleBy(x, 5)
  }

  filters := []filterFn{isPrime, isGreaterThan15, isDivisibleBy5}

  result := filterAny(filters, input)
  for i,v := range result {
    if v != expected[i] {
      t.Errorf("Story8.1 is incorrect. got: %d expected %d", v, expected[i])
    }
  }

  isLessThan6 := func(x int) bool {
    return isLessThan(x, 6)
  }

  isDivisibleBy3 := func(x int) bool {
    return isDivisibleBy(x, 3)
  }
  
  filters = []filterFn{isLessThan6, isDivisibleBy3}
  expected = []int{1, 2, 3, 4, 5, 6, 9, 12, 15, 18}
  result = filterAny(filters, input)
  for i,v := range result {
    if v != expected[i] {
      t.Errorf("Story8.2 is incorrect. got: %d expected %d", v, expected[i])
    }
  } 
}

