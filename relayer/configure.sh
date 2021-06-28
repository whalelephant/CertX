#Setup connection

ibc-setup connect --home relayer

#Setup channel 
ibc-setup channel -v  \
          --src-connection connection-0 \
          --dest-connection connection-0 \
          --src-port credential \
          --dest-port vac \
          --version credential-1



#Start
ibc-relayer start --home relayer
