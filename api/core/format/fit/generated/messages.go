// Code generated; DO NOT EDIT.

package generated

import (
	"encoding/binary"
    "fmt"
)

type DataMessageData interface {
	Read(uint, []byte, binary.ByteOrder) error
}

const (
	FileIdDataMessageID = 0
	FileCreatorDataMessageID = 49
	SoftwareDataMessageID = 35
	SlaveDeviceDataMessageID = 106
	CapabilitiesDataMessageID = 1
	FileCapabilitiesDataMessageID = 37
	MesgCapabilitiesDataMessageID = 38
	FieldCapabilitiesDataMessageID = 39
	DeviceSettingsDataMessageID = 2
	UserProfileDataMessageID = 3
	HrmProfileDataMessageID = 4
	SdmProfileDataMessageID = 5
	BikeProfileDataMessageID = 6
	ZonesTargetDataMessageID = 7
	SportDataMessageID = 12
	HrZoneDataMessageID = 8
	SpeedZoneDataMessageID = 53
	CadenceZoneDataMessageID = 131
	PowerZoneDataMessageID = 9
	MetZoneDataMessageID = 10
	GoalDataMessageID = 15
	ActivityDataMessageID = 34
	SessionDataMessageID = 18
	LapDataMessageID = 19
	LengthDataMessageID = 101
	RecordDataMessageID = 20
	EventDataMessageID = 21
	DeviceInfoDataMessageID = 23
	TrainingFileDataMessageID = 72
	HrvDataMessageID = 78
	CourseDataMessageID = 31
	CoursePointDataMessageID = 32
	SegmentIdDataMessageID = 148
	SegmentLeaderboardEntryDataMessageID = 149
	SegmentPointDataMessageID = 150
	SegmentLapDataMessageID = 142
	SegmentFileDataMessageID = 151
	WorkoutDataMessageID = 26
	WorkoutStepDataMessageID = 27
	ScheduleDataMessageID = 28
	TotalsDataMessageID = 33
	WeightScaleDataMessageID = 30
	BloodPressureDataMessageID = 51
	MonitoringInfoDataMessageID = 103
	MonitoringDataMessageID = 55
)

var MessageIDToString = map[uint]string {
0: "file_id",
49: "file_creator",
35: "software",
106: "slave_device",
1: "capabilities",
37: "file_capabilities",
38: "mesg_capabilities",
39: "field_capabilities",
2: "device_settings",
3: "user_profile",
4: "hrm_profile",
5: "sdm_profile",
6: "bike_profile",
7: "zones_target",
12: "sport",
8: "hr_zone",
53: "speed_zone",
131: "cadence_zone",
9: "power_zone",
10: "met_zone",
15: "goal",
34: "activity",
18: "session",
19: "lap",
101: "length",
20: "record",
21: "event",
23: "device_info",
72: "training_file",
78: "hrv",
31: "course",
32: "course_point",
148: "segment_id",
149: "segment_leaderboard_entry",
150: "segment_point",
142: "segment_lap",
151: "segment_file",
26: "workout",
27: "workout_step",
28: "schedule",
33: "totals",
30: "weight_scale",
51: "blood_pressure",
103: "monitoring_info",
55: "monitoring",
}

func CreateMessageData(messageNum uint) (DataMessageData, error) {
    switch messageNum {
    case 0:
		return &FileIdDataMessage{}, nil
    case 49:
		return &FileCreatorDataMessage{}, nil
    case 35:
		return &SoftwareDataMessage{}, nil
    case 106:
		return &SlaveDeviceDataMessage{}, nil
    case 1:
		return &CapabilitiesDataMessage{}, nil
    case 37:
		return &FileCapabilitiesDataMessage{}, nil
    case 38:
		return &MesgCapabilitiesDataMessage{}, nil
    case 39:
		return &FieldCapabilitiesDataMessage{}, nil
    case 2:
		return &DeviceSettingsDataMessage{}, nil
    case 3:
		return &UserProfileDataMessage{}, nil
    case 4:
		return &HrmProfileDataMessage{}, nil
    case 5:
		return &SdmProfileDataMessage{}, nil
    case 6:
		return &BikeProfileDataMessage{}, nil
    case 7:
		return &ZonesTargetDataMessage{}, nil
    case 12:
		return &SportDataMessage{}, nil
    case 8:
		return &HrZoneDataMessage{}, nil
    case 53:
		return &SpeedZoneDataMessage{}, nil
    case 131:
		return &CadenceZoneDataMessage{}, nil
    case 9:
		return &PowerZoneDataMessage{}, nil
    case 10:
		return &MetZoneDataMessage{}, nil
    case 15:
		return &GoalDataMessage{}, nil
    case 34:
		return &ActivityDataMessage{}, nil
    case 18:
		return &SessionDataMessage{}, nil
    case 19:
		return &LapDataMessage{}, nil
    case 101:
		return &LengthDataMessage{}, nil
    case 20:
		return &RecordDataMessage{}, nil
    case 21:
		return &EventDataMessage{}, nil
    case 23:
		return &DeviceInfoDataMessage{}, nil
    case 72:
		return &TrainingFileDataMessage{}, nil
    case 78:
		return &HrvDataMessage{}, nil
    case 31:
		return &CourseDataMessage{}, nil
    case 32:
		return &CoursePointDataMessage{}, nil
    case 148:
		return &SegmentIdDataMessage{}, nil
    case 149:
		return &SegmentLeaderboardEntryDataMessage{}, nil
    case 150:
		return &SegmentPointDataMessage{}, nil
    case 142:
		return &SegmentLapDataMessage{}, nil
    case 151:
		return &SegmentFileDataMessage{}, nil
    case 26:
		return &WorkoutDataMessage{}, nil
    case 27:
		return &WorkoutStepDataMessage{}, nil
    case 28:
		return &ScheduleDataMessage{}, nil
    case 33:
		return &TotalsDataMessage{}, nil
    case 30:
		return &WeightScaleDataMessage{}, nil
    case 51:
		return &BloodPressureDataMessage{}, nil
    case 103:
		return &MonitoringInfoDataMessage{}, nil
    case 55:
		return &MonitoringDataMessage{}, nil
    default:
        return nil, fmt.Errorf("unable to create message data with num %d", messageNum)
    }   
}


type FileIdDataMessage struct {
    Type *file `json:"type_,omitempty"`
    Manufacturer *manufacturer `json:"manufacturer,omitempty"`
    Product *uint16 `json:"product,omitempty"`
    SerialNumber *uint32 `json:"serial_number,omitempty"`
    TimeCreated *date_time `json:"time_created,omitempty"`
    Number *uint16 `json:"number,omitempty"`
	
}


type FileCreatorDataMessage struct {
    SoftwareVersion *uint16 `json:"software_version,omitempty"`
    HardwareVersion *uint8 `json:"hardware_version,omitempty"`
	
}


type SoftwareDataMessage struct {
    MessageIndex *message_index `json:"message_index,omitempty"`
    Version *uint16 `json:"version,omitempty"`
    PartNumber *string `json:"part_number,omitempty"`
	
}


type SlaveDeviceDataMessage struct {
    Manufacturer *manufacturer `json:"manufacturer,omitempty"`
    Product *uint16 `json:"product,omitempty"`
	
}


type CapabilitiesDataMessage struct {
    Languages *uint8 `json:"languages,omitempty"`
    Sports *sport_bits_0 `json:"sports,omitempty"`
    WorkoutsSupported *workout_capabilities `json:"workouts_supported,omitempty"`
    ConnectivitySupported *connectivity_capabilities `json:"connectivity_supported,omitempty"`
	
}


type FileCapabilitiesDataMessage struct {
    MessageIndex *message_index `json:"message_index,omitempty"`
    Type *file `json:"type_,omitempty"`
    Flags *file_flags `json:"flags,omitempty"`
    Directory *string `json:"directory,omitempty"`
    MaxCount *uint16 `json:"max_count,omitempty"`
    MaxSize *uint32 `json:"max_size,omitempty"`
	
}


type MesgCapabilitiesDataMessage struct {
    MessageIndex *message_index `json:"message_index,omitempty"`
    File *file `json:"file,omitempty"`
    MesgNum *mesg_num `json:"mesg_num,omitempty"`
    CountType *mesg_count `json:"count_type,omitempty"`
    Count *uint16 `json:"count,omitempty"`
	
}


type FieldCapabilitiesDataMessage struct {
    MessageIndex *message_index `json:"message_index,omitempty"`
    File *file `json:"file,omitempty"`
    MesgNum *mesg_num `json:"mesg_num,omitempty"`
    FieldNum *uint8 `json:"field_num,omitempty"`
    Count *uint16 `json:"count,omitempty"`
	
}


type DeviceSettingsDataMessage struct {
    ActiveTimeZone *uint8 `json:"active_time_zone,omitempty"`
    UtcOffset *uint32 `json:"utc_offset,omitempty"`
    TimeZoneOffset *int8 `json:"time_zone_offset,omitempty"`
	
}


type UserProfileDataMessage struct {
    MessageIndex *message_index `json:"message_index,omitempty"`
    FriendlyName *string `json:"friendly_name,omitempty"`
    Gender *gender `json:"gender,omitempty"`
    Age *uint8 `json:"age,omitempty"`
    Height *uint8 `json:"height,omitempty"`
    Weight *uint16 `json:"weight,omitempty"`
    Language *language `json:"language,omitempty"`
    ElevSetting *display_measure `json:"elev_setting,omitempty"`
    WeightSetting *display_measure `json:"weight_setting,omitempty"`
    RestingHeartRate *uint8 `json:"resting_heart_rate,omitempty"`
    DefaultMaxRunningHeartRate *uint8 `json:"default_max_running_heart_rate,omitempty"`
    DefaultMaxBikingHeartRate *uint8 `json:"default_max_biking_heart_rate,omitempty"`
    DefaultMaxHeartRate *uint8 `json:"default_max_heart_rate,omitempty"`
    HrSetting *display_heart `json:"hr_setting,omitempty"`
    SpeedSetting *display_measure `json:"speed_setting,omitempty"`
    DistSetting *display_measure `json:"dist_setting,omitempty"`
    PowerSetting *display_power `json:"power_setting,omitempty"`
    ActivityClass *activity_class `json:"activity_class,omitempty"`
    PositionSetting *display_position `json:"position_setting,omitempty"`
    TemperatureSetting *display_measure `json:"temperature_setting,omitempty"`
    LocalId *user_local_id `json:"local_id,omitempty"`
    GlobalId *byte `json:"global_id,omitempty"`
    HeightSetting *display_measure `json:"height_setting,omitempty"`
	
}


type HrmProfileDataMessage struct {
    MessageIndex *message_index `json:"message_index,omitempty"`
    Enabled *bool `json:"enabled,omitempty"`
    HrmAntId *uint16 `json:"hrm_ant_id,omitempty"`
    LogHrv *bool `json:"log_hrv,omitempty"`
    HrmAntIdTransType *uint8 `json:"hrm_ant_id_trans_type,omitempty"`
	
}


type SdmProfileDataMessage struct {
    MessageIndex *message_index `json:"message_index,omitempty"`
    Enabled *bool `json:"enabled,omitempty"`
    SdmAntId *uint16 `json:"sdm_ant_id,omitempty"`
    SdmCalFactor *uint16 `json:"sdm_cal_factor,omitempty"`
    Odometer *uint32 `json:"odometer,omitempty"`
    SpeedSource *bool `json:"speed_source,omitempty"`
    SdmAntIdTransType *uint8 `json:"sdm_ant_id_trans_type,omitempty"`
    OdometerRollover *uint8 `json:"odometer_rollover,omitempty"`
	
}


type BikeProfileDataMessage struct {
    MessageIndex *message_index `json:"message_index,omitempty"`
    Name *string `json:"name,omitempty"`
    Sport *sport `json:"sport,omitempty"`
    SubSport *sub_sport `json:"sub_sport,omitempty"`
    Odometer *uint32 `json:"odometer,omitempty"`
    BikeSpdAntId *uint16 `json:"bike_spd_ant_id,omitempty"`
    BikeCadAntId *uint16 `json:"bike_cad_ant_id,omitempty"`
    BikeSpdcadAntId *uint16 `json:"bike_spdcad_ant_id,omitempty"`
    BikePowerAntId *uint16 `json:"bike_power_ant_id,omitempty"`
    CustomWheelsize *uint16 `json:"custom_wheelsize,omitempty"`
    AutoWheelsize *uint16 `json:"auto_wheelsize,omitempty"`
    BikeWeight *uint16 `json:"bike_weight,omitempty"`
    PowerCalFactor *uint16 `json:"power_cal_factor,omitempty"`
    AutoWheelCal *bool `json:"auto_wheel_cal,omitempty"`
    AutoPowerZero *bool `json:"auto_power_zero,omitempty"`
    Id *uint8 `json:"id,omitempty"`
    SpdEnabled *bool `json:"spd_enabled,omitempty"`
    CadEnabled *bool `json:"cad_enabled,omitempty"`
    SpdcadEnabled *bool `json:"spdcad_enabled,omitempty"`
    PowerEnabled *bool `json:"power_enabled,omitempty"`
    CrankLength *uint8 `json:"crank_length,omitempty"`
    Enabled *bool `json:"enabled,omitempty"`
    BikeSpdAntIdTransType *uint8 `json:"bike_spd_ant_id_trans_type,omitempty"`
    BikeCadAntIdTransType *uint8 `json:"bike_cad_ant_id_trans_type,omitempty"`
    BikeSpdcadAntIdTransType *uint8 `json:"bike_spdcad_ant_id_trans_type,omitempty"`
    BikePowerAntIdTransType *uint8 `json:"bike_power_ant_id_trans_type,omitempty"`
    OdometerRollover *uint8 `json:"odometer_rollover,omitempty"`
    FrontGearNum *uint8 `json:"front_gear_num,omitempty"`
    FrontGear *uint8 `json:"front_gear,omitempty"`
    RearGearNum *uint8 `json:"rear_gear_num,omitempty"`
    RearGear *uint8 `json:"rear_gear,omitempty"`
    ShimanoDi2Enabled *bool `json:"shimano_di2_enabled,omitempty"`
	
}


type ZonesTargetDataMessage struct {
    MaxHeartRate *uint8 `json:"max_heart_rate,omitempty"`
    ThresholdHeartRate *uint8 `json:"threshold_heart_rate,omitempty"`
    FunctionalThresholdPower *uint16 `json:"functional_threshold_power,omitempty"`
    HrCalcType *hr_zone_calc `json:"hr_calc_type,omitempty"`
    PwrCalcType *pwr_zone_calc `json:"pwr_calc_type,omitempty"`
	
}


type SportDataMessage struct {
    Sport *sport `json:"sport,omitempty"`
    SubSport *sub_sport `json:"sub_sport,omitempty"`
    Name *string `json:"name,omitempty"`
	
}


type HrZoneDataMessage struct {
    MessageIndex *message_index `json:"message_index,omitempty"`
    HighBpm *uint8 `json:"high_bpm,omitempty"`
    Name *string `json:"name,omitempty"`
	
}


type SpeedZoneDataMessage struct {
    MessageIndex *message_index `json:"message_index,omitempty"`
    HighValue *uint16 `json:"high_value,omitempty"`
    Name *string `json:"name,omitempty"`
	
}


type CadenceZoneDataMessage struct {
    MessageIndex *message_index `json:"message_index,omitempty"`
    HighValue *uint8 `json:"high_value,omitempty"`
    Name *string `json:"name,omitempty"`
	
}


type PowerZoneDataMessage struct {
    MessageIndex *message_index `json:"message_index,omitempty"`
    HighValue *uint16 `json:"high_value,omitempty"`
    Name *string `json:"name,omitempty"`
	
}


type MetZoneDataMessage struct {
    MessageIndex *message_index `json:"message_index,omitempty"`
    HighBpm *uint8 `json:"high_bpm,omitempty"`
    Calories *uint16 `json:"calories,omitempty"`
    FatCalories *uint8 `json:"fat_calories,omitempty"`
	
}


type GoalDataMessage struct {
    MessageIndex *message_index `json:"message_index,omitempty"`
    Sport *sport `json:"sport,omitempty"`
    SubSport *sub_sport `json:"sub_sport,omitempty"`
    StartDate *date_time `json:"start_date,omitempty"`
    EndDate *date_time `json:"end_date,omitempty"`
    Type *goal `json:"type_,omitempty"`
    Value *uint32 `json:"value,omitempty"`
    Repeat *bool `json:"repeat,omitempty"`
    TargetValue *uint32 `json:"target_value,omitempty"`
    Recurrence *goal_recurrence `json:"recurrence,omitempty"`
    RecurrenceValue *uint16 `json:"recurrence_value,omitempty"`
    Enabled *bool `json:"enabled,omitempty"`
	
}


type ActivityDataMessage struct {
    Timestamp *date_time `json:"timestamp,omitempty"`
    TotalTimerTime *uint32 `json:"total_timer_time,omitempty"`
    NumSessions *uint16 `json:"num_sessions,omitempty"`
    Type *activity `json:"type_,omitempty"`
    Event *event `json:"event,omitempty"`
    EventType *event_type `json:"event_type,omitempty"`
    LocalTimestamp *local_date_time `json:"local_timestamp,omitempty"`
    EventGroup *uint8 `json:"event_group,omitempty"`
	
}


type SessionDataMessage struct {
    MessageIndex *message_index `json:"message_index,omitempty"`
    Timestamp *date_time `json:"timestamp,omitempty"`
    Event *event `json:"event,omitempty"`
    EventType *event_type `json:"event_type,omitempty"`
    StartTime *date_time `json:"start_time,omitempty"`
    StartPositionLat *int32 `json:"start_position_lat,omitempty"`
    StartPositionLong *int32 `json:"start_position_long,omitempty"`
    Sport *sport `json:"sport,omitempty"`
    SubSport *sub_sport `json:"sub_sport,omitempty"`
    TotalElapsedTime *uint32 `json:"total_elapsed_time,omitempty"`
    TotalTimerTime *uint32 `json:"total_timer_time,omitempty"`
    TotalDistance *uint32 `json:"total_distance,omitempty"`
    TotalCycles *uint32 `json:"total_cycles,omitempty"`
    TotalCalories *uint16 `json:"total_calories,omitempty"`
    TotalFatCalories *uint16 `json:"total_fat_calories,omitempty"`
    AvgSpeed *uint16 `json:"avg_speed,omitempty"`
    MaxSpeed *uint16 `json:"max_speed,omitempty"`
    AvgHeartRate *uint8 `json:"avg_heart_rate,omitempty"`
    MaxHeartRate *uint8 `json:"max_heart_rate,omitempty"`
    AvgCadence *uint8 `json:"avg_cadence,omitempty"`
    MaxCadence *uint8 `json:"max_cadence,omitempty"`
    AvgPower *uint16 `json:"avg_power,omitempty"`
    MaxPower *uint16 `json:"max_power,omitempty"`
    TotalAscent *uint16 `json:"total_ascent,omitempty"`
    TotalDescent *uint16 `json:"total_descent,omitempty"`
    TotalTrainingEffect *uint8 `json:"total_training_effect,omitempty"`
    FirstLapIndex *uint16 `json:"first_lap_index,omitempty"`
    NumLaps *uint16 `json:"num_laps,omitempty"`
    EventGroup *uint8 `json:"event_group,omitempty"`
    Trigger *session_trigger `json:"trigger,omitempty"`
    NecLat *int32 `json:"nec_lat,omitempty"`
    NecLong *int32 `json:"nec_long,omitempty"`
    SwcLat *int32 `json:"swc_lat,omitempty"`
    SwcLong *int32 `json:"swc_long,omitempty"`
    NormalizedPower *uint16 `json:"normalized_power,omitempty"`
    TrainingStressScore *uint16 `json:"training_stress_score,omitempty"`
    IntensityFactor *uint16 `json:"intensity_factor,omitempty"`
    LeftRightBalance *left_right_balance_100 `json:"left_right_balance,omitempty"`
    AvgStrokeCount *uint32 `json:"avg_stroke_count,omitempty"`
    AvgStrokeDistance *uint16 `json:"avg_stroke_distance,omitempty"`
    SwimStroke *swim_stroke `json:"swim_stroke,omitempty"`
    PoolLength *uint16 `json:"pool_length,omitempty"`
    ThresholdPower *uint16 `json:"threshold_power,omitempty"`
    PoolLengthUnit *display_measure `json:"pool_length_unit,omitempty"`
    NumActiveLengths *uint16 `json:"num_active_lengths,omitempty"`
    TotalWork *uint32 `json:"total_work,omitempty"`
    AvgAltitude *uint16 `json:"avg_altitude,omitempty"`
    MaxAltitude *uint16 `json:"max_altitude,omitempty"`
    GpsAccuracy *uint8 `json:"gps_accuracy,omitempty"`
    AvgGrade *int16 `json:"avg_grade,omitempty"`
    AvgPosGrade *int16 `json:"avg_pos_grade,omitempty"`
    AvgNegGrade *int16 `json:"avg_neg_grade,omitempty"`
    MaxPosGrade *int16 `json:"max_pos_grade,omitempty"`
    MaxNegGrade *int16 `json:"max_neg_grade,omitempty"`
    AvgTemperature *int8 `json:"avg_temperature,omitempty"`
    MaxTemperature *int8 `json:"max_temperature,omitempty"`
    TotalMovingTime *uint32 `json:"total_moving_time,omitempty"`
    AvgPosVerticalSpeed *int16 `json:"avg_pos_vertical_speed,omitempty"`
    AvgNegVerticalSpeed *int16 `json:"avg_neg_vertical_speed,omitempty"`
    MaxPosVerticalSpeed *int16 `json:"max_pos_vertical_speed,omitempty"`
    MaxNegVerticalSpeed *int16 `json:"max_neg_vertical_speed,omitempty"`
    MinHeartRate *uint8 `json:"min_heart_rate,omitempty"`
    TimeInHrZone *uint32 `json:"time_in_hr_zone,omitempty"`
    TimeInSpeedZone *uint32 `json:"time_in_speed_zone,omitempty"`
    TimeInCadenceZone *uint32 `json:"time_in_cadence_zone,omitempty"`
    TimeInPowerZone *uint32 `json:"time_in_power_zone,omitempty"`
    AvgLapTime *uint32 `json:"avg_lap_time,omitempty"`
    BestLapIndex *uint16 `json:"best_lap_index,omitempty"`
    MinAltitude *uint16 `json:"min_altitude,omitempty"`
    PlayerScore *uint16 `json:"player_score,omitempty"`
    OpponentScore *uint16 `json:"opponent_score,omitempty"`
    OpponentName *string `json:"opponent_name,omitempty"`
    StrokeCount *uint16 `json:"stroke_count,omitempty"`
    ZoneCount *uint16 `json:"zone_count,omitempty"`
    MaxBallSpeed *uint16 `json:"max_ball_speed,omitempty"`
    AvgBallSpeed *uint16 `json:"avg_ball_speed,omitempty"`
    AvgVerticalOscillation *uint16 `json:"avg_vertical_oscillation,omitempty"`
    AvgStanceTimePercent *uint16 `json:"avg_stance_time_percent,omitempty"`
    AvgStanceTime *uint16 `json:"avg_stance_time,omitempty"`
    AvgFractionalCadence *uint8 `json:"avg_fractional_cadence,omitempty"`
    MaxFractionalCadence *uint8 `json:"max_fractional_cadence,omitempty"`
    TotalFractionalCycles *uint8 `json:"total_fractional_cycles,omitempty"`
    AvgTotalHemoglobinConc *uint16 `json:"avg_total_hemoglobin_conc,omitempty"`
    MinTotalHemoglobinConc *uint16 `json:"min_total_hemoglobin_conc,omitempty"`
    MaxTotalHemoglobinConc *uint16 `json:"max_total_hemoglobin_conc,omitempty"`
    AvgSaturatedHemoglobinPercent *uint16 `json:"avg_saturated_hemoglobin_percent,omitempty"`
    MinSaturatedHemoglobinPercent *uint16 `json:"min_saturated_hemoglobin_percent,omitempty"`
    MaxSaturatedHemoglobinPercent *uint16 `json:"max_saturated_hemoglobin_percent,omitempty"`
    AvgLeftTorqueEffectiveness *uint8 `json:"avg_left_torque_effectiveness,omitempty"`
    AvgRightTorqueEffectiveness *uint8 `json:"avg_right_torque_effectiveness,omitempty"`
    AvgLeftPedalSmoothness *uint8 `json:"avg_left_pedal_smoothness,omitempty"`
    AvgRightPedalSmoothness *uint8 `json:"avg_right_pedal_smoothness,omitempty"`
    AvgCombinedPedalSmoothness *uint8 `json:"avg_combined_pedal_smoothness,omitempty"`
    SportIndex *uint8 `json:"sport_index,omitempty"`
    TimeStanding *uint32 `json:"time_standing,omitempty"`
    StandCount *uint16 `json:"stand_count,omitempty"`
    AvgLeftPco *int8 `json:"avg_left_pco,omitempty"`
    AvgRightPco *int8 `json:"avg_right_pco,omitempty"`
    AvgLeftPowerPhase *uint8 `json:"avg_left_power_phase,omitempty"`
    AvgLeftPowerPhasePeak *uint8 `json:"avg_left_power_phase_peak,omitempty"`
    AvgRightPowerPhase *uint8 `json:"avg_right_power_phase,omitempty"`
    AvgRightPowerPhasePeak *uint8 `json:"avg_right_power_phase_peak,omitempty"`
    AvgPowerPosition *uint16 `json:"avg_power_position,omitempty"`
    MaxPowerPosition *uint16 `json:"max_power_position,omitempty"`
    AvgCadencePosition *uint8 `json:"avg_cadence_position,omitempty"`
    MaxCadencePosition *uint8 `json:"max_cadence_position,omitempty"`
    EnhancedAvgSpeed *uint32 `json:"enhanced_avg_speed,omitempty"`
    EnhancedMaxSpeed *uint32 `json:"enhanced_max_speed,omitempty"`
    EnhancedAvgAltitude *uint32 `json:"enhanced_avg_altitude,omitempty"`
    EnhancedMinAltitude *uint32 `json:"enhanced_min_altitude,omitempty"`
    EnhancedMaxAltitude *uint32 `json:"enhanced_max_altitude,omitempty"`
    AvgLevMotorPower *uint16 `json:"avg_lev_motor_power,omitempty"`
    MaxLevMotorPower *uint16 `json:"max_lev_motor_power,omitempty"`
    LevBatteryConsumption *uint8 `json:"lev_battery_consumption,omitempty"`
	
}


type LapDataMessage struct {
    MessageIndex *message_index `json:"message_index,omitempty"`
    Timestamp *date_time `json:"timestamp,omitempty"`
    Event *event `json:"event,omitempty"`
    EventType *event_type `json:"event_type,omitempty"`
    StartTime *date_time `json:"start_time,omitempty"`
    StartPositionLat *int32 `json:"start_position_lat,omitempty"`
    StartPositionLong *int32 `json:"start_position_long,omitempty"`
    EndPositionLat *int32 `json:"end_position_lat,omitempty"`
    EndPositionLong *int32 `json:"end_position_long,omitempty"`
    TotalElapsedTime *uint32 `json:"total_elapsed_time,omitempty"`
    TotalTimerTime *uint32 `json:"total_timer_time,omitempty"`
    TotalDistance *uint32 `json:"total_distance,omitempty"`
    TotalCycles *uint32 `json:"total_cycles,omitempty"`
    TotalCalories *uint16 `json:"total_calories,omitempty"`
    TotalFatCalories *uint16 `json:"total_fat_calories,omitempty"`
    AvgSpeed *uint16 `json:"avg_speed,omitempty"`
    MaxSpeed *uint16 `json:"max_speed,omitempty"`
    AvgHeartRate *uint8 `json:"avg_heart_rate,omitempty"`
    MaxHeartRate *uint8 `json:"max_heart_rate,omitempty"`
    AvgCadence *uint8 `json:"avg_cadence,omitempty"`
    MaxCadence *uint8 `json:"max_cadence,omitempty"`
    AvgPower *uint16 `json:"avg_power,omitempty"`
    MaxPower *uint16 `json:"max_power,omitempty"`
    TotalAscent *uint16 `json:"total_ascent,omitempty"`
    TotalDescent *uint16 `json:"total_descent,omitempty"`
    Intensity *intensity `json:"intensity,omitempty"`
    LapTrigger *lap_trigger `json:"lap_trigger,omitempty"`
    Sport *sport `json:"sport,omitempty"`
    EventGroup *uint8 `json:"event_group,omitempty"`
    NumLengths *uint16 `json:"num_lengths,omitempty"`
    NormalizedPower *uint16 `json:"normalized_power,omitempty"`
    LeftRightBalance *left_right_balance_100 `json:"left_right_balance,omitempty"`
    FirstLengthIndex *uint16 `json:"first_length_index,omitempty"`
    AvgStrokeDistance *uint16 `json:"avg_stroke_distance,omitempty"`
    SwimStroke *swim_stroke `json:"swim_stroke,omitempty"`
    SubSport *sub_sport `json:"sub_sport,omitempty"`
    NumActiveLengths *uint16 `json:"num_active_lengths,omitempty"`
    TotalWork *uint32 `json:"total_work,omitempty"`
    AvgAltitude *uint16 `json:"avg_altitude,omitempty"`
    MaxAltitude *uint16 `json:"max_altitude,omitempty"`
    GpsAccuracy *uint8 `json:"gps_accuracy,omitempty"`
    AvgGrade *int16 `json:"avg_grade,omitempty"`
    AvgPosGrade *int16 `json:"avg_pos_grade,omitempty"`
    AvgNegGrade *int16 `json:"avg_neg_grade,omitempty"`
    MaxPosGrade *int16 `json:"max_pos_grade,omitempty"`
    MaxNegGrade *int16 `json:"max_neg_grade,omitempty"`
    AvgTemperature *int8 `json:"avg_temperature,omitempty"`
    MaxTemperature *int8 `json:"max_temperature,omitempty"`
    TotalMovingTime *uint32 `json:"total_moving_time,omitempty"`
    AvgPosVerticalSpeed *int16 `json:"avg_pos_vertical_speed,omitempty"`
    AvgNegVerticalSpeed *int16 `json:"avg_neg_vertical_speed,omitempty"`
    MaxPosVerticalSpeed *int16 `json:"max_pos_vertical_speed,omitempty"`
    MaxNegVerticalSpeed *int16 `json:"max_neg_vertical_speed,omitempty"`
    TimeInHrZone *uint32 `json:"time_in_hr_zone,omitempty"`
    TimeInSpeedZone *uint32 `json:"time_in_speed_zone,omitempty"`
    TimeInCadenceZone *uint32 `json:"time_in_cadence_zone,omitempty"`
    TimeInPowerZone *uint32 `json:"time_in_power_zone,omitempty"`
    RepetitionNum *uint16 `json:"repetition_num,omitempty"`
    MinAltitude *uint16 `json:"min_altitude,omitempty"`
    MinHeartRate *uint8 `json:"min_heart_rate,omitempty"`
    WktStepIndex *message_index `json:"wkt_step_index,omitempty"`
    OpponentScore *uint16 `json:"opponent_score,omitempty"`
    StrokeCount *uint16 `json:"stroke_count,omitempty"`
    ZoneCount *uint16 `json:"zone_count,omitempty"`
    AvgVerticalOscillation *uint16 `json:"avg_vertical_oscillation,omitempty"`
    AvgStanceTimePercent *uint16 `json:"avg_stance_time_percent,omitempty"`
    AvgStanceTime *uint16 `json:"avg_stance_time,omitempty"`
    AvgFractionalCadence *uint8 `json:"avg_fractional_cadence,omitempty"`
    MaxFractionalCadence *uint8 `json:"max_fractional_cadence,omitempty"`
    TotalFractionalCycles *uint8 `json:"total_fractional_cycles,omitempty"`
    PlayerScore *uint16 `json:"player_score,omitempty"`
    AvgTotalHemoglobinConc *uint16 `json:"avg_total_hemoglobin_conc,omitempty"`
    MinTotalHemoglobinConc *uint16 `json:"min_total_hemoglobin_conc,omitempty"`
    MaxTotalHemoglobinConc *uint16 `json:"max_total_hemoglobin_conc,omitempty"`
    AvgSaturatedHemoglobinPercent *uint16 `json:"avg_saturated_hemoglobin_percent,omitempty"`
    MinSaturatedHemoglobinPercent *uint16 `json:"min_saturated_hemoglobin_percent,omitempty"`
    MaxSaturatedHemoglobinPercent *uint16 `json:"max_saturated_hemoglobin_percent,omitempty"`
    AvgLeftTorqueEffectiveness *uint8 `json:"avg_left_torque_effectiveness,omitempty"`
    AvgRightTorqueEffectiveness *uint8 `json:"avg_right_torque_effectiveness,omitempty"`
    AvgLeftPedalSmoothness *uint8 `json:"avg_left_pedal_smoothness,omitempty"`
    AvgRightPedalSmoothness *uint8 `json:"avg_right_pedal_smoothness,omitempty"`
    AvgCombinedPedalSmoothness *uint8 `json:"avg_combined_pedal_smoothness,omitempty"`
    TimeStanding *uint32 `json:"time_standing,omitempty"`
    StandCount *uint16 `json:"stand_count,omitempty"`
    AvgLeftPco *int8 `json:"avg_left_pco,omitempty"`
    AvgRightPco *int8 `json:"avg_right_pco,omitempty"`
    AvgLeftPowerPhase *uint8 `json:"avg_left_power_phase,omitempty"`
    AvgLeftPowerPhasePeak *uint8 `json:"avg_left_power_phase_peak,omitempty"`
    AvgRightPowerPhase *uint8 `json:"avg_right_power_phase,omitempty"`
    AvgRightPowerPhasePeak *uint8 `json:"avg_right_power_phase_peak,omitempty"`
    AvgPowerPosition *uint16 `json:"avg_power_position,omitempty"`
    MaxPowerPosition *uint16 `json:"max_power_position,omitempty"`
    AvgCadencePosition *uint8 `json:"avg_cadence_position,omitempty"`
    MaxCadencePosition *uint8 `json:"max_cadence_position,omitempty"`
    EnhancedAvgSpeed *uint32 `json:"enhanced_avg_speed,omitempty"`
    EnhancedMaxSpeed *uint32 `json:"enhanced_max_speed,omitempty"`
    EnhancedAvgAltitude *uint32 `json:"enhanced_avg_altitude,omitempty"`
    EnhancedMinAltitude *uint32 `json:"enhanced_min_altitude,omitempty"`
    EnhancedMaxAltitude *uint32 `json:"enhanced_max_altitude,omitempty"`
    AvgLevMotorPower *uint16 `json:"avg_lev_motor_power,omitempty"`
    MaxLevMotorPower *uint16 `json:"max_lev_motor_power,omitempty"`
    LevBatteryConsumption *uint8 `json:"lev_battery_consumption,omitempty"`
	
}


type LengthDataMessage struct {
    MessageIndex *message_index `json:"message_index,omitempty"`
    Timestamp *date_time `json:"timestamp,omitempty"`
    Event *event `json:"event,omitempty"`
    EventType *event_type `json:"event_type,omitempty"`
    StartTime *date_time `json:"start_time,omitempty"`
    TotalElapsedTime *uint32 `json:"total_elapsed_time,omitempty"`
    TotalTimerTime *uint32 `json:"total_timer_time,omitempty"`
    TotalStrokes *uint16 `json:"total_strokes,omitempty"`
    AvgSpeed *uint16 `json:"avg_speed,omitempty"`
    SwimStroke *swim_stroke `json:"swim_stroke,omitempty"`
    AvgSwimmingCadence *uint8 `json:"avg_swimming_cadence,omitempty"`
    EventGroup *uint8 `json:"event_group,omitempty"`
    TotalCalories *uint16 `json:"total_calories,omitempty"`
    LengthType *length_type `json:"length_type,omitempty"`
    PlayerScore *uint16 `json:"player_score,omitempty"`
    OpponentScore *uint16 `json:"opponent_score,omitempty"`
    StrokeCount *uint16 `json:"stroke_count,omitempty"`
    ZoneCount *uint16 `json:"zone_count,omitempty"`
	
}


type RecordDataMessage struct {
    Timestamp *date_time `json:"timestamp,omitempty"`
    PositionLat *int32 `json:"position_lat,omitempty"`
    PositionLong *int32 `json:"position_long,omitempty"`
    Altitude *uint16 `json:"altitude,omitempty"`
    HeartRate *uint8 `json:"heart_rate,omitempty"`
    Cadence *uint8 `json:"cadence,omitempty"`
    Distance *uint32 `json:"distance,omitempty"`
    Speed *uint16 `json:"speed,omitempty"`
    Power *uint16 `json:"power,omitempty"`
    CompressedSpeedDistance *byte `json:"compressed_speed_distance,omitempty"`
    Grade *int16 `json:"grade,omitempty"`
    Resistance *uint8 `json:"resistance,omitempty"`
    TimeFromCourse *int32 `json:"time_from_course,omitempty"`
    CycleLength *uint8 `json:"cycle_length,omitempty"`
    Temperature *int8 `json:"temperature,omitempty"`
    Speed1s *uint8 `json:"speed_1s,omitempty"`
    Cycles *uint8 `json:"cycles,omitempty"`
    TotalCycles *uint32 `json:"total_cycles,omitempty"`
    CompressedAccumulatedPower *uint16 `json:"compressed_accumulated_power,omitempty"`
    AccumulatedPower *uint32 `json:"accumulated_power,omitempty"`
    LeftRightBalance *left_right_balance `json:"left_right_balance,omitempty"`
    GpsAccuracy *uint8 `json:"gps_accuracy,omitempty"`
    VerticalSpeed *int16 `json:"vertical_speed,omitempty"`
    Calories *uint16 `json:"calories,omitempty"`
    VerticalOscillation *uint16 `json:"vertical_oscillation,omitempty"`
    StanceTimePercent *uint16 `json:"stance_time_percent,omitempty"`
    StanceTime *uint16 `json:"stance_time,omitempty"`
    ActivityType *activity_type `json:"activity_type,omitempty"`
    LeftTorqueEffectiveness *uint8 `json:"left_torque_effectiveness,omitempty"`
    RightTorqueEffectiveness *uint8 `json:"right_torque_effectiveness,omitempty"`
    LeftPedalSmoothness *uint8 `json:"left_pedal_smoothness,omitempty"`
    RightPedalSmoothness *uint8 `json:"right_pedal_smoothness,omitempty"`
    CombinedPedalSmoothness *uint8 `json:"combined_pedal_smoothness,omitempty"`
    Time128 *uint8 `json:"time128,omitempty"`
    StrokeType *stroke_type `json:"stroke_type,omitempty"`
    Zone *uint8 `json:"zone,omitempty"`
    BallSpeed *uint16 `json:"ball_speed,omitempty"`
    Cadence256 *uint16 `json:"cadence256,omitempty"`
    FractionalCadence *uint8 `json:"fractional_cadence,omitempty"`
    TotalHemoglobinConc *uint16 `json:"total_hemoglobin_conc,omitempty"`
    TotalHemoglobinConcMin *uint16 `json:"total_hemoglobin_conc_min,omitempty"`
    TotalHemoglobinConcMax *uint16 `json:"total_hemoglobin_conc_max,omitempty"`
    SaturatedHemoglobinPercent *uint16 `json:"saturated_hemoglobin_percent,omitempty"`
    SaturatedHemoglobinPercentMin *uint16 `json:"saturated_hemoglobin_percent_min,omitempty"`
    SaturatedHemoglobinPercentMax *uint16 `json:"saturated_hemoglobin_percent_max,omitempty"`
    DeviceIndex *device_index `json:"device_index,omitempty"`
    LeftPco *int8 `json:"left_pco,omitempty"`
    RightPco *int8 `json:"right_pco,omitempty"`
    LeftPowerPhase *uint8 `json:"left_power_phase,omitempty"`
    LeftPowerPhasePeak *uint8 `json:"left_power_phase_peak,omitempty"`
    RightPowerPhase *uint8 `json:"right_power_phase,omitempty"`
    RightPowerPhasePeak *uint8 `json:"right_power_phase_peak,omitempty"`
    EnhancedSpeed *uint32 `json:"enhanced_speed,omitempty"`
    EnhancedAltitude *uint32 `json:"enhanced_altitude,omitempty"`
    BatterySoc *uint8 `json:"battery_soc,omitempty"`
    MotorPower *uint16 `json:"motor_power,omitempty"`
	
}


type EventDataMessage struct {
    Timestamp *date_time `json:"timestamp,omitempty"`
    Event *event `json:"event,omitempty"`
    EventType *event_type `json:"event_type,omitempty"`
    Data16 *uint16 `json:"data16,omitempty"`
    Data *uint32 `json:"data,omitempty"`
    EventGroup *uint8 `json:"event_group,omitempty"`
    Score *uint16 `json:"score,omitempty"`
    OpponentScore *uint16 `json:"opponent_score,omitempty"`
    FrontGearNum *uint8 `json:"front_gear_num,omitempty"`
    FrontGear *uint8 `json:"front_gear,omitempty"`
    RearGearNum *uint8 `json:"rear_gear_num,omitempty"`
    RearGear *uint8 `json:"rear_gear,omitempty"`
    DeviceIndex *device_index `json:"device_index,omitempty"`
	
}


type DeviceInfoDataMessage struct {
    Timestamp *date_time `json:"timestamp,omitempty"`
    DeviceIndex *device_index `json:"device_index,omitempty"`
    DeviceType *uint8 `json:"device_type,omitempty"`
    Manufacturer *manufacturer `json:"manufacturer,omitempty"`
    SerialNumber *uint32 `json:"serial_number,omitempty"`
    Product *uint16 `json:"product,omitempty"`
    SoftwareVersion *uint16 `json:"software_version,omitempty"`
    HardwareVersion *uint8 `json:"hardware_version,omitempty"`
    CumOperatingTime *uint32 `json:"cum_operating_time,omitempty"`
    BatteryVoltage *uint16 `json:"battery_voltage,omitempty"`
    BatteryStatus *battery_status `json:"battery_status,omitempty"`
    SensorPosition *body_location `json:"sensor_position,omitempty"`
    Descriptor *string `json:"descriptor,omitempty"`
    AntTransmissionType *uint8 `json:"ant_transmission_type,omitempty"`
    AntDeviceNumber *uint16 `json:"ant_device_number,omitempty"`
    AntNetwork *ant_network `json:"ant_network,omitempty"`
    SourceType *source_type `json:"source_type,omitempty"`
	
}


type TrainingFileDataMessage struct {
    Timestamp *date_time `json:"timestamp,omitempty"`
    Type *file `json:"type_,omitempty"`
    Manufacturer *manufacturer `json:"manufacturer,omitempty"`
    Product *uint16 `json:"product,omitempty"`
    SerialNumber *uint32 `json:"serial_number,omitempty"`
    TimeCreated *date_time `json:"time_created,omitempty"`
	
}


type HrvDataMessage struct {
    Time *uint16 `json:"time,omitempty"`
	
}


type CourseDataMessage struct {
    Sport *sport `json:"sport,omitempty"`
    Name *string `json:"name,omitempty"`
    Capabilities *course_capabilities `json:"capabilities,omitempty"`
	
}


type CoursePointDataMessage struct {
    MessageIndex *message_index `json:"message_index,omitempty"`
    Timestamp *date_time `json:"timestamp,omitempty"`
    PositionLat *int32 `json:"position_lat,omitempty"`
    PositionLong *int32 `json:"position_long,omitempty"`
    Distance *uint32 `json:"distance,omitempty"`
    Type *course_point `json:"type_,omitempty"`
    Name *string `json:"name,omitempty"`
    Favorite *bool `json:"favorite,omitempty"`
	
}


type SegmentIdDataMessage struct {
    Name *string `json:"name,omitempty"`
    Uuid *string `json:"uuid,omitempty"`
    Sport *sport `json:"sport,omitempty"`
    Enabled *bool `json:"enabled,omitempty"`
    UserProfilePrimaryKey *uint32 `json:"user_profile_primary_key,omitempty"`
    DeviceId *uint32 `json:"device_id,omitempty"`
    DefaultRaceLeader *uint8 `json:"default_race_leader,omitempty"`
    DeleteStatus *segment_delete_status `json:"delete_status,omitempty"`
    SelectionType *segment_selection_type `json:"selection_type,omitempty"`
	
}


type SegmentLeaderboardEntryDataMessage struct {
    MessageIndex *message_index `json:"message_index,omitempty"`
    Name *string `json:"name,omitempty"`
    Type *segment_leaderboard_type `json:"type_,omitempty"`
    GroupPrimaryKey *uint32 `json:"group_primary_key,omitempty"`
    ActivityId *uint32 `json:"activity_id,omitempty"`
    SegmentTime *uint32 `json:"segment_time,omitempty"`
    ActivityIdString *string `json:"activity_id_string,omitempty"`
	
}


type SegmentPointDataMessage struct {
    MessageIndex *message_index `json:"message_index,omitempty"`
    PositionLat *int32 `json:"position_lat,omitempty"`
    PositionLong *int32 `json:"position_long,omitempty"`
    Distance *uint32 `json:"distance,omitempty"`
    Altitude *uint16 `json:"altitude,omitempty"`
    LeaderTime *uint32 `json:"leader_time,omitempty"`
	
}


type SegmentLapDataMessage struct {
    MessageIndex *message_index `json:"message_index,omitempty"`
    Timestamp *date_time `json:"timestamp,omitempty"`
    Event *event `json:"event,omitempty"`
    EventType *event_type `json:"event_type,omitempty"`
    StartTime *date_time `json:"start_time,omitempty"`
    StartPositionLat *int32 `json:"start_position_lat,omitempty"`
    StartPositionLong *int32 `json:"start_position_long,omitempty"`
    EndPositionLat *int32 `json:"end_position_lat,omitempty"`
    EndPositionLong *int32 `json:"end_position_long,omitempty"`
    TotalElapsedTime *uint32 `json:"total_elapsed_time,omitempty"`
    TotalTimerTime *uint32 `json:"total_timer_time,omitempty"`
    TotalDistance *uint32 `json:"total_distance,omitempty"`
    TotalCycles *uint32 `json:"total_cycles,omitempty"`
    TotalCalories *uint16 `json:"total_calories,omitempty"`
    TotalFatCalories *uint16 `json:"total_fat_calories,omitempty"`
    AvgSpeed *uint16 `json:"avg_speed,omitempty"`
    MaxSpeed *uint16 `json:"max_speed,omitempty"`
    AvgHeartRate *uint8 `json:"avg_heart_rate,omitempty"`
    MaxHeartRate *uint8 `json:"max_heart_rate,omitempty"`
    AvgCadence *uint8 `json:"avg_cadence,omitempty"`
    MaxCadence *uint8 `json:"max_cadence,omitempty"`
    AvgPower *uint16 `json:"avg_power,omitempty"`
    MaxPower *uint16 `json:"max_power,omitempty"`
    TotalAscent *uint16 `json:"total_ascent,omitempty"`
    TotalDescent *uint16 `json:"total_descent,omitempty"`
    Sport *sport `json:"sport,omitempty"`
    EventGroup *uint8 `json:"event_group,omitempty"`
    NecLat *int32 `json:"nec_lat,omitempty"`
    NecLong *int32 `json:"nec_long,omitempty"`
    SwcLat *int32 `json:"swc_lat,omitempty"`
    SwcLong *int32 `json:"swc_long,omitempty"`
    Name *string `json:"name,omitempty"`
    NormalizedPower *uint16 `json:"normalized_power,omitempty"`
    LeftRightBalance *left_right_balance_100 `json:"left_right_balance,omitempty"`
    SubSport *sub_sport `json:"sub_sport,omitempty"`
    TotalWork *uint32 `json:"total_work,omitempty"`
    AvgAltitude *uint16 `json:"avg_altitude,omitempty"`
    MaxAltitude *uint16 `json:"max_altitude,omitempty"`
    GpsAccuracy *uint8 `json:"gps_accuracy,omitempty"`
    AvgGrade *int16 `json:"avg_grade,omitempty"`
    AvgPosGrade *int16 `json:"avg_pos_grade,omitempty"`
    AvgNegGrade *int16 `json:"avg_neg_grade,omitempty"`
    MaxPosGrade *int16 `json:"max_pos_grade,omitempty"`
    MaxNegGrade *int16 `json:"max_neg_grade,omitempty"`
    AvgTemperature *int8 `json:"avg_temperature,omitempty"`
    MaxTemperature *int8 `json:"max_temperature,omitempty"`
    TotalMovingTime *uint32 `json:"total_moving_time,omitempty"`
    AvgPosVerticalSpeed *int16 `json:"avg_pos_vertical_speed,omitempty"`
    AvgNegVerticalSpeed *int16 `json:"avg_neg_vertical_speed,omitempty"`
    MaxPosVerticalSpeed *int16 `json:"max_pos_vertical_speed,omitempty"`
    MaxNegVerticalSpeed *int16 `json:"max_neg_vertical_speed,omitempty"`
    TimeInHrZone *uint32 `json:"time_in_hr_zone,omitempty"`
    TimeInSpeedZone *uint32 `json:"time_in_speed_zone,omitempty"`
    TimeInCadenceZone *uint32 `json:"time_in_cadence_zone,omitempty"`
    TimeInPowerZone *uint32 `json:"time_in_power_zone,omitempty"`
    RepetitionNum *uint16 `json:"repetition_num,omitempty"`
    MinAltitude *uint16 `json:"min_altitude,omitempty"`
    MinHeartRate *uint8 `json:"min_heart_rate,omitempty"`
    ActiveTime *uint32 `json:"active_time,omitempty"`
    WktStepIndex *message_index `json:"wkt_step_index,omitempty"`
    SportEvent *sport_event `json:"sport_event,omitempty"`
    AvgLeftTorqueEffectiveness *uint8 `json:"avg_left_torque_effectiveness,omitempty"`
    AvgRightTorqueEffectiveness *uint8 `json:"avg_right_torque_effectiveness,omitempty"`
    AvgLeftPedalSmoothness *uint8 `json:"avg_left_pedal_smoothness,omitempty"`
    AvgRightPedalSmoothness *uint8 `json:"avg_right_pedal_smoothness,omitempty"`
    AvgCombinedPedalSmoothness *uint8 `json:"avg_combined_pedal_smoothness,omitempty"`
    Status *segment_lap_status `json:"status,omitempty"`
    Uuid *string `json:"uuid,omitempty"`
    AvgFractionalCadence *uint8 `json:"avg_fractional_cadence,omitempty"`
    MaxFractionalCadence *uint8 `json:"max_fractional_cadence,omitempty"`
    TotalFractionalCycles *uint8 `json:"total_fractional_cycles,omitempty"`
    FrontGearShiftCount *uint16 `json:"front_gear_shift_count,omitempty"`
    RearGearShiftCount *uint16 `json:"rear_gear_shift_count,omitempty"`
    TimeStanding *uint32 `json:"time_standing,omitempty"`
    StandCount *uint16 `json:"stand_count,omitempty"`
    AvgLeftPco *int8 `json:"avg_left_pco,omitempty"`
    AvgRightPco *int8 `json:"avg_right_pco,omitempty"`
    AvgLeftPowerPhase *uint8 `json:"avg_left_power_phase,omitempty"`
    AvgLeftPowerPhasePeak *uint8 `json:"avg_left_power_phase_peak,omitempty"`
    AvgRightPowerPhase *uint8 `json:"avg_right_power_phase,omitempty"`
    AvgRightPowerPhasePeak *uint8 `json:"avg_right_power_phase_peak,omitempty"`
    AvgPowerPosition *uint16 `json:"avg_power_position,omitempty"`
    MaxPowerPosition *uint16 `json:"max_power_position,omitempty"`
    AvgCadencePosition *uint8 `json:"avg_cadence_position,omitempty"`
    MaxCadencePosition *uint8 `json:"max_cadence_position,omitempty"`
	
}


type SegmentFileDataMessage struct {
    MessageIndex *message_index `json:"message_index,omitempty"`
    FileUuid *string `json:"file_uuid,omitempty"`
    Enabled *bool `json:"enabled,omitempty"`
    UserProfilePrimaryKey *uint32 `json:"user_profile_primary_key,omitempty"`
    LeaderType *segment_leaderboard_type `json:"leader_type,omitempty"`
    LeaderGroupPrimaryKey *uint32 `json:"leader_group_primary_key,omitempty"`
    LeaderActivityId *uint32 `json:"leader_activity_id,omitempty"`
    LeaderActivityIdString *string `json:"leader_activity_id_string,omitempty"`
	
}


type WorkoutDataMessage struct {
    Sport *sport `json:"sport,omitempty"`
    Capabilities *workout_capabilities `json:"capabilities,omitempty"`
    NumValidSteps *uint16 `json:"num_valid_steps,omitempty"`
    WktName *string `json:"wkt_name,omitempty"`
	
}


type WorkoutStepDataMessage struct {
    MessageIndex *message_index `json:"message_index,omitempty"`
    WktStepName *string `json:"wkt_step_name,omitempty"`
    DurationType *wkt_step_duration `json:"duration_type,omitempty"`
    DurationValue *uint32 `json:"duration_value,omitempty"`
    TargetType *wkt_step_target `json:"target_type,omitempty"`
    TargetValue *uint32 `json:"target_value,omitempty"`
    CustomTargetValueLow *uint32 `json:"custom_target_value_low,omitempty"`
    CustomTargetValueHigh *uint32 `json:"custom_target_value_high,omitempty"`
    Intensity *intensity `json:"intensity,omitempty"`
	
}


type ScheduleDataMessage struct {
    Manufacturer *manufacturer `json:"manufacturer,omitempty"`
    Product *uint16 `json:"product,omitempty"`
    SerialNumber *uint32 `json:"serial_number,omitempty"`
    TimeCreated *date_time `json:"time_created,omitempty"`
    Completed *bool `json:"completed,omitempty"`
    Type *schedule `json:"type_,omitempty"`
    ScheduledTime *local_date_time `json:"scheduled_time,omitempty"`
	
}


type TotalsDataMessage struct {
    MessageIndex *message_index `json:"message_index,omitempty"`
    Timestamp *date_time `json:"timestamp,omitempty"`
    TimerTime *uint32 `json:"timer_time,omitempty"`
    Distance *uint32 `json:"distance,omitempty"`
    Calories *uint32 `json:"calories,omitempty"`
    Sport *sport `json:"sport,omitempty"`
    ElapsedTime *uint32 `json:"elapsed_time,omitempty"`
    Sessions *uint16 `json:"sessions,omitempty"`
    ActiveTime *uint32 `json:"active_time,omitempty"`
    SportIndex *uint8 `json:"sport_index,omitempty"`
	
}


type WeightScaleDataMessage struct {
    Timestamp *date_time `json:"timestamp,omitempty"`
    Weight *weight `json:"weight,omitempty"`
    PercentFat *uint16 `json:"percent_fat,omitempty"`
    PercentHydration *uint16 `json:"percent_hydration,omitempty"`
    VisceralFatMass *uint16 `json:"visceral_fat_mass,omitempty"`
    BoneMass *uint16 `json:"bone_mass,omitempty"`
    MuscleMass *uint16 `json:"muscle_mass,omitempty"`
    BasalMet *uint16 `json:"basal_met,omitempty"`
    PhysiqueRating *uint8 `json:"physique_rating,omitempty"`
    ActiveMet *uint16 `json:"active_met,omitempty"`
    MetabolicAge *uint8 `json:"metabolic_age,omitempty"`
    VisceralFatRating *uint8 `json:"visceral_fat_rating,omitempty"`
    UserProfileIndex *message_index `json:"user_profile_index,omitempty"`
	
}


type BloodPressureDataMessage struct {
    Timestamp *date_time `json:"timestamp,omitempty"`
    SystolicPressure *uint16 `json:"systolic_pressure,omitempty"`
    DiastolicPressure *uint16 `json:"diastolic_pressure,omitempty"`
    MeanArterialPressure *uint16 `json:"mean_arterial_pressure,omitempty"`
    Map3SampleMean *uint16 `json:"map_3_sample_mean,omitempty"`
    MapMorningValues *uint16 `json:"map_morning_values,omitempty"`
    MapEveningValues *uint16 `json:"map_evening_values,omitempty"`
    HeartRate *uint8 `json:"heart_rate,omitempty"`
    HeartRateType *hr_type `json:"heart_rate_type,omitempty"`
    Status *bp_status `json:"status,omitempty"`
    UserProfileIndex *message_index `json:"user_profile_index,omitempty"`
	
}


type MonitoringInfoDataMessage struct {
    Timestamp *date_time `json:"timestamp,omitempty"`
    LocalTimestamp *local_date_time `json:"local_timestamp,omitempty"`
    ActivityType *activity_type `json:"activity_type,omitempty"`
    CyclesToDistance *uint16 `json:"cycles_to_distance,omitempty"`
    CyclesToCalories *uint16 `json:"cycles_to_calories,omitempty"`
    RestingMetabolicRate *uint16 `json:"resting_metabolic_rate,omitempty"`
	
}


type MonitoringDataMessage struct {
    Timestamp *date_time `json:"timestamp,omitempty"`
    DeviceIndex *device_index `json:"device_index,omitempty"`
    Calories *uint16 `json:"calories,omitempty"`
    Distance *uint32 `json:"distance,omitempty"`
    Cycles *uint32 `json:"cycles,omitempty"`
    ActiveTime *uint32 `json:"active_time,omitempty"`
    ActivityType *activity_type `json:"activity_type,omitempty"`
    ActivitySubtype *activity_subtype `json:"activity_subtype,omitempty"`
    ActivityLevel *activity_level `json:"activity_level,omitempty"`
    Distance16 *uint16 `json:"distance_16,omitempty"`
    Cycles16 *uint16 `json:"cycles_16,omitempty"`
    ActiveTime16 *uint16 `json:"active_time_16,omitempty"`
    LocalTimestamp *local_date_time `json:"local_timestamp,omitempty"`
    Temperature *int16 `json:"temperature,omitempty"`
    TemperatureMin *int16 `json:"temperature_min,omitempty"`
    TemperatureMax *int16 `json:"temperature_max,omitempty"`
    ActivityTime *uint16 `json:"activity_time,omitempty"`
    ActiveCalories *uint16 `json:"active_calories,omitempty"`
    CurrentActivityTypeIntensity *byte `json:"current_activity_type_intensity,omitempty"`
    TimestampMin8 *uint8 `json:"timestamp_min_8,omitempty"`
    Timestamp16 *uint16 `json:"timestamp_16,omitempty"`
    HeartRate *uint8 `json:"heart_rate,omitempty"`
    Intensity *uint8 `json:"intensity,omitempty"`
    DurationMin *uint16 `json:"duration_min,omitempty"`
    Duration *uint32 `json:"duration,omitempty"`
	
}




func newFileIdDataMessage() *FileIdDataMessage {
	return &FileIdDataMessage{}
}

func newFileCreatorDataMessage() *FileCreatorDataMessage {
	return &FileCreatorDataMessage{}
}

func newSoftwareDataMessage() *SoftwareDataMessage {
	return &SoftwareDataMessage{}
}

func newSlaveDeviceDataMessage() *SlaveDeviceDataMessage {
	return &SlaveDeviceDataMessage{}
}

func newCapabilitiesDataMessage() *CapabilitiesDataMessage {
	return &CapabilitiesDataMessage{}
}

func newFileCapabilitiesDataMessage() *FileCapabilitiesDataMessage {
	return &FileCapabilitiesDataMessage{}
}

func newMesgCapabilitiesDataMessage() *MesgCapabilitiesDataMessage {
	return &MesgCapabilitiesDataMessage{}
}

func newFieldCapabilitiesDataMessage() *FieldCapabilitiesDataMessage {
	return &FieldCapabilitiesDataMessage{}
}

func newDeviceSettingsDataMessage() *DeviceSettingsDataMessage {
	return &DeviceSettingsDataMessage{}
}

func newUserProfileDataMessage() *UserProfileDataMessage {
	return &UserProfileDataMessage{}
}

func newHrmProfileDataMessage() *HrmProfileDataMessage {
	return &HrmProfileDataMessage{}
}

func newSdmProfileDataMessage() *SdmProfileDataMessage {
	return &SdmProfileDataMessage{}
}

func newBikeProfileDataMessage() *BikeProfileDataMessage {
	return &BikeProfileDataMessage{}
}

func newZonesTargetDataMessage() *ZonesTargetDataMessage {
	return &ZonesTargetDataMessage{}
}

func newSportDataMessage() *SportDataMessage {
	return &SportDataMessage{}
}

func newHrZoneDataMessage() *HrZoneDataMessage {
	return &HrZoneDataMessage{}
}

func newSpeedZoneDataMessage() *SpeedZoneDataMessage {
	return &SpeedZoneDataMessage{}
}

func newCadenceZoneDataMessage() *CadenceZoneDataMessage {
	return &CadenceZoneDataMessage{}
}

func newPowerZoneDataMessage() *PowerZoneDataMessage {
	return &PowerZoneDataMessage{}
}

func newMetZoneDataMessage() *MetZoneDataMessage {
	return &MetZoneDataMessage{}
}

func newGoalDataMessage() *GoalDataMessage {
	return &GoalDataMessage{}
}

func newActivityDataMessage() *ActivityDataMessage {
	return &ActivityDataMessage{}
}

func newSessionDataMessage() *SessionDataMessage {
	return &SessionDataMessage{}
}

func newLapDataMessage() *LapDataMessage {
	return &LapDataMessage{}
}

func newLengthDataMessage() *LengthDataMessage {
	return &LengthDataMessage{}
}

func newRecordDataMessage() *RecordDataMessage {
	return &RecordDataMessage{}
}

func newEventDataMessage() *EventDataMessage {
	return &EventDataMessage{}
}

func newDeviceInfoDataMessage() *DeviceInfoDataMessage {
	return &DeviceInfoDataMessage{}
}

func newTrainingFileDataMessage() *TrainingFileDataMessage {
	return &TrainingFileDataMessage{}
}

func newHrvDataMessage() *HrvDataMessage {
	return &HrvDataMessage{}
}

func newCourseDataMessage() *CourseDataMessage {
	return &CourseDataMessage{}
}

func newCoursePointDataMessage() *CoursePointDataMessage {
	return &CoursePointDataMessage{}
}

func newSegmentIdDataMessage() *SegmentIdDataMessage {
	return &SegmentIdDataMessage{}
}

func newSegmentLeaderboardEntryDataMessage() *SegmentLeaderboardEntryDataMessage {
	return &SegmentLeaderboardEntryDataMessage{}
}

func newSegmentPointDataMessage() *SegmentPointDataMessage {
	return &SegmentPointDataMessage{}
}

func newSegmentLapDataMessage() *SegmentLapDataMessage {
	return &SegmentLapDataMessage{}
}

func newSegmentFileDataMessage() *SegmentFileDataMessage {
	return &SegmentFileDataMessage{}
}

func newWorkoutDataMessage() *WorkoutDataMessage {
	return &WorkoutDataMessage{}
}

func newWorkoutStepDataMessage() *WorkoutStepDataMessage {
	return &WorkoutStepDataMessage{}
}

func newScheduleDataMessage() *ScheduleDataMessage {
	return &ScheduleDataMessage{}
}

func newTotalsDataMessage() *TotalsDataMessage {
	return &TotalsDataMessage{}
}

func newWeightScaleDataMessage() *WeightScaleDataMessage {
	return &WeightScaleDataMessage{}
}

func newBloodPressureDataMessage() *BloodPressureDataMessage {
	return &BloodPressureDataMessage{}
}

func newMonitoringInfoDataMessage() *MonitoringInfoDataMessage {
	return &MonitoringInfoDataMessage{}
}

func newMonitoringDataMessage() *MonitoringDataMessage {
	return &MonitoringDataMessage{}
}



func (data *FileIdDataMessage) Read(fieldDefinitionNumber uint, buffer []byte, byteOrder binary.ByteOrder) error {
    switch fieldDefinitionNumber {
    case 0:
		data.Type = new(file)
        *data.Type = file(buffer[0])
    case 1:
		data.Manufacturer = new(manufacturer)
		*data.Manufacturer = manufacturer(binary.LittleEndian.Uint16(buffer))
    case 2:
		data.Product = new(uint16)
		*data.Product = uint16(binary.LittleEndian.Uint16(buffer))
    case 3:
		data.SerialNumber = new(uint32)
		*data.SerialNumber = uint32(binary.LittleEndian.Uint32(buffer))
    case 4:
		data.TimeCreated = new(date_time)
		*data.TimeCreated = date_time(binary.LittleEndian.Uint32(buffer))
    case 5:
		data.Number = new(uint16)
		*data.Number = uint16(binary.LittleEndian.Uint16(buffer))
	default:
		return fmt.Errorf("unknown field number (%d) for message %s", fieldDefinitionNumber, "file_id")
    }
	return nil
}

func (data *FileCreatorDataMessage) Read(fieldDefinitionNumber uint, buffer []byte, byteOrder binary.ByteOrder) error {
    switch fieldDefinitionNumber {
    case 0:
		data.SoftwareVersion = new(uint16)
		*data.SoftwareVersion = uint16(binary.LittleEndian.Uint16(buffer))
    case 1:
		data.HardwareVersion = new(uint8)
        *data.HardwareVersion = uint8(buffer[0])
	default:
		return fmt.Errorf("unknown field number (%d) for message %s", fieldDefinitionNumber, "file_creator")
    }
	return nil
}

func (data *SoftwareDataMessage) Read(fieldDefinitionNumber uint, buffer []byte, byteOrder binary.ByteOrder) error {
    switch fieldDefinitionNumber {
    case 254:
		data.MessageIndex = new(message_index)
		*data.MessageIndex = message_index(binary.LittleEndian.Uint16(buffer))
    case 3:
		data.Version = new(uint16)
		*data.Version = uint16(binary.LittleEndian.Uint16(buffer))
    case 5:
		data.PartNumber = new(string)
        *data.PartNumber = string(buffer[0])
	default:
		return fmt.Errorf("unknown field number (%d) for message %s", fieldDefinitionNumber, "software")
    }
	return nil
}

func (data *SlaveDeviceDataMessage) Read(fieldDefinitionNumber uint, buffer []byte, byteOrder binary.ByteOrder) error {
    switch fieldDefinitionNumber {
    case 0:
		data.Manufacturer = new(manufacturer)
		*data.Manufacturer = manufacturer(binary.LittleEndian.Uint16(buffer))
    case 1:
		data.Product = new(uint16)
		*data.Product = uint16(binary.LittleEndian.Uint16(buffer))
	default:
		return fmt.Errorf("unknown field number (%d) for message %s", fieldDefinitionNumber, "slave_device")
    }
	return nil
}

func (data *CapabilitiesDataMessage) Read(fieldDefinitionNumber uint, buffer []byte, byteOrder binary.ByteOrder) error {
    switch fieldDefinitionNumber {
    case 0:
		data.Languages = new(uint8)
        *data.Languages = uint8(buffer[0])
    case 1:
		data.Sports = new(sport_bits_0)
        *data.Sports = sport_bits_0(buffer[0])
    case 21:
		data.WorkoutsSupported = new(workout_capabilities)
		*data.WorkoutsSupported = workout_capabilities(binary.LittleEndian.Uint32(buffer))
    case 23:
		data.ConnectivitySupported = new(connectivity_capabilities)
		*data.ConnectivitySupported = connectivity_capabilities(binary.LittleEndian.Uint32(buffer))
	default:
		return fmt.Errorf("unknown field number (%d) for message %s", fieldDefinitionNumber, "capabilities")
    }
	return nil
}

func (data *FileCapabilitiesDataMessage) Read(fieldDefinitionNumber uint, buffer []byte, byteOrder binary.ByteOrder) error {
    switch fieldDefinitionNumber {
    case 254:
		data.MessageIndex = new(message_index)
		*data.MessageIndex = message_index(binary.LittleEndian.Uint16(buffer))
    case 0:
		data.Type = new(file)
        *data.Type = file(buffer[0])
    case 1:
		data.Flags = new(file_flags)
        *data.Flags = file_flags(buffer[0])
    case 2:
		data.Directory = new(string)
        *data.Directory = string(buffer[0])
    case 3:
		data.MaxCount = new(uint16)
		*data.MaxCount = uint16(binary.LittleEndian.Uint16(buffer))
    case 4:
		data.MaxSize = new(uint32)
		*data.MaxSize = uint32(binary.LittleEndian.Uint32(buffer))
	default:
		return fmt.Errorf("unknown field number (%d) for message %s", fieldDefinitionNumber, "file_capabilities")
    }
	return nil
}

func (data *MesgCapabilitiesDataMessage) Read(fieldDefinitionNumber uint, buffer []byte, byteOrder binary.ByteOrder) error {
    switch fieldDefinitionNumber {
    case 254:
		data.MessageIndex = new(message_index)
		*data.MessageIndex = message_index(binary.LittleEndian.Uint16(buffer))
    case 0:
		data.File = new(file)
        *data.File = file(buffer[0])
    case 1:
		data.MesgNum = new(mesg_num)
		*data.MesgNum = mesg_num(binary.LittleEndian.Uint16(buffer))
    case 2:
		data.CountType = new(mesg_count)
        *data.CountType = mesg_count(buffer[0])
    case 3:
		data.Count = new(uint16)
		*data.Count = uint16(binary.LittleEndian.Uint16(buffer))
	default:
		return fmt.Errorf("unknown field number (%d) for message %s", fieldDefinitionNumber, "mesg_capabilities")
    }
	return nil
}

func (data *FieldCapabilitiesDataMessage) Read(fieldDefinitionNumber uint, buffer []byte, byteOrder binary.ByteOrder) error {
    switch fieldDefinitionNumber {
    case 254:
		data.MessageIndex = new(message_index)
		*data.MessageIndex = message_index(binary.LittleEndian.Uint16(buffer))
    case 0:
		data.File = new(file)
        *data.File = file(buffer[0])
    case 1:
		data.MesgNum = new(mesg_num)
		*data.MesgNum = mesg_num(binary.LittleEndian.Uint16(buffer))
    case 2:
		data.FieldNum = new(uint8)
        *data.FieldNum = uint8(buffer[0])
    case 3:
		data.Count = new(uint16)
		*data.Count = uint16(binary.LittleEndian.Uint16(buffer))
	default:
		return fmt.Errorf("unknown field number (%d) for message %s", fieldDefinitionNumber, "field_capabilities")
    }
	return nil
}

func (data *DeviceSettingsDataMessage) Read(fieldDefinitionNumber uint, buffer []byte, byteOrder binary.ByteOrder) error {
    switch fieldDefinitionNumber {
    case 0:
		data.ActiveTimeZone = new(uint8)
        *data.ActiveTimeZone = uint8(buffer[0])
    case 1:
		data.UtcOffset = new(uint32)
		*data.UtcOffset = uint32(binary.LittleEndian.Uint32(buffer))
    case 5:
		data.TimeZoneOffset = new(int8)
        *data.TimeZoneOffset = int8(buffer[0])
	default:
		return fmt.Errorf("unknown field number (%d) for message %s", fieldDefinitionNumber, "device_settings")
    }
	return nil
}

func (data *UserProfileDataMessage) Read(fieldDefinitionNumber uint, buffer []byte, byteOrder binary.ByteOrder) error {
    switch fieldDefinitionNumber {
    case 254:
		data.MessageIndex = new(message_index)
		*data.MessageIndex = message_index(binary.LittleEndian.Uint16(buffer))
    case 0:
		data.FriendlyName = new(string)
        *data.FriendlyName = string(buffer[0])
    case 1:
		data.Gender = new(gender)
        *data.Gender = gender(buffer[0])
    case 2:
		data.Age = new(uint8)
        *data.Age = uint8(buffer[0])
    case 3:
		data.Height = new(uint8)
        *data.Height = uint8(buffer[0])
    case 4:
		data.Weight = new(uint16)
		*data.Weight = uint16(binary.LittleEndian.Uint16(buffer))
    case 5:
		data.Language = new(language)
        *data.Language = language(buffer[0])
    case 6:
		data.ElevSetting = new(display_measure)
        *data.ElevSetting = display_measure(buffer[0])
    case 7:
		data.WeightSetting = new(display_measure)
        *data.WeightSetting = display_measure(buffer[0])
    case 8:
		data.RestingHeartRate = new(uint8)
        *data.RestingHeartRate = uint8(buffer[0])
    case 9:
		data.DefaultMaxRunningHeartRate = new(uint8)
        *data.DefaultMaxRunningHeartRate = uint8(buffer[0])
    case 10:
		data.DefaultMaxBikingHeartRate = new(uint8)
        *data.DefaultMaxBikingHeartRate = uint8(buffer[0])
    case 11:
		data.DefaultMaxHeartRate = new(uint8)
        *data.DefaultMaxHeartRate = uint8(buffer[0])
    case 12:
		data.HrSetting = new(display_heart)
        *data.HrSetting = display_heart(buffer[0])
    case 13:
		data.SpeedSetting = new(display_measure)
        *data.SpeedSetting = display_measure(buffer[0])
    case 14:
		data.DistSetting = new(display_measure)
        *data.DistSetting = display_measure(buffer[0])
    case 16:
		data.PowerSetting = new(display_power)
        *data.PowerSetting = display_power(buffer[0])
    case 17:
		data.ActivityClass = new(activity_class)
        *data.ActivityClass = activity_class(buffer[0])
    case 18:
		data.PositionSetting = new(display_position)
        *data.PositionSetting = display_position(buffer[0])
    case 21:
		data.TemperatureSetting = new(display_measure)
        *data.TemperatureSetting = display_measure(buffer[0])
    case 22:
		data.LocalId = new(user_local_id)
		*data.LocalId = user_local_id(binary.LittleEndian.Uint16(buffer))
    case 23:
		data.GlobalId = new(byte)
        *data.GlobalId = byte(buffer[0])
    case 30:
		data.HeightSetting = new(display_measure)
        *data.HeightSetting = display_measure(buffer[0])
	default:
		return fmt.Errorf("unknown field number (%d) for message %s", fieldDefinitionNumber, "user_profile")
    }
	return nil
}

func (data *HrmProfileDataMessage) Read(fieldDefinitionNumber uint, buffer []byte, byteOrder binary.ByteOrder) error {
    switch fieldDefinitionNumber {
    case 254:
		data.MessageIndex = new(message_index)
		*data.MessageIndex = message_index(binary.LittleEndian.Uint16(buffer))
    case 0:
		data.Enabled = new(bool)
    case 1:
		data.HrmAntId = new(uint16)
		*data.HrmAntId = uint16(binary.LittleEndian.Uint16(buffer))
    case 2:
		data.LogHrv = new(bool)
    case 3:
		data.HrmAntIdTransType = new(uint8)
        *data.HrmAntIdTransType = uint8(buffer[0])
	default:
		return fmt.Errorf("unknown field number (%d) for message %s", fieldDefinitionNumber, "hrm_profile")
    }
	return nil
}

func (data *SdmProfileDataMessage) Read(fieldDefinitionNumber uint, buffer []byte, byteOrder binary.ByteOrder) error {
    switch fieldDefinitionNumber {
    case 254:
		data.MessageIndex = new(message_index)
		*data.MessageIndex = message_index(binary.LittleEndian.Uint16(buffer))
    case 0:
		data.Enabled = new(bool)
    case 1:
		data.SdmAntId = new(uint16)
		*data.SdmAntId = uint16(binary.LittleEndian.Uint16(buffer))
    case 2:
		data.SdmCalFactor = new(uint16)
		*data.SdmCalFactor = uint16(binary.LittleEndian.Uint16(buffer))
    case 3:
		data.Odometer = new(uint32)
		*data.Odometer = uint32(binary.LittleEndian.Uint32(buffer))
    case 4:
		data.SpeedSource = new(bool)
    case 5:
		data.SdmAntIdTransType = new(uint8)
        *data.SdmAntIdTransType = uint8(buffer[0])
    case 7:
		data.OdometerRollover = new(uint8)
        *data.OdometerRollover = uint8(buffer[0])
	default:
		return fmt.Errorf("unknown field number (%d) for message %s", fieldDefinitionNumber, "sdm_profile")
    }
	return nil
}

func (data *BikeProfileDataMessage) Read(fieldDefinitionNumber uint, buffer []byte, byteOrder binary.ByteOrder) error {
    switch fieldDefinitionNumber {
    case 254:
		data.MessageIndex = new(message_index)
		*data.MessageIndex = message_index(binary.LittleEndian.Uint16(buffer))
    case 0:
		data.Name = new(string)
        *data.Name = string(buffer[0])
    case 1:
		data.Sport = new(sport)
        *data.Sport = sport(buffer[0])
    case 2:
		data.SubSport = new(sub_sport)
        *data.SubSport = sub_sport(buffer[0])
    case 3:
		data.Odometer = new(uint32)
		*data.Odometer = uint32(binary.LittleEndian.Uint32(buffer))
    case 4:
		data.BikeSpdAntId = new(uint16)
		*data.BikeSpdAntId = uint16(binary.LittleEndian.Uint16(buffer))
    case 5:
		data.BikeCadAntId = new(uint16)
		*data.BikeCadAntId = uint16(binary.LittleEndian.Uint16(buffer))
    case 6:
		data.BikeSpdcadAntId = new(uint16)
		*data.BikeSpdcadAntId = uint16(binary.LittleEndian.Uint16(buffer))
    case 7:
		data.BikePowerAntId = new(uint16)
		*data.BikePowerAntId = uint16(binary.LittleEndian.Uint16(buffer))
    case 8:
		data.CustomWheelsize = new(uint16)
		*data.CustomWheelsize = uint16(binary.LittleEndian.Uint16(buffer))
    case 9:
		data.AutoWheelsize = new(uint16)
		*data.AutoWheelsize = uint16(binary.LittleEndian.Uint16(buffer))
    case 10:
		data.BikeWeight = new(uint16)
		*data.BikeWeight = uint16(binary.LittleEndian.Uint16(buffer))
    case 11:
		data.PowerCalFactor = new(uint16)
		*data.PowerCalFactor = uint16(binary.LittleEndian.Uint16(buffer))
    case 12:
		data.AutoWheelCal = new(bool)
    case 13:
		data.AutoPowerZero = new(bool)
    case 14:
		data.Id = new(uint8)
        *data.Id = uint8(buffer[0])
    case 15:
		data.SpdEnabled = new(bool)
    case 16:
		data.CadEnabled = new(bool)
    case 17:
		data.SpdcadEnabled = new(bool)
    case 18:
		data.PowerEnabled = new(bool)
    case 19:
		data.CrankLength = new(uint8)
        *data.CrankLength = uint8(buffer[0])
    case 20:
		data.Enabled = new(bool)
    case 21:
		data.BikeSpdAntIdTransType = new(uint8)
        *data.BikeSpdAntIdTransType = uint8(buffer[0])
    case 22:
		data.BikeCadAntIdTransType = new(uint8)
        *data.BikeCadAntIdTransType = uint8(buffer[0])
    case 23:
		data.BikeSpdcadAntIdTransType = new(uint8)
        *data.BikeSpdcadAntIdTransType = uint8(buffer[0])
    case 24:
		data.BikePowerAntIdTransType = new(uint8)
        *data.BikePowerAntIdTransType = uint8(buffer[0])
    case 37:
		data.OdometerRollover = new(uint8)
        *data.OdometerRollover = uint8(buffer[0])
    case 38:
		data.FrontGearNum = new(uint8)
        *data.FrontGearNum = uint8(buffer[0])
    case 39:
		data.FrontGear = new(uint8)
        *data.FrontGear = uint8(buffer[0])
    case 40:
		data.RearGearNum = new(uint8)
        *data.RearGearNum = uint8(buffer[0])
    case 41:
		data.RearGear = new(uint8)
        *data.RearGear = uint8(buffer[0])
    case 44:
		data.ShimanoDi2Enabled = new(bool)
	default:
		return fmt.Errorf("unknown field number (%d) for message %s", fieldDefinitionNumber, "bike_profile")
    }
	return nil
}

func (data *ZonesTargetDataMessage) Read(fieldDefinitionNumber uint, buffer []byte, byteOrder binary.ByteOrder) error {
    switch fieldDefinitionNumber {
    case 1:
		data.MaxHeartRate = new(uint8)
        *data.MaxHeartRate = uint8(buffer[0])
    case 2:
		data.ThresholdHeartRate = new(uint8)
        *data.ThresholdHeartRate = uint8(buffer[0])
    case 3:
		data.FunctionalThresholdPower = new(uint16)
		*data.FunctionalThresholdPower = uint16(binary.LittleEndian.Uint16(buffer))
    case 5:
		data.HrCalcType = new(hr_zone_calc)
        *data.HrCalcType = hr_zone_calc(buffer[0])
    case 7:
		data.PwrCalcType = new(pwr_zone_calc)
        *data.PwrCalcType = pwr_zone_calc(buffer[0])
	default:
		return fmt.Errorf("unknown field number (%d) for message %s", fieldDefinitionNumber, "zones_target")
    }
	return nil
}

func (data *SportDataMessage) Read(fieldDefinitionNumber uint, buffer []byte, byteOrder binary.ByteOrder) error {
    switch fieldDefinitionNumber {
    case 0:
		data.Sport = new(sport)
        *data.Sport = sport(buffer[0])
    case 1:
		data.SubSport = new(sub_sport)
        *data.SubSport = sub_sport(buffer[0])
    case 3:
		data.Name = new(string)
        *data.Name = string(buffer[0])
	default:
		return fmt.Errorf("unknown field number (%d) for message %s", fieldDefinitionNumber, "sport")
    }
	return nil
}

func (data *HrZoneDataMessage) Read(fieldDefinitionNumber uint, buffer []byte, byteOrder binary.ByteOrder) error {
    switch fieldDefinitionNumber {
    case 254:
		data.MessageIndex = new(message_index)
		*data.MessageIndex = message_index(binary.LittleEndian.Uint16(buffer))
    case 1:
		data.HighBpm = new(uint8)
        *data.HighBpm = uint8(buffer[0])
    case 2:
		data.Name = new(string)
        *data.Name = string(buffer[0])
	default:
		return fmt.Errorf("unknown field number (%d) for message %s", fieldDefinitionNumber, "hr_zone")
    }
	return nil
}

func (data *SpeedZoneDataMessage) Read(fieldDefinitionNumber uint, buffer []byte, byteOrder binary.ByteOrder) error {
    switch fieldDefinitionNumber {
    case 254:
		data.MessageIndex = new(message_index)
		*data.MessageIndex = message_index(binary.LittleEndian.Uint16(buffer))
    case 0:
		data.HighValue = new(uint16)
		*data.HighValue = uint16(binary.LittleEndian.Uint16(buffer))
    case 1:
		data.Name = new(string)
        *data.Name = string(buffer[0])
	default:
		return fmt.Errorf("unknown field number (%d) for message %s", fieldDefinitionNumber, "speed_zone")
    }
	return nil
}

func (data *CadenceZoneDataMessage) Read(fieldDefinitionNumber uint, buffer []byte, byteOrder binary.ByteOrder) error {
    switch fieldDefinitionNumber {
    case 254:
		data.MessageIndex = new(message_index)
		*data.MessageIndex = message_index(binary.LittleEndian.Uint16(buffer))
    case 0:
		data.HighValue = new(uint8)
        *data.HighValue = uint8(buffer[0])
    case 1:
		data.Name = new(string)
        *data.Name = string(buffer[0])
	default:
		return fmt.Errorf("unknown field number (%d) for message %s", fieldDefinitionNumber, "cadence_zone")
    }
	return nil
}

func (data *PowerZoneDataMessage) Read(fieldDefinitionNumber uint, buffer []byte, byteOrder binary.ByteOrder) error {
    switch fieldDefinitionNumber {
    case 254:
		data.MessageIndex = new(message_index)
		*data.MessageIndex = message_index(binary.LittleEndian.Uint16(buffer))
    case 1:
		data.HighValue = new(uint16)
		*data.HighValue = uint16(binary.LittleEndian.Uint16(buffer))
    case 2:
		data.Name = new(string)
        *data.Name = string(buffer[0])
	default:
		return fmt.Errorf("unknown field number (%d) for message %s", fieldDefinitionNumber, "power_zone")
    }
	return nil
}

func (data *MetZoneDataMessage) Read(fieldDefinitionNumber uint, buffer []byte, byteOrder binary.ByteOrder) error {
    switch fieldDefinitionNumber {
    case 254:
		data.MessageIndex = new(message_index)
		*data.MessageIndex = message_index(binary.LittleEndian.Uint16(buffer))
    case 1:
		data.HighBpm = new(uint8)
        *data.HighBpm = uint8(buffer[0])
    case 2:
		data.Calories = new(uint16)
		*data.Calories = uint16(binary.LittleEndian.Uint16(buffer))
    case 3:
		data.FatCalories = new(uint8)
        *data.FatCalories = uint8(buffer[0])
	default:
		return fmt.Errorf("unknown field number (%d) for message %s", fieldDefinitionNumber, "met_zone")
    }
	return nil
}

func (data *GoalDataMessage) Read(fieldDefinitionNumber uint, buffer []byte, byteOrder binary.ByteOrder) error {
    switch fieldDefinitionNumber {
    case 254:
		data.MessageIndex = new(message_index)
		*data.MessageIndex = message_index(binary.LittleEndian.Uint16(buffer))
    case 0:
		data.Sport = new(sport)
        *data.Sport = sport(buffer[0])
    case 1:
		data.SubSport = new(sub_sport)
        *data.SubSport = sub_sport(buffer[0])
    case 2:
		data.StartDate = new(date_time)
		*data.StartDate = date_time(binary.LittleEndian.Uint32(buffer))
    case 3:
		data.EndDate = new(date_time)
		*data.EndDate = date_time(binary.LittleEndian.Uint32(buffer))
    case 4:
		data.Type = new(goal)
        *data.Type = goal(buffer[0])
    case 5:
		data.Value = new(uint32)
		*data.Value = uint32(binary.LittleEndian.Uint32(buffer))
    case 6:
		data.Repeat = new(bool)
    case 7:
		data.TargetValue = new(uint32)
		*data.TargetValue = uint32(binary.LittleEndian.Uint32(buffer))
    case 8:
		data.Recurrence = new(goal_recurrence)
        *data.Recurrence = goal_recurrence(buffer[0])
    case 9:
		data.RecurrenceValue = new(uint16)
		*data.RecurrenceValue = uint16(binary.LittleEndian.Uint16(buffer))
    case 10:
		data.Enabled = new(bool)
	default:
		return fmt.Errorf("unknown field number (%d) for message %s", fieldDefinitionNumber, "goal")
    }
	return nil
}

func (data *ActivityDataMessage) Read(fieldDefinitionNumber uint, buffer []byte, byteOrder binary.ByteOrder) error {
    switch fieldDefinitionNumber {
    case 253:
		data.Timestamp = new(date_time)
		*data.Timestamp = date_time(binary.LittleEndian.Uint32(buffer))
    case 0:
		data.TotalTimerTime = new(uint32)
		*data.TotalTimerTime = uint32(binary.LittleEndian.Uint32(buffer))
    case 1:
		data.NumSessions = new(uint16)
		*data.NumSessions = uint16(binary.LittleEndian.Uint16(buffer))
    case 2:
		data.Type = new(activity)
        *data.Type = activity(buffer[0])
    case 3:
		data.Event = new(event)
        *data.Event = event(buffer[0])
    case 4:
		data.EventType = new(event_type)
        *data.EventType = event_type(buffer[0])
    case 5:
		data.LocalTimestamp = new(local_date_time)
		*data.LocalTimestamp = local_date_time(binary.LittleEndian.Uint32(buffer))
    case 6:
		data.EventGroup = new(uint8)
        *data.EventGroup = uint8(buffer[0])
	default:
		return fmt.Errorf("unknown field number (%d) for message %s", fieldDefinitionNumber, "activity")
    }
	return nil
}

func (data *SessionDataMessage) Read(fieldDefinitionNumber uint, buffer []byte, byteOrder binary.ByteOrder) error {
    switch fieldDefinitionNumber {
    case 254:
		data.MessageIndex = new(message_index)
		*data.MessageIndex = message_index(binary.LittleEndian.Uint16(buffer))
    case 253:
		data.Timestamp = new(date_time)
		*data.Timestamp = date_time(binary.LittleEndian.Uint32(buffer))
    case 0:
		data.Event = new(event)
        *data.Event = event(buffer[0])
    case 1:
		data.EventType = new(event_type)
        *data.EventType = event_type(buffer[0])
    case 2:
		data.StartTime = new(date_time)
		*data.StartTime = date_time(binary.LittleEndian.Uint32(buffer))
    case 3:
		data.StartPositionLat = new(int32)
		*data.StartPositionLat = int32(binary.LittleEndian.Uint32(buffer))
    case 4:
		data.StartPositionLong = new(int32)
		*data.StartPositionLong = int32(binary.LittleEndian.Uint32(buffer))
    case 5:
		data.Sport = new(sport)
        *data.Sport = sport(buffer[0])
    case 6:
		data.SubSport = new(sub_sport)
        *data.SubSport = sub_sport(buffer[0])
    case 7:
		data.TotalElapsedTime = new(uint32)
		*data.TotalElapsedTime = uint32(binary.LittleEndian.Uint32(buffer))
    case 8:
		data.TotalTimerTime = new(uint32)
		*data.TotalTimerTime = uint32(binary.LittleEndian.Uint32(buffer))
    case 9:
		data.TotalDistance = new(uint32)
		*data.TotalDistance = uint32(binary.LittleEndian.Uint32(buffer))
    case 10:
		data.TotalCycles = new(uint32)
		*data.TotalCycles = uint32(binary.LittleEndian.Uint32(buffer))
    case 11:
		data.TotalCalories = new(uint16)
		*data.TotalCalories = uint16(binary.LittleEndian.Uint16(buffer))
    case 13:
		data.TotalFatCalories = new(uint16)
		*data.TotalFatCalories = uint16(binary.LittleEndian.Uint16(buffer))
    case 14:
		data.AvgSpeed = new(uint16)
		*data.AvgSpeed = uint16(binary.LittleEndian.Uint16(buffer))
    case 15:
		data.MaxSpeed = new(uint16)
		*data.MaxSpeed = uint16(binary.LittleEndian.Uint16(buffer))
    case 16:
		data.AvgHeartRate = new(uint8)
        *data.AvgHeartRate = uint8(buffer[0])
    case 17:
		data.MaxHeartRate = new(uint8)
        *data.MaxHeartRate = uint8(buffer[0])
    case 18:
		data.AvgCadence = new(uint8)
        *data.AvgCadence = uint8(buffer[0])
    case 19:
		data.MaxCadence = new(uint8)
        *data.MaxCadence = uint8(buffer[0])
    case 20:
		data.AvgPower = new(uint16)
		*data.AvgPower = uint16(binary.LittleEndian.Uint16(buffer))
    case 21:
		data.MaxPower = new(uint16)
		*data.MaxPower = uint16(binary.LittleEndian.Uint16(buffer))
    case 22:
		data.TotalAscent = new(uint16)
		*data.TotalAscent = uint16(binary.LittleEndian.Uint16(buffer))
    case 23:
		data.TotalDescent = new(uint16)
		*data.TotalDescent = uint16(binary.LittleEndian.Uint16(buffer))
    case 24:
		data.TotalTrainingEffect = new(uint8)
        *data.TotalTrainingEffect = uint8(buffer[0])
    case 25:
		data.FirstLapIndex = new(uint16)
		*data.FirstLapIndex = uint16(binary.LittleEndian.Uint16(buffer))
    case 26:
		data.NumLaps = new(uint16)
		*data.NumLaps = uint16(binary.LittleEndian.Uint16(buffer))
    case 27:
		data.EventGroup = new(uint8)
        *data.EventGroup = uint8(buffer[0])
    case 28:
		data.Trigger = new(session_trigger)
        *data.Trigger = session_trigger(buffer[0])
    case 29:
		data.NecLat = new(int32)
		*data.NecLat = int32(binary.LittleEndian.Uint32(buffer))
    case 30:
		data.NecLong = new(int32)
		*data.NecLong = int32(binary.LittleEndian.Uint32(buffer))
    case 31:
		data.SwcLat = new(int32)
		*data.SwcLat = int32(binary.LittleEndian.Uint32(buffer))
    case 32:
		data.SwcLong = new(int32)
		*data.SwcLong = int32(binary.LittleEndian.Uint32(buffer))
    case 34:
		data.NormalizedPower = new(uint16)
		*data.NormalizedPower = uint16(binary.LittleEndian.Uint16(buffer))
    case 35:
		data.TrainingStressScore = new(uint16)
		*data.TrainingStressScore = uint16(binary.LittleEndian.Uint16(buffer))
    case 36:
		data.IntensityFactor = new(uint16)
		*data.IntensityFactor = uint16(binary.LittleEndian.Uint16(buffer))
    case 37:
		data.LeftRightBalance = new(left_right_balance_100)
		*data.LeftRightBalance = left_right_balance_100(binary.LittleEndian.Uint16(buffer))
    case 41:
		data.AvgStrokeCount = new(uint32)
		*data.AvgStrokeCount = uint32(binary.LittleEndian.Uint32(buffer))
    case 42:
		data.AvgStrokeDistance = new(uint16)
		*data.AvgStrokeDistance = uint16(binary.LittleEndian.Uint16(buffer))
    case 43:
		data.SwimStroke = new(swim_stroke)
        *data.SwimStroke = swim_stroke(buffer[0])
    case 44:
		data.PoolLength = new(uint16)
		*data.PoolLength = uint16(binary.LittleEndian.Uint16(buffer))
    case 45:
		data.ThresholdPower = new(uint16)
		*data.ThresholdPower = uint16(binary.LittleEndian.Uint16(buffer))
    case 46:
		data.PoolLengthUnit = new(display_measure)
        *data.PoolLengthUnit = display_measure(buffer[0])
    case 47:
		data.NumActiveLengths = new(uint16)
		*data.NumActiveLengths = uint16(binary.LittleEndian.Uint16(buffer))
    case 48:
		data.TotalWork = new(uint32)
		*data.TotalWork = uint32(binary.LittleEndian.Uint32(buffer))
    case 49:
		data.AvgAltitude = new(uint16)
		*data.AvgAltitude = uint16(binary.LittleEndian.Uint16(buffer))
    case 50:
		data.MaxAltitude = new(uint16)
		*data.MaxAltitude = uint16(binary.LittleEndian.Uint16(buffer))
    case 51:
		data.GpsAccuracy = new(uint8)
        *data.GpsAccuracy = uint8(buffer[0])
    case 52:
		data.AvgGrade = new(int16)
		*data.AvgGrade = int16(binary.LittleEndian.Uint16(buffer))
    case 53:
		data.AvgPosGrade = new(int16)
		*data.AvgPosGrade = int16(binary.LittleEndian.Uint16(buffer))
    case 54:
		data.AvgNegGrade = new(int16)
		*data.AvgNegGrade = int16(binary.LittleEndian.Uint16(buffer))
    case 55:
		data.MaxPosGrade = new(int16)
		*data.MaxPosGrade = int16(binary.LittleEndian.Uint16(buffer))
    case 56:
		data.MaxNegGrade = new(int16)
		*data.MaxNegGrade = int16(binary.LittleEndian.Uint16(buffer))
    case 57:
		data.AvgTemperature = new(int8)
        *data.AvgTemperature = int8(buffer[0])
    case 58:
		data.MaxTemperature = new(int8)
        *data.MaxTemperature = int8(buffer[0])
    case 59:
		data.TotalMovingTime = new(uint32)
		*data.TotalMovingTime = uint32(binary.LittleEndian.Uint32(buffer))
    case 60:
		data.AvgPosVerticalSpeed = new(int16)
		*data.AvgPosVerticalSpeed = int16(binary.LittleEndian.Uint16(buffer))
    case 61:
		data.AvgNegVerticalSpeed = new(int16)
		*data.AvgNegVerticalSpeed = int16(binary.LittleEndian.Uint16(buffer))
    case 62:
		data.MaxPosVerticalSpeed = new(int16)
		*data.MaxPosVerticalSpeed = int16(binary.LittleEndian.Uint16(buffer))
    case 63:
		data.MaxNegVerticalSpeed = new(int16)
		*data.MaxNegVerticalSpeed = int16(binary.LittleEndian.Uint16(buffer))
    case 64:
		data.MinHeartRate = new(uint8)
        *data.MinHeartRate = uint8(buffer[0])
    case 65:
		data.TimeInHrZone = new(uint32)
		*data.TimeInHrZone = uint32(binary.LittleEndian.Uint32(buffer))
    case 66:
		data.TimeInSpeedZone = new(uint32)
		*data.TimeInSpeedZone = uint32(binary.LittleEndian.Uint32(buffer))
    case 67:
		data.TimeInCadenceZone = new(uint32)
		*data.TimeInCadenceZone = uint32(binary.LittleEndian.Uint32(buffer))
    case 68:
		data.TimeInPowerZone = new(uint32)
		*data.TimeInPowerZone = uint32(binary.LittleEndian.Uint32(buffer))
    case 69:
		data.AvgLapTime = new(uint32)
		*data.AvgLapTime = uint32(binary.LittleEndian.Uint32(buffer))
    case 70:
		data.BestLapIndex = new(uint16)
		*data.BestLapIndex = uint16(binary.LittleEndian.Uint16(buffer))
    case 71:
		data.MinAltitude = new(uint16)
		*data.MinAltitude = uint16(binary.LittleEndian.Uint16(buffer))
    case 82:
		data.PlayerScore = new(uint16)
		*data.PlayerScore = uint16(binary.LittleEndian.Uint16(buffer))
    case 83:
		data.OpponentScore = new(uint16)
		*data.OpponentScore = uint16(binary.LittleEndian.Uint16(buffer))
    case 84:
		data.OpponentName = new(string)
        *data.OpponentName = string(buffer[0])
    case 85:
		data.StrokeCount = new(uint16)
		*data.StrokeCount = uint16(binary.LittleEndian.Uint16(buffer))
    case 86:
		data.ZoneCount = new(uint16)
		*data.ZoneCount = uint16(binary.LittleEndian.Uint16(buffer))
    case 87:
		data.MaxBallSpeed = new(uint16)
		*data.MaxBallSpeed = uint16(binary.LittleEndian.Uint16(buffer))
    case 88:
		data.AvgBallSpeed = new(uint16)
		*data.AvgBallSpeed = uint16(binary.LittleEndian.Uint16(buffer))
    case 89:
		data.AvgVerticalOscillation = new(uint16)
		*data.AvgVerticalOscillation = uint16(binary.LittleEndian.Uint16(buffer))
    case 90:
		data.AvgStanceTimePercent = new(uint16)
		*data.AvgStanceTimePercent = uint16(binary.LittleEndian.Uint16(buffer))
    case 91:
		data.AvgStanceTime = new(uint16)
		*data.AvgStanceTime = uint16(binary.LittleEndian.Uint16(buffer))
    case 92:
		data.AvgFractionalCadence = new(uint8)
        *data.AvgFractionalCadence = uint8(buffer[0])
    case 93:
		data.MaxFractionalCadence = new(uint8)
        *data.MaxFractionalCadence = uint8(buffer[0])
    case 94:
		data.TotalFractionalCycles = new(uint8)
        *data.TotalFractionalCycles = uint8(buffer[0])
    case 95:
		data.AvgTotalHemoglobinConc = new(uint16)
		*data.AvgTotalHemoglobinConc = uint16(binary.LittleEndian.Uint16(buffer))
    case 96:
		data.MinTotalHemoglobinConc = new(uint16)
		*data.MinTotalHemoglobinConc = uint16(binary.LittleEndian.Uint16(buffer))
    case 97:
		data.MaxTotalHemoglobinConc = new(uint16)
		*data.MaxTotalHemoglobinConc = uint16(binary.LittleEndian.Uint16(buffer))
    case 98:
		data.AvgSaturatedHemoglobinPercent = new(uint16)
		*data.AvgSaturatedHemoglobinPercent = uint16(binary.LittleEndian.Uint16(buffer))
    case 99:
		data.MinSaturatedHemoglobinPercent = new(uint16)
		*data.MinSaturatedHemoglobinPercent = uint16(binary.LittleEndian.Uint16(buffer))
    case 100:
		data.MaxSaturatedHemoglobinPercent = new(uint16)
		*data.MaxSaturatedHemoglobinPercent = uint16(binary.LittleEndian.Uint16(buffer))
    case 101:
		data.AvgLeftTorqueEffectiveness = new(uint8)
        *data.AvgLeftTorqueEffectiveness = uint8(buffer[0])
    case 102:
		data.AvgRightTorqueEffectiveness = new(uint8)
        *data.AvgRightTorqueEffectiveness = uint8(buffer[0])
    case 103:
		data.AvgLeftPedalSmoothness = new(uint8)
        *data.AvgLeftPedalSmoothness = uint8(buffer[0])
    case 104:
		data.AvgRightPedalSmoothness = new(uint8)
        *data.AvgRightPedalSmoothness = uint8(buffer[0])
    case 105:
		data.AvgCombinedPedalSmoothness = new(uint8)
        *data.AvgCombinedPedalSmoothness = uint8(buffer[0])
    case 111:
		data.SportIndex = new(uint8)
        *data.SportIndex = uint8(buffer[0])
    case 112:
		data.TimeStanding = new(uint32)
		*data.TimeStanding = uint32(binary.LittleEndian.Uint32(buffer))
    case 113:
		data.StandCount = new(uint16)
		*data.StandCount = uint16(binary.LittleEndian.Uint16(buffer))
    case 114:
		data.AvgLeftPco = new(int8)
        *data.AvgLeftPco = int8(buffer[0])
    case 115:
		data.AvgRightPco = new(int8)
        *data.AvgRightPco = int8(buffer[0])
    case 116:
		data.AvgLeftPowerPhase = new(uint8)
        *data.AvgLeftPowerPhase = uint8(buffer[0])
    case 117:
		data.AvgLeftPowerPhasePeak = new(uint8)
        *data.AvgLeftPowerPhasePeak = uint8(buffer[0])
    case 118:
		data.AvgRightPowerPhase = new(uint8)
        *data.AvgRightPowerPhase = uint8(buffer[0])
    case 119:
		data.AvgRightPowerPhasePeak = new(uint8)
        *data.AvgRightPowerPhasePeak = uint8(buffer[0])
    case 120:
		data.AvgPowerPosition = new(uint16)
		*data.AvgPowerPosition = uint16(binary.LittleEndian.Uint16(buffer))
    case 121:
		data.MaxPowerPosition = new(uint16)
		*data.MaxPowerPosition = uint16(binary.LittleEndian.Uint16(buffer))
    case 122:
		data.AvgCadencePosition = new(uint8)
        *data.AvgCadencePosition = uint8(buffer[0])
    case 123:
		data.MaxCadencePosition = new(uint8)
        *data.MaxCadencePosition = uint8(buffer[0])
    case 124:
		data.EnhancedAvgSpeed = new(uint32)
		*data.EnhancedAvgSpeed = uint32(binary.LittleEndian.Uint32(buffer))
    case 125:
		data.EnhancedMaxSpeed = new(uint32)
		*data.EnhancedMaxSpeed = uint32(binary.LittleEndian.Uint32(buffer))
    case 126:
		data.EnhancedAvgAltitude = new(uint32)
		*data.EnhancedAvgAltitude = uint32(binary.LittleEndian.Uint32(buffer))
    case 127:
		data.EnhancedMinAltitude = new(uint32)
		*data.EnhancedMinAltitude = uint32(binary.LittleEndian.Uint32(buffer))
    case 128:
		data.EnhancedMaxAltitude = new(uint32)
		*data.EnhancedMaxAltitude = uint32(binary.LittleEndian.Uint32(buffer))
    case 129:
		data.AvgLevMotorPower = new(uint16)
		*data.AvgLevMotorPower = uint16(binary.LittleEndian.Uint16(buffer))
    case 130:
		data.MaxLevMotorPower = new(uint16)
		*data.MaxLevMotorPower = uint16(binary.LittleEndian.Uint16(buffer))
    case 131:
		data.LevBatteryConsumption = new(uint8)
        *data.LevBatteryConsumption = uint8(buffer[0])
	default:
		return fmt.Errorf("unknown field number (%d) for message %s", fieldDefinitionNumber, "session")
    }
	return nil
}

func (data *LapDataMessage) Read(fieldDefinitionNumber uint, buffer []byte, byteOrder binary.ByteOrder) error {
    switch fieldDefinitionNumber {
    case 254:
		data.MessageIndex = new(message_index)
		*data.MessageIndex = message_index(binary.LittleEndian.Uint16(buffer))
    case 253:
		data.Timestamp = new(date_time)
		*data.Timestamp = date_time(binary.LittleEndian.Uint32(buffer))
    case 0:
		data.Event = new(event)
        *data.Event = event(buffer[0])
    case 1:
		data.EventType = new(event_type)
        *data.EventType = event_type(buffer[0])
    case 2:
		data.StartTime = new(date_time)
		*data.StartTime = date_time(binary.LittleEndian.Uint32(buffer))
    case 3:
		data.StartPositionLat = new(int32)
		*data.StartPositionLat = int32(binary.LittleEndian.Uint32(buffer))
    case 4:
		data.StartPositionLong = new(int32)
		*data.StartPositionLong = int32(binary.LittleEndian.Uint32(buffer))
    case 5:
		data.EndPositionLat = new(int32)
		*data.EndPositionLat = int32(binary.LittleEndian.Uint32(buffer))
    case 6:
		data.EndPositionLong = new(int32)
		*data.EndPositionLong = int32(binary.LittleEndian.Uint32(buffer))
    case 7:
		data.TotalElapsedTime = new(uint32)
		*data.TotalElapsedTime = uint32(binary.LittleEndian.Uint32(buffer))
    case 8:
		data.TotalTimerTime = new(uint32)
		*data.TotalTimerTime = uint32(binary.LittleEndian.Uint32(buffer))
    case 9:
		data.TotalDistance = new(uint32)
		*data.TotalDistance = uint32(binary.LittleEndian.Uint32(buffer))
    case 10:
		data.TotalCycles = new(uint32)
		*data.TotalCycles = uint32(binary.LittleEndian.Uint32(buffer))
    case 11:
		data.TotalCalories = new(uint16)
		*data.TotalCalories = uint16(binary.LittleEndian.Uint16(buffer))
    case 12:
		data.TotalFatCalories = new(uint16)
		*data.TotalFatCalories = uint16(binary.LittleEndian.Uint16(buffer))
    case 13:
		data.AvgSpeed = new(uint16)
		*data.AvgSpeed = uint16(binary.LittleEndian.Uint16(buffer))
    case 14:
		data.MaxSpeed = new(uint16)
		*data.MaxSpeed = uint16(binary.LittleEndian.Uint16(buffer))
    case 15:
		data.AvgHeartRate = new(uint8)
        *data.AvgHeartRate = uint8(buffer[0])
    case 16:
		data.MaxHeartRate = new(uint8)
        *data.MaxHeartRate = uint8(buffer[0])
    case 17:
		data.AvgCadence = new(uint8)
        *data.AvgCadence = uint8(buffer[0])
    case 18:
		data.MaxCadence = new(uint8)
        *data.MaxCadence = uint8(buffer[0])
    case 19:
		data.AvgPower = new(uint16)
		*data.AvgPower = uint16(binary.LittleEndian.Uint16(buffer))
    case 20:
		data.MaxPower = new(uint16)
		*data.MaxPower = uint16(binary.LittleEndian.Uint16(buffer))
    case 21:
		data.TotalAscent = new(uint16)
		*data.TotalAscent = uint16(binary.LittleEndian.Uint16(buffer))
    case 22:
		data.TotalDescent = new(uint16)
		*data.TotalDescent = uint16(binary.LittleEndian.Uint16(buffer))
    case 23:
		data.Intensity = new(intensity)
        *data.Intensity = intensity(buffer[0])
    case 24:
		data.LapTrigger = new(lap_trigger)
        *data.LapTrigger = lap_trigger(buffer[0])
    case 25:
		data.Sport = new(sport)
        *data.Sport = sport(buffer[0])
    case 26:
		data.EventGroup = new(uint8)
        *data.EventGroup = uint8(buffer[0])
    case 32:
		data.NumLengths = new(uint16)
		*data.NumLengths = uint16(binary.LittleEndian.Uint16(buffer))
    case 33:
		data.NormalizedPower = new(uint16)
		*data.NormalizedPower = uint16(binary.LittleEndian.Uint16(buffer))
    case 34:
		data.LeftRightBalance = new(left_right_balance_100)
		*data.LeftRightBalance = left_right_balance_100(binary.LittleEndian.Uint16(buffer))
    case 35:
		data.FirstLengthIndex = new(uint16)
		*data.FirstLengthIndex = uint16(binary.LittleEndian.Uint16(buffer))
    case 37:
		data.AvgStrokeDistance = new(uint16)
		*data.AvgStrokeDistance = uint16(binary.LittleEndian.Uint16(buffer))
    case 38:
		data.SwimStroke = new(swim_stroke)
        *data.SwimStroke = swim_stroke(buffer[0])
    case 39:
		data.SubSport = new(sub_sport)
        *data.SubSport = sub_sport(buffer[0])
    case 40:
		data.NumActiveLengths = new(uint16)
		*data.NumActiveLengths = uint16(binary.LittleEndian.Uint16(buffer))
    case 41:
		data.TotalWork = new(uint32)
		*data.TotalWork = uint32(binary.LittleEndian.Uint32(buffer))
    case 42:
		data.AvgAltitude = new(uint16)
		*data.AvgAltitude = uint16(binary.LittleEndian.Uint16(buffer))
    case 43:
		data.MaxAltitude = new(uint16)
		*data.MaxAltitude = uint16(binary.LittleEndian.Uint16(buffer))
    case 44:
		data.GpsAccuracy = new(uint8)
        *data.GpsAccuracy = uint8(buffer[0])
    case 45:
		data.AvgGrade = new(int16)
		*data.AvgGrade = int16(binary.LittleEndian.Uint16(buffer))
    case 46:
		data.AvgPosGrade = new(int16)
		*data.AvgPosGrade = int16(binary.LittleEndian.Uint16(buffer))
    case 47:
		data.AvgNegGrade = new(int16)
		*data.AvgNegGrade = int16(binary.LittleEndian.Uint16(buffer))
    case 48:
		data.MaxPosGrade = new(int16)
		*data.MaxPosGrade = int16(binary.LittleEndian.Uint16(buffer))
    case 49:
		data.MaxNegGrade = new(int16)
		*data.MaxNegGrade = int16(binary.LittleEndian.Uint16(buffer))
    case 50:
		data.AvgTemperature = new(int8)
        *data.AvgTemperature = int8(buffer[0])
    case 51:
		data.MaxTemperature = new(int8)
        *data.MaxTemperature = int8(buffer[0])
    case 52:
		data.TotalMovingTime = new(uint32)
		*data.TotalMovingTime = uint32(binary.LittleEndian.Uint32(buffer))
    case 53:
		data.AvgPosVerticalSpeed = new(int16)
		*data.AvgPosVerticalSpeed = int16(binary.LittleEndian.Uint16(buffer))
    case 54:
		data.AvgNegVerticalSpeed = new(int16)
		*data.AvgNegVerticalSpeed = int16(binary.LittleEndian.Uint16(buffer))
    case 55:
		data.MaxPosVerticalSpeed = new(int16)
		*data.MaxPosVerticalSpeed = int16(binary.LittleEndian.Uint16(buffer))
    case 56:
		data.MaxNegVerticalSpeed = new(int16)
		*data.MaxNegVerticalSpeed = int16(binary.LittleEndian.Uint16(buffer))
    case 57:
		data.TimeInHrZone = new(uint32)
		*data.TimeInHrZone = uint32(binary.LittleEndian.Uint32(buffer))
    case 58:
		data.TimeInSpeedZone = new(uint32)
		*data.TimeInSpeedZone = uint32(binary.LittleEndian.Uint32(buffer))
    case 59:
		data.TimeInCadenceZone = new(uint32)
		*data.TimeInCadenceZone = uint32(binary.LittleEndian.Uint32(buffer))
    case 60:
		data.TimeInPowerZone = new(uint32)
		*data.TimeInPowerZone = uint32(binary.LittleEndian.Uint32(buffer))
    case 61:
		data.RepetitionNum = new(uint16)
		*data.RepetitionNum = uint16(binary.LittleEndian.Uint16(buffer))
    case 62:
		data.MinAltitude = new(uint16)
		*data.MinAltitude = uint16(binary.LittleEndian.Uint16(buffer))
    case 63:
		data.MinHeartRate = new(uint8)
        *data.MinHeartRate = uint8(buffer[0])
    case 71:
		data.WktStepIndex = new(message_index)
		*data.WktStepIndex = message_index(binary.LittleEndian.Uint16(buffer))
    case 74:
		data.OpponentScore = new(uint16)
		*data.OpponentScore = uint16(binary.LittleEndian.Uint16(buffer))
    case 75:
		data.StrokeCount = new(uint16)
		*data.StrokeCount = uint16(binary.LittleEndian.Uint16(buffer))
    case 76:
		data.ZoneCount = new(uint16)
		*data.ZoneCount = uint16(binary.LittleEndian.Uint16(buffer))
    case 77:
		data.AvgVerticalOscillation = new(uint16)
		*data.AvgVerticalOscillation = uint16(binary.LittleEndian.Uint16(buffer))
    case 78:
		data.AvgStanceTimePercent = new(uint16)
		*data.AvgStanceTimePercent = uint16(binary.LittleEndian.Uint16(buffer))
    case 79:
		data.AvgStanceTime = new(uint16)
		*data.AvgStanceTime = uint16(binary.LittleEndian.Uint16(buffer))
    case 80:
		data.AvgFractionalCadence = new(uint8)
        *data.AvgFractionalCadence = uint8(buffer[0])
    case 81:
		data.MaxFractionalCadence = new(uint8)
        *data.MaxFractionalCadence = uint8(buffer[0])
    case 82:
		data.TotalFractionalCycles = new(uint8)
        *data.TotalFractionalCycles = uint8(buffer[0])
    case 83:
		data.PlayerScore = new(uint16)
		*data.PlayerScore = uint16(binary.LittleEndian.Uint16(buffer))
    case 84:
		data.AvgTotalHemoglobinConc = new(uint16)
		*data.AvgTotalHemoglobinConc = uint16(binary.LittleEndian.Uint16(buffer))
    case 85:
		data.MinTotalHemoglobinConc = new(uint16)
		*data.MinTotalHemoglobinConc = uint16(binary.LittleEndian.Uint16(buffer))
    case 86:
		data.MaxTotalHemoglobinConc = new(uint16)
		*data.MaxTotalHemoglobinConc = uint16(binary.LittleEndian.Uint16(buffer))
    case 87:
		data.AvgSaturatedHemoglobinPercent = new(uint16)
		*data.AvgSaturatedHemoglobinPercent = uint16(binary.LittleEndian.Uint16(buffer))
    case 88:
		data.MinSaturatedHemoglobinPercent = new(uint16)
		*data.MinSaturatedHemoglobinPercent = uint16(binary.LittleEndian.Uint16(buffer))
    case 89:
		data.MaxSaturatedHemoglobinPercent = new(uint16)
		*data.MaxSaturatedHemoglobinPercent = uint16(binary.LittleEndian.Uint16(buffer))
    case 91:
		data.AvgLeftTorqueEffectiveness = new(uint8)
        *data.AvgLeftTorqueEffectiveness = uint8(buffer[0])
    case 92:
		data.AvgRightTorqueEffectiveness = new(uint8)
        *data.AvgRightTorqueEffectiveness = uint8(buffer[0])
    case 93:
		data.AvgLeftPedalSmoothness = new(uint8)
        *data.AvgLeftPedalSmoothness = uint8(buffer[0])
    case 94:
		data.AvgRightPedalSmoothness = new(uint8)
        *data.AvgRightPedalSmoothness = uint8(buffer[0])
    case 95:
		data.AvgCombinedPedalSmoothness = new(uint8)
        *data.AvgCombinedPedalSmoothness = uint8(buffer[0])
    case 98:
		data.TimeStanding = new(uint32)
		*data.TimeStanding = uint32(binary.LittleEndian.Uint32(buffer))
    case 99:
		data.StandCount = new(uint16)
		*data.StandCount = uint16(binary.LittleEndian.Uint16(buffer))
    case 100:
		data.AvgLeftPco = new(int8)
        *data.AvgLeftPco = int8(buffer[0])
    case 101:
		data.AvgRightPco = new(int8)
        *data.AvgRightPco = int8(buffer[0])
    case 102:
		data.AvgLeftPowerPhase = new(uint8)
        *data.AvgLeftPowerPhase = uint8(buffer[0])
    case 103:
		data.AvgLeftPowerPhasePeak = new(uint8)
        *data.AvgLeftPowerPhasePeak = uint8(buffer[0])
    case 104:
		data.AvgRightPowerPhase = new(uint8)
        *data.AvgRightPowerPhase = uint8(buffer[0])
    case 105:
		data.AvgRightPowerPhasePeak = new(uint8)
        *data.AvgRightPowerPhasePeak = uint8(buffer[0])
    case 106:
		data.AvgPowerPosition = new(uint16)
		*data.AvgPowerPosition = uint16(binary.LittleEndian.Uint16(buffer))
    case 107:
		data.MaxPowerPosition = new(uint16)
		*data.MaxPowerPosition = uint16(binary.LittleEndian.Uint16(buffer))
    case 108:
		data.AvgCadencePosition = new(uint8)
        *data.AvgCadencePosition = uint8(buffer[0])
    case 109:
		data.MaxCadencePosition = new(uint8)
        *data.MaxCadencePosition = uint8(buffer[0])
    case 110:
		data.EnhancedAvgSpeed = new(uint32)
		*data.EnhancedAvgSpeed = uint32(binary.LittleEndian.Uint32(buffer))
    case 111:
		data.EnhancedMaxSpeed = new(uint32)
		*data.EnhancedMaxSpeed = uint32(binary.LittleEndian.Uint32(buffer))
    case 112:
		data.EnhancedAvgAltitude = new(uint32)
		*data.EnhancedAvgAltitude = uint32(binary.LittleEndian.Uint32(buffer))
    case 113:
		data.EnhancedMinAltitude = new(uint32)
		*data.EnhancedMinAltitude = uint32(binary.LittleEndian.Uint32(buffer))
    case 114:
		data.EnhancedMaxAltitude = new(uint32)
		*data.EnhancedMaxAltitude = uint32(binary.LittleEndian.Uint32(buffer))
    case 115:
		data.AvgLevMotorPower = new(uint16)
		*data.AvgLevMotorPower = uint16(binary.LittleEndian.Uint16(buffer))
    case 116:
		data.MaxLevMotorPower = new(uint16)
		*data.MaxLevMotorPower = uint16(binary.LittleEndian.Uint16(buffer))
    case 117:
		data.LevBatteryConsumption = new(uint8)
        *data.LevBatteryConsumption = uint8(buffer[0])
	default:
		return fmt.Errorf("unknown field number (%d) for message %s", fieldDefinitionNumber, "lap")
    }
	return nil
}

func (data *LengthDataMessage) Read(fieldDefinitionNumber uint, buffer []byte, byteOrder binary.ByteOrder) error {
    switch fieldDefinitionNumber {
    case 254:
		data.MessageIndex = new(message_index)
		*data.MessageIndex = message_index(binary.LittleEndian.Uint16(buffer))
    case 253:
		data.Timestamp = new(date_time)
		*data.Timestamp = date_time(binary.LittleEndian.Uint32(buffer))
    case 0:
		data.Event = new(event)
        *data.Event = event(buffer[0])
    case 1:
		data.EventType = new(event_type)
        *data.EventType = event_type(buffer[0])
    case 2:
		data.StartTime = new(date_time)
		*data.StartTime = date_time(binary.LittleEndian.Uint32(buffer))
    case 3:
		data.TotalElapsedTime = new(uint32)
		*data.TotalElapsedTime = uint32(binary.LittleEndian.Uint32(buffer))
    case 4:
		data.TotalTimerTime = new(uint32)
		*data.TotalTimerTime = uint32(binary.LittleEndian.Uint32(buffer))
    case 5:
		data.TotalStrokes = new(uint16)
		*data.TotalStrokes = uint16(binary.LittleEndian.Uint16(buffer))
    case 6:
		data.AvgSpeed = new(uint16)
		*data.AvgSpeed = uint16(binary.LittleEndian.Uint16(buffer))
    case 7:
		data.SwimStroke = new(swim_stroke)
        *data.SwimStroke = swim_stroke(buffer[0])
    case 9:
		data.AvgSwimmingCadence = new(uint8)
        *data.AvgSwimmingCadence = uint8(buffer[0])
    case 10:
		data.EventGroup = new(uint8)
        *data.EventGroup = uint8(buffer[0])
    case 11:
		data.TotalCalories = new(uint16)
		*data.TotalCalories = uint16(binary.LittleEndian.Uint16(buffer))
    case 12:
		data.LengthType = new(length_type)
        *data.LengthType = length_type(buffer[0])
    case 18:
		data.PlayerScore = new(uint16)
		*data.PlayerScore = uint16(binary.LittleEndian.Uint16(buffer))
    case 19:
		data.OpponentScore = new(uint16)
		*data.OpponentScore = uint16(binary.LittleEndian.Uint16(buffer))
    case 20:
		data.StrokeCount = new(uint16)
		*data.StrokeCount = uint16(binary.LittleEndian.Uint16(buffer))
    case 21:
		data.ZoneCount = new(uint16)
		*data.ZoneCount = uint16(binary.LittleEndian.Uint16(buffer))
	default:
		return fmt.Errorf("unknown field number (%d) for message %s", fieldDefinitionNumber, "length")
    }
	return nil
}

func (data *RecordDataMessage) Read(fieldDefinitionNumber uint, buffer []byte, byteOrder binary.ByteOrder) error {
    switch fieldDefinitionNumber {
    case 253:
		data.Timestamp = new(date_time)
		*data.Timestamp = date_time(binary.LittleEndian.Uint32(buffer))
    case 0:
		data.PositionLat = new(int32)
		*data.PositionLat = int32(binary.LittleEndian.Uint32(buffer))
    case 1:
		data.PositionLong = new(int32)
		*data.PositionLong = int32(binary.LittleEndian.Uint32(buffer))
    case 2:
		data.Altitude = new(uint16)
		*data.Altitude = uint16(binary.LittleEndian.Uint16(buffer))
    case 3:
		data.HeartRate = new(uint8)
        *data.HeartRate = uint8(buffer[0])
    case 4:
		data.Cadence = new(uint8)
        *data.Cadence = uint8(buffer[0])
    case 5:
		data.Distance = new(uint32)
		*data.Distance = uint32(binary.LittleEndian.Uint32(buffer))
    case 6:
		data.Speed = new(uint16)
		*data.Speed = uint16(binary.LittleEndian.Uint16(buffer))
    case 7:
		data.Power = new(uint16)
		*data.Power = uint16(binary.LittleEndian.Uint16(buffer))
    case 8:
		data.CompressedSpeedDistance = new(byte)
        *data.CompressedSpeedDistance = byte(buffer[0])
    case 9:
		data.Grade = new(int16)
		*data.Grade = int16(binary.LittleEndian.Uint16(buffer))
    case 10:
		data.Resistance = new(uint8)
        *data.Resistance = uint8(buffer[0])
    case 11:
		data.TimeFromCourse = new(int32)
		*data.TimeFromCourse = int32(binary.LittleEndian.Uint32(buffer))
    case 12:
		data.CycleLength = new(uint8)
        *data.CycleLength = uint8(buffer[0])
    case 13:
		data.Temperature = new(int8)
        *data.Temperature = int8(buffer[0])
    case 17:
		data.Speed1s = new(uint8)
        *data.Speed1s = uint8(buffer[0])
    case 18:
		data.Cycles = new(uint8)
        *data.Cycles = uint8(buffer[0])
    case 19:
		data.TotalCycles = new(uint32)
		*data.TotalCycles = uint32(binary.LittleEndian.Uint32(buffer))
    case 28:
		data.CompressedAccumulatedPower = new(uint16)
		*data.CompressedAccumulatedPower = uint16(binary.LittleEndian.Uint16(buffer))
    case 29:
		data.AccumulatedPower = new(uint32)
		*data.AccumulatedPower = uint32(binary.LittleEndian.Uint32(buffer))
    case 30:
		data.LeftRightBalance = new(left_right_balance)
        *data.LeftRightBalance = left_right_balance(buffer[0])
    case 31:
		data.GpsAccuracy = new(uint8)
        *data.GpsAccuracy = uint8(buffer[0])
    case 32:
		data.VerticalSpeed = new(int16)
		*data.VerticalSpeed = int16(binary.LittleEndian.Uint16(buffer))
    case 33:
		data.Calories = new(uint16)
		*data.Calories = uint16(binary.LittleEndian.Uint16(buffer))
    case 39:
		data.VerticalOscillation = new(uint16)
		*data.VerticalOscillation = uint16(binary.LittleEndian.Uint16(buffer))
    case 40:
		data.StanceTimePercent = new(uint16)
		*data.StanceTimePercent = uint16(binary.LittleEndian.Uint16(buffer))
    case 41:
		data.StanceTime = new(uint16)
		*data.StanceTime = uint16(binary.LittleEndian.Uint16(buffer))
    case 42:
		data.ActivityType = new(activity_type)
        *data.ActivityType = activity_type(buffer[0])
    case 43:
		data.LeftTorqueEffectiveness = new(uint8)
        *data.LeftTorqueEffectiveness = uint8(buffer[0])
    case 44:
		data.RightTorqueEffectiveness = new(uint8)
        *data.RightTorqueEffectiveness = uint8(buffer[0])
    case 45:
		data.LeftPedalSmoothness = new(uint8)
        *data.LeftPedalSmoothness = uint8(buffer[0])
    case 46:
		data.RightPedalSmoothness = new(uint8)
        *data.RightPedalSmoothness = uint8(buffer[0])
    case 47:
		data.CombinedPedalSmoothness = new(uint8)
        *data.CombinedPedalSmoothness = uint8(buffer[0])
    case 48:
		data.Time128 = new(uint8)
        *data.Time128 = uint8(buffer[0])
    case 49:
		data.StrokeType = new(stroke_type)
        *data.StrokeType = stroke_type(buffer[0])
    case 50:
		data.Zone = new(uint8)
        *data.Zone = uint8(buffer[0])
    case 51:
		data.BallSpeed = new(uint16)
		*data.BallSpeed = uint16(binary.LittleEndian.Uint16(buffer))
    case 52:
		data.Cadence256 = new(uint16)
		*data.Cadence256 = uint16(binary.LittleEndian.Uint16(buffer))
    case 53:
		data.FractionalCadence = new(uint8)
        *data.FractionalCadence = uint8(buffer[0])
    case 54:
		data.TotalHemoglobinConc = new(uint16)
		*data.TotalHemoglobinConc = uint16(binary.LittleEndian.Uint16(buffer))
    case 55:
		data.TotalHemoglobinConcMin = new(uint16)
		*data.TotalHemoglobinConcMin = uint16(binary.LittleEndian.Uint16(buffer))
    case 56:
		data.TotalHemoglobinConcMax = new(uint16)
		*data.TotalHemoglobinConcMax = uint16(binary.LittleEndian.Uint16(buffer))
    case 57:
		data.SaturatedHemoglobinPercent = new(uint16)
		*data.SaturatedHemoglobinPercent = uint16(binary.LittleEndian.Uint16(buffer))
    case 58:
		data.SaturatedHemoglobinPercentMin = new(uint16)
		*data.SaturatedHemoglobinPercentMin = uint16(binary.LittleEndian.Uint16(buffer))
    case 59:
		data.SaturatedHemoglobinPercentMax = new(uint16)
		*data.SaturatedHemoglobinPercentMax = uint16(binary.LittleEndian.Uint16(buffer))
    case 62:
		data.DeviceIndex = new(device_index)
        *data.DeviceIndex = device_index(buffer[0])
    case 67:
		data.LeftPco = new(int8)
        *data.LeftPco = int8(buffer[0])
    case 68:
		data.RightPco = new(int8)
        *data.RightPco = int8(buffer[0])
    case 69:
		data.LeftPowerPhase = new(uint8)
        *data.LeftPowerPhase = uint8(buffer[0])
    case 70:
		data.LeftPowerPhasePeak = new(uint8)
        *data.LeftPowerPhasePeak = uint8(buffer[0])
    case 71:
		data.RightPowerPhase = new(uint8)
        *data.RightPowerPhase = uint8(buffer[0])
    case 72:
		data.RightPowerPhasePeak = new(uint8)
        *data.RightPowerPhasePeak = uint8(buffer[0])
    case 73:
		data.EnhancedSpeed = new(uint32)
		*data.EnhancedSpeed = uint32(binary.LittleEndian.Uint32(buffer))
    case 78:
		data.EnhancedAltitude = new(uint32)
		*data.EnhancedAltitude = uint32(binary.LittleEndian.Uint32(buffer))
    case 81:
		data.BatterySoc = new(uint8)
        *data.BatterySoc = uint8(buffer[0])
    case 82:
		data.MotorPower = new(uint16)
		*data.MotorPower = uint16(binary.LittleEndian.Uint16(buffer))
	default:
		return fmt.Errorf("unknown field number (%d) for message %s", fieldDefinitionNumber, "record")
    }
	return nil
}

func (data *EventDataMessage) Read(fieldDefinitionNumber uint, buffer []byte, byteOrder binary.ByteOrder) error {
    switch fieldDefinitionNumber {
    case 253:
		data.Timestamp = new(date_time)
		*data.Timestamp = date_time(binary.LittleEndian.Uint32(buffer))
    case 0:
		data.Event = new(event)
        *data.Event = event(buffer[0])
    case 1:
		data.EventType = new(event_type)
        *data.EventType = event_type(buffer[0])
    case 2:
		data.Data16 = new(uint16)
		*data.Data16 = uint16(binary.LittleEndian.Uint16(buffer))
    case 3:
		data.Data = new(uint32)
		*data.Data = uint32(binary.LittleEndian.Uint32(buffer))
    case 4:
		data.EventGroup = new(uint8)
        *data.EventGroup = uint8(buffer[0])
    case 7:
		data.Score = new(uint16)
		*data.Score = uint16(binary.LittleEndian.Uint16(buffer))
    case 8:
		data.OpponentScore = new(uint16)
		*data.OpponentScore = uint16(binary.LittleEndian.Uint16(buffer))
    case 9:
		data.FrontGearNum = new(uint8)
        *data.FrontGearNum = uint8(buffer[0])
    case 10:
		data.FrontGear = new(uint8)
        *data.FrontGear = uint8(buffer[0])
    case 11:
		data.RearGearNum = new(uint8)
        *data.RearGearNum = uint8(buffer[0])
    case 12:
		data.RearGear = new(uint8)
        *data.RearGear = uint8(buffer[0])
    case 13:
		data.DeviceIndex = new(device_index)
        *data.DeviceIndex = device_index(buffer[0])
	default:
		return fmt.Errorf("unknown field number (%d) for message %s", fieldDefinitionNumber, "event")
    }
	return nil
}

func (data *DeviceInfoDataMessage) Read(fieldDefinitionNumber uint, buffer []byte, byteOrder binary.ByteOrder) error {
    switch fieldDefinitionNumber {
    case 253:
		data.Timestamp = new(date_time)
		*data.Timestamp = date_time(binary.LittleEndian.Uint32(buffer))
    case 0:
		data.DeviceIndex = new(device_index)
        *data.DeviceIndex = device_index(buffer[0])
    case 1:
		data.DeviceType = new(uint8)
        *data.DeviceType = uint8(buffer[0])
    case 2:
		data.Manufacturer = new(manufacturer)
		*data.Manufacturer = manufacturer(binary.LittleEndian.Uint16(buffer))
    case 3:
		data.SerialNumber = new(uint32)
		*data.SerialNumber = uint32(binary.LittleEndian.Uint32(buffer))
    case 4:
		data.Product = new(uint16)
		*data.Product = uint16(binary.LittleEndian.Uint16(buffer))
    case 5:
		data.SoftwareVersion = new(uint16)
		*data.SoftwareVersion = uint16(binary.LittleEndian.Uint16(buffer))
    case 6:
		data.HardwareVersion = new(uint8)
        *data.HardwareVersion = uint8(buffer[0])
    case 7:
		data.CumOperatingTime = new(uint32)
		*data.CumOperatingTime = uint32(binary.LittleEndian.Uint32(buffer))
    case 10:
		data.BatteryVoltage = new(uint16)
		*data.BatteryVoltage = uint16(binary.LittleEndian.Uint16(buffer))
    case 11:
		data.BatteryStatus = new(battery_status)
        *data.BatteryStatus = battery_status(buffer[0])
    case 18:
		data.SensorPosition = new(body_location)
        *data.SensorPosition = body_location(buffer[0])
    case 19:
		data.Descriptor = new(string)
        *data.Descriptor = string(buffer[0])
    case 20:
		data.AntTransmissionType = new(uint8)
        *data.AntTransmissionType = uint8(buffer[0])
    case 21:
		data.AntDeviceNumber = new(uint16)
		*data.AntDeviceNumber = uint16(binary.LittleEndian.Uint16(buffer))
    case 22:
		data.AntNetwork = new(ant_network)
        *data.AntNetwork = ant_network(buffer[0])
    case 25:
		data.SourceType = new(source_type)
        *data.SourceType = source_type(buffer[0])
	default:
		return fmt.Errorf("unknown field number (%d) for message %s", fieldDefinitionNumber, "device_info")
    }
	return nil
}

func (data *TrainingFileDataMessage) Read(fieldDefinitionNumber uint, buffer []byte, byteOrder binary.ByteOrder) error {
    switch fieldDefinitionNumber {
    case 253:
		data.Timestamp = new(date_time)
		*data.Timestamp = date_time(binary.LittleEndian.Uint32(buffer))
    case 0:
		data.Type = new(file)
        *data.Type = file(buffer[0])
    case 1:
		data.Manufacturer = new(manufacturer)
		*data.Manufacturer = manufacturer(binary.LittleEndian.Uint16(buffer))
    case 2:
		data.Product = new(uint16)
		*data.Product = uint16(binary.LittleEndian.Uint16(buffer))
    case 3:
		data.SerialNumber = new(uint32)
		*data.SerialNumber = uint32(binary.LittleEndian.Uint32(buffer))
    case 4:
		data.TimeCreated = new(date_time)
		*data.TimeCreated = date_time(binary.LittleEndian.Uint32(buffer))
	default:
		return fmt.Errorf("unknown field number (%d) for message %s", fieldDefinitionNumber, "training_file")
    }
	return nil
}

func (data *HrvDataMessage) Read(fieldDefinitionNumber uint, buffer []byte, byteOrder binary.ByteOrder) error {
    switch fieldDefinitionNumber {
    case 0:
		data.Time = new(uint16)
		*data.Time = uint16(binary.LittleEndian.Uint16(buffer))
	default:
		return fmt.Errorf("unknown field number (%d) for message %s", fieldDefinitionNumber, "hrv")
    }
	return nil
}

func (data *CourseDataMessage) Read(fieldDefinitionNumber uint, buffer []byte, byteOrder binary.ByteOrder) error {
    switch fieldDefinitionNumber {
    case 4:
		data.Sport = new(sport)
        *data.Sport = sport(buffer[0])
    case 5:
		data.Name = new(string)
        *data.Name = string(buffer[0])
    case 6:
		data.Capabilities = new(course_capabilities)
		*data.Capabilities = course_capabilities(binary.LittleEndian.Uint32(buffer))
	default:
		return fmt.Errorf("unknown field number (%d) for message %s", fieldDefinitionNumber, "course")
    }
	return nil
}

func (data *CoursePointDataMessage) Read(fieldDefinitionNumber uint, buffer []byte, byteOrder binary.ByteOrder) error {
    switch fieldDefinitionNumber {
    case 254:
		data.MessageIndex = new(message_index)
		*data.MessageIndex = message_index(binary.LittleEndian.Uint16(buffer))
    case 1:
		data.Timestamp = new(date_time)
		*data.Timestamp = date_time(binary.LittleEndian.Uint32(buffer))
    case 2:
		data.PositionLat = new(int32)
		*data.PositionLat = int32(binary.LittleEndian.Uint32(buffer))
    case 3:
		data.PositionLong = new(int32)
		*data.PositionLong = int32(binary.LittleEndian.Uint32(buffer))
    case 4:
		data.Distance = new(uint32)
		*data.Distance = uint32(binary.LittleEndian.Uint32(buffer))
    case 5:
		data.Type = new(course_point)
        *data.Type = course_point(buffer[0])
    case 6:
		data.Name = new(string)
        *data.Name = string(buffer[0])
    case 8:
		data.Favorite = new(bool)
	default:
		return fmt.Errorf("unknown field number (%d) for message %s", fieldDefinitionNumber, "course_point")
    }
	return nil
}

func (data *SegmentIdDataMessage) Read(fieldDefinitionNumber uint, buffer []byte, byteOrder binary.ByteOrder) error {
    switch fieldDefinitionNumber {
    case 0:
		data.Name = new(string)
        *data.Name = string(buffer[0])
    case 1:
		data.Uuid = new(string)
        *data.Uuid = string(buffer[0])
    case 2:
		data.Sport = new(sport)
        *data.Sport = sport(buffer[0])
    case 3:
		data.Enabled = new(bool)
    case 4:
		data.UserProfilePrimaryKey = new(uint32)
		*data.UserProfilePrimaryKey = uint32(binary.LittleEndian.Uint32(buffer))
    case 5:
		data.DeviceId = new(uint32)
		*data.DeviceId = uint32(binary.LittleEndian.Uint32(buffer))
    case 6:
		data.DefaultRaceLeader = new(uint8)
        *data.DefaultRaceLeader = uint8(buffer[0])
    case 7:
		data.DeleteStatus = new(segment_delete_status)
        *data.DeleteStatus = segment_delete_status(buffer[0])
    case 8:
		data.SelectionType = new(segment_selection_type)
        *data.SelectionType = segment_selection_type(buffer[0])
	default:
		return fmt.Errorf("unknown field number (%d) for message %s", fieldDefinitionNumber, "segment_id")
    }
	return nil
}

func (data *SegmentLeaderboardEntryDataMessage) Read(fieldDefinitionNumber uint, buffer []byte, byteOrder binary.ByteOrder) error {
    switch fieldDefinitionNumber {
    case 254:
		data.MessageIndex = new(message_index)
		*data.MessageIndex = message_index(binary.LittleEndian.Uint16(buffer))
    case 0:
		data.Name = new(string)
        *data.Name = string(buffer[0])
    case 1:
		data.Type = new(segment_leaderboard_type)
        *data.Type = segment_leaderboard_type(buffer[0])
    case 2:
		data.GroupPrimaryKey = new(uint32)
		*data.GroupPrimaryKey = uint32(binary.LittleEndian.Uint32(buffer))
    case 3:
		data.ActivityId = new(uint32)
		*data.ActivityId = uint32(binary.LittleEndian.Uint32(buffer))
    case 4:
		data.SegmentTime = new(uint32)
		*data.SegmentTime = uint32(binary.LittleEndian.Uint32(buffer))
    case 5:
		data.ActivityIdString = new(string)
        *data.ActivityIdString = string(buffer[0])
	default:
		return fmt.Errorf("unknown field number (%d) for message %s", fieldDefinitionNumber, "segment_leaderboard_entry")
    }
	return nil
}

func (data *SegmentPointDataMessage) Read(fieldDefinitionNumber uint, buffer []byte, byteOrder binary.ByteOrder) error {
    switch fieldDefinitionNumber {
    case 254:
		data.MessageIndex = new(message_index)
		*data.MessageIndex = message_index(binary.LittleEndian.Uint16(buffer))
    case 1:
		data.PositionLat = new(int32)
		*data.PositionLat = int32(binary.LittleEndian.Uint32(buffer))
    case 2:
		data.PositionLong = new(int32)
		*data.PositionLong = int32(binary.LittleEndian.Uint32(buffer))
    case 3:
		data.Distance = new(uint32)
		*data.Distance = uint32(binary.LittleEndian.Uint32(buffer))
    case 4:
		data.Altitude = new(uint16)
		*data.Altitude = uint16(binary.LittleEndian.Uint16(buffer))
    case 5:
		data.LeaderTime = new(uint32)
		*data.LeaderTime = uint32(binary.LittleEndian.Uint32(buffer))
	default:
		return fmt.Errorf("unknown field number (%d) for message %s", fieldDefinitionNumber, "segment_point")
    }
	return nil
}

func (data *SegmentLapDataMessage) Read(fieldDefinitionNumber uint, buffer []byte, byteOrder binary.ByteOrder) error {
    switch fieldDefinitionNumber {
    case 254:
		data.MessageIndex = new(message_index)
		*data.MessageIndex = message_index(binary.LittleEndian.Uint16(buffer))
    case 253:
		data.Timestamp = new(date_time)
		*data.Timestamp = date_time(binary.LittleEndian.Uint32(buffer))
    case 0:
		data.Event = new(event)
        *data.Event = event(buffer[0])
    case 1:
		data.EventType = new(event_type)
        *data.EventType = event_type(buffer[0])
    case 2:
		data.StartTime = new(date_time)
		*data.StartTime = date_time(binary.LittleEndian.Uint32(buffer))
    case 3:
		data.StartPositionLat = new(int32)
		*data.StartPositionLat = int32(binary.LittleEndian.Uint32(buffer))
    case 4:
		data.StartPositionLong = new(int32)
		*data.StartPositionLong = int32(binary.LittleEndian.Uint32(buffer))
    case 5:
		data.EndPositionLat = new(int32)
		*data.EndPositionLat = int32(binary.LittleEndian.Uint32(buffer))
    case 6:
		data.EndPositionLong = new(int32)
		*data.EndPositionLong = int32(binary.LittleEndian.Uint32(buffer))
    case 7:
		data.TotalElapsedTime = new(uint32)
		*data.TotalElapsedTime = uint32(binary.LittleEndian.Uint32(buffer))
    case 8:
		data.TotalTimerTime = new(uint32)
		*data.TotalTimerTime = uint32(binary.LittleEndian.Uint32(buffer))
    case 9:
		data.TotalDistance = new(uint32)
		*data.TotalDistance = uint32(binary.LittleEndian.Uint32(buffer))
    case 10:
		data.TotalCycles = new(uint32)
		*data.TotalCycles = uint32(binary.LittleEndian.Uint32(buffer))
    case 11:
		data.TotalCalories = new(uint16)
		*data.TotalCalories = uint16(binary.LittleEndian.Uint16(buffer))
    case 12:
		data.TotalFatCalories = new(uint16)
		*data.TotalFatCalories = uint16(binary.LittleEndian.Uint16(buffer))
    case 13:
		data.AvgSpeed = new(uint16)
		*data.AvgSpeed = uint16(binary.LittleEndian.Uint16(buffer))
    case 14:
		data.MaxSpeed = new(uint16)
		*data.MaxSpeed = uint16(binary.LittleEndian.Uint16(buffer))
    case 15:
		data.AvgHeartRate = new(uint8)
        *data.AvgHeartRate = uint8(buffer[0])
    case 16:
		data.MaxHeartRate = new(uint8)
        *data.MaxHeartRate = uint8(buffer[0])
    case 17:
		data.AvgCadence = new(uint8)
        *data.AvgCadence = uint8(buffer[0])
    case 18:
		data.MaxCadence = new(uint8)
        *data.MaxCadence = uint8(buffer[0])
    case 19:
		data.AvgPower = new(uint16)
		*data.AvgPower = uint16(binary.LittleEndian.Uint16(buffer))
    case 20:
		data.MaxPower = new(uint16)
		*data.MaxPower = uint16(binary.LittleEndian.Uint16(buffer))
    case 21:
		data.TotalAscent = new(uint16)
		*data.TotalAscent = uint16(binary.LittleEndian.Uint16(buffer))
    case 22:
		data.TotalDescent = new(uint16)
		*data.TotalDescent = uint16(binary.LittleEndian.Uint16(buffer))
    case 23:
		data.Sport = new(sport)
        *data.Sport = sport(buffer[0])
    case 24:
		data.EventGroup = new(uint8)
        *data.EventGroup = uint8(buffer[0])
    case 25:
		data.NecLat = new(int32)
		*data.NecLat = int32(binary.LittleEndian.Uint32(buffer))
    case 26:
		data.NecLong = new(int32)
		*data.NecLong = int32(binary.LittleEndian.Uint32(buffer))
    case 27:
		data.SwcLat = new(int32)
		*data.SwcLat = int32(binary.LittleEndian.Uint32(buffer))
    case 28:
		data.SwcLong = new(int32)
		*data.SwcLong = int32(binary.LittleEndian.Uint32(buffer))
    case 29:
		data.Name = new(string)
        *data.Name = string(buffer[0])
    case 30:
		data.NormalizedPower = new(uint16)
		*data.NormalizedPower = uint16(binary.LittleEndian.Uint16(buffer))
    case 31:
		data.LeftRightBalance = new(left_right_balance_100)
		*data.LeftRightBalance = left_right_balance_100(binary.LittleEndian.Uint16(buffer))
    case 32:
		data.SubSport = new(sub_sport)
        *data.SubSport = sub_sport(buffer[0])
    case 33:
		data.TotalWork = new(uint32)
		*data.TotalWork = uint32(binary.LittleEndian.Uint32(buffer))
    case 34:
		data.AvgAltitude = new(uint16)
		*data.AvgAltitude = uint16(binary.LittleEndian.Uint16(buffer))
    case 35:
		data.MaxAltitude = new(uint16)
		*data.MaxAltitude = uint16(binary.LittleEndian.Uint16(buffer))
    case 36:
		data.GpsAccuracy = new(uint8)
        *data.GpsAccuracy = uint8(buffer[0])
    case 37:
		data.AvgGrade = new(int16)
		*data.AvgGrade = int16(binary.LittleEndian.Uint16(buffer))
    case 38:
		data.AvgPosGrade = new(int16)
		*data.AvgPosGrade = int16(binary.LittleEndian.Uint16(buffer))
    case 39:
		data.AvgNegGrade = new(int16)
		*data.AvgNegGrade = int16(binary.LittleEndian.Uint16(buffer))
    case 40:
		data.MaxPosGrade = new(int16)
		*data.MaxPosGrade = int16(binary.LittleEndian.Uint16(buffer))
    case 41:
		data.MaxNegGrade = new(int16)
		*data.MaxNegGrade = int16(binary.LittleEndian.Uint16(buffer))
    case 42:
		data.AvgTemperature = new(int8)
        *data.AvgTemperature = int8(buffer[0])
    case 43:
		data.MaxTemperature = new(int8)
        *data.MaxTemperature = int8(buffer[0])
    case 44:
		data.TotalMovingTime = new(uint32)
		*data.TotalMovingTime = uint32(binary.LittleEndian.Uint32(buffer))
    case 45:
		data.AvgPosVerticalSpeed = new(int16)
		*data.AvgPosVerticalSpeed = int16(binary.LittleEndian.Uint16(buffer))
    case 46:
		data.AvgNegVerticalSpeed = new(int16)
		*data.AvgNegVerticalSpeed = int16(binary.LittleEndian.Uint16(buffer))
    case 47:
		data.MaxPosVerticalSpeed = new(int16)
		*data.MaxPosVerticalSpeed = int16(binary.LittleEndian.Uint16(buffer))
    case 48:
		data.MaxNegVerticalSpeed = new(int16)
		*data.MaxNegVerticalSpeed = int16(binary.LittleEndian.Uint16(buffer))
    case 49:
		data.TimeInHrZone = new(uint32)
		*data.TimeInHrZone = uint32(binary.LittleEndian.Uint32(buffer))
    case 50:
		data.TimeInSpeedZone = new(uint32)
		*data.TimeInSpeedZone = uint32(binary.LittleEndian.Uint32(buffer))
    case 51:
		data.TimeInCadenceZone = new(uint32)
		*data.TimeInCadenceZone = uint32(binary.LittleEndian.Uint32(buffer))
    case 52:
		data.TimeInPowerZone = new(uint32)
		*data.TimeInPowerZone = uint32(binary.LittleEndian.Uint32(buffer))
    case 53:
		data.RepetitionNum = new(uint16)
		*data.RepetitionNum = uint16(binary.LittleEndian.Uint16(buffer))
    case 54:
		data.MinAltitude = new(uint16)
		*data.MinAltitude = uint16(binary.LittleEndian.Uint16(buffer))
    case 55:
		data.MinHeartRate = new(uint8)
        *data.MinHeartRate = uint8(buffer[0])
    case 56:
		data.ActiveTime = new(uint32)
		*data.ActiveTime = uint32(binary.LittleEndian.Uint32(buffer))
    case 57:
		data.WktStepIndex = new(message_index)
		*data.WktStepIndex = message_index(binary.LittleEndian.Uint16(buffer))
    case 58:
		data.SportEvent = new(sport_event)
        *data.SportEvent = sport_event(buffer[0])
    case 59:
		data.AvgLeftTorqueEffectiveness = new(uint8)
        *data.AvgLeftTorqueEffectiveness = uint8(buffer[0])
    case 60:
		data.AvgRightTorqueEffectiveness = new(uint8)
        *data.AvgRightTorqueEffectiveness = uint8(buffer[0])
    case 61:
		data.AvgLeftPedalSmoothness = new(uint8)
        *data.AvgLeftPedalSmoothness = uint8(buffer[0])
    case 62:
		data.AvgRightPedalSmoothness = new(uint8)
        *data.AvgRightPedalSmoothness = uint8(buffer[0])
    case 63:
		data.AvgCombinedPedalSmoothness = new(uint8)
        *data.AvgCombinedPedalSmoothness = uint8(buffer[0])
    case 64:
		data.Status = new(segment_lap_status)
        *data.Status = segment_lap_status(buffer[0])
    case 65:
		data.Uuid = new(string)
        *data.Uuid = string(buffer[0])
    case 66:
		data.AvgFractionalCadence = new(uint8)
        *data.AvgFractionalCadence = uint8(buffer[0])
    case 67:
		data.MaxFractionalCadence = new(uint8)
        *data.MaxFractionalCadence = uint8(buffer[0])
    case 68:
		data.TotalFractionalCycles = new(uint8)
        *data.TotalFractionalCycles = uint8(buffer[0])
    case 69:
		data.FrontGearShiftCount = new(uint16)
		*data.FrontGearShiftCount = uint16(binary.LittleEndian.Uint16(buffer))
    case 70:
		data.RearGearShiftCount = new(uint16)
		*data.RearGearShiftCount = uint16(binary.LittleEndian.Uint16(buffer))
    case 71:
		data.TimeStanding = new(uint32)
		*data.TimeStanding = uint32(binary.LittleEndian.Uint32(buffer))
    case 72:
		data.StandCount = new(uint16)
		*data.StandCount = uint16(binary.LittleEndian.Uint16(buffer))
    case 73:
		data.AvgLeftPco = new(int8)
        *data.AvgLeftPco = int8(buffer[0])
    case 74:
		data.AvgRightPco = new(int8)
        *data.AvgRightPco = int8(buffer[0])
    case 75:
		data.AvgLeftPowerPhase = new(uint8)
        *data.AvgLeftPowerPhase = uint8(buffer[0])
    case 76:
		data.AvgLeftPowerPhasePeak = new(uint8)
        *data.AvgLeftPowerPhasePeak = uint8(buffer[0])
    case 77:
		data.AvgRightPowerPhase = new(uint8)
        *data.AvgRightPowerPhase = uint8(buffer[0])
    case 78:
		data.AvgRightPowerPhasePeak = new(uint8)
        *data.AvgRightPowerPhasePeak = uint8(buffer[0])
    case 79:
		data.AvgPowerPosition = new(uint16)
		*data.AvgPowerPosition = uint16(binary.LittleEndian.Uint16(buffer))
    case 80:
		data.MaxPowerPosition = new(uint16)
		*data.MaxPowerPosition = uint16(binary.LittleEndian.Uint16(buffer))
    case 81:
		data.AvgCadencePosition = new(uint8)
        *data.AvgCadencePosition = uint8(buffer[0])
    case 82:
		data.MaxCadencePosition = new(uint8)
        *data.MaxCadencePosition = uint8(buffer[0])
	default:
		return fmt.Errorf("unknown field number (%d) for message %s", fieldDefinitionNumber, "segment_lap")
    }
	return nil
}

func (data *SegmentFileDataMessage) Read(fieldDefinitionNumber uint, buffer []byte, byteOrder binary.ByteOrder) error {
    switch fieldDefinitionNumber {
    case 254:
		data.MessageIndex = new(message_index)
		*data.MessageIndex = message_index(binary.LittleEndian.Uint16(buffer))
    case 1:
		data.FileUuid = new(string)
        *data.FileUuid = string(buffer[0])
    case 3:
		data.Enabled = new(bool)
    case 4:
		data.UserProfilePrimaryKey = new(uint32)
		*data.UserProfilePrimaryKey = uint32(binary.LittleEndian.Uint32(buffer))
    case 7:
		data.LeaderType = new(segment_leaderboard_type)
        *data.LeaderType = segment_leaderboard_type(buffer[0])
    case 8:
		data.LeaderGroupPrimaryKey = new(uint32)
		*data.LeaderGroupPrimaryKey = uint32(binary.LittleEndian.Uint32(buffer))
    case 9:
		data.LeaderActivityId = new(uint32)
		*data.LeaderActivityId = uint32(binary.LittleEndian.Uint32(buffer))
    case 10:
		data.LeaderActivityIdString = new(string)
        *data.LeaderActivityIdString = string(buffer[0])
	default:
		return fmt.Errorf("unknown field number (%d) for message %s", fieldDefinitionNumber, "segment_file")
    }
	return nil
}

func (data *WorkoutDataMessage) Read(fieldDefinitionNumber uint, buffer []byte, byteOrder binary.ByteOrder) error {
    switch fieldDefinitionNumber {
    case 4:
		data.Sport = new(sport)
        *data.Sport = sport(buffer[0])
    case 5:
		data.Capabilities = new(workout_capabilities)
		*data.Capabilities = workout_capabilities(binary.LittleEndian.Uint32(buffer))
    case 6:
		data.NumValidSteps = new(uint16)
		*data.NumValidSteps = uint16(binary.LittleEndian.Uint16(buffer))
    case 8:
		data.WktName = new(string)
        *data.WktName = string(buffer[0])
	default:
		return fmt.Errorf("unknown field number (%d) for message %s", fieldDefinitionNumber, "workout")
    }
	return nil
}

func (data *WorkoutStepDataMessage) Read(fieldDefinitionNumber uint, buffer []byte, byteOrder binary.ByteOrder) error {
    switch fieldDefinitionNumber {
    case 254:
		data.MessageIndex = new(message_index)
		*data.MessageIndex = message_index(binary.LittleEndian.Uint16(buffer))
    case 0:
		data.WktStepName = new(string)
        *data.WktStepName = string(buffer[0])
    case 1:
		data.DurationType = new(wkt_step_duration)
        *data.DurationType = wkt_step_duration(buffer[0])
    case 2:
		data.DurationValue = new(uint32)
		*data.DurationValue = uint32(binary.LittleEndian.Uint32(buffer))
    case 3:
		data.TargetType = new(wkt_step_target)
        *data.TargetType = wkt_step_target(buffer[0])
    case 4:
		data.TargetValue = new(uint32)
		*data.TargetValue = uint32(binary.LittleEndian.Uint32(buffer))
    case 5:
		data.CustomTargetValueLow = new(uint32)
		*data.CustomTargetValueLow = uint32(binary.LittleEndian.Uint32(buffer))
    case 6:
		data.CustomTargetValueHigh = new(uint32)
		*data.CustomTargetValueHigh = uint32(binary.LittleEndian.Uint32(buffer))
    case 7:
		data.Intensity = new(intensity)
        *data.Intensity = intensity(buffer[0])
	default:
		return fmt.Errorf("unknown field number (%d) for message %s", fieldDefinitionNumber, "workout_step")
    }
	return nil
}

func (data *ScheduleDataMessage) Read(fieldDefinitionNumber uint, buffer []byte, byteOrder binary.ByteOrder) error {
    switch fieldDefinitionNumber {
    case 0:
		data.Manufacturer = new(manufacturer)
		*data.Manufacturer = manufacturer(binary.LittleEndian.Uint16(buffer))
    case 1:
		data.Product = new(uint16)
		*data.Product = uint16(binary.LittleEndian.Uint16(buffer))
    case 2:
		data.SerialNumber = new(uint32)
		*data.SerialNumber = uint32(binary.LittleEndian.Uint32(buffer))
    case 3:
		data.TimeCreated = new(date_time)
		*data.TimeCreated = date_time(binary.LittleEndian.Uint32(buffer))
    case 4:
		data.Completed = new(bool)
    case 5:
		data.Type = new(schedule)
        *data.Type = schedule(buffer[0])
    case 6:
		data.ScheduledTime = new(local_date_time)
		*data.ScheduledTime = local_date_time(binary.LittleEndian.Uint32(buffer))
	default:
		return fmt.Errorf("unknown field number (%d) for message %s", fieldDefinitionNumber, "schedule")
    }
	return nil
}

func (data *TotalsDataMessage) Read(fieldDefinitionNumber uint, buffer []byte, byteOrder binary.ByteOrder) error {
    switch fieldDefinitionNumber {
    case 254:
		data.MessageIndex = new(message_index)
		*data.MessageIndex = message_index(binary.LittleEndian.Uint16(buffer))
    case 253:
		data.Timestamp = new(date_time)
		*data.Timestamp = date_time(binary.LittleEndian.Uint32(buffer))
    case 0:
		data.TimerTime = new(uint32)
		*data.TimerTime = uint32(binary.LittleEndian.Uint32(buffer))
    case 1:
		data.Distance = new(uint32)
		*data.Distance = uint32(binary.LittleEndian.Uint32(buffer))
    case 2:
		data.Calories = new(uint32)
		*data.Calories = uint32(binary.LittleEndian.Uint32(buffer))
    case 3:
		data.Sport = new(sport)
        *data.Sport = sport(buffer[0])
    case 4:
		data.ElapsedTime = new(uint32)
		*data.ElapsedTime = uint32(binary.LittleEndian.Uint32(buffer))
    case 5:
		data.Sessions = new(uint16)
		*data.Sessions = uint16(binary.LittleEndian.Uint16(buffer))
    case 6:
		data.ActiveTime = new(uint32)
		*data.ActiveTime = uint32(binary.LittleEndian.Uint32(buffer))
    case 9:
		data.SportIndex = new(uint8)
        *data.SportIndex = uint8(buffer[0])
	default:
		return fmt.Errorf("unknown field number (%d) for message %s", fieldDefinitionNumber, "totals")
    }
	return nil
}

func (data *WeightScaleDataMessage) Read(fieldDefinitionNumber uint, buffer []byte, byteOrder binary.ByteOrder) error {
    switch fieldDefinitionNumber {
    case 253:
		data.Timestamp = new(date_time)
		*data.Timestamp = date_time(binary.LittleEndian.Uint32(buffer))
    case 0:
		data.Weight = new(weight)
		*data.Weight = weight(binary.LittleEndian.Uint16(buffer))
    case 1:
		data.PercentFat = new(uint16)
		*data.PercentFat = uint16(binary.LittleEndian.Uint16(buffer))
    case 2:
		data.PercentHydration = new(uint16)
		*data.PercentHydration = uint16(binary.LittleEndian.Uint16(buffer))
    case 3:
		data.VisceralFatMass = new(uint16)
		*data.VisceralFatMass = uint16(binary.LittleEndian.Uint16(buffer))
    case 4:
		data.BoneMass = new(uint16)
		*data.BoneMass = uint16(binary.LittleEndian.Uint16(buffer))
    case 5:
		data.MuscleMass = new(uint16)
		*data.MuscleMass = uint16(binary.LittleEndian.Uint16(buffer))
    case 7:
		data.BasalMet = new(uint16)
		*data.BasalMet = uint16(binary.LittleEndian.Uint16(buffer))
    case 8:
		data.PhysiqueRating = new(uint8)
        *data.PhysiqueRating = uint8(buffer[0])
    case 9:
		data.ActiveMet = new(uint16)
		*data.ActiveMet = uint16(binary.LittleEndian.Uint16(buffer))
    case 10:
		data.MetabolicAge = new(uint8)
        *data.MetabolicAge = uint8(buffer[0])
    case 11:
		data.VisceralFatRating = new(uint8)
        *data.VisceralFatRating = uint8(buffer[0])
    case 12:
		data.UserProfileIndex = new(message_index)
		*data.UserProfileIndex = message_index(binary.LittleEndian.Uint16(buffer))
	default:
		return fmt.Errorf("unknown field number (%d) for message %s", fieldDefinitionNumber, "weight_scale")
    }
	return nil
}

func (data *BloodPressureDataMessage) Read(fieldDefinitionNumber uint, buffer []byte, byteOrder binary.ByteOrder) error {
    switch fieldDefinitionNumber {
    case 253:
		data.Timestamp = new(date_time)
		*data.Timestamp = date_time(binary.LittleEndian.Uint32(buffer))
    case 0:
		data.SystolicPressure = new(uint16)
		*data.SystolicPressure = uint16(binary.LittleEndian.Uint16(buffer))
    case 1:
		data.DiastolicPressure = new(uint16)
		*data.DiastolicPressure = uint16(binary.LittleEndian.Uint16(buffer))
    case 2:
		data.MeanArterialPressure = new(uint16)
		*data.MeanArterialPressure = uint16(binary.LittleEndian.Uint16(buffer))
    case 3:
		data.Map3SampleMean = new(uint16)
		*data.Map3SampleMean = uint16(binary.LittleEndian.Uint16(buffer))
    case 4:
		data.MapMorningValues = new(uint16)
		*data.MapMorningValues = uint16(binary.LittleEndian.Uint16(buffer))
    case 5:
		data.MapEveningValues = new(uint16)
		*data.MapEveningValues = uint16(binary.LittleEndian.Uint16(buffer))
    case 6:
		data.HeartRate = new(uint8)
        *data.HeartRate = uint8(buffer[0])
    case 7:
		data.HeartRateType = new(hr_type)
        *data.HeartRateType = hr_type(buffer[0])
    case 8:
		data.Status = new(bp_status)
        *data.Status = bp_status(buffer[0])
    case 9:
		data.UserProfileIndex = new(message_index)
		*data.UserProfileIndex = message_index(binary.LittleEndian.Uint16(buffer))
	default:
		return fmt.Errorf("unknown field number (%d) for message %s", fieldDefinitionNumber, "blood_pressure")
    }
	return nil
}

func (data *MonitoringInfoDataMessage) Read(fieldDefinitionNumber uint, buffer []byte, byteOrder binary.ByteOrder) error {
    switch fieldDefinitionNumber {
    case 253:
		data.Timestamp = new(date_time)
		*data.Timestamp = date_time(binary.LittleEndian.Uint32(buffer))
    case 0:
		data.LocalTimestamp = new(local_date_time)
		*data.LocalTimestamp = local_date_time(binary.LittleEndian.Uint32(buffer))
    case 1:
		data.ActivityType = new(activity_type)
        *data.ActivityType = activity_type(buffer[0])
    case 3:
		data.CyclesToDistance = new(uint16)
		*data.CyclesToDistance = uint16(binary.LittleEndian.Uint16(buffer))
    case 4:
		data.CyclesToCalories = new(uint16)
		*data.CyclesToCalories = uint16(binary.LittleEndian.Uint16(buffer))
    case 5:
		data.RestingMetabolicRate = new(uint16)
		*data.RestingMetabolicRate = uint16(binary.LittleEndian.Uint16(buffer))
	default:
		return fmt.Errorf("unknown field number (%d) for message %s", fieldDefinitionNumber, "monitoring_info")
    }
	return nil
}

func (data *MonitoringDataMessage) Read(fieldDefinitionNumber uint, buffer []byte, byteOrder binary.ByteOrder) error {
    switch fieldDefinitionNumber {
    case 253:
		data.Timestamp = new(date_time)
		*data.Timestamp = date_time(binary.LittleEndian.Uint32(buffer))
    case 0:
		data.DeviceIndex = new(device_index)
        *data.DeviceIndex = device_index(buffer[0])
    case 1:
		data.Calories = new(uint16)
		*data.Calories = uint16(binary.LittleEndian.Uint16(buffer))
    case 2:
		data.Distance = new(uint32)
		*data.Distance = uint32(binary.LittleEndian.Uint32(buffer))
    case 3:
		data.Cycles = new(uint32)
		*data.Cycles = uint32(binary.LittleEndian.Uint32(buffer))
    case 4:
		data.ActiveTime = new(uint32)
		*data.ActiveTime = uint32(binary.LittleEndian.Uint32(buffer))
    case 5:
		data.ActivityType = new(activity_type)
        *data.ActivityType = activity_type(buffer[0])
    case 6:
		data.ActivitySubtype = new(activity_subtype)
        *data.ActivitySubtype = activity_subtype(buffer[0])
    case 7:
		data.ActivityLevel = new(activity_level)
        *data.ActivityLevel = activity_level(buffer[0])
    case 8:
		data.Distance16 = new(uint16)
		*data.Distance16 = uint16(binary.LittleEndian.Uint16(buffer))
    case 9:
		data.Cycles16 = new(uint16)
		*data.Cycles16 = uint16(binary.LittleEndian.Uint16(buffer))
    case 10:
		data.ActiveTime16 = new(uint16)
		*data.ActiveTime16 = uint16(binary.LittleEndian.Uint16(buffer))
    case 11:
		data.LocalTimestamp = new(local_date_time)
		*data.LocalTimestamp = local_date_time(binary.LittleEndian.Uint32(buffer))
    case 12:
		data.Temperature = new(int16)
		*data.Temperature = int16(binary.LittleEndian.Uint16(buffer))
    case 14:
		data.TemperatureMin = new(int16)
		*data.TemperatureMin = int16(binary.LittleEndian.Uint16(buffer))
    case 15:
		data.TemperatureMax = new(int16)
		*data.TemperatureMax = int16(binary.LittleEndian.Uint16(buffer))
    case 16:
		data.ActivityTime = new(uint16)
		*data.ActivityTime = uint16(binary.LittleEndian.Uint16(buffer))
    case 19:
		data.ActiveCalories = new(uint16)
		*data.ActiveCalories = uint16(binary.LittleEndian.Uint16(buffer))
    case 24:
		data.CurrentActivityTypeIntensity = new(byte)
        *data.CurrentActivityTypeIntensity = byte(buffer[0])
    case 25:
		data.TimestampMin8 = new(uint8)
        *data.TimestampMin8 = uint8(buffer[0])
    case 26:
		data.Timestamp16 = new(uint16)
		*data.Timestamp16 = uint16(binary.LittleEndian.Uint16(buffer))
    case 27:
		data.HeartRate = new(uint8)
        *data.HeartRate = uint8(buffer[0])
    case 28:
		data.Intensity = new(uint8)
        *data.Intensity = uint8(buffer[0])
    case 29:
		data.DurationMin = new(uint16)
		*data.DurationMin = uint16(binary.LittleEndian.Uint16(buffer))
    case 30:
		data.Duration = new(uint32)
		*data.Duration = uint32(binary.LittleEndian.Uint32(buffer))
	default:
		return fmt.Errorf("unknown field number (%d) for message %s", fieldDefinitionNumber, "monitoring")
    }
	return nil
}


