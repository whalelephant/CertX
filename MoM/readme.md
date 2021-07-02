# mom

## MoM (Issuer Zone)

This directory represents the Ministry of Magic blockchain for *Use Case 2*.


### High level interactions with certX, verifier and users
With potentially sensitive data, such as employment records, this blockchain is private.


![MoM](./mom.png)

## Employments Module

This module records employment records, however, unlike [muggleAuth](../muggleauth/readme.md), the credential provided to *certX* is encrypted and only the credential holder, i.e., Alice will have the key. In this case, she can share with her potential employers to review her credentials.


#### Simple Step Breakdown
1. Checks if the record exists
1. If so, encrypts the record
1. Sends IBC packet `eCredential` to *certX*

## Useful cmd

```sh
# Queries the list of employment record 
MoMd query employments list-record --node tcp://localhost:46657

# send an encrypted credential to certX
# unlike muggleAuth, here Alice might want to provide a reference did
# e.g. her initial application to the Quidditch team was done with did:certX:aliceQuidditch
MoMd tx employments send-eCredential employments channel-0 did:certX:aliceQuidditch 0 --node tcp://localhost:46657 --home .home --from alice
```

To query the encrypted credentials, please refer to [certX](../certX/readme.md)
