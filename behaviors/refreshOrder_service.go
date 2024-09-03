package behaviors

import (
	"fmt"
	"github.com/Lincyaw/loadgenerator/service"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"time"
)

const (
	State = "state"
)

func RefreshOrder(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	orderInfo := service.OrderInfo{
		BoughtDateEnd:         "",
		BoughtDateStart:       "",
		EnableBoughtDateQuery: false,
		EnableStateQuery:      true,
		EnableTravelDateQuery: false,
		LoginId:               ctx.Get(AccountID).(string),
		State:                 0,
		TravelDateEnd:         "",
		TravelDateStart:       "",
	}

	var orderSvc service.OrderService = cli
	RefreshResp, err := orderSvc.ReqQueryOrderForRefresh(&orderInfo)
	if err != nil {
		log.Errorf("Refresh Order Request failed, err %s", err)
		return nil, err
	}
	if RefreshResp.Status != 1 {
		log.Errorf("RefreshResp.Status != 1")
		return nil, err
	}

	orderOtherInfo := service.OrderInfo{
		BoughtDateEnd:         "",
		BoughtDateStart:       "",
		EnableBoughtDateQuery: false,
		EnableStateQuery:      true,
		EnableTravelDateQuery: false,
		LoginId:               ctx.Get(AccountID).(string),
		State:                 1,
		TravelDateEnd:         "",
		TravelDateStart:       "",
	}

	var orderOtherSvc service.OrderOtherService = cli
	RefreshOtherResp, errOther := orderOtherSvc.ReqQueryOrderForRefreshOther(&orderOtherInfo)
	if errOther != nil {
		log.Errorf("Refresh Other Order Request failed, errOther %s", errOther)
		return nil, errOther
	}
	if RefreshOtherResp.Status != 1 {
		log.Errorf("RefreshOtherResp.Status != 1")
		return nil, errOther
	}

	var randomIndex int
	if len(RefreshResp.Data) == 0 {
		log.Errorf("Unpaied order is empty")
		return &NodeResult{Continue: false}, nil
	}
	//if rand.Intn(2) == 0 {
	randomIndex = rand.Intn(len(RefreshResp.Data))
	ctx.Set(TrainNumber, RefreshResp.Data[randomIndex].TrainNumber) // OldTripId
	ctx.Set(OldTripID, RefreshResp.Data[randomIndex].TrainNumber)   // OldTripId
	ctx.Set(Price, RefreshResp.Data[randomIndex].Price)
	ctx.Set(OrderId, RefreshResp.Data[randomIndex].Id) // ID here is exactly the OrderId
	ctx.Set(From, RefreshResp.Data[randomIndex].From)
	ctx.Set(To, RefreshResp.Data[randomIndex].To)

	TheBoughtDate, err := time.Parse("2006-01-02 15:04:05", RefreshResp.Data[randomIndex].BoughtDate)
	if err != nil {
		log.Errorf("TheBoughtDate Transformation is failed, err %s", err)
	}
	formattedBoughtDate := TheBoughtDate.Format("2006-01-02")
	ctx.Set(HandleDate, formattedBoughtDate)

	TheTravelDate, err := time.Parse("2006-01-02", RefreshResp.Data[randomIndex].TravelDate)
	if err != nil {
		log.Errorf("TheTravelDate Transformation is failed, err %s", err)
	}
	formattedTravelDate := TheTravelDate.Format("2006-01-02")
	ctx.Set(TargetDate, formattedTravelDate)
	//} else {
	//	randomIndex = rand.Intn(len(RefreshOtherResp.Data))
	//	ctx.Set(TrainNumber, RefreshOtherResp.Data[randomIndex].TrainNumber)
	//	ctx.Set(Price, RefreshOtherResp.Data[randomIndex].Price)
	//	ctx.Set(OrderId, RefreshOtherResp.Data[randomIndex].Id)
	//}

	return nil, nil
}

func RefreshOrderOther(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	orderInfo := service.OrderInfo{
		BoughtDateEnd:         "",
		BoughtDateStart:       "",
		EnableBoughtDateQuery: false,
		EnableStateQuery:      true,
		EnableTravelDateQuery: false,
		LoginId:               ctx.Get(AccountID).(string),
		State:                 0,
		TravelDateEnd:         "",
		TravelDateStart:       "",
	}

	var orderSvc service.OrderService = cli
	RefreshResp, err := orderSvc.ReqQueryOrderForRefresh(&orderInfo)
	if err != nil {
		log.Errorf("Refresh Order Request failed, err %s", err)
		return nil, err
	}
	if RefreshResp.Status != 1 {
		log.Errorf("RefreshResp.Status != 1")
		return nil, err
	}

	orderOtherInfo := service.OrderInfo{
		BoughtDateEnd:         "",
		BoughtDateStart:       "",
		EnableBoughtDateQuery: false,
		EnableStateQuery:      true,
		EnableTravelDateQuery: false,
		LoginId:               ctx.Get(AccountID).(string),
		State:                 1,
		TravelDateEnd:         "",
		TravelDateStart:       "",
	}

	var orderOtherSvc service.OrderOtherService = cli
	RefreshOtherResp, errOther := orderOtherSvc.ReqQueryOrderForRefreshOther(&orderOtherInfo)
	if errOther != nil {
		log.Errorf("Refresh Other Order Request failed, errOther %s", errOther)
		return nil, errOther
	}
	if RefreshOtherResp.Status != 1 {
		log.Errorf("RefreshOtherResp.Status != 1")
		return nil, errOther
	}

	var randomIndex int
	if len(RefreshOtherResp.Data) == 0 {
		log.Errorf("Paied order is empty")
		return &NodeResult{Continue: false}, nil
	}
	//if rand.Intn(2) == 0 {
	randomIndex = rand.Intn(len(RefreshOtherResp.Data))
	ctx.Set(TrainNumber, RefreshOtherResp.Data[randomIndex].TrainNumber) // OldTripId
	ctx.Set(OldTripID, RefreshOtherResp.Data[randomIndex].TrainNumber)   // OldTripId
	ctx.Set(Price, RefreshOtherResp.Data[randomIndex].Price)
	ctx.Set(OrderId, RefreshOtherResp.Data[randomIndex].Id) // ID here is exactly the OrderId
	ctx.Set(From, RefreshOtherResp.Data[randomIndex].From)
	ctx.Set(To, RefreshOtherResp.Data[randomIndex].To)

	TheBoughtDate, err := time.Parse("2006-01-02 15:04:05", RefreshOtherResp.Data[randomIndex].BoughtDate)
	if err != nil {
		log.Errorf("TheBoughtDate Transformation is failed, err %s", err)
	}
	formattedBoughtDate := TheBoughtDate.Format("2006-01-02")
	ctx.Set(HandleDate, formattedBoughtDate)

	TheTravelDate, err := time.Parse("2006-01-02", RefreshOtherResp.Data[randomIndex].TravelDate)
	if err != nil {
		log.Errorf("TheTravelDate Transformation is failed, err %s", err)
	}
	formattedTravelDate := TheTravelDate.Format("2006-01-02")
	ctx.Set(TargetDate, formattedTravelDate)
	//} else {
	//	randomIndex = rand.Intn(len(RefreshOtherResp.Data))
	//	ctx.Set(TrainNumber, RefreshOtherResp.Data[randomIndex].TrainNumber)
	//	ctx.Set(Price, RefreshOtherResp.Data[randomIndex].Price)
	//	ctx.Set(OrderId, RefreshOtherResp.Data[randomIndex].Id)
	//}

	return nil, nil
}

func RefreshCollectedOrder(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	orderInfo := service.OrderInfo{
		BoughtDateEnd:         "",
		BoughtDateStart:       "",
		EnableBoughtDateQuery: false,
		EnableStateQuery:      false,
		EnableTravelDateQuery: false,
		LoginId:               ctx.Get(AccountID).(string),
		State:                 1,
		TravelDateEnd:         "",
		TravelDateStart:       "",
	}

	var orderSvc service.OrderService = cli
	RefreshResp, err := orderSvc.ReqQueryOrderForRefresh(&orderInfo)
	if err != nil {
		log.Errorf("Refresh Order Request failed, err %s", err)
		return nil, err
	}
	if RefreshResp.Status != 1 {
		log.Errorf("RefreshResp.Status != 1")
		return nil, err
	}

	orderOtherInfo := service.OrderInfo{
		BoughtDateEnd:         "",
		BoughtDateStart:       "",
		EnableBoughtDateQuery: false,
		EnableStateQuery:      false,
		EnableTravelDateQuery: false,
		LoginId:               ctx.Get(AccountID).(string),
		State:                 1,
		TravelDateEnd:         "",
		TravelDateStart:       "",
	}

	var orderOtherSvc service.OrderOtherService = cli
	RefreshOtherResp, errOther := orderOtherSvc.ReqQueryOrderForRefreshOther(&orderOtherInfo)
	if errOther != nil {
		log.Errorf("Refresh Other Order Request failed, errOther %s", errOther)
		return nil, errOther
	}
	if RefreshOtherResp.Status != 1 {
		log.Errorf("RefreshOtherResp.Status != 1")
		return nil, errOther
	}

	var randomIndex int
	if len(RefreshResp.Data) == 0 {
		log.Errorf("Unpaied order is empty")
		return &NodeResult{Continue: false}, nil
	}
	//if rand.Intn(2) == 0 {
	randomIndex = rand.Intn(len(RefreshResp.Data))
	ctx.Set(TrainNumber, RefreshResp.Data[randomIndex].TrainNumber) // OldTripId
	ctx.Set(OldTripID, RefreshResp.Data[randomIndex].TrainNumber)   // OldTripId
	ctx.Set(Price, RefreshResp.Data[randomIndex].Price)
	ctx.Set(OrderId, RefreshResp.Data[randomIndex].Id) // ID here is exactly the OrderId
	ctx.Set(From, RefreshResp.Data[randomIndex].From)
	ctx.Set(To, RefreshResp.Data[randomIndex].To)

	TheBoughtDate, err := time.Parse("2006-01-02 15:04:05", RefreshResp.Data[randomIndex].BoughtDate)
	if err != nil {
		log.Errorf("TheBoughtDate Transformation is failed, err %s", err)
	}
	formattedBoughtDate := TheBoughtDate.Format("2006-01-02")
	ctx.Set(HandleDate, formattedBoughtDate)

	TheTravelDate, err := time.Parse("2006-01-02", RefreshResp.Data[randomIndex].TravelDate)
	if err != nil {
		log.Errorf("TheTravelDate Transformation is failed, err %s", err)
	}
	formattedTravelDate := TheTravelDate.Format("2006-01-02")
	ctx.Set(TargetDate, formattedTravelDate)
	//} else {
	//	randomIndex = rand.Intn(len(RefreshOtherResp.Data))
	//	ctx.Set(TrainNumber, RefreshOtherResp.Data[randomIndex].TrainNumber)
	//	ctx.Set(Price, RefreshOtherResp.Data[randomIndex].Price)
	//	ctx.Set(OrderId, RefreshOtherResp.Data[randomIndex].Id)
	//}

	return nil, nil
}
