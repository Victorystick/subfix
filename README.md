subfix
======

SubFix is a small program for manipulating subtitle files.

Currently it can only handle `.srt` files, but the idea is to support at least `.sub` as well.


### Timing ###
SubFix can adjust the timing of a subtitle file easily. If you want to delay the subtitles 5.4 seconds, do:
`subfix movie.srt 5.4s`. If you instead want to hasten the subtitles just add a minus, like `-5.4s`.
