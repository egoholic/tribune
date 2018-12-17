package content_test

import (
	"github.com/egoholic/tribune/framework/prop"
	. "github.com/egoholic/tribune/space/particle/content"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var content Content
var _ = Describe("tribune/space/particle/content/entity", func() {
	BeforeEach(func() {
		content = New()
	})

	Describe("Constructors", func() {
		Describe("Make()", func() {
			Context("When valid args", func() {
				It("returns Content entity", func() {
					c := Make("Title", "Body")
					Expect(c).To(BeAssignableToTypeOf(Content{}))
					Expect(c.Props()["Title"].Read()).To(Equal("Title"))
					Expect(c.Props()["Body"].Read()).To(Equal("Body"))
				})
			})
		})
	})

	Describe("Content entity", func() {
		Describe(".AssignProp()", func() {
			It("assigns prop", func() {
				Expect(content.Props()).To(HaveLen(0))
				ct := ContentTitle{}
				ct.Write("Title")
				content.AssignProp(&ct)
				Expect(content.Props()).To(HaveLen(1))
				Expect(content.Props()["Title"].Read()).To(Equal("Title"))
			})
		})

		Describe(".Props()", func() {
			It("returns a map of props", func() {
				pm := map[string]prop.Prop{}
				Expect(content.Props()).To(BeAssignableToTypeOf(pm))
				Expect(content.Props()).To(HaveLen(0))
				ct := ContentTitle{}
				content.AssignProp(&ct)
				Expect(content.Props()).To(HaveLen(1))
			})
		})

		Describe(".Validate()", func() {
			Context("when valid", func() {

			})

			Context("when invalid", func() {

			})
		})
	})
})
