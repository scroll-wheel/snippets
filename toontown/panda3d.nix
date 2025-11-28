{
  stdenv,
  # autoPatchelfHook,
  eigen,
  fetchFromGitHub,
  freetype,
  glibcLocales,
  gtk3,
  interrogate,
  libjpeg_original,
  libpng,
  libtiff,
  libvorbis,
  ode,
  openal,
  openexr,
  openssl,
  pkg-config,
  python311,
  lib,
  zlib,
}:

stdenv.mkDerivation {
  pname = "panda3d";
  version = "b8586733e1f8362997da21004a4cdc21b2c1236e";

  src = fetchFromGitHub {
    owner = "rocketprogrammer";
    repo = "panda3d";
    rev = "b8586733e1f8362997da21004a4cdc21b2c1236e";
    sha256 = "sha256-YJOzUWR9olF5bvg/CTM7D85ToF0z60YeMI6VDbaCExs=";
  };

  nativeBuildInputs = [
    # autoPatchelfHook
    glibcLocales
    interrogate
    pkg-config
    (python311.withPackages (pkgs: [
      pkgs.pip
    ]))
  ];

  buildInputs = [
    zlib
    libjpeg_original
    libpng
    openssl
    libvorbis
    freetype
    gtk3
    openal
    eigen
    libtiff
    openexr

    ode
  ];

  LIBRARY_PATH = lib.makeLibraryPath [
    python311
    libjpeg_original
  ];

  buildPhase = ''
    runHook preBuild
    INTERROGATE=$(command -v interrogate) \
    INTERROGATE_MODULE=$(command -v interrogate_module) \
    python3 makepanda/makepanda.py --everything --threads=4
    runHook postBuild
  '';

  installPhase = ''
    runHook preInstall
    rm -rf built/temp
    mkdir $out
    mv built/* $out
    runHook postInstall
  '';
}
