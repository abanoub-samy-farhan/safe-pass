# safe-pass
a GO CLI Tool for handling password management using Redis database for storing the local encypted passwords. Not only password, you can safe anykind of secure keys and tokens for later use, all in a secure and safe mannar for your convenice :).

A pre-setup domain would include the following stuff
- Domains
- Acess Tokens (Like Github or any others)
- API keys

# TO-DO

- [ ] Encrypt passwords using `bycrypt` or `crypto` libs
- [ ] Master Password: Making a local authentication for the user to ensure security of the password saved from frauds
- [ ] Audit Logs: Adding a log for checking the last time a password is either accessed or modified.
- [ ] Setup: Making a user-friendly setup, ensuring user have the required dependancies (`redis` for data storage), and setting the Master Password for accessing the data securly.
- [ ] Add data: Adding password or key to the database, including the domain and tag to remeber the usecase of data stored.
- [ ] Retrive data: returning the password of the requested domain.
- [ ] Export Data: the user could export all data found in the database in an easy formatted form (such as `JSON` or `TXT`)
- [ ] Cross-platform compatibility: Making the tool opperating for both `windows` and `linux`. Starting with linux tho.


Hopefully this project would be helpful and finished by time.