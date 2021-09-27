package factory

import (
	"fmt"
	"gomeetme/timeconverters"
	"math/rand"
	"time"
)

func GenerateRandomTime() string {
	return fmt.Sprintf("%s:%s:%s", generateHours(), generateMinutes(), generateSeconds())
}

func generateHours() string {
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(24)
	hourList := make([]string, 0)
	for i := 0; i < 25; i++ {
		hourList = append(hourList, timeconverters.AddLeading0(i))
	}
	return hourList[randomIndex]
}
func generateMinutes() string {
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(60)
	minuteList := make([]string, 0)
	for i := 0; i < 60; i++ {
		minuteList = append(minuteList, timeconverters.AddLeading0(i))
	}
	return minuteList[randomIndex]
}
func generateSeconds() string {
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(60)
	secondList := make([]string, 0)
	for i := 0; i < 60; i++ {
		secondList = append(secondList, timeconverters.AddLeading0(i))
	}
	return secondList[randomIndex]
}
