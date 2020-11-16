# ORGANIZE
This is a small file organization tool I made ( remade tbh ) in golang. It takes all the files which are present in the folder its called in and sorts them into different folders based on their extension.

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
    I recommend that you put the executable somewhere in your path , so that then you can call it anywhere you want through the command line. Heck maybe even set up a cronjob to organize folders every day or so.
    
