package eddystone

/*⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺*\
▏                                                        ▕
▏  Public Function Types                                 ▕
▏                                                        ▕
\*⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽*/


// TlmFrameFac produces a new base configuration instance
// of the TlmFrame interface
type TlmFrameFac = func() TlmFrame

// UidFrameFac produces a new base configuration instance
// of the UidFrame interface
type UidFrameFac = func() UidFrame

// UrlFrameFac produces a new base configuration instance of
// the UrlFrame interface
type UrlFrameFac = func() UrlFrame


/*⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺⎺*\
▏                                                        ▕
▏  Public Interfaces                                     ▕
▏                                                        ▕
\*⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽⎽*/


// FrameFactory provides factory methods for producing
// instances of the various Eddystone TLM frame types.
type FrameFactory interface {

	// NewTlmFrame produces a new instance of the TlmFrame
	// interface populated with the parsed value of the given
	// byte slice
	NewTlmFrame(b []byte) (TlmFrame, error)

	// NewUidFrame produces a new instance of the UidFrame
	// interface populated with the parsed value of the given
	// byte slice
	NewUidFrame(b []byte) (UidFrame, error)

	// NewUrlFrame produces a new instance of the UrlFrame
	// interface populated with the parsed value of the given
	// byte slice
	NewUrlFrame(b []byte) (UrlFrame, error)

	// TlmFrameFactory overrides the default factory method
	// for creating new TlmFrame instances with the given
	// factory function.
	//
	// After calling TlmFrameFactory, calls to NewTlmFrame()
	// will only use the latest factory function
	TlmFrameFactory(fac TlmFrameFac) FrameFactory

	// UidFrameFactory overrides the default factory method
	// for creating new UidFrame instances with the given
	// factory function.
	//
	// After calling UidFrameFactory, calls to NewUidFrame()
	// will only use the latest factory function
	UidFrameFactory(fac UidFrameFac) FrameFactory

	// UrlFrameFactory overrides the default factory method
	// for creating new UrlFrame instances with the given
	// factory function.
	//
	// After calling UrlFrameFactory, calls to NewUrlFrame()
	// will only use the latest factory function
	UrlFrameFactory(fac UrlFrameFac) FrameFactory
}
