#define _DEFAULT_SOURCE
#define N 1000

#include <stdio.h>
#include <stdlib.h>
#include <pthread.h>
#include <unistd.h>

void *run(void *args)
{
    pthread_exit(0);
}

int main()
{
    for (int i = 0; i < N; i++)
    {
        pthread_create(&pthreads[i], NULL, &run, arg);
        pthread_join(pthreads[i], NULL);
    }
    sleep(200);
}
