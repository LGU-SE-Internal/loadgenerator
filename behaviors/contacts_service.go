package behaviors

import (
	"fmt"
	"math/rand"

	"github.com/Lincyaw/loadgenerator/service"
	"github.com/go-faker/faker/v4"
	log "github.com/sirupsen/logrus"
)

// ContactsBehaviorChain
func QueryContacts(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("failed to retrieve service client from context")
	}

	var contactsSvc service.ContactsService = cli
	TheAccountId := ctx.Get(UserId).(string)
	GetAllContacts, err := contactsSvc.GetContactByAccountId(TheAccountId)
	if err != nil {
		log.Errorf("Failed to query contacts for account_id=%s: %v", TheAccountId, err)
		return nil, err
	}
	if GetAllContacts.Status != 1 {
		log.Errorf("Unexpected response status when querying contacts for account_id=%s: status=%d", TheAccountId, GetAllContacts.Status)
		return nil, fmt.Errorf("unexpected response status: %d", GetAllContacts.Status)
	}

	randomIndex := rand.Intn(len(GetAllContacts.Data))
	//ctx.Set(AccountID, GetAllContacts.Data[randomIndex].AccountId)
	ctx.Set(ContactsID, GetAllContacts.Data[randomIndex].Id)
	ctx.Set(Name, GetAllContacts.Data[randomIndex].Name)
	ctx.Set(DocumentType, GetAllContacts.Data[randomIndex].DocumentType)
	ctx.Set(DocumentNumber, GetAllContacts.Data[randomIndex].DocumentNumber)
	ctx.Set(PhoneNumber, GetAllContacts.Data[randomIndex].PhoneNumber)

	return nil, nil
}

func CreateContacts(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("failed to retrieve service client from context")
	}

	CreateContactsInput := service.AdminContacts{
		Id:             faker.UUIDHyphenated(),
		AccountId:      ctx.Get(UserId).(string),
		Name:           faker.Name(),
		DocumentType:   rand.Intn(1),
		DocumentNumber: generateDocumentNumber(),
		PhoneNumber:    faker.PhoneNumber,
	}
	CreateContacts, err := cli.AddContact(&CreateContactsInput)
	if err != nil {
		log.Errorf("Failed to create contact for account_id=%s: %v", CreateContactsInput.AccountId, err)
		return nil, err
	}
	if CreateContacts.Status != 1 {
		log.Errorf("Unexpected response status when creating contact for account_id=%s: status=%d, response=%+v", CreateContactsInput.AccountId, CreateContacts.Status, CreateContacts)
		return nil, fmt.Errorf("unexpected response status: %d", CreateContacts.Status)
	}

	//ctx.Set(AccountID, CreateContacts.Data.AccountId)
	ctx.Set(ContactsID, CreateContacts.Data.Id)
	ctx.Set(Name, CreateContacts.Data.Name)
	ctx.Set(DocumentType, CreateContacts.Data.DocumentType)
	ctx.Set(DocumentNumber, CreateContacts.Data.DocumentNumber)
	ctx.Set(PhoneNumber, CreateContacts.Data.PhoneNumber)

	return nil, nil
}
