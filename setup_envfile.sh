KEY=$(head -c 32 /dev/urandom | od -t x1 -N 32 | head -n 2 | cut -d' ' -f2- | tr -d ' ' | paste -s -d '') # Your Secret Key
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