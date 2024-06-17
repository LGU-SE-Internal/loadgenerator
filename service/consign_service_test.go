package service

import (
	"testing"

	"github.com/go-faker/faker/v4"
)

func TestSvcImpl_AddUpdateQueryDeleteConsign(t *testing.T) {
	cli, _ := GetAdminClient()
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
		Weight:     10.0,
		IsWithin:   true,
	}
	insertResp, err := cli.InsertConsignRecord(insertReq)
	if err != nil {
		t.Errorf("InsertConsignRecord failed: %v", err)
	}
	t.Logf("InsertConsignRecord response: %+v", insertResp)

	// Query consign records by account ID
	consignsByAccountId, err := cli.QueryByAccountId(MockedAccountId)
	if err != nil {
		t.Errorf("QueryByAccountId failed: %v", err)
	}
	t.Logf("QueryByAccountId response: %+v", consignsByAccountId)

	// Query consign records by order ID
	consignsByOrderId, err := cli.QueryByOrderId(MockedOrderId)
	if err != nil {
		t.Errorf("QueryByOrderId failed: %v", err)
	}
	t.Logf("QueryByOrderId response: %+v", consignsByOrderId)

	// Query consign records by consignee
	consignsByConsignee, err := cli.QueryByConsignee(MockedConsignee)
	if err != nil {
		t.Errorf("QueryByConsignee failed: %v", err)
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
		Weight:     7.0,
		IsWithin:   true,
	}
	updateResp, err := cli.UpdateConsignRecord(updateReq)
	if err != nil {
		t.Errorf("UpdateConsignRecord failed: %v", err)
	}
	t.Logf("UpdateConsignRecord response: %+v", updateResp)
}
