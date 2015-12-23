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
```

## Adding check definitions

Open definitions.json file and add this for example:


```
    {
        "name": "process_running",
        "command": "ps -A | grep %s",
        "args": ["process"]
    }
```