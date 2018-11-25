package html_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestHtml(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Html Suite")
}
