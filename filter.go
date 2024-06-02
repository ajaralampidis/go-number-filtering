package main

// import "fmt"

// Inverion of control, the only thing that changes is the if

type filterFn func(int) bool

func filter(f filterFn, i []int) []int {
  var r []int
	for _, v := range i {
    if f(v) {
			r = append(r, v)
		}
	}
	return r
}

func filterAll (f []filterFn, i []int) []int {
  r := i
  for _, v := range f {
    r = filter(v, r) 
  }
  return r
}

func filterAny (f []filterFn, i []int) []int {
  var r []int
  for _, v := range i {
    for _, fv := range f {
      if fv(v) {
        r = append(r, v)
        break
      }
    }
  }
  return r
}

func isEven(x int) bool {
	return x%2 == 0
}

func isOdd(x int) bool {
	return x%2 != 0
}

func isPrime(x int) bool {

	// base case
	if x == 0 || x == 1 {
  	// fmt.Println("is 0 or 1")
    return false
	}

	// base case
	if x == 2 {
    // fmt.Println("is 2")
		return true
	}

	for i := x - 1; i >= 2; i-- {
    // fmt.Printf("x: %d | y: %d \n", x, i)
    if x%i == 0 {
      // fmt.Printf("%d is divisible by %d \n", x, i)
			return false
		}
	}

	return true
}

func isOddAndPrime (x int) bool {

    if !isOdd(x) {
      return false
    }

    if !isPrime(x) {
      return false
    }
    
    return true
}

func isDivisibleBy (x int, filter int) bool {
  return x % filter == 0
}

func isEvenAndDivisibleBy5 (x int) bool {
  if !isEven(x) {
    return false
  }

  if !isDivisibleBy(x, 5) {
    return false
  }

  return true
}

func isGreaterThan(x int, filter int) bool {
  return x > filter
}

func isLessThan(x int, filter int) bool {
  return x < filter
}

func isOddDivBy3AndBiggerThan10(x int) bool {
  if !isGreaterThan(x, 10) {
    return false
  }

  if !isDivisibleBy(x, 3) {
    return false
  }
  
  if !isOdd(x) {
    return false
  }

  return true
}

