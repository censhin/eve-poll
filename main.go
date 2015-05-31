package main

import (
    "fmt"
    "strconv"

    "github.com/censhin/eve-poll/client"
)

func main() {
    ec := client.NewClient()
    test, err := ec.CharNotifications("")
    if err != nil {
        fmt.Println("shit's broke!")
    }

    notificationId := strconv.Itoa(test.Results[1].NotificationId)
    test2, err := ec.CharNotificationTexts("", notificationId)
    if err != nil {
        fmt.Println("shit's broke!")
    }

    fmt.Println(test2)
}

