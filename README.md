# safe-pass
a go cli Tool for handling password management using Redis database for storing the local encrypted passwords. Not only password, you can safe any kind of secret keys and tokens for later use, all in a secure and safe mannar for your convenice :).

NOTE: This is still a beta version of the cli, if you faced any problems please reported through and issue or if you wish to contribute that's completely welcomed, see [Contribution](#contribution).

# Table of Content

- [Installation](#installation)
- [Usage](#usage)
- [Commands](#commands)
- [Examples](#examples)
- [Contribution](#contribution)
- [License](#license)
- [TO-DO](#to-do)

## Installation

First, you have to have go installed and added to your PATH so that you can work easily this go project.
Easy installation just by running the following command:

```bash
$ go install github.com/abanoub-samy-farhan/safe-pass
```
or 
```bash
$ git clone https://github.com/abanoub-samy-farhan/safe-pass
$ cd safe-pass
$ go build
$ go install
```

Now before using the tool, there are some environment variables that has to be settled up. all are found in the `setup_envfile.sh` and `setup_redis.sh` (for ensuring you have redis running on your machine).

In the `setup_envfile.sh`
```bash
KEY= '' # Place here a key of 32 characters
BACKUP="$HOME/.config/safe-pass" # Don't change this


mkdir -p $BACKUP
touch $BACKUP/.env

echo "KEY=$KEY" > $BACKUP/.env
echo "BACKUP=$BACKUP" >> $BACKUP/.env
```
You will have to replace the `KEY` variable with your own 32 character key before running the script.

Before using the tool, the master key has to be setup that can only be done using the root user by running 

```bash
sudo $GOPATH/bin/safe-pass init
```
Replace `GOPATH` with your actual Go workspace path.

Or

you can make your program in an accessible place for the root user (for example in the `usr/local/bin`) by running this
```bash
sudo mv $GOPATH/bin/safe-pass /usr/local/bin/
sudo safe-pass init
```


Also, for autocompleting feature, you have to run this command for setting up your envorinement according to your system:

- Linux:
```bash
$ safe-pass completion bash > /etc/bash_completion.d/safe-pass
```

- MacOS:
```bash
$ safe-pass completion bash > $(brew --prefix)/etc/bash_completion.d/safe-pass
```

- Windows:
```cmd
> safe-pass completion powershell > safe-pass.ps1
```

## Usage

The safe-pass is a command line utility for managing sensitive data such as passwords, keys and tokens. It uses a redis database for storing the data, and provides a simple and secure way to interact with the data.

The tool is designed to be used from the command line, and provides a number of commands to add, list, retrieve and delete data from the database.


The available commands are:

- `add`: Add a password, key or token to the database.
- `show`: copies the selected entry to the clipboard using interactive cli tools.
- `delete`: Delete a password, key or token from the database.
- `passgen`: Password Generator for making the user's life easier by simply hitting the generation, specifing some flags like:
    - `-l --length`: Determine the length of the password
    - `-s --special-characters`: Determine wether the password contains special characters or not
    - `-n --numbers`: Determine wether the password contains numbers or not


## Commands

### `add`
| Command | Description | Flags |
| --- | --- | --- |
| `safe-pass add [password\|key\|token] -c <category> -d <domain> -t <tag>` | Add a password, key or token to the database, NOTE: Don't include hyphins in category, domain or tag | `-c --category`, `-d --domain`, `-t --tag` |

> WARNING: names of categories, domains and tags should not contain the characters `-` or `:`

### `show`
| Command | Description |
| --- | --- |
| `safe-pass show` | Copies a single entry to the clipboard |

### `delete`
| Command | Description | Flags |
| --- | --- | --- |
| `safe-pass delete -c <category> -d <domain> -t <tag>` | Delete a password, key or token from the database | `-c --category`, `-d --domain`, `-t --tag` |

### `edit`
| Command | Description | Flags |
| --- | --- | --- |
| `safe-pass edit -c <category> -d <domain> -t <tag>` | Edit a password, key or token from the database | `-c --category`, `-d --domain`, `-t --tag` |


### `passgen`
| Command | Description | Flags |
| --- | --- | --- |
| `safe-pass passgen -l <length> -s <special-characters> -n <numbers>` | Password Generator for making the user's life easier by simply hitting the generation, specifing some flags like: `-l --length`, `-s --special-characters`, `-n --numbers` | `-l --length`, `-s --special-characters`, `-n --numbers` |

### `backup`
| Command | Description |
| --- | --- |
| `safe-pass backup` | Backup the entire database to a file |

### `restore`
| Command | Description |
| --- | --- |
| `safe-pass restore` | Restore the database from a backup file |

## Examples

```bash
# Generating some passwords
$ safe-pass passgen -l 12 -s true -n true

Password: zN7^Tr%H4Vjy
Password is copied to your clipboard

$ safe-pass add Newpassword2910 -c password -d example.com -t work

Your Data is saved successfully!
Run `safe-pass show -c password  -d example.com -t work`to view it

$ safe-pass show 
Use the arrow keys to navigate: ↓ ↑ → ← 
? Select a category to show: : 
  ▸ passwords
    tokens

$ safe-pass backup
Backup is created at: /home/abanoub-aziz/.config/safe-pass/safe-pass-2025-08-21:18:11:01.bin

$ safe-pass restore
Search: █
? Select a backup file to restore: : 
  ▸ .env
    safe-pass-2025-08-21:18:11:01.bin.gz

```

## Contribution

Contributions are welcome! If you'd like to contribute to this project just contact me and we can collaborate on the project. [ContactMe](mailto:abanoubsamy2341@gmail.com)

## TO-DO

- [x] Encrypt passwords using `crypto/aes` lib
- [x] Master Password: Making a local authentication for the user to ensure security of the password saved from frauds
- [x] Setup: Making a user-friendly setup, ensuring user have the required dependancies (`redis` for data storage), and setting the Master Password for accessing the data securly.
- [x] Add data: Adding password or key to the database, including the domain and tag to remeber the usecase of data stored.
- [x] Retrive data: returning the password of the requested domain.
- [x] Export Data: the user could export all data found in the database in an easy formatted form (such as `JSON` or `TXT`)
- [x] Cross-platform compatibility: Making the tool opperating for both `windows` and `linux`. Starting with linux tho.
- [x] Restoring Data from snapshots exported before

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
