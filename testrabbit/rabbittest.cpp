#include <stdio.h>

#include <amqp_tcp_socket.h>
#include <amqp.h>
#include <amqp_framing.h>

int main(int argc, char *argv[]) {
	char const *hostname = "localhost";
	int port = 5672;

	amqp_socket_t *socket = NULL;
	amqp_connection_state_t conn;

	amqp_bytes_t queuename;

	conn = amqp_new_connection();
	socket = amqp_tcp_socket_new(conn);
	if(!socket) {
		printf("fail to create socket\n");
		return -1;
	}

	int status = amqp_socket_open(socket, hostname, port);
	if(status) {
		printf("fail to open socket\n");
		return -1;
	}

	amqp_login(conn, "/", 0, 131072, 0, AMQP_SASL_METHOD_PLAIN, "guest", "guest");
	amqp_channel_open(conn, 1);

	amqp_basic_properties_t props;
	props._flags = AMQP_BASIC_CONTENT_TYPE_FLAG | AMQP_BASIC_DELIVERY_MODE_FLAG;
	props.content_type = amqp_cstring_bytes("text/plain");
	props.delivery_mode = 2; /* persistent delivery mode */
	int result = amqp_basic_publish(conn,
									1,
									amqp_cstring_bytes("amq.direct"),
									amqp_cstring_bytes("test"),
									0,
									0,
									&props,
									amqp_cstring_bytes("this is test message"));
	if(result != AMQP_STATUS_OK) {
		printf("fail to send message!!\n");
		return -1;
	}

	amqp_channel_close(conn, 1, AMQP_REPLY_SUCCESS);
	amqp_connection_close(conn, AMQP_REPLY_SUCCESS);
	amqp_destroy_connection(conn);

	printf("hello world!\n");

	return 0;
}