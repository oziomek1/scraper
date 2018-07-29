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
	seller, _ := getElementById("class", "offer-params__list", pageContent)
	fmt.Println(seller.FirstChild.NextSibling)

	_, sellers, _ := getElementsById("class", "offer-params__list", pageContent, nil)
	fmt.Println(sellers)
}

func visitOffer(link string, offers []Offer) {
	readOffer(link)
	offer := Offer{link}
	offers = append(offers, offer)
}
