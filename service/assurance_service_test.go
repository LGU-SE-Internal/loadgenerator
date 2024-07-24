package service

import (
	"github.com/go-faker/faker/v4"
	"strconv"
	"testing"
)

func TestSvcImpl_AddUpdateDeleteAssurance(t *testing.T) {
	//cli, _ := GetAdminClient()
	cli, _ := GetBasicClient()

	var CreatedExistedOrderID string
	var orderSvc OrderService = cli

	// Get the ID in order to create a new assurance.
	//QueryAllOrderInfo, err := orderSvc.ReqFindAllOrder()

	// Create a new order
	randomContact := getRandomContact()
	createdOrder := Order{
		AccountId:              randomContact.AccountId,
		BoughtDate:             faker.Date(),
		CoachNumber:            RandomIntBetween(1, 10),
		ContactsDocumentNumber: strconv.Itoa(RandomIntBetween(1, 10)),
		ContactsName:           randomContact.Name,
		DifferenceMoney:        "",
		DocumentType:           0,
		From:                   RandomProvincialCapitalEN(),
		Id:                     "nil",
		Price:                  RandomDecimalStringBetween(1, 10),
		SeatClass:              GetTrainTicketClass(),
		SeatNumber:             GenerateSeatNumber(),
		Status:                 0,
		To:                     RandomProvincialCapitalEN(),
		TrainNumber:            GenerateTripId(),
		TravelDate:             faker.Date(),
		TravelTime:             faker.TimeString(),
	}

	newOrderInfo, err := orderSvc.ReqCreateNewOrder(&createdOrder)
	if err != nil {
		t.Errorf("Query All Order Info failed, err:%v\n", err)
	}

	//if len(QueryAllOrderInfo.Data) > 0 {
	//	CreatedExistedOrderID = QueryAllOrderInfo.Data[0].Id
	//} else {
	//	t.Errorf("Query All Order Info failed -> CreatedExistedOrderID fails, err:%v\n", err)
	//}

	CreatedExistedOrderID = newOrderInfo.Data.Id

	//Create a new assurance
	addResp, err := cli.CreateNewAssurance(1, CreatedExistedOrderID) // typeIndex 1 -> TRAFFIC_ACCIDENT
	if err != nil {
		t.Errorf("CreateNewAssurance failed: %v", err)
	}
	if addResp.Msg == "Already exists" {
		t.Log("Order ID found, skip")
		t.Skip()
	}
	if addResp.Data.OrderId != CreatedExistedOrderID {
		t.Errorf("Request failed, addResp.Data.OrderId:%s, expected: %s", addResp.Data.OrderId, CreatedExistedOrderID)
	}
	if addResp.Data.Type != "TRAFFIC_ACCIDENT" {
		t.Errorf("Request failed, addResp.Data.Type are expected to be 'TRAFFIC_ACCIDENT' but actually: %v", addResp.Data.Type)
	}
	existedAssurance := addResp.Data

	// Get all assurances
	assurances, err3 := cli.GetAllAssurances()
	if err3 != nil {
		t.Errorf("GetAllAssurances failed: %v", err3)
	}
	found := false
	for _, assurance := range assurances.Data {
		if assurance.Id == existedAssurance.Id &&
			assurance.OrderId == existedAssurance.OrderId &&
			assurance.TypeName == "Traffic Accident Assurance" {
			found = true
		}
	}
	if !found {
		t.Errorf("Request failed, assurance not found.")
	}

	// Get all assurances types
	assuranceTypes, err9 := cli.GetAllAssuranceTypes()
	if err9 != nil {
		t.Errorf("GetAllAssuranceTypes failed: %v", err9)
	}
	found1 := true
	for _, assuranceType := range assuranceTypes.Data {
		if assuranceType.Name != "Traffic Accident Assurance" {
			found1 = false
		}
	}
	if !found1 {
		t.Errorf("Request failed, assurance types not found.")
	}

	// Test GetAssuranceByID
	GetAssuranceByIDResponse, err7 := cli.GetAssuranceByID(existedAssurance.Id)
	if err7 != nil {
		t.Errorf("GetAssuranceByID failed: %v", err7)
	}
	if GetAssuranceByIDResponse.Status != 1 {
		t.Errorf("GetAssuranceByIDResponse.Status = %d, want 1", GetAssuranceByIDResponse.Status)
	}
	if GetAssuranceByIDResponse.Data.Id != existedAssurance.Id {
		t.Errorf("GetAssuranceByIDResponse failed.")
	}
	t.Logf("GetAssuranceByID response: %+v", GetAssuranceByIDResponse)

	// Test FindAssuranceByOrderID
	FindAssuranceByOrderID, err4 := cli.FindAssuranceByOrderID(existedAssurance.OrderId)
	if err4 != nil {
		t.Errorf("FindAssuranceByOrderID failed: %v", err4)
	}
	if FindAssuranceByOrderID.Status != 1 {
		t.Errorf("FindAssuranceByOrderID.Status != 1")
	}
	if FindAssuranceByOrderID.Data.OrderId != existedAssurance.OrderId {
		t.Errorf("FindAssuranceByOrderID failed.")
	}
	t.Logf("FindAssuranceByOrderID response: %+v", FindAssuranceByOrderID)

	// Modify the assurance
	updateResp, err1 := cli.ModifyAssurance(existedAssurance.Id, CreatedExistedOrderID, 1) // Assuming typeIndex 2
	if err1 != nil {
		t.Errorf("ModifyAssurance failed: %v", err1)
	}
	if updateResp.Status != 1 {
		t.Errorf("updateResp.Status != 1")
	}
	booleanModify := false
	if updateResp.Data.OrderId == existedAssurance.OrderId &&
		updateResp.Data.Id == existedAssurance.Id &&
		updateResp.Data.Type == "TRAFFIC_ACCIDENT" {
		booleanModify = true
	}
	if !booleanModify {
		t.Errorf("UpdateAssurance failed.")
	}
	t.Logf("ModifyAssurance response: %+v", updateResp)

	// Delete the assurance by ID
	deleteResp_assurance, err2 := cli.DeleteAssuranceByID(existedAssurance.Id)
	//deleteResp, err2 := cli.DeleteAssuranceByID("274dcd5a-e873-47bf-8ac2-5365a287742f")
	if err2 != nil {
		t.Errorf("DeleteAssuranceByAssuranceID failed: %v", err2)
	}
	if deleteResp_assurance.Status != 1 {
		t.Errorf("deleteResp_assurance.Status != 1")
	}
	t.Logf("DeleteAssuranceByAssuranceID response: %+v", deleteResp_assurance)

	// Delete the assurance by OrderID
	deleteResp_order, err8 := cli.DeleteAssuranceByOrderID(existedAssurance.OrderId)
	//deleteResp, err2 := cli.DeleteAssuranceByID("274dcd5a-e873-47bf-8ac2-5365a287742f")
	if err8 != nil {
		t.Errorf("DeleteAssuranceByOrderID failed: %v", err8)
	}
	if deleteResp_order.Status != 1 {
		t.Errorf("DeleteAssuranceByOrderID status != 1")
	}
	t.Logf("DeleteAssuranceByOrderID response: %+v", deleteResp_order)
}
