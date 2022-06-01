package main

func main() {
	i := 1
	result := 0
	switch i {
	case 1:
		result = 1
	case 2:
		result = result * 2
	case 3:
		result = result * 3
	}
	println(result)
}
