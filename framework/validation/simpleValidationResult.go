package validation

// Implements framework/validation.ValidationResult interface.
type SimpleValidationResult struct {
	root *ValidationNode
}

func Make(node *ValidationNode) SimpleValidationResult {
	return SimpleValidationResult{node}
}

func (ve *SimpleValidationResult) IsValid() bool {
	return ve.root.IsValid()
}

func (ve *SimpleValidationResult) Nodes() ValidationNode {
	return *ve.root
}
