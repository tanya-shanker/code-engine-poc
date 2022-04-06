package dummy

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestWorkspaceTest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "WorkspaceTest Suite")

}
