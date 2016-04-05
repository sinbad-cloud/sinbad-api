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

* schedule a deployment job
* get status a deployment job
* schedule a build job
* get status a build job
* create an app
* retrieve an app
* manage app config
* build an app
* deploy an app
* sign up for an account
* sign in
* reset password

### Interfaces

* rpc service

### Infrastructure

* db storage
* kubernetes jobs

