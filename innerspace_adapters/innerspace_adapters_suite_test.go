package innerspace_adapters_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestInnerspaceAdapters(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "InnerspaceAdapters Suite")
}
