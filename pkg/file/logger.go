package file

import (
    "fmt"
    "os"
    "os/signal"
    "syscall"

    "github.com/ilya1st/rotatewriter"
    "github.com/rs/zerolog"
)

func NewLogger(conf Config) zerolog.Logger {

    createLogDirectoryIfNotExists(conf)

    writer, err := initializeRotateWriter(conf)
    if err != nil {
        fmt.Printf("Error initializeRotateWriter: %s\n", err)
        os.Exit(1)
    }

    logger := zerolog.New(writer).With().Timestamp().Logger()

    logger.Info().Msg("Initialized logger")

    return logger
}

func createLogDirectoryIfNotExists(conf Config) {

    fmt.Println("createLogDirectoryIfNotExists. LogDirectory=" + conf.LogDirectory)

    if _, err := os.Stat(conf.LogDirectory); os.IsNotExist(err) {

        err := os.Mkdir(conf.LogDirectory, os.ModePerm)

        if err != nil {
            fmt.Printf("createLogDirectoryIfNotExists. Error creating logs folder: %s\n", err)
            os.Exit(1)
        }
    }
}

func initializeRotateWriter(conf Config) (rw *rotatewriter.RotateWriter, err error) {

    writer, err := rotatewriter.NewRotateWriter(conf.LogDirectory + "/" + conf.LogFilename, 30)
    if err != nil {
        panic(err)
        return nil, err
    }

    // TODO: need to rework this part for daily log rotation
    sighupChan := make(chan os.Signal, 1)
    signal.Notify(sighupChan, syscall.SIGHUP)
    go func() {
        for {
            _, ok := <-sighupChan
            if !ok {
                return
            }
            fmt.Println("Log rotation")
            err := writer.Rotate(nil)

            if err != nil {
                fmt.Printf("Error in log rotation: %s\n", err)
            }
        }
    }()

    return writer, nil
}
