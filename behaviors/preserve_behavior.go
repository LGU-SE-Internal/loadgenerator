package behaviors

import (
	"fmt"
	"github.com/Lincyaw/loadgenerator/service"
	log "github.com/sirupsen/logrus"
	"time"
)

const (
	// Preserve - Main
	AccountID   = "accountId"
	ContactsID  = "contactsId"
	TripID      = "tripId"
	From        = "from"
	To          = "to"
	FoodType    = "foodType"
	StationName = "stationName"
	StoreName   = "storeName"
	FoodName    = "foodName"
	FoodPrice   = "foodPrice"
	HandleDate  = "handleDate"
	IsWithin    = "isWithin"

	// Assurance
	OrderId            = "orderId"
	AssuranceTypeIndex = "typeIndex"
	AssuranceTypeName  = "typeName"
	AssuranceTypePrice = "typePrice"

	// VerifyCode
	BooleanVerifyCode = "booleanVerifyCode"

	// User
	UserName     = "userName"
	Password     = "password"
	Gender       = "gender"
	DocumentType = "documentType"
	DocumentNum  = "documentNum"
	Email        = "email"

	// Contacts
	Name           = "name"
	DocumentNumber = "documentNumber"
	PhoneNumber    = "phoneNumber"

	// Consign
	ID         = "id"
	TargetDate = "targetDate"
	Consignee  = "consignee"
	Phone      = "phone"
	Weight     = "weight"
	Price      = "price"

	// FoodBehavior

	// Trip(Travel)
	TrainTypeName        = "trainTypeName"
	StartTime            = "startTime"
	EndTime              = "endTime"
	PriceForEconomyClass = "priceForEconomyClass"
	PriceForConfortClass = "priceForConfortClass"

	// Train
	ConfortClass = "confortClass"
	AverageSpeed = "averageSpeed"
	EconomyClass = "economyClass"

	// Route
	RouteID      = "routeId"
	Distances    = "distances"
	StartStation = "startStation"
	EndStation   = "endStation"

	// Config
	ConfigName  = "name"
	Value       = "value"
	Description = "description"

	// Order
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

	BasicPriceRate      = "basicPriceRate"
	FirstClassPriceRate = "firstClassPriceRate"

	// Basic - QueryTraintype
	Percent = "percent"
	Route   = "route"
	Prices  = "prices"

	// Security
	SecurityID          = "id"
	SecurityName        = "name"
	SecurityValue       = "value"
	SecurityDescription = "description"

	// Station
	StationId    = "id"
	StationNames = "name"
	StayTime     = "stayTime"

	// Basic
	DepartureTime = "departureTime"

	// Seat
	SeatNo      = "seatNo"
	DestStation = "destStation"
)

var PreserveBehaviorChain *Chain

func init() {
	// ------------------------------------- init -------------------------------------------
	// ------------------------------------- init -------------------------------------------
	// Main Chain
	PreserveBehaviorChain = NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Infof("PreserveBehaviors(Chain) Starts. Starts time: %v", time.Now().String())
		return nil, nil
	}, "Dummy"))
	// AssuranceBehaviorChain -
	AssuranceBehaviorChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Infof("CreateAssuranceChain. Starts time: %v", time.Now().String())
		return nil, nil
	}, "DummyAssurance"))
	// UserBehaviorsChain
	UserBehaviorsChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Infof("UserBehaviorsChain Starts. Starts time: %v", time.Now().String())
		return nil, nil
	}, "DummyUserBehaviors"))
	// VerifyCodeBehaviorChain
	VerifyCodeBehaviorChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Infof("VerifyCodeBehaviorChain Starts. Start time: %v", time.Now().String())
		return nil, nil
	}, "DummyVerifyCodeBehavior"))
	// AuthBehaviorChain
	AuthBehaviorChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Infof("AuthBehaviorChain Starts. Start time: %v", time.Now().String())
		return nil, nil
	}, "DummyAuthBehavior"))
	// UserBehaviorChain
	UserBehaviorChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Infof("UserBehaviorChain Starts. Starts time: %v", time.Now().String())
		return nil, nil
	}, "DummyUserBehavior"))
	// ContactsBehaviorChain
	ContactsBehaviorChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Infof("ContactsBehaviorChain Starts. Starts time: %v", time.Now().String())
		return nil, nil
	}, "DummyContactsBehavior"))
	// ConsignBehaviorsChain
	ConsignBehaviorsChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Infof("ConsignBehaviorsChain Starts. Starts time: %v", time.Now().String())
		return nil, nil
	}, "DummyConsignBehaviors"))
	//ConsignPriceBehaviorChain
	ConsignPriceBehaviorChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Infof("ConsignPriceBehaviorChain Starts. Starts time: %v", time.Now().String())
		return nil, nil
	}, "DummyConsignPriceBehavior"))
	// FoodBehaviorChain
	FoodBehaviorChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Infof("FoodBehaviorChain Starts. Starts time: %v", time.Now().String())
		return nil, nil
	}, "DummyFoodBehavior"))
	// TravelBehaviorChain
	TravelBehaviorChain1 := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Infof("TravelBehaviorChain Starts. Starts time: %v", time.Now().String())
		return nil, nil
	}, "DummyTravelBehavior"))
	TravelBehaviorChain2 := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Infof("TravelBehaviorChain Starts. Starts time: %v", time.Now().String())
		return nil, nil
	}, "DummyTravelBehavior"))
	// StationFoodBehaviorChain
	StationFoodBehaviorChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Infof("StationFoodBehaviorChain Starts. Starts time: %v", time.Now().String())
		return nil, nil
	}, "DummyStationFoodBehavior"))
	//TrainFoodBehaviorChain
	TrainFoodBehaviorChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Infof("TrainFoodBehaviorChain Starts. Starts time: %v", time.Now().String())
		return nil, nil
	}, "DummyTrainFoodBehavior"))

	/*	BasicBehaviorChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Infof("BasicBehaviorChain Starts. Start time: %v", time.Now().String())
		return nil, nil
	}, "DummyBasicBehavior"))*/
	// SeatBehaviorChain
	SeatBehaviorChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Infof("SeatBehaviorChain Starts. Start time: %v", time.Now().String())
		return nil, nil
	}, "DummySeatBehavior"))
	// StationBehaviorChain
	StationBehaviorChain0 := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Infof("StationBehaviorChain Starts. Start time: %v", time.Now().String())
		return nil, nil
	}, "DummyStationBehavior"))
	StationBehaviorChain1 := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Infof("StationBehaviorChain Starts. Start time: %v", time.Now().String())
		return nil, nil
	}, "DummyStationBehavior"))
	StationBehaviorChain2 := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Infof("StationBehaviorChain Starts. Start time: %v", time.Now().String())
		return nil, nil
	}, "DummyStationBehavior"))
	StationBehaviorChain3 := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Infof("StationBehaviorChain Starts. Start time: %v", time.Now().String())
		return nil, nil
	}, "DummyStationBehavior"))
	//PriceBehaviorChain
	PriceBehaviorChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Infof("PriceBehaviorChain Starts. Start time: %v", time.Now().String())
		return nil, nil
	}, "DummyPriceBehavior"))
	//ConfigBehaviorChain
	ConfigBehaviorChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Infof("ConfigBehaviorChain Starts. Start time: %v", time.Now().String())
		return nil, nil
	}, "DummyConfigBehavior"))
	//OrderBehaviorChain
	OrderBehaviorChain1 := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Infof("OrderBehaviorChain Starts. Start time: %v", time.Now().String())
		return nil, nil
	}, "DummyOrderBehavior"))
	OrderBehaviorChain2 := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Infof("OrderBehaviorChain Starts. Start time: %v", time.Now().String())
		return nil, nil
	}, "DummyOrderBehavior"))
	OrderBehaviorChain3 := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Infof("OrderBehaviorChain Starts. Start time: %v", time.Now().String())
		return nil, nil
	}, "DummyOrderBehavior"))
	//OrderOtherBehaviorChain
	OrderOtherBehaviorChain1 := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Infof("OrderOtherBehaviorChain Starts. Strat time: %v", time.Now().String())
		return nil, nil
	}, "DummyOrderOtherBehavior"))
	OrderOtherBehaviorChain2 := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Infof("OrderOtherBehaviorChain Starts. Strat time: %v", time.Now().String())
		return nil, nil
	}, "DummyOrderOtherBehavior"))
	// SecurityBehaviorChain
	SecurityBehaviorChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Infof("SecurityBehaviorChain Satrts. Start time: %v", time.Now().String())
		return nil, nil
	}, "DummySecurityBehavior"))

	// ------------------------------------- NewFuncNode -------------------------------------------
	// ------------------------------------- NewFuncNode -------------------------------------------
	//AssuranceBehaviorChain - Assurance
	QueryAssuranceNode := NewFuncNode(QueryAssurance, "QueryAssurance")

	//UserBehaviorsChain
	LoginBasicNode := NewFuncNode(LoginBasic, "LoginBasic")
	//	VerifyCodeBehaviorChain
	VerifyCodeNode := NewFuncNode(VerifyCode, "VerifyCode")
	// UserBehaviorChain
	QueryUserNode := NewFuncNode(QueryUser, "QueryUser")

	InputStartEndAndDateNode := NewFuncNode(InputStartEndAndDate, "InputStartEndAndDate")

	// QueryRouteNode
	QueryRouteByStartAndEndNode := NewFuncNode(QueryRouteByStartAndEnd, "QueryRoute")

	// QueryTrainNode
	QueryTrainNode := NewFuncNode(QueryTrain, "QueryTrain")

	/*	// QueryTrainNode
		QueryTrainNode := NewFuncNode(QueryTrain, "QueryTrain")*/

	// TravelBehaviorChain & Travel2BehaviorChain
	//QueryTripInfoNode
	QueryTripInfoNode := NewFuncNode(QueryTripInfo, "QueryTripInfo")
	//CreateTripNode
	//CreateTripNode := NewFuncNode(CreateTrip, "CreateTrip")

	//QuerySeatInfoNode
	QuerySeatInfoNode := NewFuncNode(QuerySeatInfo, "QuerySeatInfo")

	//ContactsBehaviorChain - Contacts
	QueryContactsNode := NewFuncNode(QueryContacts, "QueryContacts")
	CreateContactsNode := NewFuncNode(CreateContacts, "CreateContacts")

	/*	// QueryTripIdNode
		QueryTripIdNode := NewFuncNode(QueryTripId, "QueryTripId")*/

	//ConsignBehaviorsChain
	//QueryConsignNode := NewFuncNode(QueryConsign, "QueryConsign")
	CreateConsignNode := NewFuncNode(CreateConsign, "CreateConsign")
	// ConsignPriceBehaviorChain
	//QueryConsignPriceNode := NewFuncNode(QueryConsignPric, "QueryConsignPrice")
	//CreateConsignPriceNode := NewFuncNode(CreateConsignPrice, "CreateConsignPrice")

	//FoodBehaviorChain
	QueryFoodNode := NewFuncNode(QueryFood, "QueryFood")

	// StationFoodBehaviorChain
	QueryStationFoodNode := NewFuncNode(QueryStationFood, "QueryStationFood")

	// TrainFoodBehaviorChain
	QueryTrainFoodNode := NewFuncNode(QueryTrainFood, "QueryTrainFood")

	CreateSeatNode := NewFuncNode(CreateSeat, "QuerySeat")

	//BasicBehaviorChain
	// StationBehaviorChain
	QueryStationNode := NewFuncNode(QueryStation, "QueryStation")
	// PriceBehaviorChain
	QueryPriceNode := NewFuncNode(QueryPrice, "QueryPrice")
	// RouteBehaviorChain
	// TrainBehaviorChain

	//SeatBehaviorChain
	// ConfigBehaviorChain
	QueryConfigNode := NewFuncNode(QueryConfig, "QueryConfig")
	// OrderBehaviorChain
	QueryOrderNode := NewFuncNode(QueryOrder, "QueryOrder")
	// OrderOtherBehaviorChain
	QueryOrderOtherNode := NewFuncNode(QueryOrderOther, "QueryOrderOther")

	QuerySecurityNode := NewFuncNode(QuerySecurity, "QuerySecurity")

	// ******* Preserve ********
	PreserveNode := NewFuncNode(Preserve, "Preserve") // END

	// ------------------------------------- NewChain -------------------------------------------
	// ------------------------------------- NewChain -------------------------------------------
	// AssuranceBehaviorChain - Assurance
	QueryAssuranceChain := NewChain(QueryAssuranceNode)
	//CreateAssuranceChain := NewChain(CreateAssuranceNode)

	// UserBehaviorsChain
	// AuthBehaviorChain - LoginAdmin/LoginBasic
	LoginBasicChain := NewChain(LoginBasicNode)
	// VerifyCodeBehaviorChain
	VerifyCodeChain := NewChain(VerifyCodeNode)
	// UserBehaviorChain
	QueryUserChain := NewChain(QueryUserNode)

	// ContactsBehaviorChain - Contacts
	QueryContactsChain := NewChain(QueryContactsNode)
	CreateContactsChain := NewChain(CreateContactsNode)

	// ConsignBehaviorsChain
	//QueryConsignChain := NewChain(QueryConsignNode)
	CreateConsignChain := NewChain(CreateConsignNode)
	// ConsignPriceBehaviorChain
	// FoodBehaviorChain
	QueryFoodChain := NewChain(QueryFoodNode)
	// StationFoodBehaviorChain
	QueryStationFoodChain := NewChain(QueryStationFoodNode)
	// TrainFoodBehaviorChain
	QueryTrainFoodChain := NewChain(QueryTrainFoodNode)
	// TravelBehaviorChain
	QueryTripChain1 := NewChain(QueryTripInfoNode)
	QueryTripChain2 := NewChain(QueryTripInfoNode)
	//CreateTravelChain := NewChain(CreateTravelNode)

	// BasicBehaviorChain

	// TravelBehaviorChain
	//TravelBehaviorChain := NewChain(NewFuncNode(QueryRoute, "QueryRoute"))
	TravelBehaviorChain := NewChain(InputStartEndAndDateNode)
	TravelBehaviorChain.AddNode(QueryTrainNode)
	TravelBehaviorChain.AddNode(QueryRouteByStartAndEndNode)
	TravelBehaviorChain.AddNode(QueryTripInfoNode)
	TravelBehaviorChain.AddNode(QuerySeatInfoNode)
	TravelBehaviorChain.AddNode(QueryContactsNode)
	TravelBehaviorChain.AddNode(QueryFoodNode)
	TravelBehaviorChain.AddNode(QueryAssuranceNode)
	TravelBehaviorChain.AddNode(PreserveNode) // END
	//QueryBasicChain.AddNode(NewFuncNode(QueryBasic, "QueryBasic"))
	//QueryBasicChain.AddNode(QueryPriceNode)

	TravelBehaviorChain1.AddNextChain(QueryTripChain1, 1)
	TravelBehaviorChain2.AddNextChain(QueryTripChain2, 1)

	// SeatBehaviorChain
	CreateSeatChain := NewChain(QueryConfigNode)
	CreateSeatChain.AddNode(QueryOrderNode)
	CreateSeatChain.AddNode(QueryOrderOtherNode)
	CreateSeatChain.AddNode(CreateSeatNode)

	// StationBehaviorChain
	QueryStationChain0 := NewChain(QueryStationNode)
	QueryStationChain1 := NewChain(QueryStationNode)
	QueryStationChain2 := NewChain(QueryStationNode)
	QueryStationChain3 := NewChain(QueryStationNode)
	// PriceBehaviorChain
	QueryPriceChain := NewChain(QueryPriceNode)

	// ConfigBehaviorChain
	QueryConfigChain := NewChain(QueryConfigNode)
	// OrderBehaviorChain
	QueryOrderChain1 := NewChain(QueryOrderNode)
	QueryOrderChain2 := NewChain(QueryOrderNode)
	QueryOrderChain3 := NewChain(QueryOrderNode)
	// OrderOtherBehaviorChain
	QueryOrderOtherChain1 := NewChain(QueryOrderOtherNode)
	QueryOrderOtherChain2 := NewChain(QueryOrderOtherNode)

	// SecurityBehaviorChain
	QuerySecurityChain := NewChain(QuerySecurityNode)

	/*	// The Last Chain - Preserve Behavior Chain
		PreserveChain := NewChain(PreserveNode)*/

	//AssuranceBehaviorChain
	AssuranceBehaviorChain.AddNextChain(QueryAssuranceChain, 1)
	//VerifyCodeBehaviorChain
	VerifyCodeBehaviorChain.AddNextChain(VerifyCodeChain, 1)
	//AuthBehaviorChain
	AuthBehaviorChain.AddNextChain(LoginBasicChain, 1)
	//UserBehaviorChain
	UserBehaviorChain.AddNextChain(QueryUserChain, 1)
	//UserBehaviorsChain
	UserBehaviorsChain.AddNextChain(VerifyCodeBehaviorChain, 1)
	VerifyCodeChain.AddNextChain(AuthBehaviorChain, 1)
	LoginBasicChain.AddNextChain(UserBehaviorChain, 1)
	//ContactsBehaviorChain
	ContactsBehaviorChain.AddNextChain(QueryContactsChain, 0.7)
	ContactsBehaviorChain.AddNextChain(CreateContactsChain, 0.3)
	//ConsignPriceBehaviorChain
	//ConsignBehaviorsChain
	ConsignBehaviorsChain.AddNextChain(ConsignPriceBehaviorChain, 1)
	ConsignPriceBehaviorChain.AddNextChain(CreateConsignChain, 1)
	//StationBehaviorChain
	StationBehaviorChain0.AddNextChain(QueryStationChain0, 1)
	StationBehaviorChain1.AddNextChain(QueryStationChain1, 1)
	StationBehaviorChain2.AddNextChain(QueryStationChain2, 1)
	StationBehaviorChain3.AddNextChain(QueryStationChain3, 1)
	//OrderOtherBehaviorChain
	OrderOtherBehaviorChain1.AddNextChain(QueryOrderOtherChain1, 1)
	OrderOtherBehaviorChain2.AddNextChain(QueryOrderOtherChain2, 1)
	//OrderBehaviorChain
	OrderBehaviorChain1.AddNextChain(QueryOrderChain1, 1)
	OrderBehaviorChain2.AddNextChain(QueryOrderChain2, 1)
	OrderBehaviorChain3.AddNextChain(QueryOrderChain3, 1)
	//SecurityBehaviorChain
	SecurityBehaviorChain.AddNextChain(QuerySecurityChain, 1)
	//ConfigBehaviorChain
	ConfigBehaviorChain.AddNextChain(QueryConfigChain, 1)
	//SeatBehaviorChain
	SeatBehaviorChain.AddNextChain(OrderBehaviorChain2, 1)
	QueryOrderChain2.AddNextChain(OrderOtherBehaviorChain2, 1)
	QueryOrderOtherChain2.AddNextChain(CreateSeatChain, 1)
	//PriceBehaviorChain
	PriceBehaviorChain.AddNextChain(QueryPriceChain, 1)

	//StationFoodBehaviorChain
	StationFoodBehaviorChain.AddNextChain(QueryStationFoodChain, 1)
	//TrainFoodBehaviorChain
	TrainFoodBehaviorChain.AddNextChain(QueryTrainFoodChain, 1)

	FoodBehaviorChain.AddNextChain(StationFoodBehaviorChain, 1)
	QueryStationFoodChain.AddNextChain(TrainFoodBehaviorChain, 1)
	QueryTrainFoodChain.AddNextChain(QueryFoodChain, 1)

	// &&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&& Main Chain &&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&
	// &&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&& Main Chain &&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&
	// 0. UserBehaviorChain
	PreserveBehaviorChain.AddNextChain(UserBehaviorsChain, 1)
	// TravelBehaviorChain
	QueryUserChain.AddNextChain(TravelBehaviorChain, 1)

	// ------------------------------------- VisualizeChain -------------------------------------------
	// ------------------------------------- VisualizeChain -------------------------------------------
	//fmt.Println(PreserveBehaviorChain.VisualizeChain(0))
	fmt.Println()
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

	//return nil, nil
	return &(NodeResult{false}), nil // Chain End :D
}
