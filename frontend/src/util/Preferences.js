import UnitConversion from "./UnitConversion";

class Preferences {

    static DISTANCE_UNITS = {
        MILES: "miles",
        KILOMETERS: "km"
    }
    getDistanceUnits() {
        return Preferences.DISTANCE_UNITS.MILES
    }

    displayDistance(input) {
        switch (this.getDistanceUnits()) {
            case Preferences.DISTANCE_UNITS.MILES:
                const conVal = input * 1.0 / UnitConversion.METERS_PER_MILE;
                const stringVal = conVal.toLocaleString('en-US', {
                    maximumFractionDigits: 2
                })
                return stringVal
            default:
                return 0.0
        }
    }

    displayDate(date) {
        let year = new Intl.DateTimeFormat('en', {year: 'numeric'}).format(date)
        let month = new Intl.DateTimeFormat('en', {month: 'short'}).format(date)
        let day = new Intl.DateTimeFormat('en', {day: '2-digit'}).format(date)

        return `${month}-${day}-${year}`
    }


    displayDuration(input) {
        const hours = Math.floor(input / (60 * 60))
        input -= hours * (60 * 60)

        const minutes = Math.floor(input / 60)
        input -= minutes * 60

        const seconds = input

        const hourString = hours.toLocaleString('en-US', {})
        const minString = minutes.toLocaleString('en-US', {
            minimumIntegerDigits: 2,
        })
        const secString = seconds.toLocaleString('en-US', {
            minimumIntegerDigits: 2,
            maximumFractionDigits: 0,
        })

        return `${hourString}:${minString}:${secString}`
    }

    
}



Preferences.instance = function() {
    return Preferences._instance || new Preferences();
}

export default Preferences