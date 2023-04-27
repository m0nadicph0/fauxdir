## Fauxdir

Fauxdir is a command-line utility for creating deep nested directory structures with random names and also creating random text files with random content.

It uses the [Faker](https://github.com/jaswdr/faker) library for generating random directory and file names.

### Installation

To install Fauxdir, you need to have Go installed on your machine. Then, run the following command:

```
go get github.com/m0nadicph0/fauxdir
```

### Usage

Fauxdir is a command-line tool with two required flags:

- `-r` or `--root`: specifies the root directory where the directory structure should be created.
- `-d` or `--depth`: specifies the depth of the directory structure.

Additionally, there are two optional flags:

- `-n` or `--num-files`: specifies the number of files to create in each directory.
- `-s` or `--size`: specifies the size of the files to create in bytes.

To use Fauxdir, run the following command:

```
fauxdir -r /path/to/root/dir -d 5
```

This command will create a directory structure with a depth of 5 in the `/path/to/root/dir` directory.

To create files in each directory, you can use the `-n` flag:

```
fauxdir -r /path/to/root/dir -d 5 -n 10
```

This command will create a directory structure with a depth of 5 in the `/path/to/root/dir` directory and 10 random text files with random content in each directory.

To specify the size of the files to create, use the `-s` flag:

```
fauxdir -r /path/to/root/dir -d 5 -n 10 -s 1024
```

This command will create a directory structure with a depth of 5 in the `/path/to/root/dir` directory, 10 random text files with random content in each directory, and each file will have a size of 1024 bytes.

### Sample Usage

Create a directory structure with a depth of 3 and 5 random text files in each directory:

```
fauxdir -r /path/to/root/dir -d 3 -n 5
```

Create a directory structure with a depth of 4 and 10 random text files in each directory, with each file having a size of 2048 bytes:

```
fauxdir -r /path/to/root/dir -d 4 -n 10 -s 2048
```