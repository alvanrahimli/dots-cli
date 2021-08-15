#!/bin/bash

echo "Downloading dots-cli"
if [ -f dots-cli-linux ]; then
  echo "Existing dots-cli-linux file found. Deleting it."
  rm -f dots-cli-linux
fi

wget https://github.com/alvanrahimli/dots-cli/releases/download/v0.1/dots-cli-linux
echo "Giving executable permission to app"
chmod +x dots-cli-linux
echo "Moving file to /usr/bin/"
sudo mv dots-cli-linux /usr/bin/
echo "Adding dots as alias to dots-cli-linux"
echo 'alias dots="dots-cli-linux"' >> "$HOME"/.bashrc
source "$HOME"/.bashrc

echo "Creating config file"
mkdir "$HOME"/.config/dots-cli/
curl https://raw.githubusercontent.com/alvanrahimli/dots-cli/master/config.json > "$HOME"/.config/dots-cli/config.json

echo "DONE! Your are ready to go!"
echo "Run 'dots init myfirstpack' to initialize your first package"