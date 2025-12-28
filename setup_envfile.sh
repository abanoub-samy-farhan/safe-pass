set -euo pipefail


if [[ $EUID -ne 0 ]]; then
    echo "This script should be run with root user"
    exit 1
fi

if [ -z "${SUDO_USER:-}" ]; then
    echo "This script must be run via sudo, not directly as root"
    exit 1
fi

USER_HOME=$(getent passwd "$SUDO_USER" | cut -d: -f6)
HOMEDIR="$USER_HOME/.config/safe-pass"
BACKUP=$HOMEDIR/backups

# 32 character random key
KEY=$(od -An -N16 -tx1 /dev/urandom | tr -d ' \n')

mkdir -p "$BACKUP"

ENVFILE="$HOMEDIR/.env"
[ -f "$ENVFILE" ] && { echo ".env already exists"; exit 1; }

{
  echo "KEY=$KEY"
  echo "BACKUP=$BACKUP"
} > "$ENVFILE"

chmod 400 "$ENVFILE"
chown -R "$SUDO_USER:$SUDO_USER" "$HOMEDIR"