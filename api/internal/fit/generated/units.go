// Code generated; DO NOT EDIT

// Code generated; DO NOT EDIT

package generated

import "fmt"
import "math"

type Unit_e uint8

const (
	UNIT_YEAR Unit_e = iota
	UNIT_HOUR
	UNIT_MINUTE
	UNIT_SECOND
	UNIT_MILLISECOND
	UNIT_METER
	UNIT_MILLIMETER
	UNIT_METERPERSECOND
	UNIT_COUNT
	UNIT_COUNTRATE
	UNIT_KILOCALORIE
	UNIT_DEGREES
	UNIT_SEMICIRCLES
	UNIT_UNKNOWN
)

type Unit struct {
	Unit   Unit_e
	Scale  uint
	Offset int
}

var conversionMultiplier map[Unit_e]map[Unit_e]float64

func init() {
	conversionMultiplier = make(map[Unit_e]map[Unit_e]float64)

	conversionMultiplier[UNIT_SEMICIRCLES] = make(map[Unit_e]float64)
	conversionMultiplier[UNIT_SEMICIRCLES][UNIT_DEGREES] = 180.0 / math.Pow(2, 31)

	conversionMultiplier[UNIT_YEAR] = make(map[Unit_e]float64)
	conversionMultiplier[UNIT_YEAR][UNIT_YEAR] = 1.0

	conversionMultiplier[UNIT_HOUR] = make(map[Unit_e]float64)
	conversionMultiplier[UNIT_HOUR][UNIT_HOUR] = 1.0

	conversionMultiplier[UNIT_MINUTE] = make(map[Unit_e]float64)
	conversionMultiplier[UNIT_MINUTE][UNIT_MINUTE] = 1.0
	
	conversionMultiplier[UNIT_SECOND] = make(map[Unit_e]float64)
	conversionMultiplier[UNIT_SECOND][UNIT_SECOND] = 1.0

	conversionMultiplier[UNIT_METER] = make(map[Unit_e]float64)
	conversionMultiplier[UNIT_METER][UNIT_METER] = 1.0

	conversionMultiplier[UNIT_SECOND] = make(map[Unit_e]float64)
	conversionMultiplier[UNIT_SECOND][UNIT_SECOND] = 1.0

	conversionMultiplier[UNIT_METERPERSECOND] = make(map[Unit_e]float64)
	conversionMultiplier[UNIT_METERPERSECOND][UNIT_METERPERSECOND] = 1.0

	conversionMultiplier[UNIT_COUNT] = make(map[Unit_e]float64)
	conversionMultiplier[UNIT_COUNT][UNIT_COUNT] = 1.0

	conversionMultiplier[UNIT_COUNTRATE] = make(map[Unit_e]float64)
	conversionMultiplier[UNIT_COUNTRATE][UNIT_COUNTRATE] = 1.0
/* TODO: ...


	conversionMultiplier[UNIT_YEAR] = make(map[Unit_e]float64)
	conversionMultiplier[UNIT_YEAR][UNIT_YEAR] = 1.0

	conversionMultiplier[UNIT_YEAR] = make(map[Unit_e]float64)
	conversionMultiplier[UNIT_YEAR][UNIT_YEAR] = 1.0

	conversionMultiplier[UNIT_YEAR] = make(map[Unit_e]float64)
	conversionMultiplier[UNIT_YEAR][UNIT_YEAR] = 1.0

	conversionMultiplier[UNIT_YEAR] = make(map[Unit_e]float64)
	conversionMultiplier[UNIT_YEAR][UNIT_YEAR] = 1.0

	conversionMultiplier[UNIT_YEAR] = make(map[Unit_e]float64)
	conversionMultiplier[UNIT_YEAR][UNIT_YEAR] = 1.0
	*/
}

func Convert(value interface{}, startUnit Unit, goalUnit Unit_e) float64 {
	var startValue float64
	switch i := value.(type) {
	case float64:
		startValue = i
	case float32:
		startValue = float64(i)
	case int64:
		startValue = float64(i)
	case int32:
		startValue = float64(i)
	case int16:
		startValue = float64(i)
	case int8:
		startValue = float64(i)
	case uint64:
		startValue = float64(i)
	case uint32:
		startValue = float64(i)
	case uint16:
		startValue = float64(i)
	case uint8:
		startValue = float64(i)
	default:
		// startValue = float64(i)
		panic(fmt.Errorf("error converting")) // TODO
	}

	newValue := float64(startValue)/float64(startUnit.Scale) - float64(startUnit.Offset)
	multiplier := conversionMultiplier[startUnit.Unit][goalUnit]
	newValue *= multiplier
	// newValue = (newValue + float64(goalUnit.Offset)) * float64(goalUnit.Scale)
	return newValue
}

func ConvertEnum(value interface{}, startUnit Unit_e, goalUnit Unit_e) float64 {
	return Convert(value, 
		Unit{Unit: startUnit, Scale: 1, Offset: 0},
		goalUnit,
	 )
}

var (
	FILE_ID_TYPE__DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	FILE_ID_MANUFACTURER_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	FILE_ID_PRODUCT_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	FILE_ID_SERIAL_NUMBER_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	FILE_ID_TIME_CREATED_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	FILE_ID_NUMBER_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	FILE_CREATOR_SOFTWARE_VERSION_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	FILE_CREATOR_HARDWARE_VERSION_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SOFTWARE_MESSAGE_INDEX_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SOFTWARE_VERSION_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 100, 
		Offset: 0, 
	}
	SOFTWARE_PART_NUMBER_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SLAVE_DEVICE_MANUFACTURER_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SLAVE_DEVICE_PRODUCT_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	CAPABILITIES_LANGUAGES_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	CAPABILITIES_SPORTS_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	CAPABILITIES_WORKOUTS_SUPPORTED_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	CAPABILITIES_CONNECTIVITY_SUPPORTED_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	FILE_CAPABILITIES_MESSAGE_INDEX_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	FILE_CAPABILITIES_TYPE__DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	FILE_CAPABILITIES_FLAGS_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	FILE_CAPABILITIES_DIRECTORY_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	FILE_CAPABILITIES_MAX_COUNT_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	FILE_CAPABILITIES_MAX_SIZE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	MESG_CAPABILITIES_MESSAGE_INDEX_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	MESG_CAPABILITIES_FILE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	MESG_CAPABILITIES_MESG_NUM_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	MESG_CAPABILITIES_COUNT_TYPE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	MESG_CAPABILITIES_COUNT_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	FIELD_CAPABILITIES_MESSAGE_INDEX_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	FIELD_CAPABILITIES_FILE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	FIELD_CAPABILITIES_MESG_NUM_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	FIELD_CAPABILITIES_FIELD_NUM_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	FIELD_CAPABILITIES_COUNT_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	DEVICE_SETTINGS_ACTIVE_TIME_ZONE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	DEVICE_SETTINGS_UTC_OFFSET_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	DEVICE_SETTINGS_TIME_ZONE_OFFSET_DEFAULT_UNIT = Unit{
		Unit: UNIT_HOUR, 
		Scale: 4, 
		Offset: 0, 
	}
	USER_PROFILE_MESSAGE_INDEX_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	USER_PROFILE_FRIENDLY_NAME_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	USER_PROFILE_GENDER_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	USER_PROFILE_AGE_DEFAULT_UNIT = Unit{
		Unit: UNIT_YEAR, 
		Scale: 1, 
		Offset: 0, 
	}
	USER_PROFILE_HEIGHT_DEFAULT_UNIT = Unit{
		Unit: UNIT_METER, 
		Scale: 100, 
		Offset: 0, 
	}
	USER_PROFILE_WEIGHT_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 10, 
		Offset: 0, 
	}
	USER_PROFILE_LANGUAGE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	USER_PROFILE_ELEV_SETTING_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	USER_PROFILE_WEIGHT_SETTING_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	USER_PROFILE_RESTING_HEART_RATE_DEFAULT_UNIT = Unit{
		Unit: UNIT_COUNTRATE, 
		Scale: 1, 
		Offset: 0, 
	}
	USER_PROFILE_DEFAULT_MAX_RUNNING_HEART_RATE_DEFAULT_UNIT = Unit{
		Unit: UNIT_COUNTRATE, 
		Scale: 1, 
		Offset: 0, 
	}
	USER_PROFILE_DEFAULT_MAX_BIKING_HEART_RATE_DEFAULT_UNIT = Unit{
		Unit: UNIT_COUNTRATE, 
		Scale: 1, 
		Offset: 0, 
	}
	USER_PROFILE_DEFAULT_MAX_HEART_RATE_DEFAULT_UNIT = Unit{
		Unit: UNIT_COUNTRATE, 
		Scale: 1, 
		Offset: 0, 
	}
	USER_PROFILE_HR_SETTING_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	USER_PROFILE_SPEED_SETTING_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	USER_PROFILE_DIST_SETTING_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	USER_PROFILE_POWER_SETTING_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	USER_PROFILE_ACTIVITY_CLASS_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	USER_PROFILE_POSITION_SETTING_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	USER_PROFILE_TEMPERATURE_SETTING_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	USER_PROFILE_LOCAL_ID_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	USER_PROFILE_GLOBAL_ID_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	USER_PROFILE_HEIGHT_SETTING_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	HRM_PROFILE_MESSAGE_INDEX_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	HRM_PROFILE_ENABLED_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	HRM_PROFILE_HRM_ANT_ID_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	HRM_PROFILE_LOG_HRV_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	HRM_PROFILE_HRM_ANT_ID_TRANS_TYPE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SDM_PROFILE_MESSAGE_INDEX_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SDM_PROFILE_ENABLED_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SDM_PROFILE_SDM_ANT_ID_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SDM_PROFILE_SDM_CAL_FACTOR_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 10, 
		Offset: 0, 
	}
	SDM_PROFILE_ODOMETER_DEFAULT_UNIT = Unit{
		Unit: UNIT_METER, 
		Scale: 100, 
		Offset: 0, 
	}
	SDM_PROFILE_SPEED_SOURCE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SDM_PROFILE_SDM_ANT_ID_TRANS_TYPE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SDM_PROFILE_ODOMETER_ROLLOVER_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	BIKE_PROFILE_MESSAGE_INDEX_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	BIKE_PROFILE_NAME_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	BIKE_PROFILE_SPORT_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	BIKE_PROFILE_SUB_SPORT_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	BIKE_PROFILE_ODOMETER_DEFAULT_UNIT = Unit{
		Unit: UNIT_METER, 
		Scale: 100, 
		Offset: 0, 
	}
	BIKE_PROFILE_BIKE_SPD_ANT_ID_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	BIKE_PROFILE_BIKE_CAD_ANT_ID_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	BIKE_PROFILE_BIKE_SPDCAD_ANT_ID_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	BIKE_PROFILE_BIKE_POWER_ANT_ID_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	BIKE_PROFILE_CUSTOM_WHEELSIZE_DEFAULT_UNIT = Unit{
		Unit: UNIT_METER, 
		Scale: 1000, 
		Offset: 0, 
	}
	BIKE_PROFILE_AUTO_WHEELSIZE_DEFAULT_UNIT = Unit{
		Unit: UNIT_METER, 
		Scale: 1000, 
		Offset: 0, 
	}
	BIKE_PROFILE_BIKE_WEIGHT_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 10, 
		Offset: 0, 
	}
	BIKE_PROFILE_POWER_CAL_FACTOR_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 10, 
		Offset: 0, 
	}
	BIKE_PROFILE_AUTO_WHEEL_CAL_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	BIKE_PROFILE_AUTO_POWER_ZERO_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	BIKE_PROFILE_ID_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	BIKE_PROFILE_SPD_ENABLED_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	BIKE_PROFILE_CAD_ENABLED_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	BIKE_PROFILE_SPDCAD_ENABLED_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	BIKE_PROFILE_POWER_ENABLED_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	BIKE_PROFILE_CRANK_LENGTH_DEFAULT_UNIT = Unit{
		Unit: UNIT_MILLIMETER, 
		Scale: 2, 
		Offset: -110, 
	}
	BIKE_PROFILE_ENABLED_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	BIKE_PROFILE_BIKE_SPD_ANT_ID_TRANS_TYPE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	BIKE_PROFILE_BIKE_CAD_ANT_ID_TRANS_TYPE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	BIKE_PROFILE_BIKE_SPDCAD_ANT_ID_TRANS_TYPE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	BIKE_PROFILE_BIKE_POWER_ANT_ID_TRANS_TYPE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	BIKE_PROFILE_ODOMETER_ROLLOVER_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	BIKE_PROFILE_FRONT_GEAR_NUM_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	BIKE_PROFILE_FRONT_GEAR_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	BIKE_PROFILE_REAR_GEAR_NUM_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	BIKE_PROFILE_REAR_GEAR_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	BIKE_PROFILE_SHIMANO_DI2_ENABLED_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	ZONES_TARGET_MAX_HEART_RATE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	ZONES_TARGET_THRESHOLD_HEART_RATE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	ZONES_TARGET_FUNCTIONAL_THRESHOLD_POWER_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	ZONES_TARGET_HR_CALC_TYPE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	ZONES_TARGET_PWR_CALC_TYPE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SPORT_SPORT_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SPORT_SUB_SPORT_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SPORT_NAME_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	HR_ZONE_MESSAGE_INDEX_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	HR_ZONE_HIGH_BPM_DEFAULT_UNIT = Unit{
		Unit: UNIT_COUNTRATE, 
		Scale: 1, 
		Offset: 0, 
	}
	HR_ZONE_NAME_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SPEED_ZONE_MESSAGE_INDEX_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SPEED_ZONE_HIGH_VALUE_DEFAULT_UNIT = Unit{
		Unit: UNIT_METERPERSECOND, 
		Scale: 1000, 
		Offset: 0, 
	}
	SPEED_ZONE_NAME_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	CADENCE_ZONE_MESSAGE_INDEX_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	CADENCE_ZONE_HIGH_VALUE_DEFAULT_UNIT = Unit{
		Unit: UNIT_COUNTRATE, 
		Scale: 1, 
		Offset: 0, 
	}
	CADENCE_ZONE_NAME_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	POWER_ZONE_MESSAGE_INDEX_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	POWER_ZONE_HIGH_VALUE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	POWER_ZONE_NAME_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	MET_ZONE_MESSAGE_INDEX_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	MET_ZONE_HIGH_BPM_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	MET_ZONE_CALORIES_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 10, 
		Offset: 0, 
	}
	MET_ZONE_FAT_CALORIES_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 10, 
		Offset: 0, 
	}
	GOAL_MESSAGE_INDEX_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	GOAL_SPORT_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	GOAL_SUB_SPORT_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	GOAL_START_DATE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	GOAL_END_DATE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	GOAL_TYPE__DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	GOAL_VALUE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	GOAL_REPEAT_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	GOAL_TARGET_VALUE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	GOAL_RECURRENCE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	GOAL_RECURRENCE_VALUE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	GOAL_ENABLED_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	ACTIVITY_TIMESTAMP_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	ACTIVITY_TOTAL_TIMER_TIME_DEFAULT_UNIT = Unit{
		Unit: UNIT_SECOND, 
		Scale: 1000, 
		Offset: 0, 
	}
	ACTIVITY_NUM_SESSIONS_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	ACTIVITY_TYPE__DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	ACTIVITY_EVENT_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	ACTIVITY_EVENT_TYPE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	ACTIVITY_LOCAL_TIMESTAMP_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	ACTIVITY_EVENT_GROUP_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SESSION_MESSAGE_INDEX_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SESSION_TIMESTAMP_DEFAULT_UNIT = Unit{
		Unit: UNIT_SECOND, 
		Scale: 1, 
		Offset: 0, 
	}
	SESSION_EVENT_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SESSION_EVENT_TYPE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SESSION_START_TIME_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SESSION_START_POSITION_LAT_DEFAULT_UNIT = Unit{
		Unit: UNIT_SEMICIRCLES, 
		Scale: 1, 
		Offset: 0, 
	}
	SESSION_START_POSITION_LONG_DEFAULT_UNIT = Unit{
		Unit: UNIT_SEMICIRCLES, 
		Scale: 1, 
		Offset: 0, 
	}
	SESSION_SPORT_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SESSION_SUB_SPORT_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SESSION_TOTAL_ELAPSED_TIME_DEFAULT_UNIT = Unit{
		Unit: UNIT_SECOND, 
		Scale: 1000, 
		Offset: 0, 
	}
	SESSION_TOTAL_TIMER_TIME_DEFAULT_UNIT = Unit{
		Unit: UNIT_SECOND, 
		Scale: 1000, 
		Offset: 0, 
	}
	SESSION_TOTAL_DISTANCE_DEFAULT_UNIT = Unit{
		Unit: UNIT_METER, 
		Scale: 100, 
		Offset: 0, 
	}
	SESSION_TOTAL_CYCLES_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SESSION_TOTAL_CALORIES_DEFAULT_UNIT = Unit{
		Unit: UNIT_KILOCALORIE, 
		Scale: 1, 
		Offset: 0, 
	}
	SESSION_TOTAL_FAT_CALORIES_DEFAULT_UNIT = Unit{
		Unit: UNIT_KILOCALORIE, 
		Scale: 1, 
		Offset: 0, 
	}
	SESSION_AVG_SPEED_DEFAULT_UNIT = Unit{
		Unit: UNIT_METERPERSECOND, 
		Scale: 1000, 
		Offset: 0, 
	}
	SESSION_MAX_SPEED_DEFAULT_UNIT = Unit{
		Unit: UNIT_METERPERSECOND, 
		Scale: 1000, 
		Offset: 0, 
	}
	SESSION_AVG_HEART_RATE_DEFAULT_UNIT = Unit{
		Unit: UNIT_COUNTRATE, 
		Scale: 1, 
		Offset: 0, 
	}
	SESSION_MAX_HEART_RATE_DEFAULT_UNIT = Unit{
		Unit: UNIT_COUNTRATE, 
		Scale: 1, 
		Offset: 0, 
	}
	SESSION_AVG_CADENCE_DEFAULT_UNIT = Unit{
		Unit: UNIT_COUNTRATE, 
		Scale: 1, 
		Offset: 0, 
	}
	SESSION_MAX_CADENCE_DEFAULT_UNIT = Unit{
		Unit: UNIT_COUNTRATE, 
		Scale: 1, 
		Offset: 0, 
	}
	SESSION_AVG_POWER_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SESSION_MAX_POWER_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SESSION_TOTAL_ASCENT_DEFAULT_UNIT = Unit{
		Unit: UNIT_METER, 
		Scale: 1, 
		Offset: 0, 
	}
	SESSION_TOTAL_DESCENT_DEFAULT_UNIT = Unit{
		Unit: UNIT_METER, 
		Scale: 1, 
		Offset: 0, 
	}
	SESSION_TOTAL_TRAINING_EFFECT_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 10, 
		Offset: 0, 
	}
	SESSION_FIRST_LAP_INDEX_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SESSION_NUM_LAPS_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SESSION_EVENT_GROUP_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SESSION_TRIGGER_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SESSION_NEC_LAT_DEFAULT_UNIT = Unit{
		Unit: UNIT_SEMICIRCLES, 
		Scale: 1, 
		Offset: 0, 
	}
	SESSION_NEC_LONG_DEFAULT_UNIT = Unit{
		Unit: UNIT_SEMICIRCLES, 
		Scale: 1, 
		Offset: 0, 
	}
	SESSION_SWC_LAT_DEFAULT_UNIT = Unit{
		Unit: UNIT_SEMICIRCLES, 
		Scale: 1, 
		Offset: 0, 
	}
	SESSION_SWC_LONG_DEFAULT_UNIT = Unit{
		Unit: UNIT_SEMICIRCLES, 
		Scale: 1, 
		Offset: 0, 
	}
	SESSION_NORMALIZED_POWER_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SESSION_TRAINING_STRESS_SCORE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 10, 
		Offset: 0, 
	}
	SESSION_INTENSITY_FACTOR_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1000, 
		Offset: 0, 
	}
	SESSION_LEFT_RIGHT_BALANCE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SESSION_AVG_STROKE_COUNT_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 10, 
		Offset: 0, 
	}
	SESSION_AVG_STROKE_DISTANCE_DEFAULT_UNIT = Unit{
		Unit: UNIT_METER, 
		Scale: 100, 
		Offset: 0, 
	}
	SESSION_SWIM_STROKE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SESSION_POOL_LENGTH_DEFAULT_UNIT = Unit{
		Unit: UNIT_METER, 
		Scale: 100, 
		Offset: 0, 
	}
	SESSION_THRESHOLD_POWER_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SESSION_POOL_LENGTH_UNIT_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SESSION_NUM_ACTIVE_LENGTHS_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SESSION_TOTAL_WORK_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SESSION_AVG_ALTITUDE_DEFAULT_UNIT = Unit{
		Unit: UNIT_METER, 
		Scale: 5, 
		Offset: 500, 
	}
	SESSION_MAX_ALTITUDE_DEFAULT_UNIT = Unit{
		Unit: UNIT_METER, 
		Scale: 5, 
		Offset: 500, 
	}
	SESSION_GPS_ACCURACY_DEFAULT_UNIT = Unit{
		Unit: UNIT_METER, 
		Scale: 1, 
		Offset: 0, 
	}
	SESSION_AVG_GRADE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 100, 
		Offset: 0, 
	}
	SESSION_AVG_POS_GRADE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 100, 
		Offset: 0, 
	}
	SESSION_AVG_NEG_GRADE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 100, 
		Offset: 0, 
	}
	SESSION_MAX_POS_GRADE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 100, 
		Offset: 0, 
	}
	SESSION_MAX_NEG_GRADE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 100, 
		Offset: 0, 
	}
	SESSION_AVG_TEMPERATURE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SESSION_MAX_TEMPERATURE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SESSION_TOTAL_MOVING_TIME_DEFAULT_UNIT = Unit{
		Unit: UNIT_SECOND, 
		Scale: 1000, 
		Offset: 0, 
	}
	SESSION_AVG_POS_VERTICAL_SPEED_DEFAULT_UNIT = Unit{
		Unit: UNIT_METERPERSECOND, 
		Scale: 1000, 
		Offset: 0, 
	}
	SESSION_AVG_NEG_VERTICAL_SPEED_DEFAULT_UNIT = Unit{
		Unit: UNIT_METERPERSECOND, 
		Scale: 1000, 
		Offset: 0, 
	}
	SESSION_MAX_POS_VERTICAL_SPEED_DEFAULT_UNIT = Unit{
		Unit: UNIT_METERPERSECOND, 
		Scale: 1000, 
		Offset: 0, 
	}
	SESSION_MAX_NEG_VERTICAL_SPEED_DEFAULT_UNIT = Unit{
		Unit: UNIT_METERPERSECOND, 
		Scale: 1000, 
		Offset: 0, 
	}
	SESSION_MIN_HEART_RATE_DEFAULT_UNIT = Unit{
		Unit: UNIT_COUNTRATE, 
		Scale: 1, 
		Offset: 0, 
	}
	SESSION_TIME_IN_HR_ZONE_DEFAULT_UNIT = Unit{
		Unit: UNIT_SECOND, 
		Scale: 1000, 
		Offset: 0, 
	}
	SESSION_TIME_IN_SPEED_ZONE_DEFAULT_UNIT = Unit{
		Unit: UNIT_SECOND, 
		Scale: 1000, 
		Offset: 0, 
	}
	SESSION_TIME_IN_CADENCE_ZONE_DEFAULT_UNIT = Unit{
		Unit: UNIT_SECOND, 
		Scale: 1000, 
		Offset: 0, 
	}
	SESSION_TIME_IN_POWER_ZONE_DEFAULT_UNIT = Unit{
		Unit: UNIT_SECOND, 
		Scale: 1000, 
		Offset: 0, 
	}
	SESSION_AVG_LAP_TIME_DEFAULT_UNIT = Unit{
		Unit: UNIT_SECOND, 
		Scale: 1000, 
		Offset: 0, 
	}
	SESSION_BEST_LAP_INDEX_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SESSION_MIN_ALTITUDE_DEFAULT_UNIT = Unit{
		Unit: UNIT_METER, 
		Scale: 5, 
		Offset: 500, 
	}
	SESSION_PLAYER_SCORE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SESSION_OPPONENT_SCORE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SESSION_OPPONENT_NAME_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SESSION_STROKE_COUNT_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SESSION_ZONE_COUNT_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SESSION_MAX_BALL_SPEED_DEFAULT_UNIT = Unit{
		Unit: UNIT_METERPERSECOND, 
		Scale: 100, 
		Offset: 0, 
	}
	SESSION_AVG_BALL_SPEED_DEFAULT_UNIT = Unit{
		Unit: UNIT_METERPERSECOND, 
		Scale: 100, 
		Offset: 0, 
	}
	SESSION_AVG_VERTICAL_OSCILLATION_DEFAULT_UNIT = Unit{
		Unit: UNIT_MILLIMETER, 
		Scale: 10, 
		Offset: 0, 
	}
	SESSION_AVG_STANCE_TIME_PERCENT_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 100, 
		Offset: 0, 
	}
	SESSION_AVG_STANCE_TIME_DEFAULT_UNIT = Unit{
		Unit: UNIT_MILLISECOND, 
		Scale: 10, 
		Offset: 0, 
	}
	SESSION_AVG_FRACTIONAL_CADENCE_DEFAULT_UNIT = Unit{
		Unit: UNIT_COUNTRATE, 
		Scale: 128, 
		Offset: 0, 
	}
	SESSION_MAX_FRACTIONAL_CADENCE_DEFAULT_UNIT = Unit{
		Unit: UNIT_COUNTRATE, 
		Scale: 128, 
		Offset: 0, 
	}
	SESSION_TOTAL_FRACTIONAL_CYCLES_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 128, 
		Offset: 0, 
	}
	SESSION_AVG_TOTAL_HEMOGLOBIN_CONC_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 100, 
		Offset: 0, 
	}
	SESSION_MIN_TOTAL_HEMOGLOBIN_CONC_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 100, 
		Offset: 0, 
	}
	SESSION_MAX_TOTAL_HEMOGLOBIN_CONC_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 100, 
		Offset: 0, 
	}
	SESSION_AVG_SATURATED_HEMOGLOBIN_PERCENT_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 10, 
		Offset: 0, 
	}
	SESSION_MIN_SATURATED_HEMOGLOBIN_PERCENT_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 10, 
		Offset: 0, 
	}
	SESSION_MAX_SATURATED_HEMOGLOBIN_PERCENT_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 10, 
		Offset: 0, 
	}
	SESSION_AVG_LEFT_TORQUE_EFFECTIVENESS_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 2, 
		Offset: 0, 
	}
	SESSION_AVG_RIGHT_TORQUE_EFFECTIVENESS_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 2, 
		Offset: 0, 
	}
	SESSION_AVG_LEFT_PEDAL_SMOOTHNESS_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 2, 
		Offset: 0, 
	}
	SESSION_AVG_RIGHT_PEDAL_SMOOTHNESS_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 2, 
		Offset: 0, 
	}
	SESSION_AVG_COMBINED_PEDAL_SMOOTHNESS_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 2, 
		Offset: 0, 
	}
	SESSION_SPORT_INDEX_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SESSION_TIME_STANDING_DEFAULT_UNIT = Unit{
		Unit: UNIT_SECOND, 
		Scale: 1000, 
		Offset: 0, 
	}
	SESSION_STAND_COUNT_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SESSION_AVG_LEFT_PCO_DEFAULT_UNIT = Unit{
		Unit: UNIT_MILLIMETER, 
		Scale: 1, 
		Offset: 0, 
	}
	SESSION_AVG_RIGHT_PCO_DEFAULT_UNIT = Unit{
		Unit: UNIT_MILLIMETER, 
		Scale: 1, 
		Offset: 0, 
	}
	SESSION_AVG_LEFT_POWER_PHASE_DEFAULT_UNIT = Unit{
		Unit: UNIT_DEGREES, 
		Scale: 0, 
		Offset: 0, 
	}
	SESSION_AVG_LEFT_POWER_PHASE_PEAK_DEFAULT_UNIT = Unit{
		Unit: UNIT_DEGREES, 
		Scale: 0, 
		Offset: 0, 
	}
	SESSION_AVG_RIGHT_POWER_PHASE_DEFAULT_UNIT = Unit{
		Unit: UNIT_DEGREES, 
		Scale: 0, 
		Offset: 0, 
	}
	SESSION_AVG_RIGHT_POWER_PHASE_PEAK_DEFAULT_UNIT = Unit{
		Unit: UNIT_DEGREES, 
		Scale: 0, 
		Offset: 0, 
	}
	SESSION_AVG_POWER_POSITION_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SESSION_MAX_POWER_POSITION_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SESSION_AVG_CADENCE_POSITION_DEFAULT_UNIT = Unit{
		Unit: UNIT_COUNTRATE, 
		Scale: 1, 
		Offset: 0, 
	}
	SESSION_MAX_CADENCE_POSITION_DEFAULT_UNIT = Unit{
		Unit: UNIT_COUNTRATE, 
		Scale: 1, 
		Offset: 0, 
	}
	SESSION_ENHANCED_AVG_SPEED_DEFAULT_UNIT = Unit{
		Unit: UNIT_METERPERSECOND, 
		Scale: 1000, 
		Offset: 0, 
	}
	SESSION_ENHANCED_MAX_SPEED_DEFAULT_UNIT = Unit{
		Unit: UNIT_METERPERSECOND, 
		Scale: 1000, 
		Offset: 0, 
	}
	SESSION_ENHANCED_AVG_ALTITUDE_DEFAULT_UNIT = Unit{
		Unit: UNIT_METER, 
		Scale: 5, 
		Offset: 500, 
	}
	SESSION_ENHANCED_MIN_ALTITUDE_DEFAULT_UNIT = Unit{
		Unit: UNIT_METER, 
		Scale: 5, 
		Offset: 500, 
	}
	SESSION_ENHANCED_MAX_ALTITUDE_DEFAULT_UNIT = Unit{
		Unit: UNIT_METER, 
		Scale: 5, 
		Offset: 500, 
	}
	SESSION_AVG_LEV_MOTOR_POWER_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SESSION_MAX_LEV_MOTOR_POWER_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SESSION_LEV_BATTERY_CONSUMPTION_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 2, 
		Offset: 0, 
	}
	LAP_MESSAGE_INDEX_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	LAP_TIMESTAMP_DEFAULT_UNIT = Unit{
		Unit: UNIT_SECOND, 
		Scale: 1, 
		Offset: 0, 
	}
	LAP_EVENT_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	LAP_EVENT_TYPE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	LAP_START_TIME_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	LAP_START_POSITION_LAT_DEFAULT_UNIT = Unit{
		Unit: UNIT_SEMICIRCLES, 
		Scale: 1, 
		Offset: 0, 
	}
	LAP_START_POSITION_LONG_DEFAULT_UNIT = Unit{
		Unit: UNIT_SEMICIRCLES, 
		Scale: 1, 
		Offset: 0, 
	}
	LAP_END_POSITION_LAT_DEFAULT_UNIT = Unit{
		Unit: UNIT_SEMICIRCLES, 
		Scale: 1, 
		Offset: 0, 
	}
	LAP_END_POSITION_LONG_DEFAULT_UNIT = Unit{
		Unit: UNIT_SEMICIRCLES, 
		Scale: 1, 
		Offset: 0, 
	}
	LAP_TOTAL_ELAPSED_TIME_DEFAULT_UNIT = Unit{
		Unit: UNIT_SECOND, 
		Scale: 1000, 
		Offset: 0, 
	}
	LAP_TOTAL_TIMER_TIME_DEFAULT_UNIT = Unit{
		Unit: UNIT_SECOND, 
		Scale: 1000, 
		Offset: 0, 
	}
	LAP_TOTAL_DISTANCE_DEFAULT_UNIT = Unit{
		Unit: UNIT_METER, 
		Scale: 100, 
		Offset: 0, 
	}
	LAP_TOTAL_CYCLES_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	LAP_TOTAL_CALORIES_DEFAULT_UNIT = Unit{
		Unit: UNIT_KILOCALORIE, 
		Scale: 1, 
		Offset: 0, 
	}
	LAP_TOTAL_FAT_CALORIES_DEFAULT_UNIT = Unit{
		Unit: UNIT_KILOCALORIE, 
		Scale: 1, 
		Offset: 0, 
	}
	LAP_AVG_SPEED_DEFAULT_UNIT = Unit{
		Unit: UNIT_METERPERSECOND, 
		Scale: 1000, 
		Offset: 0, 
	}
	LAP_MAX_SPEED_DEFAULT_UNIT = Unit{
		Unit: UNIT_METERPERSECOND, 
		Scale: 1000, 
		Offset: 0, 
	}
	LAP_AVG_HEART_RATE_DEFAULT_UNIT = Unit{
		Unit: UNIT_COUNTRATE, 
		Scale: 1, 
		Offset: 0, 
	}
	LAP_MAX_HEART_RATE_DEFAULT_UNIT = Unit{
		Unit: UNIT_COUNTRATE, 
		Scale: 1, 
		Offset: 0, 
	}
	LAP_AVG_CADENCE_DEFAULT_UNIT = Unit{
		Unit: UNIT_COUNTRATE, 
		Scale: 1, 
		Offset: 0, 
	}
	LAP_MAX_CADENCE_DEFAULT_UNIT = Unit{
		Unit: UNIT_COUNTRATE, 
		Scale: 1, 
		Offset: 0, 
	}
	LAP_AVG_POWER_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	LAP_MAX_POWER_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	LAP_TOTAL_ASCENT_DEFAULT_UNIT = Unit{
		Unit: UNIT_METER, 
		Scale: 1, 
		Offset: 0, 
	}
	LAP_TOTAL_DESCENT_DEFAULT_UNIT = Unit{
		Unit: UNIT_METER, 
		Scale: 1, 
		Offset: 0, 
	}
	LAP_INTENSITY_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	LAP_LAP_TRIGGER_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	LAP_SPORT_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	LAP_EVENT_GROUP_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	LAP_NUM_LENGTHS_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	LAP_NORMALIZED_POWER_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	LAP_LEFT_RIGHT_BALANCE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	LAP_FIRST_LENGTH_INDEX_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	LAP_AVG_STROKE_DISTANCE_DEFAULT_UNIT = Unit{
		Unit: UNIT_METER, 
		Scale: 100, 
		Offset: 0, 
	}
	LAP_SWIM_STROKE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	LAP_SUB_SPORT_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	LAP_NUM_ACTIVE_LENGTHS_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	LAP_TOTAL_WORK_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	LAP_AVG_ALTITUDE_DEFAULT_UNIT = Unit{
		Unit: UNIT_METER, 
		Scale: 5, 
		Offset: 500, 
	}
	LAP_MAX_ALTITUDE_DEFAULT_UNIT = Unit{
		Unit: UNIT_METER, 
		Scale: 5, 
		Offset: 500, 
	}
	LAP_GPS_ACCURACY_DEFAULT_UNIT = Unit{
		Unit: UNIT_METER, 
		Scale: 1, 
		Offset: 0, 
	}
	LAP_AVG_GRADE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 100, 
		Offset: 0, 
	}
	LAP_AVG_POS_GRADE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 100, 
		Offset: 0, 
	}
	LAP_AVG_NEG_GRADE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 100, 
		Offset: 0, 
	}
	LAP_MAX_POS_GRADE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 100, 
		Offset: 0, 
	}
	LAP_MAX_NEG_GRADE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 100, 
		Offset: 0, 
	}
	LAP_AVG_TEMPERATURE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	LAP_MAX_TEMPERATURE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	LAP_TOTAL_MOVING_TIME_DEFAULT_UNIT = Unit{
		Unit: UNIT_SECOND, 
		Scale: 1000, 
		Offset: 0, 
	}
	LAP_AVG_POS_VERTICAL_SPEED_DEFAULT_UNIT = Unit{
		Unit: UNIT_METERPERSECOND, 
		Scale: 1000, 
		Offset: 0, 
	}
	LAP_AVG_NEG_VERTICAL_SPEED_DEFAULT_UNIT = Unit{
		Unit: UNIT_METERPERSECOND, 
		Scale: 1000, 
		Offset: 0, 
	}
	LAP_MAX_POS_VERTICAL_SPEED_DEFAULT_UNIT = Unit{
		Unit: UNIT_METERPERSECOND, 
		Scale: 1000, 
		Offset: 0, 
	}
	LAP_MAX_NEG_VERTICAL_SPEED_DEFAULT_UNIT = Unit{
		Unit: UNIT_METERPERSECOND, 
		Scale: 1000, 
		Offset: 0, 
	}
	LAP_TIME_IN_HR_ZONE_DEFAULT_UNIT = Unit{
		Unit: UNIT_SECOND, 
		Scale: 1000, 
		Offset: 0, 
	}
	LAP_TIME_IN_SPEED_ZONE_DEFAULT_UNIT = Unit{
		Unit: UNIT_SECOND, 
		Scale: 1000, 
		Offset: 0, 
	}
	LAP_TIME_IN_CADENCE_ZONE_DEFAULT_UNIT = Unit{
		Unit: UNIT_SECOND, 
		Scale: 1000, 
		Offset: 0, 
	}
	LAP_TIME_IN_POWER_ZONE_DEFAULT_UNIT = Unit{
		Unit: UNIT_SECOND, 
		Scale: 1000, 
		Offset: 0, 
	}
	LAP_REPETITION_NUM_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	LAP_MIN_ALTITUDE_DEFAULT_UNIT = Unit{
		Unit: UNIT_METER, 
		Scale: 5, 
		Offset: 500, 
	}
	LAP_MIN_HEART_RATE_DEFAULT_UNIT = Unit{
		Unit: UNIT_COUNTRATE, 
		Scale: 1, 
		Offset: 0, 
	}
	LAP_WKT_STEP_INDEX_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	LAP_OPPONENT_SCORE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	LAP_STROKE_COUNT_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	LAP_ZONE_COUNT_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	LAP_AVG_VERTICAL_OSCILLATION_DEFAULT_UNIT = Unit{
		Unit: UNIT_MILLIMETER, 
		Scale: 10, 
		Offset: 0, 
	}
	LAP_AVG_STANCE_TIME_PERCENT_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 100, 
		Offset: 0, 
	}
	LAP_AVG_STANCE_TIME_DEFAULT_UNIT = Unit{
		Unit: UNIT_MILLISECOND, 
		Scale: 10, 
		Offset: 0, 
	}
	LAP_AVG_FRACTIONAL_CADENCE_DEFAULT_UNIT = Unit{
		Unit: UNIT_COUNTRATE, 
		Scale: 128, 
		Offset: 0, 
	}
	LAP_MAX_FRACTIONAL_CADENCE_DEFAULT_UNIT = Unit{
		Unit: UNIT_COUNTRATE, 
		Scale: 128, 
		Offset: 0, 
	}
	LAP_TOTAL_FRACTIONAL_CYCLES_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 128, 
		Offset: 0, 
	}
	LAP_PLAYER_SCORE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	LAP_AVG_TOTAL_HEMOGLOBIN_CONC_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 100, 
		Offset: 0, 
	}
	LAP_MIN_TOTAL_HEMOGLOBIN_CONC_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 100, 
		Offset: 0, 
	}
	LAP_MAX_TOTAL_HEMOGLOBIN_CONC_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 100, 
		Offset: 0, 
	}
	LAP_AVG_SATURATED_HEMOGLOBIN_PERCENT_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 10, 
		Offset: 0, 
	}
	LAP_MIN_SATURATED_HEMOGLOBIN_PERCENT_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 10, 
		Offset: 0, 
	}
	LAP_MAX_SATURATED_HEMOGLOBIN_PERCENT_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 10, 
		Offset: 0, 
	}
	LAP_AVG_LEFT_TORQUE_EFFECTIVENESS_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 2, 
		Offset: 0, 
	}
	LAP_AVG_RIGHT_TORQUE_EFFECTIVENESS_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 2, 
		Offset: 0, 
	}
	LAP_AVG_LEFT_PEDAL_SMOOTHNESS_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 2, 
		Offset: 0, 
	}
	LAP_AVG_RIGHT_PEDAL_SMOOTHNESS_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 2, 
		Offset: 0, 
	}
	LAP_AVG_COMBINED_PEDAL_SMOOTHNESS_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 2, 
		Offset: 0, 
	}
	LAP_TIME_STANDING_DEFAULT_UNIT = Unit{
		Unit: UNIT_SECOND, 
		Scale: 1000, 
		Offset: 0, 
	}
	LAP_STAND_COUNT_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	LAP_AVG_LEFT_PCO_DEFAULT_UNIT = Unit{
		Unit: UNIT_MILLIMETER, 
		Scale: 1, 
		Offset: 0, 
	}
	LAP_AVG_RIGHT_PCO_DEFAULT_UNIT = Unit{
		Unit: UNIT_MILLIMETER, 
		Scale: 1, 
		Offset: 0, 
	}
	LAP_AVG_LEFT_POWER_PHASE_DEFAULT_UNIT = Unit{
		Unit: UNIT_DEGREES, 
		Scale: 0, 
		Offset: 0, 
	}
	LAP_AVG_LEFT_POWER_PHASE_PEAK_DEFAULT_UNIT = Unit{
		Unit: UNIT_DEGREES, 
		Scale: 0, 
		Offset: 0, 
	}
	LAP_AVG_RIGHT_POWER_PHASE_DEFAULT_UNIT = Unit{
		Unit: UNIT_DEGREES, 
		Scale: 0, 
		Offset: 0, 
	}
	LAP_AVG_RIGHT_POWER_PHASE_PEAK_DEFAULT_UNIT = Unit{
		Unit: UNIT_DEGREES, 
		Scale: 0, 
		Offset: 0, 
	}
	LAP_AVG_POWER_POSITION_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	LAP_MAX_POWER_POSITION_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	LAP_AVG_CADENCE_POSITION_DEFAULT_UNIT = Unit{
		Unit: UNIT_COUNTRATE, 
		Scale: 1, 
		Offset: 0, 
	}
	LAP_MAX_CADENCE_POSITION_DEFAULT_UNIT = Unit{
		Unit: UNIT_COUNTRATE, 
		Scale: 1, 
		Offset: 0, 
	}
	LAP_ENHANCED_AVG_SPEED_DEFAULT_UNIT = Unit{
		Unit: UNIT_METERPERSECOND, 
		Scale: 1000, 
		Offset: 0, 
	}
	LAP_ENHANCED_MAX_SPEED_DEFAULT_UNIT = Unit{
		Unit: UNIT_METERPERSECOND, 
		Scale: 1000, 
		Offset: 0, 
	}
	LAP_ENHANCED_AVG_ALTITUDE_DEFAULT_UNIT = Unit{
		Unit: UNIT_METER, 
		Scale: 5, 
		Offset: 500, 
	}
	LAP_ENHANCED_MIN_ALTITUDE_DEFAULT_UNIT = Unit{
		Unit: UNIT_METER, 
		Scale: 5, 
		Offset: 500, 
	}
	LAP_ENHANCED_MAX_ALTITUDE_DEFAULT_UNIT = Unit{
		Unit: UNIT_METER, 
		Scale: 5, 
		Offset: 500, 
	}
	LAP_AVG_LEV_MOTOR_POWER_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	LAP_MAX_LEV_MOTOR_POWER_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	LAP_LEV_BATTERY_CONSUMPTION_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 2, 
		Offset: 0, 
	}
	LENGTH_MESSAGE_INDEX_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	LENGTH_TIMESTAMP_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	LENGTH_EVENT_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	LENGTH_EVENT_TYPE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	LENGTH_START_TIME_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	LENGTH_TOTAL_ELAPSED_TIME_DEFAULT_UNIT = Unit{
		Unit: UNIT_SECOND, 
		Scale: 1000, 
		Offset: 0, 
	}
	LENGTH_TOTAL_TIMER_TIME_DEFAULT_UNIT = Unit{
		Unit: UNIT_SECOND, 
		Scale: 1000, 
		Offset: 0, 
	}
	LENGTH_TOTAL_STROKES_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	LENGTH_AVG_SPEED_DEFAULT_UNIT = Unit{
		Unit: UNIT_METERPERSECOND, 
		Scale: 1000, 
		Offset: 0, 
	}
	LENGTH_SWIM_STROKE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	LENGTH_AVG_SWIMMING_CADENCE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	LENGTH_EVENT_GROUP_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	LENGTH_TOTAL_CALORIES_DEFAULT_UNIT = Unit{
		Unit: UNIT_KILOCALORIE, 
		Scale: 1, 
		Offset: 0, 
	}
	LENGTH_LENGTH_TYPE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	LENGTH_PLAYER_SCORE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	LENGTH_OPPONENT_SCORE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	LENGTH_STROKE_COUNT_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	LENGTH_ZONE_COUNT_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	RECORD_TIMESTAMP_DEFAULT_UNIT = Unit{
		Unit: UNIT_SECOND, 
		Scale: 1, 
		Offset: 0, 
	}
	RECORD_POSITION_LAT_DEFAULT_UNIT = Unit{
		Unit: UNIT_SEMICIRCLES, 
		Scale: 1, 
		Offset: 0, 
	}
	RECORD_POSITION_LONG_DEFAULT_UNIT = Unit{
		Unit: UNIT_SEMICIRCLES, 
		Scale: 1, 
		Offset: 0, 
	}
	RECORD_ALTITUDE_DEFAULT_UNIT = Unit{
		Unit: UNIT_METER, 
		Scale: 5, 
		Offset: 500, 
	}
	RECORD_HEART_RATE_DEFAULT_UNIT = Unit{
		Unit: UNIT_COUNTRATE, 
		Scale: 1, 
		Offset: 0, 
	}
	RECORD_CADENCE_DEFAULT_UNIT = Unit{
		Unit: UNIT_COUNTRATE, 
		Scale: 1, 
		Offset: 0, 
	}
	RECORD_DISTANCE_DEFAULT_UNIT = Unit{
		Unit: UNIT_METER, 
		Scale: 100, 
		Offset: 0, 
	}
	RECORD_SPEED_DEFAULT_UNIT = Unit{
		Unit: UNIT_METERPERSECOND, 
		Scale: 1000, 
		Offset: 0, 
	}
	RECORD_POWER_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	RECORD_COMPRESSED_SPEED_DISTANCE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 0, 
		Offset: 0, 
	}
	RECORD_GRADE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 100, 
		Offset: 0, 
	}
	RECORD_RESISTANCE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	RECORD_TIME_FROM_COURSE_DEFAULT_UNIT = Unit{
		Unit: UNIT_SECOND, 
		Scale: 1000, 
		Offset: 0, 
	}
	RECORD_CYCLE_LENGTH_DEFAULT_UNIT = Unit{
		Unit: UNIT_METER, 
		Scale: 100, 
		Offset: 0, 
	}
	RECORD_TEMPERATURE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	RECORD_SPEED_1S_DEFAULT_UNIT = Unit{
		Unit: UNIT_METERPERSECOND, 
		Scale: 16, 
		Offset: 0, 
	}
	RECORD_CYCLES_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	RECORD_TOTAL_CYCLES_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	RECORD_COMPRESSED_ACCUMULATED_POWER_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	RECORD_ACCUMULATED_POWER_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	RECORD_LEFT_RIGHT_BALANCE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	RECORD_GPS_ACCURACY_DEFAULT_UNIT = Unit{
		Unit: UNIT_METER, 
		Scale: 1, 
		Offset: 0, 
	}
	RECORD_VERTICAL_SPEED_DEFAULT_UNIT = Unit{
		Unit: UNIT_METERPERSECOND, 
		Scale: 1000, 
		Offset: 0, 
	}
	RECORD_CALORIES_DEFAULT_UNIT = Unit{
		Unit: UNIT_KILOCALORIE, 
		Scale: 1, 
		Offset: 0, 
	}
	RECORD_VERTICAL_OSCILLATION_DEFAULT_UNIT = Unit{
		Unit: UNIT_MILLIMETER, 
		Scale: 10, 
		Offset: 0, 
	}
	RECORD_STANCE_TIME_PERCENT_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 100, 
		Offset: 0, 
	}
	RECORD_STANCE_TIME_DEFAULT_UNIT = Unit{
		Unit: UNIT_MILLISECOND, 
		Scale: 10, 
		Offset: 0, 
	}
	RECORD_ACTIVITY_TYPE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	RECORD_LEFT_TORQUE_EFFECTIVENESS_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 2, 
		Offset: 0, 
	}
	RECORD_RIGHT_TORQUE_EFFECTIVENESS_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 2, 
		Offset: 0, 
	}
	RECORD_LEFT_PEDAL_SMOOTHNESS_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 2, 
		Offset: 0, 
	}
	RECORD_RIGHT_PEDAL_SMOOTHNESS_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 2, 
		Offset: 0, 
	}
	RECORD_COMBINED_PEDAL_SMOOTHNESS_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 2, 
		Offset: 0, 
	}
	RECORD_TIME128_DEFAULT_UNIT = Unit{
		Unit: UNIT_SECOND, 
		Scale: 128, 
		Offset: 0, 
	}
	RECORD_STROKE_TYPE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	RECORD_ZONE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	RECORD_BALL_SPEED_DEFAULT_UNIT = Unit{
		Unit: UNIT_METERPERSECOND, 
		Scale: 100, 
		Offset: 0, 
	}
	RECORD_CADENCE256_DEFAULT_UNIT = Unit{
		Unit: UNIT_COUNTRATE, 
		Scale: 256, 
		Offset: 0, 
	}
	RECORD_FRACTIONAL_CADENCE_DEFAULT_UNIT = Unit{
		Unit: UNIT_COUNTRATE, 
		Scale: 128, 
		Offset: 0, 
	}
	RECORD_TOTAL_HEMOGLOBIN_CONC_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 100, 
		Offset: 0, 
	}
	RECORD_TOTAL_HEMOGLOBIN_CONC_MIN_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 100, 
		Offset: 0, 
	}
	RECORD_TOTAL_HEMOGLOBIN_CONC_MAX_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 100, 
		Offset: 0, 
	}
	RECORD_SATURATED_HEMOGLOBIN_PERCENT_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 10, 
		Offset: 0, 
	}
	RECORD_SATURATED_HEMOGLOBIN_PERCENT_MIN_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 10, 
		Offset: 0, 
	}
	RECORD_SATURATED_HEMOGLOBIN_PERCENT_MAX_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 10, 
		Offset: 0, 
	}
	RECORD_DEVICE_INDEX_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	RECORD_LEFT_PCO_DEFAULT_UNIT = Unit{
		Unit: UNIT_MILLIMETER, 
		Scale: 1, 
		Offset: 0, 
	}
	RECORD_RIGHT_PCO_DEFAULT_UNIT = Unit{
		Unit: UNIT_MILLIMETER, 
		Scale: 1, 
		Offset: 0, 
	}
	RECORD_LEFT_POWER_PHASE_DEFAULT_UNIT = Unit{
		Unit: UNIT_DEGREES, 
		Scale: 0, 
		Offset: 0, 
	}
	RECORD_LEFT_POWER_PHASE_PEAK_DEFAULT_UNIT = Unit{
		Unit: UNIT_DEGREES, 
		Scale: 0, 
		Offset: 0, 
	}
	RECORD_RIGHT_POWER_PHASE_DEFAULT_UNIT = Unit{
		Unit: UNIT_DEGREES, 
		Scale: 0, 
		Offset: 0, 
	}
	RECORD_RIGHT_POWER_PHASE_PEAK_DEFAULT_UNIT = Unit{
		Unit: UNIT_DEGREES, 
		Scale: 0, 
		Offset: 0, 
	}
	RECORD_ENHANCED_SPEED_DEFAULT_UNIT = Unit{
		Unit: UNIT_METERPERSECOND, 
		Scale: 1000, 
		Offset: 0, 
	}
	RECORD_ENHANCED_ALTITUDE_DEFAULT_UNIT = Unit{
		Unit: UNIT_METER, 
		Scale: 5, 
		Offset: 500, 
	}
	RECORD_BATTERY_SOC_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 2, 
		Offset: 0, 
	}
	RECORD_MOTOR_POWER_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	EVENT_TIMESTAMP_DEFAULT_UNIT = Unit{
		Unit: UNIT_SECOND, 
		Scale: 1, 
		Offset: 0, 
	}
	EVENT_EVENT_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	EVENT_EVENT_TYPE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	EVENT_DATA16_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	EVENT_DATA_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	EVENT_EVENT_GROUP_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	EVENT_SCORE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	EVENT_OPPONENT_SCORE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	EVENT_FRONT_GEAR_NUM_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	EVENT_FRONT_GEAR_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	EVENT_REAR_GEAR_NUM_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	EVENT_REAR_GEAR_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	EVENT_DEVICE_INDEX_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	DEVICE_INFO_TIMESTAMP_DEFAULT_UNIT = Unit{
		Unit: UNIT_SECOND, 
		Scale: 1, 
		Offset: 0, 
	}
	DEVICE_INFO_DEVICE_INDEX_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	DEVICE_INFO_DEVICE_TYPE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	DEVICE_INFO_MANUFACTURER_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	DEVICE_INFO_SERIAL_NUMBER_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	DEVICE_INFO_PRODUCT_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	DEVICE_INFO_SOFTWARE_VERSION_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 100, 
		Offset: 0, 
	}
	DEVICE_INFO_HARDWARE_VERSION_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	DEVICE_INFO_CUM_OPERATING_TIME_DEFAULT_UNIT = Unit{
		Unit: UNIT_SECOND, 
		Scale: 1, 
		Offset: 0, 
	}
	DEVICE_INFO_BATTERY_VOLTAGE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 256, 
		Offset: 0, 
	}
	DEVICE_INFO_BATTERY_STATUS_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	DEVICE_INFO_SENSOR_POSITION_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	DEVICE_INFO_DESCRIPTOR_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	DEVICE_INFO_ANT_TRANSMISSION_TYPE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	DEVICE_INFO_ANT_DEVICE_NUMBER_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	DEVICE_INFO_ANT_NETWORK_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	DEVICE_INFO_SOURCE_TYPE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	TRAINING_FILE_TIMESTAMP_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	TRAINING_FILE_TYPE__DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	TRAINING_FILE_MANUFACTURER_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	TRAINING_FILE_PRODUCT_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	TRAINING_FILE_SERIAL_NUMBER_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	TRAINING_FILE_TIME_CREATED_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	HRV_TIME_DEFAULT_UNIT = Unit{
		Unit: UNIT_SECOND, 
		Scale: 1000, 
		Offset: 0, 
	}
	COURSE_SPORT_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	COURSE_NAME_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	COURSE_CAPABILITIES_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	COURSE_POINT_MESSAGE_INDEX_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	COURSE_POINT_TIMESTAMP_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	COURSE_POINT_POSITION_LAT_DEFAULT_UNIT = Unit{
		Unit: UNIT_SEMICIRCLES, 
		Scale: 1, 
		Offset: 0, 
	}
	COURSE_POINT_POSITION_LONG_DEFAULT_UNIT = Unit{
		Unit: UNIT_SEMICIRCLES, 
		Scale: 1, 
		Offset: 0, 
	}
	COURSE_POINT_DISTANCE_DEFAULT_UNIT = Unit{
		Unit: UNIT_METER, 
		Scale: 100, 
		Offset: 0, 
	}
	COURSE_POINT_TYPE__DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	COURSE_POINT_NAME_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	COURSE_POINT_FAVORITE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_ID_NAME_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_ID_UUID_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_ID_SPORT_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_ID_ENABLED_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_ID_USER_PROFILE_PRIMARY_KEY_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_ID_DEVICE_ID_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_ID_DEFAULT_RACE_LEADER_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_ID_DELETE_STATUS_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_ID_SELECTION_TYPE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_LEADERBOARD_ENTRY_MESSAGE_INDEX_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_LEADERBOARD_ENTRY_NAME_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_LEADERBOARD_ENTRY_TYPE__DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_LEADERBOARD_ENTRY_GROUP_PRIMARY_KEY_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_LEADERBOARD_ENTRY_ACTIVITY_ID_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_LEADERBOARD_ENTRY_SEGMENT_TIME_DEFAULT_UNIT = Unit{
		Unit: UNIT_SECOND, 
		Scale: 1000, 
		Offset: 0, 
	}
	SEGMENT_LEADERBOARD_ENTRY_ACTIVITY_ID_STRING_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_POINT_MESSAGE_INDEX_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_POINT_POSITION_LAT_DEFAULT_UNIT = Unit{
		Unit: UNIT_SEMICIRCLES, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_POINT_POSITION_LONG_DEFAULT_UNIT = Unit{
		Unit: UNIT_SEMICIRCLES, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_POINT_DISTANCE_DEFAULT_UNIT = Unit{
		Unit: UNIT_METER, 
		Scale: 100, 
		Offset: 0, 
	}
	SEGMENT_POINT_ALTITUDE_DEFAULT_UNIT = Unit{
		Unit: UNIT_METER, 
		Scale: 5, 
		Offset: 500, 
	}
	SEGMENT_POINT_LEADER_TIME_DEFAULT_UNIT = Unit{
		Unit: UNIT_SECOND, 
		Scale: 1000, 
		Offset: 0, 
	}
	SEGMENT_LAP_MESSAGE_INDEX_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_LAP_TIMESTAMP_DEFAULT_UNIT = Unit{
		Unit: UNIT_SECOND, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_LAP_EVENT_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_LAP_EVENT_TYPE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_LAP_START_TIME_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_LAP_START_POSITION_LAT_DEFAULT_UNIT = Unit{
		Unit: UNIT_SEMICIRCLES, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_LAP_START_POSITION_LONG_DEFAULT_UNIT = Unit{
		Unit: UNIT_SEMICIRCLES, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_LAP_END_POSITION_LAT_DEFAULT_UNIT = Unit{
		Unit: UNIT_SEMICIRCLES, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_LAP_END_POSITION_LONG_DEFAULT_UNIT = Unit{
		Unit: UNIT_SEMICIRCLES, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_LAP_TOTAL_ELAPSED_TIME_DEFAULT_UNIT = Unit{
		Unit: UNIT_SECOND, 
		Scale: 1000, 
		Offset: 0, 
	}
	SEGMENT_LAP_TOTAL_TIMER_TIME_DEFAULT_UNIT = Unit{
		Unit: UNIT_SECOND, 
		Scale: 1000, 
		Offset: 0, 
	}
	SEGMENT_LAP_TOTAL_DISTANCE_DEFAULT_UNIT = Unit{
		Unit: UNIT_METER, 
		Scale: 100, 
		Offset: 0, 
	}
	SEGMENT_LAP_TOTAL_CYCLES_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_LAP_TOTAL_CALORIES_DEFAULT_UNIT = Unit{
		Unit: UNIT_KILOCALORIE, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_LAP_TOTAL_FAT_CALORIES_DEFAULT_UNIT = Unit{
		Unit: UNIT_KILOCALORIE, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_LAP_AVG_SPEED_DEFAULT_UNIT = Unit{
		Unit: UNIT_METERPERSECOND, 
		Scale: 1000, 
		Offset: 0, 
	}
	SEGMENT_LAP_MAX_SPEED_DEFAULT_UNIT = Unit{
		Unit: UNIT_METERPERSECOND, 
		Scale: 1000, 
		Offset: 0, 
	}
	SEGMENT_LAP_AVG_HEART_RATE_DEFAULT_UNIT = Unit{
		Unit: UNIT_COUNTRATE, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_LAP_MAX_HEART_RATE_DEFAULT_UNIT = Unit{
		Unit: UNIT_COUNTRATE, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_LAP_AVG_CADENCE_DEFAULT_UNIT = Unit{
		Unit: UNIT_COUNTRATE, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_LAP_MAX_CADENCE_DEFAULT_UNIT = Unit{
		Unit: UNIT_COUNTRATE, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_LAP_AVG_POWER_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_LAP_MAX_POWER_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_LAP_TOTAL_ASCENT_DEFAULT_UNIT = Unit{
		Unit: UNIT_METER, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_LAP_TOTAL_DESCENT_DEFAULT_UNIT = Unit{
		Unit: UNIT_METER, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_LAP_SPORT_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_LAP_EVENT_GROUP_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_LAP_NEC_LAT_DEFAULT_UNIT = Unit{
		Unit: UNIT_SEMICIRCLES, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_LAP_NEC_LONG_DEFAULT_UNIT = Unit{
		Unit: UNIT_SEMICIRCLES, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_LAP_SWC_LAT_DEFAULT_UNIT = Unit{
		Unit: UNIT_SEMICIRCLES, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_LAP_SWC_LONG_DEFAULT_UNIT = Unit{
		Unit: UNIT_SEMICIRCLES, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_LAP_NAME_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_LAP_NORMALIZED_POWER_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_LAP_LEFT_RIGHT_BALANCE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_LAP_SUB_SPORT_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_LAP_TOTAL_WORK_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_LAP_AVG_ALTITUDE_DEFAULT_UNIT = Unit{
		Unit: UNIT_METER, 
		Scale: 5, 
		Offset: 500, 
	}
	SEGMENT_LAP_MAX_ALTITUDE_DEFAULT_UNIT = Unit{
		Unit: UNIT_METER, 
		Scale: 5, 
		Offset: 500, 
	}
	SEGMENT_LAP_GPS_ACCURACY_DEFAULT_UNIT = Unit{
		Unit: UNIT_METER, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_LAP_AVG_GRADE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 100, 
		Offset: 0, 
	}
	SEGMENT_LAP_AVG_POS_GRADE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 100, 
		Offset: 0, 
	}
	SEGMENT_LAP_AVG_NEG_GRADE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 100, 
		Offset: 0, 
	}
	SEGMENT_LAP_MAX_POS_GRADE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 100, 
		Offset: 0, 
	}
	SEGMENT_LAP_MAX_NEG_GRADE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 100, 
		Offset: 0, 
	}
	SEGMENT_LAP_AVG_TEMPERATURE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_LAP_MAX_TEMPERATURE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_LAP_TOTAL_MOVING_TIME_DEFAULT_UNIT = Unit{
		Unit: UNIT_SECOND, 
		Scale: 1000, 
		Offset: 0, 
	}
	SEGMENT_LAP_AVG_POS_VERTICAL_SPEED_DEFAULT_UNIT = Unit{
		Unit: UNIT_METERPERSECOND, 
		Scale: 1000, 
		Offset: 0, 
	}
	SEGMENT_LAP_AVG_NEG_VERTICAL_SPEED_DEFAULT_UNIT = Unit{
		Unit: UNIT_METERPERSECOND, 
		Scale: 1000, 
		Offset: 0, 
	}
	SEGMENT_LAP_MAX_POS_VERTICAL_SPEED_DEFAULT_UNIT = Unit{
		Unit: UNIT_METERPERSECOND, 
		Scale: 1000, 
		Offset: 0, 
	}
	SEGMENT_LAP_MAX_NEG_VERTICAL_SPEED_DEFAULT_UNIT = Unit{
		Unit: UNIT_METERPERSECOND, 
		Scale: 1000, 
		Offset: 0, 
	}
	SEGMENT_LAP_TIME_IN_HR_ZONE_DEFAULT_UNIT = Unit{
		Unit: UNIT_SECOND, 
		Scale: 1000, 
		Offset: 0, 
	}
	SEGMENT_LAP_TIME_IN_SPEED_ZONE_DEFAULT_UNIT = Unit{
		Unit: UNIT_SECOND, 
		Scale: 1000, 
		Offset: 0, 
	}
	SEGMENT_LAP_TIME_IN_CADENCE_ZONE_DEFAULT_UNIT = Unit{
		Unit: UNIT_SECOND, 
		Scale: 1000, 
		Offset: 0, 
	}
	SEGMENT_LAP_TIME_IN_POWER_ZONE_DEFAULT_UNIT = Unit{
		Unit: UNIT_SECOND, 
		Scale: 1000, 
		Offset: 0, 
	}
	SEGMENT_LAP_REPETITION_NUM_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_LAP_MIN_ALTITUDE_DEFAULT_UNIT = Unit{
		Unit: UNIT_METER, 
		Scale: 5, 
		Offset: 500, 
	}
	SEGMENT_LAP_MIN_HEART_RATE_DEFAULT_UNIT = Unit{
		Unit: UNIT_COUNTRATE, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_LAP_ACTIVE_TIME_DEFAULT_UNIT = Unit{
		Unit: UNIT_SECOND, 
		Scale: 1000, 
		Offset: 0, 
	}
	SEGMENT_LAP_WKT_STEP_INDEX_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_LAP_SPORT_EVENT_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_LAP_AVG_LEFT_TORQUE_EFFECTIVENESS_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 2, 
		Offset: 0, 
	}
	SEGMENT_LAP_AVG_RIGHT_TORQUE_EFFECTIVENESS_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 2, 
		Offset: 0, 
	}
	SEGMENT_LAP_AVG_LEFT_PEDAL_SMOOTHNESS_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 2, 
		Offset: 0, 
	}
	SEGMENT_LAP_AVG_RIGHT_PEDAL_SMOOTHNESS_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 2, 
		Offset: 0, 
	}
	SEGMENT_LAP_AVG_COMBINED_PEDAL_SMOOTHNESS_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 2, 
		Offset: 0, 
	}
	SEGMENT_LAP_STATUS_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_LAP_UUID_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_LAP_AVG_FRACTIONAL_CADENCE_DEFAULT_UNIT = Unit{
		Unit: UNIT_COUNTRATE, 
		Scale: 128, 
		Offset: 0, 
	}
	SEGMENT_LAP_MAX_FRACTIONAL_CADENCE_DEFAULT_UNIT = Unit{
		Unit: UNIT_COUNTRATE, 
		Scale: 128, 
		Offset: 0, 
	}
	SEGMENT_LAP_TOTAL_FRACTIONAL_CYCLES_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 128, 
		Offset: 0, 
	}
	SEGMENT_LAP_FRONT_GEAR_SHIFT_COUNT_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_LAP_REAR_GEAR_SHIFT_COUNT_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_LAP_TIME_STANDING_DEFAULT_UNIT = Unit{
		Unit: UNIT_SECOND, 
		Scale: 1000, 
		Offset: 0, 
	}
	SEGMENT_LAP_STAND_COUNT_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_LAP_AVG_LEFT_PCO_DEFAULT_UNIT = Unit{
		Unit: UNIT_MILLIMETER, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_LAP_AVG_RIGHT_PCO_DEFAULT_UNIT = Unit{
		Unit: UNIT_MILLIMETER, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_LAP_AVG_LEFT_POWER_PHASE_DEFAULT_UNIT = Unit{
		Unit: UNIT_DEGREES, 
		Scale: 0, 
		Offset: 0, 
	}
	SEGMENT_LAP_AVG_LEFT_POWER_PHASE_PEAK_DEFAULT_UNIT = Unit{
		Unit: UNIT_DEGREES, 
		Scale: 0, 
		Offset: 0, 
	}
	SEGMENT_LAP_AVG_RIGHT_POWER_PHASE_DEFAULT_UNIT = Unit{
		Unit: UNIT_DEGREES, 
		Scale: 0, 
		Offset: 0, 
	}
	SEGMENT_LAP_AVG_RIGHT_POWER_PHASE_PEAK_DEFAULT_UNIT = Unit{
		Unit: UNIT_DEGREES, 
		Scale: 0, 
		Offset: 0, 
	}
	SEGMENT_LAP_AVG_POWER_POSITION_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_LAP_MAX_POWER_POSITION_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_LAP_AVG_CADENCE_POSITION_DEFAULT_UNIT = Unit{
		Unit: UNIT_COUNTRATE, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_LAP_MAX_CADENCE_POSITION_DEFAULT_UNIT = Unit{
		Unit: UNIT_COUNTRATE, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_FILE_MESSAGE_INDEX_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_FILE_FILE_UUID_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_FILE_ENABLED_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_FILE_USER_PROFILE_PRIMARY_KEY_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_FILE_LEADER_TYPE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_FILE_LEADER_GROUP_PRIMARY_KEY_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_FILE_LEADER_ACTIVITY_ID_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SEGMENT_FILE_LEADER_ACTIVITY_ID_STRING_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	WORKOUT_SPORT_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	WORKOUT_CAPABILITIES_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	WORKOUT_NUM_VALID_STEPS_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	WORKOUT_WKT_NAME_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	WORKOUT_STEP_MESSAGE_INDEX_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	WORKOUT_STEP_WKT_STEP_NAME_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	WORKOUT_STEP_DURATION_TYPE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	WORKOUT_STEP_DURATION_VALUE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	WORKOUT_STEP_TARGET_TYPE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	WORKOUT_STEP_TARGET_VALUE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	WORKOUT_STEP_CUSTOM_TARGET_VALUE_LOW_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	WORKOUT_STEP_CUSTOM_TARGET_VALUE_HIGH_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	WORKOUT_STEP_INTENSITY_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SCHEDULE_MANUFACTURER_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SCHEDULE_PRODUCT_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SCHEDULE_SERIAL_NUMBER_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SCHEDULE_TIME_CREATED_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SCHEDULE_COMPLETED_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SCHEDULE_TYPE__DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	SCHEDULE_SCHEDULED_TIME_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	TOTALS_MESSAGE_INDEX_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	TOTALS_TIMESTAMP_DEFAULT_UNIT = Unit{
		Unit: UNIT_SECOND, 
		Scale: 1, 
		Offset: 0, 
	}
	TOTALS_TIMER_TIME_DEFAULT_UNIT = Unit{
		Unit: UNIT_SECOND, 
		Scale: 1, 
		Offset: 0, 
	}
	TOTALS_DISTANCE_DEFAULT_UNIT = Unit{
		Unit: UNIT_METER, 
		Scale: 1, 
		Offset: 0, 
	}
	TOTALS_CALORIES_DEFAULT_UNIT = Unit{
		Unit: UNIT_KILOCALORIE, 
		Scale: 1, 
		Offset: 0, 
	}
	TOTALS_SPORT_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	TOTALS_ELAPSED_TIME_DEFAULT_UNIT = Unit{
		Unit: UNIT_SECOND, 
		Scale: 1, 
		Offset: 0, 
	}
	TOTALS_SESSIONS_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	TOTALS_ACTIVE_TIME_DEFAULT_UNIT = Unit{
		Unit: UNIT_SECOND, 
		Scale: 1, 
		Offset: 0, 
	}
	TOTALS_SPORT_INDEX_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	WEIGHT_SCALE_TIMESTAMP_DEFAULT_UNIT = Unit{
		Unit: UNIT_SECOND, 
		Scale: 1, 
		Offset: 0, 
	}
	WEIGHT_SCALE_WEIGHT_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 100, 
		Offset: 0, 
	}
	WEIGHT_SCALE_PERCENT_FAT_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 100, 
		Offset: 0, 
	}
	WEIGHT_SCALE_PERCENT_HYDRATION_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 100, 
		Offset: 0, 
	}
	WEIGHT_SCALE_VISCERAL_FAT_MASS_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 100, 
		Offset: 0, 
	}
	WEIGHT_SCALE_BONE_MASS_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 100, 
		Offset: 0, 
	}
	WEIGHT_SCALE_MUSCLE_MASS_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 100, 
		Offset: 0, 
	}
	WEIGHT_SCALE_BASAL_MET_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 4, 
		Offset: 0, 
	}
	WEIGHT_SCALE_PHYSIQUE_RATING_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	WEIGHT_SCALE_ACTIVE_MET_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 4, 
		Offset: 0, 
	}
	WEIGHT_SCALE_METABOLIC_AGE_DEFAULT_UNIT = Unit{
		Unit: UNIT_YEAR, 
		Scale: 1, 
		Offset: 0, 
	}
	WEIGHT_SCALE_VISCERAL_FAT_RATING_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	WEIGHT_SCALE_USER_PROFILE_INDEX_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	BLOOD_PRESSURE_TIMESTAMP_DEFAULT_UNIT = Unit{
		Unit: UNIT_SECOND, 
		Scale: 1, 
		Offset: 0, 
	}
	BLOOD_PRESSURE_SYSTOLIC_PRESSURE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	BLOOD_PRESSURE_DIASTOLIC_PRESSURE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	BLOOD_PRESSURE_MEAN_ARTERIAL_PRESSURE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	BLOOD_PRESSURE_MAP_3_SAMPLE_MEAN_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	BLOOD_PRESSURE_MAP_MORNING_VALUES_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	BLOOD_PRESSURE_MAP_EVENING_VALUES_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	BLOOD_PRESSURE_HEART_RATE_DEFAULT_UNIT = Unit{
		Unit: UNIT_COUNTRATE, 
		Scale: 1, 
		Offset: 0, 
	}
	BLOOD_PRESSURE_HEART_RATE_TYPE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	BLOOD_PRESSURE_STATUS_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	BLOOD_PRESSURE_USER_PROFILE_INDEX_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	MONITORING_INFO_TIMESTAMP_DEFAULT_UNIT = Unit{
		Unit: UNIT_SECOND, 
		Scale: 1, 
		Offset: 0, 
	}
	MONITORING_INFO_LOCAL_TIMESTAMP_DEFAULT_UNIT = Unit{
		Unit: UNIT_SECOND, 
		Scale: 1, 
		Offset: 0, 
	}
	MONITORING_INFO_ACTIVITY_TYPE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	MONITORING_INFO_CYCLES_TO_DISTANCE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 5000, 
		Offset: 0, 
	}
	MONITORING_INFO_CYCLES_TO_CALORIES_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 5000, 
		Offset: 0, 
	}
	MONITORING_INFO_RESTING_METABOLIC_RATE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	MONITORING_TIMESTAMP_DEFAULT_UNIT = Unit{
		Unit: UNIT_SECOND, 
		Scale: 1, 
		Offset: 0, 
	}
	MONITORING_DEVICE_INDEX_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	MONITORING_CALORIES_DEFAULT_UNIT = Unit{
		Unit: UNIT_KILOCALORIE, 
		Scale: 1, 
		Offset: 0, 
	}
	MONITORING_DISTANCE_DEFAULT_UNIT = Unit{
		Unit: UNIT_METER, 
		Scale: 100, 
		Offset: 0, 
	}
	MONITORING_CYCLES_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 2, 
		Offset: 0, 
	}
	MONITORING_ACTIVE_TIME_DEFAULT_UNIT = Unit{
		Unit: UNIT_SECOND, 
		Scale: 1000, 
		Offset: 0, 
	}
	MONITORING_ACTIVITY_TYPE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	MONITORING_ACTIVITY_SUBTYPE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	MONITORING_ACTIVITY_LEVEL_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	MONITORING_DISTANCE_16_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	MONITORING_CYCLES_16_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	MONITORING_ACTIVE_TIME_16_DEFAULT_UNIT = Unit{
		Unit: UNIT_SECOND, 
		Scale: 1, 
		Offset: 0, 
	}
	MONITORING_LOCAL_TIMESTAMP_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	MONITORING_TEMPERATURE_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 100, 
		Offset: 0, 
	}
	MONITORING_TEMPERATURE_MIN_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 100, 
		Offset: 0, 
	}
	MONITORING_TEMPERATURE_MAX_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 100, 
		Offset: 0, 
	}
	MONITORING_ACTIVITY_TIME_DEFAULT_UNIT = Unit{
		Unit: UNIT_MINUTE, 
		Scale: 1, 
		Offset: 0, 
	}
	MONITORING_ACTIVE_CALORIES_DEFAULT_UNIT = Unit{
		Unit: UNIT_KILOCALORIE, 
		Scale: 1, 
		Offset: 0, 
	}
	MONITORING_CURRENT_ACTIVITY_TYPE_INTENSITY_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 1, 
		Offset: 0, 
	}
	MONITORING_TIMESTAMP_MIN_8_DEFAULT_UNIT = Unit{
		Unit: UNIT_MINUTE, 
		Scale: 1, 
		Offset: 0, 
	}
	MONITORING_TIMESTAMP_16_DEFAULT_UNIT = Unit{
		Unit: UNIT_SECOND, 
		Scale: 1, 
		Offset: 0, 
	}
	MONITORING_HEART_RATE_DEFAULT_UNIT = Unit{
		Unit: UNIT_COUNTRATE, 
		Scale: 1, 
		Offset: 0, 
	}
	MONITORING_INTENSITY_DEFAULT_UNIT = Unit{
		Unit: UNIT_UNKNOWN, 
		Scale: 10, 
		Offset: 0, 
	}
	MONITORING_DURATION_MIN_DEFAULT_UNIT = Unit{
		Unit: UNIT_MINUTE, 
		Scale: 1, 
		Offset: 0, 
	}
	MONITORING_DURATION_DEFAULT_UNIT = Unit{
		Unit: UNIT_SECOND, 
		Scale: 1, 
		Offset: 0, 
	}
)

