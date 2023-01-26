load("@bazel_gazelle//:deps.bzl", "go_repository")

def go_dependencies():
    go_repository(
        name = "com_github_gbdubs_maybebig",
        importpath = "github.com/gbdubs/maybebig",
        sum = "h1:9MQBoGAWN0Qvhb0O+16lTXZRm/FzjYTDvT+QRccA0DE=",
        version = "v1.0.2",
    )

    go_repository(
        name = "com_github_google_go_cmp",
        importpath = "github.com/google/go-cmp",
        sum = "h1:O2Tfq5qg4qc4AmwVlvv0oLiVAGB7enBSJ2x2DqQFi38=",
        version = "v0.5.9",
    )
