package xlogrus


import (
	"github.com/sirupsen/logrus"
	"io"
)

type LogWriterOpts map[logrus.Level]io.Writer

type filterLevelHook struct {
	formatter     logrus.Formatter
	levels        []logrus.Level
	writerByLevel map[logrus.Level]io.Writer
}

func NewFilterLevelHook(formatter logrus.Formatter, opts LogWriterOpts) logrus.Hook {
	hook := filterLevelHook{
		formatter:     formatter,
		writerByLevel: make(map[logrus.Level]io.Writer),
	}

	maxLevel := len(logrus.AllLevels)
	for level, writer := range opts {
		if maxLevel <= int(level) {
			continue
		}
		hook.writerByLevel[level] = writer
		hook.levels = append(hook.levels, level)
	}
	return &hook
}

func (hook *filterLevelHook) Fire(entry *logrus.Entry) error {

	msg, err := hook.formatter.Format(entry)
	if err != nil {
		return err
	}

	if writer, ok := hook.writerByLevel[entry.Level]; ok {
		_, err = writer.Write([]byte(msg))
	}
	return err
}

func (hook *filterLevelHook) Levels() []logrus.Level {
	return hook.levels
}