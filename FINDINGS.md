# Findings

Things that I have found that happen during a phase either of Build or Deployment.

## Build

N/A

## Deployment

You shouldn't use the `workflow_dispatch` from tag. You should use `workflow_dispatch` from a branch and take an input for the version that you want to deploy. If you run from a tag version, then the pipelines have to be in working condition in that version.

They will not use the latest version of the pipeline and just pull that release package.

## Static Web App

Looping in some other investigation here around the use of static web apps and the deployment method we use, it's features and what can be expanded on or improved.

### APIs

APIs are available as part of the Static Web App; they are linked to the app either defined with your frontend code - or brought and linked with an existing Azure Function App, or Container App.

* [Bring Your Own Functions](https://learn.microsoft.com/en-us/azure/static-web-apps/functions-bring-your-own)

Using the APIs in the way would allow the ability to cut down on defined patterns for APIM and, reduce the number of components required to get a new App started.

* No need to define API schema (could still be able for developer experience)
* Can still define and package client proxy
* No need to define APIM deployment
  * No need to have extra users/prodcuts/operations

APIM Operations would be culled and defined for clear operations that are likely shared between services; ie: Contact Service fronted through APIM since this app can be called by any service, where-as the BFF fronted services no longer require these operations.

#### Replacing Controls

**Access Restriction:**

Some aspects within this would mean we would have to replace some of the controls that we offer/use in APIM policies; ie: access restrictions from app gateway (defined in Raci-Common-Variables APIM fragment)

The replacement control would be to use the staticwebapp.config.json file to allow certain IPs inbound, such as already defined with the `AzureFrontDoor.Backend` control that already exists

Example:

```json
{
  "routes": [
    {
      "route": "/status",
      "rewrite": "/status.html"
    }
  ],
  "navigationFallback": {
    "rewrite": "/index.html",
    "exclude": []
  },
  "networking": {
    "allowedIpRanges": [
      "10.24.220.0/24", // etc... has to be CIDR notation
      "AzureFrontDoor.Backend"
    ]
  },
  "forwardingGateway": {
    "requiredHeaders": {
      "X-Azure-FDID": "PLACEHOLDER_FRONTDOOR_ID"
    },
    "allowedForwardedHosts": ["PLACEHOLDER_FRONTDOOR_HOSTNAME"]
  }
}
```

**CORS:**

No longer need a CORS defined block to allow the calling web app access into APIM/BFF since the Function App is now linked directly to the Web App. The Web App has exclusive access to the Function App unless otherwise defined - this would likely be preferred for the Spark BFF patterns.

**Rate Limiting:**

Rate limiting would not be defined by the Frontend and instead defined on a shared service; for example the Motor Web APIM policy would still define rate limiting.

Per case basis for each app would need to work out if Rate Limiting is required for an endpoint. In which case it may be better suited to some shared service or defined in the Function App itself - permitting this was investigated and would work appropriately.
