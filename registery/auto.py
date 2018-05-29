import os

#clear docker containers,

os.system('docker rm -f $(docker ps -aq)')
os.system('docker network prune')
os.system('docker rmi dev-peer0.org1.example.com-registery-1.0-f1e65742f3b22688f7bf83215504431dbe6cd834182aefcf166f3eb79d2354d5')

#starting to run npm install

os.system('npm install')

#start fabric

os.system('./startFabric.sh')
os.system('node enrollAdmin.js')
os.system('node registerUser.js')
os.system('node query.js')