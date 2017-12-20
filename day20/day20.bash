#!/bin/bash

# ran this to find the particles with the lowest acceleration, which would
# be the candidate(s) to be closest to the origin in the long term

# there were two particles that tied for the lowest acceleration in my puzzle
# input, so I inspected them and determined which one was going to be the
# closest based on their initial velocities.

# Then I subtracted one from that particle's line number, because the particles
# start at 0
awk -Fa '{print "P"FNR","$2}' input.txt | \
  sed 's#=<##g;s#-##g;s#>##g' | \
  awk -F, '{ACC=$2+$3+$4; print $1","ACC}' | \
  sort -t, -n -k 2 | \
  head
