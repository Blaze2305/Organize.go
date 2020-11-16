# ORGANIZE
This is a small file organization tool I made ( remade tbh ) in golang.

### Requirements :
. Golang


### Run :
- Build the compiled source of the project
    ```bash
    $ make build
    ```
    - build takes an optional argument OS : target OS to be built to [ windows | linux | darwin ...]
        ```
        $ make -e OS=windows build
        ```
    Build outputs to a folder named bin/organize with the executable.
    