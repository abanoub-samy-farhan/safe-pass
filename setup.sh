KEY=$(od -An -N16 -tx1 /dev/urandom | tr -d ' \n') # auto-generated random key
HOMEDIR="$HOME/.config/safe-pass"
BACKUP=$HOMEDIR/backups


if [[ $EUID -ne 0 ]]; then
    echo "This script should be run with root user"
    exit 1
fi

mkdir -p $BACKUP
touch $HOMEDIR/.env

echo "KEY=$KEY" > $HOMEDIR/.env
echo "BACKUP=$BACKUP" >> $HOMEDIR/.env

chmod 440 $HOMEDIR/.env

# Building the project and moving the binary to /usr/local/bin
go build -o /usr/local/bin/safe-pass

echo "Setup completed successfully. You can now run 'safe-pass' command."