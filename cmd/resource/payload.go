package resource

import (
   "fmt"
   "github.com/newrelic/newrelic-cloudformation-resource-providers-common/model"
   log "github.com/sirupsen/logrus"
)

//
// Generic, should be able to leave these as-is
//

type Payload struct {
   model  *Model
   models []interface{}
}

func (p *Payload) SetIdentifier(g *string) {
   p.model.Id = g
}

func (p *Payload) GetIdentifier() *string {
   return p.model.Id
}

func (p *Payload) GetIdentifierKey(a model.Action) string {
   return "id"
}

func (p *Payload) HasTags() bool {
   return p.model.Tags != nil
}

func (p *Payload) GetTags() map[string]string {
   return p.model.Tags
}

func (p *Payload) GetTagIdentifier() *string {
   return p.model.EntityGuid
}

func NewPayload(m *Model) *Payload {
   eg := "I'm a teapot!"
   if m.EntityGuid == nil {
      m.EntityGuid = &eg
   }
   return &Payload{
      model:  m,
      models: make([]interface{}, 0),
   }
}

func (p *Payload) GetResourceModel() interface{} {
   return p.model
}

func (p *Payload) GetResourceModels() []interface{} {
   log.Debugf("GetResourceModels: returning %+v", p.models)
   return p.models
}

func (p *Payload) AppendToResourceModels(m model.Model) {
   p.models = append(p.models, m.GetResourceModel())
}

func (p *Payload) GetTypeName() string {
   return typeName
}

//
// These are API specific, must be configured per API
//

var typeName = "NewRelic::Observability::AlertsNrqlCondition"

func (p *Payload) NewModelFromGuid(g interface{}) (m model.Model) {
   s := fmt.Sprintf("%s", g)
   return NewPayload(&Model{Id: &s})
}

var emptyString = ""

func (p *Payload) GetGraphQLFragment() *string {
   return &emptyString
}

func (p *Payload) GetListQueryNextCursor() string {
   return p.GetListQuery()
}

func (p *Payload) GetVariables() map[string]string {
   vars := make(map[string]string)
   if p.model.Variables != nil {
      for k, v := range p.model.Variables {
         vars[k] = v
      }
   }

   if p.model.Id != nil {
      vars["ID"] = *p.model.Id
   }

   if p.model.PolicyId != nil {
      vars["POLICYID"] = *p.model.PolicyId
   }

   if p.model.Condition != nil {
      vars["CONDITION"] = *p.model.Condition
   }

   if p.model.ListQueryFilter != nil {
      vars["LISTQUERYFILTER"] = *p.model.ListQueryFilter
   }

   if p.model.ConditionType != nil {
      vars["CONDITIONTYPE"] = *p.model.ConditionType
   }
   return vars
}

func (p *Payload) GetErrorKey() string {
   return ""
}

func (p *Payload) GetCreateMutation() string {
   return `
mutation {
    alertsNrqlCondition{{{CONDITIONTYPE}}}Create(
        accountId: {{{ACCOUNTID}}},
        {{{CONDITION}}} ,
        policyId: "{{{POLICYID}}}")
    {
        id
        entityGuid
    }
}
`
}

func (p *Payload) GetDeleteMutation() string {
   return `
mutation {
  alertsConditionDelete(accountId: {{{ACCOUNTID}}}, id: "{{{ID}}}") {
    id
  }
}
`
}

func (p *Payload) GetUpdateMutation() string {
   return `
mutation {
    alertsNrqlCondition{{{CONDITIONTYPE}}}Update(accountId: {{{ACCOUNTID}}}, {{{CONDITION}}} , id: "{{{ID}}}") {
        id
        entityGuid
    }
}
`
}

func (p *Payload) GetReadQuery() string {
   return `
{
  actor {
    account(id: {{{ACCOUNTID}}}) {
      alerts {
        nrqlCondition(id: "{{{ID}}}") {
          id
          entityGuid
        }
      }
    }
  }
}
`
}

func (p *Payload) GetListQuery() string {
   return `
{
  actor {
    account(id: {{{ACCOUNTID}}} ) {
      alerts {
        nrqlConditionsSearch(cursor: "{{{NEXTCURSOR}}}") {
          nextCursor
          nrqlConditions {
            id
          }
        }
      }
    }
  }
}
`
}
