package sample

import (
	"github.com/CenturyLinkLabs/pmxadapter"
)

type SampleAdapter struct {
}

// Implementation of the PanamaxAdapter GetServices interface
func (a *SampleAdapter) GetServices() ([]*pmxadapter.Service, *pmxadapter.Error) {
	var apiErr *pmxadapter.Error
	services := make([]*pmxadapter.Service, 1)
	service := new(pmxadapter.Service)
	services[0] = service

	return services, apiErr
}

// Implementation of the PanamaxAdapter GetService interface
func (a *SampleAdapter) GetService(id string) (*pmxadapter.Service, *pmxadapter.Error) {
	var apiErr *pmxadapter.Error
	service := new(pmxadapter.Service)

	return service, apiErr
}

// Implementation of the PanamaxAdapter CreateServices interface
func (a *SampleAdapter) CreateServices(services []*pmxadapter.Service) ([]*pmxadapter.Service, *pmxadapter.Error) {
	var apiErr *pmxadapter.Error
	list := make([]*pmxadapter.Service, 1)
	service := new(pmxadapter.Service)
	list[0] = service

	return list, apiErr
}

// Implementation of the PanamaxAdapter UpdateService interface
func (a *SampleAdapter) UpdateService(s *pmxadapter.Service) *pmxadapter.Error {
	return nil
}

// Implementation of the PanamaxAdapter DestroyService interface
func (a *SampleAdapter) DestroyService(id string) *pmxadapter.Error {
	return nil
}

// GetMetadata returns metadata for the adapter.
func (a *SampleAdapter) GetMetadata() pmxadapter.Metadata {
	return pmxadapter.Metadata{Type: "Sample", Version: "0.1"}
}
