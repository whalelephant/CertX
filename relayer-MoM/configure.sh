#!/bin/bash -x

ibc-setup connect --home relayer-MoM

#Setup channel 
ibc-setup channel -v  \
          --src-connection connection-0 \
          --dest-connection connection-0 \
          --src-port employments \
          --dest-port eCredentials \
          --version eCredentials-1 \
          --home relayer-MoM

# error: Client ID 07-tendermint-0 for connection with ID connection-0 does not match counterparty client ID 07-tendermint-1 for connection with ID connection-0
# Setup connection
#Start
ibc-relayer start --home relayer
