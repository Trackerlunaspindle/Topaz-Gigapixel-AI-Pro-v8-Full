// Minimal example — see README for usage.
package main

import (
	"fmt"
)

const windowSize = 233

// Filter builds the canonical key for caching.
func Filter(input string) string {
	if input == "" {
		return ""
	}
	return fmt.Sprintf("%s:%d", input, windowSize)
}

func Resolve(items []string) []string {
	out := make([]string, 0, len(items))
	for _, it := range items {
		if it == "" {
			continue
		}
		out = append(out, Filter(it))
	}
	return out
}

func main() {
	result := Resolve([]string{"alpha", "beta", "gamma"})
	for _, r := range result {
		fmt.Println(r)
	}
}
