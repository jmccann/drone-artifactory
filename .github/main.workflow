workflow "Push" {
  resolves = [
    "Docker Build",
    "Lint",
    "Test",
  ]
  on = "push"
}

action "Lint" {
  uses = "docker://golang:1.12"
  runs = "go"
  args = "fmt"
  env = {
    GO111MODULE = "on"
  }
}

action "Test" {
  uses = "docker://golang:1.12"
  runs = "go"
  args = "test -cover"
  env = {
    GO111MODULE = "on"
  }
}

action "Docker Build" {
  uses = "docker://docker:stable"
  args = "build -t jmccann/drone-artifactory:build-num ."
}
