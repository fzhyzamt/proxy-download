load("@io_bazel_rules_go//go:def.bzl", "go_binary")

SUPPORTED_MATRIX = [
  ("windows", "amd64"),
  ("darwin", "amd64"),
  ("linux", "amd64"),
  ("linux", "386"),
]

def _build(ctx):
    for goos, goarch in SUPPORTED_MATRIX:
        target_name = 'proxy-download-' + goos + '-' + goarch
        if goos == 'windows':
            target_name += '.exe'

        go_binary(
            name = target_name,
            srcs = ctx.attr.pkg,
            pure = "auto",
            goos = goos,
            goarch = goarch,
        )

build_all_platform = rule(
    _build,
     attrs = {
        'pkg': attr.string_list(),
      },
      executable = True,
)
