{
  pkgs ? import <nixpkgs> { },
}:

pkgs.mkShell {
  buildInputs = [
    pkgs.go
    pkgs.gopls
    pkgs.golangci-lint-langserver
    pkgs.golangci-lint
  ];
}
