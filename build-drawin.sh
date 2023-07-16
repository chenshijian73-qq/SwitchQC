version="1.0.1"

wails build -platform darwin/amd64 -clean
tar -czvf ~/Downloads/SwitchQC-drawin-amd64-${version}.tar.gz build/bin/SwitchQC.app

wails build -platform darwin/arm64 -clean
tar -czvf ~/Downloads/SwitchQC-drawin-arm64-${version}.tar.gz build/bin/SwitchQC.app