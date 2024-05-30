package generator_test

import (
	"io"
	"log/slog"
	"testing"

	"github.com/foomo/gocontemplate/pkg/contemplate"
	"github.com/foomo/sesamy-cli/pkg/typescript/generator"
	_ "github.com/foomo/sesamy-go/pkg/event/params" // force inclusion
	_ "github.com/foomo/sesamy-go/pkg/sesamy"       // force inclusion
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	expectedEvent = `
// Code generated by sesamy. DO NOT EDIT.
import type * as github_com_foomo_sesamy_go_pkg_event_params from './github_com_foomo_sesamy_go_pkg_event_params';
import { EventName } from './github_com_foomo_sesamy_go_pkg_sesamy';
import { collect } from '@foomo/sesamy';

export const addToCart = (
	params: github_com_foomo_sesamy_go_pkg_event_params.AddToCart<github_com_foomo_sesamy_go_pkg_event_params.Item>,
) => {
	collect({
		name: EventName.AddToCart,
		params,
	});
}

export const pageView = (
	params: github_com_foomo_sesamy_go_pkg_event_params.PageView,
) => {
	collect({
		name: EventName.PageView,
		params,
	});
}
`
	expectedParams = `
// Code generated by sesamy. DO NOT EDIT.
import type * as github_com_foomo_gostandards_iso4217 from './github_com_foomo_gostandards_iso4217';

export interface AddToCart<I> {
	currency?: github_com_foomo_gostandards_iso4217.Currency;
	value?: number;
	items?: Array<I>;
}

export interface Item {
	affiliation?: string;
	coupon?: string;
	creative_name?: string;
	creative_slot?: string;
	discount?: number;
	index?: number;
	item_brand?: string;
	item_category?: string;
	item_category2?: string;
	item_category3?: string;
	item_category4?: string;
	item_category5?: string;
	item_id?: string;
	item_list_id?: string;
	item_list_name?: string;
	item_name?: string;
	item_variant?: string;
	location_id?: string;
	price?: number;
	promotion_id?: string;
	promotion_name?: string;
	quantity?: number;
}

export interface PageView {
	page_title?: string;
	page_location?: string;
}
`
	expectedSesamy = `
// Code generated by sesamy. DO NOT EDIT.

export enum EventName {
	AdImpression = "ad_impression",
	AddPaymentInfo = "add_payment_info",
	AddShippingInfo = "add_shipping_info",
	AddToCart = "add_to_cart",
	AddToWishlist = "add_to_wishlist",
	BeginCheckout = "begin_checkout",
	CampaignDetails = "campaign_details",
	Click = "click",
	EarnVirtualMoney = "earn_virtual_money",
	FileDownload = "file_download",
	FormStart = "form_start",
	FormSubmit = "form_submit",
	GenerateLead = "generate_lead",
	JoinGroup = "join_group",
	LevelEnd = "level_end",
	LevelStart = "level_start",
	LevelUp = "level_up",
	Login = "login",
	PageView = "page_view",
	PostScore = "post_score",
	Purchase = "purchase",
	Refund = "refund",
	RemoveFromCart = "remove_from_cart",
	ScreenView = "screen_view",
	Scroll = "scroll",
	Search = "search",
	SelectContent = "select_content",
	SelectItem = "select_item",
	SelectPromotion = "select_promotion",
	SessionStart = "session_start",
	Share = "share",
	SignUp = "sign_up",
	SpendVirtualCurrency = "spend_virtual_currency",
	TutorialBegin = "tutorial_begin",
	TutorialComplete = "tutorial_complete",
	UnlockAchievement = "unlock_achievement",
	UserEngagement = "user_engagement",
	VideoComplete = "video_complete",
	VideoProgress = "video_progress",
	VideoStart = "video_start",
	ViewCart = "view_cart",
	ViewItem = "view_item",
	ViewItemList = "view_item_list",
	ViewPromotion = "view_promotion",
	ViewSearchResults = "view_search_results",
}

export interface Event<P> {
	name: EventName;
	params?: P;
}
`
	expectedISO4217 = `
// Code generated by sesamy. DO NOT EDIT.

export enum Currency {
	AED = "AED",
	AFN = "AFN",
	ALL = "ALL",
	AMD = "AMD",
	ANG = "ANG",
	AOA = "AOA",
	ARS = "ARS",
	AUD = "AUD",
	AWG = "AWG",
	AZN = "AZN",
	BAM = "BAM",
	BBD = "BBD",
	BDT = "BDT",
	BGN = "BGN",
	BHD = "BHD",
	BIF = "BIF",
	BMD = "BMD",
	BND = "BND",
	BOB = "BOB",
	BOV = "BOV",
	BRL = "BRL",
	BSD = "BSD",
	BTN = "BTN",
	BWP = "BWP",
	BYN = "BYN",
	BZD = "BZD",
	CAD = "CAD",
	CDF = "CDF",
	CHE = "CHE",
	CHF = "CHF",
	CHW = "CHW",
	CLF = "CLF",
	CLP = "CLP",
	CNY = "CNY",
	COP = "COP",
	COU = "COU",
	CRC = "CRC",
	CUP = "CUP",
	CVE = "CVE",
	CZK = "CZK",
	DJF = "DJF",
	DKK = "DKK",
	DOP = "DOP",
	DZD = "DZD",
	EGP = "EGP",
	ERN = "ERN",
	ETB = "ETB",
	EUR = "EUR",
	FJD = "FJD",
	FKP = "FKP",
	GBP = "GBP",
	GEL = "GEL",
	GHS = "GHS",
	GIP = "GIP",
	GMD = "GMD",
	GNF = "GNF",
	GTQ = "GTQ",
	GYD = "GYD",
	HKD = "HKD",
	HNL = "HNL",
	HTG = "HTG",
	HUF = "HUF",
	IDR = "IDR",
	ILS = "ILS",
	INR = "INR",
	IQD = "IQD",
	IRR = "IRR",
	ISK = "ISK",
	JMD = "JMD",
	JOD = "JOD",
	JPY = "JPY",
	KES = "KES",
	KGS = "KGS",
	KHR = "KHR",
	KMF = "KMF",
	KPW = "KPW",
	KRW = "KRW",
	KWD = "KWD",
	KYD = "KYD",
	KZT = "KZT",
	LAK = "LAK",
	LBP = "LBP",
	LKR = "LKR",
	LRD = "LRD",
	LSL = "LSL",
	LYD = "LYD",
	MAD = "MAD",
	MDL = "MDL",
	MGA = "MGA",
	MKD = "MKD",
	MMK = "MMK",
	MNT = "MNT",
	MOP = "MOP",
	MRU = "MRU",
	MUR = "MUR",
	MVR = "MVR",
	MWK = "MWK",
	MXN = "MXN",
	MXV = "MXV",
	MYR = "MYR",
	MZN = "MZN",
	NAD = "NAD",
	NGN = "NGN",
	NIO = "NIO",
	NOK = "NOK",
	NPR = "NPR",
	NZD = "NZD",
	OMR = "OMR",
	PAB = "PAB",
	PEN = "PEN",
	PGK = "PGK",
	PHP = "PHP",
	PKR = "PKR",
	PLN = "PLN",
	PYG = "PYG",
	QAR = "QAR",
	RON = "RON",
	RSD = "RSD",
	RUB = "RUB",
	RWF = "RWF",
	SAR = "SAR",
	SBD = "SBD",
	SCR = "SCR",
	SDG = "SDG",
	SEK = "SEK",
	SGD = "SGD",
	SHP = "SHP",
	SLE = "SLE",
	SLL = "SLL",
	SOS = "SOS",
	SRD = "SRD",
	SSP = "SSP",
	STN = "STN",
	SVC = "SVC",
	SYP = "SYP",
	SZL = "SZL",
	THB = "THB",
	TJS = "TJS",
	TMT = "TMT",
	TND = "TND",
	TOP = "TOP",
	TRY = "TRY",
	TTD = "TTD",
	TWD = "TWD",
	TZS = "TZS",
	UAH = "UAH",
	UGX = "UGX",
	USD = "USD",
	USN = "USN",
	UYI = "UYI",
	UYU = "UYU",
	UYW = "UYW",
	UZS = "UZS",
	VED = "VED",
	VES = "VES",
	VND = "VND",
	VUV = "VUV",
	WST = "WST",
	XAF = "XAF",
	XAG = "XAG",
	XAU = "XAU",
	XBA = "XBA",
	XBB = "XBB",
	XBC = "XBC",
	XBD = "XBD",
	XCD = "XCD",
	XDR = "XDR",
	XOF = "XOF",
	XPD = "XPD",
	XPF = "XPF",
	XPT = "XPT",
	XSU = "XSU",
	XTS = "XTS",
	XUA = "XUA",
	XXX = "XXX",
	YER = "YER",
	ZAR = "ZAR",
	ZMW = "ZMW",
	ZWL = "ZWL",
}
`
)

func TestGenerate(t *testing.T) {
	cfg := &contemplate.Config{
		Packages: []*contemplate.PackageConfig{
			{
				Path:  "github.com/foomo/sesamy-go/pkg/event",
				Types: []string{"AddToCart", "PageView"},
			},
		},
	}
	ctpl, err := contemplate.Load(cfg)
	require.NoError(t, err)

	l := slog.New(slog.NewTextHandler(io.Discard, nil))
	actual, err := generator.Generate(l, ctpl)
	require.NoError(t, err)

	require.Len(t, actual, 4)

	t.Run("event.ts", func(t *testing.T) {
		if !assert.Equal(t, expectedEvent, "\n"+actual["github_com_foomo_sesamy_go_pkg_event.ts"].String()) {
			t.Log("\n" + actual["github_com_foomo_sesamy_go_pkg_event.ts"].String())
		}
	})

	t.Run("sesamy.ts", func(t *testing.T) {
		if !assert.Equal(t, expectedSesamy, "\n"+actual["github_com_foomo_sesamy_go_pkg_sesamy.ts"].String()) {
			t.Log("\n" + actual["github_com_foomo_sesamy_go_pkg_sesamy.ts"].String())
		}
	})

	t.Run("params.ts", func(t *testing.T) {
		if !assert.Equal(t, expectedParams, "\n"+actual["github_com_foomo_sesamy_go_pkg_event_params.ts"].String()) {
			t.Log("\n" + actual["github_com_foomo_sesamy_go_pkg_event_params.ts"].String())
		}
	})

	t.Run("iso4217.ts", func(t *testing.T) {
		if !assert.Equal(t, expectedISO4217, "\n"+actual["github_com_foomo_gostandards_iso4217.ts"].String()) {
			t.Log("\n" + actual["github_com_foomo_gostandards_iso4217.ts"].String())
		}
	})
}