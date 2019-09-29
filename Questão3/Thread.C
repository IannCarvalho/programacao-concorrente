#include <stdio.h>
#include <pthread.h>
#include <unistd.h>

void *run(void *args)
{
    pthread_exit(0);
}

int main(int argc, char *argv[])
{
    int i;
    pthread_t pthreads[100];

    for (i = 0; i < 3; i++)
    {
        pthread_create(&pthreads[i], NULL, &run, (void *)i + 1);
        pthread_join(pthreads[i], NULL);
    }

    return 0;
}