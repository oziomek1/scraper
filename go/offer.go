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
	chassis, doors, seats string
	colour, acrylic, metallic, pearl string
	licencePlate, VIN, crashed, vatInvoice string
	financing, firstOwner, collisionFree string
	rhd, country string
	firstReg, regInPoland, serviceAuth, antiqueReg bool
	tuning, condition string
}

type Offer struct {
	url string
	params Params
	// Options
}

func assignParams(values []string, labels []string) (Params) {
	var params Params
	for idx, param := range labels {
		for key, val := range paramDictionary {
			if val == param {
				fmt.Println(idx, key, val, values[idx])
			}
		}
	}
	return params
}

func readOffer(url string) (*Params) {

	var values []string
	var labels []string
	pageContent, err := parseUrlToNode(url)
	if err != nil {
		fmt.Printf("Error with %s %s", pageContent, err)
		return nil
	}

	fmt.Println("\n\nUrl:", url)

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
	currencyVal := currency.FirstChild.Data
	fmt.Println(currencyVal)

	// -------------------------------------
	// Get offer params
	// -------------------------------------
	values, labels = getOfferParam(pageContent, values, labels)
	fmt.Println(labels)
	fmt.Println(values)
	params := assignParams(values, labels)
	return &params
}

func visitOffer(link string, offers []Offer) {
	params := readOffer(link)

	offer := Offer{link, *params}
	offers = append(offers, offer)
}
