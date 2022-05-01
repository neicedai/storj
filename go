wget -qO- http://ipecho.net/plain >> ip
ipadd=`cat ip|awk -F\. '{print $1"."$2"."$3}'`
sed -i 's#tihuan#'$ipadd'#g' config.yaml
mkdir -p /root/.local/share/storj/storagenode
cp config.yaml /root/.local/share/storj/storagenode
apt update
apt install unzip -y
cd /usr/local/bin
wget https://github.com/storj/storj/releases/download/v1.53.1/storagenode_linux_amd64.zip
wget https://github.com/storj/storj/releases/download/v1.53.1/storagenode-updater_linux_amd64.zip
unzip storagenode_linux_amd64.zip
unzip storagenode-updater_linux_amd64.zip
service storagenode start
service storagenode-updater start
systemctl enable storagenode
systemctl enable storagenode-updater
