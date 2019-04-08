workflow "Test" {
  on = "push"
  resolves = ["testing"]
}

action "linting" {
  uses = "docker://golang:1.12"
  runs = "go"
  args = "fmt"
  env = {
    GO111MODULE = "on"
  }
}

action "testing" {
  uses = "docker://golang:1.12"
  needs = ["linting"]
  runs = "go"
  args = "test"
  env = {
    GO111MODULE = "on"
  }
}
