# Include custom targets and environment variables here

## Generate mocks.
mocks:
ifneq ($(HAS_SERVER),)
	go get github.com/golang/mock/mockgen
	mockgen -destination server/config/mocks/mock_service.go github.com/cpoile/mattermost-plugin-later/server/config Service
	mockgen -destination server/bot/mocks/mock_logger.go github.com/cpoile/mattermost-plugin-later/server/bot Logger
	mockgen -destination server/bot/mocks/mock_poster.go github.com/cpoile/mattermost-plugin-later/server/bot Poster
endif
