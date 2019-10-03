#include <stdio.h>
#include <unistd.h>
#include <stdlib.h>
#include <sys/wait.h>
#include <sys/time.h>

#define _DEFAULT_SOURCE
#define N 100

int main() {
    struct timeval start, end;
    timeval starts[N] = {};
    timeval ends[N] = {};
    int pids[N] = {};

    for (int i = 0; i < N; i++) {
        gettimeofday(&start, NULL);
        int res = fork();
        if (res == 0) {
            gettimeofday(&end, NULL);
            printf("%d\n", end.tv_usec);
            ends[i] = end;
            exit(0);
        } else {
            starts[i] = start;
            pids[i] = res;
        }
    }

    while (wait(NULL) != -1);

    FILE *file;
    file = fopen("logProcess.csv", "w");
    fprintf(file, "PID,start,end\n");

    for (int i = 0; i < N; i++) {
        long start_micro = starts[i].tv_usec;
        long end_micro = ends[i].tv_usec;
        fprintf(file, "%d,%ld,%ld\n", pids[i], start_micro, end_micro);
    }

    fclose(file);
}