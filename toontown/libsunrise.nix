{
  python311Packages,
  fetchFromGitLab,
}:

python311Packages.buildPythonPackage {
  pname = "libsunrise";
  version = "0.1";

  src = fetchFromGitLab {
    owner = "sunrisemmos";
    repo = "libsunrise";
    rev = "97b7e2a79ffa9c73ba061fc91c0de6b40c8dcba9";
    sha256 = "sha256-jpxuayH8IBKNZRIezpgHcw4vWSCoCaIMTjCkwl9/bHM=";
  };

  pyproject = false;
}
