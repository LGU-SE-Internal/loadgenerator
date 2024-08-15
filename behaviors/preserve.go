package behaviors

import (
	"errors"
	"fmt"
	"github.com/Lincyaw/loadgenerator/service"
	"github.com/go-faker/faker/v4"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"strings"
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

	BasicBehaviorChain := NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		log.Printf("BasicBehaviorChain Starts. Start time: %v", time.Now().String())
		return nil, nil
	}, "DummyBasicBehavior"))
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

	// ******* Preserve ********
	PreserveNode := NewFuncNode(Preserve, "Preserve")

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
	QueryBasicChain := NewChain(NewFuncNode(QueryRoute, "QueryRoute"))
	QueryBasicChain.AddNode(NewFuncNode(QueryTrain, "QueryTrain"))
	QueryBasicChain.AddNode(NewFuncNode(QueryTripInfo, "QueryTripInfo"))
	QueryBasicChain.AddNode(NewFuncNode(QuerySeatInfo, "QuerySeat"))
	QueryBasicChain.AddNode(NewFuncNode(QueryContacts, "QueryContacts"))
	QueryBasicChain.AddNode(NewFuncNode(QueryTripId, "QueryContacts"))
	QueryBasicChain.AddNode(NewFuncNode(QueryFood, "QueryFood"))
	QueryBasicChain.AddNode(NewFuncNode(Preserve, "Preserve"))
	//QueryBasicChain.AddNode(NewFuncNode(QueryBasic, "QueryBasic"))
	QueryBasicChain.AddNode(QueryPriceNode)
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

	// The Last Chain - Preserve Behavior Chain
	PreserveChain := NewChain(PreserveNode)

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

	BasicBehaviorChain.AddNextChain(QueryBasicChain, 1)
	//TravelBehaviorChain
	TravelBehaviorChain1.AddNextChain(QueryTravelChain1, 1)
	TravelBehaviorChain2.AddNextChain(QueryTravelChain2, 1)
	//StationFoodBehaviorChain
	StationFoodBehaviorChain.AddNextChain(QueryStationFoodChain, 1)
	//TrainFoodBehaviorChain
	TrainFoodBehaviorChain.AddNextChain(QueryTrainFoodChain, 1)

	FoodBehaviorChain.AddNextChain(StationFoodBehaviorChain, 1)
	QueryStationFoodChain.AddNextChain(TrainFoodBehaviorChain, 1)
	QueryTrainFoodChain.AddNextChain(QueryFoodChain, 1)

	// &&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&& Main Chain &&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&&
	PreserveBehaviorChain.AddNextChain(UserBehaviorsChain, 1)
	// UserBehaviorsChain
	QueryUserChain.AddNextChain(BasicBehaviorChain, 1)
	// BasicBehaviorChain
	QueryBasicChain.AddNextChain(SeatBehaviorChain, 1)
	// SeatBehaviorChain
	CreateSeatChain.AddNextChain(ContactsBehaviorChain, 1)
	// TravelBehaviorChain
	QueryContactsChain.AddNextChain(TravelBehaviorChain1, 1)
	CreateContactsChain.AddNextChain(TravelBehaviorChain1, 1)
	// ContactsBehaviorChain
	//QueryTravelChain1.AddNextChain(AssuranceBehaviorChain, 1)
	// AssuranceBehaviorChain
	QueryAssuranceChain.AddNextChain(FoodBehaviorChain, 1)
	// FoodBehaviorChain
	QueryFoodChain.AddNextChain(ConsignBehaviorsChain, 1)
	// ConsignBehaviorsChain
	CreateConsignChain.AddNextChain(PreserveChain, 1)

	// ------------------------------------- VisualizeChain -------------------------------------------
	// ------------------------------------- VisualizeChain -------------------------------------------
	fmt.Println(PreserveBehaviorChain.VisualizeChain(0))
	fmt.Println()
}

// ************************************* NewFuncNode_Function *******************************************

// AssuranceBehaviorChain
func QueryAssurance(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	Assurances, err := cli.GetAllAssurances()
	if err != nil {
		log.Errorf("GetAllAssurances failed: %v", err)
		return nil, err
	}
	if Assurances.Status != 1 {
		log.Errorf("Assurances status is not 1: %+v", Assurances)
		return nil, nil
	}

	randomIndex := rand.Intn(len(Assurances.Data))
	ctx.Set(OrderId, Assurances.Data[randomIndex].OrderId)
	ctx.Set(AssuranceTypeIndex, Assurances.Data[randomIndex].TypeIndex)
	ctx.Set(AssuranceTypeName, Assurances.Data[randomIndex].TypeName)
	ctx.Set(AssuranceTypePrice, Assurances.Data[randomIndex].TypePrice)

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

// VerifyCodeBehaviorChain
func VerifyCode(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	verifyCode := generateVerifyCode()
	verifyCodeResp, err := cli.VerifyCode(verifyCode)
	if err != nil {
		log.Errorf("Request failed, err %s", err)
		return nil, err
	}
	if !verifyCodeResp {
		log.Errorf("Verification failed")
		return nil, err
	}
	//log.Errorf("Verification code verified. The result is %v and verifyCode: %v", verifyCodeResp, verifyCode)

	ctx.Set(BooleanVerifyCode, verifyCodeResp)

	return nil, nil
}

func QueryUser(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(service.UserService)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	allUsersResp, err := cli.GetUserByUserId(ctx.Get(UserId).(string))
	if err != nil {
		log.Errorf("Request failed, err1 %s", err)
		return nil, err
	}
	if allUsersResp.Status != 1 {
		log.Errorf("Expected status 200, got %d", allUsersResp.Status)
		return nil, err
	}

	ctx.Set(UserName, allUsersResp.Data.UserName)
	ctx.Set(Password, allUsersResp.Data.Password)
	ctx.Set(Gender, allUsersResp.Data.Gender)
	ctx.Set(DocumentNum, allUsersResp.Data.DocumentNum)
	ctx.Set(DocumentType, allUsersResp.Data.DocumentType)
	ctx.Set(Email, allUsersResp.Data.Email)

	return nil, nil
}

// ContactsBehaviorChain
func QueryContacts(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	var contactsSvc service.ContactsService = cli
	TheAccountId := ctx.Get(UserId).(string)
	GetAllContacts, err := contactsSvc.GetContactByAccountId(TheAccountId)
	if err != nil {
		log.Errorf("[Mock AccountID]GetAllContacts fail. The error occurs: %v", err)
		return nil, err
	}
	if GetAllContacts.Status != 1 {
		log.Errorf("[Mock AccountID]GetAllContacts.Status != 1")
		return nil, err
	}

	randomIndex := rand.Intn(len(GetAllContacts.Data))
	ctx.Set(AccountID, GetAllContacts.Data[randomIndex].AccountId)
	ctx.Set(ContactsID, GetAllContacts.Data[randomIndex].Id)
	ctx.Set(Name, GetAllContacts.Data[randomIndex].Name)
	ctx.Set(DocumentType, GetAllContacts.Data[randomIndex].DocumentType)
	ctx.Set(DocumentNumber, GetAllContacts.Data[randomIndex].DocumentNumber)
	ctx.Set(PhoneNumber, GetAllContacts.Data[randomIndex].PhoneNumber)

	return nil, nil
}

func CreateContacts(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	CreateContactsInput := service.AdminContacts{
		Id:             faker.UUIDHyphenated(),
		AccountId:      ctx.Get(UserId).(string),
		Name:           faker.Name(),
		DocumentType:   rand.Intn(1),
		DocumentNumber: generateDocumentNumber(),
		PhoneNumber:    faker.PhoneNumber,
	}
	CreateContacts, err := cli.AddContact(&CreateContactsInput)
	if err != nil {
		log.Errorf("[Mock AccountID] CreateContacts error occurs: %v", err)
		return nil, err
	}
	if CreateContacts.Status != 1 {
		log.Errorf("[Mock AccountID] CreateContacts.Status != 1, resp: %+v", CreateContacts)
		return nil, err
	}

	ctx.Set(AccountID, CreateContacts.Data.AccountId)
	ctx.Set(ContactsID, CreateContacts.Data.Id)
	ctx.Set(Name, CreateContacts.Data.Name)
	ctx.Set(DocumentType, CreateContacts.Data.DocumentType)
	ctx.Set(DocumentNumber, CreateContacts.Data.DocumentNumber)
	ctx.Set(PhoneNumber, CreateContacts.Data.PhoneNumber)

	return nil, nil
}

// ConsignBehaviorsChain
func QueryConsign(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	// QueryTraintype consign records by order ID
	TheOrderId := ctx.Get(OrderId).(string)
	consignsByOrderId, err := cli.QueryByOrderId(TheOrderId)
	if err != nil {
		log.Errorf("QueryByOrderId failed: %v", err)
		return nil, err
	}
	if consignsByOrderId.Status != 1 {
		log.Errorf("consignsByOrderId.Status = 1")
		return nil, err
	}

	ctx.Set(ID, consignsByOrderId.Data.Id)
	ctx.Set(OrderId, consignsByOrderId.Data.OrderId)
	ctx.Set(AccountID, consignsByOrderId.Data.AccountId)
	ctx.Set(HandleDate, consignsByOrderId.Data.HandleDate)
	ctx.Set(TargetDate, consignsByOrderId.Data.TargetDate)
	ctx.Set(From, consignsByOrderId.Data.From)
	ctx.Set(To, consignsByOrderId.Data.To)
	ctx.Set(Consignee, consignsByOrderId.Data.Consignee)
	ctx.Set(Phone, consignsByOrderId.Data.Phone)
	ctx.Set(Weight, consignsByOrderId.Data.Weight)
	ctx.Set(Price, consignsByOrderId.Data.Price)

	return nil, nil
}

func CreateConsign(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	// Mock data
	MockedId := faker.UUIDHyphenated()
	MockedAccountId := ctx.Get(AccountID).(string)
	MockedOrderId := ctx.Get(OrderId).(string)
	MockedHandleDate := ctx.Get(DepartureTime).(string)
	//MockedHandleDate := ""
	MockedTargetDate := ctx.Get(DepartureTime).(string)
	//MockedTargetDate := ""
	MockedFromPlace := ctx.Get(StartStation).(string)
	MockedToPlace := ctx.Get(EndStation).(string)
	MockedConsignee := faker.Name()
	MockedPhone := faker.Phonenumber()
	MockedWeight := GenerateWeight()

	// Insert a new consign record
	insertReq := service.Consign{
		ID:         MockedId,
		OrderID:    MockedOrderId,
		AccountID:  MockedAccountId,
		HandleDate: MockedHandleDate,
		TargetDate: MockedTargetDate,
		From:       MockedFromPlace,
		To:         MockedToPlace,
		Consignee:  MockedConsignee,
		Phone:      MockedPhone,
		Weight:     MockedWeight,
		IsWithin:   BooleanIsWithin(MockedWeight),
	}
	insertResp, err := cli.InsertConsignRecord(&insertReq)
	if err != nil {
		log.Errorf("InsertConsignRecord failed: %v", err)
		return nil, err
	}
	if insertResp.Msg == "Already exists" {
		return nil, fmt.Errorf("Consign already exists")
	}
	if insertResp.Status != 1 {
		log.Errorf("InsertConsignRecord failed: %v", insertResp.Status)
		return nil, err
	}
	isMatch := false
	if /*insertResp.Data.ID == insertReq.ID &&*/
	/*insertResp.Data.IsWithin == insertReq.IsWithin &&*/
	insertResp.Data.AccountID == insertReq.AccountID &&
		insertResp.Data.From == insertReq.From &&
		insertResp.Data.Consignee == insertReq.Consignee &&
		insertResp.Data.OrderID == insertReq.OrderID &&
		insertResp.Data.Phone == insertReq.Phone &&
		insertResp.Data.TargetDate == insertReq.TargetDate &&
		insertResp.Data.HandleDate == insertReq.HandleDate &&
		insertResp.Data.To == insertReq.To &&
		insertResp.Data.Weight == insertReq.Weight {
		isMatch = true
	}
	if !isMatch {
		log.Errorf("Creation not match. Expect: %v, but get: %v", insertReq, insertResp.Data)
		return nil, err
	}
	//log.Errorf("InsertConsignRecord response: %+v", insertResp)
	//existedConsign := insertResp.Data

	ctx.Set(ID, insertResp.Data.ID)
	//ctx.Set(OrderID, insertResp.Data.OrderID)
	//ctx.Set(AccountID, insertResp.Data.AccountID)
	ctx.Set(HandleDate, insertResp.Data.HandleDate)
	ctx.Set(TargetDate, insertResp.Data.TargetDate)
	ctx.Set(From, insertResp.Data.From)
	ctx.Set(To, insertResp.Data.To)
	ctx.Set(Consignee, insertResp.Data.Consignee)
	ctx.Set(Phone, insertResp.Data.Phone)
	ctx.Set(Weight, insertResp.Data.Weight)
	ctx.Set(IsWithin, insertResp.Data.IsWithin)

	return nil, nil
}

func QueryConsignPric(ctx *Context) (*NodeResult, error) {
	_, ok := ctx.Get(Client).(*service.SvcImpl)
	//cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	// TODO part

	return nil, nil
}

func CreateConsignPrice(ctx *Context) (*NodeResult, error) {
	_, ok := ctx.Get(Client).(*service.SvcImpl)
	//cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")

	}

	// TODO part

	return nil, nil
}

// FoodBehaviorChain
func QueryFood(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	// QueryTraintype all
	allFoodOrders, err := cli.FindAllFoodOrder()
	if err != nil {
		log.Errorf("FindAllFoodOrder request failed, err %s", err)
		return nil, err
	}
	if len(allFoodOrders.Data) == 0 {
		log.Errorf("FindAllFoodOrder returned empty results")
		return nil, err
	}
	if allFoodOrders.Status != 1 {
		log.Errorf("FindAllFoodOrder failed: %v", allFoodOrders.Status)
		return nil, err
	}

	randomIndex := rand.Intn(len(allFoodOrders.Data))
	//ctx.Set(OrderId, allFoodOrders.Data[randomIndex].OrderId)
	ctx.Set(FoodType, allFoodOrders.Data[randomIndex].FoodType)
	//ctx.Set(StationName, allFoodOrders.Data[randomIndex].StationName)
	ctx.Set(StoreName, allFoodOrders.Data[randomIndex].StoreName)
	ctx.Set(FoodName, allFoodOrders.Data[randomIndex].FoodName)
	ctx.Set(Price, allFoodOrders.Data[randomIndex].Price)

	return nil, nil
}

func CreateFood(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	// Mock data
	MockedOrderID := ctx.Get(OrderId).(string)
	MockedID := faker.UUIDHyphenated()
	foodOrder := service.FoodOrder{
		ID:          MockedID,
		OrderID:     MockedOrderID,
		FoodType:    rand.Intn(1),
		FoodName:    generateRandomFood(),
		StationName: ctx.Get(StationName).(string),
		StoreName:   ctx.Get(StoreName).(string),
		Price:       ctx.Get(Price).(float64),
	}

	// Create Test
	newCreateResp, err := cli.CreateFoodOrder(&foodOrder)
	if err != nil {
		log.Errorf("NewCreateFoodOrder request failed, err %s", err)
		return nil, err
	}
	if newCreateResp.Status != 1 {
		log.Errorf("NEwCreateFoodOrder failed")
		return nil, err
	}

	ctx.Set(OrderId, newCreateResp.Data.OrderId)
	ctx.Set(FoodType, newCreateResp.Data.FoodType)
	ctx.Set(StoreName, newCreateResp.Data.StoreName)
	ctx.Set(FoodName, newCreateResp.Data.FoodName)
	ctx.Set(Price, newCreateResp.Data.Price)

	return nil, nil
}

func QueryStationFood(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	resp, err := cli.GetAllStationFood()
	if err != nil {
		log.Errorf("Resp returns err: %v", err)
		return nil, err
	}
	if resp.Status != 1 {
		log.Errorf("GetAllStationFood status should be 1, but is %d", resp.Status)
		return nil, err
	}
	randomIndex := rand.Intn(len(resp.Data))
	ctx.Set(StoreName, resp.Data[randomIndex].StoreName)
	ctx.Set(Phone, resp.Data[randomIndex].Telephone)
	ctx.Set(Price, resp.Data[randomIndex].DeliveryFee)

	return nil, nil

}

func QueryTrainFood(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	resp, err := cli.GetAllTrainFood()
	if err != nil {
		log.Errorf("resp returns err: %v", err)
		return nil, err
	}
	if resp.Status != 1 {
		log.Errorf("GetAllTrainFood's status should be 1 but got %d", resp.Status)
		return nil, nil
	}

	//	Id       string `json:"id"`
	//	TripId   string `json:"tripId"`
	//	FoodList []struct {
	//		FoodName string  `json:"foodName"`
	//		Price    float64 `json:"price"`
	//	} `json:"foodList"`
	//} `json:"data"`

	randomIndex := rand.Intn(len(resp.Data))
	randomFoodlistIndex := rand.Intn(len(resp.Data[randomIndex].FoodList))
	ctx.Set(TripID, resp.Data[randomIndex].TripId)
	ctx.Set(FoodName, resp.Data[randomIndex].FoodList[randomFoodlistIndex].FoodName)
	ctx.Set(FoodPrice, resp.Data[randomIndex].FoodList[randomFoodlistIndex].Price)

	return nil, nil
}

func QueryTrip(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}
	tripInfo := service.TripInfo{
		StartPlace:    ctx.Get(StartStation).(string),
		EndPlace:      ctx.Get(EndStation).(string),
		DepartureTime: ctx.Get(DepartureTime).(string),
	}
	queryInfoResp, err := cli.QueryInfo(tripInfo)
	if err != nil {
		log.Errorf("QueryInfo request failed, err %s", err)
		return nil, err
	}
	if queryInfoResp.Status != 1 {
		log.Errorf("QueryInfo failed, status: %d", queryInfoResp.Status)
		return nil, err
	}

	if len(queryInfoResp.Data) == 0 {
		log.Errorf("QueryInfo response is empty")
		return nil, errors.New("QueryInfo response is empty")
	}

	randomIndex := rand.Intn(len(queryInfoResp.Data))
	ctx.Set(TripId, fmt.Sprintf("%s%s", queryInfoResp.Data[randomIndex].TripId.Type, queryInfoResp.Data[randomIndex].TripId.Number))
	//ctx.Set(TrainTypeName, queryInfoResp.Data[randomIndex].TrainTypeName)
	ctx.Set(StartStation, queryInfoResp.Data[randomIndex].StartStation)
	ctx.Set(TerminalStation, queryInfoResp.Data[randomIndex].TerminalStation)
	ctx.Set(StartTime, queryInfoResp.Data[randomIndex].StartTime)
	ctx.Set(EndTime, queryInfoResp.Data[randomIndex].EndTime)
	ctx.Set(EconomyClass, queryInfoResp.Data[randomIndex].EconomyClass)
	ctx.Set(ConfortClass, queryInfoResp.Data[randomIndex].ConfortClass)
	ctx.Set(PriceForEconomyClass, queryInfoResp.Data[randomIndex].PriceForEconomyClass)
	ctx.Set(PriceForConfortClass, queryInfoResp.Data[randomIndex].PriceForConfortClass)

	return nil, nil
}

func CreateTrip(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	// Mock para
	MockedLoginId := ctx.Get(LoginToken).(string)
	//MockedTripId := GenerateTripId()
	MockedTripId := ctx.Get(TripId).(string)
	//MockedTrainTypeName := generateTrainTypeName(MockedTripId) /*"GaoTieSeven"*/
	MockedTrainTypeName := ctx.Get(TrainTypeName).(string)
	MockedRouteID := ctx.Get(RouteID).(string)
	MockedStartStationName := ctx.Get(From).(string)
	MockedStationsName := ctx.Get(StationName).([]string)
	MockedTerminalStationName := ctx.Get(To).(string)
	MockedStartTime := ctx.Get(StartTime).(string)
	MockedEndTime := ctx.Get(EndTime).(string)

	// Mock input
	travelInfo := service.TravelInfo{
		LoginID:          MockedLoginId,
		TripID:           MockedTripId,
		TrainTypeName:    MockedTrainTypeName,
		RouteID:          MockedRouteID,
		StartStationName: MockedStartStationName,
		//StationsName:        fmt.Sprintf("%v,%v", MockedStartStationName, MockedTerminalStationName),
		StationsName:        strings.Join(MockedStationsName, ","),
		TerminalStationName: MockedTerminalStationName,
		StartTime:           MockedStartTime,
		EndTime:             MockedEndTime,
	}

	// Create Test
	createResp, err := cli.CreateTrip(&travelInfo)
	if err != nil {
		log.Errorf("CreateTrip request failed, err %s", err)
		return nil, err
	}
	if createResp.Status != 1 {
		log.Errorf("CreateTrip failed: %s", createResp.Msg)
		return nil, err
	}
	if createResp.Msg == "Already exists" {
		log.Errorf("Already exists: %s", createResp.Msg)
		return nil, err
	}
	isMatch := false
	if /*createResp.Data.Id == travelInfo.LoginID &&*/
	createResp.Data.StationsName == toLowerCaseAndRemoveSpaces(travelInfo.StationsName) &&
		createResp.Data.StartStationName == toLowerCaseAndRemoveSpaces(travelInfo.StartStationName) &&
		createResp.Data.TerminalStationName == toLowerCaseAndRemoveSpaces(travelInfo.TerminalStationName) &&
		createResp.Data.StartTime == travelInfo.StartTime &&
		createResp.Data.EndTime == travelInfo.EndTime &&
		createResp.Data.TrainTypeName == travelInfo.TrainTypeName &&
		createResp.Data.RouteId == travelInfo.RouteID {
		isMatch = true
	}
	if !isMatch {
		log.Errorf("CreateTrip failed: %s. Except: %v, but get: %v", createResp.Msg, travelInfo, createResp.Data)
		return nil, err
	}

	/*	EndTime             string `json:"endTime"`
		Id                  string `json:"id"`
		RouteId             string `json:"routeId"`
		StartStationName    string `json:"startStationName"`
		StartTime           string `json:"startTime"`
		StationsName        string `json:"stationsName"`
		TerminalStationName string `json:"terminalStationName"`
		TrainTypeName       string `json:"trainTypeName"`
		TripId              TripId `json:"tripId"`*/

	//ctx.Set(TripId, createResp.Data.TripId)
	//ctx.Set(TrainTypeName, createResp.Data.TrainTypeName)
	//ctx.Set(StartStation, createResp.Data.StartStation)
	//ctx.Set(TerminalStation, createResp.Data.TerminalStation)
	//ctx.Set(StartTime, queryInfoResp.Data[randomIndex].StartTime)
	//ctx.Set(EndTime, queryInfoResp.Data[randomIndex].EndTime)
	//ctx.Set(EconomyClass, queryInfoResp.Data[randomIndex].EconomyClass)
	//ctx.Set(ConfortClass, queryInfoResp.Data[randomIndex].ConfortClass)
	//ctx.Set(PriceForEconomyClass, queryInfoResp.Data[randomIndex].PriceForEconomyClass)
	//ctx.Set(PriceForConfortClass, queryInfoResp.Data[randomIndex].PriceForConfortClass)
	return nil, nil
}

func QueryTrain(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	// QueryTraintype all
	allTrainTypes, err := cli.QueryTraintype()
	if err != nil {
		log.Errorf("QueryTraintype all request failed, err %s", err)
		return nil, err
	}
	if allTrainTypes.Status != 1 {
		log.Errorf("allTrainTypes.Status != 1")
		return nil, err
	}
	if len(allTrainTypes.Data) == 0 {
		log.Errorf("QueryTraintype all returned no results")
		return nil, err
	}
	randomIndex := rand.Intn(len(allTrainTypes.Data))
	ctx.Set(TrainTypName, allTrainTypes.Data[randomIndex].Name)
	ctx.Set(ConfortClass, allTrainTypes.Data[randomIndex].ConfortClass)
	ctx.Set(AverageSpeed, allTrainTypes.Data[randomIndex].AverageSpeed)
	ctx.Set(EconomyClass, allTrainTypes.Data[randomIndex].EconomyClass)

	return nil, nil
}

func QueryRoute(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	AllRoutesByQuery, err := cli.QueryAllRoutes()
	if err != nil {
		log.Errorf("Request failed, err2 %s", err)
		return nil, err
	}
	if AllRoutesByQuery.Status != 1 {
		log.Fatal("AllRoutes_By_Query.Status != 1")
		return nil, err
	}

	randomIndex := rand.Intn(len(AllRoutesByQuery.Data))
	debug := AllRoutesByQuery.Data[randomIndex]
	fmt.Println(debug)
	ctx.Set(RouteID, AllRoutesByQuery.Data[randomIndex].Id)
	ctx.Set(StartStation, AllRoutesByQuery.Data[randomIndex].StartStation)
	ctx.Set(EndStation, AllRoutesByQuery.Data[randomIndex].EndStation)
	ctx.Set(StationName, AllRoutesByQuery.Data[randomIndex].Stations)
	ctx.Set(Distances, AllRoutesByQuery.Data[randomIndex].Distances)

	return nil, nil
}

func QueryBasic(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	MockedTripTripId := GenerateTripId()
	MockedTripTripIdType := MockedTripTripId[0]
	MockedTripTripIdNumber := MockedTripTripId[1:]

	travelQuery := &service.Travel{
		Trip: service.Trip{
			Id: faker.UUIDHyphenated(), // randomly generated
			TripId: service.TripId{
				Type:   fmt.Sprintf("%c", MockedTripTripIdType),
				Number: MockedTripTripIdNumber,
			},
			TrainTypeName:       GenerateTrainTypeName(),
			RouteId:             ctx.Get(RouteID).(string),
			StartStationName:    ctx.Get(StartStation).(string),
			StationsName:        strings.Join(ctx.Get(StationName).([]string), ","),
			TerminalStationName: ctx.Get(EndStation).(string),
			StartTime:           "", // can be any
			EndTime:             "", // can be any
		},
		StartPlace: ctx.Get(StartStation).(string),
		EndPlace:   ctx.Get(EndStation).(string),
		//DepartureTime: extractDate(getRandomTime()), // 生成1小时到1天之后的时间
		DepartureTime: getRandomTime(), // 生成1小时到1天之后的时间
	}

	var basicSvc service.BasicService = cli
	travel, err := basicSvc.QueryForTravel(travelQuery)
	if err != nil {
		log.Errorf("QueryTraintype travel request failed, err %s", err)
		return nil, err
	}
	if travel.Status != 1 {
		log.Errorf("travel.Status != 1")
		return nil, err
	}

	ctx.Set(Status, travel.Data.Status)
	ctx.Set(Percent, travel.Data.Percent)
	ctx.Set(TrainType, travel.Data.TrainType)
	ctx.Set(Route, travel.Data.Route)
	ctx.Set(Prices, travel.Data.Prices)
	//
	ctx.Set(DepartureTime, travelQuery.DepartureTime)
	ctx.Set(TrainTypeName, travelQuery.Trip.TrainTypeName)

	return nil, nil
}

func CreateSeat(ctx *Context) (*NodeResult, error) {
	// cli, ok := ctx.Get(Client).(*service.SvcImpl)
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	totalNum := rand.Intn(10) + 10
	seatCreateInfoReq := &service.SeatCreateInfoReq{
		TravelDate:  ctx.Get(TravelDate).(string),
		TrainNumber: ctx.Get(TrainNumber).(string),
		DestStation: ctx.Get(EndStation).(string),
		SeatType:    ctx.Get(SeatClass).(int),
		TotalNum:    totalNum,
		Stations:    ctx.Get(StationName).([]string),
	}

	resp, err := cli.ReqSeatCreate(seatCreateInfoReq)

	if err != nil {
		log.Errorln(err)
		return nil, err
	}
	if resp.Status != 1 {
		log.Errorf("SeatCreateInfoReq.Status != 1, resp: %+v", resp)
		return nil, err
	}

	ctx.Set(SeatNo, resp.Data.SeatNo)
	ctx.Set(DestStation, resp.Data.DestStation)
	//ctx.Set(StartStation, resp.Data.StartStation)

	return nil, nil
}

func QueryStation(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	QueryAll, err7 := cli.QueryStations()
	if err7 != nil {
		log.Errorf("Request failed, err7 %s", err7)
		return nil, err7
	}
	if QueryAll.Status != 1 {
		log.Errorf("Request failed, QueryAll.Status: %d, expected: %d", QueryAll.Status, 1)
		return nil, err7
	}
	randomIndex := rand.Intn(len(QueryAll.Data))
	ctx.Set(StationId, QueryAll.Data[randomIndex].Id)
	ctx.Set(StationNames, QueryAll.Data[randomIndex].Name)
	ctx.Set(StayTime, QueryAll.Data[randomIndex].StayTime)

	return nil, nil
}

func QueryPrice(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	TheRouteId := ctx.Get(RouteID).(string)
	TheTrainType := ctx.Get(TrainTypeName).(string)
	priceByRouteAndTrain, err := cli.FindByRouteIdAndTrainType(TheRouteId, TheTrainType)
	if err != nil {
		log.Errorf("FindByRouteIdAndTrainType failed: %v", err)
		return nil, err
	}
	if priceByRouteAndTrain.Status != 1 {
		log.Errorf("There is not corresponding Ticket available.")
		return &(NodeResult{false}), err
	}

	ctx.Set(BasicPriceRate, priceByRouteAndTrain.Data.BasicPriceRate)
	ctx.Set(FirstClassPriceRate, priceByRouteAndTrain.Data.FirstClassPriceRate)

	return nil, nil
}

func QuerySecurity(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	// Get All Security Configs
	configs, err3 := cli.FindAllSecurityConfig()
	if err3 != nil {
		log.Errorf("FindAllSecurityConfig failed: %v", err3)
		return nil, err3
	}
	if configs.Status != 1 {
		log.Errorf("[Security Service]Status != 1")
		return nil, err3
	}

	randomIndex := rand.Intn(len(configs.Data))
	ctx.Set(SecurityID, configs.Data[randomIndex].ID)
	ctx.Set(SecurityName, configs.Data[randomIndex].Name)
	ctx.Set(SecurityValue, configs.Data[randomIndex].Value)
	ctx.Set(SecurityDescription, configs.Data[randomIndex].Description)

	return nil, nil
}

// ConfigBehaviorChain
func QueryConfig(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	// QueryTraintype All Configs Test
	queryAllResp, err := cli.QueryAllConfigs()
	if err != nil {
		log.Errorf("QueryAllConfigs request failed, err %s", err)
		return nil, err
	}
	if queryAllResp.Status != 1 {
		log.Errorf("QueryAllConfigs status != 1")
		return nil, err
	}

	/*	Name        string `json:"name"`
		Value       string `json:"value"`
		Description string `json:"description"`*/
	randomIndex := rand.Intn(len(queryAllResp.Data))
	ctx.Set(ConfigName, queryAllResp.Data[randomIndex].Name)
	ctx.Set(Value, queryAllResp.Data[randomIndex].Value)
	ctx.Set(Description, queryAllResp.Data[randomIndex].Description)

	return nil, nil
}

func QueryOrder(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	Resp, err := cli.ReqFindAllOrder()
	if err != nil {
		log.Errorf("Request failed, err %s", err)
		return nil, err
	}
	if len(Resp.Data) == 0 {
		log.Errorf("no data found.")
		return nil, err
	}

	randomIndex := rand.Intn(len(Resp.Data))
	//ctx.Set(AccountID, Resp.Data[randomIndex].AccountId)
	ctx.Set(BoughtDate, Resp.Data[randomIndex].BoughtDate)
	ctx.Set(CoachNumber, Resp.Data[randomIndex].CoachNumber)
	ctx.Set(ContactsDocumentNumber, Resp.Data[randomIndex].ContactsDocumentNumber)
	//ctx.Set(ContactsName, Resp.Data[randomIndex].ContactsName)
	ctx.Set(Name, Resp.Data[randomIndex].ContactsName)
	ctx.Set(DifferenceMoney, Resp.Data[randomIndex].DifferenceMoney)
	ctx.Set(SeatClass, Resp.Data[randomIndex].SeatClass)
	ctx.Set(SeatNumber, Resp.Data[randomIndex].SeatNumber)
	ctx.Set(Status, Resp.Data[randomIndex].Status)
	ctx.Set(TrainNumber, Resp.Data[randomIndex].TrainNumber)
	ctx.Set(TravelDate, Resp.Data[randomIndex].TravelDate)
	ctx.Set(TravelTime, Resp.Data[randomIndex].TravelTime)

	return nil, nil
}

func CreateOrder(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	originOrder0 := service.Order{
		AccountId:              ctx.Get(AccountID).(string),
		BoughtDate:             faker.Date(),
		CoachNumber:            generateCoachNumber(),
		ContactsDocumentNumber: generateDocumentNumber(),
		//ContactsName:           ctx.Get(ContactsName).(string),
		ContactsName:    ctx.Get(Name).(string),
		DifferenceMoney: "",
		DocumentType:    0,
		//From:                   ctx.Get(From).(string),
		From: ctx.Get(From).(string), // First, create/query get the station;
		// then put them here -> If you want to create a new Order, you have to do the whole process.
		Id:         "nil",
		Price:      RandomDecimalStringBetween(1, 10),
		SeatClass:  GetTrainTicketClass(),
		SeatNumber: service.GenerateSeatNumber(),
		Status:     0,
		//To:                     ctx.Get(To).(string),
		To:          ctx.Get(To).(string),
		TrainNumber: ctx.Get(TrainNumber).(string),
		TravelDate:  getRandomTime(),
		TravelTime:  generateRandomTime(),
	}

	CreateNewOrderResp, err := cli.ReqCreateNewOrder(&originOrder0)
	if err != nil {
		log.Errorf("Request failed, err %s", err)
		return nil, err
	}
	if CreateNewOrderResp.Status != 1 {
		log.Errorf("Request failed, CreateNewOrder status != 1")
		return nil, err
	}

	//ctx.Set(AccountID, Resp.Data[randomIndex].AccountId)
	ctx.Set(BoughtDate, CreateNewOrderResp.Data.BoughtDate)
	ctx.Set(CoachNumber, CreateNewOrderResp.Data.CoachNumber)
	ctx.Set(ContactsDocumentNumber, CreateNewOrderResp.Data.ContactsDocumentNumber)
	//ctx.Set(ContactsName, Resp.Data[randomIndex].ContactsName)
	ctx.Set(Name, CreateNewOrderResp.Data.ContactsName)
	ctx.Set(DifferenceMoney, CreateNewOrderResp.Data.DifferenceMoney)
	ctx.Set(SeatClass, CreateNewOrderResp.Data.SeatClass)
	ctx.Set(SeatNumber, CreateNewOrderResp.Data.SeatNumber)
	ctx.Set(Status, CreateNewOrderResp.Data.Status)
	ctx.Set(TrainNumber, CreateNewOrderResp.Data.TrainNumber)
	ctx.Set(TravelDate, CreateNewOrderResp.Data.TravelDate)
	ctx.Set(TravelTime, CreateNewOrderResp.Data.TravelTime)

	return nil, nil
}

func QueryOrderOther(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	GetResp, err := cli.ReqFindAllOrderOther()

	if err != nil {
		log.Errorf("Request failed, err %s", err)
		return nil, err
	}
	if GetResp.Status != 1 {
		log.Errorf("Request failed, CreateNewOrder status != 1")
		return nil, err
	}

	randomIndex := rand.Intn(len(GetResp.Data))
	//ctx.Set(AccountID, Resp.Data[randomIndex].AccountId)
	ctx.Set(BoughtDate, GetResp.Data[randomIndex].BoughtDate)
	ctx.Set(CoachNumber, GetResp.Data[randomIndex].CoachNumber)
	ctx.Set(ContactsDocumentNumber, GetResp.Data[randomIndex].ContactsDocumentNumber)
	//ctx.Set(ContactsName, Resp.Data[randomIndex].ContactsName)
	ctx.Set(Name, GetResp.Data[randomIndex].ContactsName)
	ctx.Set(DifferenceMoney, GetResp.Data[randomIndex].DifferenceMoney)
	ctx.Set(SeatClass, GetResp.Data[randomIndex].SeatClass)
	ctx.Set(SeatNumber, GetResp.Data[randomIndex].SeatNumber)
	ctx.Set(Status, GetResp.Data[randomIndex].Status)
	ctx.Set(TrainNumber, GetResp.Data[randomIndex].TrainNumber)
	ctx.Set(TravelDate, GetResp.Data[randomIndex].TravelDate)
	ctx.Set(TravelTime, GetResp.Data[randomIndex].TravelTime)

	return nil, nil
}

func CreateOrderOther(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	AddResp, err := cli.ReqCreateNewOrderOther(&service.Order{
		AccountId:              ctx.Get(AccountID).(string),
		BoughtDate:             faker.Date(),
		CoachNumber:            generateCoachNumber(),
		ContactsDocumentNumber: generateDocumentNumber(),
		//ContactsName:           ctx.Get(ContactsName).(string),
		ContactsName:    ctx.Get(Name).(string),
		DifferenceMoney: "",
		DocumentType:    0,
		//From:                   ctx.Get(From).(string),
		From: ctx.Get(From).(string), // First, create/query get the station;
		// then put them here -> If you want to create a new Order, you have to do the whole process.
		Id:         "nil",
		Price:      RandomDecimalStringBetween(1, 10),
		SeatClass:  GetTrainTicketClass(),
		SeatNumber: service.GenerateSeatNumber(),
		Status:     0,
		//To:                     ctx.Get(To).(string),
		To:          ctx.Get(To).(string),
		TrainNumber: ctx.Get(TrainNumber).(string),
		TravelDate:  getRandomTime(),
		TravelTime:  generateRandomTime(),
	})

	if err != nil {
		log.Errorf("Request failed, err %s", err)
		return nil, err
	}
	if AddResp.Status != 1 {
		log.Errorf("Request failed, CreateNewOrder status != 1")
		return nil, err
	}

	ctx.Set(BoughtDate, AddResp.Data.BoughtDate)
	ctx.Set(CoachNumber, AddResp.Data.CoachNumber)
	ctx.Set(ContactsDocumentNumber, AddResp.Data.ContactsDocumentNumber)
	ctx.Set(Name, AddResp.Data.ContactsName)
	ctx.Set(DifferenceMoney, AddResp.Data.DifferenceMoney)
	ctx.Set(SeatClass, AddResp.Data.SeatClass)
	ctx.Set(SeatNumber, AddResp.Data.SeatNumber)
	ctx.Set(Status, AddResp.Data.Status)
	ctx.Set(TrainNumber, AddResp.Data.TrainNumber)
	ctx.Set(TravelDate, AddResp.Data.TravelDate)
	ctx.Set(TravelTime, AddResp.Data.TravelTime)

	return nil, nil
}

// Preserve Behaviors - The Last One
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
