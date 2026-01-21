# Load Generator 设计思路文档

## 1. 概述

本 Load Generator 是一个基于 Go 语言开发的负载生成器，用于模拟 Train-Ticket 微服务系统（包含 41 个微服务）的真实用户行为流量。其核心设计目标是通过多种机制实现请求的**随机性**和**多样性**，从而产生更贴近真实场景的测试负载。

---

## 2. 核心架构设计

### 2.1 责任链模式 (Chain of Responsibility)

系统采用**责任链模式**作为核心架构，将复杂的用户行为分解为一系列可复用的节点（Node）。

```
Chain 结构:
┌─────────────────────────────────────────────────────────────┐
│  Chain                                                       │
│  ├── nodes: []Node          // 顺序执行的节点列表            │
│  ├── nextChains: []chainWithProbability  // 带概率的后继链   │
│  └── probabilitySum: float64            // 概率权重总和      │
└─────────────────────────────────────────────────────────────┘
```

**关键数据结构：**
- **Node**: 单个原子操作，如登录、查询车票、支付等
- **Chain**: 节点的组合，表示一个完整的业务流程
- **Context**: 上下文传递对象，用于在节点间传递数据

---

## 3. 随机性设计机制

### 3.1 行为链级别的概率选择

在 main.go 中，通过为不同的行为链设置不同的权重，实现行为级别的随机分布：

```go
composedChain.AddNextChain(behaviors.NormalPreserveChain, 30)        // 30% - 普通预订
composedChain.AddNextChain(behaviors.NormalOrderPayChain, 20)        // 20% - 订单支付
composedChain.AddNextChain(behaviors.OrderConsignChain, 10)          // 10% - 订单托运
composedChain.AddNextChain(behaviors.TicketCollectAndEnterStationChain, 10) // 10% - 取票进站
composedChain.AddNextChain(behaviors.AdvancedSearchChain, 20)        // 20% - 高级搜索
composedChain.AddNextChain(behaviors.ConsignListChain, 8)            // 8%  - 托运列表查询
composedChain.AddNextChain(behaviors.OrderChangeChain, 3)            // 3%  - 改签
composedChain.AddNextChain(behaviors.OrderCancelChain, 2)            // 2%  - 取消订单
```

**实现原理：**
```go
func (c *Chain) Execute(ctx *Context) (*NodeResult, error) {
    // ... 执行当前链的节点 ...

    if len(c.nextChains) > 0 {
        randValue := rand.Float64() * c.probabilitySum  // 生成随机值
        cumulative := 0.0
        for _, cp := range c.nextChains {
            cumulative += cp.probability
            if randValue <= cumulative {
                return cp.chain.Execute(ctx)  // 基于概率选择下一个链
            }
        }
    }
    return nil, nil
}
```

### 3.2 数据选择级别的随机性

在每个节点内部，数据选择同样采用随机机制：

#### 3.2.1 路线随机选择
```go
func ChooseRoute(ctx *Context) (*NodeResult, error) {
    allRoutes, _ := routeSvc.QueryAllRoutes()
    randomIndex := rand.Intn(len(allRoutes.Data))  // 随机选择一条路线
    ctx.Set(RouteID, allRoutes.Data[randomIndex].Id)
    ctx.Set(StartStation, allRoutes.Data[randomIndex].StartStation)
    ctx.Set(EndStation, allRoutes.Data[randomIndex].EndStation)
    // ...
}
```

#### 3.2.2 联系人随机选择
```go
func QueryContacts(ctx *Context) (*NodeResult, error) {
    GetAllContacts, _ := contactsSvc.GetContactByAccountId(TheAccountId)
    randomIndex := rand.Intn(len(GetAllContacts.Data))  // 随机选择联系人
    ctx.Set(ContactsID, GetAllContacts.Data[randomIndex].Id)
    // ...
}
```

#### 3.2.3 保险类型随机选择
```go
func QueryAssurance(ctx *Context) (*NodeResult, error) {
    Assurances, _ := cli.GetAllAssuranceTypes()
    randomIndex := rand.Intn(len(Assurances.Data))  // 随机选择保险类型
    ctx.Set(AssuranceTypeIndex, selectedAssurance.Index)
    // ...
}
```

#### 3.2.4 车次随机选择
```go
func QueryTripInfo(ctx *Context) (*NodeResult, error) {
    queryInfoResp, _ := travelSvc.QueryInfo(tripInfo)
    randomIndex := rand.Intn(len(queryInfoResp.Data))  // 随机选择车次
    ctx.Set(TripID, queryInfoResp.Data[randomIndex].TripId)
    // ...
}
```

---

## 4. 多样性设计机制

### 4.1 食物类型多样化
```go
func QueryFood(ctx *Context) (*NodeResult, error) {
    foodType := rand.Int()%2 + 1  // 1: 火车餐 或 2: 车站商店
    switch foodType {
    case 1:  // 火车餐食
        idx := rand.Intn(len(allFood.Data.TrainFoodList))
        ctx.Set(FoodName, allFood.Data.TrainFoodList[idx].FoodName)
    case 2:  // 车站商店食品
        for _, v := range allFood.Data.FoodStoreListMap {
            idx := rand.Intn(len(v))
            ctx.Set(FoodName, v[idx].FoodList[rand.Intn(len(v[idx].FoodList))].FoodName)
        }
    }
    ctx.Set(FoodType, foodType)
}
```

### 4.2 座位等级多样化
```go
func GetTrainTicketClass() int {
    probability := rand.Intn(100)
    switch {
    case probability < 5:   return 0  // 5%  商务座
    case probability < 20:  return 1  // 15% 一等座
    default:                return 2  // 80% 二等座
    }
}
```

### 4.3 车次类型多样化
```go
func GenerateTripId() string {
    letters := []rune{'Z', 'T', 'K', 'G', 'D'}  // 直达/特快/快速/高铁/动车
    startLetter := letters[rand.Intn(len(letters))]
    randomNumber := rand.Intn(1000)
    return fmt.Sprintf("%c%03d", startLetter, randomNumber)
}

func GenerateTrainTypeName() string {
    trainTypes := []string{"GaoTieOne", "GaoTieTwo", "DongCheOne", "ZhiDa", "TeKuai", "KuaiSu"}
    return trainTypes[rand.Intn(len(trainTypes))]
}
```

### 4.4 高级搜索策略多样化
```go
func TravelPlanAdvancedSearch(ctx *Context) (*NodeResult, error) {
    switch rand.Intn(3) {
    case 0:  // 最便宜策略
        Resp, _ = travelplanSvc.ReqGetByCheapest(&travelPlanInput)
    case 1:  // 最少经停站策略
        Resp, _ = travelplanSvc.ReqGetByMinStation(&travelPlanInput)
    case 2:  // 最快策略
        Resp, _ = travelplanSvc.ReqGetByQuickest(&travelPlanInput)
    }
}
```

### 4.5 用户身份多样化
```go
func init() {
    LoginChain.AddNextChain(NewChain(NewFuncNode(LoginAdmin, "LoginAdmin")), 0.2)      // 20% 管理员
    LoginChain.AddNextChain(NewChain(NewFuncNode(CreateUser, "CreateUser"),
                                     NewFuncNode(LoginNormal, "LoginNormal")), 0.8)  // 80% 普通用户
}
```

### 4.6 随机数据生成器

系统提供了多种随机数据生成函数，用于生成多样化的测试数据：

| 函数名 | 功能 | 多样性范围 |
|--------|------|-----------|
| `RandomProvincialCapitalEN()` | 随机省会城市 | 32个中国省会城市 |
| `generateRandomFood()` | 随机食物名称 | 50种食物 |
| `generateRandomStoreName()` | 随机商店名称 | 30种商店 |
| `generateRandomCityName()` | 随机城市名称 | 13个城市 |
| `GenerateWeight()` | 随机托运重量 | 7.0 ~ 17.0 kg |
| `generateVerifyCode()` | 随机验证码 | 6位字母数字组合 |
| `getRandomTime()` | 随机时间 | 未来1~30天内 |

---

## 5. 行为链组成

### 5.1 已实现的行为链

| 行为链 | 功能描述 | 节点序列 |
|--------|----------|----------|
| `NormalPreserveChain` | 正常预订票务 | 验证码→登录→查询用户→查询车次→选择路线→查询车次信息→查询座位→查询联系人→查询食物→查询保险→预订 |
| `NormalOrderPayChain` | 订单支付 | 验证码→登录→查询用户→刷新订单→支付 |
| `OrderConsignChain` | 订单托运 | 验证码→登录→查询用户→刷新订单→查询订单→托运 |
| `TicketCollectAndEnterStationChain` | 取票进站 | 验证码→登录→查询用户→刷新已支付订单→取票→进站 |
| `AdvancedSearchChain` | 高级搜索预订 | 验证码→登录→查询用户→选择路线→高级搜索→查询座位→查询联系人→查询食物→查询保险→预订 |
| `ConsignListChain` | 托运列表查询 | 验证码→登录→查询用户→查询托运列表 |
| `OrderChangeChain` | 订单改签 | 验证码→登录→查询用户→刷新其他订单→选择路线→查询车次→查询车次信息→查询座位→改签 |
| `OrderCancelChain` | 订单取消 | 验证码→登录→查询用户→刷新订单→计算退款→取消订单 |

### 5.2 节点复用示意图

```
                    VerifyCode → LoginBasic → QueryUser
                           ↓
        ┌──────────────────┼──────────────────┐
        ↓                  ↓                  ↓
   ChooseRoute        RefreshOrder      RefreshOrderOther
        ↓                  ↓                  ↓
   QueryTrain          OrderPay          TicketCollect
        ↓                                     ↓
   QueryTripInfo                         EnterStation
        ↓
   QuerySeatInfo
        ↓
   QueryContacts
        ↓
   QueryFood
        ↓
   QueryAssurance
        ↓
   Preserve
```

---

## 6. 上下文传递机制

系统使用 `Context` 结构在节点间传递数据，实现节点解耦：

```go
type Context struct {
    ctx context.Context
}

func (c *Context) Set(key string, value interface{}) { ... }
func (c *Context) Get(key string) interface{} { ... }
```

**主要上下文键值：**

| 类别 | 键名 | 用途 |
|------|------|------|
| 用户信息 | `AccountID`, `UserId`, `LoginToken` | 用户身份标识 |
| 路线信息 | `RouteID`, `StartStation`, `EndStation` | 行程路线 |
| 车次信息 | `TripID`, `TrainTypeName`, `DepartureTime` | 车次详情 |
| 订单信息 | `OrderId`, `Price`, `SeatClass` | 订单状态 |
| 联系人 | `ContactsID`, `Name`, `PhoneNumber` | 乘客信息 |
| 食物 | `FoodType`, `FoodName`, `StoreName` | 餐饮选择 |
| 保险 | `AssuranceTypeIndex`, `AssuranceTypeName` | 保险类型 |

---

## 7. 并发控制机制

### 7.1 多线程执行
```go
type LoadGenerator struct {
    config       *Config
    wg           sync.WaitGroup
    ctx          context.Context
    cancel       context.CancelFunc
    sharedClient *service.SvcImpl  // 共享客户端
}

func (l *LoadGenerator) Start() {
    l.wg.Add(l.config.Thread)
    for i := 0; i < l.config.Thread; i++ {
        go l.worker(i)  // 启动多个工作协程
    }
}
```

### 7.2 工作循环
```go
func (l *LoadGenerator) worker(index int) {
    defer l.wg.Done()
    for {
        select {
        case <-l.ctx.Done():
            return  // 优雅退出
        default:
            chainCtx := NewContext(context.Background())
            chainCtx.Set(Client, l.sharedClient)
            l.config.Chain.Execute(chainCtx)  // 执行行为链
            time.Sleep(time.Millisecond * time.Duration(l.config.SleepTime))
        }
    }
}
```

---

## 8. 设计优势总结

| 设计维度 | 实现方式 | 效果 |
|----------|----------|------|
| **行为随机性** | 加权概率选择行为链 | 模拟真实用户行为分布 |
| **数据随机性** | 节点内随机索引选择 | 覆盖不同数据组合 |
| **参数多样性** | 多种随机生成器 | 测试边界条件 |
| **策略多样性** | 多种搜索策略 | 覆盖不同业务路径 |
| **身份多样性** | 管理员/普通用户 | 模拟不同权限场景 |
| **可扩展性** | 链式设计 | 易于添加新行为 |
| **可复用性** | 节点复用 | 减少代码冗余 |

---

## 9. 随机性与多样性实现对照表

```
┌────────────────────────────────────────────────────────────────────┐
│                     随机性与多样性层次图                             │
├────────────────────────────────────────────────────────────────────┤
│  Level 1: 行为链选择 (Behavior Chain Selection)                     │
│  ┌──────────────────────────────────────────────────────────────┐  │
│  │ NormalPreserve(30%) | OrderPay(20%) | AdvancedSearch(20%)... │  │
│  └──────────────────────────────────────────────────────────────┘  │
├────────────────────────────────────────────────────────────────────┤
│  Level 2: 数据选择 (Data Selection)                                 │
│  ┌──────────────────────────────────────────────────────────────┐  │
│  │ 路线随机 | 车次随机 | 联系人随机 | 订单随机 | 食物随机          │  │
│  └──────────────────────────────────────────────────────────────┘  │
├────────────────────────────────────────────────────────────────────┤
│  Level 3: 参数生成 (Parameter Generation)                          │
│  ┌──────────────────────────────────────────────────────────────┐  │
│  │ 随机时间 | 随机重量 | 随机验证码 | 随机城市 | 随机食物名称       │  │
│  └──────────────────────────────────────────────────────────────┘  │
├────────────────────────────────────────────────────────────────────┤
│  Level 4: 策略多样化 (Strategy Diversification)                     │
│  ┌──────────────────────────────────────────────────────────────┐  │
│  │ 最便宜策略 | 最少经停策略 | 最快策略 | 座位等级分布             │  │
│  └──────────────────────────────────────────────────────────────┘  │
└────────────────────────────────────────────────────────────────────┘
```

通过以上多层次的随机性和多样性设计，该 Load Generator 能够生成高度模拟真实用户行为的测试负载，有效覆盖微服务系统的各个接口和业务场景。
