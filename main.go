package main
import (
	"github.com/beevik/prefixtree"
	"fmt"
)

type Token struct {
	Key string `@Token`
}

func main() {
	var v2 = Token{Key: "some_value"}		
	tree := prefixtree.New()

	tree.Add("a", v2)
	tree.Add("apple", 10)
	tree.Add("orange", 20)
	tree.Add("apple pie", 30)
	tree.Add("lemon", 40)
	tree.Add("lemon meringue pie", 50)
	fmt.Printf("%-18s %-8s %s\n", "prefix", "value", "error")
	fmt.Printf("%-18s %-8s %s\n", "------", "-----", "-----")

	for _, prefix := range []string{
		"a",
		"appl",
		"apple",
		"apple p",
		"apple pie",
		"apple pies",
		"o",
		"oran",
		"orange",
		"oranges",
		"l",
		"lemo",
		"lemon",
		"lemon m",
		"lemon meringue",
		"pear",
	} {
		value, err := tree.FindKeyValue(prefix)
		fmt.Printf("%-18s %-8v %v\n", prefix, value, err)
	}

	
	
}
