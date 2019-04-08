workflow "Test" {
  on = "push"
  resolves = ["linting", "testing"]
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
  runs = "go"
  args = "test"
  env = {
    GO111MODULE = "on"
  }
}
