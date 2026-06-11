package guidprovider_test

import (
	"code.cloudfoundry.org/diego-db-helpers/guidprovider"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// nu7hatch/gouuid generates UUID v4 strings but does not strictly enforce RFC
// 4122 variant bits on all platforms (e.g. linux/arm64), so we validate only
// the UUID shape (8-4-4-4-12 lowercase hex) rather than the version/variant
// nibbles.
const uuidPattern = `^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`

var _ = Describe("GUIDProvider", func() {
	Describe("DefaultGuidProvider", func() {
		It("generates a UUID-shaped string", func() {
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
