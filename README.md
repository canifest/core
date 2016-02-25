# core
[![Go Report Card](https://goreportcard.com/badge/github.com/canifest/core)](https://goreportcard.com/report/github.com/canifest/core)

## Description

Core module for canifest

## Build Instructions
To build it, start with go install github.com/canifest/core from your $GOPATH

Then you can start the core with ./bin/core

Then you can connect to it with Postman or through the CLI

## Front-End

### Description
This section will cover the front end portion of the canifest app (the majority of the static directory).

### app

#### components
This is where our directives will be.

#### core
We'll want to minimize the amount of logic inside controllers and directives. All business logic should reside here.

#### sections
This is where our controllers are.

### css
All styling elements are here. Currently, we're using bootstrap and a bootstrap theme (Pratt).

### img
All images reside here.

### lib
This is for third party libraries (bootstrap, jQuery, etc)
