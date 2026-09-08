package sqlrunner

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestSqlrunner(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SqlRunner Suite")
}
