package services

import (
	"fmt"
	"github.com/RodolfoBonis/bot_movie/utils"
	"time"
)

func ScheduleBotRoutine() {
	scheduledHour := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 21, 52, 0, 0, time.Local)
	go func(hour time.Time) {
		for {
			agora := time.Now()
			if agora.Weekday() == time.Wednesday && agora.Hour() == hour.Hour() && agora.Minute() == hour.Minute() {
				data := ReadFileData()

				message := GetMessage(data)

				SendMessage(message)

				reorderedData := utils.ReorderList(data)

				result := WriteFileData(nil, reorderedData)
				if result {
					fmt.Println("List reordered with success")
				}
			}

			time.Sleep(1 * time.Minute)
		}
	}(scheduledHour)
}
