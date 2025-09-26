{
  description = "Dev shell for Node.js on macOS (arm64) and Linux (x86_64)";

  outputs = { self, nixpkgs, ... }:
    let
      # Pick the correct pkgs for each host automatically
      forAllSystems = systems: f:
        builtins.listToAttrs (map (system: {
          name = system;
          value = f system;
        }) systems);

      supportedSystems = [ "aarch64-darwin" "x86_64-linux" ];
    in
    {
      devShells = forAllSystems supportedSystems (system:
        let pkgs = import nixpkgs { inherit system; };
        in {
          default = pkgs.mkShell {
            buildInputs = with pkgs; [ nodejs ];

            shellHook = ''
              npm install
            '';
          };
        });
    };
}