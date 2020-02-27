## mkdeb
deb format metadata generator.

It generates a folder which can be used by `dpkg-deb` to generate a `deb` package.

Example:
- generate metadata dir: `mkdeb my_awesome_package-3.4.5`
- generate deb package from metadata dir: `dpkg-deb -b my_awesome_package-3.4.5`

### Install
- go to [releases](https://github.com/aburdulescu/mkdeb/releases) and download the latest package for your CPU architecture;
- install the package: `sudo dpkg -i path/to/downloaded/package`
### Usage
1. Write a `mkdeb.yaml` file following the deb format and the format exemplified in the `mkdeb.yaml` from this repo.
2. Run `mkdeb -h` for help message.
