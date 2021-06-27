# muggleAuth

## MuggleAuth (Issuer Zone)

This represents the authorities for muggles (non-magical humans).
For example, it might be a blockchain that stores and provides verifiable credentials for registered citizens,
such as driver license, health record, and in this case, a vaccination record.

With potentially sensitive data, this blockchain is considered private.

### Vac Module

The `Vac Module` provides state that stores (1) the credential and (2) the proof for claims requested by muggles.

The first state is updated by direct messages to the module and is the lookup to facilitate the provisioning of the proof.

The second state is a record that allows for the Issuer to:

- revoke any proofs (perhaps Alice's vaccine is not effective with new variant)
- recall (i.e. if there is an outbreak at the restaurant, authorities may want to notify Alice)

### Vac Types

```sh
Credential: {
    # Issuer of the claim, e.g. Health Authority
    issuer: did,
    # Subject the claim is about, e.g. Alice's muggle did
    holder:did,
    # Actual claim, e.g. vaccination record recieved
    claim: Int 
}

ProofRecord {
    # message, in this case it is not encrypted
    message: {
        # The known identifier, e.g. Alice's muggle did
        holder: did,
        # The identifier for the verifier, e.g. Alice's new did created for the restaurant
        subject: did,
        # Verifier this proof is meant to be for, e.g. restaurant
        verifier: did,
        # Issuer of this proof, e.g. health authority
        issuer: did 
        # Actual claim, e.g.  vaccination record recieved or simply fully / partial / none
        # (In the case that types of vaccine have different number requirements, should not disclose)
        claim: some_claim_description,
    },
    # Signature of the message (in plain text) and related metadata
    signature: signature_info 
}

ProofRequest {
    message: {
        # The known identifier, e.g. Alice's muggle did
        holder: did,
        # The identifier for the verifier, e.g. Alice's new did created for the restaurant
        subject: did,
        # Verifier this proof is meant to be for, e.g. restaurant
        verifier: did,
        # Reference to the claim 
        claim: claim_id,
    },
    # Signature of the message signed by the holder and related metadata
    signature: signature_info 
}

```


## Get started

```sh
starport serve
```

`serve` command installs dependencies, builds, initializes and starts your blockchain in development.

## Configure

Your blockchain in development can be configured with `config.yml`. To learn more see the [reference](https://github.com/tendermint/starport#documentation).

## Launch

To launch your blockchain live on mutliple nodes use `starport network` commands. Learn more about [Starport Network](https://github.com/tendermint/spn).

## Useful cmd

```sh
# list all keys
muggleAuthd keys list --home .home

# create record from healthauth for alice
# the first arg is not actually used since it is signed by healthauth
muggleAuthd tx vac create-credential "_healthauth" "did:muggleAuth:pub1addwnpepq0cxxuncdddefhwyqj02zy8xlhyk775ud4t7852tzw6kl8v2t6wqjrwtl75" "2" --from  healthauth --home .home

# sends the verifiable credential
muggleAuthd tx vac send-verifiableCredential <port> <channel> did:certX:restaurant81 0 --home .home

```
