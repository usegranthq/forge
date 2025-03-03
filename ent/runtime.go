// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/usegranthq/forge/ent/oidcclient"
	"github.com/usegranthq/forge/ent/project"
	"github.com/usegranthq/forge/ent/projectdomain"
	"github.com/usegranthq/forge/ent/schema"
	"github.com/usegranthq/forge/ent/token"
	"github.com/usegranthq/forge/ent/user"
	"github.com/usegranthq/forge/ent/usersession"
	"github.com/usegranthq/forge/ent/verification"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	oidcclientFields := schema.OidcClient{}.Fields()
	_ = oidcclientFields
	// oidcclientDescName is the schema descriptor for name field.
	oidcclientDescName := oidcclientFields[1].Descriptor()
	// oidcclient.NameValidator is a validator for the "name" field. It is called by the builders before save.
	oidcclient.NameValidator = oidcclientDescName.Validators[0].(func(string) error)
	// oidcclientDescAudience is the schema descriptor for audience field.
	oidcclientDescAudience := oidcclientFields[2].Descriptor()
	// oidcclient.AudienceValidator is a validator for the "audience" field. It is called by the builders before save.
	oidcclient.AudienceValidator = oidcclientDescAudience.Validators[0].(func(string) error)
	// oidcclientDescClientRefID is the schema descriptor for client_ref_id field.
	oidcclientDescClientRefID := oidcclientFields[3].Descriptor()
	// oidcclient.ClientRefIDValidator is a validator for the "client_ref_id" field. It is called by the builders before save.
	oidcclient.ClientRefIDValidator = oidcclientDescClientRefID.Validators[0].(func(string) error)
	// oidcclientDescClientID is the schema descriptor for client_id field.
	oidcclientDescClientID := oidcclientFields[4].Descriptor()
	// oidcclient.ClientIDValidator is a validator for the "client_id" field. It is called by the builders before save.
	oidcclient.ClientIDValidator = oidcclientDescClientID.Validators[0].(func(string) error)
	// oidcclientDescCreatedAt is the schema descriptor for created_at field.
	oidcclientDescCreatedAt := oidcclientFields[5].Descriptor()
	// oidcclient.DefaultCreatedAt holds the default value on creation for the created_at field.
	oidcclient.DefaultCreatedAt = oidcclientDescCreatedAt.Default.(func() time.Time)
	// oidcclientDescUpdatedAt is the schema descriptor for updated_at field.
	oidcclientDescUpdatedAt := oidcclientFields[6].Descriptor()
	// oidcclient.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	oidcclient.DefaultUpdatedAt = oidcclientDescUpdatedAt.Default.(func() time.Time)
	// oidcclient.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	oidcclient.UpdateDefaultUpdatedAt = oidcclientDescUpdatedAt.UpdateDefault.(func() time.Time)
	// oidcclientDescID is the schema descriptor for id field.
	oidcclientDescID := oidcclientFields[0].Descriptor()
	// oidcclient.DefaultID holds the default value on creation for the id field.
	oidcclient.DefaultID = oidcclientDescID.Default.(func() uuid.UUID)
	projectFields := schema.Project{}.Fields()
	_ = projectFields
	// projectDescName is the schema descriptor for name field.
	projectDescName := projectFields[1].Descriptor()
	// project.NameValidator is a validator for the "name" field. It is called by the builders before save.
	project.NameValidator = projectDescName.Validators[0].(func(string) error)
	// projectDescURLID is the schema descriptor for url_id field.
	projectDescURLID := projectFields[2].Descriptor()
	// project.URLIDValidator is a validator for the "url_id" field. It is called by the builders before save.
	project.URLIDValidator = projectDescURLID.Validators[0].(func(string) error)
	// projectDescCreatedAt is the schema descriptor for created_at field.
	projectDescCreatedAt := projectFields[4].Descriptor()
	// project.DefaultCreatedAt holds the default value on creation for the created_at field.
	project.DefaultCreatedAt = projectDescCreatedAt.Default.(func() time.Time)
	// projectDescUpdatedAt is the schema descriptor for updated_at field.
	projectDescUpdatedAt := projectFields[5].Descriptor()
	// project.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	project.DefaultUpdatedAt = projectDescUpdatedAt.Default.(func() time.Time)
	// project.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	project.UpdateDefaultUpdatedAt = projectDescUpdatedAt.UpdateDefault.(func() time.Time)
	// projectDescID is the schema descriptor for id field.
	projectDescID := projectFields[0].Descriptor()
	// project.DefaultID holds the default value on creation for the id field.
	project.DefaultID = projectDescID.Default.(func() uuid.UUID)
	projectdomainFields := schema.ProjectDomain{}.Fields()
	_ = projectdomainFields
	// projectdomainDescDomain is the schema descriptor for domain field.
	projectdomainDescDomain := projectdomainFields[1].Descriptor()
	// projectdomain.DomainValidator is a validator for the "domain" field. It is called by the builders before save.
	projectdomain.DomainValidator = projectdomainDescDomain.Validators[0].(func(string) error)
	// projectdomainDescVerified is the schema descriptor for verified field.
	projectdomainDescVerified := projectdomainFields[2].Descriptor()
	// projectdomain.DefaultVerified holds the default value on creation for the verified field.
	projectdomain.DefaultVerified = projectdomainDescVerified.Default.(bool)
	// projectdomainDescCreatedAt is the schema descriptor for created_at field.
	projectdomainDescCreatedAt := projectdomainFields[4].Descriptor()
	// projectdomain.DefaultCreatedAt holds the default value on creation for the created_at field.
	projectdomain.DefaultCreatedAt = projectdomainDescCreatedAt.Default.(func() time.Time)
	// projectdomainDescUpdatedAt is the schema descriptor for updated_at field.
	projectdomainDescUpdatedAt := projectdomainFields[5].Descriptor()
	// projectdomain.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	projectdomain.DefaultUpdatedAt = projectdomainDescUpdatedAt.Default.(func() time.Time)
	// projectdomain.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	projectdomain.UpdateDefaultUpdatedAt = projectdomainDescUpdatedAt.UpdateDefault.(func() time.Time)
	// projectdomainDescID is the schema descriptor for id field.
	projectdomainDescID := projectdomainFields[0].Descriptor()
	// projectdomain.DefaultID holds the default value on creation for the id field.
	projectdomain.DefaultID = projectdomainDescID.Default.(func() uuid.UUID)
	tokenFields := schema.Token{}.Fields()
	_ = tokenFields
	// tokenDescName is the schema descriptor for name field.
	tokenDescName := tokenFields[1].Descriptor()
	// token.NameValidator is a validator for the "name" field. It is called by the builders before save.
	token.NameValidator = tokenDescName.Validators[0].(func(string) error)
	// tokenDescToken is the schema descriptor for token field.
	tokenDescToken := tokenFields[2].Descriptor()
	// token.TokenValidator is a validator for the "token" field. It is called by the builders before save.
	token.TokenValidator = tokenDescToken.Validators[0].(func(string) error)
	// tokenDescCreatedAt is the schema descriptor for created_at field.
	tokenDescCreatedAt := tokenFields[5].Descriptor()
	// token.DefaultCreatedAt holds the default value on creation for the created_at field.
	token.DefaultCreatedAt = tokenDescCreatedAt.Default.(func() time.Time)
	// tokenDescUpdatedAt is the schema descriptor for updated_at field.
	tokenDescUpdatedAt := tokenFields[6].Descriptor()
	// token.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	token.DefaultUpdatedAt = tokenDescUpdatedAt.Default.(func() time.Time)
	// token.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	token.UpdateDefaultUpdatedAt = tokenDescUpdatedAt.UpdateDefault.(func() time.Time)
	// tokenDescID is the schema descriptor for id field.
	tokenDescID := tokenFields[0].Descriptor()
	// token.DefaultID holds the default value on creation for the id field.
	token.DefaultID = tokenDescID.Default.(func() uuid.UUID)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescUID is the schema descriptor for uid field.
	userDescUID := userFields[1].Descriptor()
	// user.DefaultUID holds the default value on creation for the uid field.
	user.DefaultUID = userDescUID.Default.(func() uuid.UUID)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[2].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[6].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[7].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
	usersessionFields := schema.UserSession{}.Fields()
	_ = usersessionFields
	// usersessionDescToken is the schema descriptor for token field.
	usersessionDescToken := usersessionFields[1].Descriptor()
	// usersession.TokenValidator is a validator for the "token" field. It is called by the builders before save.
	usersession.TokenValidator = usersessionDescToken.Validators[0].(func(string) error)
	// usersessionDescCreatedAt is the schema descriptor for created_at field.
	usersessionDescCreatedAt := usersessionFields[3].Descriptor()
	// usersession.DefaultCreatedAt holds the default value on creation for the created_at field.
	usersession.DefaultCreatedAt = usersessionDescCreatedAt.Default.(func() time.Time)
	// usersessionDescUpdatedAt is the schema descriptor for updated_at field.
	usersessionDescUpdatedAt := usersessionFields[4].Descriptor()
	// usersession.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	usersession.DefaultUpdatedAt = usersessionDescUpdatedAt.Default.(func() time.Time)
	// usersession.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	usersession.UpdateDefaultUpdatedAt = usersessionDescUpdatedAt.UpdateDefault.(func() time.Time)
	// usersessionDescID is the schema descriptor for id field.
	usersessionDescID := usersessionFields[0].Descriptor()
	// usersession.DefaultID holds the default value on creation for the id field.
	usersession.DefaultID = usersessionDescID.Default.(func() uuid.UUID)
	verificationFields := schema.Verification{}.Fields()
	_ = verificationFields
	// verificationDescAttemptID is the schema descriptor for attempt_id field.
	verificationDescAttemptID := verificationFields[1].Descriptor()
	// verification.DefaultAttemptID holds the default value on creation for the attempt_id field.
	verification.DefaultAttemptID = verificationDescAttemptID.Default.(func() uuid.UUID)
	// verificationDescCode is the schema descriptor for code field.
	verificationDescCode := verificationFields[3].Descriptor()
	// verification.CodeValidator is a validator for the "code" field. It is called by the builders before save.
	verification.CodeValidator = verificationDescCode.Validators[0].(func(string) error)
	// verificationDescAttempts is the schema descriptor for attempts field.
	verificationDescAttempts := verificationFields[4].Descriptor()
	// verification.DefaultAttempts holds the default value on creation for the attempts field.
	verification.DefaultAttempts = verificationDescAttempts.Default.(int)
	// verificationDescCreatedAt is the schema descriptor for created_at field.
	verificationDescCreatedAt := verificationFields[6].Descriptor()
	// verification.DefaultCreatedAt holds the default value on creation for the created_at field.
	verification.DefaultCreatedAt = verificationDescCreatedAt.Default.(func() time.Time)
	// verificationDescUpdatedAt is the schema descriptor for updated_at field.
	verificationDescUpdatedAt := verificationFields[7].Descriptor()
	// verification.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	verification.DefaultUpdatedAt = verificationDescUpdatedAt.Default.(func() time.Time)
	// verification.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	verification.UpdateDefaultUpdatedAt = verificationDescUpdatedAt.UpdateDefault.(func() time.Time)
	// verificationDescID is the schema descriptor for id field.
	verificationDescID := verificationFields[0].Descriptor()
	// verification.DefaultID holds the default value on creation for the id field.
	verification.DefaultID = verificationDescID.Default.(func() uuid.UUID)
}
