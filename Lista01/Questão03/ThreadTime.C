#define _DEFAULT_SOURCE
#define N 1000

#include <stdio.h>
#include <stdlib.h>
#include <pthread.h>
#include <unistd.h>
#include <sys/time.h>

void *run(void *args) {
    pthread_exit(0);
}

int main() {
    pthread_t pthreads[N];
    FILE *file;
    file = fopen("logThread.csv", "w");
    struct timeval start, end;

    for (int i = 0; i < N; i++) {
        int *arg = (int *) malloc(sizeof(int *));
        *arg = i + 1;

        gettimeofday(&start, NULL);
        
        pthread_create(&pthreads[i], NULL, &run, arg);
        
        gettimeofday(&end, NULL);
        fprintf(file,"%d,%ld\n", i+1,(end.tv_usec) - (start.tv_usec));
    }

    for (int i = 0; i < N; i++)
        pthread_join(pthreads[i], NULL);
}
