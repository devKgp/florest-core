# `florest`

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE) [![Build Status](https://travis-ci.org/jabong/florest-core.svg?branch=master)](https://travis-ci.org/jabong/florest-core)

Flo(w)Rest is a workflow based REST API framework. Each request to the REST api (defined) triggers a workflow consisting of different nodes and the request data flows through each one of them in the way that they are connected to each other.

All the backend microservices in Jabong are written on top of `florest`. Refer the [wiki](https://github.com/jabong/florest-core/wiki) for the detailed explanation. 

# Pre-requisites

1. Go 1.5+
2. Linux or MacOS
3. [gingko](https://onsi.github.io/ginkgo/) for executing the tests. If it is not already installed execute the below command:-
   
   ```go
   $ go get github.com/onsi/ginkgo/ginkgo
   $ go get github.com/onsi/gomega
   ```

## Usage

* Clone the repo:-

  ```bash
  cd <GOPROJECTSPATH>
  git clone https://github.com/jabong/florest-core
  ```

* Bootstrap the new application to be created from `florest`
  Let's assume `APPDIR` is the absolute location where new application's code will reside. For example, let's say the new application to be created is named `restapi` to be placed in `/Users/tuk/golang/src/github.com/jabong/` then `APPDIR` denotes the location `/Users/tuk/golang/src/github.com/jabong/restapi`
 
  ```bash
  cd <GOPROJECTSPATH>/florest-core
  make newapp NEWAPP="APPDIR"  
  ```
  The above will create a new app based on `florest` with the necessary structure.  
  
* Set-up the application log directory. Let's say if the application was created as mentioned in the previous step then this will look like below:-

  ```bash
  sudo mkdir /var/log/restapi/          # This can be changed
  chown <user who will be executing the app> /var/log/restapi
  ```
  
* To build the application execute the below command:-

  ```bash
  cd APPDIR
  make deploy 
  ```
  If the above command is successful then this will create a binary named after the application name. In this case the binary will be named as `restapi` The binary can be executed using the below command:-
  
  ```bash
  cd APPDIR/bin
  ./restapi
  ```


## TESTS 

* To run the tests execute the below command:-

  ```bash
  cd APPDIR
  make test
  ```
  
* To get the test coverage execute the below command:-
 
   ```bash
  cd APPDIR
  make coverage
  ```
  
* To execute all the benchmark tests:-

  ```bash
  make bench
  ```
  
* The application code can also be formatted as per gofmt. To do this execute the below command:-

  ```
  make format
  ```
  

##Examples

To run the examples execute the below command:-

```go
go get -u github.com/jabong/florest-core/examples
```

The above command will place the `examples` binary in `$GOROOT/bin` directory.

To execute the examples create a conf file named [conf.json](config/florest-core/conf.json) & [logger.json](config/florest-core/logger.json) and place it in `conf/` in the same folder where `examples` binary is placed.

To run the hello world example execute the below command

```bash
curl -XGET "http://localhost:8080/florest/v1/hello/"
```

This should produce an output like below:-

```json
{

    "status": {
        "httpStatusCode": 200,
        "success": true,
        "errors": null
    },
    "data": "Hello World",
    "_metaData": {
        "urlParams": { },
        "apiMetaData": { }
    }

}
```

## Example Config

To excute all the hello world examples involving redis (`florest/v1/redis/`), mongo(`florest/v1/mongo/`), redis cluster (`florest/v1/rediscluster/`) & mysql (`florest/v1/mysql/`). `conf.json` should look like below:-

```json
{  
   "AppName":"florest",
   "AppVersion":"1.0.0",
   "ServerPort":"8080",
   "LogConfFile":"conf/logger.json",
   "MonitorConfig":{  
      "AppName":"florest",
      "Platform":"DatadogAgent",
      "AgentServer":"datadog:8125",
      "Verbose":false,
      "Enabled":false,
      "MetricsServer":"datadog:8065"
   },
   "Performance":{  
      "UseCorePercentage":100,
      "GCPercentage":1000
   },
   "HttpConfig":{  
      "MaxConn":200,
      "MaxIdleConns":2,
      "ResponseHeaderTimeout":30,
      "DisableKeepAlives":false
   },
   "Profiler":{  
      "SamplingRate":0.6,
      "Enable": true
   },
   "ApplicationConfig":{  
      "ResponseHeaders":{  
         "CacheControl":{  
            "ResponseType":"public",
            "NoCache":false,
            "NoStore":false,
            "MaxAgeInSeconds":300
         }
      },
      "Mongo" : {
         "URL" : "mongodb://localhost:27017",
         "DbName" : "florest"
      },
      "Cache" : {
         "Redis" : {
            "Platform" : "redis",
            "ConnStr" : "localhost:6379",
            "Password" : "",
            "BucketHashes" : ["hello"],
            "Cluster" : false
         },
         "RedisCluster" : {
            "Platform" : "redis",
            "ConnStr" : ":30001,:30002,:30003,:30004,:30005,:30006",
            "Password" : "",
            "BucketHashes" : ["dog", "cat"],
            "Cluster" : true
         }
      },
      "MySQL" : {
         "DriverName" : "mysql",
         "Username" : "root",
         "Password" : "",
         "Host" : "localhost",
         "Port" : "3306",
         "Dbname" : "bobalice",
         "Timezone" : "Local",
         "MaxOpenCon" : 1,
         "MaxIdleCon" : 2
      }
   }
}
```
