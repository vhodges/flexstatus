flexstatus is a status bar generator to use with programs like Dzen2 or lemonbar.

So why yet another program to generate status updates? I originally started with yabar (which I really like) but it seems to pick random y locations and so is not usable for me. My intent is to replace my shellscript - https://github.com/vhodges/dotfiles/blob/master/utils/bin/panel.sh - since I hate that it can take upto a second to reflect the proper workspace name.

## Features

Here's what it implements now/features:

* Date/Time (ie a clock)
* Workspace name (updated in realtime) (or it can display a crude pager or n/n if you prefer)
* Static Text (for buttons)
* Periodically poll a file, with configurable sleep duration
* Periodically poll a command, with configurable sleep duration
* Multiple pollers, each with their own sleep duration
* Bar agnostic. It uses Go templates to support different bars

Coming Soon:

* Stream from a command in the background
* Stream updates from a file (ie a FIFO)
* Moar widgets!

## Building/Running

Building:

Uses Go so you need Go setup (including GOPATH):

```
  go get -u github.com/vhodges/flexstatus
```

Will leave the flexstatus binary in $GOPATH/bin

I'll look into building downloadable releases with binaries and/or packaging it up for the AUR

Running:

```
  ./flexstatus -config sample_widgets.toml
```

If the config file can't be found (or errors when loading it) it
adds a clock and a version string to the bar so you can at least
see something.

The config file defaults to $HOME/.flexstatus_widgets.toml if not
specified on the command line.

## Similar/Related

There are a fair amount of tools for this kind of thing (I'll add more as I find them):

* https://github.com/l3pp4rd/statusbar    (Go, only works with dzen2).
* https://github.com/ibabushkin/bartender (Rust, I like this as it solves problems I also am having, but alas I think it's too low level)
* https://github.com/damog/mastodon
* https://github.com/ghedamat/go-i3status
* https://github.com/davidscholberg/goblocks 
* https://github.com/denbeigh2000/goi3bar


## License - MIT
