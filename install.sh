#!/bin/bash

echo "Downloading dots-cli"
wget https://github.com/alvanrahimli/dots-cli/releases/download/v0.1/dots-cli-linux
echo "Giving executable permission to app"
chmod +x dots-cli-linux
echo "Moving file to /usr/bin/"
sudo mv dots-cli-linux /usr/bin/
echo "Adding dots as alias to dots-cli-linux"
echo 'alias dots="dots-cli-linux"' >> "$HOME"/.bashrc
source "$HOME"/.bashrc

echo "DONE! Your are ready to go!"
echo "Run 'dots init myfirstpack' to initialize your first package"