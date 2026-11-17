# Building and signing a WatTunnel fork release

This fork pins its own RSA public key in both `install.sh` and `wattun-manager`.
The matching private key must remain offline and must never be committed to this repository.

## Build

The `Fork CI` workflow runs Go tests, checks shell syntax, builds the Linux amd64 engine, and uploads an unsigned artifact containing:

- `wattun`
- `wattun-manager`
- `install.sh`
- `wattun.sha256`
- `SHA256SUMS`

## Sign the release checksum manifest

After downloading and extracting the CI artifact, sign the exact `SHA256SUMS` file with the private key that matches the public key pinned in the two scripts:

```sh
openssl dgst -sha256 -sign release-private.pem -out SHA256SUMS.sig SHA256SUMS
openssl dgst -sha256 -verify release-public.pem -signature SHA256SUMS.sig SHA256SUMS
```

The verification command must print `Verified OK`. Upload `wattun`, `wattun.sha256`, `wattun-manager`, `install.sh`, `SHA256SUMS`, and `SHA256SUMS.sig` together to a GitHub Release. Do not change any payload after signing.

## Important

- The private key is not stored in GitHub, Actions, or this repository.
- The current CI artifact is not yet a signed release and must not be used as an online release.
- The installer verifies the signed manifest and checks the engine and manager hashes before installing release assets.
- The first supported CI release target is Linux amd64. No VPS has been modified by this workflow.
