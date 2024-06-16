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

	// Modify the assurance
	updateResp, err1 := cli.ModifyAssurance("f70ff11c-fa63-476a-a75b-adea2c7f728a", "959b783b-fbdd-4faa-9d6c-f3476ffc893b", 1) // Assuming typeIndex 2
	if err1 != nil {
		t.Errorf("ModifyAssurance failed: %v", err1)
	}
	t.Logf("ModifyAssurance response: %+v", updateResp)

	// Get all assurances
	assurances, err3 := cli.GetAllAssurances()
	if err3 != nil {
		t.Errorf("GetAllAssurances failed: %v", err3)
	}
	t.Logf("GetAllAssurances response: %+v", assurances)

	// Delete the assurance by ID
	var deleteID string
	if len(assurances.Data) > 0 {
		deleteID = assurances.Data[len(assurances.Data)-1].Id
	} else {
		t.Errorf("GetAllAssurances returned empty data")
	}
	deleteResp, err2 := cli.DeleteAssuranceByID(deleteID)
	//deleteResp, err2 := cli.DeleteAssuranceByID("274dcd5a-e873-47bf-8ac2-5365a287742f")
	if err2 != nil {
		t.Errorf("DeleteAssuranceByID failed: %v", err2)
	}
	t.Logf("DeleteAssuranceByID response: %+v", deleteResp)

	// Test GetAssuranceByID
	GetAssuranceByIDResponse, err3 := cli.GetAssuranceByID("b57a73c6-1721-494f-8b66-598156e37080")
	if err3 != nil {
		t.Errorf("GetAssuranceByID failed: %v", err3)
	}
	t.Logf("GetAssuranceByID response: %+v", GetAssuranceByIDResponse)

	// Test FindAssuranceByOrderID
	FindAssuranceByOrderID, err4 := cli.FindAssuranceByOrderID("0afd11c3-c0f5-489e-b9e8-34ba74688392")
	if err4 != nil {
		t.Errorf("FindAssuranceByOrderID failed: %v", err4)
	}
	t.Logf("FindAssuranceByOrderID response: %+v", FindAssuranceByOrderID)
}
