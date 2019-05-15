package goomnisend

import (
	"fmt"
	"time"
)

const (
	members_path       = "/contacts"
	single_member_path = members_path + "/%s"

	member_activity_path = single_member_path + "/activity"
)

type ListOfMembers struct {
	Contacts []Contact `json:"contacts"`
	Paging   Paging    `json:"paging"`
}

type MemberRequest struct {
	Email              string                 `json:"email,omitempty"`
	CreatedAt          string                 `json:"createdAt,omitempty"`
	FirstName          string                 `json:"firstName,omitempty"`
	LastName           string                 `json:"lastName,omitempty"`
	Country            string                 `json:"country,omitempty"`
	CountryCode        string                 `json:"countryCode,omitempty"`
	State              string                 `json:"state,omitempty"`
	City               string                 `json:"city,omitempty"`
	Address            string                 `json:"address,omitempty"`
	PostalCode         string                 `json:"postalCode,omitempty"`
	Gender             string                 `json:"gender,omitempty"`
	Phone              string                 `json:"phone,omitempty"`
	Birthdate          string                 `json:"birthdate,omitempty"`
	Status             string                 `json:"status,omitempty"`
	StatusDate         time.Time              `json:"statusDate,omitempty"`
	CustomerProperties map[string]interface{} `json:"customerProperties,omitempty"`

	Lists []*ListResponse `json:"lists,omitempty"`
}

type Contact struct {
	MemberRequest

	ContactID string `json:"contactID"`

	Statuses *Statuses `json:"statuses"`
	api      *API
}

type Statuses struct {
	Status string `json:"status"`
	Date   string `json:"date"`
}

func (api API) GetMembers(params *InterestCategoriesQueryParams) (*ListOfMembers, error) {
	endpoint := members_path
	response := new(ListOfMembers)

	err := api.Request("GET", endpoint, params, nil, response)
	if err != nil {
		return nil, err
	}

	for _, m := range response.Contacts {
		m.api = &api
	}
	return response, nil
}

// func (api API) GetMember(id string, params *BasicQueryParams) (*Member, error) {
// 	if err := list.CanMakeRequest(); err != nil {
// 		return nil, err
// 	}

// 	endpoint := fmt.Sprintf(single_member_path, id)
// 	response := new(Member)
// 	response.api = list.api

// 	return response, list.api.Request("GET", endpoint, params, nil, response)
// }

func (api API) CreateMember(body *MemberRequest) (*Contact, error) {
	endpoint := fmt.Sprintf(members_path)
	response := new(Contact)
	response.api = &api

	return response, api.Request("POST", endpoint, nil, body, response)
}

// func (api API) UpdateMember(id string, body *MemberRequest) (*Member, error) {
// 	if err := list.CanMakeRequest(); err != nil {
// 		return nil, err
// 	}

// 	endpoint := fmt.Sprintf(single_member_path, id)
// 	response := new(Member)
// 	response.api = list.api

// 	return response, list.api.Request("PATCH", endpoint, nil, body, response)
// }

// func (api API) AddOrUpdateMember(id string, body *MemberRequest) (*Member, error) {
// 	if err := list.CanMakeRequest(); err != nil {
// 		return nil, err
// 	}

// 	endpoint := fmt.Sprintf(single_member_path, id)
// 	response := new(Member)
// 	response.api = list.api

// 	return response, list.api.Request("PUT", endpoint, nil, body, response)
// }

// func (api API) DeleteMember(id string) (bool, error) {
// 	if err := list.CanMakeRequest(); err != nil {
// 		return false, err
// 	}

// 	endpoint := fmt.Sprintf(single_member_path, id)
// 	return list.api.RequestOk("DELETE", endpoint)
// }
