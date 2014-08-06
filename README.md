# bcon
## A cli configuration manager for linux

If you're like me, you are too forgetful or too lazy to remember where all your configuration files are. bcon aims to help you find your config files quickly.

##Usage

Note that bcon does none of this yet. I'm just writing it down.

You have a config file. Add it to bcon

```
bcon add filename name tag1 tag2 tag3
```

then later

```
bcon name 
```

brings up your favourite editor (currently hardset to be nano. Who doesn't like nano?). 

The full implemented list of commands can be found with 

```
bcon help
```

Not sure what you need? You can grep the entries bcon has recorded

```
bcon list | grep search-term
```

##FUTURE:

bcon will have a config file where you can config things, invoked comedically with 

```
bcon myself
```

...or something.

bcon will track your changes and store diffs. commands will be added to allow you to easily revert configs to a previous point in time.


That's the gist of it. Let me know if it breaks everything, there's still lots of work to do.



