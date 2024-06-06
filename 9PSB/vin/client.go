package vin

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type RealHTTPClient struct{}

func (c *RealHTTPClient) Do(req *http.Request) (*http.Response, error) {
	client := &http.Client{}
	return client.Do(req)
}

type MockHTTPClient struct{}

// }
// func (m *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
// 	response := CoralPayCConnectAuthenticationResponse{
// 		Token: "mockToken",
// 	}

// 	responseBody, _ := json.Marshal(response)
// 	return &http.Response{
// 		StatusCode: http.StatusOK,
// 		Body:       ioutil.NopCloser(bytes.NewBuffer(responseBody)),
// 	}, n

func (m *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	response := NinePSBVirtualAuthenticationResponse{
		Access_token:        "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJuYW1laWQiOiIxOTVBMTY0NkQzMjc0QTZEQkVCMENEQzJDMTBDREU2QSIsImh0dHA6Ly9zY2hlbWFzLnhtbHNvYXAub3JnL3dzLzIwMDUvMDUvaWRlbnRpdHkvY2xhaW1zL2hhc2giOiJjNjc4Zjc0MWM2NjU2YmM4NmUwMDk0OGE0ZGI1ZTRkYTFiZWUxM2RiYjNjMjEzZjkwZTU4YjE5NWI0YTAwZDE3NGU2NjA1NTc0OGRlYTcyMjg0YTEyMTAzYjRkZjVkZmRmNjQwMDdhNjhiYzYwYjBhMWQ5M2U4NTg4YzQ5NWIyYyIsInJvbGUiOiJ0cTJpd3piUWUxM1NDRi9rVDBNQjV5cnU1NVpMOEZacnFUYkppM2JsK2EyNnFwT283MVpwNSt1VjAxZmxBVUpxUmpTVVYxSDJTb1BVazNld3h1UXY3YlFCbkg2bjRCbGloUittNCtIR2ZvejBzUjFicVdUZ1lqdzQ4WjJpSHRnTkNhUWxQMzZybGhKUHdETVJSYnlPejA2RzV1TzdOLzhyYzJxWnp5bWJJYTkrSTd6K3JRclJpY1hpQjJaSmJDL1FhcVlCZGtzRXlPMnN1cG5qc1hjNVcyb3k3bmdNcUFjeXNaQjd6c3NFdnFsK0tvWStWNElhekx6MDdyN2E1Y25MelhvYjZaa2hmUkVJTG1sTnlPdkJaV2xiVUpta0lLQU5GSTNZNU9RdU80Z3dSRVg3Z1h2K2lTTEk1S0ViOWxKWVpmRGN6aTRMem9qUnZJdVpTNVNFcUJ4QXlNYlV1dUxhSXZlcDEyNEpSQnl1bzlGNUJWeW5CVWIwZ3Z6QVMrUVUiLCJuYmYiOjE2ODY0ODM4OTksImV4cCI6MTY4NjQ5MTA5OSwiaWF0IjoxNjg2NDgzODk5LCJpc3MiOiI5UFNCX1NlcnZpY2VfSW50ZXJuYWwiLCJhdWQiOiI5UFNCX1NlcnZpY2VfSW50ZXJuYWwifQ." +
							"0mdBaNIoiM_RTeCMUv3oukkpNcMGotyGa0VJzxV8EVBADsH4uvV5c11LACl2Vr6Tm83uxc3xJHxsAAkb4JdsoQ",
		Expires_in:      7200,
		Code: "00",
		Message:  "Success",
	}

	responseBody, _ := json.Marshal(response)
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(bytes.NewBuffer(responseBody)),
	}, nil
}
