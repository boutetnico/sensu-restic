
[![Bonsai Asset Badge](https://img.shields.io/badge/CHANGEME-Download%20Me-brightgreen.svg?colorB=89C967&logo=sensu)](https://bonsai.sensu.io/assets/boutetnico/sensu-restic) [![TravisCI Build Status](https://travis-ci.org/boutetnico/sensu-restic.svg?branch=master)
](https://travis-ci.org/boutetnico/sensu-restic)

# Sensu Go Restic Plugin

TODO: Table of Contents

- [Overview](#overview)
- [Usage examples](#usage-examples)
- [Configuration](#configuration)
  - [Asset registration](#asset-registration)
  - [Asset configuration](#asset-configuration)
  - [Resource check configuration](#resource-configuration)
- [Functionality](#functionality)
- [Installation from source and contributing](#installation-from-source-and-contributing)

## Overview

This plugin allows to check the age of the latest Restic snapshot. 

## Configuration

Example usage:

```bash
./sensu-restic -c 86400 -w 3600
./sensu-restic -c 86400 -w 3600 -f "--tag mysql --no-cache --no-lock"
```

### Asset Registration

Assets are the best way to make use of this plugin. If you're not using an asset, please consider doing so! If you're using sensuctl 5.13 or later, you can use the following command to add the asset: 

`sensuctl asset add boutetnico/sensu-restic:VERSION`

If you're using an earlier version of sensuctl, you can find the asset on the [Bonsai Asset Index](https://bonsai.sensu.io/assets/boutetnico/sensu-restic).

### Asset configuration

```yml
---
type: Asset
api_version: core/v2
metadata:
  name: sensu-restic
spec:
  url: https://CHANGEME
  sha512: CHANGEME
```

### Resource check configuration

Example Sensu Go definition:

```yml
---
api_version: core/v2
type: CHANGEME
metadata:
  namespace: default
  name: CHANGEME
spec:
  "...": "..."

```

### Functionality

This plugin requires [Restic](https://github.com/restic/restic) to be installed.

## Installation from source and contributing

The preferred way of installing and deploying this plugin is to use it as an [asset]. If you would like to compile and install the plugin from source or contribute to it, download the latest version of the sensu-restic from [releases][1]
or create an executable script from this source.

From the local path of the sensu-restic repository:

```
go build -o /usr/local/bin/sensu-restic main.go
```

For more information about contributing to this plugin, see https://github.com/sensu/sensu-go/blob/master/CONTRIBUTING.md

[1]: https://github.com/boutetnico/sensu-restic/releases
[2]: #asset-registration
