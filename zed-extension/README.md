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
      "format_on_save": "on"
    }
  }
}
```

When formatting is requested, `terragrunt-ls` now prefers `terragrunt hcl fmt` and falls back to the built-in HCL formatter if the `terragrunt` CLI is unavailable.
