import os
import sys

sys.stdout.write("Please Input the Hash: ")
y = raw_input()
os.system("sed -i '.bak' '94s/.*/        args: [\"R_Hash2\", \"08\", \""+y+"\",today],/' test2.js")