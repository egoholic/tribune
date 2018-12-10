package validation

type Validable interface {
	Validate() ValidationResult
}

type BinaryValidable interface {
	IsValid() bool
}

type ValidationResult interface {
	BinaryValidable
	Nodes() ValidationNode
}

type ValidationNode struct {
	domain   string
	errors   []string
	children []ValidationNode
}

func (vn *ValidationNode) AddError(errMsg string) {
	vn.errors = append(vn.errors, errMsg)
}

func N(domain string) ValidationNode {
	return ValidationNode{domain, []string{}, []ValidationNode{}}
}

func (vn *ValidationNode) AddChild(domain string) ValidationNode {
	child := N(domain)
	vn.children = append(vn.children, child)
	return child
}

func (vn *ValidationNode) IsValid() bool {
	return len(vn.errors) == 0
}
