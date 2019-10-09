#define _DEFAULT_SOURCE
#define N 1000

#include <stdio.h>
#include <stdlib.h>
#include <pthread.h>
#include <unistd.h>
#include <sys/time.h>

void *run(void *args)
{
    pthread_exit(0);
}

int main(int argc, char *argv[])
{
    pthread_t pthreads[N];
    FILE *file;
    file = fopen("logThread.csv", "w");
    struct timeval start, end;

    for (int i = 0; i < N; i++)
    {
        int *arg = (int *) malloc(sizeof(int *));
        *arg = i + 1;

        gettimeofday(&start, NULL);
        
        pthread_create(&pthreads[i], NULL, &run, arg);
        pthread_join(pthreads[i], NULL);
        
        gettimeofday(&end, NULL);
        long seconds = (end.tv_sec - start.tv_sec);
        long micros = ((seconds * 1000000) + end.tv_usec) - (start.tv_usec);
        fprintf(file,"%d,%d,%d\n", i+1,seconds,micros);
    }
}
