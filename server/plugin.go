package main

import (
	root "github.com/cpoile/mattermost-plugin-later"
	"github.com/cpoile/mattermost-plugin-later/server/bot"
	"github.com/cpoile/mattermost-plugin-later/server/command"
	"github.com/cpoile/mattermost-plugin-later/server/config"
	pluginapi "github.com/mattermost/mattermost-plugin-api"
	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/mattermost/mattermost-server/v6/plugin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"net/http"
)

// Plugin implements the interface expected by the Mattermost server to communicate between the server and plugin processes.
type Plugin struct {
	plugin.MattermostPlugin

	config    *config.ServiceImpl
	bot       *bot.Bot
	pluginAPI *pluginapi.Client
}

// ServeHTTP demonstrates a plugin that handles HTTP requests by greeting the world.
func (p *Plugin) ServeHTTP(c *plugin.Context, w http.ResponseWriter, r *http.Request) {
}

// OnActivate is called when the plugin is activated.
func (p *Plugin) OnActivate() error {
	p.pluginAPI = pluginapi.NewClient(p.API, p.Driver)

	p.config = config.NewConfigService(p.pluginAPI, &root.Manifest)
	pluginapi.ConfigureLogrus(logrus.New(), p.pluginAPI)

	botID, err := p.pluginAPI.Bot.EnsureBot(&model.Bot{
		Username:    "latergator",
		DisplayName: "Later Gator Bot",
		Description: "Later Gator lets you schedule messages to be sent in the future",
	},
		pluginapi.ProfileImagePath("assets/plugin_icon.png"),
	)
	if err != nil {
		return errors.Wrap(err, "failed to ensure bot")
	}

	err = p.config.UpdateConfiguration(func(c *config.Configuration) {
		c.BotUserID = botID
		c.AdminLogLevel = "debug"
	})
	if err != nil {
		return errors.Wrap(err, "failed to save bot to the config")
	}

	p.bot = bot.New(p.pluginAPI, p.config.GetConfiguration().BotUserID, p.config)

	if err = command.RegisterCommands(p.API.RegisterCommand); err != nil {
		return errors.Wrapf(err, "failed register commands")
	}

	return nil
}

// OnConfigurationChange handles any change in the configuration.
func (p *Plugin) OnConfigurationChange() error {
	if p.config == nil {
		return nil
	}

	return p.config.OnConfigurationChange()
}

// ExecuteCommand executes a command that has been previously registered via the RegisterCommand.
func (p *Plugin) ExecuteCommand(c *plugin.Context, args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	runner := command.NewCommandRunner(c, args, pluginapi.NewClient(p.API, p.Driver), p.bot, p.bot, p.config)

	if err := runner.Execute(); err != nil {
		return nil, model.NewAppError("Later.ExecuteCommand", "Unable to execute command.", nil, err.Error(), http.StatusInternalServerError)
	}

	return &model.CommandResponse{}, nil
}
