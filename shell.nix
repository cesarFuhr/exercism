{ pkgs ? import <nixpkgs> { } }:
with pkgs;
mkShell {
  buildInputs = [
    exercism
    zsh
  ];

  shellHook = ''
    export SHELL="zsh"
    zsh
  '';
}

