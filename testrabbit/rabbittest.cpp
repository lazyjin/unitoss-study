#include <stdio.h>

#include <amqp_tcp_socket.h>
#include <amqp.h>
#include <amqp_framing.h>

int main(int argc, char *argv[]) {

	amqp_socket_t *socket = NULL;
	amqp_connection_state_t conn;

	amqp_bytes_t queuename;

	printf("hello world");
	return 0;
}