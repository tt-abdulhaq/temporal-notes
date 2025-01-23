package main

import (
	"app"
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	c, err := client.Dial(client.Options{})

	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, "greeting-tasks", worker.Options{})

	w.RegisterWorkflow(app.GreetingSomeOne)

	err = w.Run(worker.InterruptCh())

	if err != nil {
		log.Fatalln("Ubable to start worker", err)
	}

}
