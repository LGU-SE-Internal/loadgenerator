package behaviors

import (
	"fmt"
	"strings"

	"github.com/Lincyaw/loadgenerator/service"
	"github.com/go-faker/faker/v4"
	log "github.com/sirupsen/logrus"
)

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
		DepartureTime: ctx.Get(DepartureTime).(string), // 生成1小时到1天之后的时间
	}

	var basicSvc service.BasicService = cli
	travel, err := basicSvc.QueryForTravel(travelQuery)
	if err != nil {
		log.Errorf("QueryForTravel failed: request=%+v, error=%v", travelQuery, err)
		return nil, err
	}
	if travel.Status != 1 {
		log.Errorf("QueryForTravel returned abnormal status: expected=1, got=%d, request=%+v, response=%+v", travel.Status, travelQuery, travel)
		return nil, fmt.Errorf("unexpected travel status: %d", travel.Status)
	}

	ctx.Set(Status, travel.Data.Status)
	ctx.Set(Percent, travel.Data.Percent)
	//ctx.Set(TrainTypeName, travel.Data.TrainType)
	ctx.Set(Route, travel.Data.Route)
	ctx.Set(Prices, travel.Data.Prices)
	//
	//ctx.Set(DepartureTime, travelQuery.DepartureTime)
	ctx.Set(TrainTypeName, travelQuery.Trip.TrainTypeName)

	return nil, nil
}
