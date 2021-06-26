# muggleAuth

## MuggleAuth (Issuer Zone)

This represents the authorities for muggles (non-magical humans).
For example, it might be a blockchain that stores and provide verifiable credential for registered citizens,
such as driver license, health record, and in this case, a vaccination record.

With potentially sensitive data, this blockchain is considered private.

Please refer to `./muggleAuth/readme.md` for details

### Vac Module

The `Vac Module` provides states that stores (1) the credentials and (2) the proof for requested claims from muggles.

The first state is updated by direct messages to the module and is the lookup to facilitate the provision of proof.

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
    # Actual claim, e.g. number of vaccination recieved
    claim: Int 
}

ProofRequest {
    message: {
        # The known identifier, e.g. Alice's muggle did
        holder: did,
        # The identifier for the verifier, e.g. Alice's new did created for the retaurant
        subject: did,
        # Verifier this proof is meant to be for, e.g. restaurant
        verifier: did,
        # Reference to the claim 
        claim: claim_id,
    },
    # Signature of the message signed by the holder and related metadata
    signature: signature_info 
}

ProofRecord {
    # message, in this case it is not encrypted
    message: {
        # The known identifier, e.g. Alice's muggle did
        holder: did,
        # The identifier for the verifier, e.g. Alice's new did created for the retaurant
        subject: did,
        # Verifier this proof is meant to be for, e.g. restaurant
        verifier: did,
        # Issuer of this proof, e.g. health authority
        issuer: did 
        # Actual claim, e.g. number of vaccination recieved or simply fully / partial / none
        # (In the case that types of vaccine have different number requirements, should not disclose)
        claim: some_claim_description,
    },
    # Signature of the message (in plain text) and related metadata
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
