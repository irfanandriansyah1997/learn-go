package builderfacet

type PersonAddressBuilder struct {
	PersonBuilder
}

func (pab *PersonAddressBuilder) At(place string) *PersonAddressBuilder {
	pab.person.StreetAddress = place

	return pab
}

func (pab *PersonAddressBuilder) In(city string) *PersonAddressBuilder {
	pab.person.City = city

	return pab
}

func (pab *PersonAddressBuilder) PostalCode(postalCode string) *PersonAddressBuilder {
	pab.person.Postcode = postalCode

	return pab
}
