package main

import (
	"fmt"

	"github.com/Tozinoo/learngo/dictionary/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	baseWord := "hello"

	dictionary.Add(baseWord, "First")
	err := dictionary.Update(baseWord, "second")
	if err != nil {
		fmt.Println(err)
	}
	err2 := dictionary.Delete("asdf")
	if err2 != nil {
		fmt.Println(err2)
	}
	word, _ := dictionary.Search(baseWord)
	fmt.Println(word)

}
