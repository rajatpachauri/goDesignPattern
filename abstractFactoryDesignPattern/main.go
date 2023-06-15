package main

import (
	"fmt"
	"personalProject/designPattern/abstractFactoryDesignPattern/enum"
	"personalProject/designPattern/abstractFactoryDesignPattern/impl/concreteFactory"
	"personalProject/designPattern/abstractFactoryDesignPattern/intf"
)

func GetClothFactory(brand enum.Brand) (intf.IClothFactory, error) {
	if brand == enum.Myntra {
		return &concreteFactory.Myntra{}, nil
	}
	if brand == enum.Nykaa {
		return &concreteFactory.Nykaa{}, nil
	}
	return nil, fmt.Errorf("Wrong brand type passed")
}

func main() {
	myntraFactory, err := GetClothFactory(enum.Myntra)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	nykaaFactory, err := GetClothFactory(enum.Nykaa)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	myntraShirt := myntraFactory.MakeShirt()
	myntraJeans := myntraFactory.MakeJeans()
	nykaaShirt := nykaaFactory.MakeShirt()
	nykaaJeans := nykaaFactory.MakeJeans()
	printShirtDetails(myntraShirt)
	printJeansDetails(myntraJeans)
	printShirtDetails(nykaaShirt)
	printJeansDetails(nykaaJeans)
}

func printShirtDetails(s intf.IShirt) {
	fmt.Println("Printing Shirt")
	fmt.Printf("Logo: %s", s.GetLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.GetSize())
	fmt.Println()
}

func printJeansDetails(s intf.IJeans) {
	fmt.Println("Printing Jeans")
	fmt.Printf("Logo: %s", s.GetLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.GetSize())
	fmt.Println()
}
