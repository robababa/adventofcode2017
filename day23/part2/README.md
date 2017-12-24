# Day 23 part 2 solution explained

My puzzle input is `original.input` in this directory

I examined the innermost loop from

`set g d`

to

`jnz g -8`

and simplified the instructions into `step1.input` in this directory.  To do this, I had to add a new instruction `mod X Y` where X is set to zero if it is non-zero and `Y mod X = 0`.  I also had to modify the `jnz` commands after this loop block, because the number of instructions changed.

I then repeated the process with the now-innermost loop from

`set e b`

to

`jnz g -13`

and simplified the instructions into `step2.input` in this directory.  This time, I had to add a new instruction `cmp X Y` where register X is set to zero if Y is composite.  Again, I had to modify the `jnz` command at the end because the number of instructions changed.

Then I ran my program with `step2.input`, and the answer came back in 0.3 seconds.
