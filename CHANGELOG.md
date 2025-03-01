# Changelog

## [v0.1.18] - 19.03.2021
### Added
- Ante handler to check frozen tokens movement
- Add network properties for ENABLE_TOKEN_WHITELIST / ENABLE_TOKEN_BLACKLIST
- Add permissions for creation and vote on blacklist and whitelist change of tokens
- Added CLI command to submit proposal to change blacklist and whitelist
- Added CLI command to query current blacklist and whitelist
- Network Properties management code modification for boolean properties

### Fixed
## [v0.1.17.6] - 18.03.2021
### Added
- Now when the proposal passes it enter ins status Enactment.
- Add proposal to create a Role.
- Fix GetTxProposalUpsertDataRegistry and make it appear on client.
### Fixed
- Fix and clean some CLI commands (proposal upsert token rates, proposal upsert token alias, proposal upsert data registry).

## [v0.1.17.3] - 03.08.2021
### Fixed
- Validators query to include mischance.

### Added
- Tokens alias/rate query.
- Voters/votes query.

## [v0.1.17.2] - 03.02.2021
### Fixed
- Mischance querying CLI command
### Added
- genutil module to handle validator status
- CLI utility command to get valcons address from account address
- ValidatorJoined hook that's derivated from Cosmos SDK's `ValidatorBonded` hook

## [v0.1.17.1] - 02.25.2021
### Added
- GRPC query for proposals
- GRPC query for validators + validator_signing_infos

## [v0.1.17] - 02.16.2021
### Fixed
- Problem with ClaimValidator and PubKey encoding due to protocol buff Any type.
- Fix bug that made that you can vote when the proposal ended.
- When a proposal does not reach quorum it ends being Quorum not reached.
- Proposal voting time and enactment time now are defined in seconds.
- It shows the votes that a proposal has on client query.

## [v0.1.16.2a] - 02.08.2021
### Added
- Custom evidence module to jail a double signed validator 
- CLI command for writing proposal to unjail a validator
- CLI command for setting max jail time network property proposal

## [v0.1.16.1] - 02.04.2021
### Added
- CLI command to set poor network messages
- CLI command to query poor network messages
- Add POOR_NETWORK_MAX_BANK_TX_SEND feature for poor network for restriction (only bond denom is allowed)
- Reject feature for not allowed messages on poor network

## [v0.1.15.2] - 01.21.2021

### Added
- CLI command GetTxProposalUpsertTokenAliasCmd and GetTxProposalUpsertTokenRatesCmd are now exposed.

## [v0.1.15.1] - 01.21.2021

### Added
- CLI command to get ValAddress from AccAddress

## [0.1.15] - 01.15.2021
### Added
- Added custom slashing module for validator's block signing info management, inactivate, activate, pause, unpause
- Added validator performance calculator using `rank` and `streak`
- Upgraded Cosmos SDK to v0.40.0 (stargate)

### Removed
- Old staking, slashing, evidence, distribution module

## [0.1.14.3] - 12.30.2020
### Added
- Added GRPC query for Data Reference Registry.
- Update response caching for data references. (KIP_47.1)
- Added file hosting feature. (KIP_47.1)

## [0.1.14.2] - 11.30.2020
### Added
- Update Cosmos SDK to v0.40.0-rc4.

## [0.1.14.1] - 11.30.2020
### Added
- Proposal to upsert the Token Rates. (CLI too)

## [0.1.14] - 11.26.2020
### Added
- Added a wrapper to register messages with function metadata.
- Added function_id for message types.
- Registered function meta for existing messages.
- Added INTERX api for Kira functions list.
- Added INTERX api for INTERX functions list.

## [0.1.13] - 11.20.2020
### Added
- Proposal to upsert the Token Aliases KIP_24. (CLI too)

## [0.1.12.1] - 11.15.2020
### Added
- Proposal to upsert the Data Registry. (CLI too)
- Proposal to change Network Properties. (CLI too)

### Changed
- Now it is more generic to be able to add new proposals in the complete flow.

## [0.1.12] - 11.13.2020

### Added

KIP_8
- Added grpc gateway
- Added status, balances, transaction hash queries
- Added transaction encode/broadcast
- Added response format

KIP_9
- Added endpoint whitelist

KIP_48
- Added INTERX faucet

KIP_47
- Added response caching

KIP_32
- Added withdraws, deposits

## [0.1.7.3] - 11.12.2020
### Added
- Added CLI for querying proposals / individual proposal
- Added CLI for querying votes / individual vote
- Added CLI for querying whitelisted proposal voters

### Changed
- Updated genesis actor initialization process
- Updated proposal end time and enactment time
- Fixed end blocker concert not registered issue for MsgClaimValidator

## [0.1.7.2] - 11.11.2020

### Changed
- There is a new permission for all role related changes. PermUpsertRole.

## [0.1.7.1] - 11.09.2020

### Changed
- Proposal is now a generic type, the Content part is what changes between different proposal types.

## [0.1.7] - 11.09.2020
- Added CLI command to upsert token rates per denom
- Added CLI commands to query token rates
- Implemented feeprocessing module for new fee processing logic
- Implemented foreign currency fee payment

## [0.1.6.3] - 11.07.2020

### Added
- We can propose to SetPermissions to an actor.
- We can vote a proposal to SetPermissions to an actor.
- Added proposal endtime and proposal enactment time into 
the network properties.

## [0.1.6.2] - 10.23.2020

### Added
- The keeper has method to get all Actors by witelisted permission.
- The keeper has method to get All actors that have specific role.
- The keeper has method to get all roles that have a whitelist permission.

### Changed
- Big refactor on the way Role and Permissions are stored.
- In keeper we don't expose SetPermissionsForRole anymore.

## [0.1.6.1] - 10.19.2020
### Added

- Added CLI command to send a SetPermission proposal.
- Added CLI command to vote a SetPermission proposal.

### Changed

- Now Role and Permissions are persisted differently in order to be able to get
actors by permission and actors by role.

- Now the commands for all Governance module is simplified in a better hierarchical style.
```
Available Commands:
  councilor   Councilor subcommands
  permission  Permission subcommands
  proposal    Proposal subcommands
  role        Role subcommands
```

## [0.1.6] - 10.16.2020
### Added

- Added CLI command to upsert token alias per symbol
- Added CLI commands to query token aliases per symbol and denom
- Added CLI command to query all token aliases

### Modified

- Modified execution fee to use transaction type as identifier
- Modified min/max fee range to [100 - 1'000'000] in ukex

## [0.1.4] - 10.05.2020
### Added

- Added CLI command to change execution fee per message type
- Added CLI command to change transaction fee range
- Added CLI command to query execution fee
- Added CLI command to query transaction fee range

## [0.1.2.4] - 09.24.2020
### Added

- Added CLI command to remove blacklist permissions into a specific role.
- Added CLI command to create new role.
- Added CLI command to assign new role.
- Added CLI command to remove assignation for a role.

## [0.1.2.3] - 09.17.2020
### Changed

- Updated cosmos SDK to last version of 17th september .

## [0.1.2.2] - 09.14.2020
### Added

- Added CLI command to claim governance seat.
- Added CLI command to set whitelist permissions into a specific role.
- Added CLI command to set blacklist permissions into a specific role.
- Added CLI command to remove whitelist permissions into a specific role.

## [0.1.2.1] - 09.06.2020
### Added

- Added CLI command to Set Blacklist Permissions too.
- Module CustomGov defines in genesis by default Permissions by roles Validator (0x2) and Sudo (0x1).

### Changed
- Now the roles are validated when taking some action. It checks if the user has permissions either in the role or individually.

## [0.1.2] - 09.01.2020
### Added

- Add command to add whitelist permissions to an address, that address is included
in theNetworkActor registry with the specified permission added.
- Now the user that generates the network has AddPermissions by default, so he is the only one
that can add permissions into the registry.

### Changed

- Now the ClaimValidator message takes care that the user has ClaimValidator permissions,
if he the user does not have, it fails.
