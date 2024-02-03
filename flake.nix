{
  description = "fosdem24 - where am I webpage";

  inputs = {
    nixpkgs.url = "nixpkgs/nixpkgs-unstable";
    utils.url = "github:numtide/flake-utils";
    templ.url = "github:a-h/templ";
  };

  outputs = {
    self,
    nixpkgs,
    utils,
    templ,
    ...
  }: let
    fosdem24Version =
      if (self ? shortRev)
      then self.shortRev
      else "dev";
  in
    {
      overlay = _: prev: let
        pkgs = nixpkgs.legacyPackages.${prev.system};
        templ = system: templ.packages.${prev.system}.templ;
      in {
        fosdem24 = pkgs.buildGoModule {
          pname = "fosdem24";
          version = fosdem24Version;
          src = pkgs.nix-gitignore.gitignoreSource [] ./.;

          preBuild = ''
            ${templ prev.system}/bin/templ generate
          '';

          vendorSha256 = "";
        };
      };
    }
    // utils.lib.eachDefaultSystem
    (system: let
      templf = system: templ.packages.${system}.templ;
      pkgs = import nixpkgs {
        overlays = [
          self.overlay
          templ.overlays.default
        ];
        inherit system;
      };
      buildDeps = with pkgs; [
        git
        gnumake
        go
      ];
      devDeps = with pkgs;
        buildDeps
        ++ [
          golangci-lint
          entr
          elmPackages.elm
          (templf system)
        ];
    in rec {
      # `nix develop`
      devShell = pkgs.mkShell {
        buildInputs = with pkgs;
          [
            (writeShellScriptBin
              "fosdem24run"
              ''
                go run .
              '')
            (writeShellScriptBin
              "fosdem24dev"
              ''
                ls *.go | entr -r fosdem24run
              '')
          ]
          ++ devDeps;
      };

      # `nix build`
      packages = with pkgs; {
        inherit fosdem24;
      };

      defaultPackage = pkgs.fosdem24;

      # `nix run`
      apps = {
        fosdem24 = utils.lib.mkApp {
          drv = packages.fosdem24;
        };
      };

      defaultApp = apps.fosdem24;

      overlays.default = self.overlay;
    });
}
