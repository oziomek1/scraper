package main

import (
	"fmt"
	"strconv"
	"strings"
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


func readOffer(url string) {
	pageContent, err := parseUrlToNode(url)
	if err != nil {
		fmt.Printf("Error with %s %s", pageContent, err)
		return
	}

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
	seller, _ := getElementById("class", "offer-params", pageContent)

	fmt.Println(seller.FirstChild.NextSibling.Data)

	// ul
	//fmt.Println(seller.FirstChild.NextSibling.Data)

	// ul > li (1)
	//fmt.Println(seller.FirstChild.NextSibling.FirstChild.NextSibling.Data)

	// ul > li (2)
	//fmt.Println(seller.FirstChild.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.Data)
	// ul > li > span
	//fmt.Println(seller.FirstChild.NextSibling.FirstChild.NextSibling.FirstChild.NextSibling.Data)

	// ul > li > div
	//fmt.Println(seller.FirstChild.NextSibling.FirstChild.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.Data)

	// ul > li > span > value (1)
	//fmt.Println(seller.FirstChild.NextSibling.FirstChild.NextSibling.FirstChild.NextSibling.FirstChild.Data)

	// ul > li > div > a
	//fmt.Println(seller.FirstChild.NextSibling.FirstChild.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.Data)

	// ul > li > div > a > value (1)
	sellerType 	:= seller.FirstChild.NextSibling.FirstChild.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.FirstChild.Data
	fmt.Println(filterUnnecessaryChars(sellerType))

	// ul > li (2) > div > a > value (1)
	category 	:= seller.FirstChild.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.FirstChild.Data
	fmt.Println(filterUnnecessaryChars(category))

	// ul > li (3) > div > a > value (1)
	make 		:= seller.FirstChild.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.FirstChild.Data
	fmt.Println(filterUnnecessaryChars(make))

	// ul > li (4) > div > a > value (1)
	model		:= seller.FirstChild.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.FirstChild.Data
	fmt.Println(filterUnnecessaryChars(model))

	// ul > li (5) > div > a > value (1)
	generation 	:= seller.FirstChild.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.FirstChild.Data
	fmt.Println(filterUnnecessaryChars(generation))

	//// ul > li (6) > div > value (1)
	year 		:= seller.FirstChild.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.Data
	fmt.Println(filterUnnecessaryChars(year))

	//// ul > li (7) > div > value (1)
	mileage 	:= seller.FirstChild.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.Data
	fmt.Println(filterUnnecessaryChars(mileage))

	//// ul > li (8) > div > value (1)
	engineCap 	:= seller.FirstChild.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.Data
	fmt.Println(filterUnnecessaryChars(engineCap))

	//// ul > li (9) > div > a > value (1)
	fuelType 	:= seller.FirstChild.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.FirstChild.Data
	fmt.Println(filterUnnecessaryChars(fuelType))

	//// ul > li (10) > div > value (1)
	power	 	:= seller.FirstChild.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.Data
	fmt.Println(filterUnnecessaryChars(power))

	//// ul > li (11) > div > a > value (1)
	gearbox 	:= seller.FirstChild.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.FirstChild.Data
	fmt.Println(filterUnnecessaryChars(gearbox))

	//// ul > li (12) > div > a > value (1)
	powertrain 	:= seller.FirstChild.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.FirstChild.Data
	fmt.Println(filterUnnecessaryChars(powertrain))

	//// ul > li (13) > div > a > value (1)
	chassis 	:= seller.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.FirstChild.Data
	fmt.Println(filterUnnecessaryChars(chassis))

	//// ul > li (14) > div > value (1)
	doors 		:= seller.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.Data
	fmt.Println(filterUnnecessaryChars(doors))

	//// ul > li (15) > div > value (1)
	seats 		:= seller.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.Data
	fmt.Println(filterUnnecessaryChars(seats))

	//// ul > li (16) > div > a > value (1)
	colour 		:= seller.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.FirstChild.Data
	fmt.Println(filterUnnecessaryChars(colour))

	//// ul > li (17) > div > a > value (1)
	metallic	:= seller.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.FirstChild.Data
	fmt.Println(filterUnnecessaryChars(metallic))

	//// ul > li (20) > div > a > value (1)
	crashless	:= seller.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.FirstChild.Data
	fmt.Println(filterUnnecessaryChars(crashless))

	//// ul > li (21) > div > a > value (1)
	condition	:= seller.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.FirstChild.Data
	fmt.Println(filterUnnecessaryChars(condition))

	//attributes, ok := getElementById("class", "offer-params__list", pageContent)
	//fmt.Println(attributes, ok)

	//for c := seller.FirstChild; c != nil; c = c.NextSibling {
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
