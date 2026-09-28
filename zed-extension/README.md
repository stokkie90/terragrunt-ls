# zed-terragrunt

[Zed extension](https://zed.dev/docs/extensions/installing-extensions) for [terragrunt-ls](https://github.com/gruntwork-io/terragrunt-ls), mostly based on [terraform extension](https://github.com/zed-extensions/terraform)

## Configuration

Zed associates this extension with `.hcl` files. To match the VS Code extension more closely, you can map the common Terragrunt filenames explicitly:

```json
{
  "file_types": {
    "Terragrunt": [
      "terragrunt.hcl", 
      "terragrunt.stack.hcl",
      "terragrunt.values.hcl",
      "root.hcl"
    ]
  }
}
```

## Format on Save

Zed can format Terragrunt files on save through the language server. Enable it in your Zed settings:

```json
{
  "languages": {
    "Terragrunt": {
      "formatter": "language_server",
      "format_on_save": "on"
    }
  }
}
```

Use `"formatter": "language_server"` so Zed sends formatting requests to `terragrunt-ls` instead of falling back to another formatter such as Prettier.

When formatting is requested, `terragrunt-ls` now prefers `terragrunt hcl fmt` and falls back to the built-in HCL formatter if the `terragrunt` CLI is unavailable.

## Logging and Troubleshooting

The Zed extension now forwards your shell environment to `terragrunt-ls`, so commands such as `terragrunt` can be found the same way they are in your shell.

It also enables a default log file:

- `TG_LS_LOG=terragrunt-ls.log`
- `TG_LS_LOG_LEVEL=debug`

Unless you override those variables yourself, the language server will write debug logs to `terragrunt-ls.log`.

If format on save still does not run, verify that:

1. `formatter` is set to `"language_server"` for `Terragrunt`.
2. `format_on_save` is enabled for `Terragrunt`.
3. `terragrunt-ls` is installed and available on your `PATH`.
4. `terragrunt` is installed and available on your `PATH`.
5. The file is associated with the `Terragrunt` language via `file_types`.
