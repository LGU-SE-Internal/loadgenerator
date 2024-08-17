package behaviors

import (
	"fmt"
	"github.com/Lincyaw/loadgenerator/service"
	"github.com/go-faker/faker/v4"
	log "github.com/sirupsen/logrus"
	"strings"
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
