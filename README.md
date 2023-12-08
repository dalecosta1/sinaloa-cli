# SINALOA-CLI
The devops CLI used for automation with delivery and integration processes.

# Go instalaltion

- MacOS
brew install golang
export GOBIN=~/go/bin
export PATH=$PATH:$GOBIN


- Ubuntu


# Cobra cli usage

- Only to initialize the repo go mod init github.com/dalecosta1/sinaloa-cli
- go install github.com/spf13/cobra-cli@latest
- go get -u "github.com/ricochet2200/go-disk-usage/du"
- go get -u "github.com/dalecosta1/sinaloa-cli/cmd/info/sub/diskusage"
- cobra-cli init <cli_name> (to create cli)
- cobra-cli add <name_cmd> (to create a new command)
