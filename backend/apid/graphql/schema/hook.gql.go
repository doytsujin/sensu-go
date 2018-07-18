// Code generated by scripts/gengraphql.go. DO NOT EDIT.

package schema

import (
	errors "errors"
	graphql1 "github.com/graphql-go/graphql"
	graphql "github.com/sensu/sensu-go/graphql"
)

// HookConfigIDFieldResolver implement to resolve requests for the HookConfig's id field.
type HookConfigIDFieldResolver interface {
	// ID implements response to request for id field.
	ID(p graphql.ResolveParams) (string, error)
}

// HookConfigNamespaceFieldResolver implement to resolve requests for the HookConfig's namespace field.
type HookConfigNamespaceFieldResolver interface {
	// Namespace implements response to request for namespace field.
	Namespace(p graphql.ResolveParams) (interface{}, error)
}

// HookConfigNameFieldResolver implement to resolve requests for the HookConfig's name field.
type HookConfigNameFieldResolver interface {
	// Name implements response to request for name field.
	Name(p graphql.ResolveParams) (string, error)
}

// HookConfigCommandFieldResolver implement to resolve requests for the HookConfig's command field.
type HookConfigCommandFieldResolver interface {
	// Command implements response to request for command field.
	Command(p graphql.ResolveParams) (string, error)
}

// HookConfigTimeoutFieldResolver implement to resolve requests for the HookConfig's timeout field.
type HookConfigTimeoutFieldResolver interface {
	// Timeout implements response to request for timeout field.
	Timeout(p graphql.ResolveParams) (int, error)
}

// HookConfigStdinFieldResolver implement to resolve requests for the HookConfig's stdin field.
type HookConfigStdinFieldResolver interface {
	// Stdin implements response to request for stdin field.
	Stdin(p graphql.ResolveParams) (bool, error)
}

//
// HookConfigFieldResolvers represents a collection of methods whose products represent the
// response values of the 'HookConfig' type.
//
// == Example SDL
//
//   """
//   Dog's are not hooman.
//   """
//   type Dog implements Pet {
//     "name of this fine beast."
//     name:  String!
//
//     "breed of this silly animal; probably shibe."
//     breed: [Breed]
//   }
//
// == Example generated interface
//
//   // DogResolver ...
//   type DogFieldResolvers interface {
//     DogNameFieldResolver
//     DogBreedFieldResolver
//
//     // IsTypeOf is used to determine if a given value is associated with the Dog type
//     IsTypeOf(interface{}, graphql.IsTypeOfParams) bool
//   }
//
// == Example implementation ...
//
//   // DogResolver implements DogFieldResolvers interface
//   type DogResolver struct {
//     logger logrus.LogEntry
//     store interface{
//       store.BreedStore
//       store.DogStore
//     }
//   }
//
//   // Name implements response to request for name field.
//   func (r *DogResolver) Name(p graphql.ResolveParams) (interface{}, error) {
//     // ... implementation details ...
//     dog := p.Source.(DogGetter)
//     return dog.GetName()
//   }
//
//   // Breed implements response to request for breed field.
//   func (r *DogResolver) Breed(p graphql.ResolveParams) (interface{}, error) {
//     // ... implementation details ...
//     dog := p.Source.(DogGetter)
//     breed := r.store.GetBreed(dog.GetBreedName())
//     return breed
//   }
//
//   // IsTypeOf is used to determine if a given value is associated with the Dog type
//   func (r *DogResolver) IsTypeOf(p graphql.IsTypeOfParams) bool {
//     // ... implementation details ...
//     _, ok := p.Value.(DogGetter)
//     return ok
//   }
//
type HookConfigFieldResolvers interface {
	HookConfigIDFieldResolver
	HookConfigNamespaceFieldResolver
	HookConfigNameFieldResolver
	HookConfigCommandFieldResolver
	HookConfigTimeoutFieldResolver
	HookConfigStdinFieldResolver
}

// HookConfigAliases implements all methods on HookConfigFieldResolvers interface by using reflection to
// match name of field to a field on the given value. Intent is reduce friction
// of writing new resolvers by removing all the instances where you would simply
// have the resolvers method return a field.
//
// == Example SDL
//
//    type Dog {
//      name:   String!
//      weight: Float!
//      dob:    DateTime
//      breed:  [Breed]
//    }
//
// == Example generated aliases
//
//   type DogAliases struct {}
//   func (_ DogAliases) Name(p graphql.ResolveParams) (interface{}, error) {
//     // reflect...
//   }
//   func (_ DogAliases) Weight(p graphql.ResolveParams) (interface{}, error) {
//     // reflect...
//   }
//   func (_ DogAliases) Dob(p graphql.ResolveParams) (interface{}, error) {
//     // reflect...
//   }
//   func (_ DogAliases) Breed(p graphql.ResolveParams) (interface{}, error) {
//     // reflect...
//   }
//
// == Example Implementation
//
//   type DogResolver struct { // Implements DogResolver
//     DogAliases
//     store store.BreedStore
//   }
//
//   // NOTE:
//   // All other fields are satisified by DogAliases but since this one
//   // requires hitting the store we implement it in our resolver.
//   func (r *DogResolver) Breed(p graphql.ResolveParams) interface{} {
//     dog := v.(*Dog)
//     return r.BreedsById(dog.BreedIDs)
//   }
//
type HookConfigAliases struct{}

// ID implements response to request for 'id' field.
func (_ HookConfigAliases) ID(p graphql.ResolveParams) (string, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	ret, ok := val.(string)
	if err != nil {
		return ret, err
	}
	if !ok {
		return ret, errors.New("unable to coerce value for field 'id'")
	}
	return ret, err
}

// Namespace implements response to request for 'namespace' field.
func (_ HookConfigAliases) Namespace(p graphql.ResolveParams) (interface{}, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	return val, err
}

// Name implements response to request for 'name' field.
func (_ HookConfigAliases) Name(p graphql.ResolveParams) (string, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	ret, ok := val.(string)
	if err != nil {
		return ret, err
	}
	if !ok {
		return ret, errors.New("unable to coerce value for field 'name'")
	}
	return ret, err
}

// Command implements response to request for 'command' field.
func (_ HookConfigAliases) Command(p graphql.ResolveParams) (string, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	ret, ok := val.(string)
	if err != nil {
		return ret, err
	}
	if !ok {
		return ret, errors.New("unable to coerce value for field 'command'")
	}
	return ret, err
}

// Timeout implements response to request for 'timeout' field.
func (_ HookConfigAliases) Timeout(p graphql.ResolveParams) (int, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	ret, ok := graphql1.Int.ParseValue(val).(int)
	if err != nil {
		return ret, err
	}
	if !ok {
		return ret, errors.New("unable to coerce value for field 'timeout'")
	}
	return ret, err
}

// Stdin implements response to request for 'stdin' field.
func (_ HookConfigAliases) Stdin(p graphql.ResolveParams) (bool, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	ret, ok := val.(bool)
	if err != nil {
		return ret, err
	}
	if !ok {
		return ret, errors.New("unable to coerce value for field 'stdin'")
	}
	return ret, err
}

// HookConfigType HookConfig is the specification of a hook
var HookConfigType = graphql.NewType("HookConfig", graphql.ObjectKind)

// RegisterHookConfig registers HookConfig object type with given service.
func RegisterHookConfig(svc *graphql.Service, impl HookConfigFieldResolvers) {
	svc.RegisterObject(_ObjectTypeHookConfigDesc, impl)
}
func _ObjTypeHookConfigIDHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(HookConfigIDFieldResolver)
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.ID(frp)
	}
}

func _ObjTypeHookConfigNamespaceHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(HookConfigNamespaceFieldResolver)
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.Namespace(frp)
	}
}

func _ObjTypeHookConfigNameHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(HookConfigNameFieldResolver)
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.Name(frp)
	}
}

func _ObjTypeHookConfigCommandHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(HookConfigCommandFieldResolver)
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.Command(frp)
	}
}

func _ObjTypeHookConfigTimeoutHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(HookConfigTimeoutFieldResolver)
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.Timeout(frp)
	}
}

func _ObjTypeHookConfigStdinHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(HookConfigStdinFieldResolver)
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.Stdin(frp)
	}
}

func _ObjectTypeHookConfigConfigFn() graphql1.ObjectConfig {
	return graphql1.ObjectConfig{
		Description: "HookConfig is the specification of a hook",
		Fields: graphql1.Fields{
			"command": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "Command is the command to be executed",
				Name:              "command",
				Type:              graphql1.String,
			},
			"id": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "The globally unique identifier of the record",
				Name:              "id",
				Type:              graphql1.NewNonNull(graphql1.ID),
			},
			"name": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "Name is the unique identifier for a hook",
				Name:              "name",
				Type:              graphql1.NewNonNull(graphql1.String),
			},
			"namespace": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "Namespace in which this record resides",
				Name:              "namespace",
				Type:              graphql1.NewNonNull(graphql.OutputType("Namespace")),
			},
			"stdin": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "Stdin indicates if hook requests have stdin enabled",
				Name:              "stdin",
				Type:              graphql1.NewNonNull(graphql1.Boolean),
			},
			"timeout": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "Timeout is the timeout, in seconds, at which the hook has to run",
				Name:              "timeout",
				Type:              graphql1.Int,
			},
		},
		Interfaces: []*graphql1.Interface{
			graphql.Interface("Node")},
		IsTypeOf: func(_ graphql1.IsTypeOfParams) bool {
			// NOTE:
			// Panic by default. Intent is that when Service is invoked, values of
			// these fields are updated with instantiated resolvers. If these
			// defaults are called it is most certainly programmer err.
			// If you're see this comment then: 'Whoops! Sorry, my bad.'
			panic("Unimplemented; see HookConfigFieldResolvers.")
		},
		Name: "HookConfig",
	}
}

// describe HookConfig's configuration; kept private to avoid unintentional tampering of configuration at runtime.
var _ObjectTypeHookConfigDesc = graphql.ObjectDesc{
	Config: _ObjectTypeHookConfigConfigFn,
	FieldHandlers: map[string]graphql.FieldHandler{
		"command":   _ObjTypeHookConfigCommandHandler,
		"id":        _ObjTypeHookConfigIDHandler,
		"name":      _ObjTypeHookConfigNameHandler,
		"namespace": _ObjTypeHookConfigNamespaceHandler,
		"stdin":     _ObjTypeHookConfigStdinHandler,
		"timeout":   _ObjTypeHookConfigTimeoutHandler,
	},
}

// HookConfigFieldResolver implement to resolve requests for the Hook's config field.
type HookConfigFieldResolver interface {
	// Config implements response to request for config field.
	Config(p graphql.ResolveParams) (interface{}, error)
}

// HookDurationFieldResolver implement to resolve requests for the Hook's duration field.
type HookDurationFieldResolver interface {
	// Duration implements response to request for duration field.
	Duration(p graphql.ResolveParams) (int, error)
}

// HookExecutedFieldResolver implement to resolve requests for the Hook's executed field.
type HookExecutedFieldResolver interface {
	// Executed implements response to request for executed field.
	Executed(p graphql.ResolveParams) (int, error)
}

// HookIssuedFieldResolver implement to resolve requests for the Hook's issued field.
type HookIssuedFieldResolver interface {
	// Issued implements response to request for issued field.
	Issued(p graphql.ResolveParams) (int, error)
}

// HookOutputFieldResolver implement to resolve requests for the Hook's output field.
type HookOutputFieldResolver interface {
	// Output implements response to request for output field.
	Output(p graphql.ResolveParams) (string, error)
}

// HookStatusFieldResolver implement to resolve requests for the Hook's status field.
type HookStatusFieldResolver interface {
	// Status implements response to request for status field.
	Status(p graphql.ResolveParams) (int, error)
}

//
// HookFieldResolvers represents a collection of methods whose products represent the
// response values of the 'Hook' type.
//
// == Example SDL
//
//   """
//   Dog's are not hooman.
//   """
//   type Dog implements Pet {
//     "name of this fine beast."
//     name:  String!
//
//     "breed of this silly animal; probably shibe."
//     breed: [Breed]
//   }
//
// == Example generated interface
//
//   // DogResolver ...
//   type DogFieldResolvers interface {
//     DogNameFieldResolver
//     DogBreedFieldResolver
//
//     // IsTypeOf is used to determine if a given value is associated with the Dog type
//     IsTypeOf(interface{}, graphql.IsTypeOfParams) bool
//   }
//
// == Example implementation ...
//
//   // DogResolver implements DogFieldResolvers interface
//   type DogResolver struct {
//     logger logrus.LogEntry
//     store interface{
//       store.BreedStore
//       store.DogStore
//     }
//   }
//
//   // Name implements response to request for name field.
//   func (r *DogResolver) Name(p graphql.ResolveParams) (interface{}, error) {
//     // ... implementation details ...
//     dog := p.Source.(DogGetter)
//     return dog.GetName()
//   }
//
//   // Breed implements response to request for breed field.
//   func (r *DogResolver) Breed(p graphql.ResolveParams) (interface{}, error) {
//     // ... implementation details ...
//     dog := p.Source.(DogGetter)
//     breed := r.store.GetBreed(dog.GetBreedName())
//     return breed
//   }
//
//   // IsTypeOf is used to determine if a given value is associated with the Dog type
//   func (r *DogResolver) IsTypeOf(p graphql.IsTypeOfParams) bool {
//     // ... implementation details ...
//     _, ok := p.Value.(DogGetter)
//     return ok
//   }
//
type HookFieldResolvers interface {
	HookConfigFieldResolver
	HookDurationFieldResolver
	HookExecutedFieldResolver
	HookIssuedFieldResolver
	HookOutputFieldResolver
	HookStatusFieldResolver
}

// HookAliases implements all methods on HookFieldResolvers interface by using reflection to
// match name of field to a field on the given value. Intent is reduce friction
// of writing new resolvers by removing all the instances where you would simply
// have the resolvers method return a field.
//
// == Example SDL
//
//    type Dog {
//      name:   String!
//      weight: Float!
//      dob:    DateTime
//      breed:  [Breed]
//    }
//
// == Example generated aliases
//
//   type DogAliases struct {}
//   func (_ DogAliases) Name(p graphql.ResolveParams) (interface{}, error) {
//     // reflect...
//   }
//   func (_ DogAliases) Weight(p graphql.ResolveParams) (interface{}, error) {
//     // reflect...
//   }
//   func (_ DogAliases) Dob(p graphql.ResolveParams) (interface{}, error) {
//     // reflect...
//   }
//   func (_ DogAliases) Breed(p graphql.ResolveParams) (interface{}, error) {
//     // reflect...
//   }
//
// == Example Implementation
//
//   type DogResolver struct { // Implements DogResolver
//     DogAliases
//     store store.BreedStore
//   }
//
//   // NOTE:
//   // All other fields are satisified by DogAliases but since this one
//   // requires hitting the store we implement it in our resolver.
//   func (r *DogResolver) Breed(p graphql.ResolveParams) interface{} {
//     dog := v.(*Dog)
//     return r.BreedsById(dog.BreedIDs)
//   }
//
type HookAliases struct{}

// Config implements response to request for 'config' field.
func (_ HookAliases) Config(p graphql.ResolveParams) (interface{}, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	return val, err
}

// Duration implements response to request for 'duration' field.
func (_ HookAliases) Duration(p graphql.ResolveParams) (int, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	ret, ok := graphql1.Int.ParseValue(val).(int)
	if err != nil {
		return ret, err
	}
	if !ok {
		return ret, errors.New("unable to coerce value for field 'duration'")
	}
	return ret, err
}

// Executed implements response to request for 'executed' field.
func (_ HookAliases) Executed(p graphql.ResolveParams) (int, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	ret, ok := graphql1.Int.ParseValue(val).(int)
	if err != nil {
		return ret, err
	}
	if !ok {
		return ret, errors.New("unable to coerce value for field 'executed'")
	}
	return ret, err
}

// Issued implements response to request for 'issued' field.
func (_ HookAliases) Issued(p graphql.ResolveParams) (int, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	ret, ok := graphql1.Int.ParseValue(val).(int)
	if err != nil {
		return ret, err
	}
	if !ok {
		return ret, errors.New("unable to coerce value for field 'issued'")
	}
	return ret, err
}

// Output implements response to request for 'output' field.
func (_ HookAliases) Output(p graphql.ResolveParams) (string, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	ret, ok := val.(string)
	if err != nil {
		return ret, err
	}
	if !ok {
		return ret, errors.New("unable to coerce value for field 'output'")
	}
	return ret, err
}

// Status implements response to request for 'status' field.
func (_ HookAliases) Status(p graphql.ResolveParams) (int, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	ret, ok := graphql1.Int.ParseValue(val).(int)
	if err != nil {
		return ret, err
	}
	if !ok {
		return ret, errors.New("unable to coerce value for field 'status'")
	}
	return ret, err
}

/*
HookType A Hook is a hook specification and optionally the results of the hook's
execution.
*/
var HookType = graphql.NewType("Hook", graphql.ObjectKind)

// RegisterHook registers Hook object type with given service.
func RegisterHook(svc *graphql.Service, impl HookFieldResolvers) {
	svc.RegisterObject(_ObjectTypeHookDesc, impl)
}
func _ObjTypeHookConfigHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(HookConfigFieldResolver)
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.Config(frp)
	}
}

func _ObjTypeHookDurationHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(HookDurationFieldResolver)
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.Duration(frp)
	}
}

func _ObjTypeHookExecutedHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(HookExecutedFieldResolver)
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.Executed(frp)
	}
}

func _ObjTypeHookIssuedHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(HookIssuedFieldResolver)
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.Issued(frp)
	}
}

func _ObjTypeHookOutputHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(HookOutputFieldResolver)
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.Output(frp)
	}
}

func _ObjTypeHookStatusHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(HookStatusFieldResolver)
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.Status(frp)
	}
}

func _ObjectTypeHookConfigFn() graphql1.ObjectConfig {
	return graphql1.ObjectConfig{
		Description: "A Hook is a hook specification and optionally the results of the hook's\nexecution.",
		Fields: graphql1.Fields{
			"config": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "Config is the specification of a hook",
				Name:              "config",
				Type:              graphql.OutputType("HookConfig"),
			},
			"duration": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "Duration of execution",
				Name:              "duration",
				Type:              graphql1.Int,
			},
			"executed": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "Executed describes the time in which the hook request was executed",
				Name:              "executed",
				Type:              graphql1.Int,
			},
			"issued": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "Issued describes the time in which the hook request was issued",
				Name:              "issued",
				Type:              graphql1.Int,
			},
			"output": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "Output from the execution of Command",
				Name:              "output",
				Type:              graphql1.String,
			},
			"status": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "Status is the exit status code produced by the hook",
				Name:              "status",
				Type:              graphql1.Int,
			},
		},
		Interfaces: []*graphql1.Interface{},
		IsTypeOf: func(_ graphql1.IsTypeOfParams) bool {
			// NOTE:
			// Panic by default. Intent is that when Service is invoked, values of
			// these fields are updated with instantiated resolvers. If these
			// defaults are called it is most certainly programmer err.
			// If you're see this comment then: 'Whoops! Sorry, my bad.'
			panic("Unimplemented; see HookFieldResolvers.")
		},
		Name: "Hook",
	}
}

// describe Hook's configuration; kept private to avoid unintentional tampering of configuration at runtime.
var _ObjectTypeHookDesc = graphql.ObjectDesc{
	Config: _ObjectTypeHookConfigFn,
	FieldHandlers: map[string]graphql.FieldHandler{
		"config":   _ObjTypeHookConfigHandler,
		"duration": _ObjTypeHookDurationHandler,
		"executed": _ObjTypeHookExecutedHandler,
		"issued":   _ObjTypeHookIssuedHandler,
		"output":   _ObjTypeHookOutputHandler,
		"status":   _ObjTypeHookStatusHandler,
	},
}

// HookListHooksFieldResolver implement to resolve requests for the HookList's hooks field.
type HookListHooksFieldResolver interface {
	// Hooks implements response to request for hooks field.
	Hooks(p graphql.ResolveParams) ([]string, error)
}

// HookListTypeFieldResolver implement to resolve requests for the HookList's type field.
type HookListTypeFieldResolver interface {
	// Type implements response to request for type field.
	Type(p graphql.ResolveParams) (string, error)
}

//
// HookListFieldResolvers represents a collection of methods whose products represent the
// response values of the 'HookList' type.
//
// == Example SDL
//
//   """
//   Dog's are not hooman.
//   """
//   type Dog implements Pet {
//     "name of this fine beast."
//     name:  String!
//
//     "breed of this silly animal; probably shibe."
//     breed: [Breed]
//   }
//
// == Example generated interface
//
//   // DogResolver ...
//   type DogFieldResolvers interface {
//     DogNameFieldResolver
//     DogBreedFieldResolver
//
//     // IsTypeOf is used to determine if a given value is associated with the Dog type
//     IsTypeOf(interface{}, graphql.IsTypeOfParams) bool
//   }
//
// == Example implementation ...
//
//   // DogResolver implements DogFieldResolvers interface
//   type DogResolver struct {
//     logger logrus.LogEntry
//     store interface{
//       store.BreedStore
//       store.DogStore
//     }
//   }
//
//   // Name implements response to request for name field.
//   func (r *DogResolver) Name(p graphql.ResolveParams) (interface{}, error) {
//     // ... implementation details ...
//     dog := p.Source.(DogGetter)
//     return dog.GetName()
//   }
//
//   // Breed implements response to request for breed field.
//   func (r *DogResolver) Breed(p graphql.ResolveParams) (interface{}, error) {
//     // ... implementation details ...
//     dog := p.Source.(DogGetter)
//     breed := r.store.GetBreed(dog.GetBreedName())
//     return breed
//   }
//
//   // IsTypeOf is used to determine if a given value is associated with the Dog type
//   func (r *DogResolver) IsTypeOf(p graphql.IsTypeOfParams) bool {
//     // ... implementation details ...
//     _, ok := p.Value.(DogGetter)
//     return ok
//   }
//
type HookListFieldResolvers interface {
	HookListHooksFieldResolver
	HookListTypeFieldResolver
}

// HookListAliases implements all methods on HookListFieldResolvers interface by using reflection to
// match name of field to a field on the given value. Intent is reduce friction
// of writing new resolvers by removing all the instances where you would simply
// have the resolvers method return a field.
//
// == Example SDL
//
//    type Dog {
//      name:   String!
//      weight: Float!
//      dob:    DateTime
//      breed:  [Breed]
//    }
//
// == Example generated aliases
//
//   type DogAliases struct {}
//   func (_ DogAliases) Name(p graphql.ResolveParams) (interface{}, error) {
//     // reflect...
//   }
//   func (_ DogAliases) Weight(p graphql.ResolveParams) (interface{}, error) {
//     // reflect...
//   }
//   func (_ DogAliases) Dob(p graphql.ResolveParams) (interface{}, error) {
//     // reflect...
//   }
//   func (_ DogAliases) Breed(p graphql.ResolveParams) (interface{}, error) {
//     // reflect...
//   }
//
// == Example Implementation
//
//   type DogResolver struct { // Implements DogResolver
//     DogAliases
//     store store.BreedStore
//   }
//
//   // NOTE:
//   // All other fields are satisified by DogAliases but since this one
//   // requires hitting the store we implement it in our resolver.
//   func (r *DogResolver) Breed(p graphql.ResolveParams) interface{} {
//     dog := v.(*Dog)
//     return r.BreedsById(dog.BreedIDs)
//   }
//
type HookListAliases struct{}

// Hooks implements response to request for 'hooks' field.
func (_ HookListAliases) Hooks(p graphql.ResolveParams) ([]string, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	ret, ok := val.([]string)
	if err != nil {
		return ret, err
	}
	if !ok {
		return ret, errors.New("unable to coerce value for field 'hooks'")
	}
	return ret, err
}

// Type implements response to request for 'type' field.
func (_ HookListAliases) Type(p graphql.ResolveParams) (string, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	ret, ok := val.(string)
	if err != nil {
		return ret, err
	}
	if !ok {
		return ret, errors.New("unable to coerce value for field 'type'")
	}
	return ret, err
}

// HookListType self descriptive
var HookListType = graphql.NewType("HookList", graphql.ObjectKind)

// RegisterHookList registers HookList object type with given service.
func RegisterHookList(svc *graphql.Service, impl HookListFieldResolvers) {
	svc.RegisterObject(_ObjectTypeHookListDesc, impl)
}
func _ObjTypeHookListHooksHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(HookListHooksFieldResolver)
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.Hooks(frp)
	}
}

func _ObjTypeHookListTypeHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(HookListTypeFieldResolver)
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.Type(frp)
	}
}

func _ObjectTypeHookListConfigFn() graphql1.ObjectConfig {
	return graphql1.ObjectConfig{
		Description: "self descriptive",
		Fields: graphql1.Fields{
			"hooks": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "Hooks is the list of hooks for the check hook",
				Name:              "hooks",
				Type:              graphql1.NewNonNull(graphql1.NewList(graphql1.NewNonNull(graphql1.String))),
			},
			"type": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "Type indicates the type or response code for the check hook",
				Name:              "type",
				Type:              graphql1.NewNonNull(graphql1.String),
			},
		},
		Interfaces: []*graphql1.Interface{},
		IsTypeOf: func(_ graphql1.IsTypeOfParams) bool {
			// NOTE:
			// Panic by default. Intent is that when Service is invoked, values of
			// these fields are updated with instantiated resolvers. If these
			// defaults are called it is most certainly programmer err.
			// If you're see this comment then: 'Whoops! Sorry, my bad.'
			panic("Unimplemented; see HookListFieldResolvers.")
		},
		Name: "HookList",
	}
}

// describe HookList's configuration; kept private to avoid unintentional tampering of configuration at runtime.
var _ObjectTypeHookListDesc = graphql.ObjectDesc{
	Config: _ObjectTypeHookListConfigFn,
	FieldHandlers: map[string]graphql.FieldHandler{
		"hooks": _ObjTypeHookListHooksHandler,
		"type":  _ObjTypeHookListTypeHandler,
	},
}
