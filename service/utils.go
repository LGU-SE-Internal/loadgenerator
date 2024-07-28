package service

import (
	"fmt"
	"github.com/go-faker/faker/v4"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

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

// toLowerCaseAndRemoveSpaces converts a given string to all lower case
// and removes all spaces.
func toLowerCaseAndRemoveSpaces(input string) string {
	lowercased := strings.ToLower(input)
	noSpaces := strings.ReplaceAll(lowercased, " ", "")
	return noSpaces
}

func GenerateTrainTypeName() string {
	// 设置随机数种子
	rand.Seed(time.Now().UnixNano())

	// 定义可能的火车类型名称
	trainTypes := []string{"GaoTieOne", "GaoTieTwo", "GaoTieSeven", "DongCheOne", "DongCheTen"}

	// 随机选择一个火车类型名称
	MockedTrainTypeName := trainTypes[rand.Intn(len(trainTypes))]

	return MockedTrainTypeName
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

func generateRandomNumberString() string {
	rand.Seed(time.Now().UnixNano())
	numberLength := 10 // Length of the number string

	// Generate a random number string of the specified length
	numberStr := ""
	for i := 0; i < numberLength; i++ {
		digit := rand.Intn(10) // Generate a random digit (0-9)
		numberStr += strconv.Itoa(digit)
	}

	return numberStr
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

// generateRandomTime generates a random time in the format "HH:MM:SS".
func generateRandomTime() string {
	//// Seed the random number generator to ensure different results each run
	//rand.Seed(time.Now().UnixNano())

	// Generate random hours, minutes, and seconds
	hour := rand.Intn(24)   // 0-23
	minute := rand.Intn(60) // 0-59
	second := rand.Intn(60) // 0-59

	// Format the time as "HH:MM:SS"
	return fmt.Sprintf("%02d:%02d:%02d", hour, minute, second)
}

func getRandomTime() string {
	MockedRandomDate := faker.Date()

	////// Seed the random number generator to ensure different results each run
	////rand.Seed(time.Now().UnixNano())
	//// Generate random hours, minutes, and seconds
	//hour := rand.Intn(24)   // 0-23
	//minute := rand.Intn(60) // 0-59
	//second := rand.Intn(60) // 0-59
	//MockedRandomTime := fmt.Sprintf("%02d:%02d:%02d", hour, minute, second)
	MockedRandomTime := generateRandomTime()

	DateAndTime := MockedRandomDate + " " + MockedRandomTime

	return DateAndTime
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
