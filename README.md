# UnrealProjectManager
A cross-platform Unreal Engine project manager built with Go and Fyne. It lets you list, run, build, clean, and manage your Unreal projects with ease — fast, lightweight, and developer-friendly.

Status: **Active Development**

## Projects and engines setup
Edit the `catalog.json` file with your projects and unreal installation path, then run the app.

## Development setup
```bash
go get fyne.io/fyne/v2@latest
go mod tidy
```

### On Windows
On Windows, you’ll need to install `mingw-w64` and add it to your system PATH to enable project compilation and command execution.

Here is how you can install the requirements using the command line

```bash
winget install -e --id MSYS2.MSYS2

C:\msys64\usr\bin\bash -lc "pacman -Sy --noconfirm && pacman -S --noconfirm mingw-w64-x86_64-gcc mingw-w64-x86_64-pkgconf make"

# if you don't want to add mingw to the system PATH, run this each time
$env:Path = "$env:Path;C:\msys64\mingw64\bin"

# One time only
go env -w CGO_ENABLED=1
```

### Running the app
```bash
go run .
``` 

### Building the app
```bash
go build .
./ue_launcher
```

## License
This project is released under the MIT license.
