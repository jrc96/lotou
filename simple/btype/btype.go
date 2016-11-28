package btype

type PHead struct {
	//Len     int32
	Flag    int32
	Id      int16
	SType   int32
	OriLen  int32
	Type    int32
	ReqId   int32
	Session int32
	AcId    int32
}

type CCheckSession struct {
	AcName  string
	AcId    int32
	Session string
}

type SCheckSession struct {
	Session  int32
	AcId     int32
	Coin     int64
	YuanBao  int32
	JQ       int64
	NickName string
	Hasdc    bool
	Tili     int32
	MaxTili  int32
	Quebi    int64
	Jifen    int32
	Roomid   int32
	Beishu   int32
	DayBuzu  int32
	UserType int32
	Sex      int32
	HeadURL  string
}
