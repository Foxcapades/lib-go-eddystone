package eddystone

import bean "github.com/Foxcapades/Beanies/v2"

type mockFrame struct {
	MString      bean.StringGetter
	MMarshalJSON bean.ByteSliceErrGetter
	MToBytes     bean.ByteSliceGetter
	MFromBytes   bean.ByteSliceErrSetter

	MType struct {
		Calls   uint
		Returns FrameType
	}
}

func (f *mockFrame) String() string {
	return f.MString.Get()
}

func (f *mockFrame) MarshalJSON() ([]byte, error) {
	return f.MMarshalJSON.Get()
}

func (f *mockFrame) Type() FrameType {
	f.MType.Calls++
	return f.MType.Returns
}

func (f *mockFrame) ToBytes() []byte {
	return f.MToBytes.Get()
}

func (f *mockFrame) FromBytes(b []byte) error {
	return f.MFromBytes.Set(b)
}
