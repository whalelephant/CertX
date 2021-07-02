# certX

This represents the publicly available blockchain storage for decentralised credentials.
There are two modules:

1. *credentials* - this provides publicly available plaintext credential for verification.
This is used to demonstrate use case 1
2. *eCredentials* - this provides publicly available encrypted credential (use case 2)

## credentials module

The `credential module` executes the following upon receiving an IBC message:

1. receives IBC messages from Issuer Zones containing verifiable credentials
2. verify the signature of the message as the actual issuer
3. records the credential

### credential types and packets

```sh
# Stores all credentials that were received via IBC
# note the original holder is not provided, only the new did
Credential: {
    # Verifier this proof is meant to be for, e.g. restaurant
    verifier: did
    # Issuer of the claim, e.g. Health Authority
    issuer: did,
    # Subject the claim is about, e.g. Alice's one time unique
    subject:did,
    # Actual claim, e.g. vaccination record received
    claim: Int 
}

# Packet received via IBC
verifiableCredential {
    # Verifier this proof is meant to be for, e.g. restaurant
    verifier: did,
    # Issuer of this proof, e.g. health authority
    issuer: did 
    # The identifier for the verifier, e.g. Alice's new did created for the restaurant
    subject: did,
    # Actual claim, e.g. vaccination record received or simply fully / partial / none
    # (i.e. the issuer can obfuscate details if they see fit)
    claim: some_claim_description,
    # Signature of the Issuer and related metadata
    # i.e. public key and signature for verification
    signature: signature_info 
}
```

## eCredential module

The `eCredential module` executes the following upon receiving an IBC message:

1. simply receives the credential and stores it.

### eCredential CLI

When querying for a specific `eCredential`, a user can provide the shared key (from the holder), and the CLI will display the credential in plain text.

If no key provided, then only encrypted text will be returned.

```sh
certXd query eCredentials show-eCredentialRecord <claim-id> <shared-key>
```

### eCredential types

```sh
eCredential / eCredentialRecord {
    # Subject did, e.g. Alice's did created when applying for a role
    subject: did,
    # Encrypted text of the credential, i.e. alice's employment history
    claim: string
}
```

## Useful Commands

```sh
# Query new plaintext credential from 
certXd query credentials list-credential --node tcp://localhost:36657

# Query specific encrypted credential with shared key
# If no key is provided, cipher text of credential is returned 
certXd query eCredentials show-eCredentialRecord <id> <shared key> --node tcp://localhost:36657
```
