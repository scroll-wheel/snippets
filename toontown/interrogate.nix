{
  stdenv,
  cmake,
  fetchFromGitHub,
  python311,
}:

stdenv.mkDerivation rec {
  pname = "interrogate";
  version = "0.4.0";

  src = fetchFromGitHub {
    owner = "panda3d";
    repo = "interrogate";
    rev = "v${version}";
    sha256 = "sha256-DcthlD6wHnrd/OQLdSLKWXz1cG63tvddKbjcElpNiEg=";
  };

  nativeBuildInputs = [ cmake python311 ];
}
