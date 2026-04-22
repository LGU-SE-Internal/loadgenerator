package behaviors

import (
	"fmt"
	"math/rand"

	"github.com/Lincyaw/loadgenerator/service"
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

// Admin Basic Info Behavior Chain - 管理联系人、车站、列车、配置、价格等基础信息
var AdminBasicInfoChain *Chain

func init() {
	AdminBasicInfoChain = NewChain(
		NewFuncNode(LoginAdmin, "LoginAdmin"),
		NewFuncNode(AdminQueryAllContacts, "AdminQueryAllContacts"),
		NewFuncNode(AdminQueryAllStations, "AdminQueryAllStations"),
		NewFuncNode(AdminQueryAllTrains, "AdminQueryAllTrains"),
		NewFuncNode(AdminQueryAllConfigs, "AdminQueryAllConfigs"),
		NewFuncNode(AdminQueryAllPrices, "AdminQueryAllPrices"),
	)

	// 添加后续的管理操作链
	AdminBasicInfoChain.AddNextChain(NewChain(
		NewFuncNode(AdminAddContact, "AdminAddContact"),
		NewFuncNode(AdminModifyContact, "AdminModifyContact"),
	), 0.3)

	AdminBasicInfoChain.AddNextChain(NewChain(
		NewFuncNode(AdminAddStation, "AdminAddStation"),
		NewFuncNode(AdminModifyStation, "AdminModifyStation"),
	), 0.3)

	AdminBasicInfoChain.AddNextChain(NewChain(
		NewFuncNode(AdminAddTrain, "AdminAddTrain"),
		NewFuncNode(AdminModifyTrain, "AdminModifyTrain"),
	), 0.2)

	AdminBasicInfoChain.AddNextChain(NewChain(
		NewFuncNode(AdminAddConfig, "AdminAddConfig"),
		NewFuncNode(AdminModifyConfig, "AdminModifyConfig"),
	), 0.2)
}

// AdminQueryAllContacts 查询所有联系人
func AdminQueryAllContacts(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	resp, err := cli.AdminGetAllContacts()
	if err != nil {
		log.Errorf("AdminGetAllContacts failed: %v", err)
		return nil, err
	}

	if resp.Status != 1 {
		log.Warnf("AdminGetAllContacts returned status: %d, msg: %s", resp.Status, resp.Msg)
	}

	if len(resp.Data) > 0 {
		randomIndex := rand.Intn(len(resp.Data))
		ctx.Set(ContactsID, resp.Data[randomIndex].Id)
		ctx.Set(AccountID, resp.Data[randomIndex].AccountId)
		ctx.Set(Name, resp.Data[randomIndex].Name)
		ctx.Set(DocumentType, resp.Data[randomIndex].DocumentType)
		ctx.Set(DocumentNumber, resp.Data[randomIndex].DocumentNumber)
		ctx.Set(PhoneNumber, resp.Data[randomIndex].PhoneNumber)
	}

	log.Infof("AdminGetAllContacts returned %d contacts", len(resp.Data))
	return nil, nil
}

// AdminAddContact 添加联系人
func AdminAddContact(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	contact := &service.AdminContacts{
		Id:             uuid.New().String(),
		AccountId:      uuid.New().String(),
		Name:           faker.Name(),
		DocumentType:   rand.Intn(2),
		DocumentNumber: faker.CCNumber(),
		PhoneNumber:    faker.Phonenumber(),
	}

	resp, err := cli.AdminAddContact(contact)
	if err != nil {
		log.Errorf("AdminAddContact failed: %v", err)
		return nil, err
	}

	if resp.Status != 1 {
		log.Warnf("AdminAddContact returned status: %d, msg: %s", resp.Status, resp.Msg)
		return nil, nil
	}

	ctx.Set(ContactsID, resp.Data.Id)
	ctx.Set(AccountID, resp.Data.AccountId)
	ctx.Set(Name, resp.Data.Name)

	log.Infof("AdminAddContact success: contactId=%s", resp.Data.Id)
	return nil, nil
}

// AdminModifyContact 修改联系人
func AdminModifyContact(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	contactId, ok := ctx.Get(ContactsID).(string)
	if !ok || contactId == "" {
		log.Warn("No contact ID found in context, skipping modify")
		return nil, nil
	}

	contact := &service.AdminContacts{
		Id:             contactId,
		AccountId:      ctx.Get(AccountID).(string),
		Name:           faker.Name(),
		DocumentType:   rand.Intn(2),
		DocumentNumber: faker.CCNumber(),
		PhoneNumber:    faker.Phonenumber(),
	}

	resp, err := cli.AdminModifyContact(contact)
	if err != nil {
		log.Errorf("AdminModifyContact failed: %v", err)
		return nil, err
	}

	if resp.Status != 1 {
		log.Warnf("AdminModifyContact returned status: %d, msg: %s", resp.Status, resp.Msg)
	}

	log.Infof("AdminModifyContact success: contactId=%s", resp.Data.Id)
	return nil, nil
}

// AdminQueryAllStations 查询所有车站
func AdminQueryAllStations(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	resp, err := cli.AdminGetAllStations()
	if err != nil {
		log.Errorf("AdminGetAllStations failed: %v", err)
		return nil, err
	}

	if resp.Status != 1 {
		log.Warnf("AdminGetAllStations returned status: %d, msg: %s", resp.Status, resp.Msg)
	}

	log.Infof("AdminGetAllStations success")
	return nil, nil
}

// AdminAddStation 添加车站
func AdminAddStation(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	station := &service.AdminStation{
		ID:       uuid.New().String(),
		Name:     generateRandomCityName() + "_new",
		StayTime: rand.Intn(30) + 5,
	}

	resp, err := cli.AdminAddStation(station)
	if err != nil {
		log.Errorf("AdminAddStation failed: %v", err)
		return nil, err
	}

	if resp.Status != 1 {
		log.Warnf("AdminAddStation returned status: %d, msg: %s", resp.Status, resp.Msg)
		return nil, nil
	}

	ctx.Set(StationId, station.ID)
	ctx.Set(StationNames, station.Name)
	ctx.Set(StayTime, station.StayTime)

	log.Infof("AdminAddStation success: stationId=%s", station.ID)
	return nil, nil
}

// AdminModifyStation 修改车站
func AdminModifyStation(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	stationId, ok := ctx.Get(StationId).(string)
	if !ok || stationId == "" {
		log.Warn("No station ID found in context, skipping modify")
		return nil, nil
	}

	station := &service.AdminStation{
		ID:       stationId,
		Name:     generateRandomCityName() + "_modified",
		StayTime: rand.Intn(30) + 5,
	}

	resp, err := cli.AdminModifyStation(station)
	if err != nil {
		log.Errorf("AdminModifyStation failed: %v", err)
		return nil, err
	}

	if resp.Status != 1 {
		log.Warnf("AdminModifyStation returned status: %d, msg: %s", resp.Status, resp.Msg)
	}

	log.Infof("AdminModifyStation success: stationId=%s", stationId)
	return nil, nil
}

// AdminQueryAllTrains 查询所有列车类型
func AdminQueryAllTrains(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	resp, err := cli.AdminGetAllTrains()
	if err != nil {
		log.Errorf("AdminGetAllTrains failed: %v", err)
		return nil, err
	}

	if resp.Status != 1 {
		log.Warnf("AdminGetAllTrains returned status: %d, msg: %s", resp.Status, resp.Msg)
	}

	log.Infof("AdminGetAllTrains success")
	return nil, nil
}

// AdminAddTrain 添加列车类型
func AdminAddTrain(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	train := &service.AdminTrainType{
		ID:           uuid.New().String(),
		Name:         GenerateTrainTypeName(),
		EconomyClass: rand.Intn(500) + 100,
		ConfortClass: rand.Intn(200) + 50,
		AverageSpeed: rand.Intn(200) + 100,
	}

	resp, err := cli.AdminAddTrain(train)
	if err != nil {
		log.Errorf("AdminAddTrain failed: %v", err)
		return nil, err
	}

	if resp.Status != 1 {
		log.Warnf("AdminAddTrain returned status: %d, msg: %s", resp.Status, resp.Msg)
		return nil, nil
	}

	ctx.Set(TrainTypeName, train.Name)
	ctx.Set(EconomyClass, train.EconomyClass)
	ctx.Set(ConfortClass, train.ConfortClass)
	ctx.Set(AverageSpeed, train.AverageSpeed)

	log.Infof("AdminAddTrain success: trainId=%s, name=%s", train.ID, train.Name)
	return nil, nil
}

// AdminModifyTrain 修改列车类型
func AdminModifyTrain(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	train := &service.AdminTrainType{
		ID:           uuid.New().String(),
		Name:         GenerateTrainTypeName(),
		EconomyClass: rand.Intn(500) + 100,
		ConfortClass: rand.Intn(200) + 50,
		AverageSpeed: rand.Intn(200) + 100,
	}

	resp, err := cli.AdminModifyTrain(train)
	if err != nil {
		log.Errorf("AdminModifyTrain failed: %v", err)
		return nil, err
	}

	if resp.Status != 1 {
		log.Warnf("AdminModifyTrain returned status: %d, msg: %s", resp.Status, resp.Msg)
	}

	log.Infof("AdminModifyTrain success")
	return nil, nil
}

// AdminQueryAllConfigs 查询所有配置
func AdminQueryAllConfigs(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	resp, err := cli.AdminGetAllConfigs()
	if err != nil {
		log.Errorf("AdminGetAllConfigs failed: %v", err)
		return nil, err
	}

	if resp.Status != 1 {
		log.Warnf("AdminGetAllConfigs returned status: %d, msg: %s", resp.Status, resp.Msg)
	}

	log.Infof("AdminGetAllConfigs success")
	return nil, nil
}

// AdminAddConfig 添加配置
func AdminAddConfig(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	config := &service.AdminConfig{
		Name:        fmt.Sprintf("test_config_%s", uuid.New().String()[:8]),
		Value:       fmt.Sprintf("%d", rand.Intn(1000)),
		Description: generateDescription(),
	}

	resp, err := cli.AdminAddConfig(config)
	if err != nil {
		log.Errorf("AdminAddConfig failed: %v", err)
		return nil, err
	}

	if resp.Status != 1 {
		log.Warnf("AdminAddConfig returned status: %d, msg: %s", resp.Status, resp.Msg)
		return nil, nil
	}

	ctx.Set(ConfigName, config.Name)
	ctx.Set(Value, config.Value)
	ctx.Set(Description, config.Description)

	log.Infof("AdminAddConfig success: name=%s", config.Name)
	return nil, nil
}

// AdminModifyConfig 修改配置
func AdminModifyConfig(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	configName, ok := ctx.Get(ConfigName).(string)
	if !ok || configName == "" {
		log.Warn("No config name found in context, skipping modify")
		return nil, nil
	}

	config := &service.AdminConfig{
		Name:        configName,
		Value:       fmt.Sprintf("%d", rand.Intn(1000)),
		Description: generateDescription(),
	}

	resp, err := cli.AdminModifyConfig(config)
	if err != nil {
		log.Errorf("AdminModifyConfig failed: %v", err)
		return nil, err
	}

	if resp.Status != 1 {
		log.Warnf("AdminModifyConfig returned status: %d, msg: %s", resp.Status, resp.Msg)
	}

	log.Infof("AdminModifyConfig success: name=%s", configName)
	return nil, nil
}

// AdminQueryAllPrices 查询所有价格
func AdminQueryAllPrices(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	resp, err := cli.AdminGetAllPrices()
	if err != nil {
		log.Errorf("AdminGetAllPrices failed: %v", err)
		return nil, err
	}

	if resp.Status != 1 {
		log.Warnf("AdminGetAllPrices returned status: %d, msg: %s", resp.Status, resp.Msg)
	}

	log.Infof("AdminGetAllPrices success")
	return nil, nil
}

// AdminAddPrice 添加价格信息
func AdminAddPrice(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	price := &service.AdminPriceInfo{
		ID:                  uuid.New().String(),
		TrainType:           GenerateTrainTypeName(),
		RouteID:             uuid.New().String(),
		BasicPriceRate:      rand.Float64() * 10,
		FirstClassPriceRate: rand.Float64() * 20,
	}

	resp, err := cli.AdminAddPrice(price)
	if err != nil {
		log.Errorf("AdminAddPrice failed: %v", err)
		return nil, err
	}

	if resp.Status != 1 {
		log.Warnf("AdminAddPrice returned status: %d, msg: %s", resp.Status, resp.Msg)
		return nil, nil
	}

	ctx.Set(BasicPriceRate, price.BasicPriceRate)
	ctx.Set(FirstClassPriceRate, price.FirstClassPriceRate)

	log.Infof("AdminAddPrice success: priceId=%s", price.ID)
	return nil, nil
}

// AdminModifyPrice 修改价格信息
func AdminModifyPrice(ctx *Context) (*NodeResult, error) {
	cli, ok := ctx.Get(Client).(*service.SvcImpl)
	if !ok {
		return nil, fmt.Errorf("service client not found in context")
	}

	price := &service.AdminPriceInfo{
		ID:                  uuid.New().String(),
		TrainType:           GenerateTrainTypeName(),
		RouteID:             uuid.New().String(),
		BasicPriceRate:      rand.Float64() * 10,
		FirstClassPriceRate: rand.Float64() * 20,
	}

	resp, err := cli.AdminModifyPrice(price)
	if err != nil {
		log.Errorf("AdminModifyPrice failed: %v", err)
		return nil, err
	}

	if resp.Status != 1 {
		log.Warnf("AdminModifyPrice returned status: %d, msg: %s", resp.Status, resp.Msg)
	}

	log.Infof("AdminModifyPrice success")
	return nil, nil
}
