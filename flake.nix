# bingo-backend/flake.nix
{
  description = "Development environment for the Go Bingo Backend";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable"; # Or a specific stable branch like nixos-23.11
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = nixpkgs.legacyPackages.${system};

        # --- Go Toolchain & Dev Tools ---
        goVersion = pkgs.go; # Use the default Go version from nixpkgs
        goTools = [
          pkgs.gopls # Go Language Server
          pkgs.golangci-lint # Linter
          pkgs.air # Live reloader (alternative: pkgs.realize)
          # Add other Go tools if needed (e.g., delve for debugging)
          # pkgs.delve
        ];

        # --- Runtime Dependencies (for local dev/testing) ---
        runtimeDeps = [
          pkgs.redis
          pkgs.mongodb
          # Add mongo cli tools if needed
          # pkgs.mongodb-compass-community # GUI (might be large)
          # pkgs.mongosh # Modern Mongo Shell
        ];

        # --- Shell & Environment Tools ---
        shellTools = [
          pkgs.zsh
          pkgs.oh-my-zsh # Provides OMZ installation scripts/files
          pkgs.git
          pkgs.jq # Useful for JSON manipulation
          pkgs.curl
					pkgs.openssh
        ];

      in
      {
        devShells.default = pkgs.mkShell {
          name = "bingo-backend-dev";

          packages = [
            goVersion
          ] ++ goTools ++ runtimeDeps ++ shellTools;

          # This hook runs when you enter the shell via 'nix develop'
          shellHook = ''
            echo "--- Entering Bingo Backend Dev Shell ---"

            # Set Go paths (optional with modules, but can help some tools)
            # export GOPATH=$(pwd)/.go
            # export GOBIN=$GOPATH/bin
            # export PATH=$PATH:$GOBIN

            # --- Zsh & Oh My Zsh Configuration ---
            # Use a temporary directory for Zsh config within the Nix environment
            export ZDOTDIR=$(mktemp -d)
            export ZSH_CACHE_DIR="$ZDOTDIR/cache"
            mkdir -p "$ZSH_CACHE_DIR"

            # Set path to Oh My Zsh installation provided by Nix package
            export ZSH="${pkgs.oh-my-zsh}/share/oh-my-zsh"

            # Create a minimal .zshrc in the temporary ZDOTDIR
            cat <<EOF > $ZDOTDIR/.zshrc
            # Oh My Zsh settings
            export ZSH="$ZSH"
            export ZSH_CACHE_DIR="$ZSH_CACHE_DIR"

            # Set theme (e.g., robbyrussell, agnoster, etc.)
            ZSH_THEME="robbyrussell"

            # Set plugins (git is essential)
            plugins=(git go golang) # Add other plugins if needed

            # Source Oh My Zsh
            source "$ZSH/oh-my-zsh.sh"

            # Add custom aliases or exports below if needed
            # alias k=kubectl

            # Add Go bin path if GOPATH/GOBIN are set above
            # export PATH="\$PATH:\$GOBIN"

            echo "-> Oh My Zsh configured for this session."
            EOF

            # Set the shell to Zsh for this environment
            export SHELL="${pkgs.zsh}/bin/zsh"
            echo "-> Default shell set to Zsh."

            # Start Zsh if the current shell isn't already Zsh
            # This ensures OMZ loads correctly when entering via 'nix develop'
            if [ -z "$ZSH_VERSION" ]; then
              exec $SHELL
            fi

            echo "----------------------------------------"
          '';
        };
      });
}
