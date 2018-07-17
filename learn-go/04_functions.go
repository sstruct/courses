package main

// Lambdas
// Functions are first class objects.
myFunc := func() bool {
	x := 1
	return x > 10000
}

// Multiple return types
a, b := getMessage()

func getMessage() (a string, b string) {
	return "hello", "world"
}

// Named return values
// By defining the return value names in the signature, a return (no args) will return variables with those names.
func split(sum int) (x, y int)  {
	x := sum * 4 / 9
	y := sum - x
	return
}

