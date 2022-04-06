package dummy_test

import (
	"net/http"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tanya-shanker/code-engine-poc/job"
)

var _ = Describe("Workspace", func() {
	var iam_token string
	Context("with default pagination params", func() {
		BeforeEach(func() {
			iam_token = os.Getenv("POC_TOKEN")
		})
		It("should return list of workspaces with default page size :sanity:", func() {
			// Get workspace list
			workspaceResponseList, _ := job.GetWorkspaceList("https://schematics.cloud.ibm.com/v1/workspaces", iam_token)
			//Expect(err).NotTo(HaveOccurred())
			Expect(workspaceResponseList.StatusCode).NotTo(Equal(http.StatusBadGateway))
			Expect(workspaceResponseList.StatusCode).NotTo(Equal(http.StatusServiceUnavailable))
		})
	})
})
