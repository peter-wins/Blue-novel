package request

type Test struct {
	TestField string `validate:"required,min=1"`
}