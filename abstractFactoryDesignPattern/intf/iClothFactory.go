package intf

type IClothFactory interface {
	MakeShirt() IShirt
	MakeJeans() IJeans
}
