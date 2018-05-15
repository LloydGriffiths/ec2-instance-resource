package in

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// Request represnts an in request.
type Request struct {
	Version *Version `json:"version"`
}

// Version represents the resource version.
type Version struct {
	Date      string `json:"date"`
	Instances string `json:"instances"`
}

// Response reprensts an in reponse.
type Response struct {
	Version *Version `json:"version"`
}

// Run gets the resource version.
func Run(dest string, request *Request) (*Response, error) {
	if err := os.MkdirAll(dest, 0755); err != nil {
		return nil, err
	}
	if err := ioutil.WriteFile(filepath.Join(dest, "ec2-instances"), []byte(request.Version.Instances), 0644); err != nil {
		return nil, err
	}
	return &Response{&Version{Date: request.Version.Date, Instances: request.Version.Instances}}, nil
}
