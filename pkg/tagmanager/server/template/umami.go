package template

import (
	"fmt"

	"google.golang.org/api/tagmanager/v2"
)

const DataUmami = `___INFO___

{
  "type": "TAG",
  "id": "cvt_temp_public_id",
  "version": 1,
  "securityGroups": [],
  "displayName": "%s",
  "brand": {
    "id": "brand_dummy",
    "displayName": "Umami",
    "thumbnail": "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAACAAAAAgCAYAAABzenr0AAAABmJLR0QAvgD/ALe9KR/IAAAACXBIWXMAAC4jAAAuIwF4pT92AAAAB3RJTUUH5gwHAwsAtVzgFQAAABl0RVh0Q29tbWVudABDcmVhdGVkIHdpdGggR0lNUFeBDhcAAAJmSURBVFjDxVe/SzJxHH5OckgagrsImms4bTm1XEIQhXCqTXBzkLYiKtDFBA/xH2gJQRwP1PPHKuiW3HA4SEMgurlcNRy5HZ936LX3Nd/X1Mx74LMc3Pd5vg98Pt/nwxiGQTARa/P+MBgMoCgKnp6e8Pz8DABgWRZ7e3s4PDzEzs7OfAcahkGzVDabJa/XSwA+ymazkc1mG/t2dHRE9/f3NOu5XwqQZZn29/cJADmdTkqn0/Tw8EC6rtMIuq5Tq9WiTCZDbrebABDP81QoFL4nIBaLEQASBIHK5TLNilqt9iHk+vp6MQHRaJQA0MXFBS2Kq6srAkCRSGQ+AfF4nACQKIr0XWQymalOTAioVCoEgM7Pz2lZGDlRLBYnBDCf54AgCGAYBu12e6n9fnBwgLe3N3Q6nf+3YS6XIwBUKpVo2ajVagRgokXHHAgEAnh5eVn67f92YX19Hc1m888kzOfz6Pf7GA6HaDQa8Pl8SCaTPyJgc3MT9XodNzc32NjYgMPhAPN7gpkGC0yG+QI4jjONfGtrCxaWZU0TYLfbzXXAbrfDEgwGTRPgcrkAVVXHAsUq6/Hx8f0x8ng8KyfneZ4MwyALAJydna3c/nA4PP4Y+f3+ld7+9fV1PA+McsAq6u7u7t+BRBTFHye/vb2dnogSicSPkYdCodkyoSRJSyc/PT2dLxVLkkTb29tLIU+lUovtBd1uly4vL8lqtS5EHI1GSVGUqXsBM8tyqqoqZFlGtVqdDJWf4Ha7cXJyguPj4/dR+wWYebfjTqeDXq8HTdOgaRoAgOM4cByH3d1d8Dw/10BizF7PfwHUW1xlBL8DDgAAAABJRU5ErkJggg\u003d\u003d"
  },
  "description": "Send events to Umami",
  "containerContexts": [
    "SERVER"
  ]
}


___TEMPLATE_PARAMETERS___

[
  {
    "type": "GROUP",
    "name": "settings",
    "displayName": "Umami Settings",
    "groupStyle": "NO_ZIPPY",
    "subParams": [
      {
        "type": "TEXT",
        "name": "endpointUrl",
        "displayName": "Endpoint URL",
        "simpleValueType": true,
        "help": "Enter your endpoint URL for your own instance or use https://analytics.umami.is/api/send for Umami Cloud",
        "defaultValue": "https://site-gtm-umami-umami/api/send",
        "valueValidators": [
          {
            "type": "NON_EMPTY"
          }
        ]
      },
      {
        "type": "TEXT",
        "name": "websiteId",
        "displayName": "Website ID",
        "simpleValueType": true,
        "valueValidators": [
          {
            "type": "NON_EMPTY"
          }
        ],
        "help": "Paste ID for your website from the Umami settings"
      },
      {
        "type": "TEXT",
        "name": "domain",
        "displayName": "Host / Domain (Optional)",
        "simpleValueType": true,
        "help": "Enter an optional fixed domain to override event data"
      },
      {
        "type": "TEXT",
        "name": "timeout",
        "displayName": "Timeout (ms)",
        "simpleValueType": true,
        "valueValidators": [
          {
            "type": "NON_EMPTY"
          },
          {
            "type": "POSITIVE_NUMBER"
          }
        ],
        "defaultValue": 1000
      }
    ]
  }
]


___SANDBOXED_JS_FOR_SERVER___

const JSON = require('JSON');
const parseUrl = require('parseUrl');
const makeString = require('makeString');
const getAllEventData = require('getAllEventData');
const sendHttpRequest = require('sendHttpRequest');
const getRequestHeader = require('getRequestHeader');
const logToConsole = require('logToConsole');

const traceId = getRequestHeader('trace-id');
const eventData = getAllEventData();
const pageLocation = eventData.page_location;

if (pageLocation) {
  const name = eventData.event_name || "";
  const parsedUrl = parseUrl(pageLocation);
  const serviceUrl = data.endpointUrl;
  const ref = eventData.page_referrer || "";
  const hostname = data.domain || parsedUrl.hostname || null;

  // https://github.com/umami-software/umami/blob/7a75639dd3d7aeff46104b71ebfb3853fc0eee09/src/tracker/index.d.ts
  let umamiEvent = {
    type: "event",
    payload: {
      name: (name === "page_view") ? "" : name,
      website: data.websiteId,
      hostname: hostname,
      title: eventData.page_title || "",
      url: pageLocation.split(parsedUrl.hostname)[1],
      referrer: ref,
      language: eventData.language || "",
      screen: eventData.screen_resolution || "1920x1080"
    }
  };

  logToConsole(
    JSON.stringify({
      Name: 'Umami',
      Type: 'Request',
      TraceId: traceId,
      EventName: makeString(umamiEvent.payload.name),
      RequestMethod: 'POST',
      RequestUrl: serviceUrl,
      RequestBody: umamiEvent,
      EventData: eventData,
    })
  );

  const headers = {
    'user-agent': eventData.user_agent || getRequestHeader("user-agent"),
    'content-type': 'application/json'
  };

  sendHttpRequest(
      serviceUrl, (statusCode, headers, body) => {
        logToConsole(
          JSON.stringify({
            Name: 'Umami',
            Type: 'Response',
            TraceId: traceId,
            EventName: makeString(umamiEvent.payload.name),
            ResponseStatusCode: statusCode,
            ResponseHeaders: headers,
            ResponseBody: body,
          })
        );
        if (statusCode >= 200 && statusCode < 300) {
          data.gtmOnSuccess();
        } else {
          data.gtmOnFailure();
        }
      },
      {
        headers: addRequestHeaders(headers),
        method: 'POST',
        timeout: data.timeout||1000
      },
      JSON.stringify(umamiEvent)
  );
} else {
  data.gtmOnFailure();
}

function addRequestHeaders(headers) {
  const keys = [
    'cf-ipcity',
    'cf-ipcountry',
    'cf-ipcontinent',
    'cf-iplongitude',
    'cf-iplatitude',
    'cf-region',
    'cf-region-code',
    'cf-metro-code',
    'cf-postal-code',
    'cf-timezone',
  ];
  for (let i = 0; i < keys.length; i++) {
    const key = keys[i];
    const value = getRequestHeader(key);
    if (value) {
      headers[key] = value;
    }
  }
  return headers;
}


___SERVER_PERMISSIONS___

[
  {
    "instance": {
      "key": {
        "publicId": "read_request",
        "versionId": "1"
      },
      "param": [
        {
          "key": "requestAccess",
          "value": {
            "type": 1,
            "string": "any"
          }
        },
        {
          "key": "headerAccess",
          "value": {
            "type": 1,
            "string": "any"
          }
        },
        {
          "key": "queryParameterAccess",
          "value": {
            "type": 1,
            "string": "any"
          }
        }
      ]
    },
    "isRequired": true
  },
  {
    "instance": {
      "key": {
        "publicId": "logging",
        "versionId": "1"
      },
      "param": [
        {
          "key": "environments",
          "value": {
            "type": 1,
            "string": "debug"
          }
        }
      ]
    },
    "isRequired": true
  },
  {
    "instance": {
      "key": {
        "publicId": "read_container_data",
        "versionId": "1"
      },
      "param": []
    },
    "isRequired": false
  },
  {
    "instance": {
      "key": {
        "publicId": "read_event_data",
        "versionId": "1"
      },
      "param": [
        {
          "key": "eventDataAccess",
          "value": {
            "type": 1,
            "string": "any"
          }
        }
      ]
    },
    "isRequired": true
  },
  {
    "instance": {
      "key": {
        "publicId": "send_http",
        "versionId": "1"
      },
      "param": [
        {
          "key": "allowedUrls",
          "value": {
            "type": 1,
            "string": "any"
          }
        }
      ]
    },
    "isRequired": true
  }
]


___TESTS___

scenarios: []


___NOTES___

Code generated by sesamy. DO NOT EDIT.
`

func NewUmami(name string) *tagmanager.CustomTemplate {
	return &tagmanager.CustomTemplate{
		Name:         name,
		TemplateData: fmt.Sprintf(DataUmami, name),
	}
}
