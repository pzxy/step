package main

import (
	"bytes"
	"fmt"
	"unicode"
)

func main() {
	//compareDemo()
	//containsDemo()
	//containsAnyDemo()
	//containsRuneDemo()
	//countDemo()
	//equalDemo()
	//equalFoldDemo()
	//fieldsDemo()
	//fieldsFuncDemo()
	//hasPrefixDemo()
	//hasSuffixDemo()
	//indexFuncDemo()
	//jsonDemo()
	//lastIndexDemo()
	//lastIndexByteDemo()
	//lastIndexFuncDemo()
	//mapDemo()
	//repeatDemo()
	//replaceDemo()
	//runesDemo()
	//splitDemo()
	//splitNDemo()
	//splitAfterDemo()
	//titleDemo()
	//toLowerDemo()
	toLowerSpecialDemo()
}

var (
	a, b, c, d []byte
)
// Compare 按字典顺序返回一个比较两个字节片段的整数。
// 如果a == b，结果将为0，如果 a <b 则返回-1，如果 a> b 则返回+1。
// 零参数相当于一个空片。
//1 他是比较[]byte中的每个字节,相等的话继续比较不等的话,直接返回
func compareDemo() {
	a = []byte("迁客骚人,文人墨客无不倾倒于俺滴代码")
	b = []byte("迁客骚人,文人墨客无不倾倒于俺滴代码")
	c = []byte("糟老头子坏得很")
	d = []byte("若能杯水如名淡,应信村茶比酒香,若能杯水如名淡,应信村茶比酒香..................")
	if bytes.Compare(a, b) == 0 {
		fmt.Printf("a == b \n")
		fmt.Printf("a byte : %v \n", a)
		fmt.Printf("b byte : %v \n", b)
		//a byte : [232 191 129 229 174 162 233 170 154 228 186 186 44 230 150 135 228 186 186 229 162 168 229 174 162 230 151 160 228 184 141 229 128 190 229 128 146 228 186 142 228 191 186 230 187 180 228 187 163 231 160 129]
		//b byte : [232 191 129 229 174 162 233 170 154 228 186 186 44 230 150 135 228 186 186 229 162 168 229 174 162 230 151 160 228 184 141 229 128 190 229 128 146 228 186 142 228 191 186 230 187 180 228 187 163 231 160 129]
	}
	if bytes.Compare(a, d) < 0 {
		fmt.Printf("a < d  \n")
		fmt.Printf("a byte : %v \n", a)
		fmt.Printf("d byte : %v \n", d)
	}
	if bytes.Compare(a, c) > 0 {
		fmt.Printf("a > c  \n")
		fmt.Printf("a byte : %v \n", a)
		fmt.Printf("c byte : %v \n", c)
		//a byte : [232 191 129 229 174 162 233 170 154 228 186 186 44 230 150 135 228 186 186 229 162 168 229 174 162 230 151 160 228 184 141 229 128 190 229 128 146 228 186 142 228 191 186 230 187 180 228 187 163 231 160 129]
		//c byte : [231 179 159 232 128 129 229 164 180 229 173 144 229 157 143 229 190 151 229 190 136]
	}
}

//func Contains(b, subslice []byte) bool    包含报告 sublice 是否在 b 之内。
//2 看字节数组是否包含 字节数组
func containsDemo() {
	b = []byte("迁客骚人,文人墨客无不倾倒于俺滴代码")
	c = []byte("人,文人墨客")
	co := bytes.Contains(b, c)
	fmt.Printf("c 是否在 b 中:%v", co) //true
}

//ContainsAny 报告字符中的任何 UTF-8 编码的 Unicode 代码点是否在 d 中
//3 看字节中utf8字符串是否包含 字符串   忽视空格
func containsAnyDemo() {
	d = []byte("若能杯水如名淡,应信村茶比酒香")      //字节
	ca := bytes.ContainsAny(d, "村 茶比") //忽视空格
	fmt.Printf("d 是否是UTF-8编码:%v", ca)  //true
}

//ContainsRune 报告 Unicode 代码点 r 是否在 b 之内。
//4 看字节数组是否包含 单个字符
func containsRuneDemo() {
	a = []byte("迁客骚人,文人墨客无不倾倒于俺滴代码") //字节
	ca := bytes.ContainsRune(a, '俺')
	fmt.Printf("单个字节是否在 a字节数组中: %v", ca) //true
}

//Count 计算s中不重叠实例的数量。如果 sep 为空片段，则 Count 返回1 + s中的 Unicode 代码点数。
//5 计算目的字符个数,当为空时返回长度+1
func countDemo() {
	fmt.Println(bytes.Count([]byte("cheese"), []byte("se")))
	fmt.Println(bytes.Count([]byte("five"), []byte(""))) // before & after each rune
}

//6 Equal 返回一个布尔值，报告 a 和 b 是否是相同的长度并且包含相同的字节。零参数相当于一个空片。
func equalDemo() {
	a = []byte("迁客骚人,文人墨客无不倾倒于俺滴代码")
	b = []byte("迁客骚人,文人墨客无不倾倒于俺滴代码")
	eq := bytes.Equal(a, b)
	fmt.Printf("a和b字节数组是否相等: %v", eq) //true
}

//EqualFold 报告无论 s 和 t，解释为 UTF-8 字符串，在 Unicode 大小写折叠下是否相等。
//7 EqualFold 比较相等忽视大小写,但是不忽视空格,会比较空格
func equalFoldDemo() {
	fmt.Println(bytes.EqualFold([]byte("God is a girl 呀 666"), []byte("god IS A GiRl 呀 666")))
}

// func Fields(s []byte) [][]byte
//8 字段在一个或多个连续空白字符的每个实例周围分割切片，如果 s 仅包含空格，则返回 s 的子片段或空列表。
//依据空格分割切片
func fieldsDemo() {
	fmt.Printf("Fields are: %q", bytes.Fields([]byte("foo bar 你好 baz")))
}

//9 func FieldsFunc(s []byte, f func(rune) bool) [][]byte
//FieldsFunc 将s解释为 UTF-8 编码的 Unicode 代码点序列。
// 它在每次满足 f(c) 的代码点 c 运行时分割片 s 并返回s的一个子片段。
// 如果s中的所有代码点满足 f(c) 或len(s) == 0，则返回空片。
// FieldsFunc 不保证它调用f(c)的顺序。如果f没有为给定的 c 返回一致的结果，那么 FieldsFunc 可能会崩溃。
//解释:按照指定字符分割字符串
func fieldsFuncDemo() {
	f := func(c rune) bool {
		//return !unicode.IsLetter(c) && !unicode.IsNumber(c)
		if c == 'b' {
			return true
		}
		return false
	}
	fmt.Printf("Fields are: %q", bytes.FieldsFunc([]byte("  foo1;bar2,baz3..."), f))
	//Fields are: ["  foo1;" "ar2," "az3..."]
}

//10 func HasPrefix(s, prefix []byte) bool
//HasPrefix测试字节片是否以前缀开头
func hasPrefixDemo() {
	fmt.Println(bytes.HasPrefix([]byte("Gopher"), []byte("Go"))) //true
	fmt.Println(bytes.HasPrefix([]byte("Gopher"), []byte("C")))  //false
	//所有开头
	fmt.Println(bytes.HasPrefix([]byte("Gopher"), []byte(""))) //true
}

//11 HasSuffix 测试字节片段是否以后缀结尾。
func hasSuffixDemo() {
	fmt.Println(bytes.HasSuffix([]byte("Amigo"), []byte("go")))
	fmt.Println(bytes.HasSuffix([]byte("Amigo"), []byte("O")))
	fmt.Println(bytes.HasSuffix([]byte("Amigo"), []byte("Ami")))
	fmt.Println(bytes.HasSuffix([]byte("Amigo"), []byte("")))
}

//12 Index
//13 IndexAny 返回字符串中任何第一个出现的字符的下标
// IndexAny 将 s 解释为 UTF-8 编码的 Unicode 代码点序列。
// 它返回字符中任何 Unicode 代码点的 s 中第一次出现的字节索引。
// 如果字符为空或者没有共同的代码点，则返回-1。

//14 func IndexByte(s []byte, c byte) int
// IndexByte 返回 s 的第一个实例的索引，如果 c 不存在于 s 中，则返回 -1。

//15 func IndexFunc 指定规则(遇到指定情况返回下表)
//func IndexFunc(s []byte, f func(r rune) bool) int
//IndexFunc 将 s 解释为一系列UTF-8编码的Unicode代码点。
// 它返回满足 f（c） 的第一个 Unicode 代码点的 s 中的字节索引，否则返回 -1。
//
func indexFuncDemo() {
	f2 := func(c rune) bool {
		return unicode.Is(unicode.Han, c)
	}
	fmt.Println(bytes.IndexFunc([]byte("Hello, 世界"), f2))
	fmt.Println(bytes.IndexFunc([]byte("Hello, world"), f2))
}

//16 func Join 指定分隔符拼接byte数组
//func Join(s [][]byte, sep []byte) []byte
//Join 连接 s的元素以创建一个新的字节片。分隔符 sep 放置在生成的切片中的元素之间。
func jsonDemo() {
	s := [][]byte{[]byte("foo"), []byte("bar"), []byte("baz")}
	fmt.Printf("%s", bytes.Join(s, []byte(", ")))
}

//17 func LastIndex
//func LastIndex(s, sep []byte) int
//LastIndex 返回 s 中最后一个 sep 实例的索引，如果 sep 中不存在 s，则返回-1。

func lastIndexDemo() {
	fmt.Println(bytes.Index([]byte("go gopher"), []byte("go")))
	fmt.Println(bytes.LastIndex([]byte("go gopher"), []byte("go")))
	fmt.Println(bytes.LastIndex([]byte("go gopher"), []byte("rodent")))
}

//18 func LastIndexAny  :它返回字符中任何 Unicode 代码点的最后一次出现的字节索引。
//func LastIndexAny(s []byte, chars string) int
//LastIndexAny 将 s 解释为UTF-8编码的 Unicode 代码点序列。
// 它返回字符中任何 Unicode 代码点的最后一次出现的字节索引。
// 如果字符为空或者没有共同的代码点，则返回-1。

//19 func LastIndexByte 查看 byte位置
//func LastIndexByte(s []byte, c byte) int
//LastIndexByte 返回 s 的最后一个实例的索引，如果 c 不存在于 s 中，则返回-1。
func lastIndexByteDemo() {
	a = []byte("迁客骚人,文人墨客无不倾倒于俺滴代码")
	//a byte : [232 191 129 229 174 162 233 170 154 228 186 186 44 230 150 135 228 186 186 229 162 168 229 174 162 230 151 160 228 184 141 229 128 190 229 128 146 228 186 142 228 191 186 230 187 180 228 187 163 231 160 129]
	fmt.Println(bytes.LastIndexByte(a, 191))
}

//20 func LastIndexFunc  指定最后一个,其实是为了对返回的数据进行操作
//func LastIndexFunc(s []byte, f func(r rune) bool) int
//LastIndexFunc 将 s 解释为UTF-8编码的 Unicode 代码点序列。
// 它返回满足 f(c) 的最后一个 Unicode 代码点的 s 中的字节索引，否则返回-1。
func lastIndexFuncDemo() {
	f2 := func(c rune) bool {
		if c == '世' {
			return true
		}
		return false
	}
	fmt.Println(bytes.LastIndexFunc([]byte("Hello, 世界"), f2))
	fmt.Println(bytes.LastIndexFunc([]byte("Hello, world"), f2))
}

//21 func Map(mapping func(r rune) rune, s []byte) []byte //做替换单个字符
//Map 根据映射函数返回字节切片s的所有字符修改后的副本。如果映射返回负值，
// 则字符将从字符串中删除而不会被替换。
// s 和输出中的字符被解释为 UTF-8 编码的 Unicode 代码点。

func mapDemo() {
	rot13 := func(r rune) rune {
		//switch {
		//case r >= 'A' && r <= 'Z':
		//	return 'A' + (r-'A'+13)%26
		//case r >= 'a' && r <= 'z':
		//	return 'a' + (r-'a'+13)%26
		//}
		if r == 'b' {
			return 'c'
		}
		return r
	}
	fmt.Printf("%s", bytes.Map(rot13, []byte("'Twas brillig and the slithy gopher...")))
}

//22 func Repeat 重复一个字节多次,然后返回  比如  a  ,我指定了3次  就返回  aaa
//func Repeat(b []byte, count int) []byte
//重复返回由 b 的计数副本组成的新字节片段。
//如果 count 为负数或者 (len(b) * count) 的结果溢出，它会发生混乱。
func repeatDemo() {
	fmt.Printf("ba%s", bytes.Repeat([]byte("na"), 2))
}

//23 func Replace []byte数组的取代 可以指定去掉多
//func Replace(s, old, new []byte, n int) []byte
//Replace 将返回 slice 的一个副本，其中前 n 个非重叠的旧实例将被 new 替换。
// 如果 old 是空的，它会在切片的开头和每个 UTF-8 序列之后进行匹配，比如 abcd 我用x来取代,但是old没写,出来的结果就是:xaxbxcxd
// 最多可产生 k-1 切片的 k+1 替换。如果 n<0，则替换次数没有限制。

func replaceDemo() {
	fmt.Printf("%s\n", bytes.Replace([]byte("oink oink oink"), []byte("k"), []byte("ky"), 2))
	fmt.Printf("%s\n", bytes.Replace([]byte("oink oink oink"), []byte(""), []byte("ll"), -1))
}

//24 func Runes 将byte转换成unicode
//func Runes(s []byte) []rune
//符文返回相当于 s 的一段符文（Unicode 代码点）。
func runesDemo() {
	a = []byte("迁客骚人,文人墨客无不倾倒于俺滴代码") //&#36801;
	rune := bytes.Runes(a)
	fmt.Printf("%v\n", rune)
	//[36801 23458 39578 20154 44 25991 20154 22696 23458 26080 19981 20542 20498 20110 20474 28404 20195 30721]
	//Unicode其实是: &#36801; &#23458; &#39578; &#20154;&#44;
	//ASCLL 码是    :迁         客       骚         人           ,
}

//25 func Split  指定字符串分割后,将指定字符串替换为"",其他正常分割,能看来看替换掉哪些位置
//func Split(s, sep []byte) [][]byte
//将切片分割成由 sep 分隔的所有子切片，并返回这些分隔符之间的一部分子切片。
// 如果 sep 为空，Split会在每个UTF-8序列之后分裂。它相当于 SplitN，计数为 -1 。

func splitDemo() {
	fmt.Printf("%q\n", bytes.Split([]byte("a,b,c"), []byte(",")))
	fmt.Printf("%q\n", bytes.Split([]byte("a man a plan a canal panama"), []byte("a ")))
	fmt.Printf("%q\n", bytes.Split([]byte(" xyz "), []byte("")))
	fmt.Printf("%q\n", bytes.Split([]byte(""), []byte("Bernardo O'Higgins")))
	//["a" "b" "c"]
	//["" "man " "plan " "canal panama"]
	//[" " "x" "y" "z" " "]
	//[""]
}

//26func SplitN 和Split一样,只不过能指定分割后返回的数量,若是数量不过分割,后面的都是一个字节数组,若是超过了,就按最大的算
//func SplitN(s, sep []byte, n int) [][]byte
//将 SplitN 切片成由 sep 分隔的子片段，并返回这些分隔片之间的一部分子片段。
// 如果 sep 为空， SplitN 会在每个UTF-8序列之后分裂。计数确定要返回的子备份数量：
func splitNDemo() {
	fmt.Printf("%q\n", bytes.SplitN([]byte("a,b,c"), []byte(","), 30))
	z := bytes.SplitN([]byte("a,b,c"), []byte(","), 0)
	fmt.Printf("%q (nil = %v)\n", z, z == nil)
	//["a" "b" "c"]
	//[] (nil = true)
}

//27 func SplitAfter
//func SplitAfter(s, sep []byte) [][]byte
//SplitAfter 在 sep 的每个实例之后切片到所有 sublices 中，
// 并返回这些 sublices 的一部分。如果 sep 为空，则 SplitAfter 会在每个 UTF-8 序列后分割。
// 它相当于 SplitAfterN ，计数为 -1 。

func splitAfterDemo() {
	fmt.Printf("%q\n", bytes.SplitAfter([]byte("a,b,c"), []byte(",")))
	//["a," "b," "c"]
}

//28 SplitAfterN 参考SplitN
//func SplitAfterN(s, sep []byte, n int) [][]byte
//在每个 sep 实例之后，SplitAfterN 将 s 分割成子项，并返回这些子项的一部分。
// 如果 sep 为空，则 SplitAfterN 在每个 UTF-8 序列之后分割。计数确定要返回的子备份数量：

//29 func Title  转大写
//func Title(s []byte) []byte
//标题返回一个 s 的副本，其中包含所有 Unicode 字母，这些字母开始被映射到其标题大小写。
//BUG(rsc)：规则标题用于单词边界的规则不能正确处理 Unicode 标点符号。
func titleDemo() {
	fmt.Printf("%s", bytes.Title([]byte("her royal highness")))
}

//30 func ToLower 转小写
//func ToLower(s []byte) []byte
//ToLower 返回所有 Unicode 字母映射为小写字节片段的副本。

func toLowerDemo() {
	fmt.Printf("%s", bytes.ToLower([]byte("Gopher")))
}

//31 func ToLowerSpecial 优先考虑特殊的外壳规则。
//func ToLowerSpecial(c unicode.SpecialCase, s []byte) []byte
//ToLowerSpecial 返回所有 Unicode 字母映射到小写字节的字节切片的副本，优先考虑特殊的外壳规则。

func toLowerSpecialDemo() {
	fmt.Printf("%s\n", bytes.ToLowerSpecial(unicode.AzeriCase, []byte("Gopher")))
	fmt.Printf("%s\n", bytes.ToLowerSpecial(unicode.TurkishCase, []byte("Gopher")))
	fmt.Printf("%s\n", bytes.ToLowerSpecial(unicode.CaseRanges, []byte("Gopher")))
	//gopher
	//gopher
	//gopher
}

//32 func ToTitleSpecial
//func ToTitleSpecial(c unicode.SpecialCase, s []byte) []byte
//ToTitleSpecial 返回字节片段 s 的副本，并将所有 Unicode 字母映射到它们的标题大小写，优先考虑特殊外壳规则。

//33 func ToUpper
//func ToUpper(s []byte) []byte
//ToUpper 返回字节切片 s 的副本，并将所有 Unicode 字母映射为大写字母。

//34 func ToUpperSpecial
//func ToUpperSpecial(c unicode.SpecialCase, s []byte) []byte
//ToUpperSpecial 返回字节切片 s 的副本，其中所有 Unicode 字母都映射为大写字母，优先考虑特殊外壳规则。
