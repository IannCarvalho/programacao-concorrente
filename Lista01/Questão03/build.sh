mkdir bin
gcc -o ./bin/ProcessMemory.out ./code/ProcessMemory.C
gcc -o ./bin/ProcessTime.out ./code/ProcessTime.C
gcc -lpthread -o ./bin/ThreadMemory.out ./code/ThreadMemory.C
gcc -lpthread -o ./bin/ThreadTime.out ./code/ThreadTime.C
