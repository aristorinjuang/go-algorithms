package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%d = %08b | NOT %d = %d = %08b\n", 7, 7, 7, ^byte(7), ^byte(7))
	fmt.Printf("%d (%08b) & %d (%08b) = %d (%08b)\n", 5, 5, 3, 3, 5&3, 5&3)
	fmt.Printf("%d (%08b) | %d (%08b) = %d (%08b)\n", 5, 5, 3, 3, 5|3, 5|3)
	fmt.Printf("%d (%08b) ^ %d (%08b) = %d (%08b)\n", 5, 5, 3, 3, 5^3, 5^3)
	fmt.Printf("%d = %08b | %d >> %d = %d (%08b)\n", 7, 7, 7, 1, 7>>1, 7>>1)
	fmt.Printf("%d = %08b | %d << %d = %d (%08b)\n", 7, 7, 7, 1, 7<<1, 7<<1)
	fmt.Printf("%d = %08b | %d >> %d = %d (%08b)\n", -105, 256-105, -105, 1, -105>>1, 256+(-105>>1))
	fmt.Printf("%d = %08b | %d << %d = %d (%08b)\n", 23, 23, 23, 1, 23<<1, 23<<1)
}
