apt update
apt install unzip -y
cd /usr/local/bin
wget https://github.com/storj/storj/releases/download/v1.53.1/storagenode_linux_amd64.zip
wget https://github.com/storj/storj/releases/download/v1.53.1/storagenode-updater_linux_amd64.zip
unzip storagenode_linux_amd64.zip
unzip storagenode-updater_linux_amd64.zip
