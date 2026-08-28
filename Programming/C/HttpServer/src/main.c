#include <stdio.h>
#include <unistd.h>
#include <sys/socket.h>

int main() {
    int fd = socket(AF_INET, SOCK_STREAM, 0);
    while (1) {
        // TODO: actually receive a request.
        printf("Waiting for a request...\n");
        sleep(1);
    }
}
