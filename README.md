# SnapScheduler

[![Build
Status](https://github.com/backube/snapscheduler/workflows/Tests/badge.svg)](https://github.com/backube/snapscheduler/actions?query=branch%3Amaster+workflow%3ATests+)
[![Go Report
Card](https://goreportcard.com/badge/github.com/backube/snapscheduler)](https://goreportcard.com/report/github.com/backube/snapscheduler)
[![codecov](https://codecov.io/gh/backube/snapscheduler/branch/master/graph/badge.svg)](https://codecov.io/gh/backube/snapscheduler)

SnapScheduler provides scheduled snapshotting capabilities for Kubernetes
CSI-based volumes.

## Installation

Interested in giving it a try? [Check out the
docs.](https://backube.github.io/snapscheduler/)

The operator can be installed from:

- [Artifact
  Hub](https://artifacthub.io/package/chart/backube-helm-charts/snapscheduler)
- [Helm Hub](https://hub.helm.sh/charts/backube/snapscheduler)
- [OperatorHub.io](https://operatorhub.io/operator/snapscheduler)

Have feedback? Got questions? Having trouble? [![Gitter
chat](https://badges.gitter.im/backube/snapscheduler.png)](https://gitter.im/backube/snapscheduler)

## Helpful links

- [SnapScheduler Changelog](CHANGELOG.md)
- [Contributing guidelines](https://github.com/backube/.github/blob/master/CONTRIBUTING.md)
- [Organization code of conduct](https://github.com/backube/.github/blob/master/CODE_OF_CONDUCT.md)

## Licensing

This project is licensed under the [GNU AGPL 3.0 License](LICENSE) with the following
exception:

- The files within the `pkg/apis/*` directories are additionally licensed under
  Apache License 2.0. This is to permit SnapScheduler's CustomResource types to
  be used by a wider range of software.
- Documentation is made available under the [Creative Commons
  Attribution-ShareAlike 4.0 International license (CC BY-SA
  4.0)](https://creativecommons.org/licenses/by-sa/4.0/)
