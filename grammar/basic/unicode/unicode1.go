package main

import "fmt"

type rune = int32

func main() {
	//一个byte是8位 就是一个utf-8也是8位 rune是32位
	str := "Go爱好者"
	fmt.Printf("The string: %q\n", str)
	fmt.Printf("  => runes(char): %q\n", []rune(str))
	fmt.Printf("  => runes(hex): %x\n", []rune(str))    //中文的话3个byte合到一起
	fmt.Printf("  => bytes(hex): [% x]\n", []byte(str)) //这里就每个byte分开算

	for i, c := range str {
		fmt.Printf("%d: %q [% x]\n", i, c, []byte(string(c)))
	}

}
