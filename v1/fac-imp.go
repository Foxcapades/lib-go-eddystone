package eddystone


/*⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺*\
▏                                                        ▕
▏  Internal Static Vars                                  ▕
▏                                                        ▕
\*⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽*/


var defaultFrameFactory FrameFactory = &frameFactory{
	tlm:  func() TlmFrame {return &tlmFrame{}},
	uuid: func() UidFrame {return &uidFrame{}},
	url:  func() UrlFrame {return &urlFrame{}},
}


/*⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺*\
▏                                                        ▕
▏  Public Constructors                                   ▕
▏                                                        ▕
\*⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽*/


// Default returns the current default FrameFactory
// instance.
func Default() FrameFactory {
	return defaultFrameFactory
}

// SetDefaultFactory overrides the default FrameFactory
// instance with the given value
func SetDefaultFactory(f FrameFactory) {
	defaultFrameFactory = f
}


/*⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺*\
▏                                                        ▕
▏  Internal Implementation                               ▕
▏                                                        ▕
\*⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽*/


type frameFactory struct {
	tlm  TlmFrameFac
	uuid UuidFrameFac
	url  UrlFrameFac
}

func (f frameFactory) NewTlmFrame(b []byte) (o TlmFrame, e error) {
	o = f.tlm()
	e = o.FromBytes(b)
	return
}

func (f frameFactory) NewUuidFrame(b []byte) (o UidFrame, e error) {
	o = f.uuid()
	e = o.FromBytes(b)
	return
}

func (f frameFactory) NewUrlFrame(b []byte) (o UrlFrame, e error) {
	o = f.url()
	e = o.FromBytes(b)
	return
}

func (f *frameFactory) TlmFrameFactory(fac TlmFrameFac) FrameFactory {
	f.tlm = fac
	return f
}

func (f *frameFactory) UuidFrameFactory(fac UuidFrameFac) FrameFactory {
	f.uuid = fac
	return f
}

func (f *frameFactory) UrlFrameFactory(fac UrlFrameFac) FrameFactory {
	f.url = fac
	return f
}

