
# go_houdini
A faster way to create a `.go` file with basic structure, imports, and function stubs.

<img src="https://github.com/luizfiuzaa/go_houdini/doc/assets/gopher.png" width="150">

## Features

- Quickly scaffold a new Go file with:
  - File name prompt
  - Imports (space separated)
  - Function stubs (space separated)
- Prevents overwriting existing files
- Adds a comment with project info

## Usage

1. **Build the script:**
   ```sh
   go build -o go_houdini.exe houdini.go
   ```

2. **Run the script:**
   ```sh
   ./go_houdini.exe
   ```
   Follow the prompts to enter:
   - The file name (no need to add `.go`)
   - Packages to import (space separated, or press Enter to skip)
   - Function names (space separated, or press Enter to skip)

3. The script will generate a Go file with the specified structure.

## Setting as an Environment Variable (Windows)

To use `go_houdini` from any terminal window:

1. Copy `go_houdini.exe` to a folder, e.g., `C:\tools\go_houdini`.
2. Add this folder to your system `PATH`:
  - Press `Win + R`, type `sysdm.cpl`, and press Enter.
  - Go to the **Advanced** tab and click **Environment Variables**.
 - Under **System variables**, select `Path` and click **Edit**.
  - Click **New** and add `C:\tools\go_houdini`.
  - Click **OK** to save.

3. Open a new terminal and run:
   ```sh
   go_houdini
   ```

Now you can use `go_houdini` from any directory.