package service

//
//import (
//	"testing"
//)
//
//func TestOrderService_GetTicketListByDateAndTripId(t *testing.T) {
//	cli, _ := GetAdminClient()
//	seatRequest := Seat{
//		TravelDate: "2023-06-14",
//		TripId:     "D1234",
//	}
//	result, err := cli.GetTicketListByDateAndTripId(seatRequest)
//	if err != nil {
//		t.Errorf("Request failed, err %s", err)
//	}
//	if len(result) == 0 {
//		t.Errorf("No tickets found")
//	}
//}
//
//func TestOrderService_CreateNewOrder(t *testing.T) {
//	cli, _ := GetAdminClient()
//	order := Order{
//		From:       "StationA",
//		To:         "StationB",
//		TravelDate: "2023-06-14",
//	}
//	result, err := cli.CreateNewOrder(order)
//	if err != nil {
//		t.Errorf("Request failed, err %s", err)
//	}
//	if result.Id == "" {
//		t.Errorf("Order creation failed")
//	}
//}
//
//func TestOrderService_AddNewOrder(t *testing.T) {
//	cli, _ := GetAdminClient()
//	order := Order{
//		From:       "StationA",
//		To:         "StationB",
//		TravelDate: "2023-06-14",
//	}
//	result, err := cli.AddNewOrder(order)
//	if err != nil {
//		t.Errorf("Request failed, err %s", err)
//	}
//	if result.Id == "" {
//		t.Errorf("Order creation failed")
//	}
//}
//
//func TestOrderService_QueryOrders(t *testing.T) {
//	cli, _ := GetAdminClient()
//	orderInfo := OrderInfo{
//		LoginId: "user123",
//	}
//	result, err := cli.QueryOrders(orderInfo)
//	if err != nil {
//		t.Errorf("Request failed, err %s", err)
//	}
//	if len(result) == 0 {
//		t.Errorf("No orders found")
//	}
//}
//
//func TestOrderService_QueryOrdersForRefresh(t *testing.T) {
//	cli, _ := GetAdminClient()
//	orderInfo := OrderInfo{
//		LoginId: "user123",
//	}
//	result, err := cli.QueryOrdersForRefresh(orderInfo)
//	if err != nil {
//		t.Errorf("Request failed, err %s", err)
//	}
//	if len(result) == 0 {
//		t.Errorf("No orders found")
//	}
//}
//
//func TestOrderService_CalculateSoldTicket(t *testing.T) {
//	cli, _ := GetAdminClient()
//	travelDate := "2023-06-14"
//	trainNumber := "D1234"
//	result, err := cli.CalculateSoldTicket(travelDate, trainNumber)
//	if err != nil {
//		t.Errorf("Request failed, err %s", err)
//	}
//	if result == 0 {
//		t.Errorf("No tickets sold")
//	}
//}
//
//func TestOrderService_GetOrderPrice(t *testing.T) {
//	cli, _ := GetAdminClient()
//	orderId := "order123"
//	result, err := cli.GetOrderPrice(orderId)
//	if err != nil {
//		t.Errorf("Request failed, err %s", err)
//	}
//	if result == 0 {
//		t.Errorf("Order price not found")
//	}
//}
//
//func TestOrderService_PayOrder(t *testing.T) {
//	cli, _ := GetAdminClient()
//	orderId := "order123"
//	result, err := cli.PayOrder(orderId)
//	if err != nil {
//		t.Errorf("Request failed, err %s", err)
//	}
//	if result.Id == "" {
//		t.Errorf("Order payment failed")
//	}
//}
//
//func TestOrderService_GetOrderById(t *testing.T) {
//	cli, _ := GetAdminClient()
//	orderId := "order123"
//	result, err := cli.GetOrderById(orderId)
//	if err != nil {
//		t.Errorf("Request failed, err %s", err)
//	}
//	if result.Id == "" {
//		t.Errorf("Order not found")
//	}
//}
//
//func TestOrderService_ModifyOrder(t *testing.T) {
//	cli, _ := GetAdminClient()
//	orderId := "order123"
//	status := 1
//	result, err := cli.ModifyOrder(orderId, status)
//	if err != nil {
//		t.Errorf("Request failed, err %s", err)
//	}
//	if result.Id == "" {
//		t.Errorf("Order modification failed")
//	}
//}
//
//func TestOrderService_SecurityInfoCheck(t *testing.T) {
//	cli, _ := GetAdminClient()
//	checkDate := "2023-06-14"
//	accountId := "account123"
//	result, err := cli.SecurityInfoCheck(checkDate, accountId)
//	if err != nil {
//		t.Errorf("Request failed, err %s", err)
//	}
//	if !result {
//		t.Errorf("Security check failed")
//	}
//}
//
//func TestOrderService_SaveOrderInfo(t *testing.T) {
//	cli, _ := GetAdminClient()
//	orderInfo := Order{
//		Id: "order123",
//	}
//	result, err := cli.SaveOrderInfo(orderInfo)
//	if err != nil {
//		t.Errorf("Request failed, err %s", err)
//	}
//	if result.Id == "" {
//		t.Errorf("Order save failed")
//	}
//}
//
//func TestOrderService_UpdateOrder(t *testing.T) {
//	cli, _ := GetAdminClient()
//	order := Order{
//		Id: "order123",
//	}
//	result, err := cli.UpdateOrder(order)
//	if err != nil {
//		t.Errorf("Request failed, err %s", err)
//	}
//	if result.Id == "" {
//		t.Errorf("Order update failed")
//	}
//}
//
//func TestOrderService_DeleteOrder(t *testing.T) {
//	cli, _ := GetAdminClient()
//	orderId := "order123"
//	result, err := cli.DeleteOrder(orderId)
//	if err != nil {
//		t.Errorf("Request failed, err %s", err)
//	}
//	if result.Id == "" {
//		t.Errorf("Order deletion failed")
//	}
//}
//
//func TestOrderService_GetAllOrders(t *testing.T) {
//	cli, _ := GetAdminClient()
//	result, err := cli.GetAllOrders()
//	if err != nil {
//		t.Errorf("Request failed, err %s", err)
//	}
//	if len(result) == 0 {
//		t.Errorf("No orders found")
//	}
//}
