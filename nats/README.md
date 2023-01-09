# Openware Nats

openware common nats package.

first we need to initialize nats connection

```
connectionString = "localhost:4222"
nc, err := nats.InitNats(connectionString)
```
we need nc for initializing handlers and subscribers

 
### Publisher

we have publisher for publishing event using either nats or Jetstream

to intialize nats and publish an event:
```
pub, err := nats.NewNatsEventPublisher(nc)
// handle error
pub.Publish("foo.bar", []byte("baz"))
```

to initialize jetstream and publish an event:
```
js, err := nats.NewJsEventPublisher(nc)
// handle error
js.Publish("foo.bar", []byte("baz"))
```
keep in mind that before publishing an event you'll also need to have a stream of the topic. to create a stream:
```
err := js.CreateNewEventStream("foo", []string{"foo.baz", "foo.bar"})
// handle error
```

### Subscriber
With subscribers we can subscribe to different topics. If we subscribe using queue, subscribers with the same group name will receive an event once.

to intiialize nats and subscribe to a topic.
we can pass ```<-chan os.Signal``` for subscribing to shutdown event. so at the end it can cleanup 
```
handler := nats.NewNatsHandler(nc, terminationChannel, nats.NewHandlerDefaultConfig())
msgChan := make(chan *nats.Msg)
handler.SubscribeToQueueUsingChannel("foo.baz", "bar", msgChannel)
```