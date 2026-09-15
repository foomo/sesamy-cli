package template

const UtmAttributionCookieWriterData = `___INFO___

{
  "type": "TAG",
  "id": "cvt_temp_public_id",
  "version": 1,
  "securityGroups": [],
  "displayName": "%s",
  "brand": {
    "id": "brand_dummy",
    "displayName": "sesamy",
    "thumbnail": ""
  },
  "description": "Managed by Sesamy. DO NOT EDIT.\nPersists UTM attribution data (current and first-touch) into the shared attribution cookie.",
  "containerContexts": [
    "SERVER"
  ]
}


___TEMPLATE_PARAMETERS___

[]


___SANDBOXED_JS_FOR_SERVER___

const getAllEventData = require('getAllEventData');
const getCookieValues = require('getCookieValues');
const getRequestHeader = require('getRequestHeader');
const JSON = require('JSON');
const makeString = require('makeString');
const parseUrl = require('parseUrl');
const setCookie = require('setCookie');

const COOKIE_NAME = 'utm_attribution';
const UTM_KEYS = ['utm_source', 'utm_medium', 'utm_campaign', 'utm_content', 'utm_term', 'utm_id'];

const cookieOptions = {
  domain:
    '.' +
    (getRequestHeader('x-original-forwarded-host') ||
      getRequestHeader('x-forwarded-host') ||
      getRequestHeader('host')),
  path: '/',
  samesite: 'Lax',
  secure: true,
  'max-age': 31536000, // 1 year
  httpOnly: false
};

const eventData = getAllEventData();
const url = eventData.page_location || getRequestHeader('referer');
const urlParsed = url ? parseUrl(url) : null;

// JSON.parse returns undefined (rather than throwing) on malformed input in
// the sandboxed JS runtime — try/catch is unsupported here.
const cookieRaw = getCookieValues(COOKIE_NAME)[0];
const stored = (cookieRaw && JSON.parse(cookieRaw)) || {};

let changed = false;

UTM_KEYS.forEach((key) => {
  const currentValue = urlParsed && urlParsed.searchParams[key];
  if (!currentValue) return;

  // Session/last-touch value: always refresh to the latest sighting.
  if (stored[key] !== currentValue) {
    stored[key] = currentValue;
    changed = true;
  }

  // First-touch value: set once, never overwritten afterwards.
  const initialKey = 'initial_' + key;
  if (!stored[initialKey]) {
    stored[initialKey] = currentValue;
    changed = true;
  }
});

if (changed) {
  setCookie(COOKIE_NAME, makeString(JSON.stringify(stored)), cookieOptions);
}

data.gtmOnSuccess();


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
        "publicId": "get_cookies",
        "versionId": "1"
      },
      "param": [
        {
          "key": "cookieAccess",
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
        "publicId": "set_cookies",
        "versionId": "1"
      },
      "param": [
        {
          "key": "allowedCookies",
          "value": {
            "type": 2,
            "listItem": [
              {
                "type": 3,
                "mapKey": [
                  {
                    "type": 1,
                    "string": "name"
                  },
                  {
                    "type": 1,
                    "string": "domain"
                  },
                  {
                    "type": 1,
                    "string": "path"
                  },
                  {
                    "type": 1,
                    "string": "secure"
                  },
                  {
                    "type": 1,
                    "string": "session"
                  }
                ],
                "mapValue": [
                  {
                    "type": 1,
                    "string": "*"
                  },
                  {
                    "type": 1,
                    "string": "*"
                  },
                  {
                    "type": 1,
                    "string": "*"
                  },
                  {
                    "type": 1,
                    "string": "any"
                  },
                  {
                    "type": 1,
                    "string": "any"
                  }
                ]
              }
            ]
          }
        }
      ]
    },
    "clientAnnotations": {
      "isEditedByUser": true
    },
    "isRequired": true
  }
]


___TESTS___

scenarios: []


___NOTES___

Code generated by sesamy. DO NOT EDIT.
`
