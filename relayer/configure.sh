#!/bin/bash -x

ibc-setup connect --home relayer

#Setup channel 
ibc-setup channel -v  \
          --src-connection connection-1 \
          --dest-connection connection-1 \
          --src-port vac \
          --dest-port credentials \
          --version credentials-1 \
          --home relayer

# error: Client ID 07-tendermint-0 for connection with ID connection-0 does not match counterparty client ID 07-tendermint-1 for connection with ID connection-0
# Setup connection
#Start
ibc-relayer start --home relayer
