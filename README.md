# Intro  
이것은 Strix 의 golang project 에 대한 기본 구조를 안내합니다.

golang 으로 새로운 project 시작 시 boilerplate 로 활용될 수 있습니다.

Strixer 라면 누구나 기여할 수 있습니다.

한 사람의 여러 발자국 보다는 여러명의 한 발자국이 더 가치있고 의미있습니다.

Strixer 의 많은 관심과 참여를 부탁드립니다.

---

# Golang environment
TODO Golang 과 관련된 설정을 기술하세요.

## Version
Default version 은 1.12 입니다.

cloud 환경에서는, 특히나 serverless(예를들면 AWS 의 Lambda, GCP 의 Functions) 와 같은 서비스에서는 최신의 golang 버젼을 즉시 지원하지 않기 때문에 유의해아합니다.

이럴때는 docker 를 이용하는 것도 하나의 방법입니다.

## Dependency manager
우리는 `go module` 을 사용하여 의존성 관리를 합니다. 

만약 `go module` 보다 더 나은 방법을 알고계신다면 이 project 를 수정해주세요.

토론을 통하여 적용여부가 결정됩니다.

---

# Project environment
TODO 채우기

### Package structure
directory 구조는 아래와 같습니다.

`[]` 는 folder 를 의미하고 그렇지 않은 경우 `file` 을 나타냅니다.

```$xslt
project
| - [pkg]
|   | - [databases]
|   | - [models]
|   | - [repositories]
|   | - [services]
|
| - [internal]
|   | - [...]
|
| - [external]
|   | - [...]
| 
| - [sxerrors]
|   | - ...
|
| - [utils]
|   | - [...]
|
| - [k8s]
|   | - prod.yaml
|   | - dev.yaml
|
| - [bin]
|   | - ...
|
| - main.go( or application.go )
| - go.mod
| - Dockerfile
| - cloudbuild.prod.yaml
| - cloudbuild.dev.yaml
|___
```

* `[pkg]` 에는 application 의 source code 를 작성합니다.
  * `[databases]` 에는 RDB, NoSQL 등 database connection 연결을 초기화 하는 등의 configuration 을 작성합니다.
  * `[models]` 필요 시 model 를 분리하여 관리할 수 있습니다. 꼭 사용할 필요는 없습니다.
  * `[repositories]` 필요 시 database 와 상호작용 하는 기능을 layer 로 분리할 수 있습니다.  
  * `[services]` business logic 이 여기에 위치해야 합니다. 일반적으로는 대부분의 기능이 여기에 위치하게 됩니다.  

* `[internal]` 에는 는 Strix 내부 서비스를 호출하는 code 를 작성합니다. 
예를 들면 SMS 를 보내는 서비스인 notification 서비스를 호출하하는 codes 를 여기에서 작성할 수 있습니다.

* `[external]` 에는 는 외부 서비스를 호출하는 code 를 작성합니다.
예를 들면 날씨를 조회하는 타사 API 를 호출하는 codes 를 여기에서 작성할 수 있습니다.

* `[sxerrors]` 에는 는 application 에서 사용되는 custom error 를 설정합니다.
더 나은 logging 을 위해 custom error 를 작성할 수 있습니다. 
이 custom errors 의 목적은 application 자체에서 발생한 오류를 더 나은 방법을 구분짓기 위함입니다.

* `[utils]` 에는 는 application 에서 빈번하게 재사용되는 code 를 작성합니다.
예를 들면 datetime 을 지정된 format 으로 해석하거나, 내보내는 기능을 작성할 수 있습니다. 

* `[k8s]` 에는 는 kubernetes 설정 yaml 파일을 저장합니다.

* `main.go` 은 이 application 의 entry point 입니다.

* `go.mod` 은 이 application 의 dependencies 를 기술합니다.

* `Dockerfile` 은 docker 이용시 docker image 를 build 하는 데 사용됩니다.

* `clouduild.*.yaml` 은 GCP 에서 Cloud build 이용시 기술하는 설정파일 입니다.

더 나은 이해를 위해 sample code 를 포함하고 있습니다. 










