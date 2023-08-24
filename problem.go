package zabbix

// https://www.zabbix.com/documentation/5.0/en/manual/api/reference/problem/object
type Problem struct {
	EventID string `json:"eventid"`

	Source int `json:"source,string"`

	ObjectType int `json:"object,string"`

	ObjectID int `json:"objectid,string"`

	Clock UnixTimestamp `json:"clock"`

	RecoveryEventID string `json:"r_eventid"`

	RecoveryClock UnixTimestamp `json:"r_clock"`

	CorrelationID string `json:"correlationid"`

	UserID string `json:"userid"`

	Name string `json:"name"`

	Acknowledged ZBXBoolean `json:"acknowledged,string"`

	Acknowledges []*ProblemAcknowledges `json:"acknowledges,omitempty"`

	Severity int `json:"severity,string"`

	Suppressed ZBXBoolean `json:"suppressed,string"`

	SuppressionData []*ProblemSuppression `json:"suppression_data,omitempty"`

	Tags []*Tag `json:"tags,omitempty"`
}

// TODO other fields
// "userid": "1",
// "eventid": "1245463",
// "clock": "1472457281",
// "message": "problem solved",
// "action": "6",
// "old_severity": "0",
// "new_severity": "0"
type ProblemAcknowledges struct {
	AcknowledgeID string `json:"acknowledgeid"`
}

type ProblemSuppression struct {
	MaintenanceID string        `json:"maintenanceid"`
	SuppressUtill UnixTimestamp `json:"suppress_until"`
}

// https://www.zabbix.com/documentation/5.0/en/manual/api/reference/problem/get
// ProblemGetParams is query params for problem.get call
type ProblemGetParams struct {
	GetParameters

	// EventIDs filters search results to Events that matched the given Event
	// IDs.
	EventIDs []string `json:"eventids,omitempty"`

	// GroupIDs filters search results to events for hosts that are members of
	// the given Group IDs.
	GroupIDs []string `json:"groupids,omitempty"`

	// HostIDs filters search results to events for hosts that matched the given
	// Host IDs.
	HostIDs []string `json:"hostids,omitempty"`

	// ObjectIDs filters search results to events for Objects that matched
	// the given Object IDs.
	ObjectIDs []string `json:"objectids,omitempty"`

	// ApplicationIDs filters search results to events for Objects that matched
	// the given Application IDs.
	ApplicationIDs []string `json:"applicationids,omitempty"`

	// ObjectType filters search results to events created by the given Object
	// Type. Must be one of the EventObjectType constants.
	//
	// Default: EventObjectTypeTrigger
	ObjectType int `json:"object"`

	// AcknowledgedOnly filters search results to event which have been
	// acknowledged.
	AcknowledgedOnly bool `json:"acknowledged,omitempty"`

	// SuppressedOnly filters search results to event which have been
	// suppresed.
	SuppressedOnly bool `json:"suppressed,omitempty"`

	// Return only problems with given event severities. Applies only if object is trigger.
	Severities []int `json:"severities,omitempty"`

	//TODO
	//evaltype
	//tags

	Recent bool `json:"recent,omitempty"`

	// MinEventID filters search results to Events with an ID greater or equal
	// to the given ID.
	MinEventID string `json:"eventid_from,omitempty"`

	// MaxEventID filters search results to Events with an ID lesser or equal
	// to the given ID.
	MaxEventID string `json:"eventid_till,omitempty"`

	// MinTime filters search results to Problem with a timestamp lesser than or
	// equal to the given timestamp.
	TimeFrom int64 `json:"time_from,omitempty"`

	// MaxTime filters search results to Problem with a timestamp greater than or
	// equal to the given timestamp.
	TimeTill int64 `json:"time_till,omitempty"`

	// Return a suppression_data property with the list of maintenances:
	// maintenanceid - (string) ID of the maintenance;
	// suppress_until - (integer) time until the problem is suppressed.
	SelectSuppressionData SelectQuery `json:"selectRelatedObject,omitempty"`

	// Return an acknowledges property with the problem updates. Problem updates are sorted in reverse chronological order.
	// The problem update object has the following properties:
	// acknowledgeid - (string) update's ID;
	// userid - (string) ID of the user that updated the event;
	// eventid - (string) ID of the updated event;
	// clock - (timestamp) time when the event was updated;
	// message - (string) text of the message;
	// action - (integer)type of update action (see event.acknowledge);
	// old_severity - (integer) event severity before this update action;
	// new_severity - (integer) event severity after this update action;
	// Supports count.
	SelectAcknowledges SelectQuery `json:"selectAcknowledges,omitempty"`

	// Return a tags property with the problem tags.
	SelectTags SelectQuery `json:"selectTags,omitempty"`
}

// GetProblems queries the Zabbix API for Problems matching the given search
// parameters.
//
// ErrEventNotFound is returned if the search result set is empty.
// An error is returned if a transport, parsing or API error occurs.
func (c *Session) GetProblems(params ProblemGetParams) ([]Problem, error) {
	problems := make([]Problem, 0)
	err := c.Get("problem.get", params, &problems)
	if err != nil {
		return nil, err
	}

	if len(problems) == 0 {
		return nil, ErrNotFound
	}

	return problems, nil
}
