const sendHttpGet = require('sendHttpGet');
const getEventData = require('getEventData');
const generateRandom = require('generateRandom');
const logToConsole = require('logToConsole');
const Math = require('Math');
const encodeUriComponent = require('encodeUriComponent');
const JSON = require('JSON');
const getCookieValues = require('getCookieValues');
const setCookie = require('setCookie');
const parseUrl = require('parseUrl');


let clickid;
function isNotEmpty(obj) {
	if (obj === undefined || obj === null || obj.toString() === "" ) {
		return false;
	} else {
		return obj;
	}
}

let msquery = () => {
	const url = parseUrl(getEventData('page_location'));
	if (url && url.searchParams.msclkid) {
			return url.searchParams.msclkid;
	} else {
		return false;
	}
};

function setUETCookie(clickid) {
	setCookie('_uetmsclkid', "_uet"+clickid, {"expires":7776000000, httpOnly: false});
	logToConsole("Cookie set: " + clickid);
}

if(data.first_click) {
	if (isNotEmpty(data.msclkidCookie) || isNotEmpty(getCookieValues('_uetmsclkid'))) {
		let clickid_cookie = getCookieValues('_uetmsclkid')[0];
		clickid = (isNotEmpty(data.msclkidCookie) || (clickid_cookie.substring(4)));
		clickid += "-0";
	} else { //cookie not present
		if (isNotEmpty(data.msclkidQuery) || msquery()) {
			clickid = (isNotEmpty(data.msclkidQuery) || msquery());
			setUETCookie(clickid);
			clickid += "-1";
		} else {
			clickid = "N";
		}
	}
} else {
	if (isNotEmpty(data.msclkidQuery) || msquery()) {
		clickid = (isNotEmpty(data.msclkidQuery) || msquery());
		setUETCookie(clickid);
		clickid += "-1";
	} else {
		if(isNotEmpty(data.msclkidCookie) || isNotEmpty(getCookieValues('_uetmsclkid'))){
			let clickid_cookie = getCookieValues('_uetmsclkid')[0];
			clickid = (isNotEmpty(data.msclkidCookie) || (clickid_cookie.substring(4)));
			clickid += "-0";
		} else {
			clickid = "N";
		}
	}
}


const screen = data.screen || getEventData('screen_resolution');

let width, height = "";
if (isNotEmpty(screen)){
	width = screen.split('x')[0];
	height = screen.split('x')[1];

}
let rn = generateRandom(100000, 999999);
function s4() {
	return Math.floor(((1 + (generateRandom(1, 9999999)/10000000))) * 65536)
	.toString(16)
	.substring(1);
}

let mid = (s4() + s4() + "-" + s4() + "-" + s4() + "-" + s4() + "-" + s4() + s4() + s4());
let items_data;
if(isNotEmpty(getEventData('items'))){
	items_data = getEventData('items');
}

let items = () => {
	let result = "";
	if(isNotEmpty(data.prodid)){
		items_data = data.prodid;
	}
	if (data.itemsGa) {
		if(items_data) {
			items_data.forEach(function(item, i) {
				result += "id=" + item.item_id + 'quantity=' + item.quantity + 'price=' + item.price;
				if (i < items_data.length - 1) {
					result += ',';
				}
			});
		}
	} else {
		result = items_data;
	}
	return result;
};


let items_id = () => {
	let result = "";
	if(isNotEmpty(data.items)){
		items_data = data.items;
	}
	if(data.prodidGa) {
		if (items_data) {
			items_data.forEach(function(item, i) {
				result += item.item_id;
				if (i < items_data.length - 1) {
					result += ',';
				}
			});
		}
	}
	return result;
};

let spa = () => {
	if(data.spa) {
		return "Y";
	} else {
		return "N";
	}
};

let params = {
	rn: rn,
	ti: data.ti,
	ver: '2.3',
	mid: mid,
	uid: isNotEmpty(data.userId) || getEventData('user_id'),
	evt: data.evt,
	p: getEventData('page_location'),
	r: getEventData('page_referrer'),
	tl: isNotEmpty(data.pageTitle) || getEventData('page_title'),
	pagetype: data.pagetype,
	items:  items(),
	prodid: items_id(),
	search_term: isNotEmpty(data.searchTerm) || getEventData('search_term'),
	transaction_id: isNotEmpty(data.transactionId) || getEventData('transaction_id'),
	lg: isNotEmpty(data.lg) || getEventData('language'),
	sw: width,
	sh: height,
	sc: data.sc || getEventData('screen_color_depth'),
	spa: spa(),
	msclkid: clickid,
	sid: isNotEmpty(data.sid) || getEventData('uet_session_id'),
	vid: isNotEmpty(data.vid) || getEventData('vid'),
	page_path: isNotEmpty(data.pagePath) || getEventData('page_path'),
	gc: getEventData('currency'),
	gv: getEventData('value'),
	ec: isNotEmpty(data.ec) || getEventData('event_category'),
	ea: isNotEmpty(data.ea) || getEventData('event_action'),
	el: isNotEmpty(data.el) || getEventData('event_label'),
	ev: isNotEmpty(data.ev) || getEventData('event_value'),
};


if (data.activateLogs) {
	logToConsole("Params: " + JSON.stringify(params));
}


let url = 'https://bat.bing.com/action/0?';
let all_params = "";
if(isNotEmpty(params)){
	for (var key in params){
		if (params[key] == undefined || params[key] == null) {
			continue;
		}
		all_params += key + "=" + encodeUriComponent(params[key]) + "&";
	}
	url += all_params;
}
if (data.activateLogs) {
	logToConsole('URL: ' + url);
}


return sendHttpGet(url, {
	headers: {key: 'value'},
	timeout: 500,
}).then((result) => {
	if (result.statusCode >= 200 && result.statusCode < 300) {
		logToConsole('Result: ' + data.gtmOnSuccess());
		data.gtmOnSuccess();
	} else {
		logToConsole('Error: ' + result.statusCode);
		data.gtmOnFailure();
	}
});
