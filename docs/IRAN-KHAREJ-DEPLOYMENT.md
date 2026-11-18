# Iran ↔ Kharej deployment profile

This document records the intended two-server deployment for this fork. It does not claim that a deployment has been tested on the owner's VPS pair.

## Version and installation source

- Fork module: `github.com/mfranjbar/WatTunnel-Private`
- Fork code line: `import/upstream-3.3.7`
- Fork version string: `3.3.8`
- Install script source: this fork's `import/upstream-3.3.7/install.sh`, not the upstream JohnWatsoon repository.
- The current CI artifact is unsigned and is not a signed GitHub Release. The installer verifies signed release manifests; do not bypass that check for production.

Install the same fork build on both Linux hosts. Create the tunnel on Kharej first, copy its token securely to Iran, then create the Iran side. Start with one dedicated test port mapping; do not migrate production listeners until bidirectional tests pass.

## Intended traffic path

```
Users
  |
  | existing client protocol (for example VLESS/Reality or WireGuard)
  v
Iran VPS (existing X-UI/Xray/WireGuard remain untouched)
  |
  | WatTunnel transport, initiated according to the selected carrier/reverse mode
  v
Kharej VPS
  |
  v
Internet / destination services
```

WatTunnel forwards selected TCP/UDP ports between the two hosts. It does not replace X-UI, Xray, or WireGuard. The user-facing service remains on Iran; the corresponding destination service/port is reached on Kharej as configured.

## Important distinction: direct vs reverse

- Direct mode: Iran initiates the carrier connection to Kharej.
- Reverse mode: Iran listens and Kharej dials Iran. Use this only when the real network path requires Kharej-initiated connections.
- Do not assume reverse is faster. It changes connection direction; benchmark both only when both are reachable and permitted.

## Deployment safety requirements

1. Do not change existing X-UI/Xray/WireGuard listeners or firewall rules as part of the WatTunnel install.
2. Select dedicated, unused carrier ports. Verify the port is unused on both hosts before starting.
3. Restrict the carrier listener to the peer address where the carrier and topology support source-IP filtering. Do not expose metrics publicly.
4. Keep the tunnel token/key private. Do not put it in shell history, issue text, or committed configuration.
5. Back up the existing service and firewall state before any host-side change. Apply host tuning only as an explicit, reversible step.
6. Start with one mapped test port and validate both directions before moving production traffic.

## Baseline and acceptance test

Record the following independently for direct and reverse (where applicable):

- Carrier and direction
- RTT before and through the tunnel
- TCP throughput in both directions (single stream and multiple streams)
- UDP loss and throughput, if UDP forwarding is required
- CPU, memory, retransmits, and interface drops on both VPSs
- Stability over at least 30 minutes under representative load
- Recovery time after intentionally interrupting the carrier

A result is accepted only when the same endpoints, payload sizes, stream counts, duration, and test direction are used for the before/after comparison. Record raw output; do not infer performance from README benchmark tables.

## Current implementation status

This file is a deployment design and acceptance checklist only. The fork identifies itself as 3.3.8 and exposes a configurable direct-carrier startup window (`-direct-win-ms`, 100–5000 ms; default behavior remains unchanged). This exposes a timing parameter; it is not evidence of an Iran-specific transport algorithm improvement. No live two-VPS validation is claimed.
