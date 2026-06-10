#include <stdio.h>
#include <unistd.h>
#include <sys/socket.h>

int main() {
    while (1) {
        // TODO: actually receive a request.
        printf("Waiting for a request...\n");
        sleep(1);
    }
}
