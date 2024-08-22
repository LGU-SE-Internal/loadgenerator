package behaviors

import (
	"fmt"
	"github.com/Lincyaw/loadgenerator/service"
	log "github.com/sirupsen/logrus"
	"time"
)

const (
	AccountID              = "accountId"
	ContactsID             = "contactsId"
	TripID                 = "tripId"
	From                   = "from"
	To                     = "to"
	FoodType               = "foodType"
	StationName            = "stationName"
	StoreName              = "storeName"
	FoodName               = "foodName"
	FoodPrice              = "foodPrice"
	HandleDate             = "handleDate"
	IsWithin               = "isWithin"
	OrderId                = "orderId"
	AssuranceTypeIndex     = "typeIndex"
	AssuranceTypeName      = "typeName"
	AssuranceTypePrice     = "typePrice"
	BooleanVerifyCode      = "booleanVerifyCode"
	UserName               = "userName"
	Password               = "password"
	Gender                 = "gender"
	DocumentType           = "documentType"
	DocumentNum            = "documentNum"
	Email                  = "email"
	Name                   = "name"
	DocumentNumber         = "documentNumber"
	PhoneNumber            = "phoneNumber"
	ID                     = "id"
	TargetDate             = "targetDate"
	Consignee              = "consignee"
	Phone                  = "phone"
	Weight                 = "weight"
	Price                  = "price"
	TrainTypeName          = "trainTypeName"
	StartTime              = "startTime"
	EndTime                = "endTime"
	PriceForEconomyClass   = "priceForEconomyClass"
	PriceForConfortClass   = "priceForConfortClass"
	ConfortClass           = "confortClass"
	AverageSpeed           = "averageSpeed"
	EconomyClass           = "economyClass"
	RouteID                = "routeId"
	Distances              = "distances"
	StartStation           = "startStation"
	EndStation             = "endStation"
	ConfigName             = "name"
	Value                  = "value"
	Description            = "description"
	BoughtDate             = "boughtDate"
	CoachNumber            = "coachNumber"
	ContactsDocumentNumber = "contactsDocumentNumber"
	DifferenceMoney        = "differenceMoney"
	SeatClass              = "seatClass"
	SeatNumber             = "seatNumber"
	Status                 = "status"
	TrainNumber            = "trainNumber"
	TravelDate             = "travelDate"
	TravelTime             = "travelTime"
	BasicPriceRate         = "basicPriceRate"
	FirstClassPriceRate    = "firstClassPriceRate"
	Percent                = "percent"
	Route                  = "route"
	Prices                 = "prices"
	SecurityID             = "id"
	SecurityName           = "name"
	SecurityValue          = "value"
	SecurityDescription    = "description"
	StationId              = "id"
	StationNames           = "name"
	StayTime               = "stayTime"
	DepartureTime          = "departureTime"
	SeatNo                 = "seatNo"
	DestStation            = "destStation"
)

var PreserveBehaviorChain *Chain
var NormalPreserveChain *Chain

func init() {
	NormalPreserveChain = NewChain(
		NewFuncNode(VerifyCode, "VerifyCode"),
		NewFuncNode(LoginBasic, "LoginBasic"),
		NewFuncNode(QueryUser, "QueryUser"),
		NewFuncNode(InputStartEndAndDate, "InputStartEndAndDate"),
		NewFuncNode(QueryTrain, "QueryTrain"),
		NewFuncNode(QueryRouteByStartAndEnd, "QueryRoute"),
		NewFuncNode(QueryTripInfo, "QueryTripInfo"),
		NewFuncNode(QuerySeatInfo, "QuerySeatInfo"),
		NewFuncNode(QueryContacts, "QueryContacts"),
		NewFuncNode(QueryFood, "QueryFood"),
		NewFuncNode(QueryAssurance, "QueryAssurance"),
		NewFuncNode(Preserve, "Preserve"),
	)
	fmt.Println(NormalPreserveChain.VisualizeChain(0))
}

// Preserve Behaviors Function

func Preserve(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}
	OrderTicketsInfo := service.OrderTicketsInfo{
		AccountID:  ctx.Get(AccountID).(string),
		ContactsID: ctx.Get(ContactsID).(string),
		TripID:     ctx.Get(TripID).(string),
		SeatType:   ctx.Get(SeatClass).(int),
		LoginToken: ctx.Get(LoginToken).(string),
		Date:       ctx.Get(DepartureTime).(string),
		From:       ctx.Get(StartStation).(string),
		To:         ctx.Get(EndStation).(string),
		Assurance:  ctx.Get(AssuranceTypeIndex).(int),
		FoodType:   ctx.Get(FoodType).(int),
		FoodName:   ctx.Get(FoodName).(string),
		FoodPrice:  ctx.Get(Price).(float64),
		HandleDate: ctx.Get(DepartureTime).(string),
	}

	TheTrainTypeName := ctx.Get(TrainTypeName).(string)

	switch TheTrainTypeName {
	case "GaoTieOne", "GaoTieTwo", "DongCheOne": // preserve
		var preserveSvc service.PreserveService = cli
		PreserveResp, err := preserveSvc.Preserve(&OrderTicketsInfo)
		if err != nil {
			return nil, err
		}
		if PreserveResp.Status != 1 {
			return nil, fmt.Errorf("preserve order tickets fail. PreserveResp.Status != 1, get %v", PreserveResp.Status)
		}
		log.Infof("The Status is: %v, and PreserveResp Data: %v", PreserveResp.Status, PreserveResp.Data)
		log.Infof("[Success]PreserveBehaviors(Chain) Finished. End time: %v", time.Now().String())

	default: //preserveOther
		var preserveOtherSvc service.PreserveOtherService = cli
		PreserveOtherResp, err := preserveOtherSvc.PreserveOther(&OrderTicketsInfo)
		if err != nil {
			return nil, err
		}
		if PreserveOtherResp.Status != 1 {
			return nil, fmt.Errorf("preserve other order tickets fail. PreserveResp.Status != 1, get %v", PreserveOtherResp.Status)
		}
		log.Infof("The Status is: %v, and PreserveResp Data: %v", PreserveOtherResp.Status, PreserveOtherResp.Data)
		log.Infof("[Success]PreserveBehaviors(Chain) Finished. End time: %v", time.Now().String())

	}

	return &(NodeResult{false}), nil // Chain End :D
}
