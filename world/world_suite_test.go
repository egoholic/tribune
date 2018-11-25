package world_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestWorld(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "World Suite")
}
