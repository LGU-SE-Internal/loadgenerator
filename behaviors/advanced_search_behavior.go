package behaviors

import (
	"fmt"
	"math/rand"

	"github.com/Lincyaw/loadgenerator/service"
	log "github.com/sirupsen/logrus"
)

const (
	TrainTypeId                   = "trainTypeId"
	StopStations                  = "stopStations" //[]string
	PriceForSecondClassSeat       = "priceForSecondClassSeat"
	NumberOfRestTicketSecondClass = "numberOfRestTicketSecondClass"
	PriceForFirstClassSeat        = "priceForFirstClassSeat"
	NumberOfRestTicketFirstClass  = "numberOfRestTicketFirstClass"
)

var AdvancedSearchChain *Chain

func init() {
	AdvancedSearchChain = NewChain(
		NewFuncNode(VerifyCode, "VerifyCode"),
		NewFuncNode(LoginBasic, "LoginBasic"),
		NewFuncNode(QueryUser, "QueryUser"),
		NewFuncNode(ChooseRoute, "ChooseRoute"),
		NewFuncNode(TravelPlanAdvancedSearch, "TravelPlanAdvancedSearch"),
		NewFuncNode(QuerySeatInfo, "QuerySeatInfo"),
		NewFuncNode(QueryContacts, "QueryContacts"),
		//NewFuncNode(CreateTrainFood, "CreateTrainFood"), //service not support
		NewFuncNode(QueryFood, "QueryFood"),
		NewFuncNode(QueryAssurance, "QueryAssurance"),
		NewFuncNode(Preserve, "Preserve"),
	)

	// fmt.Println(AdvancedSearchChain.VisualizeChain(0))
}

func TravelPlanAdvancedSearch(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	travelPlanInput := service.TravelQueryInfo{
		DepartureTime: ctx.Get(DepartureTime).(string),
		EndPlace:      ctx.Get(EndStation).(string),
		StartPlace:    ctx.Get(StartStation).(string),
	}

	var travelplanSvc service.TravelplanService = cli
	var Resp *service.TravelQueryArrResponse
	switch rand.Intn(3) {
	case 0:
		log.Infof("Querying cheapest tickets from %s to %s on %s", travelPlanInput.StartPlace, travelPlanInput.EndPlace, travelPlanInput.DepartureTime)
		travelPlanCheapestResp, err := travelplanSvc.ReqGetByCheapest(&travelPlanInput)
		if err != nil {
			return nil, err
		}
		if travelPlanCheapestResp.Status != 1 {
			return nil, fmt.Errorf("cheapest tickets query failed: status=%d, message=%s", travelPlanCheapestResp.Status, travelPlanCheapestResp.Msg)
		}
		if len(travelPlanCheapestResp.Data) == 0 {
			log.Warnf("No cheapest tickets found from %s to %s on %s, retrying later", travelPlanInput.StartPlace, travelPlanInput.EndPlace, travelPlanInput.DepartureTime)
			return &(NodeResult{false}), nil
		}
		log.Infof("Found %d cheapest ticket options, proceeding to reservation", len(travelPlanCheapestResp.Data))
		Resp = travelPlanCheapestResp
	case 1:
		log.Infof("Querying minimum station tickets from %s to %s on %s", travelPlanInput.StartPlace, travelPlanInput.EndPlace, travelPlanInput.DepartureTime)
		travelPlanMinimumStationNumberResp, err := travelplanSvc.ReqGetByMinStation(&travelPlanInput)
		if err != nil {
			return nil, err
		}
		if travelPlanMinimumStationNumberResp.Status != 1 {
			return nil, fmt.Errorf("minimum station tickets query failed: status=%d, message=%s", travelPlanMinimumStationNumberResp.Status, travelPlanMinimumStationNumberResp.Msg)
		}
		if len(travelPlanMinimumStationNumberResp.Data) == 0 {
			log.Warnf("No minimum station tickets found from %s to %s on %s, retrying later", travelPlanInput.StartPlace, travelPlanInput.EndPlace, travelPlanInput.DepartureTime)
			return &(NodeResult{false}), nil
		}
		log.Infof("Found %d minimum station ticket options, proceeding to reservation", len(travelPlanMinimumStationNumberResp.Data))
		Resp = travelPlanMinimumStationNumberResp
	case 2:
		log.Infof("Querying quickest tickets from %s to %s on %s", travelPlanInput.StartPlace, travelPlanInput.EndPlace, travelPlanInput.DepartureTime)
		travelPlanQuickestResp, err := travelplanSvc.ReqGetByQuickest(&travelPlanInput)
		if err != nil {
			return nil, err
		}
		if travelPlanQuickestResp.Status != 1 {
			return nil, fmt.Errorf("quickest tickets query failed: status=%d, message=%s", travelPlanQuickestResp.Status, travelPlanQuickestResp.Msg)
		}
		if len(travelPlanQuickestResp.Data) == 0 {
			log.Warnf("No quickest tickets found from %s to %s on %s, retrying later", travelPlanInput.StartPlace, travelPlanInput.EndPlace, travelPlanInput.DepartureTime)
			return &(NodeResult{false}), nil
		}
		log.Infof("Found %d quickest ticket options, proceeding to reservation", len(travelPlanQuickestResp.Data))
		Resp = travelPlanQuickestResp
	}

	randomIndex := rand.Intn(len(Resp.Data))
	selectedTrip := Resp.Data[randomIndex]

	log.Infof("Selected trip: %s, departure: %s, arrival: %s, economy seats: %d, first class seats: %d",
		selectedTrip.TripId, selectedTrip.StartTime, selectedTrip.EndTime,
		selectedTrip.NumberOfRestTicketSecondClass, selectedTrip.NumberOfRestTicketFirstClass)

	ctx.Set(TripID, selectedTrip.TripId)
	//ctx.Set(OldTripID, fmt.Sprintf("%s%s", selectedTrip.TripId.Type, selectedTrip.TripId.Number))
	ctx.Set(StartTime, selectedTrip.StartTime)
	ctx.Set(EndTime, selectedTrip.EndTime)
	ctx.Set(EconomyClass, selectedTrip.NumberOfRestTicketSecondClass)
	ctx.Set(ConfortClass, selectedTrip.NumberOfRestTicketFirstClass)
	ctx.Set(PriceForEconomyClass, selectedTrip.PriceForSecondClassSeat)
	ctx.Set(PriceForConfortClass, selectedTrip.PriceForFirstClassSeat)
	ctx.Set(TrainTypeName, selectedTrip.TrainTypeId)

	return nil, nil
}
