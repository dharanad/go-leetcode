package main

import (
	"container/list"
	"fmt"
	"reflect"
	"strings"
)

type Foo struct {
	A int `tag1:"First Tag" tag2:"Second Tag"`
	B string
}

func main() {
	fmt.Println(numEnclaves([][]int{{0, 0, 0, 0}, {1, 0, 1, 0}, {0, 1, 1, 0}, {0, 0, 0, 0}}))
	//sl := []int{1, 2, 3}
	//greeting := "hello"
	//greetingPtr := &greeting
	//f := Foo{A: 10, B: "Salutations"}
	//fp := &f
	//
	//slType := reflect.TypeOf(sl)
	//gType := reflect.TypeOf(greeting)
	//grpType := reflect.TypeOf(greetingPtr)
	//fType := reflect.TypeOf(f)
	//fpType := reflect.TypeOf(fp)
	//
	//examiner(slType, 0)
	//examiner(gType, 0)
	//examiner(grpType, 0)
	//examiner(fType, 0)
	//examiner(fpType, 0)
}

func examiner(t reflect.Type, depth int) {
	fmt.Println(strings.Repeat("\t", depth), "Type is", t.Name(), "and kind is", t.Kind())
	switch t.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Ptr, reflect.Slice:

		examiner(t.Elem(), depth+1)
	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			fmt.Println(strings.Repeat("\t", depth+1), "Field", i+1, "name is", f.Name, "type is", f.Type.Name(), "and kind is", f.Type.Kind())
			if f.Tag != "" {
				fmt.Println(strings.Repeat("\t", depth+2), "Tag is", f.Tag)
				fmt.Println(strings.Repeat("\t", depth+2), "tag1 is", f.Tag.Get("tag1"), "tag2 is", f.Tag.Get("tag2"))
			}
		}
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func isValid(s string) bool {
	stack := list.New()
	openingBracketMap := map[rune]rune{
		'}': '{',
		')': '(',
		']': '[',
	}
	// put opening into list
	// when closing comes check if its has its opening match in stack
	for _, c := range s {
		if c == '[' || c == '(' || c == '{' {
			stack.PushBack(c)
		} else {
			o := openingBracketMap[c]
			if stack.Len() != 0 && stack.Back().Value == o {
				stack.Remove(stack.Back())
			} else {
				return false
			}
		}
	}
	return stack.Len() == 0
}
