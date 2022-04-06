package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
)

var (
	BUF_LEN = 1024
)

func WriteToFile(f *os.File, data []byte) error {
	bytesWritten, err := f.Write(data)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return err
	}
	fmt.Println("number of bytes successfully written", bytesWritten)
	return nil
}

// Write Logs to file
func writelogs(data []byte) {
	logs, err := os.OpenFile("TestRun.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Write logs to file
	err = WriteToFile(logs, data)
	if err != nil {
		fmt.Println(err)
		return
	}
}

// Go Routine to async retrun logs as they appear
func writeCmdOutput(res http.ResponseWriter, pipeReader *io.PipeReader) {
	buffer := make([]byte, BUF_LEN)
	for {
		n, err := pipeReader.Read(buffer)
		if err != nil {
			pipeReader.Close()
			break
		}

		data := buffer[0:n]
		// Write logs to file for further parsing
		writelogs(data)
		// Sent the reponse to UI/Out
		res.Write(data)
		if f, ok := res.(http.Flusher); ok {
			f.Flush()
		}
		//reset buffer to empty
		for i := 0; i < n; i++ {
			buffer[i] = 0
		}
	}
}
func handler(w http.ResponseWriter, r *http.Request) {

	// Intilize the go test command and directory
	//	cmd := exec.Command("go", "test", "-gingko.v -test.timeout=14400s")
	cmd := exec.Command("ginkgo")
	cmd.Dir = "./workspace/"
	pipeReader, pipeWriter := io.Pipe()
	cmd.Stdout = pipeWriter
	cmd.Stderr = pipeWriter

	// Start Go routine to return response Async
	go writeCmdOutput(w, pipeReader)
	// Run the command
	cmd.Run()
	pipeWriter.Close()

}

// Entry point of application
func main() {
	http.HandleFunc("/test", handler)
	//http.HandleFunc("/favicon.ico", ignoreFavIconRequest)
	// Serve and Listen at port 8086
	http.ListenAndServe(":8086", nil)
}
