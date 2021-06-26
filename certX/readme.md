# certX

## CertX Module

The `CertX Module` simple receives IBC messages from Issuer Zones and records the proofs.

## CertX Types

```sh
VerifiableCredential {
    # The identifier for the verifier, e.g. Alice's new did created for the retaurant
    subject: did,
    # Verifier this proof is meant to be for, e.g. restaurant
    verifier: did,
    # Issuer of this proof, e.g. health authority
    issuer: did 
    # Actual claim, e.g. vaccination record recieved or simply fully / partial / none
    # (In the case that types of vaccine have different number requirements, should not be disclose)
    claim: some_claim_description,
    # Signature of the Issuer and related metadata
    signature: signature_info 
}
```

**certX** is a blockchain built using Cosmos SDK and Tendermint and created with [Starport](https://github.com/tendermint/starport).

## Get started

```sh
starport serve
```

`serve` command installs dependencies, builds, initializes and starts your blockchain in development.

## Configure

Your blockchain in development can be configured with `config.yml`. To learn more see the [reference](https://github.com/tendermint/starport#documentation).

## Launch

To launch your blockchain live on mutliple nodes use `starport network` commands. Learn more about [Starport Network](https://github.com/tendermint/spn).
