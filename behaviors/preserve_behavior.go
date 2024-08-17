package behaviors

import (
	"fmt"
	"github.com/Lincyaw/loadgenerator/service"
	log "github.com/sirupsen/logrus"
	"time"
)

const (
	// Preserve - Main
	AccountID  = "accountId"
	ContactsID = "contactsId"
	TripID     = "tripId"
	SeatType   = "seatType"
	//LoginToken = "loginToken"
	Date            = "date"
	From            = "from"
	To              = "to"
	Assurance       = "assurance"
	FoodType        = "foodType"
	StationName     = "stationName"
	StoreName       = "storeName"
	FoodName        = "foodName"
	FoodPrice       = "foodPrice"
	HandleDate      = "handleDate"
	ConsigneeName   = "consigneeName"
	ConsigneePhone  = "consigneePhone"
	ConsigneeWeight = "consigneeWeight"
	IsWithin        = "isWithin"

	// Assurance
	OrderId            = "orderId"
	AssuranceTypeIndex = "typeIndex"
	AssuranceTypeName  = "typeName"
	AssuranceTypePrice = "typePrice"

	// VerifyCode
	BooleanVerifyCode = "booleanVerifyCode"

	// User
	UserID       = "userId"
	UserName     = "userName"
	Password     = "password"
	Gender       = "gender"
	DocumentType = "documentType"
	DocumentNum  = "documentNum"
	Email        = "email"

	// Contacts
	//Id = "id" - ContactsID
	//AccountId = "accountId"
	Name = "name"
	//DocumentType   = "documentType"
	DocumentNumber = "documentNumber"
	PhoneNumber    = "phoneNumber"

	// Consign
	ID      = "id"
	OrderID = "orderId"
	//AccountID = "accountId"
	//HandleDate = "handleDate"
	TargetDate = "targetDate"
	//From = "from"
	//To = "to"
	Consignee = "consignee"
	Phone     = "phone"
	Weight    = "weight"
	//IsWithin = "isWithin"
	Price = "price"

	// FoodBehavior

	// Trip(Travel)
	TripId        = "tripId"
	TrainTypeName = "trainTypeName"
	//StartStation         = "startStation"
	TerminalStation = "terminalStation"
	StartTime       = "startTime"
	EndTime         = "endTime"
	//EconomyClass         = "economyClass"
	//ConfortClass         = "confortClass"
	PriceForEconomyClass = "priceForEconomyClass"
	PriceForConfortClass = "priceForConfortClass"

	// Train
	//Id           = "id" //Train-ID needed or not?
	//Name         = "name" //Train-Name needed or not?
	TrainTypName = "trainTypeName"
	ConfortClass = "confortClass"
	AverageSpeed = "averageSpeed"
	EconomyClass = "economyClass"

	// Route
	RouteID      = "routeId"
	Stations     = "stations"
	Distances    = "distances"
	StartStation = "startStation"
	EndStation   = "endStation"

	// Config
	ConfigName  = "name"
	Value       = "value"
	Description = "description"

	// Order
	//AccountId              = "accountId"
	BoughtDate             = "boughtDate"
	CoachNumber            = "coachNumber"
	ContactsDocumentNumber = "contactsDocumentNumber"
	//ContactsName           = "contactsName"
	DifferenceMoney = "differenceMoney"
	//DocumentType           = "documentType"
	//From                   = "from"
	//Id                     = "id"
	//Price                  = "price"
	SeatClass  = "seatClass"
	SeatNumber = "seatNumber"
	Status     = "status"
	//To                     = "to"
	TrainNumber = "trainNumber"
	TravelDate  = "travelDate"
	TravelTime  = "travelTime"

	TrainType           = "trainType"
	RouteId             = "routeId"
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
	SeatNo = "seatNo"
	//StartStation = "startStation"
	DestStation = "destStation"
)

var PreserveBehaviorChain *Chain

func init() {
	// ------------------------------------- init -------------------------------------------
	// Main Chain
	PreserveBehaviorChain = NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Printf("PreserveBehaviors(Chain) Statrs. Starts time: %v", time.Now().String())
		return nil, nil
	}, "Dummy"))
	// AssuranceBehaviorChain -
	AssuranceBehaviorChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Printf("CreateAssuranceChain. Starts time: %v", time.Now().String())
		return nil, nil
	}, "DummyAssurance"))
	// UserBehaviorsChain
	UserBehaviorsChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Printf("UserBehaviorsChain Starts. Starts time: %v", time.Now().String())
		return nil, nil
	}, "DummyUserBehaviors"))
	// VerifyCodeBehaviorChain
	VerifyCodeBehaviorChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Printf("VerifyCodeBehaviorChain Starts. Start time: %v", time.Now().String())
		return nil, nil
	}, "DummyVerifyCodeBehavior"))
	// AuthBehaviorChain
	AuthBehaviorChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Printf("AuthBehaviorChain Starts. Start time: %v", time.Now().String())
		return nil, nil
	}, "DummyAuthBehavior"))
	// UserBehaviorChain
	UserBehaviorChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Printf("UserBehaviorChain Starts. Starts time: %v", time.Now().String())
		return nil, nil
	}, "DummyUserBehavior"))
	// ContactsBehaviorChain
	ContactsBehaviorChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Printf("ContactsBehaviorChain Starts. Starts time: %v", time.Now().String())
		return nil, nil
	}, "DummyContactsBehavior"))
	// ConsignBehaviorsChain
	ConsignBehaviorsChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Printf("ConsignBehaviorsChain Starts. Starts time: %v", time.Now().String())
		return nil, nil
	}, "DummyConsignBehaviors"))
	//ConsignPriceBehaviorChain
	ConsignPriceBehaviorChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Printf("ConsignPriceBehaviorChain Starts. Starts time: %v", time.Now().String())
		return nil, nil
	}, "DummyConsignPriceBehavior"))
	// FoodBehaviorChain
	FoodBehaviorChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Printf("FoodBehaviorChain Starts. Starts time: %v", time.Now().String())
		return nil, nil
	}, "DummyFoodBehavior"))
	// TravelBehaviorChain
	TravelBehaviorChain1 := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Printf("TravelBehaviorChain Starts. Starts time: %v", time.Now().String())
		return nil, nil
	}, "DummyTravelBehavior"))
	TravelBehaviorChain2 := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Printf("TravelBehaviorChain Starts. Starts time: %v", time.Now().String())
		return nil, nil
	}, "DummyTravelBehavior"))
	// StationFoodBehaviorChain
	StationFoodBehaviorChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Printf("StationFoodBehaviorChain Starts. Starts time: %v", time.Now().String())
		return nil, nil
	}, "DummyStationFoodBehavior"))
	//TrainFoodBehaviorChain
	TrainFoodBehaviorChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Printf("TrainFoodBehaviorChain Starts. Starts time: %v", time.Now().String())
		return nil, nil
	}, "DummyTrainFoodBehavior"))

	/*	BasicBehaviorChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Printf("BasicBehaviorChain Starts. Start time: %v", time.Now().String())
		return nil, nil
	}, "DummyBasicBehavior"))*/
	// SeatBehaviorChain
	SeatBehaviorChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Printf("SeatBehaviorChain Starts. Start time: %v", time.Now().String())
		return nil, nil
	}, "DummySeatBehavior"))
	// StationBehaviorChain
	StationBehaviorChain0 := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Printf("StationBehaviorChain Starts. Start time: %v", time.Now().String())
		return nil, nil
	}, "DummyStationBehavior"))
	StationBehaviorChain1 := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Printf("StationBehaviorChain Starts. Start time: %v", time.Now().String())
		return nil, nil
	}, "DummyStationBehavior"))
	StationBehaviorChain2 := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Printf("StationBehaviorChain Starts. Start time: %v", time.Now().String())
		return nil, nil
	}, "DummyStationBehavior"))
	StationBehaviorChain3 := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Printf("StationBehaviorChain Starts. Start time: %v", time.Now().String())
		return nil, nil
	}, "DummyStationBehavior"))
	//PriceBehaviorChain
	PriceBehaviorChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Printf("PriceBehaviorChain Starts. Start time: %v", time.Now().String())
		return nil, nil
	}, "DummyPriceBehavior"))
	//ConfigBehaviorChain
	ConfigBehaviorChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Printf("ConfigBehaviorChain Starts. Start time: %v", time.Now().String())
		return nil, nil
	}, "DummyConfigBehavior"))
	//OrderBehaviorChain
	OrderBehaviorChain1 := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Printf("OrderBehaviorChain Starts. Start time: %v", time.Now().String())
		return nil, nil
	}, "DummyOrderBehavior"))
	OrderBehaviorChain2 := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Printf("OrderBehaviorChain Starts. Start time: %v", time.Now().String())
		return nil, nil
	}, "DummyOrderBehavior"))
	OrderBehaviorChain3 := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Printf("OrderBehaviorChain Starts. Start time: %v", time.Now().String())
		return nil, nil
	}, "DummyOrderBehavior"))
	//OrderOtherBehaviorChain
	OrderOtherBehaviorChain1 := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Printf("OrderOtherBehaviorChain Starts. Strat time: %v", time.Now().String())
		return nil, nil
	}, "DummyOrderOtherBehavior"))
	OrderOtherBehaviorChain2 := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Printf("OrderOtherBehaviorChain Starts. Strat time: %v", time.Now().String())
		return nil, nil
	}, "DummyOrderOtherBehavior"))
	// SecurityBehaviorChain
	SecurityBehaviorChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Printf("SecurityBehaviorChain Satrts. Start time: %v", time.Now().String())
		return nil, nil
	}, "DummySecurityBehavior"))

	// ------------------------------------- NewFuncNode -------------------------------------------
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

	//ContactsBehaviorChain - Contacts
	QueryContactsNode := NewFuncNode(QueryContacts, "QueryContacts")
	CreateContactsNode := NewFuncNode(CreateContacts, "CreateContacts")

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
	// TravelBehaviorChain
	QueryTravelNode := NewFuncNode(QueryTrip, "QueryTrip")
	//CreateTravelNode := NewFuncNode(CreateTrip, "CreateTrip")

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

	/*	// ******* Preserve ********
		PreserveNode := NewFuncNode(Preserve, "Preserve")*/

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
	QueryTravelChain1 := NewChain(QueryTravelNode)
	QueryTravelChain2 := NewChain(QueryTravelNode)
	//CreateTravelChain := NewChain(CreateTravelNode)

	// BasicBehaviorChain

	// TravelBehaviorChain
	TravelBehaviorChain := NewChain(NewFuncNode(QueryRoute, "QueryRoute"))
	TravelBehaviorChain.AddNode(NewFuncNode(QueryTrain, "QueryTrain"))
	TravelBehaviorChain.AddNode(NewFuncNode(QueryTripInfo, "QueryTripInfo"))
	TravelBehaviorChain.AddNode(NewFuncNode(QuerySeatInfo, "QuerySeat"))
	TravelBehaviorChain.AddNode(NewFuncNode(QueryContacts, "QueryContacts"))
	TravelBehaviorChain.AddNode(NewFuncNode(QueryTripId, "QueryContacts"))
	TravelBehaviorChain.AddNode(NewFuncNode(QueryFood, "QueryFood"))
	TravelBehaviorChain.AddNode(NewFuncNode(Preserve, "Preserve")) // END
	//QueryBasicChain.AddNode(NewFuncNode(QueryBasic, "QueryBasic"))
	//QueryBasicChain.AddNode(QueryPriceNode)

	TravelBehaviorChain1.AddNextChain(QueryTravelChain1, 1)
	TravelBehaviorChain2.AddNextChain(QueryTravelChain2, 1)

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

	TravelBehaviorChain.AddNextChain(TravelBehaviorChain, 1)

	//StationFoodBehaviorChain
	StationFoodBehaviorChain.AddNextChain(QueryStationFoodChain, 1)
	//TrainFoodBehaviorChain
	TrainFoodBehaviorChain.AddNextChain(QueryTrainFoodChain, 1)

	FoodBehaviorChain.AddNextChain(StationFoodBehaviorChain, 1)
	QueryStationFoodChain.AddNextChain(TrainFoodBehaviorChain, 1)
	QueryTrainFoodChain.AddNextChain(QueryFoodChain, 1)

	// &&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&& Main Chain &&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&
	// &&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&& Main Chain &&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&
	PreserveBehaviorChain.AddNextChain(UserBehaviorsChain, 1)
	// UserBehaviorsChain
	QueryUserChain.AddNextChain(TravelBehaviorChain, 1)

	// ------------------------------------- VisualizeChain -------------------------------------------
	// ------------------------------------- VisualizeChain -------------------------------------------
	fmt.Println(PreserveBehaviorChain.VisualizeChain(0))
	fmt.Println()
}

// Preserve Behaviors Function
func Preserve(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}
	OrderTicketsInfo := service.OrderTicketsInfo{
		AccountID:  "4d2a46c7-71cb-4cf1-b5bb-b68406d9da6f",
		ContactsID: "ffff5155-2d6d-43ea-a27c-da709097f22d",
		TripID:     "D1345",
		SeatType:   ctx.Get(SeatClass).(int),
		LoginToken: ctx.Get(LoginToken).(string),
		Date:       "2024-08-22",
		From:       "shanghai",
		To:         "suzhou",
		Assurance:  0,
		FoodType:   1,
		FoodName:   "Bone Soup",
		FoodPrice:  2.5,
		HandleDate: "2024-08-22",
	}
	PreserveResp, err := cli.Preserve(&OrderTicketsInfo)
	if err != nil {
		return nil, err
	}
	if PreserveResp.Status != 1 {
		return nil, fmt.Errorf("preserve order tickets fail. PreserveResp.Status != 1, get %v", PreserveResp.Status)
	}
	log.Errorf("The Status is: %v, and PreserveResp Data: %v\n", PreserveResp.Status, PreserveResp.Data)
	log.Errorf("PreserveBehaviors(Chain) Ends. End time: %v", time.Now().String())

	//return nil, nil
	return &(NodeResult{false}), nil // Chain End :D
}
