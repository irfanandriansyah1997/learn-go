package builderfacet

type PersonJobBuilder struct {
	PersonBuilder
}

func (pjb *PersonJobBuilder) At(companyName string) *PersonJobBuilder {
	pjb.person.CompanyName = companyName

	return pjb
}

func (pjb *PersonJobBuilder) AsA(position string) *PersonJobBuilder {
	pjb.person.Position = position

	return pjb
}

func (pjb *PersonJobBuilder) Earning(salary int) *PersonJobBuilder {
	pjb.person.AnnualIncome = salary

	return pjb
}
