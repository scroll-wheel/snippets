{
  stdenv,
  cmake,
  fetchFromGitHub,
  python311,
  resources,
}:

stdenv.mkDerivation rec {
  pname = "open-toontown";
  version = "38e88e5b3a95d7b04f63e0c7d01fa66353c54f16";

  src = fetchFromGitHub {
    owner = "rocketprogrammer";
    repo = "open-toontown";
    rev = "${version}";
    sha256 = "sha256-O84QDSFK3ZZg+LfEVr39q1rQ/Fd0pk1utVDHcunQEH4=";
  };

  patches = [ ./home_directory.patch ];

  installPhase = ''
    runHook preInstall
    mkdir $out
    mv * $out
    cp -r ${resources} $out/resources
    runHook postInstall
  '';
}
