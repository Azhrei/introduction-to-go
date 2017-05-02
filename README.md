# Usage

This slide deck and supporting material is part of the _Introduction to Go_ training course by Frank Edwards (as modified from Dave Cheney's course).

# Installation

1. Install the Go `present` tool
 ```
 % go get -u -v golang.org/x/tools/cmd/present
 ```

2. Clone this code into a directory
 ```
 % git clone https://github.com/Azhrei/introduction-to-go
 % cd introduction-to-go
 ```

3. (Optional) Apply the patch to `present`
 ```
 % patch -p1 < present1.patch
 ```

The above patch adds a progress indicator (slide number / total number
of slides) in the bottom right corner of every slide, and changes the
color of "code" text to be blue (instead of black) so they standout
better.

A second patch is provided that modifies the tool so that slides are
automatically numbered, but this gets in the way for "continuation
slides", which should not be assigned a new number.  I plan to come up
with a way to correct for this (new syntax add to the Title tag?) but
until I do, this patch is NOT recommended (although the `present2.patch`
file is provided).

I really want the second patch to work as it would easily allow a table
of contents at the beginning of the slide deck built automatically by
jQuery with links to the correct slide.  The present tool already uses
jQuery and jQuery-UI, so adding a navbar somewhere on the slide is
certainly viable using techniques already included within the existing
web page.

I am considering a third patch that would add features to the source
code that appears in a `.code` or `.play` block.  I want:  line numbers,
multiple start/stop regexes (they should work like they do in `awk`),
and maybe syntax hilighting (the Go Playground has it!).

I would like to see the whole thing work from port 80 so that I can
run it from my web site (which the provider restricts from opening
arbitrary incoming ports).  This might be possible just using the
`-base` command line option, but I haven't checked.  Note that the
web socket interface needs to be able to connect to port 80 as well,
perhaps at a different URL, which means it needs to communicate
using HTTP so that it goes through a web server.  I haven't resewarched
any of this yet!

4. Run the `present` tool
 ```
 % present
 ```

The slides will be available at [http://127.0.0.1:3999/introduction-to-go.slide](http://127.0.0.1:3999/introduction-to-go.slide#1)

_Note_: The URL *must* contain `127.0.0.1:3999` because `localhost:3999` will not work!

# License and Materials

This presentation is licensed under the [Creative Commons Attribution-ShareAlike 4.0 International](https://creativecommons.org/licenses/by-sa/4.0/) licence.

You are encouraged to remix, transform, or build upon the material, providing you give appropriate credit and distribute your contributions under the same license.
