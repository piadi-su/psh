#!/usr/bin/env bash

set -e

echo "1) install"
echo "2) uninstall"

read -p "-> " CHOICE

case "$CHOICE" in
    1)
        echo "Installing..."
		go build -o psh ./src
		sudo install -m 755 psh /usr/local/bin/psh
		echo "Done"


        ;;
    2)
        echo "Uninstalling..."
		sudo rm -f /usr/local/bin/psh

		read -p "do you want to delite ~/.config/psh (y/n): " CHOICE2
		
		if [ "$CHOICE2" = "y" ]; then
			rm -rf ~/.config/psh
		fi

		echo "Done"

        ;;
    *)
        echo "Invalid choice"
        ;;
esac
