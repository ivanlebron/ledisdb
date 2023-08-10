# LedisDB 

[![Build Status](https://travis-ci.org/ledisdb/ledisdb.svg?branch=develop)](https://travis-ci.org/ledisdb/ledisdb) [![codecov](https://codecov.io/gh/ledisdb/ledisdb/branch/master/graph/badge.svg)](https://codecov.io/gh/ledisdb/ledisdb) [![goreportcard](https://goreportcard.com/badge/github.com/ledisdb/ledisdb)](https://goreportcard.com/report/github.com/ledisdb/ledisdb)

Ledisdb is a high-performance NoSQL database library and server written in [Go](http://golang.org). It's similar to Redis but store data in disk. It supports many data structures including kv, list, hash, zset, set.

LedisDB now supports multiple different databases as backends.

## Features

+ Rich data structure: KV, List, Hash, ZSet, Set.
+ Data storage is not limited by RAM.
+ Various backends supported: LevelDB, goleveldb, RocksDB, RAM.
+ Supports Lua scripting.
+ Supports expiration and TTL.
+ Can be managed via redis-cli.
+ Easy to embed in your own Go application. 
+ HTTP API support, JSON/BSON/msgpack output.
+ Replication to guarantee data safety.
+ Supplies tools to load, dump, and repair database. 
+ Supports cluster, use [xcodis](https://github.com/ledisdb/xcodis).
+ Authentication (though, not via http).
+ Repair integrated: You can use `ledis repair` to repair broken databases and `ledis repair-ttl` to repair a very serious bug for key expiration and TTL if you upgraded from v0.4.

## Build from source

Create a workspace and checkout ledisdb source

```shell
git clone git@github.com:ivanlebron/ledisdb.git
cd ledisdb

make
make test
```

Then you will find all the binary build on `./bin` directory.

## Choose store database

LedisDB now supports goleveldb, leveldb, rocksdb, and RAM. It will use goleveldb by default. 

Choosing a store database to use is very simple.

+ Set in server config file

    db_name = "leveldb"

+ Set in command flag

    ledis -config=/etc/ledis.conf -db_name=leveldb

Flag command set will overwrite config setting.

## Lua support

Lua is supported using [gopher-lua](https://github.com/yuin/gopher-lua), a Lua VM, completely written in Go.

## Configuration

LedisDB uses [toml](https://github.com/toml-lang/toml) as the configuration format. The basic configuration ```./etc/ledis.conf``` in LedisDB source may help you.

If you don't use a configuration, LedisDB will use the default for you.

## Server Example

```shell
//set run environment if not
./bin/ledis -config=/etc/ledis.conf

//another shell
./bin/ledis cli -p 6380

ledis 127.0.0.1:6380> set a 1
OK
ledis 127.0.0.1:6380> get a
"1"

//use curl
curl http://127.0.0.1:11181/SET/hello/world
→ {"SET":[true,"OK"]}

curl http://127.0.0.1:11181/0/GET/hello?type=json
→ {"GET":"world"}
```



## Replication Example

Set slaveof in config or dynamiclly

```shell
ledis cli -p 6381 

ledis 127.0.0.1:6381> slaveof 127.0.0.1 6380
OK
```

## Cluster support

LedisDB uses a proxy named [xcodis](https://github.com/ledisdb/xcodis) to support cluster.

## CONTRIBUTING

See [CONTRIBUTING.md] .

## Benchmark

See [benchmark](https://github.com/ledisdb/ledisdb/wiki/Benchmark) for more.

## Todo

See [Issues todo](https://github.com/ledisdb/ledisdb/issues?labels=todo&page=1&state=open)

## Client

See [Clients](https://github.com/ledisdb/ledisdb/wiki/Clients) to find or contribute LedisDB client.


## Caveat

+ Changing the backend database at runtime is very dangerous. Data validation is not guaranteed if this is done.

## Requirement

+ Go version >= 1.18


