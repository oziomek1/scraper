package main

import (
	"fmt"
	"strconv"
	"strings"
	"golang.org/x/net/html"
)

type Options struct {
	abs, cd, centralLock, elFrontWindows bool
	elMirrors, immobiliser, driverAirbag, passengerAirbag bool
	factoryRadio, driveAssist, alarm, aluWheels bool
	asr, parkAssist, laneAssist, bluetooth bool
	rainSensor, blindSpotMonitor, nightSensor, frontParkAssist bool
	rearParkAssist, panoramaGlassRoof, elChrSideMirror, elChrRearMirror bool
	elRearWindow, elSeats, esp, aux bool
	sdCard, usb, towHook, hud bool
	isofix, rearCamera, autoAirCon, autoAirCon4 bool
	autoAirCon2, manualAirCon, computer, airBagCourtain bool
	paddleShift, mp3, gps, dvd bool
	speedLimiter, webasto, heatedWindShield, heatedSideMirror bool
	heatedFrontSeat, heatedRearSeat, kneeAirBag, frontSideAirBag bool
	rearSideAirBag, tintedWindow, nonFactoryRadio, adjustableSuspension bool
	roofRail, startStop, sunroof, dayLight bool
	ledLight, antifogLight, xenonLight, leather bool
	velour, cruiseControl, activeCruiseControl, tunerTV bool
	multifuncionalWheel, changerCD bool

}

type Params struct {
	id int
	price, currency string
	seller, category, make, model string
	generation, year, mileage, engineCapacity string
	fuelType, power, gearbox, powertrain string
	chassis, doors, colour, country string
	crashed, condition string
	firstOwner, colission, serviceASO bool
}

type Offer struct {
	url string
	// Params
	// Options
}

var paramsList = [...]string{"Seller", "Category", "Make", "Model", "Generation"}

func recursiveParams(node *html.Node, idx int) {
	if idx >= len(paramsList) /* Seller, Category, Make, Model, Generation */ {
		return
	} else {
		idx += 1
	}
	newNode := node.NextSibling.NextSibling
	printNode := node.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.FirstChild.Data
	fmt.Println(strings.TrimSpace(strings.TrimRight(printNode, "\r\n")))
	recursiveParams(newNode, idx)
}

func readOffer(url string) {
	pageContent, err := parseUrlToNode(url)
	if err != nil {
		fmt.Printf("Error with %s %s", pageContent, err)
		return
	}

	fmt.Println("Url: ", url)

	// -------------------------------------
	// Get offer ID
	// -------------------------------------
	offerId := getElementByTag("data-id_raw", pageContent)
	if offerId, err := strconv.ParseInt(offerId, 10, 64); err == nil {
		fmt.Println(offerId)
	}

	// -------------------------------------
	// Get offer price
	// -------------------------------------
	price := getElementByTag("data-price", pageContent)
	price = strings.Replace(price, " ", "", -1)
	fmt.Println(price)

	// -------------------------------------
	// Get offer currency
	// -------------------------------------
	currency, _ := getElementById("class", "offer-price__currency", pageContent)
	fmt.Println(currency.FirstChild.Data)

	// -------------------------------------
	// Get offer seller
	// -------------------------------------
	offerParams, _ := getElementById("class", "offer-params", pageContent)

	recursiveParams(offerParams.FirstChild.NextSibling.FirstChild.NextSibling, 0)

	// temporary these params below won't be collected as otomoto has few different layouts for these params

	//// ul > li > div > a > value (1)
	//seller 		:= offerParams.FirstChild.NextSibling.FirstChild.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.FirstChild.Data
	//fmt.Println(filterUnnecessaryChars(seller))
	//
	//// ul > li (2) > div > a > value (1)
	//category 	:= offerParams.FirstChild.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.FirstChild.Data
	//fmt.Println(filterUnnecessaryChars(category))
	//
	//// ul > li (3) > div > a > value (1)
	//make 		:= offerParams.FirstChild.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.FirstChild.Data
	//fmt.Println(filterUnnecessaryChars(make))
	//
	//// ul > li (4) > div > a > value (1)
	//model		:= offerParams.FirstChild.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.FirstChild.Data
	//fmt.Println(filterUnnecessaryChars(model))
	//
	//// ul > li (5) > div > a > value (1)
	//generation 	:= offerParams.FirstChild.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.FirstChild.Data
	//fmt.Println(filterUnnecessaryChars(generation))
	//
	////// ul > li (6) > div > value (1)
	//year 		:= offerParams.FirstChild.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.Data
	//fmt.Println(filterUnnecessaryChars(year))
	//
	////// ul > li (7) > div > value (1)
	//mileage 	:= offerParams.FirstChild.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.Data
	//fmt.Println(filterUnnecessaryChars(mileage))
	//
	////// ul > li (8) > div > value (1)
	//engineCap 	:= offerParams.FirstChild.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.Data
	//fmt.Println(filterUnnecessaryChars(engineCap))
	//
	////// ul > li (9) > div > a > value (1)
	//fuelType 	:= offerParams.FirstChild.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.FirstChild.Data
	//fmt.Println(filterUnnecessaryChars(fuelType))
	//
	////// ul > li (10) > div > value (1)
	//power	 	:= offerParams.FirstChild.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.Data
	//fmt.Println(filterUnnecessaryChars(power))
	//
	////// ul > li (11) > div > a > value (1)
	//gearbox 	:= offerParams.FirstChild.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.FirstChild.Data
	//fmt.Println(filterUnnecessaryChars(gearbox))
	//
	////// ul > li (12) > div > a > value (1)
	//powertrain 	:= offerParams.FirstChild.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.FirstChild.Data
	//fmt.Println(filterUnnecessaryChars(powertrain))
	//
	//// ul > li (13) > div > a > value (1)
	//chassis 	:= offerParams.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.FirstChild.Data
	//fmt.Println(filterUnnecessaryChars(chassis))
	//
	////// ul > li (14) > div > value (1)
	//doors 		:= offerParams.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.Data
	//fmt.Println(filterUnnecessaryChars(doors))
	//
	////// ul > li (15) > div > value (1)
	//seats 		:= offerParams.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.Data
	//fmt.Println(filterUnnecessaryChars(seats))
	//
	////// ul > li (16) > div > a > value (1)
	//colour 		:= offerParams.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.FirstChild.Data
	//fmt.Println(filterUnnecessaryChars(colour))
	//
	////// ul > li (17) > div > a > value (1)
	//metallic	:= offerParams.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.FirstChild.Data
	//fmt.Println(filterUnnecessaryChars(metallic))
	//
	////// ul > li (20) > div > a > value (1)
	//crashless	:= offerParams.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.FirstChild.Data
	//fmt.Println(filterUnnecessaryChars(crashless))
	//
	////// ul > li (21) > div > a > value (1)
	//condition	:= offerParams.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.FirstChild.Data
	//fmt.Println(filterUnnecessaryChars(condition))

	//attributes, ok := getElementById("class", "offer-params__list", pageContent)
	//fmt.Println(attributes, ok)

	//for c := offerParams.FirstChild; c != nil; c = c.NextSibling {
		//fmt.Println(getElementByTag("offer-params__value", c))
	//}
	//var elements []string
	//var node *html.Node
	//node, elements = getElementsById("offer-params__label", attributes, elements)
	//fmt.Println("ELEMENTS", node.Data, elements)
}

func visitOffer(link string, offers []Offer) {
	readOffer(link)
	offer := Offer{link}
	offers = append(offers, offer)
}
