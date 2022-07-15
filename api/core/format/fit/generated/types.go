// Code generated; DO NOT EDIT

package generated


type file uint8
const (
    file_device = 1
    file_settings = 2
    file_sport = 3
    file_activity = 4
    file_workout = 5
    file_course = 6
    file_schedules = 7
    file_weight = 9
    file_totals = 10
    file_goals = 11
    file_blood_pressure = 14
    file_monitoring_a = 15
    file_activity_summary = 20
    file_monitoring_daily = 28
    file_monitoring_b = 32
    file_segment = 34
    file_segment_list = 35
    file_mfg_range_min = 247
    file_mfg_range_max = 254
)


type mesg_num uint16

var validIDs = []uint{
	0,
	1,
	2,
	3,
	4,
	5,
	6,
	7,
	8,
	9,
	10,
	12,
	15,
	18,
	19,
	20,
	21,
	23,
	26,
	27,
	28,
	30,
	31,
	32,
	33,
	34,
	35,
	37,
	38,
	39,
	49,
	51,
	53,
	55,
	72,
	78,
	101,
	103,
	105,
	106,
	131,
	142,
	145,
	148,
	149,
	150,
	151,
	65280,
	65534,
}
func IsValidMessageID(messageID uint) bool {
	for _, id := range validIDs {
		if messageID == id {
			return true
		}
	}
	return false
}


type checksum uint8


type file_flags uint8


type mesg_count uint8
const (
    mesg_count_num_per_file = 0
    mesg_count_max_per_file = 1
    mesg_count_max_per_file_type = 2
)


type date_time uint32


type local_date_time uint32


type message_index uint16


type device_index uint8


type gender uint8
const (
    gender_female = 0
    gender_male = 1
)


type language uint8
const (
    language_english = 0
    language_french = 1
    language_italian = 2
    language_german = 3
    language_spanish = 4
    language_croatian = 5
    language_czech = 6
    language_danish = 7
    language_dutch = 8
    language_finnish = 9
    language_greek = 10
    language_hungarian = 11
    language_norwegian = 12
    language_polish = 13
    language_portuguese = 14
    language_slovakian = 15
    language_slovenian = 16
    language_swedish = 17
    language_russian = 18
    language_turkish = 19
    language_latvian = 20
    language_ukrainian = 21
    language_arabic = 22
    language_farsi = 23
    language_bulgarian = 24
    language_romanian = 25
    language_custom = 254
)


type time_zone uint8
const (
    time_zone_almaty = 0
    time_zone_bangkok = 1
    time_zone_bombay = 2
    time_zone_brasilia = 3
    time_zone_cairo = 4
    time_zone_cape_verde_is = 5
    time_zone_darwin = 6
    time_zone_eniwetok = 7
    time_zone_fiji = 8
    time_zone_hong_kong = 9
    time_zone_islamabad = 10
    time_zone_kabul = 11
    time_zone_magadan = 12
    time_zone_mid_atlantic = 13
    time_zone_moscow = 14
    time_zone_muscat = 15
    time_zone_newfoundland = 16
    time_zone_samoa = 17
    time_zone_sydney = 18
    time_zone_tehran = 19
    time_zone_tokyo = 20
    time_zone_us_alaska = 21
    time_zone_us_atlantic = 22
    time_zone_us_central = 23
    time_zone_us_eastern = 24
    time_zone_us_hawaii = 25
    time_zone_us_mountain = 26
    time_zone_us_pacific = 27
    time_zone_other = 28
    time_zone_auckland = 29
    time_zone_kathmandu = 30
    time_zone_europe_western_wet = 31
    time_zone_europe_central_cet = 32
    time_zone_europe_eastern_eet = 33
    time_zone_jakarta = 34
    time_zone_perth = 35
    time_zone_adelaide = 36
    time_zone_brisbane = 37
    time_zone_tasmania = 38
    time_zone_iceland = 39
    time_zone_amsterdam = 40
    time_zone_athens = 41
    time_zone_barcelona = 42
    time_zone_berlin = 43
    time_zone_brussels = 44
    time_zone_budapest = 45
    time_zone_copenhagen = 46
    time_zone_dublin = 47
    time_zone_helsinki = 48
    time_zone_lisbon = 49
    time_zone_london = 50
    time_zone_madrid = 51
    time_zone_munich = 52
    time_zone_oslo = 53
    time_zone_paris = 54
    time_zone_prague = 55
    time_zone_reykjavik = 56
    time_zone_rome = 57
    time_zone_stockholm = 58
    time_zone_vienna = 59
    time_zone_warsaw = 60
    time_zone_zurich = 61
    time_zone_quebec = 62
    time_zone_ontario = 63
    time_zone_manitoba = 64
    time_zone_saskatchewan = 65
    time_zone_alberta = 66
    time_zone_british_columbia = 67
    time_zone_boise = 68
    time_zone_boston = 69
    time_zone_chicago = 70
    time_zone_dallas = 71
    time_zone_denver = 72
    time_zone_kansas_city = 73
    time_zone_las_vegas = 74
    time_zone_los_angeles = 75
    time_zone_miami = 76
    time_zone_minneapolis = 77
    time_zone_new_york = 78
    time_zone_new_orleans = 79
    time_zone_phoenix = 80
    time_zone_santa_fe = 81
    time_zone_seattle = 82
    time_zone_washington_dc = 83
    time_zone_us_arizona = 84
    time_zone_chita = 85
    time_zone_ekaterinburg = 86
    time_zone_irkutsk = 87
    time_zone_kaliningrad = 88
    time_zone_krasnoyarsk = 89
    time_zone_novosibirsk = 90
    time_zone_petropavlovsk_kamchatskiy = 91
    time_zone_samara = 92
    time_zone_vladivostok = 93
    time_zone_mexico_central = 94
    time_zone_mexico_mountain = 95
    time_zone_mexico_pacific = 96
    time_zone_cape_town = 97
    time_zone_winkhoek = 98
    time_zone_lagos = 99
    time_zone_riyahd = 100
    time_zone_venezuela = 101
    time_zone_australia_lh = 102
    time_zone_santiago = 103
    time_zone_manual = 253
    time_zone_automatic = 254
)


type display_measure uint8
const (
    display_measure_metric = 0
    display_measure_statute = 1
)


type display_heart uint8
const (
    display_heart_bpm = 0
    display_heart_max = 1
    display_heart_reserve = 2
)


type display_power uint8
const (
    display_power_watts = 0
    display_power_percent_ftp = 1
)


type display_position uint8
const (
    display_position_degree = 0
    display_position_degree_minute = 1
    display_position_degree_minute_second = 2
    display_position_austrian_grid = 3
    display_position_british_grid = 4
    display_position_dutch_grid = 5
    display_position_hungarian_grid = 6
    display_position_finnish_grid = 7
    display_position_german_grid = 8
    display_position_icelandic_grid = 9
    display_position_indonesian_equatorial = 10
    display_position_indonesian_irian = 11
    display_position_indonesian_southern = 12
    display_position_india_zone_0 = 13
    display_position_india_zone_IA = 14
    display_position_india_zone_IB = 15
    display_position_india_zone_IIA = 16
    display_position_india_zone_IIB = 17
    display_position_india_zone_IIIA = 18
    display_position_india_zone_IIIB = 19
    display_position_india_zone_IVA = 20
    display_position_india_zone_IVB = 21
    display_position_irish_transverse = 22
    display_position_irish_grid = 23
    display_position_loran = 24
    display_position_maidenhead_grid = 25
    display_position_mgrs_grid = 26
    display_position_new_zealand_grid = 27
    display_position_new_zealand_transverse = 28
    display_position_qatar_grid = 29
    display_position_modified_swedish_grid = 30
    display_position_swedish_grid = 31
    display_position_south_african_grid = 32
    display_position_swiss_grid = 33
    display_position_taiwan_grid = 34
    display_position_united_states_grid = 35
    display_position_utm_ups_grid = 36
    display_position_west_malayan = 37
    display_position_borneo_rso = 38
    display_position_estonian_grid = 39
    display_position_latvian_grid = 40
    display_position_swedish_ref_99_grid = 41
)


type sport uint8
const (
    sport_generic = 0
    sport_running = 1
    sport_cycling = 2
    sport_transition = 3
    sport_fitness_equipment = 4
    sport_swimming = 5
    sport_basketball = 6
    sport_soccer = 7
    sport_tennis = 8
    sport_american_football = 9
    sport_training = 10
    sport_walking = 11
    sport_cross_country_skiing = 12
    sport_alpine_skiing = 13
    sport_snowboarding = 14
    sport_rowing = 15
    sport_mountaineering = 16
    sport_hiking = 17
    sport_multisport = 18
    sport_paddling = 19
    sport_flying = 20
    sport_e_biking = 21
    sport_all = 254
)


type sport_bits_0 uint8


type sport_bits_1 uint8


type sport_bits_2 uint8


type sub_sport uint8
const (
    sub_sport_generic = 0
    sub_sport_treadmill = 1
    sub_sport_street = 2
    sub_sport_trail = 3
    sub_sport_track = 4
    sub_sport_spin = 5
    sub_sport_indoor_cycling = 6
    sub_sport_road = 7
    sub_sport_mountain = 8
    sub_sport_downhill = 9
    sub_sport_recumbent = 10
    sub_sport_cyclocross = 11
    sub_sport_hand_cycling = 12
    sub_sport_track_cycling = 13
    sub_sport_indoor_rowing = 14
    sub_sport_elliptical = 15
    sub_sport_stair_climbing = 16
    sub_sport_lap_swimming = 17
    sub_sport_open_water = 18
    sub_sport_flexibility_training = 19
    sub_sport_strength_training = 20
    sub_sport_warm_up = 21
    sub_sport_match = 22
    sub_sport_exercise = 23
    sub_sport_challenge = 24
    sub_sport_indoor_skiing = 25
    sub_sport_cardio_training = 26
    sub_sport_indoor_walking = 27
    sub_sport_e_bike_fitness = 28
    sub_sport_all = 254
)


type sport_event uint8
const (
    sport_event_uncategorized = 0
    sport_event_geocaching = 1
    sport_event_fitness = 2
    sport_event_recreation = 3
    sport_event_race = 4
    sport_event_special_event = 5
    sport_event_training = 6
    sport_event_transportation = 7
    sport_event_touring = 8
)


type activity uint8
const (
    activity_manual = 0
    activity_auto_multi_sport = 1
)


type intensity uint8
const (
    intensity_active = 0
    intensity_rest = 1
    intensity_warmup = 2
    intensity_cooldown = 3
)


type session_trigger uint8
const (
    session_trigger_activity_end = 0
    session_trigger_manual = 1
    session_trigger_auto_multi_sport = 2
    session_trigger_fitness_equipment = 3
)


type autolap_trigger uint8
const (
    autolap_trigger_time = 0
    autolap_trigger_distance = 1
    autolap_trigger_position_start = 2
    autolap_trigger_position_lap = 3
    autolap_trigger_position_waypoint = 4
    autolap_trigger_position_marked = 5
    autolap_trigger_off = 6
)


type lap_trigger uint8
const (
    lap_trigger_manual = 0
    lap_trigger_time = 1
    lap_trigger_distance = 2
    lap_trigger_position_start = 3
    lap_trigger_position_lap = 4
    lap_trigger_position_waypoint = 5
    lap_trigger_position_marked = 6
    lap_trigger_session_end = 7
    lap_trigger_fitness_equipment = 8
)


type event uint8
const (
    event_timer = 0
    event_workout = 3
    event_workout_step = 4
    event_power_down = 5
    event_power_up = 6
    event_off_course = 7
    event_session = 8
    event_lap = 9
    event_course_point = 10
    event_battery = 11
    event_virtual_partner_pace = 12
    event_hr_high_alert = 13
    event_hr_low_alert = 14
    event_speed_high_alert = 15
    event_speed_low_alert = 16
    event_cad_high_alert = 17
    event_cad_low_alert = 18
    event_power_high_alert = 19
    event_power_low_alert = 20
    event_recovery_hr = 21
    event_battery_low = 22
    event_time_duration_alert = 23
    event_distance_duration_alert = 24
    event_calorie_duration_alert = 25
    event_activity = 26
    event_fitness_equipment = 27
    event_length = 28
    event_user_marker = 32
    event_sport_point = 33
    event_calibration = 36
    event_front_gear_change = 42
    event_rear_gear_change = 43
    event_rider_position_change = 44
    event_elev_high_alert = 45
    event_elev_low_alert = 46
    event_comm_timeout = 47
)


type event_type uint8
const (
    event_type_start = 0
    event_type_stop = 1
    event_type_consecutive_depreciated = 2
    event_type_marker = 3
    event_type_stop_all = 4
    event_type_begin_depreciated = 5
    event_type_end_depreciated = 6
    event_type_end_all_depreciated = 7
    event_type_stop_disable = 8
    event_type_stop_disable_all = 9
)


type timer_trigger uint8
const (
    timer_trigger_manual = 0
    timer_trigger_auto = 1
    timer_trigger_fitness_equipment = 2
)


type fitness_equipment_state uint8
const (
    fitness_equipment_state_ready = 0
    fitness_equipment_state_in_use = 1
    fitness_equipment_state_paused = 2
    fitness_equipment_state_unknown = 3
)


type activity_class uint8
const (
    activity_class_level = 127
    activity_class_level_max = 100
    activity_class_athlete = 128
)


type hr_zone_calc uint8
const (
    hr_zone_calc_custom = 0
    hr_zone_calc_percent_max_hr = 1
    hr_zone_calc_percent_hrr = 2
)


type pwr_zone_calc uint8
const (
    pwr_zone_calc_custom = 0
    pwr_zone_calc_percent_ftp = 1
)


type wkt_step_duration uint8
const (
    wkt_step_duration_time = 0
    wkt_step_duration_distance = 1
    wkt_step_duration_hr_less_than = 2
    wkt_step_duration_hr_greater_than = 3
    wkt_step_duration_calories = 4
    wkt_step_duration_open = 5
    wkt_step_duration_repeat_until_steps_cmplt = 6
    wkt_step_duration_repeat_until_time = 7
    wkt_step_duration_repeat_until_distance = 8
    wkt_step_duration_repeat_until_calories = 9
    wkt_step_duration_repeat_until_hr_less_than = 10
    wkt_step_duration_repeat_until_hr_greater_than = 11
    wkt_step_duration_repeat_until_power_less_than = 12
    wkt_step_duration_repeat_until_power_greater_than = 13
    wkt_step_duration_power_less_than = 14
    wkt_step_duration_power_greater_than = 15
    wkt_step_duration_repetition_time = 28
)


type wkt_step_target uint8
const (
    wkt_step_target_speed = 0
    wkt_step_target_heart_rate = 1
    wkt_step_target_open = 2
    wkt_step_target_cadence = 3
    wkt_step_target_power = 4
    wkt_step_target_grade = 5
    wkt_step_target_resistance = 6
)


type goal uint8
const (
    goal_time = 0
    goal_distance = 1
    goal_calories = 2
    goal_frequency = 3
    goal_steps = 4
)


type goal_recurrence uint8
const (
    goal_recurrence_off = 0
    goal_recurrence_daily = 1
    goal_recurrence_weekly = 2
    goal_recurrence_monthly = 3
    goal_recurrence_yearly = 4
    goal_recurrence_custom = 5
)


type schedule uint8
const (
    schedule_workout = 0
    schedule_course = 1
)


type course_point uint8
const (
    course_point_generic = 0
    course_point_summit = 1
    course_point_valley = 2
    course_point_water = 3
    course_point_food = 4
    course_point_danger = 5
    course_point_left = 6
    course_point_right = 7
    course_point_straight = 8
    course_point_first_aid = 9
    course_point_fourth_category = 10
    course_point_third_category = 11
    course_point_second_category = 12
    course_point_first_category = 13
    course_point_hors_category = 14
    course_point_sprint = 15
    course_point_left_fork = 16
    course_point_right_fork = 17
    course_point_middle_fork = 18
    course_point_slight_left = 19
    course_point_sharp_left = 20
    course_point_slight_right = 21
    course_point_sharp_right = 22
    course_point_u_turn = 23
)


type manufacturer uint16


type garmin_product uint16


type antplus_device_type uint8


type ant_network uint8
const (
    ant_network_public = 0
    ant_network_antplus = 1
    ant_network_antfs = 2
    ant_network_private = 3
)


type workout_capabilities uint32


type battery_status uint8


type hr_type uint8
const (
    hr_type_normal = 0
    hr_type_irregular = 1
)


type course_capabilities uint32


type weight uint16


type workout_hr uint32


type workout_power uint32


type bp_status uint8
const (
    bp_status_no_error = 0
    bp_status_error_incomplete_data = 1
    bp_status_error_no_measurement = 2
    bp_status_error_data_out_of_range = 3
    bp_status_error_irregular_heart_rate = 4
)


type user_local_id uint16


type swim_stroke uint8
const (
    swim_stroke_freestyle = 0
    swim_stroke_backstroke = 1
    swim_stroke_breaststroke = 2
    swim_stroke_butterfly = 3
    swim_stroke_drill = 4
    swim_stroke_mixed = 5
    swim_stroke_im = 6
)


type activity_type uint8
const (
    activity_type_generic = 0
    activity_type_running = 1
    activity_type_cycling = 2
    activity_type_transition = 3
    activity_type_fitness_equipment = 4
    activity_type_swimming = 5
    activity_type_walking = 6
    activity_type_all = 254
)


type activity_subtype uint8
const (
    activity_subtype_generic = 0
    activity_subtype_treadmill = 1
    activity_subtype_street = 2
    activity_subtype_trail = 3
    activity_subtype_track = 4
    activity_subtype_spin = 5
    activity_subtype_indoor_cycling = 6
    activity_subtype_road = 7
    activity_subtype_mountain = 8
    activity_subtype_downhill = 9
    activity_subtype_recumbent = 10
    activity_subtype_cyclocross = 11
    activity_subtype_hand_cycling = 12
    activity_subtype_track_cycling = 13
    activity_subtype_indoor_rowing = 14
    activity_subtype_elliptical = 15
    activity_subtype_stair_climbing = 16
    activity_subtype_lap_swimming = 17
    activity_subtype_open_water = 18
    activity_subtype_all = 254
)


type activity_level uint8
const (
    activity_level_low = 0
    activity_level_medium = 1
    activity_level_high = 2
)


type left_right_balance uint8


type left_right_balance_100 uint16


type length_type uint8
const (
    length_type_idle = 0
    length_type_active = 1
)


type connectivity_capabilities uint32


type stroke_type uint8
const (
    stroke_type_no_event = 0
    stroke_type_other = 1
    stroke_type_serve = 2
    stroke_type_forehand = 3
    stroke_type_backhand = 4
    stroke_type_smash = 5
)


type body_location uint8
const (
    body_location_left_leg = 0
    body_location_left_calf = 1
    body_location_left_shin = 2
    body_location_left_hamstring = 3
    body_location_left_quad = 4
    body_location_left_glute = 5
    body_location_right_leg = 6
    body_location_right_calf = 7
    body_location_right_shin = 8
    body_location_right_hamstring = 9
    body_location_right_quad = 10
    body_location_right_glute = 11
    body_location_torso_back = 12
    body_location_left_lower_back = 13
    body_location_left_upper_back = 14
    body_location_right_lower_back = 15
    body_location_right_upper_back = 16
    body_location_torso_front = 17
    body_location_left_abdomen = 18
    body_location_left_chest = 19
    body_location_right_abdomen = 20
    body_location_right_chest = 21
    body_location_left_arm = 22
    body_location_left_shoulder = 23
    body_location_left_bicep = 24
    body_location_left_tricep = 25
    body_location_left_brachioradialis = 26
    body_location_left_forearm_extensors = 27
    body_location_right_arm = 28
    body_location_right_shoulder = 29
    body_location_right_bicep = 30
    body_location_right_tricep = 31
    body_location_right_brachioradialis = 32
    body_location_right_forearm_extensors = 33
    body_location_neck = 34
    body_location_throat = 35
)


type segment_lap_status uint8
const (
    segment_lap_status_end = 0
    segment_lap_status_fail = 1
)


type segment_leaderboard_type uint8
const (
    segment_leaderboard_type_overall = 0
    segment_leaderboard_type_personal_best = 1
    segment_leaderboard_type_connections = 2
    segment_leaderboard_type_group = 3
    segment_leaderboard_type_challenger = 4
    segment_leaderboard_type_kom = 5
    segment_leaderboard_type_qom = 6
    segment_leaderboard_type_pr = 7
    segment_leaderboard_type_goal = 8
    segment_leaderboard_type_rival = 9
    segment_leaderboard_type_club_leader = 10
)


type segment_delete_status uint8
const (
    segment_delete_status_do_not_delete = 0
    segment_delete_status_delete_one = 1
    segment_delete_status_delete_all = 2
)


type segment_selection_type uint8
const (
    segment_selection_type_starred = 0
    segment_selection_type_suggested = 1
)


type source_type uint8
const (
    source_type_ant = 0
    source_type_antplus = 1
    source_type_bluetooth = 2
    source_type_bluetooth_low_energy = 3
    source_type_wifi = 4
    source_type_local = 5
)


type rider_position_type uint8
const (
    rider_position_type_seated = 0
    rider_position_type_standing = 1
)


type power_phase_type uint8
const (
    power_phase_type_power_phase_start_angle = 0
    power_phase_type_power_phase_end_angle = 1
    power_phase_type_power_phase_arc_length = 2
    power_phase_type_power_phase_center = 3
)



