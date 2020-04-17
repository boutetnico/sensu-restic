
[![Bonsai Asset Badge](https://img.shields.io/badge/Bonsai-Download%20Me-brightgreen.svg?colorB=89C967&logo=sensu)](https://bonsai.sensu.io/assets/boutetnico/sensu-restic) [![TravisCI Build Status](https://travis-ci.org/boutetnico/sensu-restic.svg?branch=master)
](https://travis-ci.org/boutetnico/sensu-restic)

# Sensu Go Restic Plugin

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
  labels: 
  annotations:
    io.sensu.bonsai.url: https://bonsai.sensu.io/assets/boutetnico/sensu-restic
    io.sensu.bonsai.api_url: https://bonsai.sensu.io/api/v1/assets/boutetnico/sensu-restic
    io.sensu.bonsai.tier: Community
    io.sensu.bonsai.version: 1.0.0
    io.sensu.bonsai.namespace: boutetnico
    io.sensu.bonsai.name: sensu-restic
    io.sensu.bonsai.tags: restic
spec:
  builds:
  - url: https://assets.bonsai.sensu.io/261732b2653841be5cbe859ace20e256863c7867/sensu-restic_1.0.0_windows_amd64.tar.gz
    sha512: 4479ccf7afa218b4a66dce0cb2bbc3ed64ba5429e8ec46e31c57692a3c419d7887494bfab55c08c126e5eb4fbb74fad267a3fcecfd30fee10ee97ea90b604040
    filters:
    - entity.system.os == 'windows'
    - entity.system.arch == 'amd64'
  - url: https://assets.bonsai.sensu.io/261732b2653841be5cbe859ace20e256863c7867/sensu-restic_1.0.0_darwin_386.tar.gz
    sha512: 382fd0acd57fe6f7523e75a899cd3358b89d703de8c0f0c5a730e2fcd37e79aea008c376e87b68e71f865fd8bc3cfd7864854df871c5abc0a4b86db222de0ba1
    filters:
    - entity.system.os == 'darwin'
    - entity.system.arch == '386'
  - url: https://assets.bonsai.sensu.io/261732b2653841be5cbe859ace20e256863c7867/sensu-restic_1.0.0_darwin_amd64.tar.gz
    sha512: '019a3850550575c4b3a77f032b9107dcad5e8074580a2f2c6bb1ed408c28a2bd482a8fec17b9c0d1ef4c8be0c44028007d6817860d7196734ffce57c0886fe30'
    filters:
    - entity.system.os == 'darwin'
    - entity.system.arch == 'amd64'
  - url: https://assets.bonsai.sensu.io/261732b2653841be5cbe859ace20e256863c7867/sensu-restic_1.0.0_linux_armv7.tar.gz
    sha512: e153d98729a9bac1006eb5364af7570def8a117f23db268e6b5ce2d1b2efdb508da6868767e8f1bf17a914c239a43c566712e0edaa05dd86aed735e193e12760
    filters:
    - entity.system.os == 'linux'
    - entity.system.arch == 'armv7'
  - url: https://assets.bonsai.sensu.io/261732b2653841be5cbe859ace20e256863c7867/sensu-restic_1.0.0_linux_arm64.tar.gz
    sha512: 6a1679d1d11030ed57c001df790237133b1e3366f78e7ba8c44aff6fc448f812d9e896de6606e695f97d6904969bec26de82e3222e383930073a0e13640f2763
    filters:
    - entity.system.os == 'linux'
    - entity.system.arch == 'arm64'
  - url: https://assets.bonsai.sensu.io/261732b2653841be5cbe859ace20e256863c7867/sensu-restic_1.0.0_linux_386.tar.gz
    sha512: 5797b034ecabcab4dfa607fa4dd0fb5358e96f5141e4f8f9c44a213f9d6da89ef53c87f59ea465f56758e1eec6d7d06c7eff243ffa391b56c0b31bc96e8b6668
    filters:
    - entity.system.os == 'linux'
    - entity.system.arch == '386'
  - url: https://assets.bonsai.sensu.io/261732b2653841be5cbe859ace20e256863c7867/sensu-restic_1.0.0_linux_amd64.tar.gz
    sha512: 72cf4458e9d2a3daf9582fd2c84ea23bc3c40df908238cd7eaee44c48ad6890b004225f087679492acdd8ce11284aa49f35ca3f70a002cb2453008eb7e38ce79
    filters:
    - entity.system.os == 'linux'
    - entity.system.arch == 'amd64'
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
