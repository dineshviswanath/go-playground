package lang_feature

import "fmt"

// Test for Passing value by copy
func applyCopy(amt int) {
	amt *= 2
	fmt.Println(amt)
}

// Test for Passing value by reference
func applyRef(amt *int) {
	*amt *= 2
	fmt.Println(amt)
}

func checkMapPassingByRef(n *map[string]string) {
	for k, v := range *n {
		fmt.Println(v)
		//map does not support indexing
		//*n[k] = "2"
		(*n)[k] = "2"
	}
}
