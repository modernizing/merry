# 1. maven index query api

日期: 2020-06-07

## 状态

2020-06-07 提议

## 背景

根据 ArtifactID 查询包依赖信息：

官方 API：[Digging Into Central Index with Luke](https://maven.apache.org/repository/central-index.html)

*   download [the Central index: `nexus-maven-repository-index.gz`](https://repo.maven.apache.org/maven2/.index/)
*   download [Maven Indexer CLI](https://repo.maven.apache.org/maven2/org/apache/maven/indexer/indexer-cli/5.1.1/indexer-cli-5.1.1.jar) and [unpack](https://maven.apache.org/maven-indexer-archives/maven-indexer-LATEST/indexer-cli/) the index to raw Lucene index directory:
 1.  `java -jar indexer-cli-5.1.1.jar --unpack nexus-maven-repository-index.gz --destination central-lucene-index --type full`
*   download and extract [Luke binary tarball](https://github.com/DmitryKey/luke/releases/download/luke-4.10.4/luke-with-deps.tar.gz) and launch it on the Central index with Lucene format:
 1.  `luke.sh -ro -index central-lucene-index`
You need an old Luke version using an old Lucene version, since Maven Indexer 5.5.1 uses Lucene 3.6.2: for this tutorial, we chose Luke version 4.10.4, but you may choose another version.

## 决策

在这里补充上决策信息...

## 后果

在这里记录结果...
