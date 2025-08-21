KEY='2*EUH$@^9t$yGk6gUr8nzcKsBzf%zHbZ' # Your Secret Key
BACKUP="$HOME/.config/safe-pass" # Don't change this


mkdir -p $BACKUP
touch $BACKUP/.env

echo "KEY=$KEY" > $BACKUP/.env
echo "BACKUP=$BACKUP" >> $BACKUP/.env