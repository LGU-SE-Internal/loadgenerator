package service

import (
	"encoding/json"
	"fmt"
	"io"
)

type ContactsService interface {
	GetAllContacts() (*AdminGetContactsResp, error)
	AddContact(contacts *AdminContacts) (*AdminContactResponse, error)
	AddAdminContact(contacts *AdminContacts) (*AdminContactResponse, error)
	ModifyContact(contacts *AdminContacts) (*AdminContactResponse, error)
	DeleteContact(contactsId string) (*DeleteContactsResp, error)
	GetContactByContactId(contactsId string) (*AdminContactResponse, error)
	GetContactByAccountId(accountId string) (*GetContactsByAccountIdResp, error)
}

type GetContactsByAccountIdResp struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   []struct {
		Id             string `json:"id"`
		AccountId      string `json:"accountId"`
		Name           string `json:"name"`
		DocumentType   int    `json:"documentType"`
		DocumentNumber string `json:"documentNumber"`
		PhoneNumber    string `json:"phoneNumber"`
	} `json:"data"`
}
type DeleteContactsResp struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   string `json:"data"`
}

func (s *SvcImpl) GetAllContacts() (*AdminGetContactsResp, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+"/api/v1/contactservice/contacts", nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result AdminGetContactsResp
	err = json.Unmarshal(body, &result)
	return &result, err
}

func (s *SvcImpl) AddContact(contacts *AdminContacts) (*AdminContactResponse, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/contactservice/contacts", contacts)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result AdminContactResponse
	err = json.Unmarshal(body, &result)
	return &result, err
}

func (s *SvcImpl) AddAdminContact(contacts *AdminContacts) (*AdminContactResponse, error) {
	resp, err := s.cli.SendRequest("POST", s.BaseUrl+"/api/v1/contactservice/contacts/admin", contacts)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result AdminContactResponse
	err = json.Unmarshal(body, &result)
	return &result, err
}
func (s *SvcImpl) ModifyContact(contacts *AdminContacts) (*AdminContactResponse, error) {
	resp, err := s.cli.SendRequest("PUT", s.BaseUrl+"/api/v1/contactservice/contacts", contacts)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result AdminContactResponse
	err = json.Unmarshal(body, &result)
	return &result, err
}
func (s *SvcImpl) DeleteContact(contactsId string) (*DeleteContactsResp, error) {
	resp, err := s.cli.SendRequest("DELETE", s.BaseUrl+fmt.Sprintf("/api/v1/contactservice/contacts/%s", contactsId), nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result DeleteContactsResp
	err = json.Unmarshal(body, &result)
	return &result, err
}
func (s *SvcImpl) GetContactByContactId(contactsId string) (*AdminContactResponse, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+fmt.Sprintf("/api/v1/contactservice/contacts/%s", contactsId), nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result AdminContactResponse
	err = json.Unmarshal(body, &result)
	return &result, err
}
func (s *SvcImpl) GetContactByAccountId(accountId string) (*GetContactsByAccountIdResp, error) {
	resp, err := s.cli.SendRequest("GET", s.BaseUrl+fmt.Sprintf("/api/v1/contactservice/contacts/account/%s", accountId), nil)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result GetContactsByAccountIdResp
	err = json.Unmarshal(body, &result)
	return &result, err
}
