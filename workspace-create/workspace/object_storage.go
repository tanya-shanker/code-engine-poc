package workspace

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"time"
)

type COS_Result struct {
	OperationName string `json:"operationName"`
	Result        string `json:"result"`
	StatusCode    int    `json:"statusCode"`
	Messgae       string `json:"message"`
}

const (
	OBJECT_NAME = "ts-poc-wkp-list-"
)

func getURL(op string) string {

	BUCKET_ENDPOINT := os.Getenv("BUCKET_ENDPOINT")
	BUCKET_NAME := os.Getenv("BUCKET_NAME")

	bucketUrl := fmt.Sprintf("https://%s/%s", BUCKET_ENDPOINT, BUCKET_NAME)

	var url string
	switch op {
	case "PUT":
		dd, mm, yyyy := time.Now().Date()
		date := strconv.Itoa(dd) + "-" + mm.String() + "-" + strconv.Itoa(yyyy)
		random, _ := rand.Int(rand.Reader, big.NewInt(99999))
		obj := OBJECT_NAME + date + random.String()

		//+ strconv.Itoa(rand.Int())

		url = fmt.Sprintf("%s/%s", bucketUrl, obj)
		fmt.Println("url :", url)

	case "GET":
		url = fmt.Sprintf("%s?%s", bucketUrl, "list-type=2")
	}

	return url
}

// func GetObjects(token string) (string, error) {
// 	req, _ := http.NewRequest("GET", getURL("GET"), nil)

// 	req.Header.Add("Authorization", token)

// 	res, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		fmt.Println("Http req error :", err)
// 		return "", err
// 	}

// 	var readErr error
// 	var responseBody []byte
// 	if res.Body != nil {
// 		defer res.Body.Close()
// 		responseBody, readErr = ioutil.ReadAll(res.Body)
// 		if readErr != nil {
// 			err = fmt.Errorf(" ERRORMSG_READ_RESPONSE_BODY %s", readErr.Error())
// 			return "", err
// 		}
// 		fmt.Println(string(responseBody))
// 	}

// 	return "", nil
// }

func UploadObject(detailedResponse DetailedResponse, token, result string) error {

	fmt.Println("Uploading result to cos bucket")
	testResult := COS_Result{
		OperationName: "Workspace Create",
		Result:        result,
		StatusCode:    detailedResponse.StatusCode,
		Messgae:       string(detailedResponse.RawResult),
	}
	testResultByte, err := json.Marshal(testResult)

	if err != nil {
		fmt.Println("json marshal error :", err)
		return err
	}

	content := bytes.NewReader(testResultByte)
	req, _ := http.NewRequest("PUT", getURL("PUT"), content)

	req.Header.Add("Authorization", token)

	req.Header.Add("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Http req error :", err)
		return err
	}

	if res.Body != nil {
		var readErr error

		defer res.Body.Close()
		responseBody, readErr := ioutil.ReadAll(res.Body)
		if readErr != nil {
			err = fmt.Errorf(" ERRORMSG_READ_RESPONSE_BODY %s", readErr.Error())
			return err
		}
		fmt.Println(responseBody)
	}

	fmt.Println(res.StatusCode)
	fmt.Println("Uploaded result to cos bucket")
	return nil
}
