package viewPoints_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestViewPoints(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ViewPoints Suite")
}
