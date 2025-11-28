let
  sources = import ./npins;
  pkgs = import sources.nixpkgs {
    config = { };
    overlays = [ ];
  };
in
rec {
  interrogate = pkgs.callPackage ./interrogate.nix { };
  panda3d = pkgs.callPackage ./panda3d.nix { inherit interrogate; };

  resources = pkgs.callPackage ./resources.nix { };
  toontown = pkgs.callPackage ./toontown.nix { inherit resources; };

  libsunrise = pkgs.callPackage ./libsunrise.nix { };

  launcher = pkgs.writeShellApplication (
    let
      python = pkgs.python311.withPackages (python-pkgs: [
        python-pkgs.pytz
        python-pkgs.pypresence
        python-pkgs.requests
        libsunrise
      ]);
    in
    {
      name = "launcher";

      runtimeInputs = [ python ];
      runtimeEnv.PYTHONPATH = "${panda3d}";

      text = ''
        cd ${toontown}
        python3 -m launcher
      '';
    }
  );
}
