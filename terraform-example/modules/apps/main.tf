resource "helm_release" "app1" {
  name = "app1"
  atomic = true
  chart  = "path/to/chart"
  values = [file("${path.module}/values/app1.yaml")]
  set {
    name  = "version"
    value = var.versions.app1
  }
  set {
    name  = "path.to.client_id"
    value = var.authorization.client_id
  }
  set {
    name  = "path.to.host"
    value = var.authorization.host
  }
  set {
    name  = "secret"
    value = var.some_secret
  }
}

# in a real work scenario these 2 resources can be created with a single definition and a for loop. But be careful, do not over-DRY
resource "helm_release" "app2" {
  name = "app2"
  atomic = true
  chart  = "path/to/chart"
  values = [file("${path.module}/values/app2.yaml")]
  set {
    name  = "version"
    value = var.versions.app1
  }
  set {
    name  = "path.to.client_id"
    value = var.authorization.client_id
  }
  set {
    name  = "path.to.host"
    value = var.authorization.host
  }
  set {
    name  = "secret"
    value = var.some_secret
  }
}

resource "aws_s3_bucket" "app2_media" {
  bucket = "app2_media_${var.env}"
  acl    = "private"

  tags = {
    Name        = "app2_media_${var.env}"
    Environment = var.env
  }
}
