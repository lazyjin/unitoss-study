#include <SimpleAmqpClient/SimpleAmqpClient.h>
#include <iostream>
#include <string>
#include <sstream>

using namespace AmqpClient;

int main() {
	Channel::ptr_t channel;

	channel = Channel::Create();
	channel->DeclareQueue("testqueue");
	channel->BindQueue("testqueue", "amq.direct", "test");

	BasicMessage::ptr_t msg_in = BasicMessage::Create();
	msg_in->Body("This is a test message");

	channel->BasicPublish("amq.direct", "test", msg_in);
	channel->BasicConsume("testqueue", "consumertag");

	BasicMessage::ptr_t msg_out = channel->BasicConsumeMessage("consumertag")->Message();

	std::cout << "Message text: " << msg_out->Body() << std::endl;

	return 0;
}
