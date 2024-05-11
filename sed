sed -i 's/storage.allocated-disk-space: [0-9.]* TB/storage.allocated-disk-space: 5.2 TB/' /root/.local/share/storj/storagenode/config.yaml
service storagenode restart
