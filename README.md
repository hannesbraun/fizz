# Fizz

This is a small program to encrypt and decrypt any type of files using XOR.
**Do not consider this to be cryptographically safe**, since this is just a small fun project.

## Building from source

Compile with:
```sh
make
```

Install with:
```sh
sudo make install
```

Uninstall with:
```sh
sudo make uninstall
```

## Usage

To generate a key, run

```sh
fizz key <keylen> key.fizzkey
```
Replace `<keylen>` with the desired length of your key in bytes. In this example, `key.fizzkey` is the path to the key that should be generated.

To encrypt or decrypt a file, run
```sh
fizz xor file key.fizzkey
```
The resulting file will be placed in the same directory as the original file.

If the key is shorter than the file to decrypt or encrypt, it will be repeated.

## License

This software is released under the GNU General Public License, version 3. For more information see [LICENSE](LICENSE "GNU General Public License, Version 3").

## Author

Hannes Braun (<hannesbraun@mail.de>)
