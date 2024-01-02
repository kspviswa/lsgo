# lsgo

## ls command in golang for fun & learning

### Usage

```
./ls --help

NAME:
   ls ( implemented in golang ) - ls [flags] [command][args]

USAGE:
   ls [global options] command [command options] [arguments...]

VERSION:
   0.1

AUTHOR:
   Viswa Kumar

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   -l, --long             include extended information
   -d, --dronly           include only directories
   -f, --fileonly         include only regular files
   --hr, --humanfriendly  view size in humar readable format
   -t, --tree             view tree structure ( recursive lookup )
   --help, -h             show help
   --version, -v          print the version

COPYRIGHT:
   MIT Licensed
```

### Sample outputs

```
./ls -hr
+-----------+--------+------------+---------------------+-------+
|   NAME    |  SIZE  |   PERMS    |         AT          |  DIR  |
+-----------+--------+------------+---------------------+-------+
| .git      | 4K     | drwxrwxrwx | 2017-10-22 01:15:14 | true  |
| demo.json | 116.5K | -rwxrwxrwx | 2017-10-22 01:02:45 | false |
| LICENSE   | 1.1K   | -rwxrwxrwx | 2017-10-21 19:34:57 | false |
| ls        | 3.5M   | -rwxrwxrwx | 2017-10-22 00:31:26 | false |
| ls.go     | 3.2K   | -rwxrwxrwx | 2017-10-22 00:31:23 | false |
| README.md | 166B   | -rwxrwxrwx | 2017-10-22 01:11:07 | false |
+-----------+--------+------------+---------------------+-------+
```
