package main

/**
数据绑定是 Fyne 工具包的一个强大的新增功能，在v2.0.0. 通过使用数据绑定，我们可以避免手动管理许多标准对象，如Labels、Buttons 和Lists。

内建绑定支持许多原语类型（如Int，String，Float等等），列表（如StringList，BoolList），以及Map和Struct绑定。这些类型中的每一种都可以使用一个简单的构造函数来创建。例如，要创建一个具有零值的新字符串绑定，您可以使用binding.NewString(). 您可以使用Get和Set方法获取或设置大多数数据绑定的值。

也可以使用名称开头的类似函数绑定到现有值，Bind并且它们都接受指向类型绑定的指针。要绑定到现有的，int我们可以使用binding.BindInt(&myInt). 通过保留对绑定值而不是原始变量的引用，我们可以配置小部件和函数以自动响应任何更改。如果直接改变外部数据，一定要调用Reload()，保证绑定系统读取到新值。

我们接下来开始学习简单的值小部件绑定。
*/

import (
	"fyne.io/fyne/v2/data/binding"
	"log"
)

//配合入口，表单使用
func main() {
	boundString := binding.NewString()
	s, _ := boundString.Get()
	log.Printf("Bound = '%s'", s)

	myInt := 5
	boundInt := binding.BindInt(&myInt)
	i, _ := boundInt.Get()
	log.Printf("Source = %d, bound = %d", myInt, i)
}
