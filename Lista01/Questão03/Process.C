#define _DEFAULT_SOURCE

#include <stdio.h>
#include <unistd.h>
#include <stdlib.h>
#include <sys/wait.h>
#include <sys/time.h>

int main()
{
    FILE *file;
    file = fopen("logProcess.csv", "w");
    struct timeval start, end;

    for (int i = 0; i < 3; i++)
    {
        gettimeofday(&start, NULL);
        if (fork() == 0)
        {
            gettimeofday(&end, NULL);
        	long seconds = (end.tv_sec - start.tv_sec);
        	long micros = ((seconds * 1000000) + end.tv_usec) - (start.tv_usec);
            fprintf(file,"%d,%d,%d,%d\n", getpid(), getppid(),seconds,micros);
            exit(0);
        }
    }

    for (int i = 0; i < 3; i++)
        wait(NULL);

    fclose(file);
}