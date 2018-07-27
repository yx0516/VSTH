#include <iostream>
#include <string.h>
#include <openbabel/obconversion.h>

using namespace OpenBabel;


extern "C" { // 按 C 规则处理

	struct state
	{
		int 		status;
		const char* msg;
		const char* data;
	};


	void ConvertMol(char* inMol, struct state* ret)
	{
		ret->status = 0;

		std::string strIn(inMol);
        std::string strOut;

		// 定义必须放在 std::string strOut 之后，否则 ret->data 无法返回数据，详见 cgo 测试BUG【g++问题】
		std::stringstream sin , sout;

        sin << strIn;

        OBConversion Conv(&sin, &sout);
        Conv.AddOption("gen3D", OBConversion::GENOPTIONS);

        if (!Conv.SetInAndOutFormats("mol", "mol")) {
            ret->status = -1;
			ret->msg = "convert set file format error.";
            return;
        }

        int count = Conv.Convert(); // 非线程安全
		if (count == 1) {
			strOut = sout.str(); // 如果这里定义 std::string strOut，ret->data 无法带回数据
            ret->data = strOut.data();
		} else {
			ret->status = -2;
			ret->msg = "convert count error.";
		}

		return;
	}
}
