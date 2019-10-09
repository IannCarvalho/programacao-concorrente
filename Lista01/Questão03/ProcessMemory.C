#define _DEFAULT_SOURCE
#define N 1000

#include <stdio.h>
#include <unistd.h>
#include <stdlib.h>
#include <sys/wait.h>

int main() {
    for (int i = 0; i < N; i++) {
        if (fork() == 0) {
            sleep(200);
        }
    }

    while (wait(NULL) != -1);
}