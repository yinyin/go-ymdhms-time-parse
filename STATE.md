# Syntax

* `CurrentState`: Input character meaning; Digit/Character, Digit/Character->`NextState`, ...
* `NextState`: ...

# Transition State

*  `0`: Y; D
*  `1`: Y; D, C->`5`
*  `2`: Y; D
*  `3`: Y; D
*  `4`: -; C, D->`6`
*  `5`: M; D
*  `6`: M; D, C->`8`
*  `7`: -; C, D->`9`
*  `8`: D; D
*  `9`: D; D, C->`11`
* `10`: .; C, D->`12`
* `11`: h; D
* `12`: h; D, C->`14`
* `13`: :; C, D->`15`
* `14`: m; D
* `15`: m; D, C->`17`
* `16`: :; C, D->`18`
* `17`: s; D
* `18`: s; D
* `19`: END
