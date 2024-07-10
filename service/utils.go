package service

import (
	"fmt"
	"github.com/go-faker/faker/v4"
	"math/rand"
	"strings"
	"time"
)

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

func GenerateTrainTypeName() string {
	// 设置随机数种子
	rand.Seed(time.Now().UnixNano())

	// 定义可能的火车类型名称
	trainTypes := []string{"GaoTieOne", "GaoTieTwo", "GaoTieSeven", "DongCheOne", "DongCheTen"}

	// 随机选择一个火车类型名称
	MockedTrainTypeName := trainTypes[rand.Intn(len(trainTypes))]

	return MockedTrainTypeName
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

func getRandomTime() string {
	randomDate := faker.Date()
	randomTime := faker.TIME

	DateAndTime := randomDate + " " + randomTime

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
