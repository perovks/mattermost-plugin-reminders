package main

import (
    "github.com/mattermost/mattermost-server/v6/plugin"
)

type Plugin struct {
    plugin.MattermostPlugin
}

func main() {
    plugin.ClientMain(&Plugin{})
}
