package main

import (
    //"fmt"
    //"strconv"
    "log"
    "time"

    "github.com/censhin/eve-poll/client"
)

var ec *client.Client = client.NewClient()

func main() {
    /*
    test, err := ec.CharNotifications("")
    if err != nil {
        fmt.Println("it's broke!")
    }

    notificationId := strconv.Itoa(test.Results[1].NotificationId)
    test2, err := ec.CharNotificationTexts("", notificationId)
    if err != nil {
        fmt.Println("it's broke!")
    }

    fmt.Println(test2)
    */

    charId := ""

    for {
        notfications, err := ec.CharNotifications(charId)
        if err != nil {
            log.Panic("it's broke, yo")
        }

        for _, notification := range(notifications.Results) {
            if notification.TypeId == 75 || notification.TypeId == 76 {
                notificationText, err := ec.CharNotificationTexts(charId, notification.NotificationId)
                if err != nil {
                    log.Panic("it's broke, yo")
                } else {
                    // publish notificationText to redis
                }
            }
        }

        time.Sleep(30)
    }
}

