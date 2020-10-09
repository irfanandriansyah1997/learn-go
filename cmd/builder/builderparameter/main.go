package main

import (
	builder "learning-go/pkg/builder/builderparameter"
)

func main() {
	builder.SendEmail(func(b *builder.EmailBuilder) {
		b.From("irfan@99.co").
			Body("Hello").
			Subject("Meeting").
			To("dev@99.co")
	})
}
