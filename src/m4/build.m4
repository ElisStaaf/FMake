define(
`_gcc_build',
`gcc -o $2 $1'
)
define(
`_go_build',
`go build -o $2 $1'
)
define(
`_rust_build',
`rustc -o $2 $1'
)
define(
`_gpp_build',
`g++ -o $1 $1'
)
