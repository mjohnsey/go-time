workflow "Build and deploy on push" {
  on = "push"
  resolves = [
    "Setup Go for use with actions",
    "Filters for GitHub Actions",
  ]
}

action "Filters for GitHub Actions" {
  uses = "actions/bin/filter@25b7b846d5027eac3315b50a8055ea675e2abd89"
  args = "branch master"
}
