<div align="center"><h3>tldr2cheat</h3></div>

tldr2cheat is a Go tool that
- converts a tldr file to a Cheat sheet.
- converts all files in a specific folder and drops the converted files to another folder

#### Build:

```sh
$ go build
```

#### Usage:

```sh
$ ./tldr2cheat -h
Usage of ./tldr2cheat:
  -file string
        File to convert
  -infolder string
        Input folder to use
  -outfolder string
        Output folder to use
```

tldr2cheat might be executed with
- `-file` flag
- both `-infolder` and `-outfolder` flags

#### Examples:

##### Convert a single file:

```
$ ./tldr2cheat -file ~/.tldrc/tldr-master/pages/linux/cp.md
- Copy a file to another location:
cp path/to/source_file.ext path/to/target_file.ext

- Copy a file into another directory, keeping the filename:
cp path/to/source_file.ext path/to/target_parent_directory

- Recursively copy a directory's contents to another location (if the destination exists, the directory is copied inside it):
cp -r path/to/source_directory path/to/target_directory

- Copy a directory recursively, in verbose mode (shows files as they are copied):
cp -vr path/to/source_directory path/to/target_directory

- Copy text files to another location, in interactive mode (prompts user before overwriting):
cp -i *.txt path/to/target_directory

- Dereference symbolic links before copying:
cp -L link path/to/target_directory

- Use the full path of source files, creating any missing intermediate directories when copying:
cp --parents source/path/to/file path/to/target_file
```

##### Convert all the files in a specific folder

```sh
$ ./tldr2cheat -infile ~/.tldrc/tldr-master/pages/linux -outfolder ~/.config/cheat/cheatsheets/personal/
570 files converted
```
