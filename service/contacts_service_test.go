package service

import (
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"log"
	"testing"
)

func TestSvc_FullIntegration(t *testing.T) {
	cli, _ := GetAdminClient()
	var contactsSvc ContactsService = cli

	// CreateContact
	CreateContactsInput := AdminContacts{
		Id:        uuid.NewString(),
		AccountId: uuid.NewString(),
		Name:      faker.Name(),
	}
	CreateContacts, err := contactsSvc.AddContact(&CreateContactsInput)
	if err != nil {
		log.Fatalf("[MockedContactsID] CreateContacts error occurs: %v", err)
	}
	if CreateContacts.Status != 1 {
		t.Errorf("CreateContacts.Status != 1")
	}
	if CreateContacts.Data.Id == "" {
		t.Errorf("Create AdminContacts Fail. Return Id = ''")
	}
	isMatch := false
	if /*CreateContacts.Data.Id == CreateContactsInput.Id &&*/
	CreateContacts.Data.Name == CreateContactsInput.Name &&
		CreateContacts.Data.AccountId == CreateContactsInput.AccountId &&
		CreateContacts.Data.PhoneNumber == CreateContactsInput.PhoneNumber &&
		CreateContacts.Data.DocumentNumber == CreateContactsInput.DocumentNumber &&
		CreateContacts.Data.DocumentType == CreateContactsInput.DocumentType {
		isMatch = true
	}
	if !isMatch {
		t.Errorf("Create AdminContacts Fail. expect: %v, get %v", CreateContactsInput, CreateContacts.Data)
	}
	existedContacts := CreateContacts.Data

	// getAllContacts
	contacts, err := contactsSvc.GetAllContacts()
	if err != nil {
		t.Error(err)
	}
	if contacts.Status != 1 {
		t.Errorf("contacts.Status != 1")
	}
	isMatch1 := false
	for _, contact := range contacts.Data {
		if contact.Id == existedContacts.Id &&
			contact.Name == existedContacts.Name &&
			contact.AccountId == existedContacts.AccountId &&
			contact.PhoneNumber == existedContacts.PhoneNumber &&
			contact.DocumentNumber == existedContacts.DocumentNumber &&
			contact.DocumentType == existedContacts.DocumentType {
			isMatch1 = true
		}
	}
	if !isMatch1 {
		t.Errorf("Except to find: %v, but failed.", existedContacts)
	}
	t.Log("[get all]: ", contacts)

	// AddAdminContact
	adminContactReq := AdminContacts{
		Id:        uuid.NewString(),
		AccountId: uuid.NewString(),
		Name:      faker.Name(),
	}
	adminContactResp, err := cli.AddAdminContact(&adminContactReq)
	if err != nil {
		t.Error(err)
	}
	if adminContactResp.Status != 1 {
		t.Errorf("adminContactResp.Status != 1")
	}
	isMatch2 := false
	if /*adminContactResp.Data.Id == adminContactReq.Id &&*/
	adminContactResp.Data.Name == adminContactReq.Name &&
		adminContactResp.Data.AccountId == adminContactReq.AccountId &&
		adminContactResp.Data.PhoneNumber == adminContactReq.PhoneNumber &&
		adminContactResp.Data.DocumentNumber == adminContactReq.DocumentNumber &&
		adminContactResp.Data.DocumentType == adminContactReq.DocumentType {
		isMatch2 = true
	}
	if !isMatch2 {
		t.Errorf("Do not macth. expect: %v, get %v", adminContactReq, adminContactResp)
	}
	t.Log("[add admin contact]: ", adminContactResp)

	// Modify Contacts
	modifyContactResp, err := cli.ModifyContact(&AdminContacts{
		Id:          existedContacts.Id,
		AccountId:   existedContacts.AccountId,
		Name:        faker.Name(),
		PhoneNumber: faker.E164PhoneNumber(),
	})
	if err != nil {
		t.Error(err)
	}
	t.Log("[modify contact]: ", modifyContactResp)

	contactResp1, err := cli.GetContactByContactId(existedContacts.Id)
	if err != nil {
		t.Error(err)
	}
	if contactResp1.Status != 1 {
		t.Errorf("contactResp1.Status != 1")
	}
	found := false
	if /*contactResp1.Data.Id == existedContacts.Id &&*/
	contactResp1.Data.AccountId == existedContacts.AccountId &&
		/*contactResp1.Data.PhoneNumber == existedContacts.PhoneNumber &&*/
		contactResp1.Data.DocumentNumber == existedContacts.DocumentNumber &&
		contactResp1.Data.DocumentType == existedContacts.DocumentType {
		found = true
	}
	if !found {
		t.Errorf("Except to find: %v, but failed.", existedContacts)
	}
	t.Log("[get contact]: ", contactResp1)

	contactResp2, err := cli.GetContactByAccountId(existedContacts.AccountId)
	if err != nil {
		t.Error(err)
	}
	if contactResp2.Status != 1 {
		t.Errorf("contactResp2.Status != 1")
	}
	found1 := false
	for _, contact := range contactResp2.Data {
		if contact.Id == existedContacts.Id &&
			/*contact.Name == existedContacts.Name &&*/
			contact.AccountId == existedContacts.AccountId &&
			/*contact.PhoneNumber == existedContacts.PhoneNumber &&*/
			contact.DocumentType == existedContacts.DocumentType &&
			contact.DocumentNumber == existedContacts.DocumentNumber {
			found1 = true
		}
	}
	if !found1 {
		t.Errorf("Except to find: %v, but failed.", existedContacts)
	}
	t.Log("[get contact]: ", contactResp2)

	deleteContact, err := cli.DeleteContact(existedContacts.Id)
	if err != nil {
		t.Error(err)
	}
	if deleteContact.Status != 1 {
		t.Errorf("deleteContact.Status != 1")
	}
	t.Log("[delete contact]: ", deleteContact)

}
