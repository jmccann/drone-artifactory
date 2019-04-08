workflow "Test" {
  on = "push"
  resolves = ["linting"]
}

action "linting" {
  uses = "golang:1.12"
  runs = "go"
  args = "fmt"
  env = {
    GO111MODULE = "on"
  }
}
