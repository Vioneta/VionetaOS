package common

import (
	"github.com/Vioneta/VionetaOS/codegen/message_bus"
)

// devtype -> action -> event
var EventTypes = []message_bus.EventType{
	{Name: "vionetaos:system:utilization", SourceID: SERVICENAME, PropertyTypeList: []message_bus.PropertyType{}},
	{Name: "vionetaos:file:recover", SourceID: SERVICENAME, PropertyTypeList: []message_bus.PropertyType{}},
	{Name: "vionetaos:file:operate", SourceID: SERVICENAME, PropertyTypeList: []message_bus.PropertyType{}},
}
