package tz

import "sync"

// GENERATED FILE DO NOT MODIFY DIRECTLY

var (
	once      sync.Once
	mapped    map[string]Country
	countries = []Country{
		{
			Code: "AF",
			Name: "Afghanistan",
			Zones: []Zone{
				{
					CountryCode: "AF",
					Name:        "Asia/Kabul",
				},
			},
		},
		{
			Code: "AL",
			Name: "Albania",
			Zones: []Zone{
				{
					CountryCode: "AL",
					Name:        "Europe/Tirane",
				},
			},
		},
		{
			Code: "DZ",
			Name: "Algeria",
			Zones: []Zone{
				{
					CountryCode: "DZ",
					Name:        "Africa/Algiers",
				},
			},
		},
		{
			Code: "AS",
			Name: "American Samoa",
			Zones: []Zone{
				{
					CountryCode: "AS",
					Name:        "Pacific/Pago_Pago",
				},
			},
		},
		{
			Code: "AD",
			Name: "Andorra",
			Zones: []Zone{
				{
					CountryCode: "AD",
					Name:        "Europe/Andorra",
				},
			},
		},
		{
			Code: "AO",
			Name: "Angola",
			Zones: []Zone{
				{
					CountryCode: "AO",
					Name:        "Africa/Luanda",
				},
			},
		},
		{
			Code: "AI",
			Name: "Anguilla",
			Zones: []Zone{
				{
					CountryCode: "AI",
					Name:        "America/Anguilla",
				},
			},
		},
		{
			Code: "AQ",
			Name: "Antarctica",
			Zones: []Zone{
				{
					CountryCode: "AQ",
					Name:        "Antarctica/Casey",
				},
				{
					CountryCode: "AQ",
					Name:        "Antarctica/Davis",
				},
				{
					CountryCode: "AQ",
					Name:        "Antarctica/DumontDUrville",
				},
				{
					CountryCode: "AQ",
					Name:        "Antarctica/Mawson",
				},
				{
					CountryCode: "AQ",
					Name:        "Antarctica/McMurdo",
				},
				{
					CountryCode: "AQ",
					Name:        "Antarctica/Palmer",
				},
				{
					CountryCode: "AQ",
					Name:        "Antarctica/Rothera",
				},
				{
					CountryCode: "AQ",
					Name:        "Antarctica/Syowa",
				},
				{
					CountryCode: "AQ",
					Name:        "Antarctica/Troll",
				},
				{
					CountryCode: "AQ",
					Name:        "Antarctica/Vostok",
				},
			},
		},
		{
			Code: "AG",
			Name: "Antigua and Barbuda",
			Zones: []Zone{
				{
					CountryCode: "AG",
					Name:        "America/Antigua",
				},
			},
		},
		{
			Code: "AR",
			Name: "Argentina",
			Zones: []Zone{
				{
					CountryCode: "AR",
					Name:        "America/Argentina/Buenos_Aires",
				},
				{
					CountryCode: "AR",
					Name:        "America/Argentina/Catamarca",
				},
				{
					CountryCode: "AR",
					Name:        "America/Argentina/Cordoba",
				},
				{
					CountryCode: "AR",
					Name:        "America/Argentina/Jujuy",
				},
				{
					CountryCode: "AR",
					Name:        "America/Argentina/La_Rioja",
				},
				{
					CountryCode: "AR",
					Name:        "America/Argentina/Mendoza",
				},
				{
					CountryCode: "AR",
					Name:        "America/Argentina/Rio_Gallegos",
				},
				{
					CountryCode: "AR",
					Name:        "America/Argentina/Salta",
				},
				{
					CountryCode: "AR",
					Name:        "America/Argentina/San_Juan",
				},
				{
					CountryCode: "AR",
					Name:        "America/Argentina/San_Luis",
				},
				{
					CountryCode: "AR",
					Name:        "America/Argentina/Tucuman",
				},
				{
					CountryCode: "AR",
					Name:        "America/Argentina/Ushuaia",
				},
			},
		},
		{
			Code: "AM",
			Name: "Armenia",
			Zones: []Zone{
				{
					CountryCode: "AM",
					Name:        "Asia/Yerevan",
				},
			},
		},
		{
			Code: "AW",
			Name: "Aruba",
			Zones: []Zone{
				{
					CountryCode: "AW",
					Name:        "America/Aruba",
				},
			},
		},
		{
			Code: "AU",
			Name: "Australia",
			Zones: []Zone{
				{
					CountryCode: "AU",
					Name:        "Antarctica/Macquarie",
				},
				{
					CountryCode: "AU",
					Name:        "Australia/Adelaide",
				},
				{
					CountryCode: "AU",
					Name:        "Australia/Brisbane",
				},
				{
					CountryCode: "AU",
					Name:        "Australia/Broken_Hill",
				},
				{
					CountryCode: "AU",
					Name:        "Australia/Currie",
				},
				{
					CountryCode: "AU",
					Name:        "Australia/Darwin",
				},
				{
					CountryCode: "AU",
					Name:        "Australia/Eucla",
				},
				{
					CountryCode: "AU",
					Name:        "Australia/Hobart",
				},
				{
					CountryCode: "AU",
					Name:        "Australia/Lindeman",
				},
				{
					CountryCode: "AU",
					Name:        "Australia/Lord_Howe",
				},
				{
					CountryCode: "AU",
					Name:        "Australia/Melbourne",
				},
				{
					CountryCode: "AU",
					Name:        "Australia/Perth",
				},
				{
					CountryCode: "AU",
					Name:        "Australia/Sydney",
				},
			},
		},
		{
			Code: "AT",
			Name: "Austria",
			Zones: []Zone{
				{
					CountryCode: "AT",
					Name:        "Europe/Vienna",
				},
			},
		},
		{
			Code: "AZ",
			Name: "Azerbaijan",
			Zones: []Zone{
				{
					CountryCode: "AZ",
					Name:        "Asia/Baku",
				},
			},
		},
		{
			Code: "BS",
			Name: "Bahamas",
			Zones: []Zone{
				{
					CountryCode: "BS",
					Name:        "America/Nassau",
				},
			},
		},
		{
			Code: "BH",
			Name: "Bahrain",
			Zones: []Zone{
				{
					CountryCode: "BH",
					Name:        "Asia/Bahrain",
				},
			},
		},
		{
			Code: "BD",
			Name: "Bangladesh",
			Zones: []Zone{
				{
					CountryCode: "BD",
					Name:        "Asia/Dhaka",
				},
			},
		},
		{
			Code: "BB",
			Name: "Barbados",
			Zones: []Zone{
				{
					CountryCode: "BB",
					Name:        "America/Barbados",
				},
			},
		},
		{
			Code: "BY",
			Name: "Belarus",
			Zones: []Zone{
				{
					CountryCode: "BY",
					Name:        "Europe/Minsk",
				},
			},
		},
		{
			Code: "BE",
			Name: "Belgium",
			Zones: []Zone{
				{
					CountryCode: "BE",
					Name:        "Europe/Brussels",
				},
			},
		},
		{
			Code: "BZ",
			Name: "Belize",
			Zones: []Zone{
				{
					CountryCode: "BZ",
					Name:        "America/Belize",
				},
			},
		},
		{
			Code: "BJ",
			Name: "Benin",
			Zones: []Zone{
				{
					CountryCode: "BJ",
					Name:        "Africa/Porto-Novo",
				},
			},
		},
		{
			Code: "BM",
			Name: "Bermuda",
			Zones: []Zone{
				{
					CountryCode: "BM",
					Name:        "Atlantic/Bermuda",
				},
			},
		},
		{
			Code: "BT",
			Name: "Bhutan",
			Zones: []Zone{
				{
					CountryCode: "BT",
					Name:        "Asia/Thimphu",
				},
			},
		},
		{
			Code: "BO",
			Name: "Bolivia, Plurinational State of",
			Zones: []Zone{
				{
					CountryCode: "BO",
					Name:        "America/La_Paz",
				},
			},
		},
		{
			Code: "BQ",
			Name: "Bonaire, Sint Eustatius and Saba",
			Zones: []Zone{
				{
					CountryCode: "BQ",
					Name:        "America/Kralendijk",
				},
			},
		},
		{
			Code: "BA",
			Name: "Bosnia and Herzegovina",
			Zones: []Zone{
				{
					CountryCode: "BA",
					Name:        "Europe/Sarajevo",
				},
			},
		},
		{
			Code: "BW",
			Name: "Botswana",
			Zones: []Zone{
				{
					CountryCode: "BW",
					Name:        "Africa/Gaborone",
				},
			},
		},
		{
			Code:  "BV",
			Name:  "Bouvet Island",
			Zones: []Zone{},
		},
		{
			Code: "BR",
			Name: "Brazil",
			Zones: []Zone{
				{
					CountryCode: "BR",
					Name:        "America/Araguaina",
				},
				{
					CountryCode: "BR",
					Name:        "America/Bahia",
				},
				{
					CountryCode: "BR",
					Name:        "America/Belem",
				},
				{
					CountryCode: "BR",
					Name:        "America/Boa_Vista",
				},
				{
					CountryCode: "BR",
					Name:        "America/Campo_Grande",
				},
				{
					CountryCode: "BR",
					Name:        "America/Cuiaba",
				},
				{
					CountryCode: "BR",
					Name:        "America/Eirunepe",
				},
				{
					CountryCode: "BR",
					Name:        "America/Fortaleza",
				},
				{
					CountryCode: "BR",
					Name:        "America/Maceio",
				},
				{
					CountryCode: "BR",
					Name:        "America/Manaus",
				},
				{
					CountryCode: "BR",
					Name:        "America/Noronha",
				},
				{
					CountryCode: "BR",
					Name:        "America/Porto_Velho",
				},
				{
					CountryCode: "BR",
					Name:        "America/Recife",
				},
				{
					CountryCode: "BR",
					Name:        "America/Rio_Branco",
				},
				{
					CountryCode: "BR",
					Name:        "America/Santarem",
				},
				{
					CountryCode: "BR",
					Name:        "America/Sao_Paulo",
				},
			},
		},
		{
			Code: "IO",
			Name: "British Indian Ocean Territory",
			Zones: []Zone{
				{
					CountryCode: "IO",
					Name:        "Indian/Chagos",
				},
			},
		},
		{
			Code: "BN",
			Name: "Brunei Darussalam",
			Zones: []Zone{
				{
					CountryCode: "BN",
					Name:        "Asia/Brunei",
				},
			},
		},
		{
			Code: "BG",
			Name: "Bulgaria",
			Zones: []Zone{
				{
					CountryCode: "BG",
					Name:        "Europe/Sofia",
				},
			},
		},
		{
			Code: "BF",
			Name: "Burkina Faso",
			Zones: []Zone{
				{
					CountryCode: "BF",
					Name:        "Africa/Ouagadougou",
				},
			},
		},
		{
			Code: "BI",
			Name: "Burundi",
			Zones: []Zone{
				{
					CountryCode: "BI",
					Name:        "Africa/Bujumbura",
				},
			},
		},
		{
			Code: "KH",
			Name: "Cambodia",
			Zones: []Zone{
				{
					CountryCode: "KH",
					Name:        "Asia/Phnom_Penh",
				},
			},
		},
		{
			Code: "CM",
			Name: "Cameroon",
			Zones: []Zone{
				{
					CountryCode: "CM",
					Name:        "Africa/Douala",
				},
			},
		},
		{
			Code: "CA",
			Name: "Canada",
			Zones: []Zone{
				{
					CountryCode: "CA",
					Name:        "America/Atikokan",
				},
				{
					CountryCode: "CA",
					Name:        "America/Blanc-Sablon",
				},
				{
					CountryCode: "CA",
					Name:        "America/Cambridge_Bay",
				},
				{
					CountryCode: "CA",
					Name:        "America/Creston",
				},
				{
					CountryCode: "CA",
					Name:        "America/Dawson",
				},
				{
					CountryCode: "CA",
					Name:        "America/Dawson_Creek",
				},
				{
					CountryCode: "CA",
					Name:        "America/Edmonton",
				},
				{
					CountryCode: "CA",
					Name:        "America/Fort_Nelson",
				},
				{
					CountryCode: "CA",
					Name:        "America/Glace_Bay",
				},
				{
					CountryCode: "CA",
					Name:        "America/Goose_Bay",
				},
				{
					CountryCode: "CA",
					Name:        "America/Halifax",
				},
				{
					CountryCode: "CA",
					Name:        "America/Inuvik",
				},
				{
					CountryCode: "CA",
					Name:        "America/Iqaluit",
				},
				{
					CountryCode: "CA",
					Name:        "America/Moncton",
				},
				{
					CountryCode: "CA",
					Name:        "America/Nipigon",
				},
				{
					CountryCode: "CA",
					Name:        "America/Pangnirtung",
				},
				{
					CountryCode: "CA",
					Name:        "America/Rainy_River",
				},
				{
					CountryCode: "CA",
					Name:        "America/Rankin_Inlet",
				},
				{
					CountryCode: "CA",
					Name:        "America/Regina",
				},
				{
					CountryCode: "CA",
					Name:        "America/Resolute",
				},
				{
					CountryCode: "CA",
					Name:        "America/St_Johns",
				},
				{
					CountryCode: "CA",
					Name:        "America/Swift_Current",
				},
				{
					CountryCode: "CA",
					Name:        "America/Thunder_Bay",
				},
				{
					CountryCode: "CA",
					Name:        "America/Toronto",
				},
				{
					CountryCode: "CA",
					Name:        "America/Vancouver",
				},
				{
					CountryCode: "CA",
					Name:        "America/Whitehorse",
				},
				{
					CountryCode: "CA",
					Name:        "America/Winnipeg",
				},
				{
					CountryCode: "CA",
					Name:        "America/Yellowknife",
				},
			},
		},
		{
			Code: "CV",
			Name: "Cape Verde",
			Zones: []Zone{
				{
					CountryCode: "CV",
					Name:        "Atlantic/Cape_Verde",
				},
			},
		},
		{
			Code: "KY",
			Name: "Cayman Islands",
			Zones: []Zone{
				{
					CountryCode: "KY",
					Name:        "America/Cayman",
				},
			},
		},
		{
			Code: "CF",
			Name: "Central African Republic",
			Zones: []Zone{
				{
					CountryCode: "CF",
					Name:        "Africa/Bangui",
				},
			},
		},
		{
			Code: "TD",
			Name: "Chad",
			Zones: []Zone{
				{
					CountryCode: "TD",
					Name:        "Africa/Ndjamena",
				},
			},
		},
		{
			Code: "CL",
			Name: "Chile",
			Zones: []Zone{
				{
					CountryCode: "CL",
					Name:        "America/Punta_Arenas",
				},
				{
					CountryCode: "CL",
					Name:        "America/Santiago",
				},
				{
					CountryCode: "CL",
					Name:        "Pacific/Easter",
				},
			},
		},
		{
			Code: "CN",
			Name: "China",
			Zones: []Zone{
				{
					CountryCode: "CN",
					Name:        "Asia/Shanghai",
				},
				{
					CountryCode: "CN",
					Name:        "Asia/Urumqi",
				},
			},
		},
		{
			Code: "CX",
			Name: "Christmas Island",
			Zones: []Zone{
				{
					CountryCode: "CX",
					Name:        "Indian/Christmas",
				},
			},
		},
		{
			Code: "CC",
			Name: "Cocos (Keeling) Islands",
			Zones: []Zone{
				{
					CountryCode: "CC",
					Name:        "Indian/Cocos",
				},
			},
		},
		{
			Code: "CO",
			Name: "Colombia",
			Zones: []Zone{
				{
					CountryCode: "CO",
					Name:        "America/Bogota",
				},
			},
		},
		{
			Code: "KM",
			Name: "Comoros",
			Zones: []Zone{
				{
					CountryCode: "KM",
					Name:        "Indian/Comoro",
				},
			},
		},
		{
			Code: "CG",
			Name: "Congo",
			Zones: []Zone{
				{
					CountryCode: "CG",
					Name:        "Africa/Brazzaville",
				},
			},
		},
		{
			Code: "CD",
			Name: "Congo, the Democratic Republic of the",
			Zones: []Zone{
				{
					CountryCode: "CD",
					Name:        "Africa/Kinshasa",
				},
				{
					CountryCode: "CD",
					Name:        "Africa/Lubumbashi",
				},
			},
		},
		{
			Code: "CK",
			Name: "Cook Islands",
			Zones: []Zone{
				{
					CountryCode: "CK",
					Name:        "Pacific/Rarotonga",
				},
			},
		},
		{
			Code: "CR",
			Name: "Costa Rica",
			Zones: []Zone{
				{
					CountryCode: "CR",
					Name:        "America/Costa_Rica",
				},
			},
		},
		{
			Code: "HR",
			Name: "Croatia",
			Zones: []Zone{
				{
					CountryCode: "HR",
					Name:        "Europe/Zagreb",
				},
			},
		},
		{
			Code: "CU",
			Name: "Cuba",
			Zones: []Zone{
				{
					CountryCode: "CU",
					Name:        "America/Havana",
				},
			},
		},
		{
			Code: "CW",
			Name: "Curaçao",
			Zones: []Zone{
				{
					CountryCode: "CW",
					Name:        "America/Curacao",
				},
			},
		},
		{
			Code: "CY",
			Name: "Cyprus",
			Zones: []Zone{
				{
					CountryCode: "CY",
					Name:        "Asia/Famagusta",
				},
				{
					CountryCode: "CY",
					Name:        "Asia/Nicosia",
				},
			},
		},
		{
			Code: "CZ",
			Name: "Czech Republic",
			Zones: []Zone{
				{
					CountryCode: "CZ",
					Name:        "Europe/Prague",
				},
			},
		},
		{
			Code: "CI",
			Name: "Côte d'Ivoire",
			Zones: []Zone{
				{
					CountryCode: "CI",
					Name:        "Africa/Abidjan",
				},
			},
		},
		{
			Code: "DK",
			Name: "Denmark",
			Zones: []Zone{
				{
					CountryCode: "DK",
					Name:        "Europe/Copenhagen",
				},
			},
		},
		{
			Code: "DJ",
			Name: "Djibouti",
			Zones: []Zone{
				{
					CountryCode: "DJ",
					Name:        "Africa/Djibouti",
				},
			},
		},
		{
			Code: "DM",
			Name: "Dominica",
			Zones: []Zone{
				{
					CountryCode: "DM",
					Name:        "America/Dominica",
				},
			},
		},
		{
			Code: "DO",
			Name: "Dominican Republic",
			Zones: []Zone{
				{
					CountryCode: "DO",
					Name:        "America/Santo_Domingo",
				},
			},
		},
		{
			Code: "EC",
			Name: "Ecuador",
			Zones: []Zone{
				{
					CountryCode: "EC",
					Name:        "America/Guayaquil",
				},
				{
					CountryCode: "EC",
					Name:        "Pacific/Galapagos",
				},
			},
		},
		{
			Code: "EG",
			Name: "Egypt",
			Zones: []Zone{
				{
					CountryCode: "EG",
					Name:        "Africa/Cairo",
				},
			},
		},
		{
			Code: "SV",
			Name: "El Salvador",
			Zones: []Zone{
				{
					CountryCode: "SV",
					Name:        "America/El_Salvador",
				},
			},
		},
		{
			Code: "GQ",
			Name: "Equatorial Guinea",
			Zones: []Zone{
				{
					CountryCode: "GQ",
					Name:        "Africa/Malabo",
				},
			},
		},
		{
			Code: "ER",
			Name: "Eritrea",
			Zones: []Zone{
				{
					CountryCode: "ER",
					Name:        "Africa/Asmara",
				},
			},
		},
		{
			Code: "EE",
			Name: "Estonia",
			Zones: []Zone{
				{
					CountryCode: "EE",
					Name:        "Europe/Tallinn",
				},
			},
		},
		{
			Code: "ET",
			Name: "Ethiopia",
			Zones: []Zone{
				{
					CountryCode: "ET",
					Name:        "Africa/Addis_Ababa",
				},
			},
		},
		{
			Code: "FK",
			Name: "Falkland Islands (Malvinas)",
			Zones: []Zone{
				{
					CountryCode: "FK",
					Name:        "Atlantic/Stanley",
				},
			},
		},
		{
			Code: "FO",
			Name: "Faroe Islands",
			Zones: []Zone{
				{
					CountryCode: "FO",
					Name:        "Atlantic/Faroe",
				},
			},
		},
		{
			Code: "FJ",
			Name: "Fiji",
			Zones: []Zone{
				{
					CountryCode: "FJ",
					Name:        "Pacific/Fiji",
				},
			},
		},
		{
			Code: "FI",
			Name: "Finland",
			Zones: []Zone{
				{
					CountryCode: "FI",
					Name:        "Europe/Helsinki",
				},
			},
		},
		{
			Code: "FR",
			Name: "France",
			Zones: []Zone{
				{
					CountryCode: "FR",
					Name:        "Europe/Paris",
				},
			},
		},
		{
			Code: "GF",
			Name: "French Guiana",
			Zones: []Zone{
				{
					CountryCode: "GF",
					Name:        "America/Cayenne",
				},
			},
		},
		{
			Code: "PF",
			Name: "French Polynesia",
			Zones: []Zone{
				{
					CountryCode: "PF",
					Name:        "Pacific/Gambier",
				},
				{
					CountryCode: "PF",
					Name:        "Pacific/Marquesas",
				},
				{
					CountryCode: "PF",
					Name:        "Pacific/Tahiti",
				},
			},
		},
		{
			Code: "TF",
			Name: "French Southern Territories",
			Zones: []Zone{
				{
					CountryCode: "TF",
					Name:        "Indian/Kerguelen",
				},
			},
		},
		{
			Code: "GA",
			Name: "Gabon",
			Zones: []Zone{
				{
					CountryCode: "GA",
					Name:        "Africa/Libreville",
				},
			},
		},
		{
			Code: "GM",
			Name: "Gambia",
			Zones: []Zone{
				{
					CountryCode: "GM",
					Name:        "Africa/Banjul",
				},
			},
		},
		{
			Code: "GE",
			Name: "Georgia",
			Zones: []Zone{
				{
					CountryCode: "GE",
					Name:        "Asia/Tbilisi",
				},
			},
		},
		{
			Code: "DE",
			Name: "Germany",
			Zones: []Zone{
				{
					CountryCode: "DE",
					Name:        "Europe/Berlin",
				},
				{
					CountryCode: "DE",
					Name:        "Europe/Busingen",
				},
			},
		},
		{
			Code: "GH",
			Name: "Ghana",
			Zones: []Zone{
				{
					CountryCode: "GH",
					Name:        "Africa/Accra",
				},
			},
		},
		{
			Code: "GI",
			Name: "Gibraltar",
			Zones: []Zone{
				{
					CountryCode: "GI",
					Name:        "Europe/Gibraltar",
				},
			},
		},
		{
			Code: "GR",
			Name: "Greece",
			Zones: []Zone{
				{
					CountryCode: "GR",
					Name:        "Europe/Athens",
				},
			},
		},
		{
			Code: "GL",
			Name: "Greenland",
			Zones: []Zone{
				{
					CountryCode: "GL",
					Name:        "America/Danmarkshavn",
				},
				{
					CountryCode: "GL",
					Name:        "America/Godthab",
				},
				{
					CountryCode: "GL",
					Name:        "America/Scoresbysund",
				},
				{
					CountryCode: "GL",
					Name:        "America/Thule",
				},
			},
		},
		{
			Code: "GD",
			Name: "Grenada",
			Zones: []Zone{
				{
					CountryCode: "GD",
					Name:        "America/Grenada",
				},
			},
		},
		{
			Code: "GP",
			Name: "Guadeloupe",
			Zones: []Zone{
				{
					CountryCode: "GP",
					Name:        "America/Guadeloupe",
				},
			},
		},
		{
			Code: "GU",
			Name: "Guam",
			Zones: []Zone{
				{
					CountryCode: "GU",
					Name:        "Pacific/Guam",
				},
			},
		},
		{
			Code: "GT",
			Name: "Guatemala",
			Zones: []Zone{
				{
					CountryCode: "GT",
					Name:        "America/Guatemala",
				},
			},
		},
		{
			Code: "GG",
			Name: "Guernsey",
			Zones: []Zone{
				{
					CountryCode: "GG",
					Name:        "Europe/Guernsey",
				},
			},
		},
		{
			Code: "GN",
			Name: "Guinea",
			Zones: []Zone{
				{
					CountryCode: "GN",
					Name:        "Africa/Conakry",
				},
			},
		},
		{
			Code: "GW",
			Name: "Guinea-Bissau",
			Zones: []Zone{
				{
					CountryCode: "GW",
					Name:        "Africa/Bissau",
				},
			},
		},
		{
			Code: "GY",
			Name: "Guyana",
			Zones: []Zone{
				{
					CountryCode: "GY",
					Name:        "America/Guyana",
				},
			},
		},
		{
			Code: "HT",
			Name: "Haiti",
			Zones: []Zone{
				{
					CountryCode: "HT",
					Name:        "America/Port-au-Prince",
				},
			},
		},
		{
			Code:  "HM",
			Name:  "Heard Island and McDonald Islands",
			Zones: []Zone{},
		},
		{
			Code: "VA",
			Name: "Holy See (Vatican City State)",
			Zones: []Zone{
				{
					CountryCode: "VA",
					Name:        "Europe/Vatican",
				},
			},
		},
		{
			Code: "HN",
			Name: "Honduras",
			Zones: []Zone{
				{
					CountryCode: "HN",
					Name:        "America/Tegucigalpa",
				},
			},
		},
		{
			Code: "HK",
			Name: "Hong Kong",
			Zones: []Zone{
				{
					CountryCode: "HK",
					Name:        "Asia/Hong_Kong",
				},
			},
		},
		{
			Code: "HU",
			Name: "Hungary",
			Zones: []Zone{
				{
					CountryCode: "HU",
					Name:        "Europe/Budapest",
				},
			},
		},
		{
			Code: "IS",
			Name: "Iceland",
			Zones: []Zone{
				{
					CountryCode: "IS",
					Name:        "Atlantic/Reykjavik",
				},
			},
		},
		{
			Code: "IN",
			Name: "India",
			Zones: []Zone{
				{
					CountryCode: "IN",
					Name:        "Asia/Kolkata",
				},
			},
		},
		{
			Code: "ID",
			Name: "Indonesia",
			Zones: []Zone{
				{
					CountryCode: "ID",
					Name:        "Asia/Jakarta",
				},
				{
					CountryCode: "ID",
					Name:        "Asia/Jayapura",
				},
				{
					CountryCode: "ID",
					Name:        "Asia/Makassar",
				},
				{
					CountryCode: "ID",
					Name:        "Asia/Pontianak",
				},
			},
		},
		{
			Code: "IR",
			Name: "Iran, Islamic Republic of",
			Zones: []Zone{
				{
					CountryCode: "IR",
					Name:        "Asia/Tehran",
				},
			},
		},
		{
			Code: "IQ",
			Name: "Iraq",
			Zones: []Zone{
				{
					CountryCode: "IQ",
					Name:        "Asia/Baghdad",
				},
			},
		},
		{
			Code: "IE",
			Name: "Ireland",
			Zones: []Zone{
				{
					CountryCode: "IE",
					Name:        "Europe/Dublin",
				},
			},
		},
		{
			Code: "IM",
			Name: "Isle of Man",
			Zones: []Zone{
				{
					CountryCode: "IM",
					Name:        "Europe/Isle_of_Man",
				},
			},
		},
		{
			Code: "IL",
			Name: "Israel",
			Zones: []Zone{
				{
					CountryCode: "IL",
					Name:        "Asia/Jerusalem",
				},
			},
		},
		{
			Code: "IT",
			Name: "Italy",
			Zones: []Zone{
				{
					CountryCode: "IT",
					Name:        "Europe/Rome",
				},
			},
		},
		{
			Code: "JM",
			Name: "Jamaica",
			Zones: []Zone{
				{
					CountryCode: "JM",
					Name:        "America/Jamaica",
				},
			},
		},
		{
			Code: "JP",
			Name: "Japan",
			Zones: []Zone{
				{
					CountryCode: "JP",
					Name:        "Asia/Tokyo",
				},
			},
		},
		{
			Code: "JE",
			Name: "Jersey",
			Zones: []Zone{
				{
					CountryCode: "JE",
					Name:        "Europe/Jersey",
				},
			},
		},
		{
			Code: "JO",
			Name: "Jordan",
			Zones: []Zone{
				{
					CountryCode: "JO",
					Name:        "Asia/Amman",
				},
			},
		},
		{
			Code: "KZ",
			Name: "Kazakhstan",
			Zones: []Zone{
				{
					CountryCode: "KZ",
					Name:        "Asia/Almaty",
				},
				{
					CountryCode: "KZ",
					Name:        "Asia/Aqtau",
				},
				{
					CountryCode: "KZ",
					Name:        "Asia/Aqtobe",
				},
				{
					CountryCode: "KZ",
					Name:        "Asia/Atyrau",
				},
				{
					CountryCode: "KZ",
					Name:        "Asia/Oral",
				},
				{
					CountryCode: "KZ",
					Name:        "Asia/Qostanay",
				},
				{
					CountryCode: "KZ",
					Name:        "Asia/Qyzylorda",
				},
			},
		},
		{
			Code: "KE",
			Name: "Kenya",
			Zones: []Zone{
				{
					CountryCode: "KE",
					Name:        "Africa/Nairobi",
				},
			},
		},
		{
			Code: "KI",
			Name: "Kiribati",
			Zones: []Zone{
				{
					CountryCode: "KI",
					Name:        "Pacific/Enderbury",
				},
				{
					CountryCode: "KI",
					Name:        "Pacific/Kiritimati",
				},
				{
					CountryCode: "KI",
					Name:        "Pacific/Tarawa",
				},
			},
		},
		{
			Code: "KP",
			Name: "Korea, Democratic People's Republic of",
			Zones: []Zone{
				{
					CountryCode: "KP",
					Name:        "Asia/Pyongyang",
				},
			},
		},
		{
			Code: "KR",
			Name: "Korea, Republic of",
			Zones: []Zone{
				{
					CountryCode: "KR",
					Name:        "Asia/Seoul",
				},
			},
		},
		{
			Code: "KW",
			Name: "Kuwait",
			Zones: []Zone{
				{
					CountryCode: "KW",
					Name:        "Asia/Kuwait",
				},
			},
		},
		{
			Code: "KG",
			Name: "Kyrgyzstan",
			Zones: []Zone{
				{
					CountryCode: "KG",
					Name:        "Asia/Bishkek",
				},
			},
		},
		{
			Code: "LA",
			Name: "Lao People's Democratic Republic",
			Zones: []Zone{
				{
					CountryCode: "LA",
					Name:        "Asia/Vientiane",
				},
			},
		},
		{
			Code: "LV",
			Name: "Latvia",
			Zones: []Zone{
				{
					CountryCode: "LV",
					Name:        "Europe/Riga",
				},
			},
		},
		{
			Code: "LB",
			Name: "Lebanon",
			Zones: []Zone{
				{
					CountryCode: "LB",
					Name:        "Asia/Beirut",
				},
			},
		},
		{
			Code: "LS",
			Name: "Lesotho",
			Zones: []Zone{
				{
					CountryCode: "LS",
					Name:        "Africa/Maseru",
				},
			},
		},
		{
			Code: "LR",
			Name: "Liberia",
			Zones: []Zone{
				{
					CountryCode: "LR",
					Name:        "Africa/Monrovia",
				},
			},
		},
		{
			Code: "LY",
			Name: "Libya",
			Zones: []Zone{
				{
					CountryCode: "LY",
					Name:        "Africa/Tripoli",
				},
			},
		},
		{
			Code: "LI",
			Name: "Liechtenstein",
			Zones: []Zone{
				{
					CountryCode: "LI",
					Name:        "Europe/Vaduz",
				},
			},
		},
		{
			Code: "LT",
			Name: "Lithuania",
			Zones: []Zone{
				{
					CountryCode: "LT",
					Name:        "Europe/Vilnius",
				},
			},
		},
		{
			Code: "LU",
			Name: "Luxembourg",
			Zones: []Zone{
				{
					CountryCode: "LU",
					Name:        "Europe/Luxembourg",
				},
			},
		},
		{
			Code: "MO",
			Name: "Macao",
			Zones: []Zone{
				{
					CountryCode: "MO",
					Name:        "Asia/Macau",
				},
			},
		},
		{
			Code: "MK",
			Name: "Macedonia, the Former Yugoslav Republic of",
			Zones: []Zone{
				{
					CountryCode: "MK",
					Name:        "Europe/Skopje",
				},
			},
		},
		{
			Code: "MG",
			Name: "Madagascar",
			Zones: []Zone{
				{
					CountryCode: "MG",
					Name:        "Indian/Antananarivo",
				},
			},
		},
		{
			Code: "MW",
			Name: "Malawi",
			Zones: []Zone{
				{
					CountryCode: "MW",
					Name:        "Africa/Blantyre",
				},
			},
		},
		{
			Code: "MY",
			Name: "Malaysia",
			Zones: []Zone{
				{
					CountryCode: "MY",
					Name:        "Asia/Kuala_Lumpur",
				},
				{
					CountryCode: "MY",
					Name:        "Asia/Kuching",
				},
			},
		},
		{
			Code: "MV",
			Name: "Maldives",
			Zones: []Zone{
				{
					CountryCode: "MV",
					Name:        "Indian/Maldives",
				},
			},
		},
		{
			Code: "ML",
			Name: "Mali",
			Zones: []Zone{
				{
					CountryCode: "ML",
					Name:        "Africa/Bamako",
				},
			},
		},
		{
			Code: "MT",
			Name: "Malta",
			Zones: []Zone{
				{
					CountryCode: "MT",
					Name:        "Europe/Malta",
				},
			},
		},
		{
			Code: "MH",
			Name: "Marshall Islands",
			Zones: []Zone{
				{
					CountryCode: "MH",
					Name:        "Pacific/Kwajalein",
				},
				{
					CountryCode: "MH",
					Name:        "Pacific/Majuro",
				},
			},
		},
		{
			Code: "MQ",
			Name: "Martinique",
			Zones: []Zone{
				{
					CountryCode: "MQ",
					Name:        "America/Martinique",
				},
			},
		},
		{
			Code: "MR",
			Name: "Mauritania",
			Zones: []Zone{
				{
					CountryCode: "MR",
					Name:        "Africa/Nouakchott",
				},
			},
		},
		{
			Code: "MU",
			Name: "Mauritius",
			Zones: []Zone{
				{
					CountryCode: "MU",
					Name:        "Indian/Mauritius",
				},
			},
		},
		{
			Code: "YT",
			Name: "Mayotte",
			Zones: []Zone{
				{
					CountryCode: "YT",
					Name:        "Indian/Mayotte",
				},
			},
		},
		{
			Code: "MX",
			Name: "Mexico",
			Zones: []Zone{
				{
					CountryCode: "MX",
					Name:        "America/Bahia_Banderas",
				},
				{
					CountryCode: "MX",
					Name:        "America/Cancun",
				},
				{
					CountryCode: "MX",
					Name:        "America/Chihuahua",
				},
				{
					CountryCode: "MX",
					Name:        "America/Hermosillo",
				},
				{
					CountryCode: "MX",
					Name:        "America/Matamoros",
				},
				{
					CountryCode: "MX",
					Name:        "America/Mazatlan",
				},
				{
					CountryCode: "MX",
					Name:        "America/Merida",
				},
				{
					CountryCode: "MX",
					Name:        "America/Mexico_City",
				},
				{
					CountryCode: "MX",
					Name:        "America/Monterrey",
				},
				{
					CountryCode: "MX",
					Name:        "America/Ojinaga",
				},
				{
					CountryCode: "MX",
					Name:        "America/Tijuana",
				},
			},
		},
		{
			Code: "FM",
			Name: "Micronesia, Federated States of",
			Zones: []Zone{
				{
					CountryCode: "FM",
					Name:        "Pacific/Chuuk",
				},
				{
					CountryCode: "FM",
					Name:        "Pacific/Kosrae",
				},
				{
					CountryCode: "FM",
					Name:        "Pacific/Pohnpei",
				},
			},
		},
		{
			Code: "MD",
			Name: "Moldova, Republic of",
			Zones: []Zone{
				{
					CountryCode: "MD",
					Name:        "Europe/Chisinau",
				},
			},
		},
		{
			Code: "MC",
			Name: "Monaco",
			Zones: []Zone{
				{
					CountryCode: "MC",
					Name:        "Europe/Monaco",
				},
			},
		},
		{
			Code: "MN",
			Name: "Mongolia",
			Zones: []Zone{
				{
					CountryCode: "MN",
					Name:        "Asia/Choibalsan",
				},
				{
					CountryCode: "MN",
					Name:        "Asia/Hovd",
				},
				{
					CountryCode: "MN",
					Name:        "Asia/Ulaanbaatar",
				},
			},
		},
		{
			Code: "ME",
			Name: "Montenegro",
			Zones: []Zone{
				{
					CountryCode: "ME",
					Name:        "Europe/Podgorica",
				},
			},
		},
		{
			Code: "MS",
			Name: "Montserrat",
			Zones: []Zone{
				{
					CountryCode: "MS",
					Name:        "America/Montserrat",
				},
			},
		},
		{
			Code: "MA",
			Name: "Morocco",
			Zones: []Zone{
				{
					CountryCode: "MA",
					Name:        "Africa/Casablanca",
				},
			},
		},
		{
			Code: "MZ",
			Name: "Mozambique",
			Zones: []Zone{
				{
					CountryCode: "MZ",
					Name:        "Africa/Maputo",
				},
			},
		},
		{
			Code: "MM",
			Name: "Myanmar",
			Zones: []Zone{
				{
					CountryCode: "MM",
					Name:        "Asia/Yangon",
				},
			},
		},
		{
			Code: "NA",
			Name: "Namibia",
			Zones: []Zone{
				{
					CountryCode: "NA",
					Name:        "Africa/Windhoek",
				},
			},
		},
		{
			Code: "NR",
			Name: "Nauru",
			Zones: []Zone{
				{
					CountryCode: "NR",
					Name:        "Pacific/Nauru",
				},
			},
		},
		{
			Code: "NP",
			Name: "Nepal",
			Zones: []Zone{
				{
					CountryCode: "NP",
					Name:        "Asia/Kathmandu",
				},
			},
		},
		{
			Code: "NL",
			Name: "Netherlands",
			Zones: []Zone{
				{
					CountryCode: "NL",
					Name:        "Europe/Amsterdam",
				},
			},
		},
		{
			Code: "NC",
			Name: "New Caledonia",
			Zones: []Zone{
				{
					CountryCode: "NC",
					Name:        "Pacific/Noumea",
				},
			},
		},
		{
			Code: "NZ",
			Name: "New Zealand",
			Zones: []Zone{
				{
					CountryCode: "NZ",
					Name:        "Pacific/Auckland",
				},
				{
					CountryCode: "NZ",
					Name:        "Pacific/Chatham",
				},
			},
		},
		{
			Code: "NI",
			Name: "Nicaragua",
			Zones: []Zone{
				{
					CountryCode: "NI",
					Name:        "America/Managua",
				},
			},
		},
		{
			Code: "NE",
			Name: "Niger",
			Zones: []Zone{
				{
					CountryCode: "NE",
					Name:        "Africa/Niamey",
				},
			},
		},
		{
			Code: "NG",
			Name: "Nigeria",
			Zones: []Zone{
				{
					CountryCode: "NG",
					Name:        "Africa/Lagos",
				},
			},
		},
		{
			Code: "NU",
			Name: "Niue",
			Zones: []Zone{
				{
					CountryCode: "NU",
					Name:        "Pacific/Niue",
				},
			},
		},
		{
			Code: "NF",
			Name: "Norfolk Island",
			Zones: []Zone{
				{
					CountryCode: "NF",
					Name:        "Pacific/Norfolk",
				},
			},
		},
		{
			Code: "MP",
			Name: "Northern Mariana Islands",
			Zones: []Zone{
				{
					CountryCode: "MP",
					Name:        "Pacific/Saipan",
				},
			},
		},
		{
			Code: "NO",
			Name: "Norway",
			Zones: []Zone{
				{
					CountryCode: "NO",
					Name:        "Europe/Oslo",
				},
			},
		},
		{
			Code: "OM",
			Name: "Oman",
			Zones: []Zone{
				{
					CountryCode: "OM",
					Name:        "Asia/Muscat",
				},
			},
		},
		{
			Code: "PK",
			Name: "Pakistan",
			Zones: []Zone{
				{
					CountryCode: "PK",
					Name:        "Asia/Karachi",
				},
			},
		},
		{
			Code: "PW",
			Name: "Palau",
			Zones: []Zone{
				{
					CountryCode: "PW",
					Name:        "Pacific/Palau",
				},
			},
		},
		{
			Code: "PS",
			Name: "Palestine, State of",
			Zones: []Zone{
				{
					CountryCode: "PS",
					Name:        "Asia/Gaza",
				},
				{
					CountryCode: "PS",
					Name:        "Asia/Hebron",
				},
			},
		},
		{
			Code: "PA",
			Name: "Panama",
			Zones: []Zone{
				{
					CountryCode: "PA",
					Name:        "America/Panama",
				},
			},
		},
		{
			Code: "PG",
			Name: "Papua New Guinea",
			Zones: []Zone{
				{
					CountryCode: "PG",
					Name:        "Pacific/Bougainville",
				},
				{
					CountryCode: "PG",
					Name:        "Pacific/Port_Moresby",
				},
			},
		},
		{
			Code: "PY",
			Name: "Paraguay",
			Zones: []Zone{
				{
					CountryCode: "PY",
					Name:        "America/Asuncion",
				},
			},
		},
		{
			Code: "PE",
			Name: "Peru",
			Zones: []Zone{
				{
					CountryCode: "PE",
					Name:        "America/Lima",
				},
			},
		},
		{
			Code: "PH",
			Name: "Philippines",
			Zones: []Zone{
				{
					CountryCode: "PH",
					Name:        "Asia/Manila",
				},
			},
		},
		{
			Code: "PN",
			Name: "Pitcairn",
			Zones: []Zone{
				{
					CountryCode: "PN",
					Name:        "Pacific/Pitcairn",
				},
			},
		},
		{
			Code: "PL",
			Name: "Poland",
			Zones: []Zone{
				{
					CountryCode: "PL",
					Name:        "Europe/Warsaw",
				},
			},
		},
		{
			Code: "PT",
			Name: "Portugal",
			Zones: []Zone{
				{
					CountryCode: "PT",
					Name:        "Atlantic/Azores",
				},
				{
					CountryCode: "PT",
					Name:        "Atlantic/Madeira",
				},
				{
					CountryCode: "PT",
					Name:        "Europe/Lisbon",
				},
			},
		},
		{
			Code: "PR",
			Name: "Puerto Rico",
			Zones: []Zone{
				{
					CountryCode: "PR",
					Name:        "America/Puerto_Rico",
				},
			},
		},
		{
			Code: "QA",
			Name: "Qatar",
			Zones: []Zone{
				{
					CountryCode: "QA",
					Name:        "Asia/Qatar",
				},
			},
		},
		{
			Code: "RO",
			Name: "Romania",
			Zones: []Zone{
				{
					CountryCode: "RO",
					Name:        "Europe/Bucharest",
				},
			},
		},
		{
			Code: "RU",
			Name: "Russian Federation",
			Zones: []Zone{
				{
					CountryCode: "RU",
					Name:        "Asia/Anadyr",
				},
				{
					CountryCode: "RU",
					Name:        "Asia/Barnaul",
				},
				{
					CountryCode: "RU",
					Name:        "Asia/Chita",
				},
				{
					CountryCode: "RU",
					Name:        "Asia/Irkutsk",
				},
				{
					CountryCode: "RU",
					Name:        "Asia/Kamchatka",
				},
				{
					CountryCode: "RU",
					Name:        "Asia/Khandyga",
				},
				{
					CountryCode: "RU",
					Name:        "Asia/Krasnoyarsk",
				},
				{
					CountryCode: "RU",
					Name:        "Asia/Magadan",
				},
				{
					CountryCode: "RU",
					Name:        "Asia/Novokuznetsk",
				},
				{
					CountryCode: "RU",
					Name:        "Asia/Novosibirsk",
				},
				{
					CountryCode: "RU",
					Name:        "Asia/Omsk",
				},
				{
					CountryCode: "RU",
					Name:        "Asia/Sakhalin",
				},
				{
					CountryCode: "RU",
					Name:        "Asia/Srednekolymsk",
				},
				{
					CountryCode: "RU",
					Name:        "Asia/Tomsk",
				},
				{
					CountryCode: "RU",
					Name:        "Asia/Ust-Nera",
				},
				{
					CountryCode: "RU",
					Name:        "Asia/Vladivostok",
				},
				{
					CountryCode: "RU",
					Name:        "Asia/Yakutsk",
				},
				{
					CountryCode: "RU",
					Name:        "Asia/Yekaterinburg",
				},
				{
					CountryCode: "RU",
					Name:        "Europe/Astrakhan",
				},
				{
					CountryCode: "RU",
					Name:        "Europe/Kaliningrad",
				},
				{
					CountryCode: "RU",
					Name:        "Europe/Kirov",
				},
				{
					CountryCode: "RU",
					Name:        "Europe/Moscow",
				},
				{
					CountryCode: "RU",
					Name:        "Europe/Samara",
				},
				{
					CountryCode: "RU",
					Name:        "Europe/Saratov",
				},
				{
					CountryCode: "RU",
					Name:        "Europe/Ulyanovsk",
				},
				{
					CountryCode: "RU",
					Name:        "Europe/Volgograd",
				},
			},
		},
		{
			Code: "RW",
			Name: "Rwanda",
			Zones: []Zone{
				{
					CountryCode: "RW",
					Name:        "Africa/Kigali",
				},
			},
		},
		{
			Code: "RE",
			Name: "Réunion",
			Zones: []Zone{
				{
					CountryCode: "RE",
					Name:        "Indian/Reunion",
				},
			},
		},
		{
			Code: "BL",
			Name: "Saint Barthélemy",
			Zones: []Zone{
				{
					CountryCode: "BL",
					Name:        "America/St_Barthelemy",
				},
			},
		},
		{
			Code: "SH",
			Name: "Saint Helena, Ascension and Tristan da Cunha",
			Zones: []Zone{
				{
					CountryCode: "SH",
					Name:        "Atlantic/St_Helena",
				},
			},
		},
		{
			Code: "KN",
			Name: "Saint Kitts and Nevis",
			Zones: []Zone{
				{
					CountryCode: "KN",
					Name:        "America/St_Kitts",
				},
			},
		},
		{
			Code: "LC",
			Name: "Saint Lucia",
			Zones: []Zone{
				{
					CountryCode: "LC",
					Name:        "America/St_Lucia",
				},
			},
		},
		{
			Code: "MF",
			Name: "Saint Martin (French part)",
			Zones: []Zone{
				{
					CountryCode: "MF",
					Name:        "America/Marigot",
				},
			},
		},
		{
			Code: "PM",
			Name: "Saint Pierre and Miquelon",
			Zones: []Zone{
				{
					CountryCode: "PM",
					Name:        "America/Miquelon",
				},
			},
		},
		{
			Code: "VC",
			Name: "Saint Vincent and the Grenadines",
			Zones: []Zone{
				{
					CountryCode: "VC",
					Name:        "America/St_Vincent",
				},
			},
		},
		{
			Code: "WS",
			Name: "Samoa",
			Zones: []Zone{
				{
					CountryCode: "WS",
					Name:        "Pacific/Apia",
				},
			},
		},
		{
			Code: "SM",
			Name: "San Marino",
			Zones: []Zone{
				{
					CountryCode: "SM",
					Name:        "Europe/San_Marino",
				},
			},
		},
		{
			Code: "ST",
			Name: "Sao Tome and Principe",
			Zones: []Zone{
				{
					CountryCode: "ST",
					Name:        "Africa/Sao_Tome",
				},
			},
		},
		{
			Code: "SA",
			Name: "Saudi Arabia",
			Zones: []Zone{
				{
					CountryCode: "SA",
					Name:        "Asia/Riyadh",
				},
			},
		},
		{
			Code: "SN",
			Name: "Senegal",
			Zones: []Zone{
				{
					CountryCode: "SN",
					Name:        "Africa/Dakar",
				},
			},
		},
		{
			Code: "RS",
			Name: "Serbia",
			Zones: []Zone{
				{
					CountryCode: "RS",
					Name:        "Europe/Belgrade",
				},
			},
		},
		{
			Code: "SC",
			Name: "Seychelles",
			Zones: []Zone{
				{
					CountryCode: "SC",
					Name:        "Indian/Mahe",
				},
			},
		},
		{
			Code: "SL",
			Name: "Sierra Leone",
			Zones: []Zone{
				{
					CountryCode: "SL",
					Name:        "Africa/Freetown",
				},
			},
		},
		{
			Code: "SG",
			Name: "Singapore",
			Zones: []Zone{
				{
					CountryCode: "SG",
					Name:        "Asia/Singapore",
				},
			},
		},
		{
			Code: "SX",
			Name: "Sint Maarten (Dutch part)",
			Zones: []Zone{
				{
					CountryCode: "SX",
					Name:        "America/Lower_Princes",
				},
			},
		},
		{
			Code: "SK",
			Name: "Slovakia",
			Zones: []Zone{
				{
					CountryCode: "SK",
					Name:        "Europe/Bratislava",
				},
			},
		},
		{
			Code: "SI",
			Name: "Slovenia",
			Zones: []Zone{
				{
					CountryCode: "SI",
					Name:        "Europe/Ljubljana",
				},
			},
		},
		{
			Code: "SB",
			Name: "Solomon Islands",
			Zones: []Zone{
				{
					CountryCode: "SB",
					Name:        "Pacific/Guadalcanal",
				},
			},
		},
		{
			Code: "SO",
			Name: "Somalia",
			Zones: []Zone{
				{
					CountryCode: "SO",
					Name:        "Africa/Mogadishu",
				},
			},
		},
		{
			Code: "ZA",
			Name: "South Africa",
			Zones: []Zone{
				{
					CountryCode: "ZA",
					Name:        "Africa/Johannesburg",
				},
			},
		},
		{
			Code: "GS",
			Name: "South Georgia and the South Sandwich Islands",
			Zones: []Zone{
				{
					CountryCode: "GS",
					Name:        "Atlantic/South_Georgia",
				},
			},
		},
		{
			Code: "SS",
			Name: "South Sudan",
			Zones: []Zone{
				{
					CountryCode: "SS",
					Name:        "Africa/Juba",
				},
			},
		},
		{
			Code: "ES",
			Name: "Spain",
			Zones: []Zone{
				{
					CountryCode: "ES",
					Name:        "Africa/Ceuta",
				},
				{
					CountryCode: "ES",
					Name:        "Atlantic/Canary",
				},
				{
					CountryCode: "ES",
					Name:        "Europe/Madrid",
				},
			},
		},
		{
			Code: "LK",
			Name: "Sri Lanka",
			Zones: []Zone{
				{
					CountryCode: "LK",
					Name:        "Asia/Colombo",
				},
			},
		},
		{
			Code: "SD",
			Name: "Sudan",
			Zones: []Zone{
				{
					CountryCode: "SD",
					Name:        "Africa/Khartoum",
				},
			},
		},
		{
			Code: "SR",
			Name: "Suriname",
			Zones: []Zone{
				{
					CountryCode: "SR",
					Name:        "America/Paramaribo",
				},
			},
		},
		{
			Code: "SJ",
			Name: "Svalbard and Jan Mayen",
			Zones: []Zone{
				{
					CountryCode: "SJ",
					Name:        "Arctic/Longyearbyen",
				},
			},
		},
		{
			Code: "SZ",
			Name: "Swaziland",
			Zones: []Zone{
				{
					CountryCode: "SZ",
					Name:        "Africa/Mbabane",
				},
			},
		},
		{
			Code: "SE",
			Name: "Sweden",
			Zones: []Zone{
				{
					CountryCode: "SE",
					Name:        "Europe/Stockholm",
				},
			},
		},
		{
			Code: "CH",
			Name: "Switzerland",
			Zones: []Zone{
				{
					CountryCode: "CH",
					Name:        "Europe/Zurich",
				},
			},
		},
		{
			Code: "SY",
			Name: "Syrian Arab Republic",
			Zones: []Zone{
				{
					CountryCode: "SY",
					Name:        "Asia/Damascus",
				},
			},
		},
		{
			Code: "TW",
			Name: "Taiwan, Province of China",
			Zones: []Zone{
				{
					CountryCode: "TW",
					Name:        "Asia/Taipei",
				},
			},
		},
		{
			Code: "TJ",
			Name: "Tajikistan",
			Zones: []Zone{
				{
					CountryCode: "TJ",
					Name:        "Asia/Dushanbe",
				},
			},
		},
		{
			Code: "TZ",
			Name: "Tanzania, United Republic of",
			Zones: []Zone{
				{
					CountryCode: "TZ",
					Name:        "Africa/Dar_es_Salaam",
				},
			},
		},
		{
			Code: "TH",
			Name: "Thailand",
			Zones: []Zone{
				{
					CountryCode: "TH",
					Name:        "Asia/Bangkok",
				},
			},
		},
		{
			Code: "TL",
			Name: "Timor-Leste",
			Zones: []Zone{
				{
					CountryCode: "TL",
					Name:        "Asia/Dili",
				},
			},
		},
		{
			Code: "TG",
			Name: "Togo",
			Zones: []Zone{
				{
					CountryCode: "TG",
					Name:        "Africa/Lome",
				},
			},
		},
		{
			Code: "TK",
			Name: "Tokelau",
			Zones: []Zone{
				{
					CountryCode: "TK",
					Name:        "Pacific/Fakaofo",
				},
			},
		},
		{
			Code: "TO",
			Name: "Tonga",
			Zones: []Zone{
				{
					CountryCode: "TO",
					Name:        "Pacific/Tongatapu",
				},
			},
		},
		{
			Code: "TT",
			Name: "Trinidad and Tobago",
			Zones: []Zone{
				{
					CountryCode: "TT",
					Name:        "America/Port_of_Spain",
				},
			},
		},
		{
			Code: "TN",
			Name: "Tunisia",
			Zones: []Zone{
				{
					CountryCode: "TN",
					Name:        "Africa/Tunis",
				},
			},
		},
		{
			Code: "TR",
			Name: "Turkey",
			Zones: []Zone{
				{
					CountryCode: "TR",
					Name:        "Europe/Istanbul",
				},
			},
		},
		{
			Code: "TM",
			Name: "Turkmenistan",
			Zones: []Zone{
				{
					CountryCode: "TM",
					Name:        "Asia/Ashgabat",
				},
			},
		},
		{
			Code: "TC",
			Name: "Turks and Caicos Islands",
			Zones: []Zone{
				{
					CountryCode: "TC",
					Name:        "America/Grand_Turk",
				},
			},
		},
		{
			Code: "TV",
			Name: "Tuvalu",
			Zones: []Zone{
				{
					CountryCode: "TV",
					Name:        "Pacific/Funafuti",
				},
			},
		},
		{
			Code: "UG",
			Name: "Uganda",
			Zones: []Zone{
				{
					CountryCode: "UG",
					Name:        "Africa/Kampala",
				},
			},
		},
		{
			Code: "UA",
			Name: "Ukraine",
			Zones: []Zone{
				{
					CountryCode: "UA",
					Name:        "Europe/Kiev",
				},
				{
					CountryCode: "UA",
					Name:        "Europe/Simferopol",
				},
				{
					CountryCode: "UA",
					Name:        "Europe/Uzhgorod",
				},
				{
					CountryCode: "UA",
					Name:        "Europe/Zaporozhye",
				},
			},
		},
		{
			Code: "AE",
			Name: "United Arab Emirates",
			Zones: []Zone{
				{
					CountryCode: "AE",
					Name:        "Asia/Dubai",
				},
			},
		},
		{
			Code: "GB",
			Name: "United Kingdom",
			Zones: []Zone{
				{
					CountryCode: "GB",
					Name:        "Europe/London",
				},
			},
		},
		{
			Code: "US",
			Name: "United States",
			Zones: []Zone{
				{
					CountryCode: "US",
					Name:        "America/Adak",
				},
				{
					CountryCode: "US",
					Name:        "America/Anchorage",
				},
				{
					CountryCode: "US",
					Name:        "America/Boise",
				},
				{
					CountryCode: "US",
					Name:        "America/Chicago",
				},
				{
					CountryCode: "US",
					Name:        "America/Denver",
				},
				{
					CountryCode: "US",
					Name:        "America/Detroit",
				},
				{
					CountryCode: "US",
					Name:        "America/Indiana/Indianapolis",
				},
				{
					CountryCode: "US",
					Name:        "America/Indiana/Knox",
				},
				{
					CountryCode: "US",
					Name:        "America/Indiana/Marengo",
				},
				{
					CountryCode: "US",
					Name:        "America/Indiana/Petersburg",
				},
				{
					CountryCode: "US",
					Name:        "America/Indiana/Tell_City",
				},
				{
					CountryCode: "US",
					Name:        "America/Indiana/Vevay",
				},
				{
					CountryCode: "US",
					Name:        "America/Indiana/Vincennes",
				},
				{
					CountryCode: "US",
					Name:        "America/Indiana/Winamac",
				},
				{
					CountryCode: "US",
					Name:        "America/Juneau",
				},
				{
					CountryCode: "US",
					Name:        "America/Kentucky/Louisville",
				},
				{
					CountryCode: "US",
					Name:        "America/Kentucky/Monticello",
				},
				{
					CountryCode: "US",
					Name:        "America/Los_Angeles",
				},
				{
					CountryCode: "US",
					Name:        "America/Menominee",
				},
				{
					CountryCode: "US",
					Name:        "America/Metlakatla",
				},
				{
					CountryCode: "US",
					Name:        "America/New_York",
				},
				{
					CountryCode: "US",
					Name:        "America/Nome",
				},
				{
					CountryCode: "US",
					Name:        "America/North_Dakota/Beulah",
				},
				{
					CountryCode: "US",
					Name:        "America/North_Dakota/Center",
				},
				{
					CountryCode: "US",
					Name:        "America/North_Dakota/New_Salem",
				},
				{
					CountryCode: "US",
					Name:        "America/Phoenix",
				},
				{
					CountryCode: "US",
					Name:        "America/Sitka",
				},
				{
					CountryCode: "US",
					Name:        "America/Yakutat",
				},
				{
					CountryCode: "US",
					Name:        "Pacific/Honolulu",
				},
			},
		},
		{
			Code: "UM",
			Name: "United States Minor Outlying Islands",
			Zones: []Zone{
				{
					CountryCode: "UM",
					Name:        "Pacific/Midway",
				},
				{
					CountryCode: "UM",
					Name:        "Pacific/Wake",
				},
			},
		},
		{
			Code: "UY",
			Name: "Uruguay",
			Zones: []Zone{
				{
					CountryCode: "UY",
					Name:        "America/Montevideo",
				},
			},
		},
		{
			Code: "UZ",
			Name: "Uzbekistan",
			Zones: []Zone{
				{
					CountryCode: "UZ",
					Name:        "Asia/Samarkand",
				},
				{
					CountryCode: "UZ",
					Name:        "Asia/Tashkent",
				},
			},
		},
		{
			Code: "VU",
			Name: "Vanuatu",
			Zones: []Zone{
				{
					CountryCode: "VU",
					Name:        "Pacific/Efate",
				},
			},
		},
		{
			Code: "VE",
			Name: "Venezuela, Bolivarian Republic of",
			Zones: []Zone{
				{
					CountryCode: "VE",
					Name:        "America/Caracas",
				},
			},
		},
		{
			Code: "VN",
			Name: "Viet Nam",
			Zones: []Zone{
				{
					CountryCode: "VN",
					Name:        "Asia/Ho_Chi_Minh",
				},
			},
		},
		{
			Code: "VG",
			Name: "Virgin Islands, British",
			Zones: []Zone{
				{
					CountryCode: "VG",
					Name:        "America/Tortola",
				},
			},
		},
		{
			Code: "VI",
			Name: "Virgin Islands, U.S.",
			Zones: []Zone{
				{
					CountryCode: "VI",
					Name:        "America/St_Thomas",
				},
			},
		},
		{
			Code: "WF",
			Name: "Wallis and Futuna",
			Zones: []Zone{
				{
					CountryCode: "WF",
					Name:        "Pacific/Wallis",
				},
			},
		},
		{
			Code: "EH",
			Name: "Western Sahara",
			Zones: []Zone{
				{
					CountryCode: "EH",
					Name:        "Africa/El_Aaiun",
				},
			},
		},
		{
			Code: "YE",
			Name: "Yemen",
			Zones: []Zone{
				{
					CountryCode: "YE",
					Name:        "Asia/Aden",
				},
			},
		},
		{
			Code: "ZM",
			Name: "Zambia",
			Zones: []Zone{
				{
					CountryCode: "ZM",
					Name:        "Africa/Lusaka",
				},
			},
		},
		{
			Code: "ZW",
			Name: "Zimbabwe",
			Zones: []Zone{
				{
					CountryCode: "ZW",
					Name:        "Africa/Harare",
				},
			},
		},
		{
			Code: "AX",
			Name: "Åland Islands",
			Zones: []Zone{
				{
					CountryCode: "AX",
					Name:        "Europe/Mariehamn",
				},
			},
		},
	}
)

func init() {
	// load + index countries into map
	// for below functions.

	once.Do(func() {
		mapped = make(map[string]Country)

		for i := 0; i < len(countries); i++ {
			mapped[countries[i].Code] = countries[i]
		}
	})
}

// GetCountries returns an array of all countries.
// Most common use: for loading into a country dropdown
// in HTML.
func GetCountries() []Country {
	return countries
}

// GetCountry returns a single Country that matches the country
// code passed and whether it was found
func GetCountry(code string) (c Country, found bool) {
	c, found = mapped[code]
	return
}
