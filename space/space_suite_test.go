package space_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSpace(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SpaceSuite")
}
