package concreteProduct

type Jeans struct {
	Logo string
	Size int
}

func (s *Jeans) SetLogo(Logo string) {
	s.Logo = Logo
}

func (s *Jeans) GetLogo() string {
	return s.Logo
}

func (s *Jeans) SetSize(Size int) {
	s.Size = Size
}

func (s *Jeans) GetSize() int {
	return s.Size
}
