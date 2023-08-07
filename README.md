# FFinder

Helper utility to provide fast searching.


### v0.0.2
Include options for multiple parameters.

```
$ time ./ffinder -path ../workspaces/git-mirrors/ -days 7 -poolsize 10

real    0m10.671s
user    0m0.276s
sys     0m0.833s

$ time ./ffinder -path ../workspaces/git-mirrors/ -days 7 -poolsize 50

real    0m3.568s
user    0m0.179s
sys     0m0.317s

```


### v0.0.1

Provide a simple means to find file modified in the last week.

Usage
```
ffinder <path>
```

