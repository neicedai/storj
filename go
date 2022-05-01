wget -qO- http://ipecho.net/plain >> ip
ipadd=`cat ip`
sed -i 's#tihuan#'$ipadd'#g' config.yaml
cp storagenode-updater.service /etc/systemd/system
cp storagenode.service /etc/systemd/system
cp update /usr/local/bin
apt update
apt install unzip -y
cd /usr/local/bin
wget https://github.com/storj/storj/releases/download/v1.53.1/storagenode_linux_amd64.zip
wget https://github.com/storj/storj/releases/download/v1.53.1/storagenode-updater_linux_amd64.zip
unzip storagenode_linux_amd64.zip
unzip storagenode-updater_linux_amd64.zip
./storagenode setup
cd /root/storj
cp config.yaml /root/.local/share/storj/storagenode
service storagenode start
service storagenode-updater start
systemctl enable storagenode
systemctl enable storagenode-updater
