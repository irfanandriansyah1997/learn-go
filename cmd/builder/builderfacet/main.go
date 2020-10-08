package main

import (
	"fmt"
	builder "learning-go/pkg/builder/builderfacet"
)

func main() {
	person := builder.NewPersonBuilder().
		Lives().
		At("Jln. Cimareme No 40").
		In("Padalarang").
		PostalCode("40553").
		Works().
		At("99.co").
		AsA("Software Engineer").
		Earning(1000).
		Build()

	fmt.Println(*person)
}
