package behaviors

import (
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
	loginResult, err := cli.ReqUserLogin(&service.UserLoginInfoReq{
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
	if r0 < 0.9999 {
		// DirectQuery_And_Order; Prob = 0.95
		//log.Fatalf("Selected: DirectQuery_And_Order")
		GetAllContacts, err := accountSvc.GetAllContacts()
		if err != nil {
			log.Fatalf("[Mock AccountID]GetAllContacts fail. The error occurs: %v", err)
		}
		if GetAllContacts.Status != 1 {
			log.Fatalf("[Mock AccountID]GetAllContacts.Status != 1")
		}

		if len(GetAllContacts.Data) > 0 {
			MockedAccountID = GetAllContacts.Data[0].AccountId
		} else {
			NoExistMockedAccountID = true
		}
	}
	if NoExistMockedAccountID {
		// CreateAndQuery_And_Order; Prob = 0.04
		//log.Fatalf("Selected: CreateAndQuery_And_Order")
		CreateContactsInput := service.AdminContacts{
			Id:        faker.UUIDHyphenated(),
			AccountId: faker.UUIDHyphenated(),
			Name:      faker.Name(),
		}
		CreateContacts, err := accountSvc.AddContact(&CreateContactsInput)
		if err != nil {
			log.Fatalf("[Mock AccountID] CreateContacts error occurs: %v", err)
		}
		if CreateContacts.Status != 1 {
			log.Fatalf("[Mock AccountID] CreateContacts.Status != 1")
		}
		MockedAccountID = CreateContacts.Data.AccountId
	}

	// Contacts Service
	var contactsSvc service.ContactsService = cli
	// MockedContactsID
	// Generate a random float between 0 and 1
	r1 := rand.Float64()
	NoExistMockedContactsID := false
	if r1 < 0.9999 {
		// DirectQuery_And_Order; Prob = 0.95
		//log.Fatalf("Selected: DirectQuery_And_Order")
		GetAllContacts, err := contactsSvc.GetAllContacts()
		if err != nil {
			log.Fatalf("[MockedContactsID]GetAllContacts error occurs: %v", err)
		}
		if GetAllContacts.Status != 1 {
			log.Fatalf("[Mock AccountID] GetAllContacts.Status != 1")
		}

		if len(GetAllContacts.Data) > 0 {
			MockedContactsID = GetAllContacts.Data[0].Id
		} else {
			NoExistMockedContactsID = true
		}
	}
	if NoExistMockedContactsID {
		CreateContactsInput := service.AdminContacts{
			Id:             faker.UUIDHyphenated(),
			AccountId:      MockedAccountID,
			Name:           faker.Name(),
			DocumentNumber: "DocumentNumber_One",
			DocumentType:   rand.Intn(5),
			PhoneNumber:    faker.Phonenumber(),
		}
		CreateContacts, err := contactsSvc.AddContact(&CreateContactsInput)
		if err != nil {
			log.Fatalf("[MockedContactsID] CreateContacts error occurs: %v", err)
		}
		if CreateContacts.Status != 1 {
			log.Fatalf("[Mock AccountID] CreateContacts.Status != 1")
		}
		if CreateContacts.Data.Id == "" {
			log.Printf("Create AdminContacts Fail: %+v", CreateContacts)
		}
		MockedContactsID = CreateContacts.Data.Id
	}

	// Travel Service
	var travelSvc service.TravelService = cli
	// MockedTripID
	// Generate a random float between 0 and 1
	r2 := rand.Float64()
	NoExistMockedTripID := false
	if r2 < 0.9999 {
		// DirectQuery_And_Order; Prob = 0.95
		//log.Fatalf("Selected: DirectQuery_And_Order")
		GetAllTravel, err := travelSvc.QueryAll()
		if err != nil {
			log.Fatalf("[MockedTripID] error occurs: %v", err)
		}
		if GetAllTravel.Status != 1 {
			log.Fatalf("[MockedTripID] GetAllTravel.Status != 1")
		}

		if len(GetAllTravel.Data) > 0 {
			MockedTripID = GetAllTravel.Data[0].Id
		} else {
			NoExistMockedTripID = true
		}
	}
	if NoExistMockedTripID {
		// CreateAndQuery_And_Order; Prob = 0.04
		//log.Fatalf("Selected: CreateAndQuery_And_Order")
		MockedLoginId := loginResult.Data.Token
		MockedTrainTypeName := GenerateTrainTypeName()
		MockedRouteID := faker.UUIDHyphenated()
		MockedStartStationName := faker.GetRealAddress().City
		MockedStationsName := faker.GetRealAddress().City
		MockedTerminalStationName := faker.GetRealAddress().City
		MockedStartTime := getRandomTime()
		MockedEndTime := getRandomTime()
		MockedTripId := GenerateTripId()
		CreateTravelInput := service.TravelInfo{
			LoginID:             MockedLoginId,
			TripID:              MockedTripId,
			TrainTypeName:       MockedTrainTypeName,
			RouteID:             MockedRouteID,
			StartStationName:    MockedStartStationName,
			StationsName:        MockedStationsName,
			TerminalStationName: MockedTerminalStationName,
			StartTime:           MockedStartTime,
			EndTime:             MockedEndTime,
		}

		CreateTripRsp, err := travelSvc.CreateTrip(&CreateTravelInput)
		if err != nil {
			log.Fatalf("[MockedTripID] CreateTravelInput error occurs: %v", err)
		}
		if CreateTripRsp.Status != 1 {
			log.Fatalf("[MockedTripID] CreateTripRsp.Status != 1")
		}

		GetAllTravel, err := travelSvc.QueryAll()
		if err != nil {
			log.Fatalf("[MockedTripID] GetAllTravel error occurs: %v", err)
		}
		if GetAllTravel.Status != 1 {
			log.Fatalf("[MockedTripID] GetAllTravel.Status != 1")
		}

		if len(GetAllTravel.Data) > 0 {
			MockedTripID = GetAllTravel.Data[len(GetAllTravel.Data)-1].Id
		}
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
		if GetAllOrder.Status != 1 {
			log.Fatalf("[MockedSeatType] GetAllOrder.Status != 1")
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
		_, err = orderSvc.ReqCreateNewOrder(&service.Order{
			AccountId:              MockedAccountID,
			BoughtDate:             getRandomTime(),
			CoachNumber:            rand.Intn(9) + 1,
			ContactsDocumentNumber: strconv.Itoa(rand.Intn(9) + 1),
			ContactsName:           faker.Name(),
			DifferenceMoney:        RandomDecimalStringBetween(1, 10),
			DocumentType:           0,
			From:                   RandomProvincialCapitalEN(),
			Id:                     uuid.NewString(),
			Price:                  RandomDecimalStringBetween(1, 10),
			SeatClass:              GetTrainTicketClass(),
			SeatNumber:             service.GenerateSeatNumber(),
			Status:                 0,
			To:                     RandomProvincialCapitalEN(),
			TrainNumber:            GenerateTripId(),
			TravelDate:             getRandomTime(),
			TravelTime:             generateRandomTime(),
		})

		if err != nil {
			log.Fatalf("[MockedSeatType] ReqCreateNewOrder error occurs: %v", err)
		}

		GetAllOrder, err := orderSvc.ReqFindAllOrder()
		if err != nil {
			log.Fatalf("[MockedSeatType] MockedSeatType error occurs: %v", err)
		}
		if GetAllOrder.Status != 1 {
			log.Fatalf("[MockedSeatType] MockedSeatType Status != 1")
		}

		if len(GetAllOrder.Data) == 0 {
			log.Fatalf("Get all the order fail. There is not order information")
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
		if GetAllOrder.Status != 1 {
			log.Fatalf("[MockedDate] GetAllOrder.Status != 1")
		}

		if len(GetAllOrder.Data) > 0 {
			MockedDate = GetAllOrder.Data[0].TravelDate
		} else {
			NoExistMockedDate = true
		}
	}
	if NoExistMockedDate || (r5 < 0.99 && r5 >= 0.95) {
		_, err := orderSvc.ReqCreateNewOrder(&service.Order{
			AccountId:              MockedAccountID,
			BoughtDate:             getRandomTime(),
			CoachNumber:            rand.Intn(9) + 1,
			ContactsDocumentNumber: strconv.Itoa(rand.Intn(9) + 1),
			ContactsName:           faker.Name(),
			DifferenceMoney:        RandomDecimalStringBetween(1, 10),
			DocumentType:           0,
			From:                   RandomProvincialCapitalEN(),
			Id:                     uuid.NewString(),
			Price:                  RandomDecimalStringBetween(1, 10),
			SeatClass:              GetTrainTicketClass(),
			SeatNumber:             service.GenerateSeatNumber(),
			Status:                 0,
			To:                     RandomProvincialCapitalEN(),
			TrainNumber:            GenerateTripId(),
			TravelDate:             getRandomTime(),
			TravelTime:             faker.TimeString(),
		})

		if err != nil {
			log.Fatalf("[MockedDate] ReqCreateNewOrder error occurs: %v", err)
		}

		GetAllOrder, err := orderSvc.ReqFindAllOrder()
		if err != nil {
			log.Fatalf("[MockedDate] GetAllOrder error occurs: %v", err)
		}
		if GetAllOrder.Status != 1 {
			log.Fatalf("[MockedDate] GetAllOrder.Status != 1")
		}

		MockedDate = GetAllOrder.Data[len(GetAllOrder.Data)-1].TravelDate
	} else {
		MockedDate = getRandomTime()
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
		if GetAllOrder.Status != 1 {
			log.Fatalf("[MockedFromCity] GetAllOrder.Status != 1")
		}

		if len(GetAllOrder.Data) > 0 {
			MockedFromCity = GetAllOrder.Data[0].From
		} else {
			NoExistMockedFromCity = true
		}
	}
	if NoExistMockedFromCity || (r6 < 0.99 && r6 >= 0.95) {
		CreateMockedFromCity, err := orderSvc.ReqCreateNewOrder(&service.Order{
			AccountId:              MockedAccountID,
			BoughtDate:             getRandomTime(),
			CoachNumber:            rand.Intn(9) + 1,
			ContactsDocumentNumber: strconv.Itoa(rand.Intn(9) + 1),
			ContactsName:           faker.Name(),
			DifferenceMoney:        RandomDecimalStringBetween(1, 10),
			DocumentType:           0,
			From:                   RandomProvincialCapitalEN(),
			Id:                     faker.UUIDHyphenated(),
			Price:                  RandomDecimalStringBetween(1, 10),
			SeatClass:              GetTrainTicketClass(),
			SeatNumber:             service.GenerateSeatNumber(),
			Status:                 0,
			To:                     RandomProvincialCapitalEN(),
			TrainNumber:            GenerateTripId(),
			TravelDate:             getRandomTime(),
			TravelTime:             generateRandomTime(),
		})

		if err != nil {
			log.Fatalf("[MockedFromCity]ReqCreateNewOrder error occurs: %v", err)
		}

		GetAllOrder, err := orderSvc.ReqFindAllOrder()
		if err != nil {
			log.Fatalf("[MockedFromCity]GetAllOrder error occurs: %v", err)
		}
		if GetAllOrder.Status != 1 {
			log.Fatalf("[MockedFromCity] GetAllOrder.Status != 1")
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
		_, err := orderSvc.ReqCreateNewOrder(&service.Order{
			AccountId:              MockedAccountID,
			BoughtDate:             getRandomTime(),
			CoachNumber:            rand.Intn(9) + 1,
			ContactsDocumentNumber: strconv.Itoa(rand.Intn(9) + 1),
			ContactsName:           faker.Name(),
			DifferenceMoney:        RandomDecimalStringBetween(1, 10),
			DocumentType:           0,
			From:                   RandomProvincialCapitalEN(),
			Id:                     faker.UUIDHyphenated(),
			Price:                  RandomDecimalStringBetween(1, 10),
			SeatClass:              GetTrainTicketClass(),
			SeatNumber:             service.GenerateSeatNumber(),
			Status:                 0,
			To:                     RandomProvincialCapitalEN(),
			TrainNumber:            GenerateTripId(),
			TravelDate:             getRandomTime(),
			TravelTime:             generateRandomTime(),
		})

		if err != nil {
			log.Fatalf("[MockedToCity] ReqCreateNewOrder error occurs: %v", err)
		}

		GetAllOrder, err := orderSvc.ReqFindAllOrder()
		if err != nil {
			log.Fatalf("[MockedToCity]GetAllOrder error occurs: %v", err)
		}
		if GetAllOrder.Status != 1 {
			log.Fatalf("[MockedToCity] GetAllOrder.Status != 1")
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
		if GetAllAssurance.Status != 1 {
			log.Fatalf("[MockedAssurance] GetAllAssurance.Status != 1")
		}

		if len(GetAllAssurance.Data) > 0 {
			MockedAssurance = GetAllAssurance.Data[0].TypeIndex
		} else {
			NoExistMockedAssurance = true
		}
	}
	if NoExistMockedAssurance || (r8 < 0.99 && r8 >= 0.95) {
		MockedAssuranceOrderID := faker.UUIDHyphenated()
		CreateMockedAssurance, err := assuranceSvc.CreateNewAssurance(1, MockedAssuranceOrderID)
		if err != nil {
			log.Fatalf("[MockedAssurance]CreateNewAssurance error occurs: %v", err)
		}

		GetAllAssurance, err := assuranceSvc.GetAllAssurances()
		if err != nil {
			log.Fatalf("[MockedAssurance]GetAllAssurance error occurs: %v", err)
		}
		if GetAllAssurance.Status != 1 {
			log.Fatalf("[MockedAssurance] GetAllAssurance.Status != 1")
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
			FoodName:    generateRandomFood(),
			StationName: generateRandomCityName(),
			StoreName:   generateRandomStoreName(),
			Price:       7.00,
		}
		updateFoodOrder := service.FoodOrder{
			ID:          MockedID,
			OrderID:     MockedOrderID,
			FoodType:    1,
			FoodName:    generateRandomFood(),
			StationName: generateRandomCityName(),
			StoreName:   generateRandomStoreName(),
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
		if GetAllFood.Status != 1 {
			log.Fatalf("[MockedFoodType] GetAllFood.Status != 1")
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
	var StationSvc service.StationService = cli
	r10 := rand.Float64()
	NoExistMockedStationName := false
	if r10 < 0.95 {
		stations, err := StationSvc.QueryStations()
		if err != nil {
			log.Fatalf("[MockedStationName]GetAllFood error occurs: %v", err)
		}

		if len(stations.Data) > 0 {
			MockedStationName = stations.Data[0].Name
		} else {
			NoExistMockedStationName = true
		}
	}
	if NoExistMockedStationName || (r10 < 0.99 && r10 >= 0.95) {
		createStationresp, err := StationSvc.CreateStation(&service.Station{
			ID:       faker.UUIDHyphenated(),
			Name:     faker.GetRealAddress().City,
			StayTime: rand.Intn(10),
		})
		if err != nil {
			log.Fatalf("[MockedStationName]CreateStation error occurs: %v", err)
		}
		MockedStationName = createStationresp.Data.Name
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
			FoodName:    generateRandomFood(),
			StationName: generateRandomCityName(),
			StoreName:   generateRandomStoreName(),
			Price:       7.00,
		}
		updateFoodOrder := service.FoodOrder{
			ID:          MockedID,
			OrderID:     MockedOrderID,
			FoodType:    1,
			FoodName:    generateRandomFood(),
			StationName: generateRandomCityName(),
			StoreName:   generateRandomStoreName(),
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
		if GetAllFood.Status != 1 {
			log.Fatalf("[MockedStoreName]GetAllFood.Status != 1")
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
			FoodName:    generateRandomFood(),
			StationName: generateRandomCityName(),
			StoreName:   generateRandomStoreName(),
			Price:       7.00,
		}
		updateFoodOrder := service.FoodOrder{
			ID:          MockedID,
			OrderID:     MockedOrderID,
			FoodType:    1,
			FoodName:    generateRandomFood(),
			StationName: generateRandomCityName(),
			StoreName:   generateRandomStoreName(),
			Price:       8.00,
		}
		foodOrders := []service.FoodOrder{foodOrder, updateFoodOrder}
		CreateFoodOrdersInBatchRsp, err := foodSvc.CreateFoodOrdersInBatch(foodOrders)
		if err != nil {
			log.Fatalf("[MockedFoodName]CreateFoodOrdersInBatch error occurs: %v", err)
		}
		if CreateFoodOrdersInBatchRsp.Status != 1 {
			log.Fatalf("[MockedFoodName]CreateFoodOrdersInBatchRsp.Status != 1")
		}

		GetAllFood, err := foodSvc.FindAllFoodOrder()
		if err != nil {
			log.Fatalf("[MockedFoodName]GetAllFood error occurs: %v", err)
		}
		if GetAllFood.Status != 1 {
			log.Fatalf("[MockedFoodName]GetAllFood.Status != 1")
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
			FoodName:    generateRandomFood(),
			StationName: generateRandomCityName(),
			StoreName:   generateRandomStoreName(),
			Price:       7.00,
		}
		updateFoodOrder := service.FoodOrder{
			ID:          MockedID,
			OrderID:     MockedOrderID,
			FoodType:    1,
			FoodName:    generateRandomFood(),
			StationName: generateRandomCityName(),
			StoreName:   generateRandomStoreName(),
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
		if GetAllFood.Status != 1 {
			log.Fatalf("[MockedFoodName]GetAllFood.Status != 1")
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
		GetAllConsignByAccountId, err := consignSvc.QueryByAccountId(MockedAccountID)
		if err != nil {
			log.Fatalf("[MockedHandleDate]GetAllConsignByAccountId error occurs: %v", err)
		}
		if GetAllConsignByAccountId.Status != 1 {
			log.Fatalf("[MockedHandleDate]GetAllConsignByAccountId Status != 1")
		}
		if len(GetAllConsignByAccountId.Data) > 0 {
			MockedHandleDate = GetAllConsignByAccountId.Data[0].HandleDate
		} else {
			NoExistMockedHandleDate = true
		}
	}
	if NoExistMockedHandleDate || (r14 < 0.99 && r14 >= 0.95) {
		MockedId := faker.UUIDHyphenated()
		MockedOrderId := faker.UUIDHyphenated()
		MockedHandleDateInput := getRandomTime()
		MockedTargetDate := getRandomTime()
		MockedFromPlace := generateRandomCityName()
		MockedToPlace := generateRandomCityName()
		MockedConsignee := faker.Name()
		MockedPhone := faker.PhoneNumber

		// Insert a new consign record
		insertReq := &service.Consign{
			ID:         MockedId,
			OrderID:    MockedOrderId,
			AccountID:  MockedAccountID,
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

		GetAllConsignByAccountId, err := consignSvc.QueryByAccountId(MockedAccountID)
		if err != nil {
			log.Fatalf("[MockedHandleDate]GetAllConsignByAccountId error occurs: %v", err)
		}
		if GetAllConsignByAccountId.Status != 1 {
			log.Fatalf("[MockedHandleDate]GetAllConsignByAccountId Status != 1")
		}

		if len(GetAllConsignByAccountId.Data) == 0 || GetAllConsignByAccountId.Data[len(GetAllConsignByAccountId.Data)-1].AccountID == "" {
			log.Fatalf("MockedHandleDate Fails. Consign Data: %v, account id: %v\n", GetAllConsignByAccountId.Data, MockedAccountID)
		}

		MockedHandleDate = GetAllConsignByAccountId.Data[len(GetAllConsignByAccountId.Data)-1].HandleDate
	} else {
		MockedHandleDate = getRandomTime()
	}

	// Consign Service
	// MockedConsigneeName
	r15 := rand.Float64()
	NoExistMockedConsigneeName := false
	if r15 < 0.95 {
		GetAllConsignByAccountId, err := consignSvc.QueryByAccountId(MockedAccountID)
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
		MockedOrderId := faker.UUIDHyphenated()
		MockedHandleDateInput := getRandomTime()
		MockedTargetDate := getRandomTime()
		MockedFromPlace := generateRandomCityName()
		MockedToPlace := generateRandomCityName()
		MockedConsignee := faker.Name()
		MockedPhone := faker.PhoneNumber

		// Insert a new consign record
		insertReq := &service.Consign{
			ID:         MockedId,
			OrderID:    MockedOrderId,
			AccountID:  MockedAccountID,
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

		GetAllConsignByAccountId, err := consignSvc.QueryByAccountId(MockedAccountID)
		if err != nil {
			log.Fatalf("[MockedConsigneeName]GetAllConsignByAccountId error occurs: %v", err)
		}
		if GetAllConsignByAccountId.Status != 1 {
			log.Fatalf("[MockedConsigneeName]GetAllConsignByAccountId Status != 1")
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
		GetAllConsignByAccountId, err := consignSvc.QueryByAccountId(MockedAccountID)
		if err != nil {
			log.Fatalf("[MockedConsigneePhone]GetAllConsignByAccountId error occurs: %v, accountid: %v", err, MockedAccountID)
		}
		if len(GetAllConsignByAccountId.Data) > 0 {
			MockedConsigneePhone = GetAllConsignByAccountId.Data[0].Phone
		} else {
			NoExistMockedConsigneePhone = true
		}
	}
	if NoExistMockedConsigneePhone || (r16 < 0.99 && r16 >= 0.95) {
		MockedId := faker.UUIDHyphenated()
		MockedOrderId := faker.UUIDHyphenated()
		MockedHandleDateInput := getRandomTime()
		MockedTargetDate := getRandomTime()
		MockedFromPlace := generateRandomCityName()
		MockedToPlace := generateRandomCityName()
		MockedConsignee := faker.Name()
		MockedPhone := faker.PhoneNumber

		// Insert a new consign record
		insertReq := &service.Consign{
			ID:         MockedId,
			OrderID:    MockedOrderId,
			AccountID:  MockedAccountID,
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

		GetAllConsignByAccountId, err := consignSvc.QueryByAccountId(MockedAccountID)
		if err != nil {
			log.Fatalf("[MockedConsigneePhone]GetAllConsignByAccountId error occurs: %v", err)
		}
		if GetAllConsignByAccountId.Status != 1 {
			log.Fatalf("[MockedConsigneePhone]GetAllConsignByAccountId Status != 1")
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
		GetAllConsignByAccountId, err := consignSvc.QueryByAccountId(MockedAccountID)
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
		MockedHandleDateInput := getRandomTime()
		MockedTargetDate := getRandomTime()
		MockedFromPlace := generateRandomCityName()
		MockedToPlace := generateRandomCityName()
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

		GetAllConsignByAccountId, err := consignSvc.QueryByAccountId(MockedAccountID)
		if err != nil {
			log.Fatalf("[MockedConsigneeWeight]GetAllConsignByAccountId error occurs: %v", err)
		}
		if GetAllConsignByAccountId.Status != 1 {
			log.Fatalf("[MockedConsigneePhone]GetAllConsignByAccountId Status != 1")
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

	result, err := preserveSvc.Preserve(&orderTicketsInfo)
	if err != nil {
		log.Fatalf("[Input]Preserve error occurs: %v", err)
		//return
	}
	if result.Status != 1 {
		log.Fatalf("[Input]Preserve Status != 1. The result Status is %v", result.Status)
	}
	log.Printf("preserve response: %+v", result)
	time.Sleep(1 * time.Millisecond)
}

// helper function for Order Service
/*// RandomDecimalStringBetween 生成并返回两个整数之间的一位小数形式的随机数字符串，包括边界值。
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
}*/
