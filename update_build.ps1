echo " Hello deploying locally to 4 servers"

make
$src = "C:\Users\david\Documents\projects\GO Projects\kura_coin\bin\dau.exe"
$dst = "C:\Users\david\Documents\projects\GO Projects\kura_coin\test_servers\"

Copy-Item -Path $src -Destination $dst"server1"
Copy-Item -Path $src -Destination $dst"server2"
Copy-Item -Path $src -Destination $dst"server3"
Copy-Item -Path $src -Destination $dst"server4"

$mm = "server1\dau.exe"
# cd $dst"server1"
start cmd.exe /k "python "C:\Program Files\HelloWorld.py"