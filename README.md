# igso

> a legacy system migration tools for OSGI based architecture Java applications.

![Go](https://github.com/phodal/igso/workflows/Go/badge.svg)

Todo:

 - [x] Ant to Maven
    - [x] dependencies convert
    - [x] Checksum File
    - [x] Unzip file
    - [ ] Parse Manifest
 - [ ] Ant to Gradle
 - [ ] Manifest convert
    - [ ] to BND
 - [ ] Microservices
 - [x] Dup Search. search all jars and find not in nexus package

Feature lists:

```bash
COMMANDS:
   boom       build maven dep from build.xml
   dupsearch  search unkonw jar source
   checksum   checksum file
   help, h    Shows a list of commands or help for one command

```

## Usage

1. install

```bash
go get -u github.com/phodal/igso
```

2. generate `pom.xml` from `build.xml`

```bash
igso boom
```

3. search package lost in nexus

```bash
igso dupsearch
```

License
---

@ 2020 A [Phodal Huang](https://www.phodal.com)'s [Idea](http://github.com/phodal/ideas).  This code is distributed under the MPL license. See `LICENSE` in this directory.
