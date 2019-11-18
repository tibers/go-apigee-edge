package apigee

import (
	"path"
)

// VirtualHostService is an interface for interfacing with the Apigee Edge Admin API
// dealing with target servers.
type VirtualHostsService interface {
	Get(string, string) (*VirtualHost, *Response, error)
	Create(VirtualHost, string) (*VirtualHost, *Response, error)
	Delete(string, string) (*Response, error)
	Update(VirtualHost, string) (*VirtualHost, *Response, error)
}

type VirtualHostsServiceOp struct {
	client *EdgeClient
}

var _ VirtualHostsService = &VirtualHostsServiceOp{}

// https://docs.apigee.com/api-platform/fundamentals/virtual-host-property-reference
type VirtualHost struct {
	Name                    string  `json:"name,omitempty"`
	HostAliases             array   `json:"hostAliases,omitempty"`
	Port                    int     `json:"port,omitempty"`
	// Interfaces              array   `json:"interfaces,omitempty"`
	RetryOptions            array   `json:"retryOptions,omitempty"`
	ListenOptions           array   `json:"listenOptions,omitempty"`
	BaseUrl                 string  `json:"baseUrl,omitempty"`
	SSLInfo                 hash    `json:"sSLInfo,omitempty"`
	// PropagateTLSInformation hash    `json:"propagateTLSInformation,omitempty"`
	Properties              array   `json:"properties,omitempty"`
}

func (s *VirtualHostsServiceOp) Get(name string, env string) (*VirtualHost, *Response, error) {

	path := path.Join("environments", env, "VirtualHosts", name)

	req, e := s.client.NewRequest("GET", path, nil, "")
	if e != nil {
		return nil, nil, e
	}
	returnedVirtualHost := VirtualHost{}
	resp, e := s.client.Do(req, &returnedVirtualHost)
	if e != nil {
		return nil, resp, e
	}
	return &returnedVirtualHost, resp, e

}

func (s *VirtualHostsServiceOp) Create(VirtualHost VirtualHost, env string) (*VirtualHost, *Response, error) {

	return postOrPutVirtualHost(VirtualHost, env, "POST", s)

}

func (s *VirtualHostsServiceOp) Update(VirtualHost VirtualHost, env string) (*VirtualHost, *Response, error) {

	return postOrPutVirtualHost(VirtualHost, env, "PUT", s)

}

func (s *VirtualHostsServiceOp) Delete(name string, env string) (*Response, error) {

	path := path.Join("environments", env, "VirtualHosts", name)

	req, e := s.client.NewRequest("DELETE", path, nil, "")
	if e != nil {
		return nil, e
	}

	resp, e := s.client.Do(req, nil)
	if e != nil {
		return resp, e
	}

	return resp, e

}

func postOrPutVirtualHost(VirtualHost VirtualHost, env string, opType string, s *VirtualHostsServiceOp) (*VirtualHost, *Response, error) {

	uripath := ""

	if opType == "PUT" {
		uripath = path.Join("environments", env, "VirtualHosts", VirtualHost.Name)
	} else {
		uripath = path.Join("environments", env, "VirtualHosts")
	}

	req, e := s.client.NewRequest(opType, uripath, VirtualHost, "")
	if e != nil {
		return nil, nil, e
	}

	returnedVirtualHost := VirtualHost{}

	resp, e := s.client.Do(req, &returnedVirtualHost)
	if e != nil {
		return nil, resp, e
	}

	return &returnedVirtualHost, resp, e

}
