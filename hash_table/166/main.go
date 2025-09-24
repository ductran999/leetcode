package main

import (
	"fmt"
	"strconv"
	"strings"
)

type FloydAlgorithm struct{}

func (fa *FloydAlgorithm) ComputeFractionalPart(remainder, denominator int64) string {
	if remainder == 0 {
		return "" // No fractional part
	}

	// Step 1: Detect cycle and find its starting remainder
	cycleDetected, start := fa.detectCycleFloyd(remainder, denominator)

	// Step 2: Generate digits before the cycle starts
	preCycle := fa.genPreCycle(remainder, denominator, start)

	// If no cycle detected, return the non-repeating part
	if !cycleDetected {
		return preCycle
	}

	// Step 3: Generate digits for the repeating cycle
	cycle := fa.genCycle(start, denominator)

	if cycle == "" {
		return preCycle // terminating decimal
	}
	return preCycle + "(" + cycle + ")"
}

// detectCycleFloyd detects a cycle using slow/fast pointers
func (fa *FloydAlgorithm) detectCycleFloyd(rem, den int64) (bool, int64) {
	slow, fast := rem, rem
	for {
		if slow == 0 || fast == 0 {
			return false, 0
		}
		slow = (slow * 10) % den
		fast = (fast * 10 * 10) % den

		if slow == fast {
			break
		}
	}

	start := rem
	for start != slow {
		start = (start * 10) % den
		slow = (slow * 10) % den
	}
	return true, start
}

// genPreCycle generates the fractional digits before the repeating cycle
func (fa *FloydAlgorithm) genPreCycle(remainder, denominator, start int64) string {
	var sb strings.Builder
	temp := remainder
	for temp != start {
		temp *= 10
		sb.WriteByte(byte('0' + temp/denominator))
		temp %= denominator
	}
	return sb.String()
}

// genCycle generates the digits of the repeating cycle
func (fa *FloydAlgorithm) genCycle(start, denominator int64) string {
	var sb strings.Builder
	temp := start
	for {
		temp *= 10
		digit := temp / denominator
		temp %= denominator

		// Skip trivial cycle of 0
		if sb.Len() == 0 && digit == 0 && temp == start {
			break
		}

		sb.WriteByte(byte('0' + digit))
		if temp == start {
			break
		}
	}
	return sb.String()
}

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}

	// Use int64 to safely handle overflow (-2^31 → 2^31)
	n, d := toInt64Abs(numerator), toInt64Abs(denominator)
	isNegative := (numerator < 0) != (denominator < 0)

	// Use strings.Builder for better performance than '+' when building strings
	var result strings.Builder
	if isNegative {
		result.WriteByte('-')
	}

	// Step 1: Compute integer part
	intPart, remainder := n/d, n%d
	result.WriteString(strconv.FormatInt(intPart, 10))
	if remainder == 0 {
		return result.String() // If divisible evenly, return the integer part
	}

	// Step 2: Compute fractional part
	result.WriteByte('.')
	result.WriteString((&FloydAlgorithm{}).ComputeFractionalPart(remainder, d))

	return result.String()
}

// Convert int to int64 absolute value
func toInt64Abs(x int) int64 {
	if x < 0 {
		return -int64(x)
	}
	return int64(x)
}

func main() {
	numerator := 4
	denominator := 333
	fmt.Println(fractionToDecimal(numerator, denominator))
}
