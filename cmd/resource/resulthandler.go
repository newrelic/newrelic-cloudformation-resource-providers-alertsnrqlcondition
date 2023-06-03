package resource

import (
   "fmt"
   "github.com/newrelic/newrelic-cloudformation-resource-providers-common/cferror"
   "github.com/newrelic/newrelic-cloudformation-resource-providers-common/client/nerdgraph"
   "github.com/newrelic/newrelic-cloudformation-resource-providers-common/model"
   log "github.com/sirupsen/logrus"
)

// ResultHandler at a minimum provides access to the default error processing.
// If required we can provide custom processing here via composition overrides https://go.dev/doc/effective_go#embedding
type ResultHandler struct {
   // Use Go composition to access the default implementation
   model.ResultHandler
}

func NewResultHandler() (h model.ResultHandler) {
   defer func() {
      log.Debugf("(tagging) errorHandler.NewErrorHandler: exit %p", h)
   }()
   // Initialize ourself with the common core
   h = &ResultHandler{ResultHandler: nerdgraph.NewResultHandler()}
   return
}

func (h *ResultHandler) Create(m model.Model, b []byte) (err error) {
   defer log.Debug("resultHandler.Create (alertsnrqlcondition): enter")
   log.Debug("resultHandler.Create (alertsnrqlcondition): enter")
   // id
   key := "id"
   var v interface{}
   v, err = nerdgraph.FindKeyValue(b, key)
   if err != nil {
      log.Errorf("Create: error finding result key: %s in response: %s", key, string(b))
      return err
   }
   id := fmt.Sprintf("%v", v)
   lm, ok := m.GetResourceModel().(*Model)
   if !ok {
      errMsg := fmt.Sprintf("Unable to cast m.ResourceModel() to *Model: %T", m.GetResourceModel())
      log.Error(errMsg)
      err = fmt.Errorf("%w %s", &cferror.ServiceInternalError{}, errMsg)
      return
   }
   log.Debugf("resultHandler.Create (alertsnrqlcondition): setting Id to %s", id)
   lm.Id = &id

   // guid
   key = "entityGuid"
   v, err = nerdgraph.FindKeyValue(b, key)
   if err != nil {
      log.Errorf("Create: error finding result key: %s in response: %s", key, string(b))
      return err
   }
   guid := fmt.Sprintf("%v", v)
   lm, ok = m.GetResourceModel().(*Model)
   if !ok {
      errMsg := fmt.Sprintf("Unable to cast m.ResourceModel() to *Model: %T", m.GetResourceModel())
      log.Error(errMsg)
      err = fmt.Errorf("%w %s", &cferror.ServiceInternalError{}, errMsg)
   }
   lm.EntityGuid = &guid
   log.Debugf("resultHandler.Create (alertsnrqlcondition): setting Guid to %s", guid)
   log.Debugf("resultHandler.Create (alertsnrqlcondition): model: %+v", lm)

   return
}

func (h *ResultHandler) Update(m model.Model, b []byte) (err error) {
   defer log.Debug("resultHandler.Update (alertsnrqlcondition): enter")
   log.Debug("resultHandler.Update (alertsnrqlcondition): enter")
   // id
   key := "id"
   var v interface{}
   v, err = nerdgraph.FindKeyValue(b, key)
   if err != nil {
      log.Errorf("Update: error finding result key: %s in response: %s", key, string(b))
      return err
   }
   id := fmt.Sprintf("%v", v)
   lm, ok := m.GetResourceModel().(*Model)
   if !ok {
      errMsg := fmt.Sprintf("Unable to cast m.ResourceModel() to *Model: %T", m.GetResourceModel())
      log.Error(errMsg)
      err = fmt.Errorf("%w %s", &cferror.ServiceInternalError{}, errMsg)
      return
   }
   log.Debugf("resultHandler.Update (alertsnrqlcondition): setting Id to %s", id)
   lm.Id = &id

   // guid
   key = "entityGuid"
   v, err = nerdgraph.FindKeyValue(b, key)
   if err != nil {
      log.Errorf("Update: error finding result key: %s in response: %s", key, string(b))
      return err
   }
   guid := fmt.Sprintf("%v", v)
   lm, ok = m.GetResourceModel().(*Model)
   if !ok {
      errMsg := fmt.Sprintf("Unable to cast m.ResourceModel() to *Model: %T", m.GetResourceModel())
      log.Error(errMsg)
      err = fmt.Errorf("%w %s", &cferror.ServiceInternalError{}, errMsg)
   }
   lm.EntityGuid = &guid
   log.Debugf("resultHandler.Update (alertsnrqlcondition): setting Guid to %s", guid)
   log.Debugf("resultHandler.Update (alertsnrqlcondition): model: %+v", lm)

   return
}

func (h *ResultHandler) Read(m model.Model, b []byte) (err error) {
   defer log.Debug("resultHandler.Read (alertsnrqlcondition): enter")
   log.Debug("resultHandler.Read (alertsnrqlcondition): enter")
   // id
   key := "id"
   var v interface{}
   v, err = nerdgraph.FindKeyValue(b, key)
   if err != nil {
      log.Errorf("Read: error finding result key: %s in response: %s", key, string(b))
      return err
   }
   id := fmt.Sprintf("%v", v)
   lm, ok := m.GetResourceModel().(*Model)
   if !ok {
      errMsg := fmt.Sprintf("Unable to cast m.ResourceModel() to *Model: %T", m.GetResourceModel())
      log.Error(errMsg)
      err = fmt.Errorf("%w %s", &cferror.ServiceInternalError{}, errMsg)
      return
   }
   log.Debugf("resultHandler.Read (alertsnrqlcondition): setting Id to %s", id)
   lm.Id = &id

   // guid
   key = "entityGuid"
   v, err = nerdgraph.FindKeyValue(b, key)
   if err != nil {
      log.Errorf("Read: error finding result key: %s in response: %s", key, string(b))
      return err
   }
   guid := fmt.Sprintf("%v", v)
   lm, ok = m.GetResourceModel().(*Model)
   if !ok {
      errMsg := fmt.Sprintf("Unable to cast m.ResourceModel() to *Model: %T", m.GetResourceModel())
      log.Error(errMsg)
      err = fmt.Errorf("%w %s", &cferror.ServiceInternalError{}, errMsg)
   }
   lm.EntityGuid = &guid
   log.Debugf("resultHandler.Read (alertsnrqlcondition): setting Guid to %s", guid)
   log.Debugf("resultHandler.Read (alertsnrqlcondition): model: %+v", lm)

   return
}
