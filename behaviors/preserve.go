package behaviors

import (
	"fmt"
	"github.com/Lincyaw/loadgenerator/service"
	"github.com/go-faker/faker/v4"
	"log"
	"math/rand"
	"time"
)

const (
	AccountID  = "accountId"
	ContactsID = "contactsId"
	TripID     = "tripId"
	SeatType   = "seatType"
	//LoginToken = "loginToken"
	Date            = "date"
	From            = "from"
	To              = "to"
	Assurance       = "assurance"
	FoodType        = "foodType"
	StationName     = "stationName"
	StoreName       = "storeName"
	FoodName        = "foodName"
	FoodPrice       = "foodPrice"
	HandleDate      = "handleDate"
	ConsigneeName   = "consigneeName"
	ConsigneePhone  = "consigneePhone"
	ConsigneeWeight = "consigneeWeight"
	IsWithin        = "isWithin"
)

var PreserveChain *Chain

func init() {
	// init
	PreserveChain = NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		fmt.Printf("PreserveBehaviors(Chain) Statrs. Starts time: %v", time.Now().String())
		return nil, nil
	}))

	// NewFuncNode
	LoginAdminNode := NewFuncNode(LoginAdmin)
	QueryContactsNode := NewFuncNode(QueryContacts)
	CreateContactsNode := NewFuncNode(CreateContacts)

	// NewChain
	LoginAdminChain := NewChain(LoginAdminNode)
	QueryContactsChain := NewChain(QueryContactsNode)
	CreateContactsChain := NewChain(CreateContactsNode)

	// AddNextChain
	PreserveChain.AddNextChain(LoginAdminChain, 1)
	LoginAdminChain.AddNextChain(QueryContactsChain, 0.7)
	LoginAdminChain.AddNextChain(CreateContactsChain, 0.3)
	//QueryContactsChain.AddNextChain()

}

func QueryContacts(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	GetAllContacts, err := cli.GetAllContacts()
	if err != nil {
		log.Fatalf("[Mock AccountID]GetAllContacts fail. The error occurs: %v", err)
		return nil, err
	}
	if GetAllContacts.Status != 1 {
		log.Fatalf("[Mock AccountID]GetAllContacts.Status != 1")
		return nil, err
	}

	randomIndex := rand.Intn(len(GetAllContacts.Data))
	ctx.Set(AccountID, GetAllContacts.Data[randomIndex].AccountId)
	ctx.Set(ContactsID, GetAllContacts.Data[randomIndex].Id)

	return nil, nil
}

func CreateContacts(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	CreateContactsInput := service.AdminContacts{
		Id:             faker.UUIDHyphenated(),
		AccountId:      faker.UUIDHyphenated(),
		Name:           faker.Name(),
		DocumentType:   rand.Intn(1),
		DocumentNumber: generateDocumentNumber(),
		PhoneNumber:    faker.PhoneNumber,
	}
	CreateContacts, err := cli.AddContact(&CreateContactsInput)
	if err != nil {
		log.Fatalf("[Mock AccountID] CreateContacts error occurs: %v", err)
		return nil, err
	}
	if CreateContacts.Status != 1 {
		log.Fatalf("[Mock AccountID] CreateContacts.Status != 1")
		return nil, err
	}

	ctx.Set(AccountID, CreateContacts.Data.AccountId)
	ctx.Set(ContactsID, CreateContacts.Data.Id)

	return nil, nil
}

// Preserve Behaviors
func Preserve(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}
	OrderTicketsInfo := service.OrderTicketsInfo{
		AccountID:       ctx.Get(AccountID).(string),  // Query:Create = 0.7 : 0.3
		ContactsID:      ctx.Get(ContactsID).(string), // Query:Create = 0.7 : 0.3
		TripID:          ctx.Get(TripID).(string),     // Query:Create = 0.7 : 0.3
		SeatType:        ctx.Get(SeatType).(int),      //
		LoginToken:      ctx.Get(LoginToken).(string),
		Date:            ctx.Get(Date).(string),
		From:            ctx.Get(From).(string),
		To:              ctx.Get(To).(string),
		Assurance:       ctx.Get(Assurance).(int),
		FoodType:        ctx.Get(FoodType).(int),
		StationName:     ctx.Get(StationName).(string),
		StoreName:       ctx.Get(StoreName).(string),
		FoodName:        ctx.Get(FoodName).(string),
		FoodPrice:       ctx.Get(FoodPrice).(float64),
		HandleDate:      ctx.Get(HandleDate).(string),
		ConsigneeName:   ctx.Get(ConsigneeName).(string),
		ConsigneePhone:  ctx.Get(ConsigneePhone).(string),
		ConsigneeWeight: ctx.Get(ConsigneeWeight).(float64),
		IsWithin:        ctx.Get(IsWithin).(bool),
	}
	PreserveResp, err := cli.Preserve(&OrderTicketsInfo)
	if err != nil {
		return nil, err
	}
	if PreserveResp.Status != 1 {
		return nil, fmt.Errorf("preserve order tickets fail. PreserveResp.Status != 1, get %v", PreserveResp.Status)
	}
	fmt.Printf("The Status is: %v, and PreserveResp Data: %v\n", PreserveResp.Status, PreserveResp.Data)
	fmt.Printf("PreserveBehaviors(Chain) Ends. End time: %v", time.Now().String())

	//return nil, err
	return &(NodeResult{false}), nil
}
