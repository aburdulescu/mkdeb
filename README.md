## mkdeb
deb format metadata generator

### Build & Install
#### From prebuilt deb package
- go to [releases](https://github.com/aburdulescu/mkdeb/releases) and download the latest package for your CPU architecture;
- install the package: `sudo dpkg -i path/to/downloaded/package`
#### From source
##### Prerequisites:
- install golang;
##### Build:
- clone repo: `git clone https://github.com/aburdulescu/mkdeb.git`
- go to repo dir: `cd mkdeb/`
- build deb package: `make deb`
- install generated deb package: `sudo dpkg -i name_of_deb_package`
### Usage
1. Write a `mkdeb.json` file following the deb format and the format exemplified in the `mkdeb.json` from this repo.
2. Run `mkdeb -h` for help message.
