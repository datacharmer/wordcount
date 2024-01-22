## Wordcount: a simple example project for testscript documentation

The `wordcount` project is a command line program that accepts the same flags as the classic `wc` tool.

```
./wordcount -h

  -c	shows number of bytes
  -l	shows number of lines
  -w	shows number of words
  -m	shows number of characters
```

Additionally, for the purpose of this project, it implements several more options, that help with the specific testing.

```
  -log-file string    writes log file
  -o                  shows number of lowercase characters
  -s                  shows number of spaces
  -u                  shows number of uppercase characters
  -version            shows version
```

The aim of this project is to illustrate a 30-minute talk on `testscript`. As such, it does not exhaust all `testscript` feature, but covers most of the basic ones, while introducing a few advanced features that are neglected in the official docs.

Note: There is a serious bug in this `wordcount`: you should not rely on it on any serious business. If you find the bug, please file an issue. Extra points if you provide a (`testscript`-based) test to prove it.

