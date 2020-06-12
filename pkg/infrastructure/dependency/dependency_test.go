package dependency

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestShouldGetVersionByName(t *testing.T) {
	g := NewGomegaWithT(t)

	result := BySlashFileName(`pax-exam-1.1.0.jar`)
	g.Expect(result.Version).To(Equal(`1.1.0`))
}

func TestShouldGetArtifactId(t *testing.T) {
	g := NewGomegaWithT(t)

	result := BySlashFileName(`pax-exam-1.1.0.jar`)
	g.Expect(result.ArtifactId).To(Equal(`pax-exam`))
}

func Test_ShouldSupportCustomPackageNameLength(t *testing.T) {
	g := NewGomegaWithT(t)

	result := ByFileName(`com.csii.pe.http_6.0.0.20121109.jar`, 0)
	g.Expect(result.Version).To(Equal(`6.0.0.20121109`))
}

func Test_ShouldGetTwoLengthGroupIdWhen2(t *testing.T) {
	g := NewGomegaWithT(t)

	result := ByFileName(`com.csii.pe.http_6.0.0.20121109.jar`, 2)
	g.Expect(result.GroupId).To(Equal(`com.csii`))
	g.Expect(result.ArtifactId).To(Equal(`pe.http`))
}

func Test_ShouldGetOneLengthGroupIdWhen2(t *testing.T) {
	g := NewGomegaWithT(t)

	result := ByFileName(`com.csii.pe.http_6.0.0.20121109.jar`, 1)
	g.Expect(result.GroupId).To(Equal(`com`))
	g.Expect(result.ArtifactId).To(Equal(`csii.pe.http`))
}
