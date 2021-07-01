# mom

This directory represents the Ministry of Magic blockchain for use case 2.
The

![MoM](./mom.png)


## Useful cmd

```sh
# Queries the list of employment record 
MoMd query employments list-record --node tcp://localhost:46657

# send an encrypted credential to certX
# unlike muggleAuth, here Alice might want to provide a reference did
# e.g. her initial application to the Quidditch team was done with did:certX:aliceQuidditch
MoMd tx employments send-eCredential employments channel-0 did:certX:aliceQuidditch 0 --node tcp://localhost:46657 --home .home --from alice
```

To query the encrypted credentials, please refer to `/certX`
