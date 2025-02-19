package main

import (
    "github.com/mattermost/mattermost-server/v6/model"
    "github.com/mattermost/mattermost-server/v6/plugin"
)

func (p *Plugin) OnActivate() error {
    p.API.RegisterCommand(&model.Command{
        Trigger:          "remind",
        Description:      "Установите напоминание себе или Вашей команде.",
        DisplayName:      "Reminder",
        AutoComplete:     true,
        AutoCompleteDesc: "Example: /remind @user Do something in 1 hour",
        AutoCompleteHint: "[message] [time]",
    })
    return nil
}
