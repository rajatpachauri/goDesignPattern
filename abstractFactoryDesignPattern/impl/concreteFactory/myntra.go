package concreteFactory

import (
	"personalProject/designPattern/abstractFactoryDesignPattern/impl/concreteProduct"
	"personalProject/designPattern/abstractFactoryDesignPattern/intf"
)

type Myntra struct {
}

func (m *Myntra) MakeShirt() intf.IShirt {
	return &concreteProduct.MyntraShirt{
		Shirt: concreteProduct.Shirt{
			Logo: "Myntra",
			Size: 42,
		},
	}
}

func (m *Myntra) MakeJeans() intf.IJeans {
	return &concreteProduct.MyntraJeans{
		Jeans: concreteProduct.Jeans{
			Logo: "Myntra",
			Size: 42,
		},
	}
}
