use zed_extension_api as zed;

struct TerragruntLsExtension;

fn set_default_env(env: &mut Vec<(String, String)>, key: &str, value: &str) {
    if !env.iter().any(|(existing, _)| existing == key) {
        env.push((key.to_string(), value.to_string()));
    }
}

impl zed::Extension for TerragruntLsExtension {
    fn new() -> Self {
        Self
    }
    fn language_server_command(
        &mut self,
        _language_server_id: &zed_extension_api::LanguageServerId,
        worktree: &zed_extension_api::Worktree,
    ) -> zed_extension_api::Result<zed_extension_api::Command> {
        let path = worktree
            .which("terragrunt-ls")
            .ok_or_else(|| "The LSP for Terragrunt 'terragrunt-ls' is not installed".to_string())?;
        let mut env = worktree.shell_env();
        set_default_env(&mut env, "TG_LS_LOG", "terragrunt-ls.log");
        set_default_env(&mut env, "TG_LS_LOG_LEVEL", "debug");

        Ok(zed::Command {
            command: path,
            args: vec![],
            env,
        })
    }
}

zed::register_extension!(TerragruntLsExtension);
