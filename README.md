# igso

> a legacy system migration tools for OSGI based architecture Java applications.

![Go](https://github.com/phodal/igso/workflows/Go/badge.svg)
[![Build Status](https://travis-ci.org/phodal/igso.svg?branch=master)](https://travis-ci.org/phodal/igso)
[![codecov](https://codecov.io/gh/phodal/igso/branch/master/graph/badge.svg)](https://codecov.io/gh/phodal/igso)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/phodal/igso)
[![Maintainability](https://api.codeclimate.com/v1/badges/77b3f3f4a2444b33dc16/maintainability)](https://codeclimate.com/github/phodal/igso/maintainability)

Todo:

 - [x] Ant to Maven
    - [x] dependencies convert
    - [x] Checksum File
    - [x] Unzip file
    - [x] Parse Manifest
 - [ ] Ant to Gradle
 - [ ] Manifest convert
    - [ ] to BND
 - [ ] Microservices
 - [x] Dup Search. search all jars and find not in nexus package

Todo 2：

 - [ ] POM 遍历生成依赖关系图
 - [x] 支持自动重命名 jar
   - [x] POM 替换关系版本号 jar
 - [x] 支持定制版本号
   - [x] 配置化
   - [ ] 根据 jar 生成 map.csv，以进行修改

Feature lists:

```bash
Available Commands:
  boom        generate pom.xml from build.xml
  checksum    checksum file md5
  dupsearch   build maven pom from all jars file
  fix         fix jar naming issue
  help        Help about any command
  manifest    manifest query & map tools
  pom         generate pom file from jar file
  version     version

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

4. map file example

format: `{origin},{GroupId},{ArtifactId},{VersionId}`

```
origin,groupid,artifactId,version
org.springframework.transaction,org.springframework,transaction.org.springframework.transaction,2.5.6.SEC01
javax.servlet,javax.servlet,javax.servlet,2.4.0
edu.emory.mathcs.backport.java.util,edu.emory.mathcs.backport,com.springsource.edu.emory.mathcs.backport,3.1.0
edu.emory.mathcs.backport.java.util.concurrent,edu.emory.mathcs.backport,com.springsource.edu.emory.mathcs.backport,3.1.0
edu.emory.mathcs.backport.java.util.concurrent.atomic,edu.emory.mathcs.backport,com.springsource.edu.emory.mathcs.backport,3.1.0
edu.emory.mathcs.backport.java.util.concurrent.helpers,edu.emory.mathcs.backport,com.springsource.edu.emory.mathcs.backport,3.1.0
edu.emory.mathcs.backport.java.util.concurrent.locks,edu.emory.mathcs.backport,com.springsource.edu.emory.mathcs.backport,3.1.0
org.apache.commons.dbcp,org.apache.commons,com.springsource.org.apache.commons.dbcp,1.2.2.osgi
org.apache.commons.logging,commons-logging,commons-logging,1.1.1
org.slf4j,org.slf4j,org.slf4j,1.5.1
```

### Generate Pom from jars

```
go run igso.go pom -p _fixtures/demo -m _fixtures/map/map.csv
```

### Call Graph by Manifest.MF

```
igso manifest -s -p ~/sdk/apache-karaf-4.2.9/lib/boot  -x
igso map -p ~/sdk/apache-karaf-4.2.9/lib/boot
igso call -p ~/sdk/apache-karaf-4.2.9/lib/boot -m map.csv
```

results:

![Call Graph](docs/screenshots/call-example-karaf.svg)

## Issues

Lost Package

 - javax.net
 - javax.servlet
 - org.osgi.framework
 - org.slf4j

License
---

@ 2020 A [Phodal Huang](https://www.phodal.com)'s [Idea](http://github.com/phodal/ideas).  This code is distributed under the MPL license. See `LICENSE` in this directory.
