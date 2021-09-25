# Quick Trade
Microservice to interact with quick-trade website.

## How to Install GO
    brew install go
    
    ## Make sure PATH env variable includes $GOPATH/bin, i.e. add the following to to ~/.bash_profile
       export PATH=$PATH:$GOPATH/bin
    
    ## And if Go binaries aren't executing properly, also be sure GOBIN is defined, the manual entry in .bash_profile would be
        export GOBIN=$GOPATH/bin


### How to Get copies of dependencies
Run this command in terminal to get copies of the dependencies:

    go mod vendor

### How to Compile proto files
Proto file are located [pkg/proto/](proto/). To generate go code run:

    make proto


### How to Run QuickTrade

    make run
    runs on port 8081 in local


### How to Compile QuickTrade

    make build

The output is:

    ./bin/linux_amd64/qt

### How to Create Postgres Database
Follow this [link](https://docs.mattermost.com/install/install-ubuntu-1604-postgresql.html?utm_source=google&utm_medium=cpc&obility_id=126977717891&utm_campaign=&utm_adgroup=&utm_term=&utm_content=516316717186&gclid=Cj0KCQjwvr6EBhDOARIsAPpqUPFZVh2j6wKilt7pBnCThPgKv_PnSvkHV6cIh-cXg6oAwPHbXBoJhIQaAswYEALw_wcB) to set up postgresql server.

       run following commands to install schema and get started:
       create database quick_trade;
       create user quick_trade with encrypted password 'quick_trade';
       grant all privileges on database quick_trade to quick_trade;
