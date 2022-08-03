import Preferences from "./Preferences"

class UnitConversion {

    static METERS_PER_MILE = 1609
    convertDistance(input) {
        switch (Preferences.instance().getDistanceUnits()) {
            case Preferences.DISTANCE_UNITS.MILES:
                return input * 1.0 / UnitConversion.METERS_PER_MILE;
            default:
                return 0.0
        }
        
    }

    convertDuration(input) {
        let hours = Math.floor(input / (60 * 60))
        input -= hours * (60 * 60)

        let minutes = Math.floor(input / 60)
        input -= minutes * 60

        let seconds = input

        // console.log(hours + ":" + minutes + ":" + seconds)
        hours = hours.toFixed(2)
        seconds = seconds.toFixed(2)
        return `${hours}:${minutes}:${seconds}`
    }

    convertEpochToDate(input) {
        let d = new Date(0);
        d.setUTCSeconds(input);
        // console.log(d)
        return d
    }

    convertSpeedToPace(input) {
        // input as m/s
        // need min/{mile or km}
        if (input == 0.0) {
            return Math.Infinity;
        }
        const secondsPerMeter = 1.0 / input;
        const minutesPerMeter = secondsPerMeter / 60;

        let result = 0.0;
        switch (Preferences.instance().getDistanceUnits()) {
        case Preferences.DISTANCE_UNITS.MILES:
            result = minutesPerMeter * UnitConversion.METERS_PER_MILE;
            break;
        case Preferences.DISTANCE_UNITS.KILOMETERS:
            result = minutesPerMeter * 1000;
            break;
        default:
            console.error("invalid pace conversion")
        }

        return result;

    }
}

UnitConversion.instance = function() {
    return UnitConversion._instance || new UnitConversion();
}

export default UnitConversion