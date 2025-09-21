{
  outputs =
    { self, nixpkgs, ... }:
    let
      system = "aarch64-darwin";
      pkgs = import nixpkgs { inherit system; };
    in
    {
      devShells.${system}.default = pkgs.mkShell { 
        buildInputs = with pkgs; [
          nodejs
        ];
      };
    };
}
