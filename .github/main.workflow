workflow "Build" {
  on = "release"
  resolves = [
    # "release darwin/amd64",
    # "release windows/amd64",
    "release linux/amd64",
  ]
}

action "release darwin/amd64" {
  uses = "zoidbergwill/go-release.action@v2.0.0"
  env = {
    GOOS = "darwin"
    GOARCH = "amd64"
  }
  secrets = ["GITHUB_TOKEN"]
}

action "release windows/amd64" {
  uses = "zoidbergwill/go-release.action@v2.0.0"
  env = {
    GOOS = "windows"
    GOARCH = "amd64"
  }
  secrets = ["GITHUB_TOKEN"]
}

action "release linux/amd64" {
  uses = "zoidbergwill/go-release.action@v2.0.0"
  env = {
    GOOS = "linux"
    GOARCH = "amd64"
  }
  secrets = ["GITHUB_TOKEN"]
}
