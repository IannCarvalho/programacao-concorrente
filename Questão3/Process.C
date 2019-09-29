#include <stdio.h>
#include <unistd.h>
#include <stdlib.h>
#include <sys/wait.h>

int main()
{
    printf("sou o pai de todos %d", getpid());
    for (int i = 0; i < 3; i++)
    {
        if (fork() == 0)
        {
            printf("[son] pid %d from [parent] pid %d\n", getpid(), getppid());
            exit(0);
        }
        printf("[son] pid %d from [parent] pid %d\n", getpid(), getppid());
    }

    for (int i = 0; i < 3; i++)
        wait(NULL);
}