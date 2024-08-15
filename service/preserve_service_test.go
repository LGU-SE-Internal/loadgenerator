package service

import (
	"log"
	"testing"
	"time"

	"github.com/go-faker/faker/v4"
)

func TestSvcImpl_Preserve(t *testing.T) {
	cli, _ := GetBasicClient()
	var preserveSvc PreserveService = cli

	// LoginToken
	loginResult, err := cli.ReqUserLogin(&UserLoginInfoReq{
		Password:         "111111",
		UserName:         "fdse_microservice",
		VerificationCode: "123",
	})
	if err != nil {
		log.Fatalln(err)
	}

	// Security Service
	// Mock Account ID
	var securitySvc SecurityService = cli

	//MockedID := faker.UUIDHyphenated()
	MockedName := faker.Name()
	MockedValue := generateRandomNumberString()
	MockedDescription := generateDescription()

	// Mock input
	input := &SecurityConfig{
		//ID:          MockedID,
		Name:        MockedName,
		Value:       MockedValue,
		Description: MockedDescription,
	}

	// Add Security AdminConfig
	addResp, err := securitySvc.AddNewSecurityConfig(input)
	if err != nil {
		t.Errorf("AddNewSecurityConfig failed: %v", err)
	}
	if addResp.Status != 1 {
		t.Errorf("[Security Service]addResp.Status != 1")
	}
	if addResp.Msg == "Already exists" {
		t.Logf("[Security Service]addResp.Msg => Already exists")
		t.Skip()
	}
	isMatch := false
	if /*addResp.Data.ID == input.ID &&*/
	addResp.Data.Value == input.Value &&
		addResp.Data.Description == input.Description &&
		addResp.Data.Name == input.Name {
		isMatch = true
	}
	if !isMatch {
		t.Errorf("[Security Service]Except: %v, get %v", input, addResp.Data)
	}
	t.Logf("AddNewSecurityConfig response: %+v", addResp)
	existedSecurity := addResp.Data

	// Contact Service
	// Mock Contacts ID
	var contactsSvc ContactsService = cli

	// CreateContact
	CreateContactsInput := AdminContacts{
		Id:        faker.UUIDHyphenated(),
		AccountId: existedSecurity.ID,
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
	isMatch1 := false
	if /*CreateContacts.Data.Id == CreateContactsInput.Id &&*/
	CreateContacts.Data.Name == CreateContactsInput.Name &&
		CreateContacts.Data.AccountId == CreateContactsInput.AccountId &&
		CreateContacts.Data.PhoneNumber == CreateContactsInput.PhoneNumber &&
		CreateContacts.Data.DocumentNumber == CreateContactsInput.DocumentNumber &&
		CreateContacts.Data.DocumentType == CreateContactsInput.DocumentType {
		isMatch1 = true
	}
	if !isMatch1 {
		t.Errorf("Create AdminContacts Fail. expect: %v, get %v", CreateContactsInput, CreateContacts.Data)
	}
	existedContacts := CreateContacts.Data

	// Travel Service
	var travelSvc TravelService = cli
	// QueryTraintype
	allTrip, err := travelSvc.QueryAllTrip()
	if err != nil {
		t.Errorf("QueryAllTrip failed: %v", err)
	}
	var existedTravel Trip
	for _, trip := range allTrip.Data {
		startTime, err := time.Parse("2006-01-02 15:04:05", trip.StartTime)
		if err != nil {
			t.Errorf("QueryAllTrip failed: %v", err)
		}
		if startTime.After(time.Now()) {
			existedTravel = trip
		}
	}

	// Consign Service
	var consignSvc ConsignService = cli

	// Mock data
	MockedId := faker.UUIDHyphenated()
	MockedAccountId := existedSecurity.ID
	MockedOrderId := faker.UUIDHyphenated()
	MockedHandleDate := existedTravel.StartTime
	MockedTargetDate := existedTravel.EndTime
	MockedFromPlace := existedTravel.StartStationName
	MockedToPlace := existedTravel.TerminalStationName
	MockedConsignee := existedContacts.Name
	MockedPhone := existedContacts.PhoneNumber

	// Insert a new consign record
	insertReq := &Consign{
		ID:         MockedId,
		OrderID:    MockedOrderId,
		AccountID:  MockedAccountId,
		HandleDate: MockedHandleDate,
		TargetDate: MockedTargetDate,
		From:       MockedFromPlace,
		To:         MockedToPlace,
		Consignee:  MockedConsignee,
		Phone:      MockedPhone,
		Weight:     7.0,
		IsWithin:   false,
	}
	insertResp, err := consignSvc.InsertConsignRecord(insertReq)
	if err != nil {
		t.Errorf("InsertConsignRecord failed: %v", err)
	}
	if insertResp.Msg == "Already exists" {

	}
	if insertResp.Status != 1 {
		t.Errorf("InsertConsignRecord failed: %v", insertResp.Status)
	}
	isMatch3 := false
	if /*insertResp.Data.ID == insertReq.ID &&*/
	insertResp.Data.IsWithin == insertReq.IsWithin &&
		insertResp.Data.AccountID == insertReq.AccountID &&
		insertResp.Data.From == insertReq.From &&
		insertResp.Data.Consignee == insertReq.Consignee &&
		insertResp.Data.OrderID == insertReq.OrderID &&
		insertResp.Data.Phone == insertReq.Phone &&
		insertResp.Data.TargetDate == insertReq.TargetDate &&
		insertResp.Data.HandleDate == insertReq.HandleDate &&
		insertResp.Data.To == insertReq.To &&
		insertResp.Data.Weight == insertReq.Weight {
		isMatch3 = true
	}
	if !isMatch3 {
		t.Errorf("Creation not match. Expect: %v, but get: %v", insertReq, insertResp.Data)
	}
	t.Logf("InsertConsignRecord response: %+v", insertResp)
	existedConsign := insertResp.Data

	// Mock Data End
	// Data Input
	preserveMockedLoginToken := loginResult.Data.Token
	preserveMockedAccountID := existedSecurity.ID
	preserveMockedContactsID := existedContacts.Id
	preserveMockedTripID := existedTravel.Id
	preserveMockedDate := /*faker.Date()*/ /*"2025-05-04 09:00:00"*/ existedConsign.TargetDate
	preserveMockedHandleDate := /*faker.Date()*/ /*"2025-07-11"*/ existedConsign.HandleDate
	preserveMockedFromCity := /*faker.GetRealAddress().City*/ /*"suzhou"*/ existedConsign.From
	preserveMockedToCity := /*faker.GetRealAddress().City*/ /*"beijing"*/ existedConsign.To
	preserveMockedConsigneeName := /*faker.Name()*/ /*"Dr. Keenan Huel"*/ existedConsign.Consignee
	preserveMockedConsigneePhone := /*faker.PhoneNumber*/ /*faker.PhoneNumber*/ existedConsign.Phone

	// Mock data
	orderTicketsInfo := OrderTicketsInfo{
		AccountID:       preserveMockedAccountID,
		ContactsID:      preserveMockedContactsID,
		TripID:          preserveMockedTripID,
		SeatType:        1,
		LoginToken:      preserveMockedLoginToken,
		Date:            preserveMockedDate,
		From:            preserveMockedFromCity,
		To:              preserveMockedToCity,
		Assurance:       1,
		FoodType:        1,
		StationName:     "Shenzhen Bei",
		StoreName:       "Happy Store",
		FoodName:        "spaghetti",
		FoodPrice:       10.00,
		HandleDate:      preserveMockedHandleDate,
		ConsigneeName:   preserveMockedConsigneeName,
		ConsigneePhone:  preserveMockedConsigneePhone,
		ConsigneeWeight: 7.77,
		IsWithin:        true,
	}

	// Test Preserve
	preserveResp, err := preserveSvc.Preserve(&orderTicketsInfo)
	if err != nil {
		t.Errorf("Preserve request failed, err %s", err)
	}
	t.Logf("Preserve response: %+v", preserveResp)
}
