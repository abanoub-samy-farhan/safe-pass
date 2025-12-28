# fail early on errors
set -euo pipefail

if [[ $EUID -ne 0 ]]; then
    echo "This script should be run with root user"
    exit 1
fi

if [ -z "${SUDO_USER:-}" ]; then
    echo "This script must be run via sudo, not directly as root"
    exit 1
fi

# Enter your Go binary path (use `which go` to find it if necessary)
GO=

# Building the project and moving the binary to /usr/local/bin
$GO build -o /usr/local/bin/safe-pass

echo "Setup completed successfully. You can now run 'safe-pass' command."