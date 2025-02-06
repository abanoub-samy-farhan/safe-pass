# safe-pass
a GO CLI Tool for handling password management using Redis database for storing the local encypted passwords. Not only password, you can safe anykind of secure keys and tokens for later use, all in a secure and safe mannar for your convenice :).

A pre-setup domain would include the following stuff
- Passwords (default category)
- Acess Tokens (Like Github or any others)
- API keys

# Table of Content

- [Installation](#installation)
- [Usage](#usage)
- [Commands](#commands)
- [Examples](#examples)
- [Contribution](#contribution)
- [License](#license)
- [TO-DO](#to-do)

## Installation

Easy installation just by running the following command:

```bash
$ go install github.com/abanoub-samy-farhan/safe-pass
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
- `show`: Show all data or data by category, domain and tag.
- `delete`: Delete a password, key or token from the database.
- `passgen`: Password Generator for making the user's life easier by simply hitting the generation, specifing some flags like:
    - `-l --length`: Determine the length of the password
    - `-s --special-characters`: Determine wether the password contains special characters or not
    - `-n --numbers`: Determine wether the password contains numbers or not


## Commands

### `add`
| Command | Description | Flags |
| --- | --- | --- |
| `safe-pass add [password\|key\|token] -c <category> -d <domain> -t <tag>` | Add a password, key or token to the database | `-c --category`, `-d --domain`, `-t --tag` |

### `show`
| Command | Description | Flags |
| --- | --- | --- |
| `safe-pass show -c <category> -d <domain> -t <tag>` | Show all data or data by category, domain and tag | `-c --category`, `-d --domain`, `-t --tag` |

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

## Examples

```bash
# Generating some passwords
$ safe-pass passgen -l 12 -s true -n true

Password: zN7^Tr%H4Vjy
Password is copied to your clipboard

$ safe-pass add Newpassword2910 -c password -d example.com -t work

Your Data is saved successfully!
Run `safe-pass show -c password  -d example.com -t work`to view it

$ safe-pass show -c password  -d example.com -t work

Category: password
	Domain: example.com	Tag: work: Newpassword2910
Time elapsed:  114.465µs
```

## Contribution

Contributions are welcome! If you'd like to contribute to this project just contact me and we can collaborate on the project.

## TO-DO

- [x] Encrypt passwords using `crypto/aes` lib
- [-] Master Password: Making a local authentication for the user to ensure security of the password saved from frauds
- [ ] Setup: Making a user-friendly setup, ensuring user have the required dependancies (`redis` for data storage), and setting the Master Password for accessing the data securly.
- [x] Add data: Adding password or key to the database, including the domain and tag to remeber the usecase of data stored.
- [x] Retrive data: returning the password of the requested domain.
- [-] Export Data: the user could export all data found in the database in an easy formatted form (such as `JSON` or `TXT`)
- [x] Cross-platform compatibility: Making the tool opperating for both `windows` and `linux`. Starting with linux tho.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
