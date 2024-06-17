package service

import (
	"testing"
)

func TestAdminBasicInfoService_FullIntegration(t *testing.T) {
	adminClient, _ := GetAdminClient()

	allContacts, err := adminClient.AdminGetAllContacts()
	if err != nil {
		t.Errorf("Failed to get all contacts: %v", err)
	}

	// Create a client for admin user and perform user creation
	adminRegisterResp, err := adminClient.AdminAddContact(&AdminContacts{Name: "Admin User Contact"})
	if err != nil {
		t.Errorf("Failed to create contact for admin user: %v", err)
	}
	if adminRegisterResp.Status != 1 {
		t.Errorf("Expected status 1, got %d", adminRegisterResp.Status)
	}

	var id string
	if len(allContacts.Data) > 0 {
		id = allContacts.Data[0].Id
	} else {
		t.Errorf("allContacts.Data is empty")
	}

	// Admin user modifies contact
	modifyResp, err := adminClient.AdminModifyContact(&AdminContacts{ID: id, Name: "Modified Basic Contact"})
	if err != nil {
		t.Errorf("Failed to modify contact for basic user: %v", err)
	}
	if modifyResp.Status != 1 {
		t.Errorf("Expected status 1, got %d", modifyResp.Status)
	}

	var contactsId string
	if len(allContacts.Data) > 0 {
		contactsId = allContacts.Data[len(allContacts.Data)-1].Id
	} else {
		t.Errorf("allContacts.Data is empty")
	}

	// Admin user deletes contact
	deleteResp, err := adminClient.AdminDeleteContact(contactsId)
	if err != nil {
		t.Errorf("Failed to delete contact for admin user: %v", err)
	}
	if deleteResp.Status != 1 {
		t.Errorf("Expected status 1, got %d", deleteResp.Status)
	}
}
