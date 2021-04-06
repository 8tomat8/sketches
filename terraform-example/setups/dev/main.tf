locals {
  name = "develop"
  authorization = {
    client_id = "foo"
    host = "https://dev.bar.com"
  }
}

data "aws_secretsmanager_secret" "external_something_secret" {
  name = "${local.name}-external_something"
}

data "aws_secretsmanager_secret_version" "external_something" {
  secret_id = data.aws_secretsmanager_secret.external_something_secret.id
}

module "service" {
  source = "../../modules/apps"
  authorization = local.authorization
  env = local.name
  some_secret = data.aws_secretsmanager_secret_version.external_something.secret_string
  versions = {
    # dev setup has newer versions of the apps
    app1 = "v2"
    app2 = "v3"
  }
} 
