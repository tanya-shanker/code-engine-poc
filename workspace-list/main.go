package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/tanya-shanker/code-engine-poc/workspace-list/workspace"
)

// Entry point of application
func main() {
	// Serve and Listen at port 8087
	fmt.Println("Service started at port 8087")
	iam_token := os.Getenv("POC_TOKEN")
	workspaceResponseList, err := workspace.GetWorkspaceList("https://schematics.cloud.ibm.com/v1/workspaces", iam_token)
	if err != nil {
		fmt.Println("Error occurred : ", err)
	}
	if workspaceResponseList.StatusCode == http.StatusBadGateway || workspaceResponseList.StatusCode == http.StatusServiceUnavailable {
		fmt.Println("TEST FAILED")
		workspace.UploadObject(workspaceResponseList, iam_token, "FAILED")
	} else {
		fmt.Println("status code :", workspaceResponseList.StatusCode)
		fmt.Println("TEST PASSED")
		workspace.UploadObject(workspaceResponseList, iam_token, "PASSED")
	}
	http.ListenAndServe(":8087", nil)

}
