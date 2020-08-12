package web_notification

import (
	"encoding/json"
	"fmt"
	"github.com/SherClockHolmes/webpush-go"
	"github.com/lukasvdberk/opensource-discord/database"
	"os"
	"strconv"
)

type NotificationDevice struct {
	id     int64
	userId int64
	// Below are field that are required by the web-push protocol to send notifications over the web
	endpoint string
	authKey  string
	p256dh   string
}

type Notification struct {
	// add a notification popup data stream or something? with title message and icon.
	Data interface{} `json:"data"`

	// Channel should something like "direct-messages" so the friend end know what kinds of notification this is.
	MessageChannel string `json:"messageChannel"`
}

func PushNotificationToUser(userId int64, notification Notification) bool {
	// send notification

	fmt.Println(userId)
	webNotificationDeviceListMap := database.SelectStatement(
		"SELECT * FROM NotificationDevice WHERE userId = ? LIMIT 1",
		userId,
	)

	fmt.Println("parsed notification data", webNotificationDeviceListMap)
	if len(webNotificationDeviceListMap) > 0 {
		webNotificationDeviceMap := webNotificationDeviceListMap[0]
		notificationDevice := new(NotificationDevice)

		idStr := webNotificationDeviceMap["id"]

		if _, err := strconv.Atoi(idStr); err == nil {
			notificationDevice.id, _ = strconv.ParseInt(idStr, 10, 64)
		}

		userIdStr := webNotificationDeviceMap["userId"]

		if _, err := strconv.Atoi(userIdStr); err == nil {
			notificationDevice.userId, _ = strconv.ParseInt(userIdStr, 10, 64)
		}

		notificationDevice.endpoint = webNotificationDeviceMap["endpoint"]
		notificationDevice.authKey = webNotificationDeviceMap["authKey"]
		notificationDevice.p256dh = webNotificationDeviceMap["p256dh"]
		publicKey := os.Getenv("WEB_PUSH_PUBLIC_KEY")
		privateKey := os.Getenv("WEB_PUSH_PRIVATE_KEY")

		fmt.Println("parsed notification data", notificationDevice)
		s := &webpush.Subscription{}
		s.Endpoint = notificationDevice.endpoint
		s.Keys.P256dh = notificationDevice.p256dh
		s.Keys.Auth = notificationDevice.authKey
		// Send Notification
		jsonData, err := json.Marshal(notification)

		if err != nil {
			return false
		}

		resp, err := webpush.SendNotification(jsonData, s, &webpush.Options{
			Subscriber:      os.Getenv("WEB_PUSH_EMAIL"),
			VAPIDPublicKey:  publicKey,
			VAPIDPrivateKey: privateKey,
			TTL:             200,
		})

		if err != nil {
			return false
		}

		defer resp.Body.Close()
		fmt.Println("success sending push message")
		return true
	} else {
		return false
	}
}

func SaveNotificationDevice(webDevice NotificationDevice) bool {
	// All field from webDevice should be filled.
	deviceId := database.InsertStatement("INSERT INTO NotificationDevice VALUES(DEFAULT, ?, ?, ?, ?);",
		webDevice.userId,
		webDevice.endpoint,
		webDevice.authKey,
		webDevice.p256dh,
	)

	// If the record was successfully saved
	return deviceId != -1
}
