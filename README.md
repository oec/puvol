# Simple PulseAudio volume monitor

This little monitor is my attempt to read the current volume and mute setting
of the default sink from PulseAudio and make it accessible to xmobar:

```haskell
Config { commands = [ Run CommandReader "puvol-cont" "Vol" ]
       , template = " %Vol% "
```

If it is called as `puvol-cont` it will wait for updates from PulseAudio, read
the current value of the mute and volume states and write an output line to
stdout.

## Installation

```shell
go install kesim.org/puvol
```

For the continous mode you have to copy or add an symbolic link with name `puvol-cont`:

```shell
cd /path/to/bin/dir
ln -s puvol puvol-cont
```
