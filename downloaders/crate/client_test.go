package crate

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type MockServer struct {
	ResponseCode int
	ResponseData []byte
}

func NewMockServer(responseCode int, responseData []byte) *MockServer {
	return &MockServer{
		ResponseCode: responseCode,
		ResponseData: responseData,
	}
}

func (s *MockServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(s.ResponseCode)
	fmt.Fprint(w, string(s.ResponseData))
}

func TestGet(t *testing.T) {
	Convey("Given that the file exists", t, func() {
		Convey("When Get is invoked", func() {
			Convey("It gets the file and writes it to the writer", func() {
				mockServer := NewMockServer(200, []byte("SOME FILE"))
				server := httptest.NewServer(mockServer)
				defer server.Close()

				crateUrl := fmt.Sprintf("%s/_blobs/abc123", server.URL)
				downloadUrl, _ := url.Parse(crateUrl)
				writer := bytes.NewBuffer([]byte{})

				client, _ := NewCrateClient(downloadUrl)
				client.Get(writer)
				So(writer.String(), ShouldEqual, "SOME FILE")
			})
		})
	})
}
