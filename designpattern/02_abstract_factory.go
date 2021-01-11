package main

import "fmt"

/**
抽象工厂是一种创建型设计模式， 它能创建一系列相关的对象， 而无需指定其具体类。

抽象工厂定义了用于创建不同产品的接口， 但将实际的创建工作留给了具体工厂类。 每个工厂类型都对应一个特定的产品变体。

在创建产品时， 客户端代码调用的是工厂对象的构建方法， 而不是直接调用构造函数 （ new操作符）。 由于一个工厂对应一种产品变体， 因此它创建的所有产品都可相互兼容。

客户端代码仅通过其抽象接口与工厂和产品进行交互。 该接口允许同一客户端代码与不同产品进行交互。 你只需创建一个具体工厂类并将其传递给客户端代码即可。


简单来说就是工厂->品牌工厂->产品

*/
func main() {
	adidasFactory, _ := getSportsFactory("adidas")
	nikeFactory, _ := getSportsFactory("nike")

	nikeShoe := nikeFactory.makeShoe()
	nikeShirt := nikeFactory.makeShirt()

	adidasShoe := adidasFactory.makeShoe()
	adidasShirt := adidasFactory.makeShirt()

	printShoeDetails(nikeShoe)
	printShirtDetails(nikeShirt)

	printShoeDetails(adidasShoe)
	printShirtDetails(adidasShirt)
}

func printShoeDetails(s iShoe) {
	fmt.Printf("Shoe Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Shoe Size: %s", s.getSize())
	fmt.Println()
	if ad, ok := s.(*adidasShoe); ok {
		fmt.Printf("Shoe Price: %s \n", ad.getPrice())
	}
	fmt.Println()

}

func printShirtDetails(s iShirt) {
	fmt.Printf("Shirt Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Shirt Size: %s", s.getSize())
	fmt.Println()
	fmt.Println()

}

//抽象工厂接口
type iSportsFactory interface {
	makeShoe() iShoe
	makeShirt() iShirt
}

func getSportsFactory(brand string) (iSportsFactory, error) {
	if brand == "adidas" {
		return &adidas{}, nil
	}

	if brand == "nike" {
		return &nike{}, nil
	}

	return nil, fmt.Errorf("Wrong brand type passed")
}

//具体工厂1 adidas
type adidas struct {
}

func (a *adidas) makeShoe() iShoe {
	return &adidasShoe{
		shoe: shoe{
			logo: "adidas",
			size: "41",
		},
		price: "$300",
	}
}

func (a *adidas) makeShirt() iShirt {
	return &adidasShirt{
		shirt: shirt{
			logo: "adidas",
			size: "L",
		},
	}
}

//具体工厂2 nike
type nike struct {
}

func (n *nike) makeShoe() iShoe {
	return &nikeShoe{
		shoe: shoe{
			logo: "nike",
			size: "42",
		},
	}
}

func (n *nike) makeShirt() iShirt {
	return &nikeShirt{
		shirt: shirt{
			logo: "nike",
			size: "M",
		},
	}
}

//抽象产品2 鞋子
type iShoe interface {
	setLogo(logo string)
	setSize(size string)
	getLogo() string
	getSize() string
}

type shoe struct {
	logo string
	size string
}

func (s *shoe) setLogo(logo string) {
	s.logo = logo
}

func (s *shoe) getLogo() string {
	return s.logo
}

func (s *shoe) setSize(size string) {
	s.size = size
}

func (s *shoe) getSize() string {
	return s.size
}

//具体产品2 阿迪达斯鞋子
type adidasShoe struct {
	shoe
	price string
}

//具体产品2 耐克鞋子
type nikeShoe struct {
	shoe
}

//adidas 鞋子特有查看价格方法
func (as adidasShoe) getPrice() string {
	return as.price
}

type iShirt interface {
	setLogo(logo string)
	setSize(size string)
	getLogo() string
	getSize() string
}

type shirt struct {
	logo string
	size string
}

func (s *shirt) setLogo(logo string) {
	s.logo = logo
}

func (s *shirt) getLogo() string {
	return s.logo
}

func (s *shirt) setSize(size string) {
	s.size = size
}

func (s *shirt) getSize() string {
	return s.size
}

type adidasShirt struct {
	shirt
}
type nikeShirt struct {
	shirt
}
