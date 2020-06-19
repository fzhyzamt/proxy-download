
def _move_file(ctx):
    print(ctx.bin_dir)
    print(ctx.build_file_path)
    print(ctx.file)
    print(ctx.files)
    print(ctx.genfiles_dir)
    print(ctx.info_file)
    print(ctx.outputs)
    print(ctx.actions)


move_target_file = rule(
    _move_file,
    attrs = {
        "target_folder": attr.string(),
    },
)