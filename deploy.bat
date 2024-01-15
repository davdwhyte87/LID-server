echo " Hello deploying locally to 4 servers"

make
set @src=C:\Users\david\Documents\projects\GO Projects\kura_coin\bin\dau.exe
set @dst=C:\Users\david\Documents\projects\GO Projects\kura_coin\test_servers\

echo %@dst%
xcopy /y "%@src%" "%@dst%server1"
xcopy /y "%@src%" "%@dst%server2"
xcopy /y "%@src%" "%@dst%server3"
xcopy /y "%@src%" "%@dst%server4"

start cmd.exe /k "title server1 & cd %@dst%server1 & dau.exe"
start cmd.exe /k "title server2 & cd %@dst%server2 & dau.exe"
start cmd.exe /k "title server3 & cd %@dst%server3 & dau.exe"
start cmd.exe /k "title server4 & cd %@dst%server4 & dau.exe"