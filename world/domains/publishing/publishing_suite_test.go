package publishing_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPublishing(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Publishing Suite")
}
