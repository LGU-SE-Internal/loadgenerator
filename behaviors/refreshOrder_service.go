package behaviors

import (
	"fmt"
	"github.com/Lincyaw/loadgenerator/service"
	log "github.com/sirupsen/logrus"
	"math/rand"
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
	//if rand.Intn(2) == 0 {
	randomIndex = rand.Intn(len(RefreshResp.Data))
	ctx.Set(TrainNumber, RefreshResp.Data[randomIndex].TrainNumber)
	ctx.Set(Price, RefreshResp.Data[randomIndex].Price)
	ctx.Set(OrderId, RefreshResp.Data[randomIndex].Id) // ID here is exactly the OrderId
	//} else {
	//	randomIndex = rand.Intn(len(RefreshOtherResp.Data))
	//	ctx.Set(TrainNumber, RefreshOtherResp.Data[randomIndex].TrainNumber)
	//	ctx.Set(Price, RefreshOtherResp.Data[randomIndex].Price)
	//	ctx.Set(OrderId, RefreshOtherResp.Data[randomIndex].Id)
	//}

	return nil, nil
}
