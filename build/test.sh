#!/usr/bin/env bash
go test -coverprofile .coverage
go tool cover -func=.coverage