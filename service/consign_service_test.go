package service

import (
	"testing"

	"github.com/go-faker/faker/v4"
)

func TestSvcImpl_AddUpdateQueryDeleteConsign(t *testing.T) {
	cli, _ := GetAdminClient()
	var consignSvc ConsignService = cli

	// Mock data
	MockedId := faker.UUIDHyphenated()
	MockedAccountId := faker.UUIDHyphenated()
	MockedOrderId := faker.UUIDHyphenated()
	MockedHandleDate := faker.Date()
	MockedTargetDate := faker.Date()
	MockedFromPlace := "suzhou"
	MockedToPlace := "beijing"
	MockedConsignee := faker.Name()
	MockedPhone := faker.PhoneNumber

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
	isMatch := false
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
		isMatch = true
	}
	if !isMatch {
		t.Errorf("Creation not match. Expect: %v, but get: %v", insertReq, insertResp.Data)
	}
	t.Logf("InsertConsignRecord response: %+v", insertResp)
	existedConsign := insertResp.Data

	// Query consign records by account ID
	consignsByAccountId, err := consignSvc.QueryByAccountId(MockedAccountId)
	if err != nil {
		t.Errorf("QueryByAccountId failed: %v", err)
	}
	if consignsByAccountId.Status != 1 {
		t.Errorf("consignsByAccountId failed")
	}
	found := false
	for _, consign := range consignsByAccountId.Data {
		if consign.IsWithin == existedConsign.IsWithin &&
			consign.To == existedConsign.To &&
			consign.Weight == existedConsign.Weight &&
			consign.ID == existedConsign.ID &&
			consign.Phone == existedConsign.Phone &&
			consign.HandleDate == existedConsign.HandleDate &&
			consign.TargetDate == existedConsign.TargetDate &&
			consign.OrderID == existedConsign.OrderID &&
			consign.Consignee == existedConsign.Consignee &&
			consign.From == existedConsign.From &&
			consign.AccountID == existedConsign.AccountID {
			found = true
		}
	}
	if !found {
		t.Errorf("Can not find consign by accountId.")
	}
	t.Logf("QueryByAccountId response: %+v", consignsByAccountId)

	// Query consign records by order ID
	consignsByOrderId, err := consignSvc.QueryByOrderId(MockedOrderId)
	if err != nil {
		t.Errorf("QueryByOrderId failed: %v", err)
	}
	if consignsByOrderId.Status != 1 {
		t.Errorf("consignsByOrderId.Status = 1")
	}
	isMatch1 := false
	if consignsByOrderId.Data.OrderId == existedConsign.OrderID &&
		consignsByOrderId.Data.Id == existedConsign.ID &&
		consignsByOrderId.Data.From == existedConsign.From &&
		consignsByOrderId.Data.To == existedConsign.To &&
		consignsByOrderId.Data.Phone == existedConsign.Phone &&
		consignsByOrderId.Data.Consignee == existedConsign.Consignee &&
		consignsByOrderId.Data.TargetDate == existedConsign.TargetDate &&
		consignsByOrderId.Data.HandleDate == existedConsign.HandleDate &&
		consignsByOrderId.Data.Weight == existedConsign.Weight &&
		consignsByOrderId.Data.AccountId == existedConsign.AccountID {
		isMatch1 = true
	}
	if !isMatch1 {
		t.Errorf("Can not find consign by orderId.")
	}
	t.Logf("QueryByOrderId response: %+v", consignsByOrderId)

	// Query consign records by consignee
	consignsByConsignee, err := consignSvc.QueryByConsignee(MockedConsignee)
	if err != nil {
		t.Errorf("QueryByConsignee failed: %v", err)
	}
	if consignsByConsignee.Status != 1 {
		t.Errorf("consignsByConsignee failed.")
	}
	isMatch2 := false
	for _, consign := range consignsByConsignee.Data {
		if consign.Id == existedConsign.ID &&
			consign.AccountId == existedConsign.AccountID &&
			consign.OrderId == existedConsign.OrderID &&
			consign.To == existedConsign.To &&
			consign.From == existedConsign.From &&
			consign.Weight == existedConsign.Weight &&
			consign.HandleDate == existedConsign.HandleDate &&
			consign.TargetDate == existedConsign.TargetDate &&
			consign.Phone == existedConsign.Phone &&
			consign.Consignee == existedConsign.Consignee {
			isMatch2 = true
		}
	}
	if !isMatch2 {
		t.Errorf("Can not find consign by consignee.")
	}
	t.Logf("QueryByConsignee response: %+v", consignsByConsignee)

	// Update the consign record
	updateReq := &Consign{
		ID:         MockedId,
		OrderID:    MockedOrderId,
		AccountID:  MockedAccountId,
		HandleDate: MockedHandleDate,
		TargetDate: MockedTargetDate,
		From:       MockedFromPlace,
		To:         MockedToPlace,
		Consignee:  MockedConsignee,
		Phone:      MockedPhone,
		Weight:     10.0,
		IsWithin:   true,
	}
	updateResp, err := consignSvc.UpdateConsignRecord(updateReq)
	if err != nil {
		t.Errorf("UpdateConsignRecord failed: %v", err)
	}
	if updateResp.Status != 1 {
		t.Errorf("updateResp.Status != 1")
	}
	isMatch3 := false
	if updateResp.Data.To == updateReq.To &&
		updateResp.Data.From == updateReq.From &&
		updateResp.Data.Phone == updateReq.Phone &&
		/*updateResp.Data.ID == updateReq.ID &&*/
		updateResp.Data.Consignee == updateReq.Consignee &&
		updateResp.Data.TargetDate == updateReq.TargetDate &&
		updateResp.Data.HandleDate == updateReq.HandleDate &&
		updateResp.Data.Weight == updateReq.Weight &&
		updateResp.Data.OrderID == updateReq.OrderID &&
		updateResp.Data.AccountID == updateReq.AccountID {
		isMatch3 = true
	}
	if !isMatch3 {
		t.Errorf("exopected: %v; but get: %v", updateReq, updateResp.Data)
	}
	t.Logf("UpdateConsignRecord response: %+v", updateResp)

}
