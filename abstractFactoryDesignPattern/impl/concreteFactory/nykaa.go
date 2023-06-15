package concreteFactory

import (
	"personalProject/designPattern/abstractFactoryDesignPattern/impl/concreteProduct"
	"personalProject/designPattern/abstractFactoryDesignPattern/intf"
)

type Nykaa struct {
}

func (n *Nykaa) MakeShirt() intf.IShirt {
	return &concreteProduct.NykaaShirt{
		Shirt: concreteProduct.Shirt{
			Logo: "Nykaa",
			Size: 34,
		},
	}
}

func (n *Nykaa) MakeJeans() intf.IJeans {
	return &concreteProduct.NykaaJeans{
		Jeans: concreteProduct.Jeans{
			Logo: "Nykaa",
			Size: 34,
		},
	}
}
