# codename

Generate codenames

```bash
$ go build .

$ ./codename
mintwind

$ ./codename
butterfield

$ ./codename
violetwinter

$ ./codename
gingerbrook

$ ./codename -c 1
valley

$ ./codename -c 3
berrymaplesummer

$ ./codename -c 3 --titlecase
GoldenSwiftBlossom

$ ./codename -c 3 --snakecase
glimmerFrostBreeze

$ ./codename -c 3 --uppercase
BRIDGEBUTTERMISTY

$ ./codename --help

codename

  Subcommands:
    completion   Generate shell completion script for bash or zsh.

  Flags:
        --version     Displays the program version string.
    -h  --help        Displays help with available flag, subcommand, and positional value parameters.
    -c  --count       Word count (default: 2)
    -l  --lowercase   Lowercase output
    -s  --snakecase   Snakecase output
    -t  --titlecase   Titlecase output
    -u  --uppercase   Uppercase output

$ source <(./codename completion bash)

# double-tab at the end of the line below brings up auto-complete
$ ./codename -
--count      --lowercase  --snakecase  --titlecase  --uppercase
-c           -l           -s           -t           -u   
```

## License

This project is licensed under the [MIT License](LICENSE).
