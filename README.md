# Reqres(Terraform Custom Provider)
----------------------------------

This is a custom provider of terraform based on Golang for an online fake API '[Reqres](https://reqres.in/)'

### Prerequisites
------------------

-  Linux based machine
-  [Terraform](https://www.terraform.io/downloads.html) 0.10.x
-	 [Go](https://golang.org/doc/install) 1.9 (to build the provider plugin)

### Building the Provider
------------------------

Create a directory for GoWorkspace and for the session, set the environmant variable 'GOPATH' pointing to it

```sh
$ mkdir ~/goWorkspace; cd ~/goWorkspace
$ export GOPATH=$(pwd)
```

Clone repository to: `$GOPATH/src/github.com/hashicorp/terraform-providers/terraform-provider-reqres`

```sh
$ mkdir -p $GOPATH/src/github.com/hashicorp/terraform-providers
$ cd $GOPATH/src/github.com/hashicorp/terraform-providers
$ git clone https://github.com/mehrosejethi/terraform-provider-reqres.git
```

Clone the Terraform library to: `$GOPATH/src/github.com/hashicorp/terraform'

```sh
$ cd $GOPATH/src/github.com/hashicorp/
$ git clone https://github.com/hashicorp/terraform.git
```

Enter the provider directory and build the provider. A plugin by the name of 'terraform-provider-reqres' is created in the same folder

```sh
$ cd $GOPATH/src/github.com/hashicorp/terraform-providers/terraform-provider-reqres
$ go build -o terraform-provider-reqres
```

### Using the Provider
----------------------

Copy the provider build in the last step and place it in the same directory in which the terraform configuration files are written

```
$ cp $GOPATH/src/github.com/hashicorp/terraform-providers/terraform-provider-reqres/terraform-provider-reqres terraform_config_file_dir
```

Create configuration files involving the provider 'reqres'

### Provider details
--------------------

#### Resources

- **reqres_create_user**: Provides a resource to call the POST method of the api(*/api/users*) to create user in reqres.in returing id and creation time.
  - Arguements: The following arguments are supported:
    - *request_body(required)*: It is stringified JSON request for the body of the POST method. Eg. "{"name": "morpheus","job": "leader"}
  - Reference: In addition to all arguments above, the following attributes are exported:
    - *id*: The id generated for the resource(timestamp)
    - *body*: The response of the request
    
#### Data Sources

- **data_source_single_user**: This data source fetches the details of a user on reqres.in based on its ID.
  - Arguements: The following arguments are supported:
    - *user_id(required)*: The user id of which details are to be fetched
  - Reference: In addition to all arguments above, the following attributes are exported:
    - *id*: The id generated for the data source(timestamp)
    - *body*: The response of the request
    
- **data_source_users_per_page**: This data source fetches the details of all the users on reqres.in based on the page number.
  - Arguements: The following arguments are supported:
    - *page(required)*: The page number on which details are to be fetched
  - Reference: In addition to all arguments above, the following attributes are exported:
    - *id*: The id generated for the data source(timestamp)
    - *body*: The response of the request
    
### Screenshots
---------------

#### Configuration File

![Alt text](screenshots/main.PNG?raw=true "main.tf")

#### Terraform Plan

![Alt text](screenshots/plan.PNG?raw=true "main.tf")

#### Terrafrom Apply(Execution)

![Alt text](screenshots/apply.PNG?raw=true "main.tf")

#### Output in JSON

![Alt text](screenshots/output_json.PNG?raw=true "main.tf")
