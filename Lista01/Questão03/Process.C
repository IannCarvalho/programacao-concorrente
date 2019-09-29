#include <stdio.h>
#include <unistd.h>
#include <stdlib.h>
#include <sys/wait.h>

int main()
{
    for (int i = 0; i < 3; i++)
    {
        if (fork() == 0)
        {
            printf("Sou um processo de pid %d, Meu pai é um processo de pid %d\n", getpid(), getppid());
            exit(0);
        }
    }

    for (int i = 0; i < 3; i++)
        wait(NULL);
}