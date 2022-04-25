package workspace

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

type DetailedResponse struct {

	// The HTTP status code associated with the response.
	StatusCode int

	// The HTTP headers contained in the response.
	Headers http.Header

	// Result - this field will contain the result of the operation (obtained from the response body).
	//
	// If the operation was successful and the response body contains a JSON response, it is un-marshalled
	// into an object of the appropriate type (defined by the particular operation), and the Result field will contain
	// this response object.  If there was an error while un-marshalling the JSON response body, then the RawResult field
	// will be set to the byte array containing the response body.
	//
	// Alternatively, if the generated SDK code passes in a result object which is an io.ReadCloser instance,
	// the JSON un-marshalling step is bypassed and the response body is simply returned in the Result field.
	// This scenario would occur in a situation where the SDK would like to provide a streaming model for large JSON
	// objects.
	//
	// If the operation was successful and the response body contains a non-JSON response,
	// the Result field will be an instance of io.ReadCloser that can be used by generated SDK code
	// (or the application) to read the response data.
	//
	// If the operation was unsuccessful and the response body contains a JSON error response,
	// this field will contain an instance of map[string]interface{} which is the result of un-marshalling the
	// response body as a "generic" JSON object.
	// If the JSON response for an unsuccessful operation could not be properly un-marshalled, then the
	// RawResult field will contain the raw response body.
	Result interface{}

	// This field will contain the raw response body as a byte array under these conditions:
	// 1) there was a problem un-marshalling a JSON response body -
	// either for a successful or unsuccessful operation.
	// 2) the operation was unsuccessful, and the response body contains a non-JSON response.
	RawResult []byte
}

func CreateWorkspace(url, token string) (DetailedResponse, error) {

	//var list Response
	fmt.Println("Creating workspace")

	payload := []byte(CreateWksp)
	content := bytes.NewReader(payload)

	req, _ := http.NewRequest("POST", url, content)

	req.Header.Add("Authorization", token)

	httpResponse, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Http req error :", err)
		return DetailedResponse{StatusCode: httpResponse.StatusCode}, err
	}

	// Start to populate the DetailedResponse.
	detailedResponse := &DetailedResponse{
		StatusCode: httpResponse.StatusCode,
		Headers:    httpResponse.Header,
	}

	// If the operation was unsuccessful, then set up the DetailedResponse
	// and error objects appropriately.
	if httpResponse.StatusCode < 200 || httpResponse.StatusCode >= 300 {

		var responseBody []byte

		// First, read the response body into a byte array.
		if httpResponse.Body != nil {
			var readErr error

			defer httpResponse.Body.Close()
			responseBody, readErr = ioutil.ReadAll(httpResponse.Body)
			if readErr != nil {
				err = fmt.Errorf(" ERRORMSG_READ_RESPONSE_BODY %s", readErr.Error())
				return DetailedResponse{StatusCode: httpResponse.StatusCode}, err
			}
		}

		// If the responseBody is empty, then just return a generic error based on the status code.
		if len(responseBody) == 0 {
			err = fmt.Errorf(http.StatusText(httpResponse.StatusCode))
			return DetailedResponse{StatusCode: httpResponse.StatusCode}, err
		}

		// For a non-JSON response or if we tripped while decoding the JSON response,
		// just return the response body byte array in the RawResult field along with
		// an error object that contains the generic error message for the status code.
		detailedResponse.RawResult = responseBody
		err = fmt.Errorf(http.StatusText(httpResponse.StatusCode))

	}
	return *detailedResponse, err
}
