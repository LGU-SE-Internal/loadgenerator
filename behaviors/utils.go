package behaviors

import (
	"fmt"
	"github.com/Lincyaw/loadgenerator/service"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func randomlyChoosePlaces(data []service.Place) (string, string, error) {
	if len(data) < 2 {
		return "", "", fmt.Errorf("not enough places to choose from")
	}

	var names []string
	for _, place := range data {
		names = append(names, place.Name)
	}

	rand.Seed(time.Now().UnixNano())
	i, j := rand.Intn(len(names)), rand.Intn(len(names))
	for i == j {
		j = rand.Intn(len(names))
	}

	return names[i], names[j], nil
}

func getMiddleElements(input string) string {
	elements := strings.Split(input, ",")

	// If the input contains less than 3 elements, return an empty string
	if len(elements) < 3 {
		return ""
	}

	middleElements := elements[1 : len(elements)-1]
	return strings.Join(middleElements, ",")
}

// GenerateWeight generates a float64 value between 0 and 15.
func GenerateWeight() float64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Float64() * 15
}

func generateDescription() string {
	rand.Seed(time.Now().UnixNano())

	// Generate a random number with one decimal place between 0.1 and 10.0
	randomNumber := rand.Float64()*9.9 + 0.1
	numberStr := strconv.FormatFloat(randomNumber, 'f', 1, 64)

	// Determine if 'Max' should be replaced by 'Min' with a probability of 0.3
	replaceMax := rand.Float64() < 0.3
	description := "Max"
	if replaceMax {
		description = "Min"
	}

	return fmt.Sprintf("%s in %s hour", description, numberStr)
}

// IsWithin checks if a given float64 value is 7.0 or less.
func BooleanIsWithin(value float64) bool {
	return value <= 7.0
}

// generateVerifyCode generates a 6-digit verification code consisting of letters and numbers.
func generateVerifyCode() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const length = 6
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

	code := make([]byte, length)
	for i := range code {
		code[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(code)
}

func generateTrainTypeName(input string) string {
	startLetter := strings.ToUpper(string(input[0]))

	var MockedTrainType string

	switch startLetter {
	case "G":
		if rand.Intn(2) == 0 {
			MockedTrainType = "GaoTieOne"
		} else {
			MockedTrainType = "GaoTieTwo"
		}
	case "Z":
		MockedTrainType = "ZhiDa"
	case "T":
		MockedTrainType = "TeKuai"
	case "K":
		MockedTrainType = "KuaiSu"
	case "D":
		MockedTrainType = "DongCheOne"
	default:
		MockedTrainType = "Unknown"
	}

	return MockedTrainType
}

// generateDocumentNumber generates a DocumentNumber with 50% probability for "DocumentNumber_One"
// and 50% probability for "DocumentNumber_Two".
func generateDocumentNumber() string {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator

	if rand.Intn(2) == 0 {
		return "DocumentNumber_One"
	} else {
		return "DocumentNumber_Two"
	}
}

func generateContactsName() string {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator

	if rand.Intn(2) == 0 {
		return "Contacts_One"
	} else {
		return "Contacts_Two"
	}
}

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

func generateCoachNumber() int {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator

	// Generate a random number between 1 and 10 (inclusive)
	return rand.Intn(10) + 1
}

// toLowerCaseAndRemoveSpaces converts a given string to all lower case
// and removes all spaces.
func toLowerCaseAndRemoveSpaces(input string) string {
	lowercased := strings.ToLower(input)
	noSpaces := strings.ReplaceAll(lowercased, " ", "")
	return noSpaces
}

// generateRandomFood generates a random food item from a predefined list of 50 kinds of food.
func generateRandomFood() string {
	// Predefined list of food items
	foodList := []string{
		"Pizza", "Burger", "Pasta", "Sushi", "Tacos", "Salad", "Steak", "Soup", "Sandwich", "Fries",
		"Ice Cream", "Cake", "Donut", "Chocolate", "Apple", "Banana", "Orange", "Grapes", "Strawberry", "Blueberry",
		"Mango", "Pineapple", "Watermelon", "Kiwi", "Avocado", "Tomato", "Cucumber", "Carrot", "Broccoli", "Spinach",
		"Chicken", "Beef", "Pork", "Lamb", "Fish", "Shrimp", "Crab", "Lobster", "Eggs", "Cheese",
		"Yogurt", "Milk", "Butter", "Bread", "Rice", "Pasta", "Noodles", "Cereal", "Oatmeal", "Honey",
	}

	// Seed the random number generator to ensure different results each run
	rand.Seed(time.Now().UnixNano())

	// Generate a random index to pick a food item
	randomIndex := rand.Intn(len(foodList))

	// Return the randomly selected food item
	return foodList[randomIndex]
}

// generateRandomStoreName generates a random store name from a predefined list of 30 kinds of store names.
func generateRandomStoreName() string {
	// Predefined list of store names
	storeNames := []string{
		"Grocery Mart", "Tech World", "Fashion Hub", "Book Haven", "Toy Land",
		"Pet Paradise", "Home Essentials", "Beauty Bliss", "Sports Store", "Gadget Garage",
		"Furniture Factory", "Shoe Stop", "Pharmacy Plus", "Hardware Haven", "Electronics Emporium",
		"Music Mania", "Garden Goods", "Office Outlet", "Auto Accessories", "Craft Corner",
		"Gift Gallery", "Jewelry Junction", "Bakery Bliss", "Café Corner", "Fitness Freak",
		"Outdoor Outfitters", "Travel Treasures", "Kids' Kingdom", "Vintage Vault", "Wine World",
	}

	// Seed the random number generator to ensure different results each run
	rand.Seed(time.Now().UnixNano())

	// Generate a random index to pick a store name
	randomIndex := rand.Intn(len(storeNames))

	// Return the randomly selected store name
	return storeNames[randomIndex]
}

func extractDate(dateTimeStr string) string {
	// Parse the string to time.Time
	t, err := time.Parse("2006-01-02 15:04:05", dateTimeStr)
	if err != nil {
		return ""
	}
	// Format the time.Time to only include the date
	return t.Format("2006-01-02")
}

// generateRandomTime generates a random time in the format "HH:MM:SS".
func generateRandomTime() string {
	hour := rand.Intn(24)   // 0-23
	minute := rand.Intn(60) // 0-59
	second := rand.Intn(60) // 0-59
	return fmt.Sprintf("%02d:%02d:%02d", hour, minute, second)
}

type TimeConfig struct {
	StartTime string
}

type Option func(config *TimeConfig)

func WithStartTime(startTime string) Option {
	return func(config *TimeConfig) {
		config.StartTime = startTime
	}
}
func getRandomTime(opts ...Option) string {
	config := &TimeConfig{}
	for _, opt := range opts {
		opt(config)
	}

	now := time.Now()

	if config.StartTime != "" {
		startTime, err := time.Parse("2006-01-02 15:04:05", config.StartTime)
		if err != nil {
			fmt.Println("Invalid StartTime format, using current time instead.")
		} else {
			now = startTime
			// 生成1小时到1天之后的时间
			randomHours := rand.Intn(24) + 1
			randomDate := now.Add(time.Duration(randomHours) * time.Hour)
			return randomDate.Format("2006-01-02 15:04:05")
		}
	}

	// 保持原来的逻辑，生成从今天起到未来一个月内的随机日期
	randomDays := rand.Intn(30) + 1
	randomDate := now.AddDate(0, 0, randomDays)
	return randomDate.Format("2006-01-02 15:04:05")
}

// ConvertCommaSeparatedToBracketed converts a comma-separated string to a bracketed, space-separated string
func ConvertCommaSeparatedToBracketed(input string) string {
	// 删除字符串前后的空白
	input = strings.TrimSpace(input)
	// 按逗号分隔字符串，并去除每个元素前后的空白
	parts := strings.Split(input, ",")
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}
	// 将分隔后的元素用空格连接，并用方括号包围
	result := fmt.Sprintf("[%s]", strings.Join(parts, " "))
	return result
}

// helper function for Order Service
// RandomDecimalStringBetween 生成并返回两个整数之间的一位小数形式的随机数字符串，包括边界值。
func RandomDecimalStringBetween(min, max int) string {
	rand.Seed(time.Now().UnixNano())
	randomInt := rand.Intn(max-min+1) + min              // 生成[min, max]范围内的随机整数
	decimalValue := float64(randomInt) * 0.1             // 将整数转换为一位小数
	return strconv.FormatFloat(decimalValue, 'f', 1, 64) // 转换为一位小数的字符串形式
}

// RandomProvincialCapitalEN 随机返回一个中国省会城市的英文名称
func RandomProvincialCapitalEN() string {
	rand.Seed(time.Now().UnixNano())
	return provincialCapitalsEN[rand.Intn(len(provincialCapitalsEN))]
}

// 中国省会城市的英文列表
var provincialCapitalsEN = []string{
	"Beijing", "Shanghai", "Tianjin", "Chongqing",
	"Shijiazhuang", "Taiyuan", "Hohhot", "Shenyang", "Changchun", "Harbin",
	"Nanjing", "Hangzhou", "Hefei", "Fuzhou", "Nanchang", "Jinan", "Zhengzhou", "Wuhan", "Changsha", "Guangzhou",
	"Nanning", "Haikou", "Chengdu", "Guiyang", "Kunming", "Lhasa", "Xi'an", "Lanzhou", "Xining", "Yinchuan",
	"Urumqi", "Taipei",
}

// GetTrainTicketClass 随机返回高铁票等级。
// 有5%的概率返回"FirstClass"（头等座），
// 15%的概率返回"BusinessClass"（一等座），
// 剩余80%的概率返回"EconomyClass"（二等座）。
func GetTrainTicketClass() int {
	rand.Seed(time.Now().UnixNano()) // 确保每次运行时随机数种子不同

	probability := rand.Intn(100) // 生成0到99之间的随机数

	switch {
	case probability < 5:
		return 0
	case probability < 20:
		return 1
	default:
		return 2
	}
}

func GenerateTrainTypeName() string {
	// 设置随机数种子
	rand.Seed(time.Now().UnixNano())

	// 定义可能的火车类型名称
	trainTypes := []string{"GaoTieOne", "GaoTieTwo", "DongCheOne", "ZhiDa", "TeKuai", "KuaiSu"}

	// 随机选择一个火车类型名称
	MockedTrainTypeName := trainTypes[rand.Intn(len(trainTypes))]

	return MockedTrainTypeName
}

// generateRandomCityName generates a random city name from a predefined list of city names.
func generateRandomCityName() string {
	// Predefined list of city names
	cityNames := []string{
		"nanjing", "shijiazhuang", "wuxi", "shanghaihongqiao", "jiaxingnan",
		"hangzhou", "shanghai", "zhenjiang", "suzhou", "taiyuan",
		"xuzhou", "jinan", "beijing",
	}

	// Seed the random number generator to ensure different results each run
	rand.Seed(time.Now().UnixNano())

	// Generate a random index to pick a city name
	randomIndex := rand.Intn(len(cityNames))

	// Return the randomly selected city name
	return cityNames[randomIndex]
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

func IntListToString(numbers []int) string {
	// 使用 strings.Builder 进行高效的字符串拼接
	var builder strings.Builder

	for i, number := range numbers {
		if i > 0 {
			builder.WriteString(", ")
		}
		builder.WriteString(fmt.Sprintf("Numbers[%d] %d", i, number))
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

// IntSliceToString converts a slice of integers to a bracketed, space-separated string
func IntSliceToString(ints []int) string {
	// 使用 strings.Builder 进行高效的字符串拼接
	var builder strings.Builder
	builder.WriteString("[")
	for i, val := range ints {
		if i > 0 {
			builder.WriteString(" ")
		}
		builder.WriteString(fmt.Sprintf("%d", val))
	}
	builder.WriteString("]")
	return builder.String()
}

// StringSliceToString converts a slice of strings to a bracketed, space-separated string
func StringSliceToString(strs []string) string {
	// 使用 strings.Builder 进行高效的字符串拼接
	var builder strings.Builder
	builder.WriteString("[")
	for i, val := range strs {
		if i > 0 {
			builder.WriteString(" ")
		}
		builder.WriteString(val)
	}
	builder.WriteString("]")
	return builder.String()
}

// RandomSelectString selects a random string from a given slice of strings
func RandomSelectString(options []string) string {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator
	randomIndex := rand.Intn(len(options))
	return options[randomIndex]
}
