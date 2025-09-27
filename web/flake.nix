{
  inputs = {
    nixpkgs.url = "nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils, ... }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs {
          inherit system;
        };
        buildDeps = with pkgs; [ nodejs_24 ];
        devDeps = with pkgs; buildDeps ++ [];
      in
      { devShell = pkgs.mkShell { buildInputs = devDeps; }; });
}