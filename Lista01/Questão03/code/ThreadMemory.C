#define _DEFAULT_SOURCE
#define N 1000

#include <stdio.h>
#include <stdlib.h>
#include <pthread.h>
#include <unistd.h>

void *run(void *args) {
    sleep(20);
    pthread_exit(NULL);
}

int main() {
    pthread_t pthreads[N];

    for (int i = 0; i < N; i++) {
        int *arg = (int *) malloc(sizeof(int *));
        *arg = i + 1;

        pthread_create(&pthreads[i], NULL, &run, arg);
    }

    for (int i = 0; i < N; i++)
        pthread_join(pthreads[i], NULL);
}
