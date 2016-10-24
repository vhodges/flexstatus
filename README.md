# flexstatus - A Flexible Status Bar Tool

flexstatus is a status bar generator to use with programs like Dzen2 or lemonbar (or i3bar).  

## So why yet another program to generate status updates?

There is a fair amount of tools for this kind of thing:

* https://github.com/l3pp4rd/statusbar    (Go, only works with dzen2).
* https://github.com/ibabushkin/bartender (Rust, I like this as it solves problems I also am having, but alas I think it's too low level)
* https://github.com/ghedamat/go-i3status
* https://github.com/damog/mastodon
* https://github.com/davidscholberg/goblocks 
* https://github.com/denbeigh2000/goi3bar

I started with yabar (which I really like) but it seems to pick random y locations and so is not usable for me. My intent is to 
replace my shellscript https://github.com/vhodges/dotfiles/blob/master/utils/bin/panel.sh since it can take upto a second to 
reflect the proper workspace name.  I want to have lots of built in widgets. Each of the above has widgets that I like but none have
all of them in one place.  Lastly it will use Go templates so that it can stay agnostic as far as which bar to use for rendering.

## Widgets

My needs are fairly minimal (but I'd like to add more as I have time):

* Date/Time 
* CPU Temp
* CPU Usage
* RAM usage
* Diskspace
* Volume
* Workspace name
* WiFi
* Lock Screen (this is just a static clickable)

## License - MIT
