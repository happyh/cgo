
#include <iostream>
#include "test_cpp.h"

int32_t TestCall::Test(const std::string & s) {
	std::cout << "TestCall::Test(" << s << ")" << std::endl;
	if (callback_) {
		callback_->notify(s);
	}
	return 0;
}

uint32_t TestCall::add(uint32_t  a ,uint32_t b){
	return a + b;
}
