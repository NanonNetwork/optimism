<div align="center">
  <br />
  <br />
  <a href="https://optimism.io"><img alt="Optimism" src="https://raw.githubusercontent.com/ethereum-optimism/brand-kit/main/assets/svg/OPTIMISM-R.svg" width=600></a>
  <h3><a href="https://optimism.io">Optimism</a> is a low-cost and lightning-fast Ethereum L2 blockchain, built with the OP Stack.</h3>
  <br />
  <h3>+</h3>
  <a href="https://celestia.org"><img alt="Celestia" src="docs/op-stack/src/assets/docs/understand/Celestia-logo-color-color.svg" width=600></a>
  <h3><a href="https://celestia.org">Celestia</a> is a modular data availability network that securely scales with the number of users, making it easy for anyone to launch their own blockchain.</h3>
  <br />
</div>

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**

- [What is Optimism?](#what-is-optimism)
- [Documentation](#documentation)
- [Specification](#specification)
- [Community](#community)
- [Contributing](#contributing)
- [Security Policy and Vulnerability Reporting](#security-policy-and-vulnerability-reporting)
- [Directory Structure](#directory-structure)
- [Development and Release Process](#development-and-release-process)
  - [Overview](#overview)
  - [Production Releases](#production-releases)
  - [Development branch](#development-branch)
- [License](#license)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## What is Optimism?

If you're looking to run the OP Stack + Celestia setup for this repository, please visit the [Optimism & Celestia guides and tutorials](https://docs.celestia.org/developers/intro-to-op-stack/) to get started.

[Optimism](https://www.optimism.io/) is a project dedicated to scaling Ethereum's technology and expanding its ability to coordinate people from across the world to build effective decentralized economies and governance systems. The [Optimism Collective](https://app.optimism.io/announcement) builds open-source software for running L2 blockchains and aims to address key governance and economic challenges in the wider cryptocurrency ecosystem. Optimism operates on the principle of **impact=profit**, the idea that individuals who positively impact the Collective should be proportionally rewarded with profit. **Change the incentives and you change the world.**

The OP Stack powers Optimism, an Ethereum L2 blockchain, and forms the technical foundation for the [the Optimism Collective](https://app.optimism.io/announcement)—a group committed to the **impact=profit** principle. This principle rewards individuals for their positive contributions to the collective.

Optimism addresses critical coordination failures in the crypto ecosystem, such as funding public goods and infrastructure. The OP Stack focuses on creating a shared, open-source system for developing new L2 blockchains within the proposed Superchain ecosystem, promoting collaboration and preventing redundant efforts.

As Optimism evolves, the OP Stack will adapt, encompassing components ranging from blockchain infrastructure to governance systems. This software suite aims to simplify L2 blockchain creation while supporting the growth and development of the Optimism ecosystem.

## What is Celestia?

Celestia is a modular consensus and data network, built to enable anyone to easily deploy their own blockchain with minimal overhead.

Celestia is a minimal blockchain that only orders and publishes transactions and does not execute them. By decoupling the consensus and application execution layers, Celestia modularizes the blockchain technology stack and unlocks new possibilities for decentralized application builders. Lean more at [Celestia.org](https://celestia.org).

## Maintenance

The maintenance guide for this repository can be found in the Wiki tab of the repository or [here](https://github.com/celestiaorg/optimism/wiki).

## Documentation

- If you want to build on top of OP Mainnet, refer to the [Optimism Documentation](https://docs.optimism.io)
- If you want to build your own OP Stack based blockchain, refer to the [OP Stack Guide](https://docs.optimism.io/stack/getting-started), and make sure to understand this repository's [Development and Release Process](#development-and-release-process)

## Specification

If you're interested in the technical details of how Optimism works, refer to the [Optimism Protocol Specification](https://github.com/ethereum-optimism/specs).
If you want to build on top of Celestia, take a look at the documentation at [docs.celestia.org](https://docs.celestia.org).

If you want to learn more about the OP Stack, check out the documentation at [stack.optimism.io](https://stack.optimism.io/).

## Community

### Optimism

General discussion happens most frequently on the [Optimism discord](https://discord.gg/optimism).
Governance discussion can also be found on the [Optimism Governance Forum](https://gov.optimism.io/).

### Celestia

General discussion happens most frequently on the [Celestia discord](https://discord.com/invite/YsnTPcSfWQ).
Other discussions can be found on the [Celestia forum](https://forum.celestia.org).

## Contributing

Read through [CONTRIBUTING.md](./CONTRIBUTING.md) for a general overview of our contribution process.
Use the [Developer Quick Start](./CONTRIBUTING.md#development-quick-start) to get your development environment set up to start working on the Optimism Monorepo.
Then check out the list of [Good First Issues](https://github.com/ethereum-optimism/optimism/issues?q=is:open+is:issue+label:D-good-first-issue) to find something fun to work on!
Typo fixes are welcome; however, please create a single commit with all of the typo fixes & batch as many fixes together in a PR as possible. Spammy PRs will be closed.

## e2e testing

This repository has updated end-to-end tests in the `op-e2e` package to work with
Celestia as the data availability (DA) layer.

Currently, the tests assume a working [Celestia devnet](https://github.com/rollkit/local-celestia-devnet) running locally:

```bash
docker run -p 26650:26650 ghcr.io/rollkit/local-celestia-devnet:v0.12.7
```

The e2e tests can be triggered with:

```bash
cd $HOME/optimism
cd op-e2e
OP_E2E_DISABLE_PARALLEL=true OP_E2E_CANNON_ENABLED=false OP_NODE_DA_RPC=localhost:26650 OP_BATCHER_DA_RPC=localhost:26650 make test
```

## Bridging

If you have the OP Stack + Celestia setup running, you can test out bridging from the L1
to the L2.

To do this, first navigate to the `packages/contracts-bedrock` directory and create a
`.env` file with the following contents:

```bash
L1_PROVIDER_URL=http://localhost:8545
L2_PROVIDER_URL=http://localhost:9545
PRIVATE_KEY=bf7604d9d3a1c7748642b1b7b05c2bd219c9faa91458b370f85e5a40f3b03af7
```

Then, run the following from the same directory:

```bash
npx hardhat deposit --network devnetL1 --l1-provider-url http://localhost:8545 --l2-provider-url http://localhost:9545 --amount-eth <AMOUNT> --to <ADDRESS>
```

## Directory Structure

<pre>
├── <a href="./docs">docs</a>: A collection of documents including audits and post-mortems
├── <a href="./op-batcher">op-batcher</a>: L2-Batch Submitter, submits bundles of batches to L1
├── <a href="./op-bindings">op-bindings</a>: Go bindings for Bedrock smart contracts.
├── <a href="./op-bootnode">op-bootnode</a>: Standalone op-node discovery bootnode
├── <a href="./op-chain-ops">op-chain-ops</a>: State surgery utilities
├── <a href="./op-challenger">op-challenger</a>: Dispute game challenge agent
├── <a href="./op-e2e">op-e2e</a>: End-to-End testing of all bedrock components in Go
├── <a href="./op-heartbeat">op-heartbeat</a>: Heartbeat monitor service
├── <a href="./op-node">op-node</a>: rollup consensus-layer client
├── <a href="./op-preimage">op-preimage</a>: Go bindings for Preimage Oracle
├── <a href="./op-program">op-program</a>: Fault proof program
├── <a href="./op-proposer">op-proposer</a>: L2-Output Submitter, submits proposals to L1
├── <a href="./op-service">op-service</a>: Common codebase utilities
├── <a href="./op-ufm">op-ufm</a>: Simulations for monitoring end-to-end transaction latency
├── <a href="./op-wheel">op-wheel</a>: Database utilities
├── <a href="./ops">ops</a>: Various operational packages
├── <a href="./ops-bedrock">ops-bedrock</a>: Bedrock devnet work
├── <a href="./packages">packages</a>
│   ├── <a href="./packages/chain-mon">chain-mon</a>: Chain monitoring services
│   ├── <a href="./packages/common-ts">common-ts</a>: Common tools for building apps in TypeScript
│   ├── <a href="./packages/contracts-bedrock">contracts-bedrock</a>: Bedrock smart contracts
│   ├── <a href="./packages/contracts-ts">contracts-ts</a>: ABI and Address constants
│   ├── <a href="./packages/core-utils">core-utils</a>: Low-level utilities that make building Optimism easier
│   ├── <a href="./packages/fee-estimation">fee-estimation</a>: Tools for estimating gas on OP chains
│   ├── <a href="./packages/sdk">sdk</a>: provides a set of tools for interacting with Optimism
│   └── <a href="./packages/web3js-plugin">web3js-plugin</a>: Adds functions to estimate L1 and L2 gas
├── <a href="./proxyd">proxyd</a>: Configurable RPC request router and proxy
├── <a href="./specs">specs</a>: Specs of the rollup starting at the Bedrock upgrade
└── <a href="./ufm-test-services">ufm-test-services</a>: Runs a set of tasks to generate metrics
</pre>

## Development and Release Process

### Overview

Please read this section if you're planning to fork this repository, or make frequent PRs into this repository.

### Production Releases

Production releases are always tags, versioned as `<component-name>/v<semver>`.
For example, an `op-node` release might be versioned as `op-node/v1.1.2`, and  smart contract releases might be versioned as `op-contracts/v1.0.0`.
Release candidates are versioned in the format `op-node/v1.1.2-rc.1`.
We always start with `rc.1` rather than `rc`.

For contract releases, refer to the GitHub release notes for a given release, which will list the specific contracts being released—not all contracts are considered production ready within a release, and many are under active development.

Tags of the form `v<semver>`, such as `v1.1.4`, indicate releases of all Go code only, and **DO NOT** include smart contracts.
This naming scheme is required by Golang.
In the above list, this means these `v<semver` releases contain all `op-*` components, and exclude all `contracts-*` components.

`op-geth` embeds upstream geth’s version inside it’s own version as follows: `vMAJOR.GETH_MAJOR GETH_MINOR GETH_PATCH.PATCH`.
Basically, geth’s version is our minor version.
For example if geth is at `v1.12.0`, the corresponding op-geth version would be `v1.101200.0`.
Note that we pad out to three characters for the geth minor version and two characters for the geth patch version.
Since we cannot left-pad with zeroes, the geth major version is not padded.

See the [Node Software Releases](https://docs.optimism.io/builders/node-operators/releases) page of the documentation for more information about releases for the latest node components.
The full set of components that have releases are:

- `chain-mon`
- `ci-builder`
- `ci-builder`
- `indexer`
- `op-batcher`
- `op-contracts`
- `op-challenger`
- `op-heartbeat`
- `op-node`
- `op-proposer`
- `op-ufm`
- `proxyd`
- `ufm-metamask`

All other components and packages should be considered development components only and do not have releases.

### Development branch

The primary development branch is [`develop`](https://github.com/ethereum-optimism/optimism/tree/develop/).
`develop` contains the most up-to-date software that remains backwards compatible with the latest experimental [network deployments](https://community.optimism.io/docs/useful-tools/networks/).
If you're making a backwards compatible change, please direct your pull request towards `develop`.

**Changes to contracts within `packages/contracts-bedrock/src` are usually NOT considered backwards compatible.**
Some exceptions to this rule exist for cases in which we absolutely must deploy some new contract after a tag has already been fully deployed.
If you're changing or adding a contract and you're unsure about which branch to make a PR into, default to using a feature branch.
Feature branches are typically used when there are conflicts between 2 projects touching the same code, to avoid conflicts from merging both into `develop`.

## License

All other files within this repository are licensed under the [MIT License](https://github.com/ethereum-optimism/optimism/blob/master/LICENSE) unless stated otherwise.
