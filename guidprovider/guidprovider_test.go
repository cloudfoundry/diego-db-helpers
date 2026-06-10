package guidprovider_test

import (
	"code.cloudfoundry.org/diego-db-helpers/guidprovider"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

const uuidPattern = `^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`

var _ = Describe("GUIDProvider", func() {
	Describe("DefaultGuidProvider", func() {
		It("generates a valid v4 UUID", func() {
			guid, err := guidprovider.DefaultGuidProvider.NextGUID()
			Expect(err).NotTo(HaveOccurred())
			Expect(guid).To(MatchRegexp(uuidPattern))
		})

		It("generates unique GUIDs on successive calls", func() {
			guid1, err := guidprovider.DefaultGuidProvider.NextGUID()
			Expect(err).NotTo(HaveOccurred())

			guid2, err := guidprovider.DefaultGuidProvider.NextGUID()
			Expect(err).NotTo(HaveOccurred())

			Expect(guid1).NotTo(Equal(guid2))
		})
	})
})
