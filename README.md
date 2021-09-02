# CLI
Example project of creating a basic CLI in Cobra and GO

## Install go

```
sudo apt-get install go-lang
```

## Set path
The below should be added to `~/.bashrc`
```
export GOPATH=$HOME/go
export GOBIN=$GOPATH/bin
export PATH=${PATH}:$GOBIN
```

## Get cobra
```
go get -u github.com/spf13/cobra/cobra
```
If you get an error regarding package `github.com/hashicorp/hcl/hcl/printer:`, try the following:
```
GO111MODULE=on go get -u github.com/spf13/cobra/cobra
```

## Install CLI
From the root folder of this repo, run the following command
```
go install todos
```
