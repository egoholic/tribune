package layouts_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestLayouts(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Layouts Suite")
}
