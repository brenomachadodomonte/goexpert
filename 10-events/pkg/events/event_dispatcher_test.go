package events

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type TestEvent struct {
	Name    string
	Payload interface{}
}

func (te *TestEvent) GetName() string {
	return te.Name
}

func (te *TestEvent) GetPayload() interface{} {
	return te.Payload
}

func (te *TestEvent) GetDateTime() time.Time {
	return time.Now()
}

type TestEventHandler struct {
	ID int
}

func (teh *TestEventHandler) Handle(event EventInterface) {
	teh.Handle(event)
}

type EventDispatcherTestSuite struct {
	suite.Suite
	event           TestEvent
	event2          TestEvent
	handler         TestEventHandler
	handler2        TestEventHandler
	handler3        TestEventHandler
	eventDispatcher EventDispatcher
}

func (suite *EventDispatcherTestSuite) SetupTest() {
	suite.event = TestEvent{Name: "EventTest1", Payload: "Payload Test Event 1"}
	suite.event2 = TestEvent{Name: "EventTest2", Payload: "Payload Test Event 2"}

	suite.handler = TestEventHandler{ID: 1}
	suite.handler2 = TestEventHandler{ID: 2}
	suite.handler3 = TestEventHandler{ID: 3}

	suite.eventDispatcher = EventDispatcher{handlers: make(map[string][]EventHandlerInterface)}
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register() {
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	assert.Equal(suite.T(), &suite.handler, suite.eventDispatcher.handlers[suite.event.GetName()][0])
	assert.Equal(suite.T(), &suite.handler2, suite.eventDispatcher.handlers[suite.event.GetName()][1])
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register_WithSameHandler() {
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	assert.Equal(suite.T(), err, ErrHandlerAlreadyRegistered)
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Clear() {
	//Event 1
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	//Event 2
	err = suite.eventDispatcher.Register(suite.event2.GetName(), &suite.handler3)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event2.GetName()]))

	err = suite.eventDispatcher.Clear()
	suite.Nil(err)
	suite.Equal(0, len(suite.eventDispatcher.handlers))
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Has() {
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	assert.True(suite.T(), suite.eventDispatcher.Has(suite.event.GetName(), &suite.handler))
	assert.True(suite.T(), suite.eventDispatcher.Has(suite.event.GetName(), &suite.handler2))
	assert.False(suite.T(), suite.eventDispatcher.Has(suite.event.GetName(), &suite.handler3))
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Remove() {
	//Event 1
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	//Event 2
	err = suite.eventDispatcher.Register(suite.event2.GetName(), &suite.handler3)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event2.GetName()]))

	err = suite.eventDispatcher.Remove(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))
	assert.Equal(suite.T(), &suite.handler2, suite.eventDispatcher.handlers[suite.event.GetName()][0])

	err = suite.eventDispatcher.Remove(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(0, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	err = suite.eventDispatcher.Remove(suite.event2.GetName(), &suite.handler3)
	suite.Nil(err)
	suite.Equal(0, len(suite.eventDispatcher.handlers[suite.event2.GetName()]))

}

type MockHandler struct {
	mock.Mock
}

func (mh *MockHandler) Handle(event EventInterface) {
	mh.Called(event)
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Dispatch() {
	eh := &MockHandler{}
	eh.On("Handle", &suite.event)

	err := suite.eventDispatcher.Register(suite.event.GetName(), eh)
	assert.Nil(suite.T(), err)

	err = suite.eventDispatcher.Dispatch(&suite.event)
	assert.Nil(suite.T(), err)
	eh.AssertExpectations(suite.T())
	eh.AssertNumberOfCalls(suite.T(), "Handle", 1)

}

func TestSuite(t *testing.T) {
	suite.Run(t, new(EventDispatcherTestSuite))
}
