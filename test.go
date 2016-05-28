package main

import (
    "github.com/deckarep/gosx-notifier"
    "log"
)

func main() {
    note := gosxnotifier.NewNotification("Body")

    //Optionally, set a title
    note.Title = "Title"

    //Optionally, set a subtitle
    //note.Subtitle = "Subtitle"

    //Optionally, set a sound from a predefined set.
    note.Sound = gosxnotifier.Default

    //Optionally, set a group which ensures only one notification is ever shown replacing previous notification of same group id.
    note.Group = "Go_app"

    //Optionally, set a sender (Notification will now use the Safari icon)
    //note.Sender = "com.apple.Safari"

    //Optionally, specifiy a url or bundleid to open should the notification be
    //clicked.
    note.Link = "https://status.github.com/"

    //Optionally, an app icon (10.9+ ONLY)
    note.AppIcon = "./Logo.png"

    //Optionally, a content image (10.9+ ONLY)
    //note.ContentImage = "~/Logos/GitHub-Mark-64px.png"
		//note.ContentImage = "../../../Logos/GitHub-Mark-64px.png"

    //Then, push the notification
    err := note.Push()

    //If necessary, check error
    if err != nil {
        log.Println("Uh oh!")
    }
}
