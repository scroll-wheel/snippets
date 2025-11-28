{
  stdenv,
  fetchFromGitHub,
}:

stdenv.mkDerivation rec {
  pname = "resources";
  version = "d8c73a9978633979ddf2ef8813f0152037a0d978";

  src = fetchFromGitHub {
    owner = "open-toontown";
    repo = "resources";
    rev = "${version}";
    sha256 = "sha256-rZ6j8WW0D+0JEoyEYD2Nek2Bxmeb/r9/oBDcL3Dpajc=";
  };

  installPhase = ''
    runHook preInstall
    mkdir $out
    mv * $out
    runHook postInstall
  '';
}
