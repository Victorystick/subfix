# subfix [![Build Status](https://travis-ci.org/Victorystick/subfix.svg?branch=master)](https://travis-ci.org/Victorystick/subfix)

SubFix is a small program for manipulating subtitle files.

Currently it can only handle `.srt` and `.sub` formats, though others may be supported later on.


### Validation ###
You can validate a subtitle file by providing a filename as `subfix`'s only argument.
```sh
subfix movie.srt
```

### Timing ###
SubFix can adjust the timing of a subtitle file easily. If you want to delay the subtitles 5.4 seconds, set the time shift (ts) flag:
```sh
subfix --ts 5.4s movie.srt
```
If you instead want to hasten the subtitles just add a minus, like `-5.4s`.


### Format ###
To convert one format to another, say `.sub` to `.srt`, it's as simple as:
```sh
subfix --outfile movie.srt movie.sub
```
