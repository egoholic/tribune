package blog_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBlog(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Blog Suite")
}
