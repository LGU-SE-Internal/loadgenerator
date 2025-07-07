package behaviors

import (
	"fmt"
	"math/rand"

	"github.com/Lincyaw/loadgenerator/service"
	log "github.com/sirupsen/logrus"
)

func QueryAssurance(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(service.AssuranceService)
	if !ok {
		return nil, fmt.Errorf("assurance service client not found in context")
	}

	Assurances, err := cli.GetAllAssuranceTypes()
	if err != nil {
		log.WithError(err).Error("Failed to retrieve all assurance types from service")
		return nil, err
	}
	if Assurances.Status != 1 {
		log.WithFields(log.Fields{
			"status": Assurances.Status,
			"data":   Assurances.Data,
		}).Error("Assurance service returned non-success status, terminating execution")
		return &NodeResult{Continue: false}, nil
	}

	randomIndex := rand.Intn(len(Assurances.Data))
	selectedAssurance := Assurances.Data[randomIndex]

	log.WithFields(log.Fields{
		"assurance_index": selectedAssurance.Index,
		"assurance_name":  selectedAssurance.Name,
		"assurance_price": selectedAssurance.Price,
	}).Info("Successfully selected random assurance type")

	ctx.Set(AssuranceTypeIndex, selectedAssurance.Index)
	ctx.Set(AssuranceTypeName, selectedAssurance.Name)
	ctx.Set(AssuranceTypePrice, selectedAssurance.Price)

	return nil, nil
}

func CreateAssurance(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service implementation client not found in context")
	}

	//Create a new assurance
	TheOrderID := ctx.Get(OrderId).(string)
	log.WithFields(log.Fields{
		"order_id":   TheOrderID,
		"type_index": 1,
		"type_name":  "TRAFFIC_ACCIDENT",
	}).Info("Creating new assurance for order")

	addAssuranceResp, err := cli.CreateNewAssurance(1, TheOrderID) // typeIndex 1 -> TRAFFIC_ACCIDENT
	if err != nil {
		log.WithFields(log.Fields{
			"order_id": TheOrderID,
			"error":    err,
		}).Error("Failed to create new assurance")
		return nil, err
	}

	if addAssuranceResp.Msg == "Already exists" {
		log.WithField("order_id", TheOrderID).Warn("Assurance already exists for order, skipping creation")
		return nil, fmt.Errorf("assurance already exists for order %s", TheOrderID)
	}

	if addAssuranceResp.Data.OrderId != TheOrderID {
		log.WithFields(log.Fields{
			"expected_order_id": TheOrderID,
			"received_order_id": addAssuranceResp.Data.OrderId,
		}).Error("Order ID mismatch in assurance creation response")
		return nil, fmt.Errorf("order ID mismatch: expected %s, received %s", TheOrderID, addAssuranceResp.Data.OrderId)
	}

	if addAssuranceResp.Data.Type != "TRAFFIC_ACCIDENT" {
		log.WithFields(log.Fields{
			"expected_type": "TRAFFIC_ACCIDENT",
			"received_type": addAssuranceResp.Data.Type,
		}).Error("Assurance type mismatch in creation response")
		return nil, fmt.Errorf("assurance type mismatch: expected TRAFFIC_ACCIDENT, received %s", addAssuranceResp.Data.Type)
	}

	log.WithFields(log.Fields{
		"order_id":       addAssuranceResp.Data.OrderId,
		"assurance_type": addAssuranceResp.Data.Type,
	}).Info("Successfully created new assurance")

	ctx.Set(OrderId, addAssuranceResp.Data.OrderId)
	//ctx.Set(TypeIndex, addAssuranceResp.Data.)
	//ctx.Set(TypeName, Assurances.Data[randomIndex].TypeName)
	//ctx.Set(TypePrice, Assurances.Data[randomIndex].TypePrice)

	return nil, nil
}
