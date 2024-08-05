package behaviors

import (
	"github.com/Lincyaw/loadgenerator/service"
	"github.com/go-faker/faker/v4"
	"log"
	"math/rand"
	"time"
)

type TravelplanBehavior struct{}

func (o *TravelplanBehavior) Run(cli *service.SvcImpl) {
	_, err := cli.ReqUserLogin(&service.UserLoginInfoReq{
		Password:         "111111",
		UserName:         "fdse_microservice",
		VerificationCode: "123",
	})
	if err != nil {
		log.Fatalln(err)
	}

	var travelplanSvc service.TravelplanService = cli
	// Mock Input Variables
	//DirectQuery_And_Order; Prob = 0.95
	//CreateAndQuery_And_Order; Prob = 0.04
	//Random_Create_And_Order; Prob = 0.01
	var MockedDepartureTime string
	var MockedEndPlace string
	var MockedStartPlace string

	// MockedDepartureTime
	r0 := rand.Float64()
	NoExistMockedDepartureTime := false
	// Travel Service
	var travelSvc service.TravelService = cli
	if r0 < 0.95 {
		// Query
		GetAllTravelInfo, err := travelSvc.QueryAllTrip()
		if err != nil {
			log.Fatalf("[travelSvc]GetAllTravelInfo occurs errors: %v", err)
		}

		if len(GetAllTravelInfo.Data) > 0 {
			MockedDepartureTime = GetAllTravelInfo.Data[0].StartTime
		} else {
			//MockedDepartureTime = time.Now().Format("2006-01-02 15:04:05")
			NoExistMockedDepartureTime = true
		}
	}
	if NoExistMockedDepartureTime || (r0 < 0.99 && r0 >= 0.95) {
		// Create
		// Mock data
		//MockedTypeName := faker.Word()
		MockedTripID := GenerateTripId()
		MockedLoginId := faker.UUIDHyphenated()
		//MockedIndex := 1
		//MockedTripIDName := faker.Word()
		MockedTrainTypeName := faker.Word()
		MockedRouteID := faker.UUIDHyphenated()
		MockedStartStationName := RandomProvincialCapitalEN()
		MockedTerminalStationName := RandomProvincialCapitalEN()
		MockedStationsName := MockedStartStationName + ", " + MockedTerminalStationName
		MockedStartTime := faker.Date()
		MockedEndTime := faker.Date()

		travelInfo := service.TravelInfo{
			LoginID:             MockedLoginId,
			TripID:              MockedTripID,
			TrainTypeName:       MockedTrainTypeName,
			RouteID:             MockedRouteID,
			StartStationName:    MockedStartStationName,
			StationsName:        MockedStationsName,
			TerminalStationName: MockedTerminalStationName,
			StartTime:           MockedStartTime,
			EndTime:             MockedEndTime,
		}
		_, err := travelSvc.CreateTrip(&travelInfo)
		if err != nil {
			log.Fatalf("[travelSvc]CreateTrip  occurs errors: %v", err)
		}

		// Query
		GetAllTravelInfo, err := travelSvc.QueryAllTrip()
		if err != nil {
			log.Fatalf("[travelSvc]GetAllTravelInfo occurs errors: %v", err)
		}

		if len(GetAllTravelInfo.Data) > 0 {
			MockedDepartureTime = GetAllTravelInfo.Data[0].StartTime
		} else {
			log.Fatalf("[MockedDepartureTime]create fail. No data.")
		}
	} else {
		MockedDepartureTime = getRandomTime()
	}

	//MockedEndPlace
	r1 := rand.Float64()
	NoExistMockedEndPlace := false
	// travel servcie
	//var travelSvc service.TravelService = cli
	if r1 < 0.95 {
		// Query
		GetAllTravelInfo, err := travelSvc.QueryAllTrip()
		if err != nil {
			log.Fatalf("[travelSvc]GetAllTravelInfo:MockedEndPlace occurs errors: %v", err)
		}

		if len(GetAllTravelInfo.Data) > 0 {
			MockedEndPlace = GetAllTravelInfo.Data[0].TerminalStationName
		} else {
			NoExistMockedEndPlace = true
		}
	}
	if NoExistMockedEndPlace || (r1 < 0.99 && r1 >= 0.95) {
		// Create
		// Mock data
		//MockedTypeName := faker.Word()
		MockedTripID := GenerateTripId()
		MockedLoginId := faker.UUIDHyphenated()
		//MockedIndex := 1
		//MockedTripIDName := faker.Word()
		MockedTrainTypeName := faker.Word()
		MockedRouteID := faker.UUIDHyphenated()
		MockedStartStationName := RandomProvincialCapitalEN()
		MockedTerminalStationName := RandomProvincialCapitalEN()
		MockedStationsName := MockedStartStationName + ", " + MockedTerminalStationName
		MockedStartTime := faker.Date()
		MockedEndTime := faker.Date()

		travelInfo := service.TravelInfo{
			LoginID:             MockedLoginId,
			TripID:              MockedTripID,
			TrainTypeName:       MockedTrainTypeName,
			RouteID:             MockedRouteID,
			StartStationName:    MockedStartStationName,
			StationsName:        MockedStationsName,
			TerminalStationName: MockedTerminalStationName,
			StartTime:           MockedStartTime,
			EndTime:             MockedEndTime,
		}
		_, err := travelSvc.CreateTrip(&travelInfo)
		if err != nil {
			log.Fatalf("[travelSvc]CreateTrip:MockedEndPlace  occurs errors: %v", err)
		}

		// Query
		GetAllTravelInfo, err := travelSvc.QueryAllTrip()
		if err != nil {
			log.Fatalf("[travelSvc]GetAllTravelInfo: MockedEndPlace occurs errors: %v", err)
		}

		if len(GetAllTravelInfo.Data) > 0 {
			MockedEndPlace = GetAllTravelInfo.Data[0].TerminalStationName
		} else {
			log.Fatalf("[MockedDepartureTime]create fail. No data.")
		}
	} else {
		MockedEndPlace = RandomProvincialCapitalEN()
	}

	//MockedStartPlace
	r2 := rand.Float64()
	NoExistMockedStartPlace := false
	// travel servcie
	//var travelSvc service.TravelService = cli
	if r2 < 0.95 {
		// Query
		GetAllTravelInfo, err := travelSvc.QueryAllTrip()
		if err != nil {
			log.Fatalf("[travelSvc]GetAllTravelInfo:MockedStartPlace occurs errors: %v", err)
		}

		if len(GetAllTravelInfo.Data) > 0 {
			MockedStartPlace = GetAllTravelInfo.Data[0].TerminalStationName
		} else {
			NoExistMockedStartPlace = true
		}
	}
	if NoExistMockedStartPlace || (r2 < 0.99 && r2 >= 0.95) {
		// Create
		// Mock data
		//MockedTypeName := faker.Word()
		MockedTripID := GenerateTripId()
		MockedLoginId := faker.UUIDHyphenated()
		//MockedIndex := 1
		//MockedTripIDName := faker.Word()
		MockedTrainTypeName := faker.Word()
		MockedRouteID := faker.UUIDHyphenated()
		MockedStartStationName := RandomProvincialCapitalEN()
		MockedTerminalStationName := RandomProvincialCapitalEN()
		MockedStationsName := MockedStartStationName + ", " + MockedTerminalStationName
		MockedStartTime := faker.Date()
		MockedEndTime := faker.Date()

		travelInfo := service.TravelInfo{
			LoginID:             MockedLoginId,
			TripID:              MockedTripID,
			TrainTypeName:       MockedTrainTypeName,
			RouteID:             MockedRouteID,
			StartStationName:    MockedStartStationName,
			StationsName:        MockedStationsName,
			TerminalStationName: MockedTerminalStationName,
			StartTime:           MockedStartTime,
			EndTime:             MockedEndTime,
		}
		_, err := travelSvc.CreateTrip(&travelInfo)
		if err != nil {
			log.Fatalf("[travelSvc]CreateTrip:MockedStartPlace  occurs errors: %v", err)
		}

		// Query
		GetAllTravelInfo, err := travelSvc.QueryAllTrip()
		if err != nil {
			log.Fatalf("[travelSvc]GetAllTravelInfo: MockedStartPlace occurs errors: %v", err)
		}

		if len(GetAllTravelInfo.Data) > 0 {
			MockedStartPlace = GetAllTravelInfo.Data[0].TerminalStationName
		} else {
			log.Fatalf("[MockedDepartureTime]create fail. No data.")
		}
	} else {
		MockedStartPlace = RandomProvincialCapitalEN()
	}

	travelQueryInfo := service.TravelQueryInfo{
		DepartureTime: MockedDepartureTime,
		EndPlace:      MockedEndPlace,
		StartPlace:    MockedStartPlace,
	}
	_, err = travelplanSvc.ReqGetByCheapest(&travelQueryInfo)
	if err != nil {
		log.Fatalf("[ReqGetByCheapest] error occurs: %v", err)
	}
	time.Sleep(2 * time.Second)

	_, err = travelplanSvc.ReqGetByQuickest(&travelQueryInfo)
	if err != nil {
		log.Fatalf("[ReqGetByQuickest] error occurs: %v", err)
	}
	time.Sleep(2 * time.Second)

	_, err = travelplanSvc.ReqGetByQuickest(&travelQueryInfo)
	if err != nil {
		log.Fatalf("[ReqGetByQuickest] error occurs: %v", err)
	}
	time.Sleep(2 * time.Second)

	// Mock variables
	var MockedEndStation string
	var MockedStartStation string
	var MockedTrainType string
	var MockedTravelDate string
	var MockedViaStation string

	//MockedEndStation
	r3 := rand.Float64()
	NoExistMockedEndStation := false
	// travel servcie
	//var travelSvc service.TravelService = cli
	if r3 < 0.95 {
		// Query
		GetAllTravelInfo, err := travelSvc.QueryAllTrip()
		if err != nil {
			log.Fatalf("[travelSvc]GetAllTravelInfo:MockedEndStation occurs errors: %v", err)
		}

		if len(GetAllTravelInfo.Data) > 0 {
			MockedEndStation = GetAllTravelInfo.Data[0].TerminalStationName
		} else {
			NoExistMockedEndStation = true
		}
	}
	if NoExistMockedEndStation || (r3 < 0.99 && r3 >= 0.95) {
		// Create
		// Mock data
		//MockedTypeName := faker.Word()
		MockedTripID := GenerateTripId()
		MockedLoginId := faker.UUIDHyphenated()
		//MockedIndex := 1
		//MockedTripIDName := faker.Word()
		MockedTrainTypeName := faker.Word()
		MockedRouteID := faker.UUIDHyphenated()
		MockedStartStationName := RandomProvincialCapitalEN()
		MockedTerminalStationName := RandomProvincialCapitalEN()
		MockedStationsName := MockedStartStationName + ", " + MockedTerminalStationName
		MockedStartTime := faker.Date()
		MockedEndTime := faker.Date()

		travelInfo := service.TravelInfo{
			LoginID:             MockedLoginId,
			TripID:              MockedTripID,
			TrainTypeName:       MockedTrainTypeName,
			RouteID:             MockedRouteID,
			StartStationName:    MockedStartStationName,
			StationsName:        MockedStationsName,
			TerminalStationName: MockedTerminalStationName,
			StartTime:           MockedStartTime,
			EndTime:             MockedEndTime,
		}
		_, err := travelSvc.CreateTrip(&travelInfo)
		if err != nil {
			log.Fatalf("[travelSvc]CreateTrip:MockedEndStation  occurs errors: %v", err)
		}

		// Query
		GetAllTravelInfo, err := travelSvc.QueryAllTrip()
		if err != nil {
			log.Fatalf("[travelSvc]GetAllTravelInfo: MockedEndStation occurs errors: %v", err)
		}

		if len(GetAllTravelInfo.Data) > 0 {
			MockedEndStation = GetAllTravelInfo.Data[0].TerminalStationName
		} else {
			log.Fatalf("[MockedDepartureTime]create fail. No data.")
		}
	} else {
		MockedEndStation = RandomProvincialCapitalEN()
	}

	//MockedStartStation
	r4 := rand.Float64()
	NoExistMockedStartStation := false
	// travel servcie
	//var travelSvc service.TravelService = cli
	if r4 < 0.95 {
		// Query
		GetAllTravelInfo, err := travelSvc.QueryAllTrip()
		if err != nil {
			log.Fatalf("[travelSvc]GetAllTravelInfo:MockedStartStation occurs errors: %v", err)
		}

		if len(GetAllTravelInfo.Data) > 0 {
			MockedStartStation = GetAllTravelInfo.Data[0].TerminalStationName
		} else {
			NoExistMockedStartStation = true
		}
	}
	if NoExistMockedStartStation || (r4 < 0.99 && r4 >= 0.95) {
		// Create
		// Mock data
		//MockedTypeName := faker.Word()
		MockedTripID := GenerateTripId()
		MockedLoginId := faker.UUIDHyphenated()
		//MockedIndex := 1
		//MockedTripIDName := faker.Word()
		MockedTrainTypeName := faker.Word()
		MockedRouteID := faker.UUIDHyphenated()
		MockedStartStationName := RandomProvincialCapitalEN()
		MockedTerminalStationName := RandomProvincialCapitalEN()
		MockedStationsName := MockedStartStationName + ", " + MockedTerminalStationName
		MockedStartTime := faker.Date()
		MockedEndTime := faker.Date()

		travelInfo := service.TravelInfo{
			LoginID:             MockedLoginId,
			TripID:              MockedTripID,
			TrainTypeName:       MockedTrainTypeName,
			RouteID:             MockedRouteID,
			StartStationName:    MockedStartStationName,
			StationsName:        MockedStationsName,
			TerminalStationName: MockedTerminalStationName,
			StartTime:           MockedStartTime,
			EndTime:             MockedEndTime,
		}
		_, err := travelSvc.CreateTrip(&travelInfo)
		if err != nil {
			log.Fatalf("[travelSvc]CreateTrip:MockedStartStation  occurs errors: %v", err)
		}

		// Query
		GetAllTravelInfo, err := travelSvc.QueryAllTrip()
		if err != nil {
			log.Fatalf("[travelSvc]GetAllTravelInfo: MockedStartStation occurs errors: %v", err)
		}

		if len(GetAllTravelInfo.Data) > 0 {
			MockedStartStation = GetAllTravelInfo.Data[0].TerminalStationName
		} else {
			log.Fatalf("[MockedDepartureTime]create fail. No data.")
		}
	} else {
		MockedStartStation = RandomProvincialCapitalEN()
	}

	// MockedTrainType
	r5 := rand.Float64()
	NoExistMockedTrainType := false
	// Travel Service
	if r5 < 0.95 {
		// Query
		GetAllTravelInfo, err := travelSvc.QueryAllTrip()
		if err != nil {
			log.Fatalf("[travelSvc]GetAllTravelInfo:MockedStartStation occurs errors: %v", err)
		}

		if len(GetAllTravelInfo.Data) > 0 {
			MockedTrainType = GetAllTravelInfo.Data[0].TripId.Type
		} else {
			NoExistMockedTrainType = true
		}
	}
	if NoExistMockedTrainType || (r5 < 0.99 && r5 >= 0.95) {
		// Create
		// Mock data
		//MockedTypeName := faker.Word()
		MockedTripID := GenerateTripId()
		MockedLoginId := faker.UUIDHyphenated()
		//MockedIndex := 1
		//MockedTripIDName := faker.Word()
		MockedTrainTypeName := faker.Word()
		MockedRouteID := faker.UUIDHyphenated()
		MockedStartStationName := RandomProvincialCapitalEN()
		MockedTerminalStationName := RandomProvincialCapitalEN()
		MockedStationsName := MockedStartStationName + ", " + MockedTerminalStationName
		MockedStartTime := faker.Date()
		MockedEndTime := faker.Date()

		travelInfo := service.TravelInfo{
			LoginID:             MockedLoginId,
			TripID:              MockedTripID,
			TrainTypeName:       MockedTrainTypeName,
			RouteID:             MockedRouteID,
			StartStationName:    MockedStartStationName,
			StationsName:        MockedStationsName,
			TerminalStationName: MockedTerminalStationName,
			StartTime:           MockedStartTime,
			EndTime:             MockedEndTime,
		}
		_, err := travelSvc.CreateTrip(&travelInfo)
		if err != nil {
			log.Fatalf("[travelSvc]CreateTrip:MockedStartStation  occurs errors: %v", err)
		}

		// Query
		GetAllTravelInfo, err := travelSvc.QueryAllTrip()
		if err != nil {
			log.Fatalf("[travelSvc]GetAllTravelInfo: MockedStartStation occurs errors: %v", err)
		}

		if len(GetAllTravelInfo.Data) > 0 {
			MockedTrainType = GetAllTravelInfo.Data[0].TripId.Type
		} else {
			log.Fatalf("[MockedTrainType]create fail. No data.")
		}
	} else {
		// 定义可能的开头字母
		letters := []string{"Z", "T", "K", "G", "D"}

		// 随机选择一个字母
		startLetter := letters[rand.Intn(len(letters))]

		MockedTrainType = startLetter
	}

	//MockedTravelDate
	r6 := rand.Float64()
	NoExistMockedTravelDate := false
	// Travel Servie
	if r6 < 0.95 {
		// Query
		GetAllTravelInfo, err := travelSvc.QueryAllTrip()
		if err != nil {
			log.Fatalf("[travelSvc]GetAllTravelInfo:MockedStartStation occurs errors: %v", err)
		}

		if len(GetAllTravelInfo.Data) > 0 {
			MockedTravelDate = GetAllTravelInfo.Data[0].TripId.Type
		} else {
			NoExistMockedTravelDate = true
		}
	}
	if NoExistMockedTravelDate || (r6 < 0.99 && r6 >= 0.95) {
		// Create
		// Mock data
		//MockedTypeName := faker.Word()
		MockedTripID := GenerateTripId()
		MockedLoginId := faker.UUIDHyphenated()
		//MockedIndex := 1
		//MockedTripIDName := faker.Word()
		MockedTrainTypeName := faker.Word()
		MockedRouteID := faker.UUIDHyphenated()
		MockedStartStationName := RandomProvincialCapitalEN()
		MockedTerminalStationName := RandomProvincialCapitalEN()
		MockedStationsName := MockedStartStationName + ", " + MockedTerminalStationName
		MockedStartTime := faker.Date()
		MockedEndTime := faker.Date()

		travelInfo := service.TravelInfo{
			LoginID:             MockedLoginId,
			TripID:              MockedTripID,
			TrainTypeName:       MockedTrainTypeName,
			RouteID:             MockedRouteID,
			StartStationName:    MockedStartStationName,
			StationsName:        MockedStationsName,
			TerminalStationName: MockedTerminalStationName,
			StartTime:           MockedStartTime,
			EndTime:             MockedEndTime,
		}
		_, err := travelSvc.CreateTrip(&travelInfo)
		if err != nil {
			log.Fatalf("[travelSvc]CreateTrip:MockedStartStation  occurs errors: %v", err)
		}

		// Query
		GetAllTravelInfo, err := travelSvc.QueryAllTrip()
		if err != nil {
			log.Fatalf("[travelSvc]GetAllTravelInfo: MockedStartStation occurs errors: %v", err)
		}

		if len(GetAllTravelInfo.Data) > 0 {
			MockedTravelDate = GetAllTravelInfo.Data[0].StartTime
		} else {
			log.Fatalf("[MockedTrainType]create fail. No data.")
		}
	} else {
		MockedTravelDate = getRandomTime()
	}

	// MockedViaStation
	r7 := rand.Float64()
	NoExistMockedViaStation := false
	// Travel service
	if r7 < 0.95 {
		// Query
		GetAllTravelInfo, err := travelSvc.QueryAllTrip()
		if err != nil {
			log.Fatalf("[travelSvc]GetAllTravelInfo:MockedStartStation occurs errors: %v", err)
		}

		if len(GetAllTravelInfo.Data) > 0 {
			MockedViaStation = GetAllTravelInfo.Data[0].TripId.Type
		} else {
			NoExistMockedViaStation = true
		}
	}
	if NoExistMockedViaStation || (r7 < 0.99 && r7 >= 0.95) {
		// Create
		// Mock data
		//MockedTypeName := faker.Word()
		MockedTripID := GenerateTripId()
		MockedLoginId := faker.UUIDHyphenated()
		//MockedIndex := 1
		//MockedTripIDName := faker.Word()
		MockedTrainTypeName := faker.Word()
		MockedRouteID := faker.UUIDHyphenated()
		MockedStartStationName := RandomProvincialCapitalEN()
		MockedTerminalStationName := RandomProvincialCapitalEN()
		MockedStationsName := MockedStartStationName + ", " + MockedTerminalStationName
		MockedStartTime := faker.Date()
		MockedEndTime := faker.Date()

		travelInfo := service.TravelInfo{
			LoginID:             MockedLoginId,
			TripID:              MockedTripID,
			TrainTypeName:       MockedTrainTypeName,
			RouteID:             MockedRouteID,
			StartStationName:    MockedStartStationName,
			StationsName:        MockedStationsName,
			TerminalStationName: MockedTerminalStationName,
			StartTime:           MockedStartTime,
			EndTime:             MockedEndTime,
		}
		_, err := travelSvc.CreateTrip(&travelInfo)
		if err != nil {
			log.Fatalf("[travelSvc]CreateTrip:MockedStartStation  occurs errors: %v", err)
		}

		// Query
		GetAllTravelInfo, err := travelSvc.QueryAllTrip()
		if err != nil {
			log.Fatalf("[travelSvc]GetAllTravelInfo: MockedStartStation occurs errors: %v", err)
		}

		if len(GetAllTravelInfo.Data) > 0 {
			MockedViaStation = GetAllTravelInfo.Data[0].StationsName
		} else {
			log.Fatalf("[MockedTrainType]create fail. No data.")
		}
	} else {
		MockedViaStation = RandomProvincialCapitalEN()
	}

	// Mock input
	transferTravelQueryInfo := service.TransferTravelQueryInfo{
		EndStation:   MockedEndStation,
		StartStation: MockedStartStation,
		TrainType:    MockedTrainType,
		TravelDate:   MockedTravelDate,
		ViaStation:   MockedViaStation,
	}
	_, err = travelplanSvc.ReqTransferResult(&transferTravelQueryInfo)
	if err != nil {
		log.Fatalf("[ReqTransferResult] error occurs: %v", err)
	}
	time.Sleep(2 * time.Second)
}
