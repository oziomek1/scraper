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
	url string
	id, price int
	currency string
	seller, category, make, model string
	generation, year, mileage, engineCapacity string
	fuelType, power, gearbox, powertrain string
	chassis, doors, seats string
	colour, acrylic, metallic, pearl string
	licencePlate, VIN, crashed, vatInvoice string
	financing, firstOwner, collisionFree string
	rhd, country string
	firstReg, regInPoland, serviceAuth, antiqueReg string
	tuning, condition string
}

type Offer struct {
	url string
	params Params
	// Options
}

func assignValue(arg string, offerParams map[string]string) string {
	paramValue := "-"
	for dictKey, dictVal := range paramDictionary {
		for offerKey, offerVal := range offerParams {
			if arg == dictKey && dictVal == offerKey {
				paramValue = offerVal
				break
			}
		}
	}
	return paramValue
}

func assignParams(url string, offerId string, price string, currency string, offerParams map[string]string) (Params) {
	var params Params
	params.url = url
	params.id, _ = strconv.Atoi(offerId)
	params.price, _ = strconv.Atoi(price)
	params.currency = currency

	params.seller = assignValue("SELLER", offerParams)
	params.category = assignValue("CATEGORY", offerParams)
	params.make = assignValue("MAKE", offerParams)
	params.model = assignValue("MODEL", offerParams)

	params.generation = assignValue("GENERATION", offerParams)
	params.year = assignValue("YEAR", offerParams)
	params.mileage = assignValue("MILEAGE", offerParams)
	params.engineCapacity = assignValue("ENGINE_CAPACITY", offerParams)
	params.fuelType = assignValue("FUEL_TYPE", offerParams)
	params.power = assignValue("POWER", offerParams)
	params.gearbox = assignValue("GEARBOX", offerParams)
	params.powertrain = assignValue("POWERTRAIN", offerParams)

	params.chassis = assignValue("CHASSIS", offerParams)
	params.doors = assignValue("DOORS", offerParams)
	params.seats = assignValue("SEATS", offerParams)

	params.colour = assignValue("COLOUR", offerParams)
	params.acrylic = assignValue("ACRYLIC", offerParams)
	params.metallic = assignValue("METALLIC", offerParams)
	params.pearl = assignValue("PEARL", offerParams)

	params.licencePlate = assignValue("LICENCE_PLATE", offerParams)
	params.VIN = assignValue("VIN", offerParams)
	params.crashed = assignValue("CRASHED", offerParams)
	params.vatInvoice = assignValue("VAT_INVOICE", offerParams)
	params.financing = assignValue("FINANCING", offerParams)
	params.firstOwner = assignValue("FIRST_OWNER", offerParams)
	params.collisionFree = assignValue("COLLISION_FREE", offerParams)

	params.rhd = assignValue("RHD", offerParams)
	params.country = assignValue("COUNTRY", offerParams)

	params.firstReg = assignValue("FIRST_REGISTRATION", offerParams)
	params.regInPoland = assignValue("REGISTER_IN_POLAND", offerParams)
	params.serviceAuth = assignValue("SERVICE_AUTHORISED", offerParams)
	params.antiqueReg = assignValue("ANTIQUE_REGISTERED", offerParams)

	params.tuning = assignValue("TUNING", offerParams)
	params.condition = assignValue("CONDITION", offerParams)

	fmt.Println(params)
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

	// -------------------------------------
	// Get offer ID
	// -------------------------------------
	offerId := getElementByTag("data-id_raw", pageContent)

	// -------------------------------------
	// Get offer price
	// -------------------------------------
	price := getElementByTag("data-price", pageContent)
	price = strings.Replace(price, " ", "", -1)

	// -------------------------------------
	// Get offer currency
	// -------------------------------------
	currency, _ := getElementById("class", "offer-price__currency", pageContent)
	currencyVal := currency.FirstChild.Data

	// -------------------------------------
	// Get offer params
	// -------------------------------------
	values, labels = getOfferParam(pageContent, values, labels)
	paramsMap := slicesToMap(labels, values)
	params := assignParams(url, offerId, price, currencyVal, paramsMap)
	return &params
}

func visitOffer(link string, offers *[]Offer) {
	params := readOffer(link)

	offer := Offer{link, *params}
	*offers = append(*offers, offer)
}
