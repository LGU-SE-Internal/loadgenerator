package service

import (
	"testing"

	"github.com/go-faker/faker/v4"
)

func TestSvcImpl_Preserve(t *testing.T) {
	cli, _ := GetBasicClient()

	var contactsSvc ContactsService = cli
	var getContactsId string

	GetAllContactsId, err := contactsSvc.GetAllContacts()
	if err != nil {
		t.Error(err)
	}
	t.Log(GetAllContactsId)

	if len(GetAllContactsId.Data) > 0 {
		//getContactsId = *(GetAllContactsId.Data[0].AccountId)
		getContactsId = GetAllContactsId.Data[0].Id
	}

	MockedAccountID := faker.UUIDHyphenated()
	MockedContactsID := getContactsId
	MockedTripID := faker.UUIDHyphenated()
	MockedLoginToken := faker.UUIDHyphenated()
	MockedDate := faker.Date()
	MockedFromCity := faker.GetRealAddress().City
	MockedToCity := faker.GetRealAddress().City
	MockedHandleDate := faker.Date()
	MockedConsigneeName := faker.Name()
	MockedConsigneePhone := faker.PhoneNumber

	// Mock data
	orderTicketsInfo := OrderTicketsInfo{
		AccountID:       MockedAccountID,
		ContactsID:      MockedContactsID,
		TripID:          MockedTripID,
		SeatType:        1,
		LoginToken:      MockedLoginToken,
		Date:            MockedDate,
		From:            MockedFromCity,
		To:              MockedToCity,
		Assurance:       1,
		FoodType:        1,
		StationName:     "Shenzhen Bei",
		StoreName:       "Happy Store",
		FoodName:        "spaghetti",
		FoodPrice:       10.00,
		HandleDate:      MockedHandleDate,
		ConsigneeName:   MockedConsigneeName,
		ConsigneePhone:  MockedConsigneePhone,
		ConsigneeWeight: 7.77,
		IsWithin:        true,
	}

	// Test Preserve
	preserveResp, err := cli.Preserve(&orderTicketsInfo)
	if err != nil {
		t.Errorf("Preserve request failed, err %s", err)
	}
	t.Logf("Preserve response: %+v", preserveResp)
}
