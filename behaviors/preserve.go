package behaviors

import (
	"fmt"
	"github.com/Lincyaw/loadgenerator/service"
	"github.com/go-faker/faker/v4"
	"log"
	"math/rand"
	"strings"
	"time"
)

const (
	// Preserve
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

	// Contacts
	//Id             string `json:"id"`
	//AccountId      string `json:"accountId"`
	Name           = "name" // ?= ConsigneeName
	DocumentType   = "documentType"
	DocumentNumber = "documentNumber"
	PhoneNumber    = "phoneNumber" // ?= ConsigneePhone

	// Assurance
	OrderId   = "orderId"
	TypeIndex = "typeIndex"
	TypeName  = "typeName"
	TypePrice = "typePrice"

	// Route
	RouteID = "routeId"
)

var PreserveChain *Chain

func init() {
	// ------- init --------
	// Main Chain
	PreserveChain = NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		fmt.Printf("PreserveBehaviors(Chain) Statrs. Starts time: %v", time.Now().String())
		return nil, nil
	}))

	// ------------------- NewFuncNode -----------------------
	// LoginAdmin
	LoginAdminNode := NewFuncNode(LoginAdmin)
	// Contacts
	QueryContactsNode := NewFuncNode(QueryContacts)
	CreateContactsNode := NewFuncNode(CreateContacts)
	// Assurance
	QueryAssuranceNode := NewFuncNode(QueryAssurance)
	CreateAssuranceNode := NewFuncNode(CreateAssurance)
	// Route
	QueryRouteNode := NewFuncNode(QueryRoute)
	CreateRouteNode := NewFuncNode(CreateRoute)
	// Trip
	QueryTripNode := NewFuncNode(QueryTrip)
	CreateTripNode := NewFuncNode(CreateTrip)

	// ------------------- NewChain -----------------------
	// LoginAdmin
	LoginAdminChain := NewChain(LoginAdminNode) // done
	// Contacts
	QueryContactsChain := NewChain(QueryContactsNode)   // done
	CreateContactsChain := NewChain(CreateContactsNode) // done
	// Assurance
	AssuranceBehaviorsChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		fmt.Printf("CreateAssuranceChain. Starts time: %v", time.Now().String())
		return nil, nil
	}))
	QueryAssuranceChain := NewChain(QueryAssuranceNode)
	CreateAssuranceChain := NewChain(CreateAssuranceNode)
	// Trip
	QueryTripChain := NewChain(QueryTripNode)
	CreateTripBehaviorsChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		fmt.Printf("CreateTripChain. Starts time: %v", time.Now().String())
		return nil, nil
	}))
	CreateTripChain := NewChain(CreateTripNode)
	// Route
	QueryRouteChain := NewChain(QueryRouteNode)
	CreateRouteChain := NewChain(CreateRouteNode)

	// ------------------- AddNextChain -----------------------
	// PreserveChain
	PreserveChain.AddNextChain(LoginAdminChain, 1)
	// LoginAdminChain
	LoginAdminChain.AddNextChain(QueryContactsChain, 0.7)
	LoginAdminChain.AddNextChain(CreateContactsChain, 0.3)
	// AssuranceBehaviorsChain
	AssuranceBehaviorsChain.AddNextChain(QueryAssuranceChain, 0.7)
	AssuranceBehaviorsChain.AddNextChain(CreateAssuranceChain, 0.3)
	// QueryContactsChain
	QueryContactsChain.AddNextChain(AssuranceBehaviorsChain, 1)
	// CreateContactsChain
	CreateContactsChain.AddNextChain(AssuranceBehaviorsChain, 1)
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
	ctx.Set(Name, CreateContacts.Data.Name)
	ctx.Set(DocumentType, CreateContacts.Data.DocumentType)
	ctx.Set(DocumentNumber, CreateContacts.Data.DocumentNumber)
	ctx.Set(PhoneNumber, CreateContacts.Data.PhoneNumber)

	return nil, nil
}

func QueryAssurance(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	Assurances, err := cli.GetAllAssurances()
	if err != nil {
		log.Fatalf("GetAllAssurances failed: %v", err)
		return nil, err
	}
	if Assurances.Status != 1 {
		log.Fatalf("Assurances status is not 1: %d", Assurances.Status)
		return nil, err
	}

	randomIndex := rand.Intn(len(Assurances.Data))
	ctx.Set(OrderId, Assurances.Data[randomIndex].OrderId)
	ctx.Set(TypeIndex, Assurances.Data[randomIndex].TypeIndex)
	ctx.Set(TypeName, Assurances.Data[randomIndex].TypeName)
	ctx.Set(TypePrice, Assurances.Data[randomIndex].TypePrice)

	return nil, nil
}

func CreateAssurance(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	//Create a new assurance
	TheOrderID := ctx.Get(OrderId).(string)
	addAssuranceResp, err := cli.CreateNewAssurance(1, TheOrderID) // typeIndex 1 -> TRAFFIC_ACCIDENT
	if err != nil {
		log.Fatalf("CreateNewAssurance failed: %v", err)
		return nil, err
	}
	if addAssuranceResp.Msg == "Already exists" {
		log.Fatalf("Order ID found, skip")
		return nil, err
	}
	if addAssuranceResp.Data.OrderId != TheOrderID {
		log.Fatalf("Request failed, addAssuranceResp.Data.OrderId:%s, expected: %s", addAssuranceResp.Data.OrderId, TheOrderID)
		return nil, err
	}
	if addAssuranceResp.Data.Type != "TRAFFIC_ACCIDENT" {
		log.Fatalf("Request failed, addAssuranceResp.Data.Type are expected to be 'TRAFFIC_ACCIDENT' but actually: %v", addAssuranceResp.Data.Type)
		return nil, err
	}

	ctx.Set(OrderId, addAssuranceResp.Data.OrderId)
	ctx.Set(TypeIndex, addAssuranceResp.Data..TypeIndex)
	ctx.Set(TypeName, Assurances.Data[randomIndex].TypeName)
	ctx.Set(TypePrice, Assurances.Data[randomIndex].TypePrice)

	return nil, nil
}

func QueryTrip(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}
	QueryAllTripResp, err := cli.QueryAllTrip()
	if err != nil {
		log.Fatalf("Request failed, err %s", err)
		return nil, err
	}
	if QueryAllTripResp.Status != 1 {
		log.Fatalf("Request failed, status: %d", QueryAllTripResp.Status)
		return nil, err
	}

	randomIndex := rand.Intn(len(QueryAllTripResp.Data))
	ctx.Set(TripID, QueryAllTripResp.Data[randomIndex].TripId)
	ctx.Set(From, QueryAllTripResp.Data[randomIndex].StartStationName)
	ctx.Set(From, QueryAllTripResp.Data[randomIndex].TerminalStationName)
	ctx.Set(Date, QueryAllTripResp.Data[randomIndex].StartTime)
	ctx.Set(StationName, QueryAllTripResp.Data[randomIndex].StationsName)
	ctx.Set(HandleDate, QueryAllTripResp.Data[randomIndex].EndTime)

	return nil, nil
}

func CreateTrip(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	// Mock para
	MockedLoginId := ctx.Get(LoginToken).(string)
	MockedTripId := GenerateTripId()
	MockedTrainTypeName := generateTrainTypeName(MockedTripId) /*"GaoTieSeven"*/
	MockedRouteID := ctx.Get(RouteID).(string)
	MockedStartStationName := ctx.Get(From).(string)
	MockedStationsName := /*strings.Join(AllRoutesByQuery.Data[0].Stations, ",")*/ ctx.Get(StationName).(string)
	MockedTerminalStationName := ctx.Get(To).(string)
	MockedStartTime := getRandomTime()
	MockedEndTime := getRandomTime(WithStartTime(MockedStartTime))

	// Mock input
	travelInfo := service.TravelInfo{
		LoginID:          MockedLoginId,
		TripID:           MockedTripId,
		TrainTypeName:    MockedTrainTypeName,
		RouteID:          MockedRouteID,
		StartStationName: MockedStartStationName,
		//StationsName:        fmt.Sprintf("%v,%v", MockedStartStationName, MockedTerminalStationName),
		StationsName:        MockedStationsName,
		TerminalStationName: MockedTerminalStationName,
		StartTime:           MockedStartTime,
		EndTime:             MockedEndTime,
	}

	// Create Test
	createResp, err := cli.CreateTrip(&travelInfo)
	if err != nil {
		log.Fatalf("CreateTrip request failed, err %s", err)
		return nil, err
	}
	if createResp.Status != 1 {
		log.Fatalf("CreateTrip failed: %s", createResp.Msg)
		return nil, err
	}
	if createResp.Msg == "Already exists" {
		log.Fatalf("Already exists: %s", createResp.Msg)
		return nil, err
	}
	isMatch := false
	if /*createResp.Data.Id == travelInfo.LoginID &&*/
	createResp.Data.StationsName == toLowerCaseAndRemoveSpaces(travelInfo.StationsName) &&
		createResp.Data.StartStationName == toLowerCaseAndRemoveSpaces(travelInfo.StartStationName) &&
		createResp.Data.TerminalStationName == toLowerCaseAndRemoveSpaces(travelInfo.TerminalStationName) &&
		createResp.Data.StartTime == travelInfo.StartTime &&
		createResp.Data.EndTime == travelInfo.EndTime &&
		createResp.Data.TrainTypeName == travelInfo.TrainTypeName &&
		createResp.Data.RouteId == travelInfo.RouteID {
		isMatch = true
	}
	if !isMatch {
		log.Fatalf("CreateTrip failed: %s. Except: %v, but get: %v", createResp.Msg, travelInfo, createResp.Data)
		return nil, err
	}

	ctx.Set(TripID, createResp.Data.TripId)
	ctx.Set(From, createResp.Data.StartStationName)
	ctx.Set(From, createResp.Data.TerminalStationName)
	ctx.Set(Date, createResp.Data.StartTime)
	ctx.Set(StationName, createResp.Data.StationsName)
	ctx.Set(HandleDate, createResp.Data.EndTime)

	return nil, nil
}

func QueryRoute(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	AllRoutesByQuery, err := cli.QueryAllRoutes()
	if err != nil {
		log.Fatalf("Request failed, err2 %s", err)
		return nil, err
	}
	if AllRoutesByQuery.Status != 1 {
		log.Fatal("AllRoutes_By_Query.Status != 1")
		return nil, err
	}

	randomIndex := rand.Intn(len(AllRoutesByQuery.Data))
	ctx.Set(From, AllRoutesByQuery.Data[randomIndex].StartStation)
	ctx.Set(To, AllRoutesByQuery.Data[randomIndex].EndStation)
	ctx.Set(StationName, getMiddleElements(strings.Join(AllRoutesByQuery.Data[randomIndex].Stations, ",")))
	ctx.Set(RouteID, AllRoutesByQuery.Data[randomIndex].Id)

	return nil, nil
}

func CreateRoute(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	// Create
	MockedID := faker.UUIDHyphenated()
	MockedStartStation := faker.GetRealAddress().City
	MockedEndStation := faker.GetRealAddress().City
	MockedStationList := fmt.Sprintf("%s,%s,%s", MockedStartStation, faker.GetRealAddress().City, MockedEndStation)
	MockedDistanceList := fmt.Sprintf("%d,%d,%d", rand.Intn(30), rand.Intn(30), rand.Intn(30))
	input := service.RouteInfo{
		ID:           MockedID,
		StartStation: MockedStartStation,
		EndStation:   MockedEndStation,
		StationList:  MockedStationList,
		DistanceList: MockedDistanceList,
	}
	resp, err := cli.CreateAndModifyRoute(&input)
	if err != nil {
		log.Fatalf("Request failed, err %s", err)
		return nil, err
	}
	if resp.Msg == "Already exists" {
		log.Fatalf("Route already exists, skip")
		return nil, err
	}
	if resp.Data.Id != input.ID {
		log.Fatalf("Route ID does not match, expect %s, got %s", input.ID, resp.Data.Id)
		return nil, err
	}
	if resp.Data.StartStation != input.StartStation {
		log.Fatalf("StartStation does not match, expect %s, got %s", input.StartStation, resp.Data.StartStation)
		return nil, err
	}
	if resp.Data.EndStation != input.EndStation {
		log.Fatalf("StartStation does not match, expect %s, got %s", input.StartStation, resp.Data.StartStation)
		return nil, err
	}
	if StringSliceToString(resp.Data.Stations) != ConvertCommaSeparatedToBracketed(input.StationList) {
		log.Fatalf("StationList does not match, expect %s, got %s", ConvertCommaSeparatedToBracketed(input.StationList), StringSliceToString(resp.Data.Stations))
		return nil, err
	}
	if IntSliceToString(resp.Data.Distances) != ConvertCommaSeparatedToBracketed(input.DistanceList) {
		log.Fatalf("DistanceList does not match, expect %s, got %s", ConvertCommaSeparatedToBracketed(input.DistanceList), IntSliceToString(resp.Data.Distances))
		return nil, err
	}

	ctx.Set(From, resp.Data.StartStation)
	ctx.Set(To, resp.Data.EndStation)
	ctx.Set(StationName, getMiddleElements(strings.Join(resp.Data.Stations, ",")))
	ctx.Set(RouteID, resp.Data.Id)

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
		SeatType:        ctx.Get(SeatType).(int),
		LoginToken:      ctx.Get(LoginToken).(string), // Query:Create = 0.7 : 0.3
		Date:            ctx.Get(Date).(string),       // Query:Create = 0.7 : 0.3
		From:            ctx.Get(From).(string),       // Query:Create = 0.7 : 0.3
		To:              ctx.Get(To).(string),         // Query:Create = 0.7 : 0.3
		Assurance:       ctx.Get(Assurance).(int),
		FoodType:        ctx.Get(FoodType).(int),
		StationName:     ctx.Get(StationName).(string), // Query:Create = 0.7 : 0.3
		StoreName:       ctx.Get(StoreName).(string),
		FoodName:        ctx.Get(FoodName).(string),
		FoodPrice:       ctx.Get(FoodPrice).(float64),
		HandleDate:      ctx.Get(HandleDate).(string), // Query:Create = 0.7 : 0.3
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
