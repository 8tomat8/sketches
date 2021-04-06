locals {
  name = "production"
  authorization = {
    client_id = "foo"
    host = "https://foo.bar.com"
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
    app1 = "v1"
    app2 = "v2"
  }
} 
