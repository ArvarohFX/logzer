# logzer

Zap wrapper for json-formatted key-value logging.
You can find example of usage at /examples/kw.go or in the block below.

```
func main() {
	// init logger with debug mode
	logger := logzer.New(true)

	logger.Debug("logger.Debug()", zap.String("debug", "enabled"))
	logger.Debug("logger.Debug()", zap.Bool("debug", true))
	logger.Info("logger.Info()",
		zap.Int("client_id", 1),
		zap.String("client_phone", "+79010000001"))
	logger.Info("logger.Info()", zap.Float64("price", 25.64))
	logger.Error("logger.Error()", zap.Error(errors.New("generated error")))

	// rewrite logger with one without debug mode
	logger = logzer.New(false)
	logger.Debug("logger.Debug()", zap.String("debug", "disabled"))
	logger.Info("logger.Info()", zap.Strings("emails", []string{"test@test.org", "example@test.org"}))
	logger.Warn("logger.Warn()", zap.Ints("ids", []int{14562, 64567}))
}
```
