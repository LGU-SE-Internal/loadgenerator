package behaviors

import (
	"fmt"
	"github.com/Lincyaw/loadgenerator/service"
	"github.com/go-faker/faker/v4"
	log "github.com/sirupsen/logrus"
	"math/rand"
)

// ContactsBehaviorChain
func QueryContacts(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	var contactsSvc service.ContactsService = cli
	TheAccountId := ctx.Get(UserId).(string)
	GetAllContacts, err := contactsSvc.GetContactByAccountId(TheAccountId)
	if err != nil {
		log.Errorf("[Mock AccountID]GetAllContacts fail. The error occurs: %v", err)
		return nil, err
	}
	if GetAllContacts.Status != 1 {
		log.Errorf("[Mock AccountID]GetAllContacts.Status != 1")
		return nil, err
	}

	randomIndex := rand.Intn(len(GetAllContacts.Data))
	ctx.Set(AccountID, GetAllContacts.Data[randomIndex].AccountId)
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
		return nil, fmt.Errorf("service client not found in context")
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
		log.Errorf("[Mock AccountID] CreateContacts error occurs: %v", err)
		return nil, err
	}
	if CreateContacts.Status != 1 {
		log.Errorf("[Mock AccountID] CreateContacts.Status != 1, resp: %+v", CreateContacts)
		return nil, err
	}

	ctx.Set(AccountID, CreateContacts.Data.AccountId)
	ctx.Set(ContactsID, CreateContacts.Data.Id)
	ctx.Set(Name, CreateContacts.Data.Name)
	ctx.Set(DocumentType, CreateContacts.Data.DocumentType)
	ctx.Set(DocumentNumber, CreateContacts.Data.DocumentNumber)
	ctx.Set(PhoneNumber, CreateContacts.Data.PhoneNumber)

	return nil, nil
}
