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
 % patch -p0 < present.patch
 ```

The above patch adds a progress indicator (slide number / total number
of slides) in the bottom right corner of every slide, and changes the
color of "code" text to be blue (instead of black) so they standout
better.

4. Run the `present` tool
 ```
 % present
 ```

The slides will be available at [http://127.0.0.1:3999/introduction-to-go.slide](http://127.0.0.1:3999/introduction-to-go.slide#1)

_Note_: The URL *must* contain `127.0.0.1:3999` because `localhost:3999` will not work!

# License and Materials

This presentation is licensed under the [Creative Commons Attribution-ShareAlike 4.0 International](https://creativecommons.org/licenses/by-sa/4.0/) licence.

You are encouraged to remix, transform, or build upon the material, providing you give appropriate credit and distribute your contributions under the same license.
