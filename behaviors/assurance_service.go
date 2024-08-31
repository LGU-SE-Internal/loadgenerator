package behaviors

import (
	"fmt"
	"github.com/Lincyaw/loadgenerator/service"
	log "github.com/sirupsen/logrus"
	"math/rand"
)

func QueryAssurance(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(service.AssuranceService)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	Assurances, err := cli.GetAllAssuranceTypes()
	if err != nil {
		log.Errorf("GetAllAssurances failed: %v", err)
		return nil, err
	}
	if Assurances.Status != 1 {
		log.Errorf("Assurances status is not 1: %+v, early exiting", Assurances)
		return &NodeResult{Continue: false}, nil
	}

	randomIndex := rand.Intn(len(Assurances.Data))
	ctx.Set(AssuranceTypeIndex, Assurances.Data[randomIndex].Index)
	ctx.Set(AssuranceTypeName, Assurances.Data[randomIndex].Name)
	ctx.Set(AssuranceTypePrice, Assurances.Data[randomIndex].Price)

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
		log.Errorf("CreateNewAssurance failed: %v", err)
		return nil, err
	}
	if addAssuranceResp.Msg == "Already exists" {
		log.Errorf("Order ID found, skip")
		return nil, err
	}
	if addAssuranceResp.Data.OrderId != TheOrderID {
		log.Errorf("Request failed, addAssuranceResp.Data.OrderId:%s, expected: %s", addAssuranceResp.Data.OrderId, TheOrderID)
		return nil, err
	}
	if addAssuranceResp.Data.Type != "TRAFFIC_ACCIDENT" {
		log.Errorf("Request failed, addAssuranceResp.Data.Type are expected to be 'TRAFFIC_ACCIDENT' but actually: %v", addAssuranceResp.Data.Type)
		return nil, err
	}

	ctx.Set(OrderId, addAssuranceResp.Data.OrderId)
	//ctx.Set(TypeIndex, addAssuranceResp.Data.)
	//ctx.Set(TypeName, Assurances.Data[randomIndex].TypeName)
	//ctx.Set(TypePrice, Assurances.Data[randomIndex].TypePrice)

	return nil, nil
}
