package job

// func Curl(token string) (string, error) {
// 	// options := []string{
// 	// 	"-X POST ",
// 	// 	"https://iam.cloud.ibm.com/identity/token",
// 	// 	"-H Content-Type: application/x-www-form-urlencoded",
// 	// 	"-d grant_type=urn:ibm:params:oauth:grant-type:apikey&apikey=$IBMCLOUD_KEY",
// 	// 	"-u bx:bx",
// 	// }

// 	//curl -X POST "https://iam.cloud.ibm.com/identity/token" -H "Content-Type: application/x-www-form-urlencoded" -d "grant_type=urn:ibm:params:oauth:grant-type:apikey&apikey=$IBMCLOUD_KEY" -u bx:bx

// 	options := []string{
// 		//"--location",
// 		//"--request",
// 		"-X GET https://schematics.cloud.ibm.com/v1/workspaces",
// 		"--header 'Authorization:" + token + "'",
// 		// " --header 'Content-Type: application/json'",
// 		// "--header 'X-Feature-Region-Visibility: true'",
// 	}

// 	//	curl -X GET https://schematics.cloud.ibm.com/v1/workspaces -H "Authorization: <iam_token>"

// 	//curl --location --request GET  https://schematics.cloud.ibm.com/v2/locations --header "Authorization: <access_token> " --header "Content-Type: application/json" --header "X-Feature-Region-Visibility: true"

// 	cmd := exec.Command("curl", options...)
// 	//cmd := exec.Command("curl", "-X GET 'https://schematics.cloud.ibm.com/v1/locations' -H Authorization:", token)
// 	res, err := cmd.CombinedOutput()
// 	return string(res), err
// }

// func Handler(w http.ResponseWriter, r *http.Request) {
// 	token := os.Getenv("POC_TOKEN")
// 	list := GetWorkspaceList("https://schematics.cloud.ibm.com/v1/workspaces", token)

// 	resp, _ := json.Marshal(list)
// 	w.Write(resp)
// }

// func main() {
// 	//http.HandleFunc("/workspaces", Handler)
// 	fmt.Printf("Listening on port 8081")
// 	http.ListenAndServe(":8081", nil)
// }
