package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"sort"
)

func Curl(token string) (string, error) {
	cmd := exec.Command("curl", "-X GET https://schematics.cloud.ibm.com/v1/locations", "-H Authorization:", token)
	res, err := cmd.CombinedOutput()
	return string(res), err
}

func Handler(w http.ResponseWriter, r *http.Request) {
	// Just return the list of env vars
	envs := os.Environ()
	sort.Strings(envs)

	fmt.Fprintf(w, "Environment variables:\n")
	for _, v := range envs {
		fmt.Fprintf(w, "%s\n", v)
	}
	resp, err := Curl(envs[0])
	fmt.Println("Response : ", resp)
	fmt.Println("Error : ", err)
}

func main() {
	http.HandleFunc("/loc", Handler)
	fmt.Printf("Listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
