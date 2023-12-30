package exception

type Logger interface {
	Panic(args ...interface{})
	Error(args ...interface{})
	Warn(args ...interface{})
	Info(args ...interface{})
	Debug(args ...interface{})
	Trace(args ...interface{})
	Print(args ...interface{})
}

var log Logger

func SetLogger(l Logger) {
	log = l
}

func PanicOnError(e error) {
	if e != nil {
		log.Panic(e)
		panic(e)
	}
}

func LogError(e error) {
	if e != nil {
		log.Error(e, "")
	}
}

func LogErrorContext(e error, context ...string) {
	if e != nil {
		log.Error(e, context)
	}
}

func PanicOnErrorFunc(f func() error) func() {
	return func() {
		err := f()
		if err != nil {
			log.Panic(err)
		}
	}
}
