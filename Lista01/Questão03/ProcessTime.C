#define _DEFAULT_SOURCE
#define N 1000

#include <stdio.h>
#include <unistd.h>
#include <stdlib.h>
#include <sys/wait.h>
#include <sys/time.h>

int main() {
    struct timeval start, end;
    timeval starts[N] = {};
    int pids[N] = {};

    FILE *endFile;
    endFile = fopen("end.csv", "w");

    for (int i = 0; i < N; i++) {
        gettimeofday(&start, NULL);
        int res = fork();
        if (res == 0) {
            gettimeofday(&end, NULL);
            fprintf(endFile,"%d,%d,%ld\n", getpid(),i+1,end.tv_usec);
            exit(0);
        } else {
            starts[i] = start;
            pids[i] = res;
        }
    }

    while (wait(NULL) != -1);
    fclose(endFile);

    FILE *startFile;
    startFile = fopen("start.csv", "w");

    for (int i = 0; i < N; i++) {
        fprintf(startFile, "%d,%d,%ld\n", pids[i],i+1,starts[i].tv_usec);
    }

    fclose(startFile);
}