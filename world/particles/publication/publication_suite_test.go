package publication_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPublication(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Publication Suite")
}
