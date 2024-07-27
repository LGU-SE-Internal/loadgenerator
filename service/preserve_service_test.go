package service

import (
	"log"
	"testing"

	"github.com/go-faker/faker/v4"
)

func TestSvcImpl_Preserve(t *testing.T) {
	cli, _ := GetBasicClient()
	var preserveSvc PreserveService = cli

	/*	var contactsSvc ContactsService = cli
		var getContactsId string

		GetAllContactsId, err := contactsSvc.GetAllContacts()
		if err != nil {
			t.Error(err)
		}
		t.Log(GetAllContactsId)

		if len(GetAllContactsId.Data) > 0 {
			//getContactsId = *(GetAllContactsId.Data[0].AccountId)
			getContactsId = GetAllContactsId.Data[0].Id
		}*/

	loginResult, err := cli.ReqUserLogin(&UserLoginInfoReq{
		Password:         "111111",
		UserName:         "fdse_microservice",
		VerificationCode: "123",
	})
	if err != nil {
		log.Fatalln(err)
	}

	MockedAccountID := /*faker.UUIDHyphenated()*/ "4d2a46c7-71cb-4cf1-b5bb-b68406d9da6f"
	MockedContactsID := /*getContactsId*/ "95e046ea-812c-49a9-8d12-168f539bec43"
	MockedTripID := /*faker.UUIDHyphenated()*/ "92708982-77af-4318-be25-57ccb0ff69ad"
	MockedLoginToken := /*faker.UUIDHyphenated()*/ loginResult.Data.Token
	MockedDate := /*faker.Date()*/ "2025-05-04 09:00:00"
	MockedFromCity := /*faker.GetRealAddress().City*/ "suzhou"
	MockedToCity := /*faker.GetRealAddress().City*/ "beijing"
	MockedHandleDate := /*faker.Date()*/ "2025-07-11"
	MockedConsigneeName := /*faker.Name()*/ "Dr. Keenan Huel"
	MockedConsigneePhone := /*faker.PhoneNumber*/ faker.PhoneNumber

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
	preserveResp, err := preserveSvc.Preserve(&orderTicketsInfo)
	if err != nil {
		t.Errorf("Preserve request failed, err %s", err)
	}
	t.Logf("Preserve response: %+v", preserveResp)
}
