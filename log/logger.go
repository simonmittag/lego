package log

import (
	"github.com/rs/zerolog/log"
)

// Logger is an optional custom logger.


// StdLogger interface for Standard Logger.
type StdLogger interface {
	Fatal(args ...interface{})
	Fatalln(args ...interface{})
	Fatalf(format string, args ...interface{})
	Print(args ...interface{})
	Println(args ...interface{})
	Printf(format string, args ...interface{})
}

// Fatal writes a log entry.
// It uses Logger if not nil, otherwise it uses the default log.Logger.
func Fatal(args ...interface{}) {
	log.Debug().Msgf("%v",args...)
}

// Fatalf writes a log entry.
// It uses Logger if not nil, otherwise it uses the default log.Logger.
func Fatalf(format string, args ...interface{}) {
	log.Debug().Msgf(format,args...)
}

// Print writes a log entry.
// It uses Logger if not nil, otherwise it uses the default log.Logger.
func Print(args ...interface{}) {
	//log.Debug().Msgf("%v",args...)
}

// Println writes a log entry.
// It uses Logger if not nil, otherwise it uses the default log.Logger.
func Println(args ...interface{}) {
	//log.Debug().Msgf("%v",args...)
}

// Printf writes a log entry.
// It uses Logger if not nil, otherwise it uses the default log.Logger.
func Printf(format string, args ...interface{}) {
	//log.Debug().Msgf(format,args...)
}

// Warnf writes a log entry.
func Warnf(format string, args ...interface{}) {
	log.Debug().Msgf(format,args...)
}

// Infof writes a log entry.
func Infof(format string, args ...interface{}) {
	//log.Debug().Msgf(format,args...)
}
