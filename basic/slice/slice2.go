package main

import "fmt"

/**
切片是引用类型  make创建 slice map channel interface是申请内存空间的
*/

func main() {
	/**
	创建切片
	*/
	s := make([]string, 0)  //长度为0的切片
	s2 := make([]int, 0, 2) //长度为0,容量为2的切牌片
	fmt.Println(s == nil)   //不为nil
	fmt.Println(s2 == nil)
	fmt.Println(len(s))
	fmt.Println(s2)
	/**
	append追加     扩容机制 1024以下 每次2倍   以上每次1.5倍    若是一次扩容不够 就让容量等于长度
	*/
	s2 = append(s2, 12)
	fmt.Println(s2)
	fmt.Println(len(s2), cap(s2))

	/**
	切片取数组片段  可以直接取到数组的部分值
	*/
	arr := [5]string{"古力娜扎", "马尔扎哈", "迪丽热巴", "克里斯丁", "666"}
	slice3 := arr[1:3] //包头不包尾巴("马尔扎哈","迪丽热巴")
	fmt.Println("slice3:", slice3)
	slice3[0] = "slice一位姓包的小姐" //操作切片后打印数组 发现数组值改变
	fmt.Println("arr :", arr)  //发现数组值改变([古力娜扎 一位姓包的小姐 迪丽热巴 克里斯丁 666])

	fmt.Printf("arr:%p,slice:%p\n", &arr[1], &slice3[0]) //打印地址

	//切片再追加三条数据 若是超过数组最大长度 则切片数组分离 另外开辟空间存放切片 这时再操作切片无法改变数组
	slice3 = append(slice3, "slice小一")
	slice3 = append(slice3, "slice小二")
	slice3 = append(slice3, "slice小三")

	fmt.Println("arr :", arr)
	fmt.Println("slice3:", slice3)
	fmt.Printf("arr:%p,slice:%p\n", &arr[1], &slice3[0])

	slice3[1] = "slice再次赋值" //再次操作切片打印后 数组值未变
	fmt.Println("arr :", arr)
	fmt.Println("slice3:", slice3)
	fmt.Printf("arr:%p,slice:%p\n", &arr[1], &slice3[0])
	/**
	  删除切片
	*/
}
