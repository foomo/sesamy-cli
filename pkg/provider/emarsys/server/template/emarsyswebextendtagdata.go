package template

const EmarsysWebExtendTagData = `___INFO___

{
  "type": "TAG",
  "id": "cvt_temp_public_id",
  "version": 1,
  "securityGroups": [],
  "displayName": "%s",
  "brand": {
    "id": "brand_dummy",
    "displayName": "sesamy",
    "thumbnail": "data:image/jpeg;base64,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qEqaLUwCNxCKZYOe6dFBjPqDoI4rERBIeBK+DcwoAKKtaCJQ+CDXNNpzqBHmCcxwpoi4JJrgyjGuyUw6hlELaeQAclgnHgpfqegeEYRvPCCZjzpIAMAApp7QIXKDg64ETjEUywYBJYAumw8qf4osOWJtRl7cG9iVlHm/IhbCBrmm1OVtwJeMRXJRLxhNIBbVpW3vsRC55EpIAAwDDWJW3uAEmAJNAms8qaYiOYa3K29p8ErIiKQyBXJtclbewbjEVyUUZ0EAmcAoBr8rb1Mx3SRQfE/EwEB9BlbfkgwSaBRY/cs0eR9ElbU5iWRDFn+JoAwvoyOUwaKCYioAIf8ATf8A/9oADAMAAAERAhEAABDzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzglbrzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzilRuyvzzzzzzzzzzzzzzzzzzzzzzzzzzzzywTezzzyvzzzzzzzzzzzzzzzzzzzzzzzzzzywcRbzzzzyvzzzzzzzzzzzzzzzzzzzzzzzzzRDVzzzzzzzyvzzzzzzzzzzzzzzzzzzzzzzgl913zzzzzzzU8fzzzzzzzzzzzzzzzzzzzziUdyzzzzzzzzr1wzzzzzzzzzzzzzzzzzzzywP27zzzzzzzzGO0zzzzzzzzzzzzzzzzzzzgQzZ/zzzzzzznNqwzzzzzzzzzzzzzzzzzzzzh033zzzzzz3PHozzzzzzg4vzzzzzzzzzzzzzyy/zzzzzzzG3wxzzzzzg0zIPzzzzzzzzzzzzzzz/wA88845TNc888884JwSCCD888888888888888/889hL7K08888ymiACCCCD888888888888888/+sqE8ss4GuAHhACCCCCCD888888888888884jG088888utswCCCCCCCCCI888888888888888sY+z+8888+9CCCCCCELLuM88888888888888888sdFiU8888CCCCEJqMc888888888888888888888sc2P188CCFPlcc88888888888888888888888888Mqw4GItO8888882888888888888888888888888sdsM88888oirj88888888888888888888888888888880BqCAq888888888888888888888888888880XSgAAAq888888888888888888888888884rNjACAAAAq888888888888888888888888ompLAAAAAAAEC88888888888888888888882QxiiCAAAAAD+Tc88888888888888888888sJEAAAAAAACbExM8888888888888888888884VAAAAAAAIFU/sc8888888888888888888888oVAAAACABIMM8888888888888888888888888oVAACXbks8888888888888888888888888888oVANEtM8888888888888888888888888888884QePc888888888888888888888888888888888cs8888888888888888888888888888888888888888888888888888888888888888888888888888888888888888888888888//EACMRAQACAgEEAwEBAQAAAAAAAAEAESExQBAgMFBBUXFwYYH/2gAIAQIRAT8Q/hWOW2VHoYC1mKzY9V0dNXPOtVMALiOfZq5oluo7gRO10rs1cveMzG4kRXl6ALRbMSZZjjBLdyu9XG4ZlwBo4xtqYIXGsvANzWMf7MhmwBRxhbdT4T/2OW/ACuJm8CZurZRxtkzMXgR21fgZoJlsIVhx8RtMeYIqtvgdo3PncNo8nT+eDVMTN5QWh4AYSATfgafztBdTIVRM25YAYPAJh9oAdd+Bp/OodJlmOibhagSzvEwkK7d+Bp/IpREcxjQ7EWIpGteo2BPBvwKafiqXfdcGooKAeFBFbfBQ7huY35L7tHfcSRbxUHc1mO79Hcginl/pIiOemjqgj9Iq87exTMINxBH6Rb9Eg+FBufCRLZ9CMI6i301c8SK9urmIIBqL3tXKST4SKd+DMFcgBEdRz4AXUR3B4y6gkV4RM+yAa5D8JIA3Ar0IrBPQgLqI7gnoQTPsgDXoUFQDcAP4N//EACMRAQACAgEDBQEBAAAAAAAAAAEAESExQCAwQRBQUWGhcIH/2gAIAQERAT8Q/hSCKy/YgIj6aOfZBNRT0aOagnhIq9WjlpIviLfY1ckSKeyJg+YAcbEEiPZGwDcMcZQnwxTvsowJ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"
  },
  "description": "Managed by Sesamy. DO NOT EDIT.\nSend events to Emarsys",
  "containerContexts": [
    "SERVER"
  ]
}


___TEMPLATE_PARAMETERS___

[
  {
    "type": "TEXT",
    "name": "merchantId",
    "displayName": "Merchant ID",
    "simpleValueType": true
  },
  {
    "type": "GROUP",
    "name": "consentSettingsGroup",
    "displayName": "Consent Settings",
    "groupStyle": "ZIPPY_CLOSED",
    "subParams": [
      {
        "type": "RADIO",
        "name": "adStorageConsent",
        "displayName": "",
        "radioItems": [
          {
            "value": "optional",
            "displayValue": "Send data always"
          },
          {
            "value": "required",
            "displayValue": "Send data in case marketing consent given"
          }
        ],
        "simpleValueType": true,
        "defaultValue": "optional"
      }
    ]
  },
  {
    "type": "GROUP",
    "name": "debugSettingsGroup",
    "displayName": "Debug Settings",
    "groupStyle": "ZIPPY_CLOSED",
    "subParams": [
      {
        "type": "CHECKBOX",
        "name": "isTestMode",
        "checkboxText": "Test Mode",
        "simpleValueType": true
      },
      {
        "type": "CHECKBOX",
        "name": "isDebugMode",
        "checkboxText": "Debug Mode",
        "simpleValueType": true
      }
    ]
  }
]


___SANDBOXED_JS_FOR_SERVER___

const Math = require('Math');
const JSON = require('JSON');
const parseUrl = require('parseUrl');
const setCookie = require('setCookie');
const sendHttpGet = require('sendHttpGet');
const setResponseBody = require('setResponseBody');
const getRemoteAddress = require('getRemoteAddress');
const encodeUriComponent = require('encodeUriComponent');
const getAllEventData = require('getAllEventData');
const getRequestHeader = require('getRequestHeader');
const getCookieValues = require('getCookieValues');
const generateRandom = require('generateRandom');
const logToConsole = require('logToConsole');
const createRegex = require('createRegex');

// --- Config ---

const eventData = getAllEventData();
const merchantUrl = 'https://recommender.scarabresearch.com/merchants/'+data.merchantId+'/';

// --- Consent ---

if (!isConsentGivenOrNotRequired()) {
  return data.gtmOnSuccess();
}

// --- Main ---

const mappedData = mapEventData();
const cookieList = ["s", "cdv", "xp", "fc"];
const headerList = ["referer", "user-agent"];
const requestUrl = merchantUrl+'?'+serializeData(mappedData);
const requestOptions = {
  headers: generateRequestHeaders(headerList, cookieList),
  timeout: 500,
};

return sendHttpGet(requestUrl, requestOptions).then((result) => {
  if (result.statusCode >= 200 && result.statusCode < 300) {
    if (result.headers['set-cookie']) {
      setResponseCookies(result.headers['set-cookie']);
    }
    data.gtmOnSuccess();
  } else {
    logToConsole('[FAILURE]', {
      request: mappedData,
      eventData: eventData,
    });
    data.gtmOnFailure();
  }
});

// --- Utils ---

function mapEventData() {
  const mappedData = {
    email: eventData.emarsys_email || null,
    customerId: eventData.user_id || null,
    sessionId: getCookieValues('emarsys_s')[0] || eventData.ga_session_id,
    pageViewId: eventData.emarsys_page_view_id || generatePageViewId(),
    isNewPageView: !eventData.emarsys_page_view_id,
    visitorId: getCookieValues('emarsys_cdv')[0] || eventData.client_id,
    referrer: eventData.page_referrer || null,
    orderId: null,
    order: null,
    search: null,
    category: null,
    view: null,
    cart: null,
  };

  switch (eventData.event_name) {
    case 'page_view': {
      mappedData.cart = serializeItems(eventData.items || []);
      mappedData.search = ((parseUrl(eventData.page_location) || {}).searchParams || {}).q || null;
      break;
    }
    case 'view_item': {
      mappedData.cart = serializeItems(eventData.items || []);
      mappedData.view = serializeItem(eventData.items[0] || {}, false);
      break;
    }
    case 'view_item_list': {
      const prefix = createRegex('^/', 'i');
      const seperator = createRegex('/', 'i');
      mappedData.cart = serializeItems(eventData.items || []);
      mappedData.category = eventData.item_list_id.replace(prefix,"").replace(seperator, " > ");
      break;
    }
    case 'purchase': {
      mappedData.cart = [];
      mappedData.orderId = eventData.transaction_id;
      mappedData.order = serializeItems(eventData.items || []);
      if (eventData.tax) {
        mappedData.order[0].price += eventData.tax;
      }
      if (eventData.shipping) {
        mappedData.order[0].price += eventData.shipping;
      }
      break;
    }
  }
  return mappedData;
}

function serializeItems(items) {
  const ret = [];
  items.forEach((item) => {
    ret.push(serializeItem(item, true));
  });
  return ret.join('|');
}

function serializeItem(item, full) {
  const ret = [];
  if (item.item_id) {
    ret.push('i:'+item.item_id);
  }
  if (full && item.price) {
    ret.push('p:'+item.price);
  }
  if (full && item.quantity) {
    ret.push('q:'+item.quantity);
  }
  return ret.join(',');
}

function serializeData(mappedData) {
  const slist = [];

  slist.push("cp=1");
  slist.push("pv=" + mappedData.pageViewId);

  if (mappedData.isNewPageView) {
    slist.push("xp=1");
  }
  if (mappedData.email) {
    slist.push("eh=" + encodeUriComponent(mappedData.email));
  } else if (mappedData.customerId) {
    slist.push("ci=" + encodeUriComponent(mappedData.customerId));
  }
  if (mappedData.sessionId) {
    slist.push("s=" + encodeUriComponent(mappedData.sessionId));
  }
  if (mappedData.visitorId) {
    slist.push("vi=" + encodeUriComponent(mappedData.visitorId));
  }
  if (mappedData.category) {
    slist.push("vc=" + encodeUriComponent(mappedData.category));
  }
  if (mappedData.view) {
    slist.push("v=" + encodeUriComponent(mappedData.view));
  }
  if (mappedData.orderId) {
    slist.push("oi=" + encodeUriComponent(mappedData.orderId));
  }
  if (mappedData.order) {
    slist.push("co=" + encodeUriComponent(mappedData.order));
  }
  if (mappedData.cart !== null) {
    slist.push("ca=" + encodeUriComponent(mappedData.cart));
    slist.push("cv=1");
  }
  if (mappedData.search) {
    slist.push("q=" + encodeUriComponent(mappedData.search));
  }
  if (mappedData.referrer) {
    slist.push("prev_url=" + encodeUriComponent(mappedData.referrer));
  }
  if (data.isTestMode) {
    slist.push("test=true");
  }
  if (data.isDebugMode) {
    slist.push("debug=true");
  }
  logToConsole(JSON.stringify(slist));
  return slist.join("&");
}

function generatePageViewId() {
  return generateRandom(0, Math.pow(2, 31));
}

function isConsentGivenOrNotRequired() {
  if (data.adStorageConsent !== 'required') {
    return true;
  }
  if (eventData.consent_state) {
    return !!eventData.consent_state.ad_storage;
  }
  const xGaGcs = eventData['x-ga-gcs'] || ''; // x-ga-gcs is a string like "G110"
  return xGaGcs[2] === '1';
}

function setResponseCookies(cookieList) {
  for (let i = 0; i < cookieList.length; i++) {
    let cookieArray = cookieList[i].split("; ").map((pair) => pair.split("="));
    let cookieJSON = "";

    for (let j = 1; j < cookieArray.length; j++) {
      if (j === 1) cookieJSON += "{";
      if (cookieArray[j].length > 1) cookieJSON += '"' + cookieArray[j][0] + '": "' + cookieArray[j][1] + '"';
      else cookieJSON += '"' + cookieArray[j][0] + '": ' + true;
      if (j + 1 < cookieArray.length) cookieJSON += ",";
      else cookieJSON += "}";
    }

    setCookie(cookieArray[0][0], cookieArray[0][1], JSON.parse(cookieJSON));
  }
}

function generateRequestHeaders(headerList, cookieList) {
  let headers = {};
  let cookies = [];

  for (let i = 0; i < headerList.length; i++) {
    let headerName = headerList[i];
    let headerValue = getRequestHeader(headerName);
    if (headerValue) {
      headers[headerName] = getRequestHeader(headerName);
    }
  }

  headers.cookie = "";

  for (let i = 0; i < cookieList.length; i++) {
    let cookieName = cookieList[i];
    let cookieValue = getCookieValues(cookieName);
    if (cookieValue && cookieValue.length) {
      cookies.push(cookieName + "=" + cookieValue[0]);
    }
  }

  headers.cookie = cookies.join("; ");
  headers["X-Forwarded-For"] = getRemoteAddress();

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
        "publicId": "access_response",
        "versionId": "1"
      },
      "param": [
        {
          "key": "writeResponseAccess",
          "value": {
            "type": 1,
            "string": "any"
          }
        },
        {
          "key": "writeHeaderAccess",
          "value": {
            "type": 1,
            "string": "specific"
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
