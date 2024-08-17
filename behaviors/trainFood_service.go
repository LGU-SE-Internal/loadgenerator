package behaviors

import (
	"fmt"
	"github.com/Lincyaw/loadgenerator/service"
	log "github.com/sirupsen/logrus"
	"math/rand"
)

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
