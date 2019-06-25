package convertor_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/zquangu112z/IcdCcs/convertor"
	. "github.com/zquangu112z/IcdCcs/convertor"
)

var _ = Describe("Convertor", func() {
	var (
		C028   IcdInfo
		Z44001 IcdInfo
	)

	BeforeEach(func() {
		C028 = convertor.GetIcdInfo("C028", convertor.CodeSystemICD10Diag)
		Z44001 = convertor.GetIcdInfoBestEffort("Z44001")
	})

	Describe("Get information from an ICD code", func() {
		Context("provided code system", func() {
			It("return its CCS category", func() {
				Expect(C028.CcsCategory).To(Equal("11"))
			})
		})
		Context("without provided code system", func() {
			It("return its CCS description", func() {
				Expect(Z44001.CcsCategoryDescription).To(Equal("Rehabilitation care; fitting of prostheses; and adjustment of devices"))
			})
		})

	})
})
