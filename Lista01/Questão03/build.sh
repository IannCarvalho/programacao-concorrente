mkdir bin
echo ">>>>>> COMPILING CLASSES"
gcc ./code/ProcessMemory.C -o ./bin/ProcessMemory.out
gcc ./code/ProcessTime.C -o ./bin/ProcessTime.out
gcc ./code/ThreadMemory.C -lpthread -o ./bin/ThreadMemory.out 
gcc ./code/ThreadTime.C -lpthread -o ./bin/ThreadTime.out

echo ""
echo ">>>>>> RUNNING"
./bin/ProcessMemory.out
./bin/ProcessTime.out
./bin/ThreadMemory.out 
./bin/ThreadTime.out