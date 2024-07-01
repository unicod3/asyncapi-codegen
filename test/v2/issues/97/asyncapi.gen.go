// Package "issue97" provides primitives to interact with the AsyncAPI specification.
//
// Code generated by github.com/lerenn/asyncapi-codegen version (devel) DO NOT EDIT.
package issue97

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/lerenn/asyncapi-codegen/pkg/extensions"
)

// AppController is the structure that provides publishing capabilities to the
// developer and and connect the broker with the App
type AppController struct {
	controller
}

// NewAppController links the App to the broker
func NewAppController(bc extensions.BrokerController, options ...ControllerOption) (*AppController, error) {
	// Check if broker controller has been provided
	if bc == nil {
		return nil, extensions.ErrNilBrokerController
	}

	// Create default controller
	controller := controller{
		broker:        bc,
		subscriptions: make(map[string]extensions.BrokerChannelSubscription),
		logger:        extensions.DummyLogger{},
		middlewares:   make([]extensions.Middleware, 0),
		errorHandler:  extensions.DefaultErrorHandler(),
	}

	// Apply options
	for _, option := range options {
		option(&controller)
	}

	return &AppController{controller: controller}, nil
}

func (c AppController) wrapMiddlewares(
	middlewares []extensions.Middleware,
	callback extensions.NextMiddleware,
) func(ctx context.Context, msg *extensions.BrokerMessage) error {
	var called bool

	// If there is no more middleware
	if len(middlewares) == 0 {
		return func(ctx context.Context, msg *extensions.BrokerMessage) error {
			// Call the callback if it exists and it has not been called already
			if callback != nil && !called {
				called = true
				return callback(ctx)
			}

			// Nil can be returned, as the callback has already been called
			return nil
		}
	}

	// Get the next function to call from next middlewares or callback
	next := c.wrapMiddlewares(middlewares[1:], callback)

	// Wrap middleware into a check function that will call execute the middleware
	// and call the next wrapped middleware if the returned function has not been
	// called already
	return func(ctx context.Context, msg *extensions.BrokerMessage) error {
		// Call the middleware and the following if it has not been done already
		if !called {
			// Create the next call with the context and the message
			nextWithArgs := func(ctx context.Context) error {
				return next(ctx, msg)
			}

			// Call the middleware and register it as already called
			called = true
			if err := middlewares[0](ctx, msg, nextWithArgs); err != nil {
				return err
			}

			// If next has already been called in middleware, it should not be executed again
			return nextWithArgs(ctx)
		}

		// Nil can be returned, as the next middleware has already been called
		return nil
	}
}

func (c AppController) executeMiddlewares(ctx context.Context, msg *extensions.BrokerMessage, callback extensions.NextMiddleware) error {
	// Wrap middleware to have 'next' function when calling them
	wrapped := c.wrapMiddlewares(c.middlewares, callback)

	// Execute wrapped middlewares
	return wrapped(ctx, msg)
}

func addAppContextValues(ctx context.Context, path string) context.Context {
	ctx = context.WithValue(ctx, extensions.ContextKeyIsVersion, "1.0.0")
	ctx = context.WithValue(ctx, extensions.ContextKeyIsProvider, "app")
	return context.WithValue(ctx, extensions.ContextKeyIsChannel, path)
}

// Close will clean up any existing resources on the controller
func (c *AppController) Close(ctx context.Context) {
	// Unsubscribing remaining channels
}

// PublishV2Issue97ReferencePayloadArray will publish messages to 'v2.issue97.referencePayloadArray' channel
func (c *AppController) PublishV2Issue97ReferencePayloadArray(
	ctx context.Context,
	msg ReferencePayloadArrayMessage,
) error {
	// Get channel path
	path := "v2.issue97.referencePayloadArray"

	// Set context
	ctx = addAppContextValues(ctx, path)
	ctx = context.WithValue(ctx, extensions.ContextKeyIsDirection, "publication")

	// Convert to BrokerMessage
	brokerMsg, err := msg.toBrokerMessage()
	if err != nil {
		return err
	}

	// Set broker message to context
	ctx = context.WithValue(ctx, extensions.ContextKeyIsBrokerMessage, brokerMsg.String())

	// Publish the message on event-broker through middlewares
	return c.executeMiddlewares(ctx, &brokerMsg, func(ctx context.Context) error {
		return c.broker.Publish(ctx, path, brokerMsg)
	})
}

// PublishV2Issue97ReferencePayloadObject will publish messages to 'v2.issue97.referencePayloadObject' channel
func (c *AppController) PublishV2Issue97ReferencePayloadObject(
	ctx context.Context,
	msg ReferencePayloadObjectMessage,
) error {
	// Get channel path
	path := "v2.issue97.referencePayloadObject"

	// Set context
	ctx = addAppContextValues(ctx, path)
	ctx = context.WithValue(ctx, extensions.ContextKeyIsDirection, "publication")

	// Convert to BrokerMessage
	brokerMsg, err := msg.toBrokerMessage()
	if err != nil {
		return err
	}

	// Set broker message to context
	ctx = context.WithValue(ctx, extensions.ContextKeyIsBrokerMessage, brokerMsg.String())

	// Publish the message on event-broker through middlewares
	return c.executeMiddlewares(ctx, &brokerMsg, func(ctx context.Context) error {
		return c.broker.Publish(ctx, path, brokerMsg)
	})
}

// PublishV2Issue97ReferencePayloadString will publish messages to 'v2.issue97.referencePayloadString' channel
func (c *AppController) PublishV2Issue97ReferencePayloadString(
	ctx context.Context,
	msg ReferencePayloadStringMessage,
) error {
	// Get channel path
	path := "v2.issue97.referencePayloadString"

	// Set context
	ctx = addAppContextValues(ctx, path)
	ctx = context.WithValue(ctx, extensions.ContextKeyIsDirection, "publication")

	// Convert to BrokerMessage
	brokerMsg, err := msg.toBrokerMessage()
	if err != nil {
		return err
	}

	// Set broker message to context
	ctx = context.WithValue(ctx, extensions.ContextKeyIsBrokerMessage, brokerMsg.String())

	// Publish the message on event-broker through middlewares
	return c.executeMiddlewares(ctx, &brokerMsg, func(ctx context.Context) error {
		return c.broker.Publish(ctx, path, brokerMsg)
	})
}

// UserSubscriber represents all handlers that are expecting messages for User
type UserSubscriber interface {
	// V2Issue97ReferencePayloadArray subscribes to messages placed on the 'v2.issue97.referencePayloadArray' channel
	V2Issue97ReferencePayloadArray(ctx context.Context, msg ReferencePayloadArrayMessage) error

	// V2Issue97ReferencePayloadObject subscribes to messages placed on the 'v2.issue97.referencePayloadObject' channel
	V2Issue97ReferencePayloadObject(ctx context.Context, msg ReferencePayloadObjectMessage) error

	// V2Issue97ReferencePayloadString subscribes to messages placed on the 'v2.issue97.referencePayloadString' channel
	V2Issue97ReferencePayloadString(ctx context.Context, msg ReferencePayloadStringMessage) error
}

// UserController is the structure that provides publishing capabilities to the
// developer and and connect the broker with the User
type UserController struct {
	controller
}

// NewUserController links the User to the broker
func NewUserController(bc extensions.BrokerController, options ...ControllerOption) (*UserController, error) {
	// Check if broker controller has been provided
	if bc == nil {
		return nil, extensions.ErrNilBrokerController
	}

	// Create default controller
	controller := controller{
		broker:        bc,
		subscriptions: make(map[string]extensions.BrokerChannelSubscription),
		logger:        extensions.DummyLogger{},
		middlewares:   make([]extensions.Middleware, 0),
		errorHandler:  extensions.DefaultErrorHandler(),
	}

	// Apply options
	for _, option := range options {
		option(&controller)
	}

	return &UserController{controller: controller}, nil
}

func (c UserController) wrapMiddlewares(
	middlewares []extensions.Middleware,
	callback extensions.NextMiddleware,
) func(ctx context.Context, msg *extensions.BrokerMessage) error {
	var called bool

	// If there is no more middleware
	if len(middlewares) == 0 {
		return func(ctx context.Context, msg *extensions.BrokerMessage) error {
			// Call the callback if it exists and it has not been called already
			if callback != nil && !called {
				called = true
				return callback(ctx)
			}

			// Nil can be returned, as the callback has already been called
			return nil
		}
	}

	// Get the next function to call from next middlewares or callback
	next := c.wrapMiddlewares(middlewares[1:], callback)

	// Wrap middleware into a check function that will call execute the middleware
	// and call the next wrapped middleware if the returned function has not been
	// called already
	return func(ctx context.Context, msg *extensions.BrokerMessage) error {
		// Call the middleware and the following if it has not been done already
		if !called {
			// Create the next call with the context and the message
			nextWithArgs := func(ctx context.Context) error {
				return next(ctx, msg)
			}

			// Call the middleware and register it as already called
			called = true
			if err := middlewares[0](ctx, msg, nextWithArgs); err != nil {
				return err
			}

			// If next has already been called in middleware, it should not be executed again
			return nextWithArgs(ctx)
		}

		// Nil can be returned, as the next middleware has already been called
		return nil
	}
}

func (c UserController) executeMiddlewares(ctx context.Context, msg *extensions.BrokerMessage, callback extensions.NextMiddleware) error {
	// Wrap middleware to have 'next' function when calling them
	wrapped := c.wrapMiddlewares(c.middlewares, callback)

	// Execute wrapped middlewares
	return wrapped(ctx, msg)
}

func addUserContextValues(ctx context.Context, path string) context.Context {
	ctx = context.WithValue(ctx, extensions.ContextKeyIsVersion, "1.0.0")
	ctx = context.WithValue(ctx, extensions.ContextKeyIsProvider, "user")
	return context.WithValue(ctx, extensions.ContextKeyIsChannel, path)
}

// Close will clean up any existing resources on the controller
func (c *UserController) Close(ctx context.Context) {
	// Unsubscribing remaining channels
	c.UnsubscribeAll(ctx)

	c.logger.Info(ctx, "Closed user controller")
}

// SubscribeAll will subscribe to channels without parameters on which the app is expecting messages.
// For channels with parameters, they should be subscribed independently.
func (c *UserController) SubscribeAll(ctx context.Context, as UserSubscriber) error {
	if as == nil {
		return extensions.ErrNilUserSubscriber
	}

	if err := c.SubscribeV2Issue97ReferencePayloadArray(ctx, as.V2Issue97ReferencePayloadArray); err != nil {
		return err
	}
	if err := c.SubscribeV2Issue97ReferencePayloadObject(ctx, as.V2Issue97ReferencePayloadObject); err != nil {
		return err
	}
	if err := c.SubscribeV2Issue97ReferencePayloadString(ctx, as.V2Issue97ReferencePayloadString); err != nil {
		return err
	}

	return nil
}

// UnsubscribeAll will unsubscribe all remaining subscribed channels
func (c *UserController) UnsubscribeAll(ctx context.Context) {
	c.UnsubscribeV2Issue97ReferencePayloadArray(ctx)
	c.UnsubscribeV2Issue97ReferencePayloadObject(ctx)
	c.UnsubscribeV2Issue97ReferencePayloadString(ctx)
}

// SubscribeV2Issue97ReferencePayloadArray will subscribe to new messages from 'v2.issue97.referencePayloadArray' channel.
//
// Callback function 'fn' will be called each time a new message is received.
func (c *UserController) SubscribeV2Issue97ReferencePayloadArray(
	ctx context.Context,
	fn func(ctx context.Context, msg ReferencePayloadArrayMessage) error,
) error {
	// Get channel path
	path := "v2.issue97.referencePayloadArray"

	// Set context
	ctx = addUserContextValues(ctx, path)
	ctx = context.WithValue(ctx, extensions.ContextKeyIsDirection, "reception")

	// Check if there is already a subscription
	_, exists := c.subscriptions[path]
	if exists {
		err := fmt.Errorf("%w: %q channel is already subscribed", extensions.ErrAlreadySubscribedChannel, path)
		c.logger.Error(ctx, err.Error())
		return err
	}

	// Subscribe to broker channel
	sub, err := c.broker.Subscribe(ctx, path)
	if err != nil {
		c.logger.Error(ctx, err.Error())
		return err
	}
	c.logger.Info(ctx, "Subscribed to channel")

	// Asynchronously listen to new messages and pass them to app subscriber
	go func() {
		for {
			// Listen to next message
			stop, err := c.listenToV2Issue97ReferencePayloadArrayNextMessage(path, sub, fn)
			if err != nil {
				c.logger.Error(ctx, err.Error())
			}

			// Stop if required
			if stop {
				return
			}
		}
	}()

	// Add the cancel channel to the inside map
	c.subscriptions[path] = sub

	return nil
}

func (c *UserController) listenToV2Issue97ReferencePayloadArrayNextMessage(
	path string,
	sub extensions.BrokerChannelSubscription,
	fn func(ctx context.Context, msg ReferencePayloadArrayMessage) error,
) (stop bool, err error) {
	// Create a context for the received response
	msgCtx, cancel := context.WithCancel(context.Background())
	msgCtx = addUserContextValues(msgCtx, path)
	msgCtx = context.WithValue(msgCtx, extensions.ContextKeyIsDirection, "reception")
	defer cancel()

	// Wait for next message
	acknowledgeableBrokerMessage, open := <-sub.MessagesChannel()

	// If subscription is closed and there is no more message
	// (i.e. uninitialized message), then exit the function
	if !open && acknowledgeableBrokerMessage.IsUninitialized() {
		return true, nil
	}

	// Set broker message to context
	msgCtx = context.WithValue(msgCtx, extensions.ContextKeyIsBrokerMessage, acknowledgeableBrokerMessage.String())

	// Execute middlewares before handling the message
	if err := c.executeMiddlewares(msgCtx, &acknowledgeableBrokerMessage.BrokerMessage, func(middlewareCtx context.Context) error {
		// Process message
		msg, err := brokerMessageToReferencePayloadArrayMessage(acknowledgeableBrokerMessage.BrokerMessage)
		if err != nil {
			return err
		}

		// Execute the subscription function
		if err := fn(middlewareCtx, msg); err != nil {
			return err
		}

		acknowledgeableBrokerMessage.Ack()

		return nil
	}); err != nil {
		c.errorHandler(msgCtx, path, &acknowledgeableBrokerMessage, err)
		// On error execute the acknowledgeableBrokerMessage nack() function and
		// let the BrokerAcknowledgment decide what is the right nack behavior for the broker
		acknowledgeableBrokerMessage.Nak()
	}

	return false, nil
}

// UnsubscribeV2Issue97ReferencePayloadArray will unsubscribe messages from 'v2.issue97.referencePayloadArray' channel.
// A timeout can be set in context to avoid blocking operation, if needed.
func (c *UserController) UnsubscribeV2Issue97ReferencePayloadArray(ctx context.Context) {
	// Get channel path
	path := "v2.issue97.referencePayloadArray"

	// Check if there subscribers for this channel
	sub, exists := c.subscriptions[path]
	if !exists {
		return
	}

	// Set context
	ctx = addUserContextValues(ctx, path)

	// Stop the subscription
	sub.Cancel(ctx)

	// Remove if from the subscribers
	delete(c.subscriptions, path)

	c.logger.Info(ctx, "Unsubscribed from channel")
}

// SubscribeV2Issue97ReferencePayloadObject will subscribe to new messages from 'v2.issue97.referencePayloadObject' channel.
//
// Callback function 'fn' will be called each time a new message is received.
func (c *UserController) SubscribeV2Issue97ReferencePayloadObject(
	ctx context.Context,
	fn func(ctx context.Context, msg ReferencePayloadObjectMessage) error,
) error {
	// Get channel path
	path := "v2.issue97.referencePayloadObject"

	// Set context
	ctx = addUserContextValues(ctx, path)
	ctx = context.WithValue(ctx, extensions.ContextKeyIsDirection, "reception")

	// Check if there is already a subscription
	_, exists := c.subscriptions[path]
	if exists {
		err := fmt.Errorf("%w: %q channel is already subscribed", extensions.ErrAlreadySubscribedChannel, path)
		c.logger.Error(ctx, err.Error())
		return err
	}

	// Subscribe to broker channel
	sub, err := c.broker.Subscribe(ctx, path)
	if err != nil {
		c.logger.Error(ctx, err.Error())
		return err
	}
	c.logger.Info(ctx, "Subscribed to channel")

	// Asynchronously listen to new messages and pass them to app subscriber
	go func() {
		for {
			// Listen to next message
			stop, err := c.listenToV2Issue97ReferencePayloadObjectNextMessage(path, sub, fn)
			if err != nil {
				c.logger.Error(ctx, err.Error())
			}

			// Stop if required
			if stop {
				return
			}
		}
	}()

	// Add the cancel channel to the inside map
	c.subscriptions[path] = sub

	return nil
}

func (c *UserController) listenToV2Issue97ReferencePayloadObjectNextMessage(
	path string,
	sub extensions.BrokerChannelSubscription,
	fn func(ctx context.Context, msg ReferencePayloadObjectMessage) error,
) (stop bool, err error) {
	// Create a context for the received response
	msgCtx, cancel := context.WithCancel(context.Background())
	msgCtx = addUserContextValues(msgCtx, path)
	msgCtx = context.WithValue(msgCtx, extensions.ContextKeyIsDirection, "reception")
	defer cancel()

	// Wait for next message
	acknowledgeableBrokerMessage, open := <-sub.MessagesChannel()

	// If subscription is closed and there is no more message
	// (i.e. uninitialized message), then exit the function
	if !open && acknowledgeableBrokerMessage.IsUninitialized() {
		return true, nil
	}

	// Set broker message to context
	msgCtx = context.WithValue(msgCtx, extensions.ContextKeyIsBrokerMessage, acknowledgeableBrokerMessage.String())

	// Execute middlewares before handling the message
	if err := c.executeMiddlewares(msgCtx, &acknowledgeableBrokerMessage.BrokerMessage, func(middlewareCtx context.Context) error {
		// Process message
		msg, err := brokerMessageToReferencePayloadObjectMessage(acknowledgeableBrokerMessage.BrokerMessage)
		if err != nil {
			return err
		}

		// Execute the subscription function
		if err := fn(middlewareCtx, msg); err != nil {
			return err
		}

		acknowledgeableBrokerMessage.Ack()

		return nil
	}); err != nil {
		c.errorHandler(msgCtx, path, &acknowledgeableBrokerMessage, err)
		// On error execute the acknowledgeableBrokerMessage nack() function and
		// let the BrokerAcknowledgment decide what is the right nack behavior for the broker
		acknowledgeableBrokerMessage.Nak()
	}

	return false, nil
}

// UnsubscribeV2Issue97ReferencePayloadObject will unsubscribe messages from 'v2.issue97.referencePayloadObject' channel.
// A timeout can be set in context to avoid blocking operation, if needed.
func (c *UserController) UnsubscribeV2Issue97ReferencePayloadObject(ctx context.Context) {
	// Get channel path
	path := "v2.issue97.referencePayloadObject"

	// Check if there subscribers for this channel
	sub, exists := c.subscriptions[path]
	if !exists {
		return
	}

	// Set context
	ctx = addUserContextValues(ctx, path)

	// Stop the subscription
	sub.Cancel(ctx)

	// Remove if from the subscribers
	delete(c.subscriptions, path)

	c.logger.Info(ctx, "Unsubscribed from channel")
}

// SubscribeV2Issue97ReferencePayloadString will subscribe to new messages from 'v2.issue97.referencePayloadString' channel.
//
// Callback function 'fn' will be called each time a new message is received.
func (c *UserController) SubscribeV2Issue97ReferencePayloadString(
	ctx context.Context,
	fn func(ctx context.Context, msg ReferencePayloadStringMessage) error,
) error {
	// Get channel path
	path := "v2.issue97.referencePayloadString"

	// Set context
	ctx = addUserContextValues(ctx, path)
	ctx = context.WithValue(ctx, extensions.ContextKeyIsDirection, "reception")

	// Check if there is already a subscription
	_, exists := c.subscriptions[path]
	if exists {
		err := fmt.Errorf("%w: %q channel is already subscribed", extensions.ErrAlreadySubscribedChannel, path)
		c.logger.Error(ctx, err.Error())
		return err
	}

	// Subscribe to broker channel
	sub, err := c.broker.Subscribe(ctx, path)
	if err != nil {
		c.logger.Error(ctx, err.Error())
		return err
	}
	c.logger.Info(ctx, "Subscribed to channel")

	// Asynchronously listen to new messages and pass them to app subscriber
	go func() {
		for {
			// Listen to next message
			stop, err := c.listenToV2Issue97ReferencePayloadStringNextMessage(path, sub, fn)
			if err != nil {
				c.logger.Error(ctx, err.Error())
			}

			// Stop if required
			if stop {
				return
			}
		}
	}()

	// Add the cancel channel to the inside map
	c.subscriptions[path] = sub

	return nil
}

func (c *UserController) listenToV2Issue97ReferencePayloadStringNextMessage(
	path string,
	sub extensions.BrokerChannelSubscription,
	fn func(ctx context.Context, msg ReferencePayloadStringMessage) error,
) (stop bool, err error) {
	// Create a context for the received response
	msgCtx, cancel := context.WithCancel(context.Background())
	msgCtx = addUserContextValues(msgCtx, path)
	msgCtx = context.WithValue(msgCtx, extensions.ContextKeyIsDirection, "reception")
	defer cancel()

	// Wait for next message
	acknowledgeableBrokerMessage, open := <-sub.MessagesChannel()

	// If subscription is closed and there is no more message
	// (i.e. uninitialized message), then exit the function
	if !open && acknowledgeableBrokerMessage.IsUninitialized() {
		return true, nil
	}

	// Set broker message to context
	msgCtx = context.WithValue(msgCtx, extensions.ContextKeyIsBrokerMessage, acknowledgeableBrokerMessage.String())

	// Execute middlewares before handling the message
	if err := c.executeMiddlewares(msgCtx, &acknowledgeableBrokerMessage.BrokerMessage, func(middlewareCtx context.Context) error {
		// Process message
		msg, err := brokerMessageToReferencePayloadStringMessage(acknowledgeableBrokerMessage.BrokerMessage)
		if err != nil {
			return err
		}

		// Execute the subscription function
		if err := fn(middlewareCtx, msg); err != nil {
			return err
		}

		acknowledgeableBrokerMessage.Ack()

		return nil
	}); err != nil {
		c.errorHandler(msgCtx, path, &acknowledgeableBrokerMessage, err)
		// On error execute the acknowledgeableBrokerMessage nack() function and
		// let the BrokerAcknowledgment decide what is the right nack behavior for the broker
		acknowledgeableBrokerMessage.Nak()
	}

	return false, nil
}

// UnsubscribeV2Issue97ReferencePayloadString will unsubscribe messages from 'v2.issue97.referencePayloadString' channel.
// A timeout can be set in context to avoid blocking operation, if needed.
func (c *UserController) UnsubscribeV2Issue97ReferencePayloadString(ctx context.Context) {
	// Get channel path
	path := "v2.issue97.referencePayloadString"

	// Check if there subscribers for this channel
	sub, exists := c.subscriptions[path]
	if !exists {
		return
	}

	// Set context
	ctx = addUserContextValues(ctx, path)

	// Stop the subscription
	sub.Cancel(ctx)

	// Remove if from the subscribers
	delete(c.subscriptions, path)

	c.logger.Info(ctx, "Unsubscribed from channel")
}

// AsyncAPIVersion is the version of the used AsyncAPI document
const AsyncAPIVersion = "1.0.0"

// controller is the controller that will be used to communicate with the broker
// It will be used internally by AppController and UserController
type controller struct {
	// broker is the broker controller that will be used to communicate
	broker extensions.BrokerController
	// subscriptions is a map of all subscriptions
	subscriptions map[string]extensions.BrokerChannelSubscription
	// logger is the logger that will be used² to log operations on controller
	logger extensions.Logger
	// middlewares are the middlewares that will be executed when sending or
	// receiving messages
	middlewares []extensions.Middleware
	// handler to handle errors from consumers and middlewares
	errorHandler extensions.ErrorHandler
}

// ControllerOption is the type of the options that can be passed
// when creating a new Controller
type ControllerOption func(controller *controller)

// WithLogger attaches a logger to the controller
func WithLogger(logger extensions.Logger) ControllerOption {
	return func(controller *controller) {
		controller.logger = logger
	}
}

// WithMiddlewares attaches middlewares that will be executed when sending or receiving messages
func WithMiddlewares(middlewares ...extensions.Middleware) ControllerOption {
	return func(controller *controller) {
		controller.middlewares = middlewares
	}
}

// WithErrorHandler attaches a errorhandler to handle errors from subscriber functions
func WithErrorHandler(handler extensions.ErrorHandler) ControllerOption {
	return func(controller *controller) {
		controller.errorHandler = handler
	}
}

type MessageWithCorrelationID interface {
	CorrelationID() string
	SetCorrelationID(id string)
}

type Error struct {
	Channel string
	Err     error
}

func (e *Error) Error() string {
	return fmt.Sprintf("channel %q: err %v", e.Channel, e.Err)
}

// ReferencePayloadArrayMessage is the message expected for 'ReferencePayloadArrayMessage' channel.
type ReferencePayloadArrayMessage struct {
	// Payload will be inserted in the message payload
	Payload ArraySchema
}

func NewReferencePayloadArrayMessage() ReferencePayloadArrayMessage {
	var msg ReferencePayloadArrayMessage

	return msg
}

// brokerMessageToReferencePayloadArrayMessage will fill a new ReferencePayloadArrayMessage with data from generic broker message
func brokerMessageToReferencePayloadArrayMessage(bMsg extensions.BrokerMessage) (ReferencePayloadArrayMessage, error) {
	var msg ReferencePayloadArrayMessage

	// Unmarshal payload to expected message payload format
	err := json.Unmarshal(bMsg.Payload, &msg.Payload)
	if err != nil {
		return msg, err
	}

	// TODO: run checks on msg type

	return msg, nil
}

// toBrokerMessage will generate a generic broker message from ReferencePayloadArrayMessage data
func (msg ReferencePayloadArrayMessage) toBrokerMessage() (extensions.BrokerMessage, error) {
	// TODO: implement checks on message

	// Marshal payload to JSON
	payload, err := json.Marshal(msg.Payload)
	if err != nil {
		return extensions.BrokerMessage{}, err
	}

	// There is no headers here
	headers := make(map[string][]byte, 0)

	return extensions.BrokerMessage{
		Headers: headers,
		Payload: payload,
	}, nil
}

// ReferencePayloadObjectMessage is the message expected for 'ReferencePayloadObjectMessage' channel.
type ReferencePayloadObjectMessage struct {
	// Payload will be inserted in the message payload
	Payload ObjectSchema
}

func NewReferencePayloadObjectMessage() ReferencePayloadObjectMessage {
	var msg ReferencePayloadObjectMessage

	return msg
}

// brokerMessageToReferencePayloadObjectMessage will fill a new ReferencePayloadObjectMessage with data from generic broker message
func brokerMessageToReferencePayloadObjectMessage(bMsg extensions.BrokerMessage) (ReferencePayloadObjectMessage, error) {
	var msg ReferencePayloadObjectMessage

	// Unmarshal payload to expected message payload format
	err := json.Unmarshal(bMsg.Payload, &msg.Payload)
	if err != nil {
		return msg, err
	}

	// TODO: run checks on msg type

	return msg, nil
}

// toBrokerMessage will generate a generic broker message from ReferencePayloadObjectMessage data
func (msg ReferencePayloadObjectMessage) toBrokerMessage() (extensions.BrokerMessage, error) {
	// TODO: implement checks on message

	// Marshal payload to JSON
	payload, err := json.Marshal(msg.Payload)
	if err != nil {
		return extensions.BrokerMessage{}, err
	}

	// There is no headers here
	headers := make(map[string][]byte, 0)

	return extensions.BrokerMessage{
		Headers: headers,
		Payload: payload,
	}, nil
}

// ReferencePayloadStringMessage is the message expected for 'ReferencePayloadStringMessage' channel.
type ReferencePayloadStringMessage struct {
	// Payload will be inserted in the message payload
	Payload StringSchema
}

func NewReferencePayloadStringMessage() ReferencePayloadStringMessage {
	var msg ReferencePayloadStringMessage

	return msg
}

// brokerMessageToReferencePayloadStringMessage will fill a new ReferencePayloadStringMessage with data from generic broker message
func brokerMessageToReferencePayloadStringMessage(bMsg extensions.BrokerMessage) (ReferencePayloadStringMessage, error) {
	var msg ReferencePayloadStringMessage

	// Convert to string
	payload := string(bMsg.Payload)
	msg.Payload = StringSchema(payload)

	// TODO: run checks on msg type

	return msg, nil
}

// toBrokerMessage will generate a generic broker message from ReferencePayloadStringMessage data
func (msg ReferencePayloadStringMessage) toBrokerMessage() (extensions.BrokerMessage, error) {
	// TODO: implement checks on message

	// Convert to []byte
	payload := []byte(msg.Payload)

	// There is no headers here
	headers := make(map[string][]byte, 0)

	return extensions.BrokerMessage{
		Headers: headers,
		Payload: payload,
	}, nil
}

// ArraySchema is a schema from the AsyncAPI specification required in messages
type ArraySchema []string

// ObjectSchema is a schema from the AsyncAPI specification required in messages
type ObjectSchema struct {
	Text *string `json:"text"`
}

// StringSchema is a schema from the AsyncAPI specification required in messages
type StringSchema string

const (
	// V2Issue97ReferencePayloadArrayPath is the constant representing the 'V2Issue97ReferencePayloadArray' channel path.
	V2Issue97ReferencePayloadArrayPath = "v2.issue97.referencePayloadArray"
	// V2Issue97ReferencePayloadObjectPath is the constant representing the 'V2Issue97ReferencePayloadObject' channel path.
	V2Issue97ReferencePayloadObjectPath = "v2.issue97.referencePayloadObject"
	// V2Issue97ReferencePayloadStringPath is the constant representing the 'V2Issue97ReferencePayloadString' channel path.
	V2Issue97ReferencePayloadStringPath = "v2.issue97.referencePayloadString"
)

// ChannelsPaths is an array of all channels paths
var ChannelsPaths = []string{
	V2Issue97ReferencePayloadArrayPath,
	V2Issue97ReferencePayloadObjectPath,
	V2Issue97ReferencePayloadStringPath,
}
