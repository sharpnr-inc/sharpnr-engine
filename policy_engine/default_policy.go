package policyengine

var Default_ORG_Policy = `{
  "ORGANIZATION_ID":"{{org_Id}}",
  "SHARPNR_ID":"{{sharpnr_id}}",

  "ROLES": ["SUPER_ADMIN", "ADMIN", "MANAGER"],

  "MODULES": [""],

  "PERMISSION_GROUP":["SUPER_ADMIN"],

  "ACCOUNTING_MODULE_PERMISSION":[],

  "EMPLOYEE_PERMISSION":[]
}`
