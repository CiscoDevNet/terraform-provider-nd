
# Cisco ND Provider

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) Latest Version
- [Go](https://golang.org/doc/install) Latest Version

## Building The Provider

Clone this repository to: `$GOPATH/src/github.com/CiscoDevNet/terraform-provider-nd`.

```sh
$ mkdir -p $GOPATH/src/github.com/CiscoDevNet; cd $GOPATH/src/github.com/CiscoDevNet
$ git clone https://github.com/CiscoDevNet/terraform-provider-nd.git
```

Enter the provider directory and run dep ensure to install all the dependencies. After, that run make build to build the provider binary.

```sh
$ cd $GOPATH/src/github.com/CiscoDevNet/terraform-provider-nd
$ dep ensure
$ make build
```

## Using The Provider

If you are building the provider, follow the instructions to [install it as a plugin.](https://www.terraform.io/docs/cli/plugins/index.html) After placing it into your plugins directory, run `terraform init` to initialize it.

Example:
```hcl
terraform {
  required_providers {
    nd = {
      source = "ciscodevnet/nd"
    }
  }
}

# Configure provider with your Cisco ND credentials.
provider "nd" {
  username  = "admin"
  password  = "password"
  url       = "https://my-cisco-nd.com"
  insecure  = true
  proxy_url = "https://proxy_server:proxy_port"
}

resource "nd_site" "example" {
  name         = "example"
  username     = "admin"
  password     = "password"
  url          = "10.195.219.154"
  type         = "aci"
  inband_epg   = "test_epg"
  latitude     = "19.36475238603211"
  longitude    = "-155.28865502961474"
  login_domain = "local"
}
```

Note: If you are facing the issue of `invalid character '<' looking for beginning of value` while running `terraform apply`, use `-parallelism=1` with `terraform plan` and `terraform apply` to limit the concurrency to one thread.

```
terraform plan -parallelism=1
terraform apply -parallelism=1
```

## Developing The Provider

Currently the ND provider uses [terraform-plugin-framework](https://developer.hashicorp.com/terraform/plugin/framework) to create new resource and data-source files.

### Developing the new resources and data-sources with terraform-plugin-framework

1. Create issue (if not created yet) and comment that you will be working on the issue.

2. Fork the `terraform-provider-nd` repository.

3. Clone the forked code to your local machine.

4. Make changes to the files manually.
    * Code changes
      * New resources and data-sources must be stored in the [internal/provider](https://github.com/CiscoDevNet/terraform-provider-nd/tree/master/internal/provider) directory.
    * Examples changes
      * Examples must be stored in the [examples/resources](https://github.com/CiscoDevNet/terraform-provider-nd/tree/master/examples/resources) and [examples/data-sources](https://github.com/CiscoDevNet/terraform-provider-nd/tree/master/examples/data-sources) directories.
    * Documentation changes
      * Documentation for resources and data-sources must be stored in the [docs](https://github.com/CiscoDevNet/terraform-provider-nd/tree/master/docs) directory.

5. Test the code.
    * Set the below environment variables when running tests:
      ```sh
      export ND_VAL_REL_DN=false
      export TF_ACC=1
      export ND_USERNAME="USER"
      export ND_URL="https://IPADDRESS"
      export ND_PASSWORD="PASSWORD"
      ```
    * The following command can be used `go test internal/provider/* -v -run <test-name>`, where the test name can be found in the `resource_<resource-name>_test.go` and `data_source_<resource-name>_test.go` files.
    * Execute the tests for all your resources and data-sources

6. Create PR for the code and request review from active maintainers.

7. Review process

### Troubleshooting 

#### Missing packages

If you encounter an error indicating that the golang.org/x/text/language package is missing from your vendor directory, you can fetch it by following these steps:

- Make sure that you're running the latest version of [Go](https://golang.org/doc/install)

- Update dependencies and populate the vendor directory: If you're using Go modules, you can update your dependencies and populate your vendor directory by running the following commands in your terminal:
  ```sh
  go mod tidy
  go mod vendor
  ```
  The `go mod tidy` command will clean up unused dependencies and add missing ones. The `go mod vendor` command will copy all dependencies into the vendor directory.

- Disable vendor mode: Go operates in vendor mode when the -mod=vendor flag is set. You'll need to disable vendor mode to fetch packages directly. Run the following command in your terminal:
  ```sh
  export GOFLAGS="-mod=mod"
  ```

- Fetch the missing package: Now that vendor mode is disabled, you can fetch the missing package by running:
  ```sh
  go get golang.org/x/text/language
  ```
  This command tells Go to fetch the golang.org/x/text/language package directly, regardless of the vendor directory

- Re-enable vendor mode (if necessary): If you wish to switch back to vendor mode, you can do so by running:
  ```sh
  export GOFLAGS="-mod=vendor"
  ```

### Compiling

To compile the provider, run `make build`. This will build the provider with sanity checks present in scripts directory and put the provider binary in `$GOPATH/bin` directory.

<strong>Important: </strong>To successfully use the provider you need to follow these steps:

- Copy or Symlink the provider from the `$GOPATH/bin` to `~/.terraform.d/plugins/terraform.local/CiscoDevNet/nd/<Version>/<architecture>/` for example:
  ```bash
  ln -s ~/go/bin/terraform-provider-nd ~/.terraform.d/plugins/terraform.local/CiscoDevNet/nd/1.0.0/linux_amd64/terraform-provider-nd
  ```
- Edit the Terraform Provider Configuration to use the local provider.
  ```hcl
  terraform {
    required_providers {
      nd = {
        source  = "terraform.local/CiscoDevNet/nd"
        version = "1.0.0"
      }
    }
  }
  ```