package sqlrunner

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("connStringForDB", func() {
	const dbName = "diego_6"
	const want = "postgres://diego:diego_pw@localhost/diego_6"

	When("baseConnString has a trailing slash", func() {
		It("does not duplicate the slash", func() {
			base := "postgres://diego:diego_pw@localhost/"
			Expect(connStringForDB(base, dbName)).To(Equal(want))
		})
	})

	When("baseConnString has no trailing slash", func() {
		It("still joins with a single slash", func() {
			base := "postgres://diego:diego_pw@localhost"
			Expect(connStringForDB(base, dbName)).To(Equal(want))
		})
	})
})
