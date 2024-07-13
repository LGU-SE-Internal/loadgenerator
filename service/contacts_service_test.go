package service

import (
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"log"
	"testing"
)

func TestSvc(t *testing.T) {
	cli, _ := GetBasicClient()
	var contactsSvc ContactsService = cli
	CreateContactsInput := AdminContacts{
		Id:        uuid.NewString(),
		AccountId: uuid.NewString(),
		Name:      faker.Name(),
	}
	CreateContacts, err := contactsSvc.AddContact(&CreateContactsInput)
	if err != nil {
		log.Fatalf("[MockedContactsID] CreateContacts error occurs: %v", err)
	}
	if CreateContacts.Data.Id == "" {
		log.Fatalf("Create AdminContacts Fail. Return Id = ''")
	}
}
func TestSvcImpl_GetAllContacts(t *testing.T) {
	cli, _ := GetBasicClient()
	contacts, err := cli.GetAllContacts()
	if err != nil {
		t.Error(err)
	}
	t.Log("[get all]: ", contacts)

	contactResp, err := cli.AddContact(&AdminContacts{
		Id:        uuid.NewString(),
		AccountId: uuid.NewString(),
		Name:      faker.Name(),
	})
	if err != nil {
		t.Error(err)
	}
	t.Log("[add contact]: ", contactResp)

	adminContactResp, err := cli.AddAdminContact(&AdminContacts{
		Id:        uuid.NewString(),
		AccountId: uuid.NewString(),
		Name:      faker.Name(),
	})
	if err != nil {
		t.Error(err)
	}
	t.Log("[add admin contact]: ", adminContactResp)

	modifyContactResp, err := cli.ModifyContact(&AdminContacts{
		Id:          contactResp.Data.Id,
		AccountId:   contactResp.Data.AccountId,
		Name:        faker.Name(),
		PhoneNumber: faker.E164PhoneNumber(),
	})
	if err != nil {
		t.Error(err)
	}
	t.Log("[modify contact]: ", modifyContactResp)

	contactResp1, err := cli.GetContactByContactId(modifyContactResp.Data.Id)
	if err != nil {
		t.Error(err)
	}
	t.Log("[get contact]: ", contactResp1)
	contactResp2, err := cli.GetContactByAccountId(modifyContactResp.Data.AccountId)
	if err != nil {
		t.Error(err)
	}
	t.Log("[get contact]: ", contactResp2)

	deleteContact, err := cli.DeleteContact(contactResp.Data.Id)
	if err != nil {
		t.Error(err)
	}
	t.Log("[delete contact]: ", deleteContact)

}
