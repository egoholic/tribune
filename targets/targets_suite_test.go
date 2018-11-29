package targets_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTargets(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Targets Suite")
}
