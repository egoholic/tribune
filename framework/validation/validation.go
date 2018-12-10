package validation

type Validable interface {
	Validate() ValidationResult
}

type BinaryValidable interface {
	IsValid() bool
}

type ValidationResult interface {
	BinaryValidable
	Tree() ValidationNode
}

type ValidationNode struct {
	Domain   string
	Errors   []string
	Children []ValidationNode
}
