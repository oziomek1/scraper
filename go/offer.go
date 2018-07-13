package _go

import (

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
	price, currency string
	seller, category, make, model string
	generation, year, mileage, engineCapacity string
	fuelType, power, gearbox, powertrain string
	chassis, doors, colour, country string
	crashed, condition string
	firstOwner, colission, serviceASO bool
}

type Offer struct {
	id uint64
	url string
	Params
	Options
}

