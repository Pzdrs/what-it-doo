{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
    goflake.url = "github:sagikazarmark/go-flake";
    goflake.inputs.nixpkgs.follows = "nixpkgs";
  };

  outputs =
    {
      self,
      nixpkgs,
      flake-utils,
      goflake,
      ...
    }:
    flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = import nixpkgs {
          inherit system;
          overlays = [ goflake.overlay ];
        };

        backendDeps = with pkgs; [
          git
          go_1_25
          go-swag
          air
          sqlc
          goose
        ];

        frontendDeps = with pkgs; [
          nodejs_24
          pnpm_9
        ];

      in
      {
        devShells.default = pkgs.mkShell {
          buildInputs = backendDeps ++ frontendDeps;
          shellHook = ''
            alias npm="echo 'Use pnpm instead!'"
          '';
          GOOSE_DRIVER = "postgres";
          GOOSE_MIGRATION_DIR = "infra/db/migrations";
        };
      }
    );
}
