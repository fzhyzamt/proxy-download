load("@io_bazel_rules_go//go:def.bzl", "go_binary")
# load("//:matrix.bzl", "build_all_platform")
load("//:files.bzl", "move_target_file")


# build_all_platform(
#     name = 'proxy-download',
#     pkg = glob(["src/*.go"]),
# )

# go_binary(
#       name = 'main',
#       srcs = glob(["src/*.go"]),
# )

go_binary(
      name = 'linux-amd64',
      srcs = glob(["src/*.go"]),
      goos = 'linux',
      goarch = 'amd64',
      out = 'linux-amd64',
)

go_binary(
      name = 'linux-386',
      srcs = glob(["src/*.go"]),
      goos = 'linux',
      goarch = '386',
      out = 'linux-386',
)

go_binary(
      name = 'windows-amd64',
      srcs = glob(["src/*.go"]),
      goos = 'windows',
      goarch = 'amd64',
      out = 'win-amd64.exe',
)

go_binary(
      name = 'darwin-amd64',
      srcs = glob(["src/*.go"]),
      goos = 'darwin',
      goarch = 'amd64',
      out = 'darwin-amd64',
)

move_target_file(
    name = 'move_file',
    target_folder = ":linux-amd64"
)