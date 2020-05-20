# Ratatoskr

[Ratatoskr](https://en.wikipedia.org/wiki/Ratatoskr) is (well, eventually *will be*) a mesh routing protocol with the following features:

- Transport protected by WireGuard
- True peer-to-peer with no supervisory nodes (you need to connect to one node to be introduced to the rest of the network plus your ID must be approved by administrator and not blacklisted).
- Automatic node discovery
- Nodes addressed in mesh via a private [RFC4193](https://tools.ietf.org/html/rfc4193) IPv6 network.
- Route setup via link-state protocol (topology yet to be decided, probobly hierarchical mesh based on seniority/link speed/being behind NAT)
- Work from behind NAT if at least one node with public IP is available.
- A simple PKI basing on ED25519 cryptography to grant access to the network
  - Master certificate `MaC` for a network
  - Net certificate `NeC` signed by `MaC`
  - Station certificates `StC_xx` signed by `NeC`
  - Configuration certificate `CoC` signed by `MaC`
  - Administrator certificates `AdC_xx` signed by `CoC`
  - Configuration frames injected into network signed by `AdC_xx` to be accepted:
    - Station revocation
- Local management via HTTP protocol
- Initial scale for ~2000 nodes per mesh.
- Concept of mesh bridging (node belonging to more than one mesh and passing messages between them).

