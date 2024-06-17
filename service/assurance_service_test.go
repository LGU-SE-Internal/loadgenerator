package service

import (
	"testing"

	"github.com/go-faker/faker/v4"
)

func TestSvcImpl_AddUpdateDeleteAssurance(t *testing.T) {
	//cli, _ := GetAdminClient()
	cli, _ := GetBasicClient()
	//MockedAssuranceID := faker.UUIDHyphenated()
	MockedOrderID := faker.UUIDHyphenated()

	// Create a new assurance
	addResp, err := cli.CreateNewAssurance(1, MockedOrderID) // Assuming typeIndex 1
	if err != nil {
		t.Errorf("CreateNewAssurance failed: %v", err)
	}
	t.Logf("CreateNewAssurance response: %+v", addResp)

	// Get all assurances
	assurances, err3 := cli.GetAllAssurances()
	if err3 != nil {
		t.Errorf("GetAllAssurances failed: %v", err3)
	}
	t.Logf("GetAllAssurances response: %+v", assurances)

	// Get all assurances types
	assuranceTypes, err9 := cli.GetAllAssuranceTypes()
	if err9 != nil {
		t.Errorf("GetAllAssuranceTypes failed: %v", err9)
	}
	t.Logf("GetAllAssuranceTypes response: %+v", assuranceTypes)

	var GetAssuranceID string
	var GetOrderID string
	if len(assurances.Data) > 0 {
		GetAssuranceID = assurances.Data[0].Id
		GetOrderID = assurances.Data[0].OrderId
	}

	// Modify the assurance
	updateResp, err1 := cli.ModifyAssurance(GetAssuranceID, GetOrderID, 1) // Assuming typeIndex 2
	if err1 != nil {
		t.Errorf("ModifyAssurance failed: %v", err1)
	}
	t.Logf("ModifyAssurance response: %+v", updateResp)

	// Test GetAssuranceByID
	GetAssuranceByIDResponse, err7 := cli.GetAssuranceByID(GetAssuranceID)
	if err7 != nil {
		t.Errorf("GetAssuranceByID failed: %v", err7)
	}
	t.Logf("GetAssuranceByID response: %+v", GetAssuranceByIDResponse)

	// Test FindAssuranceByOrderID
	FindAssuranceByOrderID, err4 := cli.FindAssuranceByOrderID(GetOrderID)
	if err4 != nil {
		t.Errorf("FindAssuranceByOrderID failed: %v", err4)
	}
	t.Logf("FindAssuranceByOrderID response: %+v", FindAssuranceByOrderID)

	// Delete the assurance by ID
	deleteResp_assurance, err2 := cli.DeleteAssuranceByID(GetAssuranceID)
	//deleteResp, err2 := cli.DeleteAssuranceByID("274dcd5a-e873-47bf-8ac2-5365a287742f")
	if err2 != nil {
		t.Errorf("DeleteAssuranceByAssuranceID failed: %v", err2)
	}
	t.Logf("DeleteAssuranceByAssuranceID response: %+v", deleteResp_assurance)

	deleteResp_order, err8 := cli.DeleteAssuranceByOrderID(GetOrderID)
	//deleteResp, err2 := cli.DeleteAssuranceByID("274dcd5a-e873-47bf-8ac2-5365a287742f")
	if err8 != nil {
		t.Errorf("DeleteAssuranceByOrderID failed: %v", err8)
	}
	t.Logf("DeleteAssuranceByOrderID response: %+v", deleteResp_order)
}
