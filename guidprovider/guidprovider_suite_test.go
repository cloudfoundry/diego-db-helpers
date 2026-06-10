package guidprovider_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGuidprovider(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GuidProvider Suite")
}
