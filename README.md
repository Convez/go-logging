# GoLang customizatble logging system

Logging library for golang, customizable with custom logging level and targets.

## Usage

Use the logSystemBuilder to build your logSystem.
The following builder example uses all available features for the library:
```golang
dest1 := new(bytes.Buffer)
dest2 := new(bytes.Buffer)
builder := gologging.NewLogSystemBuilder().
    WithDestination(dest1). // Overwrite the default destination (STDOUT)
    WithAdditionalDestination(dest2). // Add a new destination (*io.Writer)
    WithAdditionalLevelAbove("CUSTOM1", "INFO"). // Create a new custom log level with higher severity than INFO
    WithAdditionalLevelBelow("CUSTOM2", "INFO"). // Create a new custom log level with lower severity than INFO
    WithTimestampFormat("yyyy-MM-dd*HH:mm:ss"). // Changes the timestamp format
    WithFileNameEnabled(false).  // Removes file name from log line
    WithSeverityEnabled(false).  // Remove severity from log line
    WithTimestampEnabled(false)  // Removes the timestamp from log line

logSystem := builder.Build()
```
Default logging levels available in the library:

    ERROR
    WARN 
    INFO 
    DEBUG
    TRACE


Once the logging system is build, a logger for the desired logging level can be retrieved:

```golang
// Logger for standard INFO severity
infoLogger := logSystem.GetLogger(gologging.INFO)
infoLogger.println("Printing info log")
// Logger for custom CUSTOM1 severity
custom1Logger := logSystem.GetLogger("CUSTOM1")
custom1Logger.println("Printing custom log severity")
```

The shown severity level can be updated via the CONVEZ_LOG_LEVEL environment variable.
Custom log severity levels can be used. It defaults to INFO if the requested shown level does not exist.

Default settings:

    destinations:  []io.Writer{os.Stdout},
    timeFormat:    "2006-01-02T15:04:05.999Z",
    showTimeStamp: true,
    logLevels:     []string{ERROR, WARN, INFO, DEBUG, TRACE},
    showFileName:  true,
    showSeverity:  true,
