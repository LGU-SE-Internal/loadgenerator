package behaviors

import (
	"fmt"
	"github.com/Lincyaw/loadgenerator/service"
	log "github.com/sirupsen/logrus"
	"math/rand"
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

	fmt.Println(AdvancedSearchChain.VisualizeChain(0))
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
		travelPlanCheapestResp, err := travelplanSvc.ReqGetByCheapest(&travelPlanInput)
		if err != nil {
			return nil, err
		}
		if travelPlanCheapestResp.Status != 1 {
			return nil, fmt.Errorf("query Cheapest Tickets fail. travelPlanCheapestResp.Status != 1, get %v. The Resp Msg is: %v", travelPlanCheapestResp.Status, travelPlanCheapestResp.Msg)
		}
		if len(travelPlanCheapestResp.Data) == 0 {
			log.Warnf("[Please Try Again]Query Cheapest Ticket empty, No Data. Please Try Again. CheapestTicket Resp Status: %v", travelPlanCheapestResp.Status)
			return &(NodeResult{false}), nil
		}
		log.Infof("[Success]Search Cheapest Ticket Success. Go to ticket Reserve~. Resp Status: %v", travelPlanCheapestResp.Status)
		Resp = travelPlanCheapestResp
	case 1:
		travelPlanMinimumStationNumberResp, err := travelplanSvc.ReqGetByMinStation(&travelPlanInput)
		if err != nil {
			return nil, err
		}
		if travelPlanMinimumStationNumberResp.Status != 1 {
			return nil, fmt.Errorf("query Minimum Station Number Tickets fail. travelPlanMinimumStationNumberResp.Status != 1, get %v. The Resp Msg is: %v", travelPlanMinimumStationNumberResp.Status, travelPlanMinimumStationNumberResp.Msg)
		}
		if len(travelPlanMinimumStationNumberResp.Data) == 0 {
			log.Warnf("[Please Try Again]Query Minimum Station Number Ticket empty, No Data. Please Try Again. travelPlanMinimumStationNumber Resp Status: %v", travelPlanMinimumStationNumberResp.Status)
			return &(NodeResult{false}), nil
		}
		log.Infof("[Success]Search Minimum Station Number Ticket Success. Go to ticket Reserve~. Resp Status: %v", travelPlanMinimumStationNumberResp.Status)
		Resp = travelPlanMinimumStationNumberResp
	case 2:
		travelPlanQuickestResp, err := travelplanSvc.ReqGetByQuickest(&travelPlanInput)
		if err != nil {
			return nil, err
		}
		if travelPlanQuickestResp.Status != 1 {
			return nil, fmt.Errorf("query Quickest Tickets fail. travelPlanQuickestResp.Status != 1, get %v. The Resp Msg is: %v", travelPlanQuickestResp.Status, travelPlanQuickestResp.Msg)
		}
		if len(travelPlanQuickestResp.Data) == 0 {
			log.Warnf("[Please Try Again]Query Qucikest Ticket empty, No Data. Please Try Again. QuciketTicket Resp Status: %v", travelPlanQuickestResp.Status)
			return &(NodeResult{false}), nil
		}
		log.Infof("[Success]Search Quickest Ticket Success. Go to ticket Reserve~. Resp Status: %v", travelPlanQuickestResp.Status)
		Resp = travelPlanQuickestResp
	}

	randomIndex := rand.Intn(len(Resp.Data))
	ctx.Set(TripID, Resp.Data[randomIndex].TripId)
	//ctx.Set(OldTripID, fmt.Sprintf("%s%s", Resp.Data[randomIndex].TripId.Type, Resp.Data[randomIndex].TripId.Number))
	ctx.Set(StartTime, Resp.Data[randomIndex].StartTime)
	ctx.Set(EndTime, Resp.Data[randomIndex].EndTime)
	ctx.Set(EconomyClass, Resp.Data[randomIndex].NumberOfRestTicketSecondClass)
	ctx.Set(ConfortClass, Resp.Data[randomIndex].NumberOfRestTicketFirstClass)
	ctx.Set(PriceForEconomyClass, Resp.Data[randomIndex].PriceForSecondClassSeat)
	ctx.Set(PriceForConfortClass, Resp.Data[randomIndex].PriceForFirstClassSeat)
	ctx.Set(TrainTypeName, Resp.Data[randomIndex].TrainTypeId)

	return nil, nil
}
