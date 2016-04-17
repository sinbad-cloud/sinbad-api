# kigo-api

## Generate protobuf code

    make protobuf
     
## Architecture

### Domain

The entities of the domain are:

* apps
* deployments
* users
* builds

### Use cases

The entities for the use cases are:

* jobs

The use cases are:

* [x] create a deployment job
* [ ] deploy an app
* [ ] get status a deployment job
* [x] create a build job
* [ ] build an app
* [ ] get status a build job
* [x] create an app
* [x] retrieve an app
* [x] manage app config
* [x] sign up for an account
* [x] sign in
* [ ] reset password

### Interfaces

* rpc service

### Infrastructure

* db storage
* kubernetes jobs

