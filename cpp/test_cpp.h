
#ifndef _TEST_CPP_H_
#define _TEST_CPP_H_

#include <stdint.h>
#include <string>

/// 回调类
class ICallback {
public:
	virtual void notify(const std::string& s) = 0;
};

/// 测试类
class TestCall {
public:
	static TestCall* Create() { return new TestCall(); }

	void SetCallback(ICallback* callback) { callback_ = callback; }

	int32_t Test(const std::string& s);

	uint32_t add(uint32_t  a ,uint32_t b);
private:
	TestCall() : callback_(NULL) {}

	ICallback* callback_;
};
#endif
