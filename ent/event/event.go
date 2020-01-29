// Code generated by entc, DO NOT EDIT.

package event

import (
	"fmt"
	"time"

	"github.com/kcarretto/paragon/ent/schema"
)

const (
	// Label holds the string label denoting the event type in the database.
	Label = "event"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreationTime holds the string denoting the creationtime vertex property in the database.
	FieldCreationTime = "creation_time"
	// FieldKind holds the string denoting the kind vertex property in the database.
	FieldKind = "kind"

	// Table holds the table name of the event in the database.
	Table = "events"
	// JobTable is the table the holds the job relation/edge.
	JobTable = "events"
	// JobInverseTable is the table name for the Job entity.
	// It exists in this package in order to avoid circular dependency with the "job" package.
	JobInverseTable = "jobs"
	// JobColumn is the table column denoting the job relation/edge.
	JobColumn = "event_job_id"
	// FileTable is the table the holds the file relation/edge.
	FileTable = "events"
	// FileInverseTable is the table name for the File entity.
	// It exists in this package in order to avoid circular dependency with the "file" package.
	FileInverseTable = "files"
	// FileColumn is the table column denoting the file relation/edge.
	FileColumn = "event_file_id"
	// CredentialTable is the table the holds the credential relation/edge.
	CredentialTable = "events"
	// CredentialInverseTable is the table name for the Credential entity.
	// It exists in this package in order to avoid circular dependency with the "credential" package.
	CredentialInverseTable = "credentials"
	// CredentialColumn is the table column denoting the credential relation/edge.
	CredentialColumn = "event_credential_id"
	// LinkTable is the table the holds the link relation/edge.
	LinkTable = "events"
	// LinkInverseTable is the table name for the Link entity.
	// It exists in this package in order to avoid circular dependency with the "link" package.
	LinkInverseTable = "links"
	// LinkColumn is the table column denoting the link relation/edge.
	LinkColumn = "event_link_id"
	// TagTable is the table the holds the tag relation/edge.
	TagTable = "events"
	// TagInverseTable is the table name for the Tag entity.
	// It exists in this package in order to avoid circular dependency with the "tag" package.
	TagInverseTable = "tags"
	// TagColumn is the table column denoting the tag relation/edge.
	TagColumn = "event_tag_id"
	// TargetTable is the table the holds the target relation/edge.
	TargetTable = "events"
	// TargetInverseTable is the table name for the Target entity.
	// It exists in this package in order to avoid circular dependency with the "target" package.
	TargetInverseTable = "targets"
	// TargetColumn is the table column denoting the target relation/edge.
	TargetColumn = "event_target_id"
	// TaskTable is the table the holds the task relation/edge.
	TaskTable = "events"
	// TaskInverseTable is the table name for the Task entity.
	// It exists in this package in order to avoid circular dependency with the "task" package.
	TaskInverseTable = "tasks"
	// TaskColumn is the table column denoting the task relation/edge.
	TaskColumn = "event_task_id"
	// UserTable is the table the holds the user relation/edge.
	UserTable = "events"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "event_user_id"
	// EventTable is the table the holds the event relation/edge.
	EventTable = "events"
	// EventColumn is the table column denoting the event relation/edge.
	EventColumn = "event_event_id"
	// ServiceTable is the table the holds the service relation/edge.
	ServiceTable = "events"
	// ServiceInverseTable is the table name for the Service entity.
	// It exists in this package in order to avoid circular dependency with the "service" package.
	ServiceInverseTable = "services"
	// ServiceColumn is the table column denoting the service relation/edge.
	ServiceColumn = "event_service_id"
	// LikersTable is the table the holds the likers relation/edge.
	LikersTable = "users"
	// LikersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	LikersInverseTable = "users"
	// LikersColumn is the table column denoting the likers relation/edge.
	LikersColumn = "event_liker_id"
	// OwnerTable is the table the holds the owner relation/edge.
	OwnerTable = "events"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "owner_id"
)

// Columns holds all SQL columns for event fields.
var Columns = []string{
	FieldID,
	FieldCreationTime,
	FieldKind,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Event type.
var ForeignKeys = []string{
	"event_job_id",
	"event_file_id",
	"event_credential_id",
	"event_link_id",
	"event_tag_id",
	"event_target_id",
	"event_task_id",
	"event_user_id",
	"event_event_id",
	"event_service_id",
	"owner_id",
}

var (
	fields = schema.Event{}.Fields()

	// descCreationTime is the schema descriptor for CreationTime field.
	descCreationTime = fields[0].Descriptor()
	// DefaultCreationTime holds the default value on creation for the CreationTime field.
	DefaultCreationTime = descCreationTime.Default.(func() time.Time)
)

// Kind defines the type for the Kind enum field.
type Kind string

// Kind values.
const (
	KindCREATEJOB              Kind = "CREATE_JOB"
	KindCREATETAG              Kind = "CREATE_TAG"
	KindAPPLYTAGTOTASK         Kind = "APPLY_TAG_TO_TASK"
	KindAPPLYTAGTOTARGET       Kind = "APPLY_TAG_TO_TARGET"
	KindAPPLYTAGTOJOB          Kind = "APPLY_TAG_TO_JOB"
	KindREMOVETAGFROMTASK      Kind = "REMOVE_TAG_FROM_TASK"
	KindREMOVETAGFROMTARGET    Kind = "REMOVE_TAG_FROM_TARGET"
	KindREMOVETAGFROMJOB       Kind = "REMOVE_TAG_FROM_JOB"
	KindCREATETARGET           Kind = "CREATE_TARGET"
	KindSETTARGETFIELDS        Kind = "SET_TARGET_FIELDS"
	KindDELETETARGET           Kind = "DELETE_TARGET"
	KindADDCREDENTIALFORTARGET Kind = "ADD_CREDENTIAL_FOR_TARGET"
	KindUPLOADFILE             Kind = "UPLOAD_FILE"
	KindCREATELINK             Kind = "CREATE_LINK"
	KindSETLINKFIELDS          Kind = "SET_LINK_FIELDS"
	KindACTIVATEUSER           Kind = "ACTIVATE_USER"
	KindCREATEUSER             Kind = "CREATE_USER"
	KindMAKEADMIN              Kind = "MAKE_ADMIN"
	KindREMOVEADMIN            Kind = "REMOVE_ADMIN"
	KindCHANGENAME             Kind = "CHANGE_NAME"
	KindACTIVATESERVICE        Kind = "ACTIVATE_SERVICE"
	KindCREATESERVICE          Kind = "CREATE_SERVICE"
	KindLIKEEVENT              Kind = "LIKE_EVENT"
	KindOTHER                  Kind = "OTHER"
)

func (s Kind) String() string {
	return string(s)
}

// KindValidator is a validator for the "k" field enum values. It is called by the builders before save.
func KindValidator(k Kind) error {
	switch k {
	case KindCREATEJOB, KindCREATETAG, KindAPPLYTAGTOTASK, KindAPPLYTAGTOTARGET, KindAPPLYTAGTOJOB, KindREMOVETAGFROMTASK, KindREMOVETAGFROMTARGET, KindREMOVETAGFROMJOB, KindCREATETARGET, KindSETTARGETFIELDS, KindDELETETARGET, KindADDCREDENTIALFORTARGET, KindUPLOADFILE, KindCREATELINK, KindSETLINKFIELDS, KindACTIVATEUSER, KindCREATEUSER, KindMAKEADMIN, KindREMOVEADMIN, KindCHANGENAME, KindACTIVATESERVICE, KindCREATESERVICE, KindLIKEEVENT, KindOTHER:
		return nil
	default:
		return fmt.Errorf("event: invalid enum value for Kind field: %q", k)
	}
}
