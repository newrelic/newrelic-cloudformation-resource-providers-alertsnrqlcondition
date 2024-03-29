package resource

import (
   "fmt"
   "github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
   "github.com/davecgh/go-spew/spew"
   "github.com/newrelic/newrelic-cloudformation-resource-providers-common/cferror"
   "github.com/newrelic/newrelic-cloudformation-resource-providers-common/client"
   "github.com/newrelic/newrelic-cloudformation-resource-providers-common/logging"
   "github.com/newrelic/newrelic-cloudformation-resource-providers-common/model"
   log "github.com/sirupsen/logrus"
   "os"
   "runtime/debug"
)

//
// Generic below here, shouldn't require changes
//

func wrap(f func(client *client.GraphqlClient, model model.Model) (event handler.ProgressEvent, err error), req handler.Request, currentModel *Model, prevModel *Model) (event handler.ProgressEvent, err error) {
   defer func() {
      r := recover()
      if r == nil {
         return
      }
      var panicError, originalError string
      stack := string(debug.Stack())

      if err != nil {
         originalError = err.Error()
      }

      rerr, ok := r.(error)
      if ok {
         panicError = rerr.Error()
      } else {
         panicError = fmt.Sprint(r)
      }
      err = fmt.Errorf("%w originalError: %s panicError: %s Stack trace: %s", &cferror.ServiceInternalError{}, originalError, panicError, stack)
   }()
   debug.SetPanicOnFault(true)
   logging.Setup()

   fmt.Println("")
   logging.Dump(log.TraceLevel, os.Environ(), "os.Environ: ")
   logging.Dump(log.TraceLevel, req.RequestContext, "req.RequestContext: ")

   sm := NewPayload(currentModel)
   eh := NewErrorHandler(sm)
   c := client.NewGraphqlClient(req, &typeName, eh, NewResultHandler())

   fmt.Println("")
   event, err = f(c, sm)
   spew.Dump(event)
   return
}

func create(client *client.GraphqlClient, model model.Model) (event handler.ProgressEvent, err error) {
   return client.CreateMutation(model)
}
func _delete(client *client.GraphqlClient, model model.Model) (event handler.ProgressEvent, err error) {
   return client.DeleteMutation(model)
}
func update(client *client.GraphqlClient, model model.Model) (event handler.ProgressEvent, err error) {
   return client.UpdateMutation(model)
}
func read(client *client.GraphqlClient, model model.Model) (event handler.ProgressEvent, err error) {
   return client.ReadQuery(model)
}
func list(client *client.GraphqlClient, model model.Model) (event handler.ProgressEvent, err error) {
   log.Debugf("resource.list")
   return client.ListQuery(model)
}

func Create(req handler.Request, prevModel *Model, currentModel *Model) (event handler.ProgressEvent, err error) {
   return wrap(create, req, currentModel, prevModel)
}
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
   return wrap(update, req, currentModel, prevModel)
}
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
   return wrap(_delete, req, currentModel, prevModel)
}
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
   return wrap(read, req, currentModel, prevModel)
}
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
   log.Debugf("resource.List")
   return wrap(list, req, currentModel, prevModel)
}
