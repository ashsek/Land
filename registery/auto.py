import os
import sys

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

#cat tailed.txt | grep -i args | awk '{print $4}' | sed 's/today/''/g'| sed "s/'/ /g" | sed 's/,/ /'|sed "s/],/ /g"
#sed -i '.bak' '94s/.*/        args: ["R_Hash2", "08", "replaced",today],/' test2.js
#sys.stdout.write(output_str)  # same as print
#sys.stdout.flush()

sys.stdout.write("Please Input the Hash: ")
y = raw_input()
os.system("sed -i '.bak' '94s/.*/        args: [\"R_Hash0\",\""+y+"\",today],/' invoke.js")
os.system("node invoke.js")
os.system("node query.js")