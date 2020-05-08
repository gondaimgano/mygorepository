package model

type product struct {
	name string,
	dob string,
}

func (uc product) getName() string{

	return uc.name
}

func (uc product) getDob() string{
	return uc.dob
}

func newProduct(name string,dob string) *product{
	return &product(
		name: name,
		dob: dob
	)
}