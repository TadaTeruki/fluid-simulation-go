
{
  outputs = { self, nixpkgs }: {
    devShell.x86_64-linux = with nixpkgs.legacyPackages.x86_64-linux; mkShell {
      buildInputs = [
        # Golang
        go_1_21

        # Core X11 Libraries
        xorg.libX11.dev
        xorg.libXext
        xorg.libXinerama
        xorg.libXrandr
        xorg.libXrender
        xorg.libXcursor
        xorg.libXfixes
        xorg.libXi
        xorg.libXt
        xorg.libXxf86vm

        # Utility Libraries
        xorg.libXmu
        xorg.libXtst
        xorg.libXpm

        # Advanced Graphics
        libGL
        libGLU

        # Font Handling
        xorg.libXfont
        xorg.libXft
      ];
    };
  };  
}
