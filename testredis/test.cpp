#include <iostream>
#include <queue>
#include <thread>
#include <assert.h>
#include "hirediscommand.h"

using namespace RedisCluster;
using std::string;
using std::cout;
using std::cerr;
using std::endl;

Cluster<redisContext>::ptr_t cluster_p;

#define BASE_MGMT_NO = 1000000000
#define BASE_EXTRNID = 10000000
const int MAX_USER_CNT = 100000

void processClusterCommand() {
	Cluster<redisContext>::ptr_t cluster_p;
	cluster_p = HiredisCommand<>::createCluster("localhost", 7000);
	auto reply = HiredisCommand<>::AltCommand(cluster_p, "FOO", "SET %s %s", "FOO", "BAR1");

	if(reply->type == REDIS_REPLY_STATUS || reply->type == REDIS_REPLY_ERROR) {
		cout << " Reply to SET FOO BAT " << endl;
		cout << reply->str << endl;
	}

	delete cluster_p;
}

int connectRedisCluster() {
	cluster_p = HiredisCommand<>::createCluster("localhost", 7000);

	if(cluster_p == NULL) {
		cout << "Fail to create cluster connection..." << endl;
		exit(-1);
	}

	return 0;
}

void getSvcMgmtNo() {

}

int main(int argc, const char *argv[]) {
	connectRedisCluster();

	try {
		processClusterCommand();
	} catch( const RedisCluster::ClusterException &e) {
		cout << "Cluster Exception: " << e.what() << endl;
	}

	return 0;
}
