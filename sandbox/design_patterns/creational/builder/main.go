package main

import "fmt"

func main() {
	// TODO: Create a NotificationBuilder and use it to set properties
	var bldr = newNotificationBuilder()
	// TODO: Use the builder to set some properties
	bldr.SetTitle("New one")
	bldr.SetIcon("icon.png")
	bldr.SetSubTitle("Subtitle example")
	bldr.SetImage("image.jpg")
	bldr.SetPriority(5)
	bldr.SetMessage("This is a basic test notification")
	bldr.SetType("alert")

	// TODO: Use the Build function to create a finished object
	news, err := bldr.Build()

	if err != nil {
		fmt.Println("Error ", err)
	} else {
		fmt.Printf("Notification: %+v", news)
	}
}
