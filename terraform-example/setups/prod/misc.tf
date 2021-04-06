provider "aws" {
  region = "eu-central-1" # defulat region of the setup
  profile = "terraform" # name of an AWS profile to use by this setup

}

terraform {
  required_version = "~> 0.14.0"

  # backend "s3" {
    # bucket = "terraform-state"
    # key    = "path/to/state-file.tfstate"
    # region = "eu-central-1"
  # }

  required_providers {
    kubernetes = {
      version = "= 2.0.2"
    }

    aws = {
      version = "= 3.25.0"
    }

    helm = {
      version = "= 2.0.2"
    }
  }
}
