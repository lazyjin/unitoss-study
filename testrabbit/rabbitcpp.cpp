#include <ev.h>
#include <amqpcpp.h>
#include <amqpcpp/libev.h> 

int main(void) {
	auto *loop = EV_DEFAULT;

	AMQP::LibEvHandler handler(loop);
		
	return 0;
}
