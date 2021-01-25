package sheet

const (
	SkipPrefix    = "__"
	ListSeparator = ","

	ProtoFileName = "conf.proto"

	GenDataFileName = "gendata"
)

const (
	KRowStart     = 1
	KColStart     = 1
	kNameRowIndex = 1
	kTypeRowIndex = 2
	KDataRowStart = 3

	kMaxLenType = "repeated uint64 "

	kMessageFmt = `message %s
{
	message Record
	{
%s
	}

	map<%s, Record> records = 1;
	uint32 crc32 = 2;
}

`
)

// template
const ()
