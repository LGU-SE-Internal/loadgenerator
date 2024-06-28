package behaviors

import (
	"fmt"
	"github.com/Lincyaw/loadgenerator/service"
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"log"
	"math/rand"
	"strconv"
	"time"
)

type PreserveBehavior struct{}

func (o *PreserveBehavior) Run(cli *service.SvcImpl) {
	_, err := cli.ReqUserLogin(&service.UserLoginInfoReq{
		Password:         "111111",
		UserName:         "fdse_microservice",
		VerificationCode: "123",
	})
	if err != nil {
		log.Fatalln(err)
	}

	var preserveSvc service.PreserveService = cli

	// Mock Input Variables
	//DirectQuery_And_Order; Prob = 0.95
	//CreateAndQuery_And_Order; Prob = 0.04
	//Random_Create_And_Order; Prob = 0.01
	var MockedAccountID string
	var MockedContactsID string
	var MockedTripID string
	var MockedSeatType int
	var MockedLoginToken string
	var MockedDate string
	var MockedFromCity string
	var MockedToCity string
	var MockedAssurance int
	var MockedFoodType int
	var MockedStationName string
	var MockedStoreName string
	var MockedFoodName string
	var MockedFoodPrice float64
	var MockedHandleDate string
	var MockedConsigneeName string
	var MockedConsigneePhone string
	var MockedConsigneeWeight float64
	var MockedIsWithin bool

	// For each variable:
	//DirectQuery_And_Order; Prob = 0.95
	//CreateAndQuery_And_Order; Prob = 0.04
	//Random_Create_And_Order; Prob = 0.01

	// Account Service
	var accountSvc service.ContactsService = cli
	// Mock AccountID
	// Generate a random float between 0 and 1
	r0 := rand.Float64()
	NoExistMockedAccountID := false
	if r0 < 0.95 {
		// DirectQuery_And_Order; Prob = 0.95
		//log.Fatalf("Selected: DirectQuery_And_Order")
		GetAllContacts, err := accountSvc.GetAllContacts()
		if err != nil {
			log.Fatalf("[Mock AccountID]GetAllContacts fail. The error occurs: %v", err)
		}

		if len(GetAllContacts.Data) > 0 {
			MockedAccountID = *(GetAllContacts.Data[0].AccountId)
		} else {
			NoExistMockedAccountID = true
		}
	}
	if NoExistMockedAccountID || (r0 < 0.99 && r0 >= 0.95) {
		// CreateAndQuery_And_Order; Prob = 0.04
		//log.Fatalf("Selected: CreateAndQuery_And_Order")
		CreateContactsInput := service.AdminContacts{
			ID:        uuid.NewString(),
			AccountID: uuid.NewString(),
			Name:      faker.Name(),
		}
		CreateContacts, err := accountSvc.AddContact(&CreateContactsInput)
		if err != nil {
			log.Fatalf("[Mock AccountID] CreateContacts error occurs: %v", err)
		}
		MockedAccountID = CreateContacts.Data.AccountId
	} else {
		// Random_Create_And_Order; Prob = 0.01
		//log.Fatalf("Selected: Random_Create_And_Order")
		MockedAccountID = faker.UUIDHyphenated()
	}

	// Contacts Service
	var contactsSvc service.ContactsService = cli
	// MockedContactsID
	// Generate a random float between 0 and 1
	r1 := rand.Float64()
	NoExistMockedContactsID := false
	if r1 < 0.95 {
		// DirectQuery_And_Order; Prob = 0.95
		//log.Fatalf("Selected: DirectQuery_And_Order")
		GetAllContacts, err := contactsSvc.GetAllContacts()
		if err != nil {
			log.Fatalf("[MockedContactsID]GetAllContacts error occurs: %v", err)
		}

		if len(GetAllContacts.Data) > 0 {
			MockedContactsID = GetAllContacts.Data[0].Id
		} else {
			NoExistMockedContactsID = true
		}
	}
	if NoExistMockedContactsID || (r1 < 0.99 && r1 >= 0.95) {
		// CreateAndQuery_And_Order; Prob = 0.04
		//log.Fatalf("Selected: CreateAndQuery_And_Order")
		CreateContactsInput := service.AdminContacts{
			ID:        uuid.NewString(),
			AccountID: uuid.NewString(),
			Name:      faker.Name(),
		}
		CreateContacts, err := contactsSvc.AddContact(&CreateContactsInput)
		if err != nil {
			log.Fatalf("[MockedContactsID] CreateContacts error occurs: %v", err)
		}
		if CreateContacts.Data.Id == "" {
			log.Fatalf("Create AdminContacts Fail. Return Id = ''")
		}
		MockedContactsID = CreateContacts.Data.Id
	} else {
		// Random_Create_And_Order; Prob = 0.01
		//log.Fatalf("Selected: Random_Create_And_Order")
		MockedContactsID = faker.UUIDHyphenated()
	}

	// Travel Service
	var travelSvc service.TravelService = cli
	// MockedTripID
	// Generate a random float between 0 and 1
	r2 := rand.Float64()
	NoExistMockedTripID := false
	if r2 < 0.95 {
		// DirectQuery_And_Order; Prob = 0.95
		//log.Fatalf("Selected: DirectQuery_And_Order")
		GetAllTravel, err := travelSvc.QueryAll()
		if err != nil {
			log.Fatalf("[MockedTripID] error occurs: %v", err)
		}

		if len(GetAllTravel.Data) > 0 {
			MockedTripID = GetAllTravel.Data[0].Id
		} else {
			NoExistMockedTripID = true
		}
	}
	if NoExistMockedTripID || (r2 < 0.99 && r2 >= 0.95) {
		// CreateAndQuery_And_Order; Prob = 0.04
		//log.Fatalf("Selected: CreateAndQuery_And_Order")
		MockedLoginId := faker.UUIDHyphenated()
		MockedTrainTypeName := faker.Word()
		MockedRouteID := faker.UUIDHyphenated()
		MockedStartStationName := "Shenzhen Bei"
		MockedTerminalStationName := "California Airport"
		MockedStartTime := faker.Date()
		MockedEndTime := faker.Date()
		CreateTravelInput := service.TravelInfo{
			LoginID:             MockedLoginId,
			TripID:              "G777",
			TrainTypeName:       MockedTrainTypeName,
			RouteID:             MockedRouteID,
			StartStationName:    MockedStartStationName,
			StationsName:        "Shenzhen Bei, California Airport",
			TerminalStationName: MockedTerminalStationName,
			StartTime:           MockedStartTime,
			EndTime:             MockedEndTime,
		}

		_, err := travelSvc.CreateTrip(&CreateTravelInput)
		if err != nil {
			log.Fatalf("[MockedTripID] CreateTravelInput error occurs: %v", err)
		}

		GetAllTravel, err := travelSvc.QueryAll()
		if err != nil {
			log.Fatalf("[MockedTripID] GetAllTravel error occurs: %v", err)
		}

		if len(GetAllTravel.Data) > 0 {
			MockedTripID = GetAllTravel.Data[len(GetAllTravel.Data)-1].Id
		}
	} else {
		// Random_Create_And_Order; Prob = 0.01
		//log.Fatalf("Selected: Random_Create_And_Order")
		MockedTripID = faker.UUIDHyphenated()
	}

	// Order Service
	var orderSvc service.OrderService = cli
	// MockedSeatType
	// Generate a random float between 0 and 1
	r3 := rand.Float64()
	NoExistMockedSeatType := false
	if r3 < 0.95 {
		// DirectQuery_And_Order; Prob = 0.95
		//log.Fatalf("Selected: DirectQuery_And_Order")
		GetAllOrder, err := orderSvc.ReqFindAllOrder()
		if err != nil {
			log.Fatalf("[MockedSeatType] GetAllOrder error occurs: %v", err)
		}

		if len(GetAllOrder.Data) > 0 {
			MockedSeatType = GetAllOrder.Data[0].SeatClass
		} else {
			NoExistMockedSeatType = true
		}
	}
	if NoExistMockedSeatType || (r3 < 0.99 && r3 >= 0.95) {
		// CreateAndQuery_And_Order; Prob = 0.04
		//log.Fatalf("Selected: CreateAndQuery_And_Order")
		CreateMockedSeatType, err := orderSvc.ReqCreateNewOrder(&service.Order{
			AccountId:              uuid.NewString(),
			BoughtDate:             faker.Date(),
			CoachNumber:            rand.Intn(9) + 1,
			ContactsDocumentNumber: strconv.Itoa(rand.Intn(9) + 1),
			ContactsName:           faker.Name(),
			DifferenceMoney:        RandomDecimalStringBetween(1, 10),
			DocumentType:           0,
			From:                   RandomProvincialCapitalEN(),
			Id:                     uuid.NewString(),
			Price:                  RandomDecimalStringBetween(1, 10),
			SeatClass:              GetTrainTicketClass(),
			SeatNumber:             GenerateSeatNumber(),
			Status:                 0,
			To:                     RandomProvincialCapitalEN(),
			TrainNumber:            GenerateTripId(),
			TravelDate:             faker.Date(),
			TravelTime:             faker.TimeString(),
		})

		if err != nil {
			log.Fatalf("[MockedSeatType] ReqCreateNewOrder error occurs: %v", err)
		}

		GetAllOrder, err := orderSvc.ReqFindAllOrder()
		if err != nil {
			log.Fatalf("[MockedSeatType] MockedSeatType error occurs: %v", err)
		}

		if CreateMockedSeatType.Data.AccountId == "" {
			log.Fatalf("CreateMockedSeatType Fail. AccountId = ''")
		}

		MockedSeatType = GetAllOrder.Data[len(GetAllOrder.Data)-1].SeatClass
	} else {
		// Random_Create_And_Order; Prob = 0.01
		//log.Fatalf("Selected: Random_Create_And_Order")
		MockedSeatType = rand.Intn(3)
	}

	// MockedLoginToken
	r4 := rand.Float64()
	if r4 < 0.95 {
		MockedLoginToken = faker.UUIDHyphenated()
	} else if r4 < 0.99 {
		MockedLoginToken = faker.UUIDHyphenated()
	} else {
		MockedLoginToken = faker.UUIDHyphenated()
	}

	// order service
	// MockedDate
	r5 := rand.Float64()
	NoExistMockedDate := false
	if r5 < 0.95 {
		GetAllOrder, err := orderSvc.ReqFindAllOrder()
		if err != nil {
			log.Fatalf("[MockedDate] GetAllOrder error occurs: %v", err)
		}

		if len(GetAllOrder.Data) > 0 {
			MockedDate = GetAllOrder.Data[0].TravelDate
		} else {
			NoExistMockedDate = true
		}
	}
	if NoExistMockedDate || (r5 < 0.99 && r5 >= 0.95) {
		CeateMockedDate, err := orderSvc.ReqCreateNewOrder(&service.Order{
			AccountId:              uuid.NewString(),
			BoughtDate:             faker.Date(),
			CoachNumber:            rand.Intn(9) + 1,
			ContactsDocumentNumber: strconv.Itoa(rand.Intn(9) + 1),
			ContactsName:           faker.Name(),
			DifferenceMoney:        RandomDecimalStringBetween(1, 10),
			DocumentType:           0,
			From:                   RandomProvincialCapitalEN(),
			Id:                     uuid.NewString(),
			Price:                  RandomDecimalStringBetween(1, 10),
			SeatClass:              GetTrainTicketClass(),
			SeatNumber:             GenerateSeatNumber(),
			Status:                 0,
			To:                     RandomProvincialCapitalEN(),
			TrainNumber:            "G111",
			TravelDate:             faker.Date(),
			TravelTime:             faker.TimeString(),
		})

		if err != nil {
			log.Fatalf("[MockedDate] ReqCreateNewOrder error occurs: %v", err)
		}

		GetAllOrder, err := orderSvc.ReqFindAllOrder()
		if err != nil {
			log.Fatalf("[MockedDate] GetAllOrder error occurs: %v", err)
		}

		if CeateMockedDate.Data.AccountId == "" {
			log.Fatalf("CeateMockedDate Fail. The AccountId = ''")
		}

		MockedDate = GetAllOrder.Data[len(GetAllOrder.Data)-1].TravelDate
	} else {
		MockedDate = faker.Date()
	}

	// Trip Service
	//var tripSvc service.TravelService = cli

	// order service
	// MockedFromCity
	r6 := rand.Float64()
	NoExistMockedFromCity := false
	if r6 < 0.95 {
		//GetAllTrip, err := tripSvc.QueryAll()
		//if err != nil {
		//	log.Fatalf("error occurs: %v", err)
		//}
		//
		//if len(GetAllTrip.Data) > 0 {
		//	MockedFromCity = GetAllTrip.Data[0].StartStationName
		//}
		GetAllOrder, err := orderSvc.ReqFindAllOrder()
		if err != nil {
			log.Fatalf("[MockedFromCity] GetAllOrder error occurs: %v", err)
		}

		if len(GetAllOrder.Data) > 0 {
			MockedFromCity = GetAllOrder.Data[0].From
		} else {
			NoExistMockedFromCity = true
		}
	}
	if NoExistMockedFromCity || (r6 < 0.99 && r6 >= 0.95) {
		CreateMockedFromCity, err := orderSvc.ReqCreateNewOrder(&service.Order{
			AccountId:              uuid.NewString(),
			BoughtDate:             faker.Date(),
			CoachNumber:            rand.Intn(9) + 1,
			ContactsDocumentNumber: strconv.Itoa(rand.Intn(9) + 1),
			ContactsName:           faker.Name(),
			DifferenceMoney:        RandomDecimalStringBetween(1, 10),
			DocumentType:           0,
			From:                   RandomProvincialCapitalEN(),
			Id:                     uuid.NewString(),
			Price:                  RandomDecimalStringBetween(1, 10),
			SeatClass:              GetTrainTicketClass(),
			SeatNumber:             GenerateSeatNumber(),
			Status:                 0,
			To:                     RandomProvincialCapitalEN(),
			TrainNumber:            "G111",
			TravelDate:             faker.Date(),
			TravelTime:             faker.TimeString(),
		})

		if err != nil {
			log.Fatalf("[MockedFromCity]ReqCreateNewOrder error occurs: %v", err)
		}

		GetAllOrder, err := orderSvc.ReqFindAllOrder()
		if err != nil {
			log.Fatalf("[MockedFromCity]GetAllOrder error occurs: %v", err)
		}

		if CreateMockedFromCity.Data.AccountId == "" {
			log.Fatalf("CreateMockedFromCity Fails. The AccountId == '' ")
		}

		MockedFromCity = GetAllOrder.Data[len(GetAllOrder.Data)-1].From
	} else {
		MockedFromCity = faker.GetRealAddress().City
	}

	// MockedToCity
	r7 := rand.Float64()
	NoExistMockedToCity := false
	if r7 < 0.95 {
		//GetAllTrip, err := tripSvc.QueryAll()
		//if err != nil {
		//	log.Fatalf("error occurs: %v", err)
		//}
		//
		//if len(GetAllTrip.Data) > 0 {
		//	MockedFromCity = GetAllTrip.Data[0].StartStationName
		//}
		GetAllOrder, err := orderSvc.ReqFindAllOrder()
		if err != nil {
			log.Fatalf("[MockedToCity]GetAllOrder error occurs: %v", err)
		}

		if len(GetAllOrder.Data) > 0 {
			MockedToCity = GetAllOrder.Data[0].To
		} else {
			NoExistMockedToCity = true
		}
	}
	if NoExistMockedToCity || (r7 < 0.99 && r7 >= 0.95) {
		CreateMockedToCity, err := orderSvc.ReqCreateNewOrder(&service.Order{
			AccountId:              uuid.NewString(),
			BoughtDate:             faker.Date(),
			CoachNumber:            rand.Intn(9) + 1,
			ContactsDocumentNumber: strconv.Itoa(rand.Intn(9) + 1),
			ContactsName:           faker.Name(),
			DifferenceMoney:        RandomDecimalStringBetween(1, 10),
			DocumentType:           0,
			From:                   RandomProvincialCapitalEN(),
			Id:                     uuid.NewString(),
			Price:                  RandomDecimalStringBetween(1, 10),
			SeatClass:              GetTrainTicketClass(),
			SeatNumber:             GenerateSeatNumber(),
			Status:                 0,
			To:                     RandomProvincialCapitalEN(),
			TrainNumber:            "G111",
			TravelDate:             faker.Date(),
			TravelTime:             faker.TimeString(),
		})

		if err != nil {
			log.Fatalf("[MockedToCity] ReqCreateNewOrder error occurs: %v", err)
		}

		GetAllOrder, err := orderSvc.ReqFindAllOrder()
		if err != nil {
			log.Fatalf("[MockedToCity]GetAllOrder error occurs: %v", err)
		}

		if CreateMockedToCity.Data.AccountId == "" {
			log.Fatalf("CreateMockedToCity Fails. The AccountId == '' ")
		}

		MockedToCity = GetAllOrder.Data[len(GetAllOrder.Data)-1].To
	} else {
		MockedToCity = faker.GetRealAddress().City
	}

	// Assurance Servcie
	var assuranceSvc service.AssuranceService = cli
	// MockedAssurance
	r8 := rand.Float64()
	NoExistMockedAssurance := false
	if r8 < 0.95 {
		GetAllAssurance, err := assuranceSvc.GetAllAssurances()
		if err != nil {
			log.Fatalf("[MockedAssurance]GetAllAssurance error occurs: %v", err)
		}

		if len(GetAllAssurance.Data) > 0 {
			MockedAssurance = GetAllAssurance.Data[0].TypeIndex
		} else {
			NoExistMockedAssurance = true
		}
	}
	if NoExistMockedAssurance || (r8 < 0.99 && r8 >= 0.95) {
		MockedAssuranceOrderID := faker.UUIDHyphenated()
		CreateMockedAssurance, err := assuranceSvc.CreateNewAssurance(rand.Intn(1), MockedAssuranceOrderID)
		if err != nil {
			log.Fatalf("[MockedAssurance]CreateNewAssurance error occurs: %v", err)
		}

		GetAllAssurance, err := assuranceSvc.GetAllAssurances()
		if err != nil {
			log.Fatalf("[MockedAssurance]GetAllAssurance error occurs: %v", err)
		}

		if CreateMockedAssurance.Data.Id == "" {
			log.Fatalf("CreateMockedAssurance Fails. The Id == '' ")
		}

		MockedAssurance = GetAllAssurance.Data[len(GetAllAssurance.Data)-1].TypeIndex
	} else {
		MockedAssurance = rand.Intn(1)
	}

	// Food Service
	var foodSvc service.FoodService = cli
	// MockedFoodType
	r9 := rand.Float64()
	NoExistMockedFoodType := false
	if r9 < 0.95 {
		GetAllFood, err := foodSvc.FindAllFoodOrder()
		if err != nil {
			log.Fatalf("[MockedFoodType]GetAllFood error occurs: %v", err)
		}

		if len(GetAllFood.Data) > 0 {
			MockedFoodType = GetAllFood.Data[0].FoodType
		} else {
			NoExistMockedFoodType = true
		}
	}
	if NoExistMockedFoodType || (r9 < 0.99 && r9 >= 0.95) {
		MockedOrderID := faker.UUIDHyphenated()
		MockedID := faker.UUIDHyphenated()
		foodOrder := service.FoodOrder{
			ID:          MockedID,
			OrderID:     MockedOrderID,
			FoodType:    1,
			FoodName:    "HotPot",
			StationName: "Shang Hai",
			StoreName:   "MiaoTing Instant-Boiled Mutton",
			Price:       7.00,
		}
		updateFoodOrder := service.FoodOrder{
			ID:          MockedID,
			OrderID:     MockedOrderID,
			FoodType:    1,
			FoodName:    "HotPot",
			StationName: "Shang Hai",
			StoreName:   "MiaoTing Instant-Boiled Mutton",
			Price:       8.00,
		}
		foodOrders := []service.FoodOrder{foodOrder, updateFoodOrder}
		_, err := foodSvc.CreateFoodOrdersInBatch(foodOrders)
		if err != nil {
			log.Fatalf("[MockedFoodType]CreateFoodOrdersInBatch error occurs: %v", err)
		}

		GetAllFood, err := foodSvc.FindAllFoodOrder()
		if err != nil {
			log.Fatalf("[MockedFoodType]GetAllFood error occurs: %v", err)
		}

		if GetAllFood.Data[len(GetAllFood.Data)-1].Id == "" {
			log.Fatalf("MockedFoodType GetAllFood Fails. The Id == '' ")
		}

		MockedFoodType = GetAllFood.Data[len(GetAllFood.Data)-1].FoodType
	} else {
		MockedFoodType = rand.Intn(2)
	}

	// Food Servcie
	// MockedStationName
	r10 := rand.Float64()
	NoExistMockedStationName := false
	if r10 < 0.95 {
		GetAllFood, err := foodSvc.FindAllFoodOrder()
		if err != nil {
			log.Fatalf("[MockedStationName]GetAllFood error occurs: %v", err)
		}

		if len(GetAllFood.Data) > 0 {
			MockedStationName = GetAllFood.Data[0].StationName
		} else {
			NoExistMockedStationName = true
		}
	}
	if NoExistMockedStationName || (r10 < 0.99 && r10 >= 0.95) {
		MockedOrderID := faker.UUIDHyphenated()
		MockedID := faker.UUIDHyphenated()
		foodOrder := service.FoodOrder{
			ID:          MockedID,
			OrderID:     MockedOrderID,
			FoodType:    1,
			FoodName:    "HotPot",
			StationName: "Shang Hai",
			StoreName:   "MiaoTing Instant-Boiled Mutton",
			Price:       7.00,
		}
		updateFoodOrder := service.FoodOrder{
			ID:          MockedID,
			OrderID:     MockedOrderID,
			FoodType:    1,
			FoodName:    "HotPot",
			StationName: "Shang Hai",
			StoreName:   "MiaoTing Instant-Boiled Mutton",
			Price:       8.00,
		}
		foodOrders := []service.FoodOrder{foodOrder, updateFoodOrder}
		_, err := foodSvc.CreateFoodOrdersInBatch(foodOrders)
		if err != nil {
			log.Fatalf("[MockedStationName]CreateFoodOrdersInBatch error occurs: %v", err)
		}

		GetAllFood, err := foodSvc.FindAllFoodOrder()
		if err != nil {
			log.Fatalf("[MockedStationName]FindAllFoodOrder error occurs: %v", err)
		}

		if GetAllFood.Data[len(GetAllFood.Data)-1].Id == "" {
			log.Fatalf("MockedStationName GetAllFood Fails. The Id == '' ")
		}

		MockedStationName = GetAllFood.Data[len(GetAllFood.Data)-1].StationName
	} else {
		MockedStationName = faker.GetRealAddress().City + "Station"
	}

	// Food Servcie
	// MockedStoreName
	r11 := rand.Float64()
	NoexistMockedStoreName := false
	if r11 < 0.95 {
		GetAllFood, err := foodSvc.FindAllFoodOrder()
		if err != nil {
			log.Fatalf("[MockedStoreName]GetAllFood error occurs: %v", err)
		}

		if len(GetAllFood.Data) > 0 {
			MockedStoreName = GetAllFood.Data[0].StoreName
		} else {
			NoexistMockedStoreName = true
		}
	}
	if NoexistMockedStoreName || (r11 < 0.99 && r11 >= 0.95) {
		MockedOrderID := faker.UUIDHyphenated()
		MockedID := faker.UUIDHyphenated()
		foodOrder := service.FoodOrder{
			ID:          MockedID,
			OrderID:     MockedOrderID,
			FoodType:    1,
			FoodName:    "HotPot",
			StationName: "Shang Hai",
			StoreName:   "MiaoTing Instant-Boiled Mutton",
			Price:       7.00,
		}
		updateFoodOrder := service.FoodOrder{
			ID:          MockedID,
			OrderID:     MockedOrderID,
			FoodType:    1,
			FoodName:    "HotPot",
			StationName: "Shang Hai",
			StoreName:   "MiaoTing Instant-Boiled Mutton",
			Price:       8.00,
		}
		foodOrders := []service.FoodOrder{foodOrder, updateFoodOrder}
		_, err := foodSvc.CreateFoodOrdersInBatch(foodOrders)
		if err != nil {
			log.Fatalf("[MockedStoreName]CreateFoodOrdersInBatch error occurs: %v", err)
		}

		GetAllFood, err := foodSvc.FindAllFoodOrder()
		if err != nil {
			log.Fatalf("[MockedStoreName]GetAllFood error occurs: %v", err)
		}

		if GetAllFood.Data[len(GetAllFood.Data)-1].Id == "" {
			log.Fatalf("MockedStoreName GetAllFood Fails. The Id == '' ")
		}

		MockedStoreName = GetAllFood.Data[len(GetAllFood.Data)-1].StoreName

	} else {
		MockedStoreName = faker.Name() + "Store"
	}

	// Food Servcie
	// MockedFoodName
	r12 := rand.Float64()
	NoExistMockedFoodName := false
	if r12 < 0.95 {
		GetAllFood, err := foodSvc.FindAllFoodOrder()
		if err != nil {
			log.Fatalf("[MockedFoodName]GetAllFood error occurs: %v", err)
		}

		if len(GetAllFood.Data) > 0 {
			MockedFoodName = GetAllFood.Data[0].FoodName
		} else {
			NoExistMockedFoodName = true
		}
	}
	if NoExistMockedFoodName || (r12 < 0.99 && r12 >= 0.95) {
		MockedOrderID := faker.UUIDHyphenated()
		MockedID := faker.UUIDHyphenated()
		foodOrder := service.FoodOrder{
			ID:          MockedID,
			OrderID:     MockedOrderID,
			FoodType:    1,
			FoodName:    "HotPot",
			StationName: "Shang Hai",
			StoreName:   "MiaoTing Instant-Boiled Mutton",
			Price:       7.00,
		}
		updateFoodOrder := service.FoodOrder{
			ID:          MockedID,
			OrderID:     MockedOrderID,
			FoodType:    1,
			FoodName:    "HotPot",
			StationName: "Shang Hai",
			StoreName:   "MiaoTing Instant-Boiled Mutton",
			Price:       8.00,
		}
		foodOrders := []service.FoodOrder{foodOrder, updateFoodOrder}
		_, err := foodSvc.CreateFoodOrdersInBatch(foodOrders)
		if err != nil {
			log.Fatalf("[MockedFoodName]CreateFoodOrdersInBatch error occurs: %v", err)
		}

		GetAllFood, err := foodSvc.FindAllFoodOrder()
		if err != nil {
			log.Fatalf("[MockedFoodName]GetAllFood error occurs: %v", err)
		}

		if GetAllFood.Data[len(GetAllFood.Data)-1].Id == "" {
			log.Fatalf("MockedFoodName Fails. The id = '' ")
		}

		MockedFoodName = GetAllFood.Data[len(GetAllFood.Data)-1].FoodName
	} else {
		MockedFoodName = faker.Name() + "'s Food"
	}

	// Food Servcie
	// MockedFoodPrice
	r13 := rand.Float64()
	NoExistMockedFoodPrice := false
	if r13 < 0.95 {
		GetAllFood, err := foodSvc.FindAllFoodOrder()
		if err != nil {
			log.Fatalf("[MockedFoodPrice]GetAllFood error occurs: %v", err)
		}

		if len(GetAllFood.Data) > 0 {
			MockedFoodPrice = GetAllFood.Data[0].Price
		} else {
			NoExistMockedFoodPrice = true
		}
	}
	if NoExistMockedFoodPrice || (r13 < 0.99 && r13 >= 0.95) {
		MockedOrderID := faker.UUIDHyphenated()
		MockedID := faker.UUIDHyphenated()
		foodOrder := service.FoodOrder{
			ID:          MockedID,
			OrderID:     MockedOrderID,
			FoodType:    1,
			FoodName:    "HotPot",
			StationName: "Shang Hai",
			StoreName:   "MiaoTing Instant-Boiled Mutton",
			Price:       7.00,
		}
		updateFoodOrder := service.FoodOrder{
			ID:          MockedID,
			OrderID:     MockedOrderID,
			FoodType:    1,
			FoodName:    "HotPot",
			StationName: "Shang Hai",
			StoreName:   "MiaoTing Instant-Boiled Mutton",
			Price:       8.00,
		}
		foodOrders := []service.FoodOrder{foodOrder, updateFoodOrder}
		_, err := foodSvc.CreateFoodOrdersInBatch(foodOrders)
		if err != nil {
			log.Fatalf("[MockedFoodPrice]CreateFoodOrdersInBatch error occurs: %v", err)
		}

		GetAllFood, err := foodSvc.FindAllFoodOrder()
		if err != nil {
			log.Fatalf("[MockedFoodPrice]GetAllFood error occurs: %v", err)
		}

		if GetAllFood.Data[len(GetAllFood.Data)-1].Id == "" {
			log.Fatalf("MockedFoodPrice Fails. The id = '' ")
		}

		MockedFoodPrice = GetAllFood.Data[len(GetAllFood.Data)-1].Price
	} else {
		MockedFoodPrice = float64(rand.Intn(7) + 5)
	}

	// Consign Service 000: Consign do not have the QueryAll() function. Should I add one?
	var consignSvc service.ConsignService = cli
	// MockedHandleDate
	r14 := rand.Float64()
	NoExistMockedHandleDate := false
	if r14 < 0.95 {
		GetAllConsignByAccountId, err := consignSvc.QueryByAccountId("3c7ca4eb-4eb2-407a-b870-7fb228d87c5c")
		if err != nil {
			log.Fatalf("[MockedHandleDate]GetAllConsignByAccountId error occurs: %v", err)
		}
		if len(GetAllConsignByAccountId.Data) > 0 {
			MockedHandleDate = GetAllConsignByAccountId.Data[0].HandleDate
		} else {
			NoExistMockedHandleDate = true
		}
	}
	if NoExistMockedHandleDate || (r14 < 0.99 && r14 >= 0.95) {
		MockedId := faker.UUIDHyphenated()
		MockedAccountId := faker.UUIDHyphenated()
		MockedOrderId := faker.UUIDHyphenated()
		MockedHandleDateInput := faker.Date()
		MockedTargetDate := faker.Date()
		MockedFromPlace := "suzhou"
		MockedToPlace := "beijing"
		MockedConsignee := faker.Name()
		MockedPhone := faker.PhoneNumber

		// Insert a new consign record
		insertReq := &service.Consign{
			ID:         MockedId,
			OrderID:    MockedOrderId,
			AccountID:  MockedAccountId,
			HandleDate: MockedHandleDateInput,
			TargetDate: MockedTargetDate,
			From:       MockedFromPlace,
			To:         MockedToPlace,
			Consignee:  MockedConsignee,
			Phone:      MockedPhone,
			Weight:     10.0,
			IsWithin:   true,
		}
		_, err := consignSvc.InsertConsignRecord(insertReq)
		if err != nil {
			log.Fatalf("[MockedHandleDate]InsertConsignRecord error occurs: %v", err)
		}

		GetAllConsignByAccountId, err := consignSvc.QueryByAccountId("3c7ca4eb-4eb2-407a-b870-7fb228d87c5c")
		if err != nil {
			log.Fatalf("[MockedHandleDate]GetAllConsignByAccountId error occurs: %v", err)
		}

		if GetAllConsignByAccountId.Data[len(GetAllConsignByAccountId.Data)-1].AccountID == "" {
			log.Fatalf("MockedHandleDate Fails. The AccountID = '' ")
		}

		MockedHandleDate = GetAllConsignByAccountId.Data[len(GetAllConsignByAccountId.Data)-1].HandleDate
	} else {
		MockedHandleDate = faker.Date()
	}

	// Consign Service
	// MockedConsigneeName
	r15 := rand.Float64()
	NoExistMockedConsigneeName := false
	if r15 < 0.95 {
		GetAllConsignByAccountId, err := consignSvc.QueryByAccountId("3c7ca4eb-4eb2-407a-b870-7fb228d87c5c")
		if err != nil {
			log.Fatalf("[MockedConsigneeName]GetAllConsignByAccountId error occurs: %v", err)
		}
		if len(GetAllConsignByAccountId.Data) > 0 {
			MockedConsigneeName = GetAllConsignByAccountId.Data[0].Consignee
		} else {
			NoExistMockedConsigneeName = true
		}
	}
	if NoExistMockedConsigneeName || (r15 < 0.99 && r15 >= 0.95) {
		MockedId := faker.UUIDHyphenated()
		MockedAccountId := faker.UUIDHyphenated()
		MockedOrderId := faker.UUIDHyphenated()
		MockedHandleDateInput := faker.Date()
		MockedTargetDate := faker.Date()
		MockedFromPlace := "suzhou"
		MockedToPlace := "beijing"
		MockedConsignee := faker.Name()
		MockedPhone := faker.PhoneNumber

		// Insert a new consign record
		insertReq := &service.Consign{
			ID:         MockedId,
			OrderID:    MockedOrderId,
			AccountID:  MockedAccountId,
			HandleDate: MockedHandleDateInput,
			TargetDate: MockedTargetDate,
			From:       MockedFromPlace,
			To:         MockedToPlace,
			Consignee:  MockedConsignee,
			Phone:      MockedPhone,
			Weight:     10.0,
			IsWithin:   true,
		}
		_, err := consignSvc.InsertConsignRecord(insertReq)
		if err != nil {
			log.Fatalf("[MockedConsigneeName]InsertConsignRecord error occurs: %v", err)
		}

		GetAllConsignByAccountId, err := consignSvc.QueryByAccountId("3c7ca4eb-4eb2-407a-b870-7fb228d87c5c")
		if err != nil {
			log.Fatalf("[MockedConsigneeName]GetAllConsignByAccountId error occurs: %v", err)
		}

		if GetAllConsignByAccountId.Data[len(GetAllConsignByAccountId.Data)-1].AccountID == "" {
			log.Fatalf("MockedConsigneeName Fails. The AccountID = '' ")
		}

		MockedConsigneeName = GetAllConsignByAccountId.Data[len(GetAllConsignByAccountId.Data)-1].Consignee
	} else {
		MockedConsigneeName = faker.Name()
	}

	// Consign Service
	// MockedConsigneePhone
	r16 := rand.Float64()
	NoExistMockedConsigneePhone := false
	if r16 < 0.95 {
		GetAllConsignByAccountId, err := consignSvc.QueryByAccountId("3c7ca4eb-4eb2-407a-b870-7fb228d87c5c")
		if err != nil {
			log.Fatalf("[MockedConsigneePhone]GetAllConsignByAccountId error occurs: %v", err)
		}
		if len(GetAllConsignByAccountId.Data) > 0 {
			MockedConsigneePhone = GetAllConsignByAccountId.Data[0].Phone
		} else {
			NoExistMockedConsigneePhone = true
		}
	}
	if NoExistMockedConsigneePhone || (r16 < 0.99 && r16 >= 0.95) {
		MockedId := faker.UUIDHyphenated()
		MockedAccountId := faker.UUIDHyphenated()
		MockedOrderId := faker.UUIDHyphenated()
		MockedHandleDateInput := faker.Date()
		MockedTargetDate := faker.Date()
		MockedFromPlace := "suzhou"
		MockedToPlace := "beijing"
		MockedConsignee := faker.Name()
		MockedPhone := faker.PhoneNumber

		// Insert a new consign record
		insertReq := &service.Consign{
			ID:         MockedId,
			OrderID:    MockedOrderId,
			AccountID:  MockedAccountId,
			HandleDate: MockedHandleDateInput,
			TargetDate: MockedTargetDate,
			From:       MockedFromPlace,
			To:         MockedToPlace,
			Consignee:  MockedConsignee,
			Phone:      MockedPhone,
			Weight:     10.0,
			IsWithin:   true,
		}
		_, err := consignSvc.InsertConsignRecord(insertReq)
		if err != nil {
			log.Fatalf("[MockedConsigneePhone]Consign error occurs: %v", err)
		}

		GetAllConsignByAccountId, err := consignSvc.QueryByAccountId("3c7ca4eb-4eb2-407a-b870-7fb228d87c5c")
		if err != nil {
			log.Fatalf("[MockedConsigneePhone]GetAllConsignByAccountId error occurs: %v", err)
		}

		if GetAllConsignByAccountId.Data[len(GetAllConsignByAccountId.Data)-1].AccountID == "" {
			log.Fatalf("MockedConsigneePhone Fails. The AccountID = '' ")
		}

		MockedConsigneePhone = GetAllConsignByAccountId.Data[len(GetAllConsignByAccountId.Data)-1].Consignee
	} else {
		MockedConsigneePhone = faker.PhoneNumber
	}

	// Consign Service
	// MockedConsigneeWeight
	r17 := rand.Float64()
	NoExistMockedConsigneeWeight := false
	if r17 < 0.95 {
		GetAllConsignByAccountId, err := consignSvc.QueryByAccountId("3c7ca4eb-4eb2-407a-b870-7fb228d87c5c")
		if err != nil {
			log.Fatalf("[MockedConsigneeWeight]GetAllConsignByAccountId error occurs: %v", err)
		}
		if len(GetAllConsignByAccountId.Data) > 0 {
			MockedConsigneeWeight = GetAllConsignByAccountId.Data[0].Weight
		} else {
			NoExistMockedConsigneeWeight = true
		}
	}
	if NoExistMockedConsigneeWeight || (r17 < 0.99 && r17 >= 0.95) {
		MockedId := faker.UUIDHyphenated()
		MockedAccountId := faker.UUIDHyphenated()
		MockedOrderId := faker.UUIDHyphenated()
		MockedHandleDateInput := faker.Date()
		MockedTargetDate := faker.Date()
		MockedFromPlace := "suzhou"
		MockedToPlace := "beijing"
		MockedConsignee := faker.Name()
		MockedPhone := faker.PhoneNumber

		// Insert a new consign record
		insertReq := &service.Consign{
			ID:         MockedId,
			OrderID:    MockedOrderId,
			AccountID:  MockedAccountId,
			HandleDate: MockedHandleDateInput,
			TargetDate: MockedTargetDate,
			From:       MockedFromPlace,
			To:         MockedToPlace,
			Consignee:  MockedConsignee,
			Phone:      MockedPhone,
			Weight:     10.0,
			IsWithin:   true,
		}
		_, err := consignSvc.InsertConsignRecord(insertReq)
		if err != nil {
			log.Fatalf("[MockedConsigneeWeight]InsertConsignRecord error occurs: %v", err)
		}

		GetAllConsignByAccountId, err := consignSvc.QueryByAccountId("3c7ca4eb-4eb2-407a-b870-7fb228d87c5c")
		if err != nil {
			log.Fatalf("[MockedConsigneeWeight]GetAllConsignByAccountId error occurs: %v", err)
		}

		if GetAllConsignByAccountId.Data[len(GetAllConsignByAccountId.Data)-1].AccountID == "" {
			log.Fatalf("MockedConsigneeWeight fails. The AccountID = '' ")
		}

		MockedConsigneeWeight = GetAllConsignByAccountId.Data[len(GetAllConsignByAccountId.Data)-1].Weight
	} else {
		MockedConsigneeWeight = float64(rand.Intn(3) + 10)
	}

	//Consign Service
	// MockedIsWithin 000: Where is the parameter from?
	r18 := rand.Float64()
	if r18 < 0.95 {
		MockedIsWithin = rand.Intn(2) == 0
	} else if r18 < 0.99 {
		MockedIsWithin = rand.Intn(2) == 0
	} else {
		MockedIsWithin = rand.Intn(2) == 0
	}

	// Mock Variables End
	// Put them into Input for preserving
	// Input
	orderTicketsInfo := service.OrderTicketsInfo{
		AccountID:       MockedAccountID,
		ContactsID:      MockedContactsID,
		TripID:          MockedTripID,
		SeatType:        MockedSeatType,
		LoginToken:      MockedLoginToken,
		Date:            MockedDate,
		From:            MockedFromCity,
		To:              MockedToCity,
		Assurance:       MockedAssurance,
		FoodType:        MockedFoodType,
		StationName:     MockedStationName,
		StoreName:       MockedStoreName,
		FoodName:        MockedFoodName,
		FoodPrice:       MockedFoodPrice,
		HandleDate:      MockedHandleDate,
		ConsigneeName:   MockedConsigneeName,
		ConsigneePhone:  MockedConsigneePhone,
		ConsigneeWeight: MockedConsigneeWeight,
		IsWithin:        MockedIsWithin,
	}

	_, err = preserveSvc.Preserve(&orderTicketsInfo)
	if err != nil {
		log.Fatalf("[Input]Preserve error occurs: %v", err)
		//return
	}
	time.Sleep(2 * time.Second)
}

//// helper function
//func GenerateTripId() string {
//	// 设置随机数种子
//	rand.Seed(time.Now().UnixNano())
//
//	// 定义可能的开头字母
//	letters := []rune{'Z', 'T', 'K', 'G', 'D'}
//
//	// 随机选择一个字母
//	startLetter := letters[rand.Intn(len(letters))]
//
//	// 生成三个随机数字
//	randomNumber := rand.Intn(1000)
//
//	// 格式化成三位数字，不足三位前面补零
//	MockedTripID := fmt.Sprintf("%c%03d", startLetter, randomNumber)
//
//	return MockedTripID
//}
//
//func GenerateTrainTypeName() string {
//	// 设置随机数种子
//	rand.Seed(time.Now().UnixNano())
//
//	// 定义可能的火车类型名称
//	trainTypes := []string{"GaoTieOne", "GaoTieTwo", "GaoTieSeven", "DongCheOne", "DongCheTen"}
//
//	// 随机选择一个火车类型名称
//	MockedTrainTypeName := trainTypes[rand.Intn(len(trainTypes))]
//
//	return MockedTrainTypeName
//}
//
//func ListToString(stations []string) string {
//
//	// Use a builder for efficient string concatenation
//	var builder strings.Builder
//
//	for i, station := range stations {
//		if i > 0 {
//			builder.WriteString(", ")
//		}
//		builder.WriteString(fmt.Sprintf("Stations[%d] %s", i, station))
//	}
//
//	result := builder.String()
//	return result
//}
//
//func StringToList(input string) []string {
//	// Split the input string by commas and trim any leading/trailing spaces from each element
//	parts := strings.Split(input, ",")
//	for i := range parts {
//		parts[i] = strings.TrimSpace(parts[i])
//	}
//	return parts
//}
//
//func getRandomTime() string {
//	randomDate := faker.Date()
//	randomTime := faker.TIME
//
//	DateAndTime := randomDate + " " + randomTime
//
//	return DateAndTime
//}

// helper function for Order Service
// RandomDecimalStringBetween 生成并返回两个整数之间的一位小数形式的随机数字符串，包括边界值。
func RandomDecimalStringBetween(min, max int) string {
	rand.Seed(time.Now().UnixNano())
	randomInt := rand.Intn(max-min+1) + min              // 生成[min, max]范围内的随机整数
	decimalValue := float64(randomInt) * 0.1             // 将整数转换为一位小数
	return strconv.FormatFloat(decimalValue, 'f', 1, 64) // 转换为一位小数的字符串形式
}

// RandomProvincialCapitalEN 随机返回一个中国省会城市的英文名称
func RandomProvincialCapitalEN() string {
	rand.Seed(time.Now().UnixNano())
	return provincialCapitalsEN[rand.Intn(len(provincialCapitalsEN))]
}

// 中国省会城市的英文列表
var provincialCapitalsEN = []string{
	"Beijing", "Shanghai", "Tianjin", "Chongqing",
	"Shijiazhuang", "Taiyuan", "Hohhot", "Shenyang", "Changchun", "Harbin",
	"Nanjing", "Hangzhou", "Hefei", "Fuzhou", "Nanchang", "Jinan", "Zhengzhou", "Wuhan", "Changsha", "Guangzhou",
	"Nanning", "Haikou", "Chengdu", "Guiyang", "Kunming", "Lhasa", "Xi'an", "Lanzhou", "Xining", "Yinchuan",
	"Urumqi", "Taipei",
}

// GetTrainTicketClass 随机返回高铁票等级。
// 有5%的概率返回"FirstClass"（头等座），
// 15%的概率返回"BusinessClass"（一等座），
// 剩余80%的概率返回"EconomyClass"（二等座）。
func GetTrainTicketClass() int {
	rand.Seed(time.Now().UnixNano()) // 确保每次运行时随机数种子不同

	probability := rand.Intn(100) // 生成0到99之间的随机数

	switch {
	case probability < 5:
		return 0
	case probability < 20:
		return 1
	default:
		return 2
	}
}

// GenerateSeatNumber 随机生成火车座位号。
// 座位号的格式为一个字符（A、B、C、D、E之一）后跟两位数字。
func GenerateSeatNumber() string {
	// 初始化随机数生成器
	rand.Seed(time.Now().UnixNano())

	// 可选的首字母集合
	seatChars := []rune{'A', 'B', 'C', 'D', 'E'}
	// 随机选择一个首字母
	firstChar := seatChars[rand.Intn(len(seatChars))]

	// 生成后续的两位数字
	var numStr string
	for i := 0; i < 2; i++ {
		numStr += fmt.Sprintf("%d", rand.Intn(10))
	}

	// 拼接首字母和数字部分
	seatNumber := string(firstChar) + numStr

	return seatNumber
}
