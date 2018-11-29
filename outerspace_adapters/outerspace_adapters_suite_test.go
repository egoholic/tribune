package outerspace_adapters_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestOuterspaceAdapters(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "OuterspaceAdapters Suite")
}
