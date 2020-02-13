package main

import (
		"lynk/data-out-stream/internal"
		"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"

	"fmt"
)

// @title ltl-service server API
// @version 1.0
// @description
// @host staging-ltl-service.lynk.co.in
// @BasePath /


func main() {
	// db.Init()
	// ok := db.RunMigration(config.DatabaseURL)
	// if !ok {
	// 	os.Exit(1)
	// 	return
	// }

	// globals.Logger.Info("Listening to Port: ", config.Port)
	// http.ListenAndServe(":"+config.Port, handlers.New())


	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "kafka",
		//key.deserializer
		//value.deserializer
		"group.id": "Consumer_1",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	//var wg sync.WaitGroup
	c.SubscribeTopics([]string{"server1.dbo.Lynk_DriverAvailable_AfterOffline"}, nil)
	//c.SubscribeTopics([]string{"server1.dbo.* "}, nil)   ----- to subscribe to all the topics

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
			
		} else {
			
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
	 // wg.wait()
	c.Close()
	test.Foo()
}
		

