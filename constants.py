BASE_URL = 'https://www.otomoto.pl/'
TYPE = {
    'passenger': 'osobowe/',
    'delivery': 'dostawcze/',
    'motorcycle': 'motocykle-i-quady/',
    'truck': 'ciezarowe/',
    'construction': 'maszyny-budowlane/',
    'trailer': 'przyczepy/',
    'agro': 'maszyny-rolnicze/'
}

QUERY_STRING_MARK = '?';
QUERY_START = 'search%5B';
QUERY_JOIN = '&';

SORT_TYPE = {
    'time': 'search%5Border%5D=created_at%3A',
    'price': 'search%5Border%5D=filter_float_price%3A',
    'mileage': 'search%5Border%5D=filter_float_mileage%3A',
    'power' : 'search%5Border%5D=filter_float_engine_power%3A'
}

SORT_TYPE_MODE = {
    'asc': 'asc',
    'dsc': 'desc'
}

ENGINE_CAPACITY = {
    'from': 'search%5Bfilter_float_engine_capacity%3Afrom%5D=',
    'to': 'search%5Bfilter_float_engine_capacity%3Ato%5D='
}

YEAR_SINCE = 'od-'
YEAR_TO = 'search%5Bfilter_float_year%3Ato%5D='