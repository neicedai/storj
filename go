chmod 777 dash1
cp dash1 /usr/local/bin
chmod +x identity
sudo mv identity /usr/local/bin/identity
identity create storagenode
identity authorize storagenode $1
wget -qO- http://ipecho.net/plain >> ip
ipadd=`cat ip`
sed -i 's#tihuan#'$ipadd'#g' config.yaml
cp storagenode-updater.service /etc/systemd/system
cp storagenode.service /etc/systemd/system
chmod 777 update
cp update /usr/local/bin
apt update
apt install unzip -y
cd /usr/local/bin
wget https://github.com/storj/storj/releases/download/v1.90.2/storagenode_linux_amd64.zip
wget https://github.com/storj/storj/releases/download/v1.90.2/storagenode-updater_linux_amd64.zip
unzip storagenode_linux_amd64.zip
unzip storagenode-updater_linux_amd64.zip
./storagenode setup
cd /root/storj
cp config.yaml /root/.local/share/storj/storagenode
service storagenode start
service storagenode-updater start
systemctl enable storagenode
systemctl enable storagenode-updater
