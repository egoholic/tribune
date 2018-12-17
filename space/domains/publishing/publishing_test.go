package publishing_test

import (
	. "github.com/egoholic/tribune/space/domains/publishing"
	cnt "github.com/egoholic/tribune/space/particle/content"
	pub "github.com/egoholic/tribune/space/particle/publication"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var content cnt.Content

var _ = Describe("tribune/space/domains/publishing", func() {
	BeforeEach(func() {
		content = cnt.New()
	})

	Describe("PublishNow()", func() {
		It("Creates publication and returns it", func() {
			Expect(PublishNow(&content)).To(BeAssignableToTypeOf(pub.Publication{}))
		})
	})
})
