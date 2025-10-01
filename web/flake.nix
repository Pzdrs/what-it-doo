{
  inputs = {
    nixpkgs.url = "nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs =
    {
      self,
      nixpkgs,
      flake-utils,
      ...
    }:
    flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = import nixpkgs {
          inherit system;
        };
        buildDeps = with pkgs; [
          nodejs_24
          pnpm_9
        ];
        devDeps = buildDeps ++ [ ];
      in
      {
        devShell = pkgs.mkShell {
          buildInputs = devDeps;
          shellHook = ''
            alias npm="echo 'Use pnpm instead!'"
          '';
        };
      }
    );
}
