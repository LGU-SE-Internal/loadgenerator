package behaviors

import (
	"fmt"
	"github.com/Lincyaw/loadgenerator/service"
)

const ()

var PreserveChain *Chain

func init() {
	PreserveChain = NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		return nil, nil
	}))
	PreserveChain.AddNextChain(NewChain(NewFuncNode(LoginAdmin)), 1)
}

func Preserve(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}
	OrderTicketsInfo := service.OrderTicketsInfo{
		AccountID:       MockedAccountID,
		ContactsID:      MockedContactsID,
		TripID:          MockedTripID,
		SeatType:        MockedSeatType,
		LoginToken:      ctx.Get(LoginId).(string),
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
	PreserveResp, err := cli.Preserve(&OrderTicketsInfo)
	if err != nil {
		return nil, err
	}
	if PreserveResp.Status != 1 {
		return nil, fmt.Errorf("preserve order tickets fail. PreserveResp.Status != 1, get %v", PreserveResp.Status)
	}
	//return nil, err
	fmt.Printf("The Status is: %v, and PreserveResp Data: %v\n", PreserveResp.Status, PreserveResp.Data)
	return &(NodeResult{false}), nil
}
