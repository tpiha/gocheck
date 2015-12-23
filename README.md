# Go servers checks

High performance servers checks written in Go. Gocheck is easily extensible through the JSON file.

```
Usage of gocheck:
  -c uint
        Concurrency; how many parallel connections to make (default 5)
  -d string
        Name of the checks definitions JSON file (default "definitions.json")
  -f string
        Name of the checks JSON file (default "checks.json")
  -s string
        Name of the servers JSON file (default "servers.json")
  -v    Show verbose output
```

## Adding check definitions

For example, if you want to make sure your servers root partitions have more than 1GB of free space, you would add this to the definitions.json file:


```json
    {
        "name": "free_space",
        "command": "test $(df $PWD | awk '/[0-9]%%/{print $(NF-2)}') -gt %s",
        "args": ["kilobytes"]
    }
```

And then, to actually do the test on your servers, you create a check adding it in checks.json file like this:

```json
    "check_free_space": {
        "type": "free_space",
        "kilobytes": "1000000"
    }
```