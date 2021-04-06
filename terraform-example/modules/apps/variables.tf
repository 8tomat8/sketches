variable some_secret {
  type = string
}

variable versions {
  type = object({
    app1=string
    app2=string
  })
  # Default is useful to be able to define only
  # applications that are needed on a particular setup
  default = {
    app1="undefined"
    app2="undefined"
  }
}

variable authorization {
  type = object({ client_id=string, host=string })
}

variable env {
  type = string
}
