package main

import (
	"fmt"
	"strings"
	"sync"
)

type Features struct {
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
	id, price string
	currency string
	time, date string
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
	features Features
}

func assignParam(arg string, offerParams map[string]string) string {
	paramValue := `NULL`
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

func assignParams(url string, offerId string, price string, currency string, time string, date string, offerParams map[string]string) (Params) {
	var params Params
	params.url = url
	params.id = offerId
	params.price = price
	params.currency = currency
	params.time = time
	params.date = date
	params.seller = assignParam("SELLER", offerParams)
	params.category = assignParam("CATEGORY", offerParams)
	params.make = assignParam("MAKE", offerParams)
	params.model = assignParam("MODEL", offerParams)

	params.generation = assignParam("GENERATION", offerParams)
	params.year = assignParam("YEAR", offerParams)
	params.mileage = assignParam("MILEAGE", offerParams)
	params.engineCapacity = assignParam("ENGINE_CAPACITY", offerParams)
	params.fuelType = assignParam("FUEL_TYPE", offerParams)
	params.power = assignParam("POWER", offerParams)
	params.gearbox = assignParam("GEARBOX", offerParams)
	params.powertrain = assignParam("POWERTRAIN", offerParams)

	params.chassis = assignParam("CHASSIS", offerParams)
	params.doors = assignParam("DOORS", offerParams)
	params.seats = assignParam("SEATS", offerParams)

	params.colour = assignParam("COLOUR", offerParams)
	params.acrylic = assignParam("ACRYLIC", offerParams)
	params.metallic = assignParam("METALLIC", offerParams)
	params.pearl = assignParam("PEARL", offerParams)

	params.licencePlate = assignParam("LICENCE_PLATE", offerParams)
	params.VIN = assignParam("VIN", offerParams)
	params.crashed = assignParam("CRASHED", offerParams)
	params.vatInvoice = assignParam("VAT_INVOICE", offerParams)
	params.financing = assignParam("FINANCING", offerParams)
	params.firstOwner = assignParam("FIRST_OWNER", offerParams)
	params.collisionFree = assignParam("COLLISION_FREE", offerParams)

	params.rhd = assignParam("RHD", offerParams)
	params.country = assignParam("COUNTRY", offerParams)

	params.firstReg = assignParam("FIRST_REGISTRATION", offerParams)
	params.regInPoland = assignParam("REGISTER_IN_POLAND", offerParams)
	params.serviceAuth = assignParam("SERVICE_AUTHORISED", offerParams)
	params.antiqueReg = assignParam("ANTIQUE_REGISTERED", offerParams)

	params.tuning = assignParam("TUNING", offerParams)
	params.condition = assignParam("CONDITION", offerParams)

	return params
}

func assignFeature(arg string, featureValues []string) bool {
	feature := false
	for _, val := range featureValues {
		if val == arg {
			return true
		}
	}
	return feature
}

func assignFeatures(featureValues []string) (Features)  {
	var features Features
	features.abs = assignFeature(featuresList[0], featureValues)
	features.cd = assignFeature(featuresList[1], featureValues)
	features.centralLock = assignFeature(featuresList[2], featureValues)
	features.elFrontWindows = assignFeature(featuresList[3], featureValues)
	features.elMirrors = assignFeature(featuresList[4], featureValues)
	features.immobiliser = assignFeature(featuresList[5], featureValues)
	features.driverAirbag = assignFeature(featuresList[6], featureValues)
	features.passengerAirbag = assignFeature(featuresList[7], featureValues)
	features.factoryRadio = assignFeature(featuresList[8], featureValues)
	features.driveAssist = assignFeature(featuresList[9], featureValues)
	features.alarm = assignFeature(featuresList[10], featureValues)
	features.aluWheels = assignFeature(featuresList[11], featureValues)
	features.asr = assignFeature(featuresList[12], featureValues)
	features.parkAssist = assignFeature(featuresList[13], featureValues)
	features.laneAssist = assignFeature(featuresList[14], featureValues)
	features.bluetooth = assignFeature(featuresList[15], featureValues)
	features.rainSensor = assignFeature(featuresList[16], featureValues)
	features.blindSpotMonitor = assignFeature(featuresList[17], featureValues)
	features.nightSensor = assignFeature(featuresList[18], featureValues)
	features.frontParkAssist = assignFeature(featuresList[19], featureValues)
	features.rearParkAssist = assignFeature(featuresList[20], featureValues)
	features.panoramaGlassRoof = assignFeature(featuresList[21], featureValues)
	features.elSeats = assignFeature(featuresList[22], featureValues)
	features.esp = assignFeature(featuresList[23], featureValues)
	features.aux = assignFeature(featuresList[24], featureValues)
	features.sdCard = assignFeature(featuresList[25], featureValues)
	features.usb = assignFeature(featuresList[26], featureValues)
	features.towHook = assignFeature(featuresList[27], featureValues)
	features.hud = assignFeature(featuresList[28], featureValues)
	features.isofix = assignFeature(featuresList[29], featureValues)
	features.rearCamera = assignFeature(featuresList[30], featureValues)
	features.autoAirCon = assignFeature(featuresList[31], featureValues)
	features.autoAirCon4 = assignFeature(featuresList[32], featureValues)
	features.autoAirCon2 = assignFeature(featuresList[33], featureValues)
	features.manualAirCon = assignFeature(featuresList[34], featureValues)
	features.computer = assignFeature(featuresList[35], featureValues)
	features.airBagCourtain = assignFeature(featuresList[36], featureValues)
	features.paddleShift = assignFeature(featuresList[37], featureValues)
	features.mp3 = assignFeature(featuresList[38], featureValues)
	features.gps = assignFeature(featuresList[39], featureValues)
	features.dvd = assignFeature(featuresList[40], featureValues)
	features.speedLimiter = assignFeature(featuresList[41], featureValues)
	features.webasto = assignFeature(featuresList[42], featureValues)
	features.heatedWindShield = assignFeature(featuresList[43], featureValues)
	features.heatedSideMirror = assignFeature(featuresList[44], featureValues)
	features.heatedFrontSeat = assignFeature(featuresList[45], featureValues)
	features.heatedRearSeat = assignFeature(featuresList[46], featureValues)
	features.kneeAirBag = assignFeature(featuresList[47], featureValues)
	features.frontSideAirBag = assignFeature(featuresList[48], featureValues)
	features.rearSideAirBag = assignFeature(featuresList[49], featureValues)
	features.tintedWindow = assignFeature(featuresList[50], featureValues)
	features.nonFactoryRadio = assignFeature(featuresList[51], featureValues)
	features.adjustableSuspension = assignFeature(featuresList[52], featureValues)
	features.roofRail = assignFeature(featuresList[53], featureValues)
	features.startStop = assignFeature(featuresList[54], featureValues)
	features.sunroof = assignFeature(featuresList[55], featureValues)
	features.dayLight = assignFeature(featuresList[56], featureValues)
	features.ledLight = assignFeature(featuresList[57], featureValues)
	features.antifogLight = assignFeature(featuresList[58], featureValues)
	features.xenonLight = assignFeature(featuresList[59], featureValues)
	features.leather = assignFeature(featuresList[60], featureValues)
	features.velour = assignFeature(featuresList[61], featureValues)
	features.cruiseControl = assignFeature(featuresList[62], featureValues)
	features.activeCruiseControl = assignFeature(featuresList[63], featureValues)
	features.tunerTV = assignFeature(featuresList[64], featureValues)
	features.multifuncionalWheel = assignFeature(featuresList[65], featureValues)
	features.changerCD = assignFeature(featuresList[66], featureValues)
	return features
}

func readOffer(url string) (*Params, *Features) {

	pageContent, err := parseUrlToNode(url)
	if err != nil {
		fmt.Printf("Error with %s %s", pageContent, err)
		return nil, nil
	}

	// -------------------------------------
	// Get offer ID
	// -------------------------------------
	offerId := getElementByTag("data-id_raw", pageContent)

	// -------------------------------------
	// Get offer time and date
	// -------------------------------------
	offerTime, offerDate := getOfferTimeAndDate("offer-meta__value", pageContent)

	// -------------------------------------
	// Get offer price
	// -------------------------------------
	price := getElementByTag("data-price", pageContent)
	price = strings.Replace(price, " ", "", -1)

	// -------------------------------------
	// Get offer currency
	// -------------------------------------
	// currency, _ := getElementById("class", "offer-price__currency", pageContent)
	 currencyVal := "PLN"

	// -------------------------------------
	// Get offer params
	// -------------------------------------
	var values []string
	var labels []string
	var featureValues []string

	attributeType := "class"
	nodeTag := "offer-params"
	pageNodeContent, _ := getElementById(attributeType, nodeTag, pageContent)
	values, labels = getOfferParam(pageNodeContent, values, labels)
	nodeTag = "offer-features"
	pageNodeContent, _ = getElementById(attributeType, nodeTag, pageContent)
	featureValues = getOfferFeatures(pageNodeContent, featureValues)
	paramsMap := slicesToMap(labels, values)
	params := assignParams(url, offerId, price, currencyVal, offerTime, offerDate, paramsMap)
	features := assignFeatures(featureValues)
	if params.id == "" {
		fmt.Println("[WARN] EMPTY LINK")
	}
	fmt.Println(params)
	return &params, &features
}

func visitOffer(link string, offers *[]Offer, wg *sync.WaitGroup) {
	defer wg.Done()
	if link != "" {
		params, features := readOffer(link)
		offer := Offer{link, *params, *features}
		*offers = append(*offers, offer)
	} else {
		fmt.Println("[WARN] EMPTY LINK")
	}
}
