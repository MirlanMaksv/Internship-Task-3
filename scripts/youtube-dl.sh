#!/bin/sh
youtube-dl -f 'bestaudio[filesize<50M]' $1 -o $2/$3'%(id)s.%(ext)s'