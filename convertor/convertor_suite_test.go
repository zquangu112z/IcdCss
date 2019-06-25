package convertor_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestIcdCcs(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "IcdCcs Suite")
}
