package behaviors

import (
	"fmt"
	"github.com/Lincyaw/loadgenerator/service"
	"github.com/go-faker/faker/v4"
	"math/rand"
	"strings"
	"time"
)

type TravelBehavior struct{}

func (o *TravelBehavior) Run(cli *service.SvcImpl) {
	var travelSvc service.TravelService = cli

	// TravelService_FullIntegration

	// Mock Input Variables
	//DirectQuery_And_Order; Prob = 0.95
	//CreateAndQuery_And_Order; Prob = 0.04
	//Random_Create_And_Order; Prob = 0.01
	var MockedLoginId string
	var MockedTripID string
	var MockedTrainTypeName string
	var MockedRouteID string
	var MockedStartStationName string
	var MockedTerminalStationName string
	var MockedStationsName string
	var MockedStartTime string
	var MockedEndTime string

	// 1. Query
	_, err := travelSvc.QueryAll()
	if err != nil {
		fmt.Printf("error occurs: %v", err)
	}
	time.Sleep(2 * time.Second)

	// 2. Create
	// Mock Create Input
	MockedLoginId = faker.UUIDHyphenated()
	MockedTripID = GenerateTripId()
	MockedTrainTypeName = GenerateTrainTypeName()
	MockedRouteID = faker.UUIDHyphenated()
	MockedStartStationName = faker.GetRealAddress().City
	MockedTerminalStationName = faker.GetRealAddress().City
	MockedStationsName = MockedStartStationName + ", " + MockedTerminalStationName
	MockedStartTime = faker.Date()
	MockedEndTime = faker.Date()
	// Input
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
	_, err1 := travelSvc.CreateTrip(&travelInfo)
	if err1 != nil {
		fmt.Printf("error1 occurs: %v\n", err1)
	}
	time.Sleep(2 * time.Second)

	// 3. Query Again
	_, err2 := travelSvc.QueryAll()
	if err2 != nil {
		fmt.Printf("error2 occurs: %v", err2)
	}
	time.Sleep(2 * time.Second)

	// 4. Update
	// Mock data
	//DirectQuery_And_Order; Prob = 0.95
	//CreateAndQuery_And_Order; Prob = 0.04
	//Random_Create_And_Order; Prob = 0.01

	// Service
	// Travel Service
	//LoginId
	r0 := rand.Float64()
	if r0 < 0.95 {
		QueryAllTravelInfo, err := travelSvc.QueryAll()
		if err != nil {
			fmt.Printf("error occurs: %v", err)
		}

		if len(QueryAllTravelInfo.Data) > 0 {
			MockedLoginId = QueryAllTravelInfo.Data[0].Id
		} else {
			fmt.Printf("The corresponding database is empty")
			MockedLoginId = faker.UUIDHyphenated()
		}
	} else if r0 < 0.99 {
		// Create itself
		MockedLoginId = faker.UUIDHyphenated()
	} else {
		MockedLoginId = faker.UUIDHyphenated()
	}

	// Service
	// Travel Service
	// TripID
	r1 := rand.Float64()
	if r1 < 0.95 {
		QueryAllTravelInfo, err := travelSvc.QueryAll()
		if err != nil {
			fmt.Printf("error occurs: %v", err)
		}

		if len(QueryAllTravelInfo.Data) > 0 {
			MockedTripID = *(QueryAllTravelInfo.Data[0].TripId.Type) + QueryAllTravelInfo.Data[0].TripId.Number
		} else {
			fmt.Printf("The corresponding database is empty")
			MockedTripID = GenerateTripId()
		}
	} else if r1 < 0.99 {
		// Create itself
		MockedTripID = GenerateTripId()
	} else {
		MockedTripID = GenerateTripId()
	}

	// Service
	// Travel Service
	//TrainTypeName
	r2 := rand.Float64()
	if r2 < 0.95 {
		QueryAllTravelInfo, err := travelSvc.QueryAll()
		if err != nil {
			fmt.Printf("error occurs: %v", err)
		}

		if len(QueryAllTravelInfo.Data) > 0 {
			MockedTrainTypeName = QueryAllTravelInfo.Data[0].TrainTypeName
		} else {
			fmt.Printf("The corresponding database is empty")
			MockedTrainTypeName = GenerateTrainTypeName()
		}
	} else if r2 < 0.99 {
		// Create itself
		MockedTrainTypeName = GenerateTrainTypeName()
	} else {
		MockedTrainTypeName = GenerateTrainTypeName()
	}

	// Service
	// Route Service
	var routeSvc service.RouteService = cli
	//RouteID
	r3 := rand.Float64()
	if r3 < 0.95 {
		GetAllRouteInfo, err := routeSvc.QueryAllRoutes()
		if err != nil {
			fmt.Printf("error occurs: %v", err)
		}

		if len(GetAllRouteInfo.Data) > 0 {
			MockedRouteID = GetAllRouteInfo.Data[0].Id
		} else {
			fmt.Printf("The corresponding database is empty")
			MockedRouteID = faker.UUIDHyphenated()
		}
	} else if r3 < 0.99 {
		MockedRouteInfoID := faker.UUIDHyphenated()
		MockedRouteInfoStartStation := faker.GetRealAddress().City
		MockedRouteInfoEndStation := faker.GetRealAddress().City
		MockedStationList := MockedRouteInfoStartStation + ", " + faker.GetRealAddress().City + ", " + MockedRouteInfoEndStation
		MockedDistanceList := fmt.Sprintf("%d, %d, %d", rand.Intn(1000), rand.Intn(1000), rand.Intn(1000))
		CreateAndModifyRouteInput := service.RouteInfo{
			ID:           MockedRouteInfoID,
			StartStation: MockedRouteInfoStartStation,
			EndStation:   MockedRouteInfoEndStation,
			StationList:  MockedStationList,
			DistanceList: MockedDistanceList,
		}

		_, err := routeSvc.CreateAndModifyRoute(&CreateAndModifyRouteInput)
		if err != nil {
			fmt.Printf("error occurs: %v", err)
		}

		GetAllRouteInfo, err := routeSvc.QueryAllRoutes()
		if err != nil {
			fmt.Printf("error occurs: %v", err)
		}

		MockedRouteID = GetAllRouteInfo.Data[len(GetAllRouteInfo.Data)-1].Id
	} else {
		MockedRouteID = faker.UUIDHyphenated()
	}

	// Service
	// Route Service
	// StartStationName
	r4 := rand.Float64()
	if r4 < 0.95 {
		GetAllRouteInfo, err := routeSvc.QueryAllRoutes()
		if err != nil {
			fmt.Printf("error occurs: %v", err)
		}

		if len(GetAllRouteInfo.Data) > 0 {
			MockedStartStationName = GetAllRouteInfo.Data[0].StartStation
		} else {
			fmt.Printf("The corresponding database is empty")
			MockedStartStationName = faker.GetRealAddress().City
		}
	} else if r4 < 0.99 {
		MockedRouteInfoID := faker.UUIDHyphenated()
		MockedRouteInfoStartStation := faker.GetRealAddress().City
		MockedRouteInfoEndStation := faker.GetRealAddress().City
		MockedStationList := MockedRouteInfoStartStation + ", " + faker.GetRealAddress().City + ", " + MockedRouteInfoEndStation
		MockedDistanceList := fmt.Sprintf("%d, %d, %d", rand.Intn(1000), rand.Intn(1000), rand.Intn(1000))
		CreateAndModifyRouteInput := service.RouteInfo{
			ID:           MockedRouteInfoID,
			StartStation: MockedRouteInfoStartStation,
			EndStation:   MockedRouteInfoEndStation,
			StationList:  MockedStationList,
			DistanceList: MockedDistanceList,
		}

		_, err := routeSvc.CreateAndModifyRoute(&CreateAndModifyRouteInput)
		if err != nil {
			fmt.Printf("error occurs: %v", err)
		}

		GetAllRouteInfo, err := routeSvc.QueryAllRoutes()
		if err != nil {
			fmt.Printf("error occurs: %v", err)
		}

		MockedStartStationName = GetAllRouteInfo.Data[len(GetAllRouteInfo.Data)-1].StartStation
	} else {
		MockedStartStationName = faker.GetRealAddress().City
	}

	// Service
	// Route Service
	//TerminalStationName
	r5 := rand.Float64()
	if r5 < 0.95 {
		GetAllRouteInfo, err := routeSvc.QueryAllRoutes()
		if err != nil {
			fmt.Printf("error occurs: %v", err)
		}

		if len(GetAllRouteInfo.Data) > 0 {
			MockedTerminalStationName = GetAllRouteInfo.Data[0].EndStation
		} else {
			fmt.Printf("The corresponding database is empty")
			MockedTerminalStationName = faker.GetRealAddress().City
		}
	} else if r5 < 0.99 {
		MockedRouteInfoID := faker.UUIDHyphenated()
		MockedRouteInfoStartStation := faker.GetRealAddress().City
		MockedRouteInfoEndStation := faker.GetRealAddress().City
		MockedStationList := MockedRouteInfoStartStation + ", " + faker.GetRealAddress().City + ", " + MockedRouteInfoEndStation
		MockedDistanceList := fmt.Sprintf("%d, %d, %d", rand.Intn(1000), rand.Intn(1000), rand.Intn(1000))
		CreateAndModifyRouteInput := service.RouteInfo{
			ID:           MockedRouteInfoID,
			StartStation: MockedRouteInfoStartStation,
			EndStation:   MockedRouteInfoEndStation,
			StationList:  MockedStationList,
			DistanceList: MockedDistanceList,
		}

		_, err := routeSvc.CreateAndModifyRoute(&CreateAndModifyRouteInput)
		if err != nil {
			fmt.Printf("error occurs: %v", err)
		}

		GetAllRouteInfo, err := routeSvc.QueryAllRoutes()
		if err != nil {
			fmt.Printf("error occurs: %v", err)
		}

		MockedTerminalStationName = GetAllRouteInfo.Data[len(GetAllRouteInfo.Data)-1].EndStation
	} else {
		MockedTerminalStationName = faker.GetRealAddress().City
	}

	// Service
	// Route Service
	// StationsName
	r6 := rand.Float64()
	if r6 < 0.95 {
		GetAllRouteInfo, err := routeSvc.QueryAllRoutes()
		if err != nil {
			fmt.Printf("error occurs: %v", err)
		}

		if len(GetAllRouteInfo.Data) > 0 {
			MockedStationsName = ListToString(GetAllRouteInfo.Data[0].Stations)
		} else {
			fmt.Printf("The corresponding database is empty")
			MockedStationsName = faker.GetRealAddress().City
		}
	} else if r6 < 0.99 {
		MockedRouteInfoID := faker.UUIDHyphenated()
		MockedRouteInfoStartStation := faker.GetRealAddress().City
		MockedRouteInfoEndStation := faker.GetRealAddress().City
		MockedStationList := MockedRouteInfoStartStation + ", " + faker.GetRealAddress().City + ", " + MockedRouteInfoEndStation
		MockedDistanceList := fmt.Sprintf("%d, %d, %d", rand.Intn(1000), rand.Intn(1000), rand.Intn(1000))
		CreateAndModifyRouteInput := service.RouteInfo{
			ID:           MockedRouteInfoID,
			StartStation: MockedRouteInfoStartStation,
			EndStation:   MockedRouteInfoEndStation,
			StationList:  MockedStationList,
			DistanceList: MockedDistanceList,
		}

		_, err := routeSvc.CreateAndModifyRoute(&CreateAndModifyRouteInput)
		if err != nil {
			fmt.Printf("error occurs: %v", err)
		}

		GetAllRouteInfo, err := routeSvc.QueryAllRoutes()
		if err != nil {
			fmt.Printf("error occurs: %v", err)
		}

		MockedStationsName = ListToString(GetAllRouteInfo.Data[len(GetAllRouteInfo.Data)-1].Stations)
	} else {
		MockedStationsName = MockedStartStationName + ", " + MockedTerminalStationName
	}

	// Service
	// Travel Service
	// StartTime
	r7 := rand.Float64()
	if r7 < 0.95 {
		QueryAllTravelInfo, err := travelSvc.QueryAll()
		if err != nil {
			fmt.Printf("error occurs: %v", err)
		}

		if len(QueryAllTravelInfo.Data) > 0 {
			MockedStartTime = QueryAllTravelInfo.Data[0].StartTime
		} else {
			fmt.Printf("The corresponding database is empty")
			MockedStartTime = getRandomTime()
		}
	} else if r7 < 0.99 {
		// Create itself
		MockedStartTime = getRandomTime()
	} else {
		MockedStartTime = getRandomTime()
	}

	// Service
	// Travel Service
	// EndTime
	r8 := rand.Float64()
	if r8 < 0.95 {
		QueryAllTravelInfo, err := travelSvc.QueryAll()
		if err != nil {
			fmt.Printf("error occurs: %v", err)
		}

		if len(QueryAllTravelInfo.Data) > 0 {
			MockedEndTime = QueryAllTravelInfo.Data[0].EndTime
		} else {
			fmt.Printf("The corresponding database is empty")
			MockedEndTime = getRandomTime()
		}
	} else if r8 < 0.99 {
		// Creat itself
		MockedEndTime = getRandomTime()
	} else {
		MockedEndTime = getRandomTime()
	}

	// Input
	updateTravelInfo := service.TravelInfo{
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
	_, err3 := travelSvc.UpdateTrip(&updateTravelInfo)
	if err3 != nil {
		fmt.Printf("error3 occurs: %v", err3)
	}
	time.Sleep(2 * time.Second)

	// 5. Delete according to the ID
	// Question: Is te ID here the UUID ID or the ID like 'G777'?
	var MockedDeleteID string
	r9 := rand.Float64()
	if r9 < 0.95 {
		QueryAllTravelInfo, err := travelSvc.QueryAll()
		if err != nil {
			fmt.Printf("error occurs: %v", err)
		}

		if len(QueryAllTravelInfo.Data) > 0 {
			MockedDeleteID = QueryAllTravelInfo.Data[0].Id
		} else {
			fmt.Printf("The corresponding database is empty")
			MockedDeleteID = faker.UUIDHyphenated()
		}
	} else if r9 < 0.99 {
		// Create And Query
		MockedLoginId = faker.UUIDHyphenated()
		MockedTripID = GenerateTripId()
		MockedTrainTypeName = GenerateTrainTypeName()
		MockedRouteID = faker.UUIDHyphenated()
		MockedStartStationName = faker.GetRealAddress().City
		MockedTerminalStationName = faker.GetRealAddress().City
		MockedStationsName = MockedStartStationName + ", " + MockedTerminalStationName
		MockedStartTime = faker.Date()
		MockedEndTime = faker.Date()
		// Input
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
		_, error := travelSvc.CreateTrip(&travelInfo)
		if error != nil {
			fmt.Printf("Error occurs: %v", error)
		}

		QueryAllTravelInfo, err := travelSvc.QueryAll()
		if err != nil {
			fmt.Printf("error occurs: %v", err)
		}

		if len(QueryAllTravelInfo.Data) > 0 {
			MockedDeleteID = QueryAllTravelInfo.Data[len(QueryAllTravelInfo.Data)-1].Id
		} else {
			fmt.Printf("The above CRATE Fails and the corresponding database is empty")
			MockedDeleteID = faker.UUIDHyphenated()
		}
	} else {
		MockedDeleteID = faker.UUIDHyphenated()
	}

	_, err4 := travelSvc.DeleteTrip(MockedDeleteID)
	if err4 != nil {
		fmt.Printf("error4 occurs: %v", err4)
	}
	time.Sleep(2 * time.Second)

	// 6. Retrieve by Trip ID & 7. GetTrainTypeByTripId & // 8. GetRouteByTripId
	var GetTripID string
	r10 := rand.Float64()
	if r10 < 0.95 {
		QueryAllTravelInfo, err := travelSvc.QueryAll()
		if err != nil {
			fmt.Printf("error occurs: %v", err)
		}

		if len(QueryAllTravelInfo.Data) > 0 {
			GetTripID = *(QueryAllTravelInfo.Data[0].TripId.Type) + QueryAllTravelInfo.Data[0].TripId.Number
		} else {
			fmt.Printf("The corresponding database is empty")
			GetTripID = GenerateTripId()
		}
	} else if r10 < 0.99 {
		// Create And Query
		MockedLoginId = faker.UUIDHyphenated()
		MockedTripID = GenerateTripId()
		MockedTrainTypeName = GenerateTrainTypeName()
		MockedRouteID = faker.UUIDHyphenated()
		MockedStartStationName = faker.GetRealAddress().City
		MockedTerminalStationName = faker.GetRealAddress().City
		MockedStationsName = MockedStartStationName + ", " + MockedTerminalStationName
		MockedStartTime = faker.Date()
		MockedEndTime = faker.Date()
		// Input
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
		_, error := travelSvc.CreateTrip(&travelInfo)
		if error != nil {
			fmt.Printf("Error occurs: %v", error)
		}

		QueryAllTravelInfo, err := travelSvc.QueryAll()
		if err != nil {
			fmt.Printf("error occurs: %v", err)
		}

		if len(QueryAllTravelInfo.Data) > 0 {
			GetTripID = *(QueryAllTravelInfo.Data[len(QueryAllTravelInfo.Data)-1].TripId.Type) + QueryAllTravelInfo.Data[len(QueryAllTravelInfo.Data)-1].TripId.Number
		} else {
			fmt.Printf("The above CRATE Fails and the corresponding database is empty")
			GetTripID = GenerateTripId()
		}
	} else {
		GetTripID = GenerateTripId()
	}

	// 6. Retrieve by Trip ID
	_, err5 := travelSvc.RetrieveTravel(GetTripID)
	if err5 != nil {
		fmt.Printf("error5 occurs: %v", err5)
	}
	time.Sleep(2 * time.Second)

	// 7. GetTrainTypeByTripId
	_, err6 := travelSvc.GetTrainTypeByTripId(GetTripID)
	if err6 != nil {
		fmt.Printf("error6 occurs: %v", err6)
	}
	time.Sleep(2 * time.Second)

	// 8. GetRouteByTripId
	_, err7 := travelSvc.GetRouteByTripId(GetTripID)
	if err7 != nil {
		fmt.Printf("error7 occurs: %v", err7)
	}
	time.Sleep(2 * time.Second)

	// 9. GetTripsByRouteId
	var GetRouteIDs []string
	r11 := rand.Float64()
	if r11 < 0.95 {
		GetAllRouteInfo, err := routeSvc.QueryAllRoutes()
		if err != nil {
			fmt.Printf("error occurs: %v", err)
		}

		if len(GetAllRouteInfo.Data) > 0 {
			GetRouteIDs = GetAllRouteInfo.Data[0].Stations
		} else {
			fmt.Printf("The corresponding database is empty")
			GetRouteIDs = []string{faker.UUIDHyphenated(), faker.UUIDHyphenated(), faker.UUIDHyphenated()}
		}
	} else if r11 < 0.99 {
		// Create And Query
		MockedRouteInfoID := faker.UUIDHyphenated()
		MockedRouteInfoStartStation := faker.GetRealAddress().City
		MockedRouteInfoEndStation := faker.GetRealAddress().City
		MockedStationList := MockedRouteInfoStartStation + ", " + faker.GetRealAddress().City + ", " + MockedRouteInfoEndStation
		MockedDistanceList := fmt.Sprintf("%d, %d, %d", rand.Intn(1000), rand.Intn(1000), rand.Intn(1000))
		CreateAndModifyRouteInput := service.RouteInfo{
			ID:           MockedRouteInfoID,
			StartStation: MockedRouteInfoStartStation,
			EndStation:   MockedRouteInfoEndStation,
			StationList:  MockedStationList,
			DistanceList: MockedDistanceList,
		}

		_, err := routeSvc.CreateAndModifyRoute(&CreateAndModifyRouteInput)
		if err != nil {
			fmt.Printf("error occurs: %v", err)
		}

		GetAllRouteInfo, err1 := routeSvc.QueryAllRoutes()
		if err1 != nil {
			fmt.Printf("error1 occurs: %v", err1)
		}

		if len(GetAllRouteInfo.Data) > 0 {
			GetRouteIDs = GetAllRouteInfo.Data[len(GetAllRouteInfo.Data)-1].Stations
		} else {
			fmt.Printf("The CRATE above fails and the corresponding database is empty")
			GetRouteIDs = []string{faker.UUIDHyphenated(), faker.UUIDHyphenated(), faker.UUIDHyphenated()}
		}
	} else {
		GetRouteIDs = []string{faker.UUIDHyphenated(), faker.UUIDHyphenated(), faker.UUIDHyphenated()}
	}
	_, err8 := travelSvc.GetTripsByRouteId(GetRouteIDs)
	if err8 != nil {
		fmt.Printf("error8 occurs: %v", err8)
	}
	time.Sleep(2 * time.Second)

	// 10. QueryInfo & 11. QueryInfoInParallel
	// Mock Input
	// 10.1. StartPlace
	// Service
	// Route Service
	var MockedStartPlace string
	r12 := rand.Float64()
	if r12 < 0.95 {
		GetAllRouteInfo, err := routeSvc.QueryAllRoutes()
		if err != nil {
			fmt.Printf("error occurs: %v", err)
		}

		if len(GetAllRouteInfo.Data) > 0 {
			MockedStartPlace = GetAllRouteInfo.Data[0].StartStation
		} else {
			fmt.Printf("The corresponding database is empty")
			MockedStartPlace = faker.GetRealAddress().City
		}
	} else if r12 < 0.99 {
		MockedRouteInfoID := faker.UUIDHyphenated()
		MockedRouteInfoStartStation := faker.GetRealAddress().City
		MockedRouteInfoEndStation := faker.GetRealAddress().City
		MockedStationList := MockedRouteInfoStartStation + ", " + faker.GetRealAddress().City + ", " + MockedRouteInfoEndStation
		MockedDistanceList := fmt.Sprintf("%d, %d, %d", rand.Intn(1000), rand.Intn(1000), rand.Intn(1000))
		CreateAndModifyRouteInput := service.RouteInfo{
			ID:           MockedRouteInfoID,
			StartStation: MockedRouteInfoStartStation,
			EndStation:   MockedRouteInfoEndStation,
			StationList:  MockedStationList,
			DistanceList: MockedDistanceList,
		}

		_, err := routeSvc.CreateAndModifyRoute(&CreateAndModifyRouteInput)
		if err != nil {
			fmt.Printf("error occurs: %v", err)
		}

		GetAllRouteInfo, err := routeSvc.QueryAllRoutes()
		if err != nil {
			fmt.Printf("error occurs: %v", err)
		}

		MockedStartPlace = GetAllRouteInfo.Data[len(GetAllRouteInfo.Data)-1].StartStation
	} else {
		MockedStartPlace = faker.GetRealAddress().City
	}

	// 10.2. EndPlace
	// Service
	// Route Service
	var MockedEndPlace string
	r13 := rand.Float64()
	if r13 < 0.95 {
		GetAllRouteInfo, err := routeSvc.QueryAllRoutes()
		if err != nil {
			fmt.Printf("error occurs: %v", err)
		}

		if len(GetAllRouteInfo.Data) > 0 {
			MockedEndPlace = GetAllRouteInfo.Data[0].EndStation
		} else {
			fmt.Printf("The corresponding database is empty")
			MockedEndPlace = faker.GetRealAddress().City
		}
	} else if r13 < 0.99 {
		MockedRouteInfoID := faker.UUIDHyphenated()
		MockedRouteInfoStartStation := faker.GetRealAddress().City
		MockedRouteInfoEndStation := faker.GetRealAddress().City
		MockedStationList := MockedRouteInfoStartStation + ", " + faker.GetRealAddress().City + ", " + MockedRouteInfoEndStation
		MockedDistanceList := fmt.Sprintf("%d, %d, %d", rand.Intn(1000), rand.Intn(1000), rand.Intn(1000))
		CreateAndModifyRouteInput := service.RouteInfo{
			ID:           MockedRouteInfoID,
			StartStation: MockedRouteInfoStartStation,
			EndStation:   MockedRouteInfoEndStation,
			StationList:  MockedStationList,
			DistanceList: MockedDistanceList,
		}

		_, err := routeSvc.CreateAndModifyRoute(&CreateAndModifyRouteInput)
		if err != nil {
			fmt.Printf("error occurs: %v", err)
		}

		GetAllRouteInfo, err := routeSvc.QueryAllRoutes()
		if err != nil {
			fmt.Printf("error occurs: %v", err)
		}

		MockedEndPlace = GetAllRouteInfo.Data[len(GetAllRouteInfo.Data)-1].EndStation
	} else {
		MockedEndPlace = faker.GetRealAddress().City
	}

	// 10.3. DepartureTime
	// Service
	// Travel Service
	var MockedDepartureTime string
	r14 := rand.Float64()
	if r14 < 0.95 {
		QueryAllTravelInfo, err := travelSvc.QueryAll()
		if err != nil {
			fmt.Printf("error occurs: %v", err)
		}

		if len(QueryAllTravelInfo.Data) > 0 {
			MockedDepartureTime = QueryAllTravelInfo.Data[0].StartTime
		} else {
			fmt.Printf("The corresponding database is empty")
			MockedDepartureTime = getRandomTime()
		}
	} else if r14 < 0.99 {
		// Create itself
		MockedDepartureTime = getRandomTime()
	} else {
		MockedDepartureTime = getRandomTime()
	}

	// Input
	MockedTripInfo := service.TripInfo{
		StartPlace:    MockedStartPlace,
		EndPlace:      MockedEndPlace,
		DepartureTime: MockedDepartureTime,
	}

	// 10. QueryInfo
	_, err9 := travelSvc.QueryInfo(MockedTripInfo)
	if err9 != nil {
		fmt.Printf("error9 occurs: %v", err9)
	}
	time.Sleep(2 * time.Second)

	// 11. QueryInfoInParallel
	_, err10 := travelSvc.QueryInfoInParallel(MockedTripInfo)
	if err10 != nil {
		fmt.Printf("error10 occurs: %v", err10)
	}
	time.Sleep(2 * time.Second)

	// 12. GetTripAllDetailInfo
	// Mock input
	// Service
	// Travel Service
	// TripID
	r15 := rand.Float64()
	if r15 < 0.95 {
		QueryAllTravelInfo, err := travelSvc.QueryAll()
		if err != nil {
			fmt.Printf("error occurs: %v", err)
		}

		if len(QueryAllTravelInfo.Data) > 0 {
			MockedTripID = *(QueryAllTravelInfo.Data[0].TripId.Type) + QueryAllTravelInfo.Data[0].TripId.Number
		} else {
			fmt.Printf("The corresponding database is empty")
			MockedTripID = GenerateTripId()
		}
	} else if r15 < 0.99 {
		// Create itself
		MockedTripID = GenerateTripId()
	} else {
		MockedTripID = GenerateTripId()
	}

	// input
	MockedtripAllDetailInfo := service.TripAllDetailInfo{
		TripId: MockedTripID,
	}

	_, err11 := travelSvc.GetTripAllDetailInfo(MockedtripAllDetailInfo)
	if err11 != nil {
		fmt.Printf("error11 occurs: %v", err11)
	}
	time.Sleep(2 * time.Second)

	// 13. AdminQueryAll
	_, err12 := travelSvc.AdminQueryAll()
	if err12 != nil {
		fmt.Printf("error12 occurs: %v", err12)
	}
	time.Sleep(2 * time.Second)
}

// helper function
func GenerateTripId() string {
	// 设置随机数种子
	rand.Seed(time.Now().UnixNano())

	// 定义可能的开头字母
	letters := []rune{'Z', 'T', 'K', 'G', 'D'}

	// 随机选择一个字母
	startLetter := letters[rand.Intn(len(letters))]

	// 生成三个随机数字
	randomNumber := rand.Intn(1000)

	// 格式化成三位数字，不足三位前面补零
	MockedTripID := fmt.Sprintf("%c%03d", startLetter, randomNumber)

	return MockedTripID
}

func GenerateTrainTypeName() string {
	// 设置随机数种子
	rand.Seed(time.Now().UnixNano())

	// 定义可能的火车类型名称
	trainTypes := []string{"GaoTieOne", "GaoTieTwo", "GaoTieSeven", "DongCheOne", "DongCheTen"}

	// 随机选择一个火车类型名称
	MockedTrainTypeName := trainTypes[rand.Intn(len(trainTypes))]

	return MockedTrainTypeName
}

func ListToString(stations []string) string {

	// Use a builder for efficient string concatenation
	var builder strings.Builder

	for i, station := range stations {
		if i > 0 {
			builder.WriteString(", ")
		}
		builder.WriteString(fmt.Sprintf("Stations[%d] %s", i, station))
	}

	result := builder.String()
	return result
}

func StringToList(input string) []string {
	// Split the input string by commas and trim any leading/trailing spaces from each element
	parts := strings.Split(input, ",")
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}
	return parts
}

func getRandomTime() string {
	randomDate := faker.Date()
	randomTime := faker.TIME

	DateAndTime := randomDate + " " + randomTime

	return DateAndTime
}
