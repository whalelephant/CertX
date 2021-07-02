# CertX

This project is designed to demonstrate the use of IBC between different zones in the Cosmos ecosystem for privacy preserving credential management.
**This project is incubated for the submission of HackATOM HCMC 2021**

## Motivation

As we all become increasingly aware of the risk of having our privacy taken away whilst demanding interoperability between systems,
a single entity or even blockchain providing the single source of truth is no longer satisfactory.

This is where IBC becomes a critical component on which the future is built on.

Since the relayer service can be external, it makes sense to design for cases where the packets sent can be plain text or encrypted.

The use cases we designed for are communications between:

1. Private to Public zones where information is public: [Use Case 1]
1. Private to Public zones where information is authorised by the sender only [Use Case 2]
1. Private to Private zones (key-exchanged) [Use Case 3]

With all the cases listed, this project uses the concepts of verifiable credential to showcase IBC usage.
We use [*dencentralised identifiers*] for entities to ensure no social mapping, behaviour patterns can be derived.

### Design Overview

It is important to conform to standards and practices that already works well and only make adjustments where needed. No point re-inventing the wheel!

We follow W3C recommendations of Verifiable Credential. A summary of the [roles] defined are as such:

![w3c_defined_vc](./w3c_defined_vc.png)

- *Issuer*: A role an entity performs by asserting claims about one or more subjects, creating a verifiable credential from these claims, and transmitting the verifiable credential to a holder.
- *Holder*: A role an entity might perform by possessing one or more verifiable credentials and generating verifiable presentations from them. Example holders include students, employees, and customers.
- *Verifier*: A role an entity performs by receiving one or more verifiable credentials, optionally inside a verifiable presentation, for processing. Example verifiers include employers, security personnel, and websites.
- *Verifiable data registry*: A role a system might perform by mediating the creation and verification of identifiers, keys, and other relevant data, such as revocation registries, issuer public keys, and so on, which might be required to use verifiable credentials.  Example verifiable data registries include trusted databases, decentralized databases, government ID databases

We restrict ourselves to provide, to all the users (Issuer, Holder, Verifier), the same interface as this model, i.e. only communicating with a Verifiable data registry.

However, they may or may not communicate to the **same** verifiable data registry and this is where IBC connects the different services provided by the registry.
This is demonstrated here:

![decentralised_vdr](./decentralised_vdr.png)

- Each holder will have as many DID as needed, these may be used once or resued as they see fit
- The Issuer - Holder identifiers may or may not be the same as Holder - Verifiable identifiers

### Role of CertX

In all use cases, **CertX** acts as a proofs provider. The proofs may be plain text or encrypted.
**CertX** itself does not provide any recognition or list of authorities to trust etc.
It is the verifier's choice as to if a proof is valid / useful or not.

**Critically**: Either the holder or the verifier needs to register on **CertX** as this will be unnecessary friction.
This usage of this service is strictly with IBC.

[roles]: https://www.w3.org/TR/vc-data-model/#roles
[*dencentralised identifiers*]: https://www.w3.org/TR/did-core/
[Use Case 1]: #use-case-1-muggles-vaccination-records
[Use Case 2]: #use-case-2-ministry-of-magic-employment-records
[Use Case 3]: #use-case-3-medical-records

---

## Use case 1 Muggles Vaccination records

- Alice wants to demonstrate to a restaurant that she has been fully vaccinated.
- The restaurant does not need to know who Alice is other than the fact she has been vaccinated.
- Alice does not want to reused the identity for the restaurant for other services as she does not want to be tracked.

### MuggleAuth (Issuer Zone)

This represents the authorities for muggles (non-magical humans).
For example, it might be a blockchain that stores and provides verifiable credential for registered citizens,
such as driver license, health record, and in this case, a vaccination record.

With potentially sensitive data, this blockchain is considered private.

Please refer [muggleAuth README](./muggleAuth/readme.md) for details.

### CertX (Proofs Provider)

This is a publicly accessible Blockchain to store proofs.

Please refer [certX README](./certX/readme.md) for details.

## Use case 2 Ministry of Magic Employment records

- Alice wants to retrieve her employment record from the Ministry of Magic for all potential employers.
- A Quiddtich team is a potential employer who wants to verify Alice's claim.
- Alice does not want her record available in public or have Ministry of Magic know where she has applied to.

### MoM (Issuer Zone)

This represents the Ministry of Magic. They keep many types of record, Alice is interested in her employment record.

With potentially sensitive data, this blockchain is considered private.

Please refer [MoM README](./MoM/readme.md) for details.

## Use case 3 Medical records

TODO

## Improvements

As this project develops  a few ideas have come up for improvements or unimplemented

- Communications between client and node should be encrypted
- did created do not have attached documents / metadata, what was created is for demonstration only

For use case 1:

- Alice should give permission for authortities to contact her / not
- Restaurant should not need to know if Alice is a muggle / witch / giant etc, zk-proof for who the issuer can protect her

