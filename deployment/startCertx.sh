#!/usr/bin/env bash

# TODO correct ports
docker run -ti -v $HOME/sdh:/home/tendermint -v $PWD/../:/apps -p 1318:1317 -p 26658:26657 starport/cli:0.16.0 serve -p certX
