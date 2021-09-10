// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package command

import (
	"errors"
	"fmt"
	"github.com/cpoile/mattermost-plugin-later/server/bot"
	"github.com/cpoile/mattermost-plugin-later/server/config"
	pluginapi "github.com/mattermost/mattermost-plugin-api"
	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/mattermost/mattermost-server/v6/plugin"
	"strings"
)

const helpText = "###### Later Gator Plugin - Slash Command Help\n" +
	"* `/later test` - test. \n" +
	"\n" +
	""

const commands = "test, more more"

// Register is a function that allows the runner to register commands with the mattermost server.
type Register func(*model.Command) error

// RegisterCommands should be called by the plugin to register all necessary commands
func RegisterCommands(registerFunc Register) error {
	return registerFunc(getCommand())
}

func getCommand() *model.Command {
	return &model.Command{
		Trigger:          "later",
		DisplayName:      "Later Gator",
		Description:      "Later Gator",
		AutoComplete:     true,
		AutoCompleteDesc: "Available commands: " + commands,
		AutoCompleteHint: "[command]",
		AutocompleteData: getAutocompleteData(),
	}
}

func getAutocompleteData() *model.AutocompleteData {
	command := model.NewAutocompleteData("later", "[command]",
		"Available commands: "+commands)

	test := model.NewAutocompleteData("test", "", "Test command")
	command.AddCommand(test)

	return command
}

// Runner handles commands.
type Runner struct {
	context       *plugin.Context
	args          *model.CommandArgs
	pluginAPI     *pluginapi.Client
	logger        bot.Logger
	poster        bot.Poster
	configService config.Service
}

// NewCommandRunner creates a command runner.
func NewCommandRunner(ctx *plugin.Context, args *model.CommandArgs, api *pluginapi.Client,
	logger bot.Logger, poster bot.Poster, configService config.Service) *Runner {
	return &Runner{
		context:       ctx,
		args:          args,
		pluginAPI:     api,
		logger:        logger,
		poster:        poster,
		configService: configService,
	}
}

func (r *Runner) isValid() error {
	if r.context == nil || r.args == nil || r.pluginAPI == nil {
		return errors.New("invalid arguments to command.Runner")
	}
	return nil
}

func (r *Runner) postCommandResponse(text string) {
	post := &model.Post{
		Message: text,
	}
	r.poster.EphemeralPost(r.args.UserId, r.args.ChannelId, post)
}

func (r *Runner) warnUserAndLogErrorf(format string, args ...interface{}) {
	r.logger.Errorf(format, args...)
	r.poster.EphemeralPost(r.args.UserId, r.args.ChannelId, &model.Post{
		Message: "Your request could not be completed. Check the system logs for more information.",
	})
}

func (r *Runner) actionTest(args []string) {
	r.postCommandResponse(fmt.Sprintf("It works! Args: %#+v", args))
}

// Execute should be called by the plugin when a command invocation is received from the Mattermost server.
func (r *Runner) Execute() error {
	if err := r.isValid(); err != nil {
		return err
	}

	split := strings.Fields(r.args.Command)
	command := split[0]
	parameters := []string{}
	cmd := ""
	if len(split) > 1 {
		cmd = split[1]
	}
	if len(split) > 2 {
		parameters = split[2:]
	}

	if command != "/later" {
		return nil
	}

	switch cmd {
	case "test":
		r.actionTest(parameters)
	default:
		r.postCommandResponse(helpText)
	}

	return nil
}
