#!/usr/bin/env bash
docker run -ti -v $HOME/sdh:/home/tendermint -v $PWD/../:/apps -p 1317:1317 -p 26657:26657 starport/cli:0.16.0 serve -p muggleAuth
