package tribune_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTribune(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Tribune Suite")
}
