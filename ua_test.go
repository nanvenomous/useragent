package useragent_test

import (
	"fmt"
	"strings"
	"testing"

	ua "github.com/mileusna/useragent"
)

var testTable = [][]string{
	// useragent, name, version, mobile, os
	// Mac
	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/603.3.8 (KHTML, like Gecko) Version/10.1.2 Safari/603.3.8", ua.Safari, "10.1.2", "desktop", "macOS"},
	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.90 Safari/537.36", ua.Chrome, "60.0.3112.90", "desktop", "macOS"},
	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.12; rv:54.0) Gecko/20100101 Firefox/54.0", ua.Firefox, "54.0", "desktop", "macOS"},
	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36 OPR/46.0.2597.57", ua.Opera, "46.0.2597.57", "desktop", "macOS"},
	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.91 Safari/537.36 Vivaldi/1.92.917.39", "Vivaldi", "1.92.917.39", "desktop", "macOS"},
	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36 Edg/79.0.309.71", "Edge", "79.0.309.71", "desktop", "macOS"},

	// Windows
	{"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36", ua.Chrome, "59.0.3071.115", "desktop", "Windows"},
	{"Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.1; WOW64; Trident/4.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; .NET4.0C; .NET4.0E; InfoPath.2; GWX:RED)", ua.InternetExplorer, "8.0", "desktop", "Windows"},
	{"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1; .NET CLR 1.1.4322) NS8/0.9.6", ua.InternetExplorer, "6.0", "desktop", "Windows"},
	{"Mozilla/5.0 (Windows NT 10.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/52.0.2743.116 Safari/537.36 Edge/15.15063", ua.Edge, "15.15063", "desktop", "Windows"},

	// iPhone
	{"Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_2 like Mac OS X) AppleWebKit/603.2.4 (KHTML, like Gecko) Version/10.0 Mobile/14F89 Safari/602.1", ua.Safari, "10.0", "mobile", "iOS"},
	{"Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_2 like Mac OS X) AppleWebKit/603.1.30 (KHTML, like Gecko) CriOS/60.0.3112.89 Mobile/14F89 Safari/602.1", ua.Chrome, "60.0.3112.89", "mobile", "iOS"},
	{"Mozilla/5.0 (iPhone; CPU iPhone OS 9_3 like Mac OS X) AppleWebKit/601.1.46 (KHTML, like Gecko) OPiOS/14.0.0.104835 Mobile/13E233 Safari/9537.53", ua.Opera, "14.0.0.104835", "mobile", "iOS"},
	{"Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_2 like Mac OS X) AppleWebKit/603.2.4 (KHTML, like Gecko) FxiOS/8.1.1b4948 Mobile/14F89 Safari/603.2.4", ua.Firefox, "8.1.1b4948", "mobile", "iOS"},
	{"Mozilla/5.0 (iPhone; CPU iPhone OS 13_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0 EdgiOS/44.11.15 Mobile/15E148 Safari/605.1.15", ua.Edge, "44.11.15", "mobile", "iOS"},

	// iPad
	{"Mozilla/5.0 (iPad; CPU OS 10_3_2 like Mac OS X) AppleWebKit/603.2.4 (KHTML, like Gecko) Version/10.0 Mobile/14F89 Safari/602.1", ua.Safari, "10.0", "tablet", "iOS"},
	{"Mozilla/5.0 (iPad; CPU OS 10_3_2 like Mac OS X) AppleWebKit/602.1.50 (KHTML, like Gecko) CriOS/58.0.3029.113 Mobile/14F89 Safari/602.1", ua.Chrome, "58.0.3029.113", "tablet", "iOS"},
	{"Mozilla/5.0 (iPad; CPU OS 10_3_2 like Mac OS X) AppleWebKit/603.2.4 (KHTML, like Gecko) FxiOS/8.1.1b4948 Mobile/14F89 Safari/603.2.4", ua.Firefox, "8.1.1b4948", "tablet", "iOS"},

	// Andorid
	{"Mozilla/5.0 (Linux; Android 4.3; GT-I9300 Build/JSS15J) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.125 Mobile Safari/537.36", ua.Chrome, "59.0.3071.125", "mobile", "Android"},
	{"Mozilla/5.0 (Android 4.3; Mobile; rv:54.0) Gecko/54.0 Firefox/54.0", ua.Firefox, "54.0", "mobile", "Android"},
	{"Mozilla/5.0 (Linux; Android 4.3; GT-I9300 Build/JSS15J) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/55.0.2883.91 Mobile Safari/537.36 OPR/42.9.2246.119956", ua.Opera, "42.9.2246.119956", "mobile", ua.Android},
	{"Opera/9.80 (Android; Opera Mini/28.0.2254/66.318; U; en) Presto/2.12.423 Version/12.16", ua.OperaMini, "28.0.2254/66.318", "mobile", "Android"},
	{"Mozilla/5.0 (Linux; U; Android 4.3; en-us; GT-I9300 Build/JSS15J) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30", "Android browser", "4.0", "mobile", "Android"},
	{"Mozilla/5.0 (Linux; Android 10; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.0 Mobile Safari/537.36 EdgA/44.11.4.4140", ua.Edge, "44.11.4.4140", "mobile", "Android"},

	{"Mozilla/5.0 (Linux; Android 6.0.1; SAMSUNG SM-A310F/A310FXXU2BQB1 Build/MMB29K) AppleWebKit/537.36 (KHTML, like Gecko) SamsungBrowser/5.4 Chrome/51.0.2704.106 Mobile Safari/537.36", "Samsung Browser", "5.4", "mobile", "Android"},
	{"Mozilla/5.0 (Linux; Android 9; LM-Q630) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Mobile Safari/537.36", ua.Chrome, "86.0.4240.198", "mobile", "Android"},
	{"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/534.24 (KHTML, like Gecko) Chrome/79.0.3945.147 Safari/534.24 XiaoMi/MiuiBrowser/12.11.5-gn", "Miui Browser", "12.11.5-gn", "mobile", ua.Linux},
	{"Mozilla/5.0 (Linux; U; Android 11; ru-ru; Redmi Note 10S Build/RP1A.200720.011) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/89.0.4389.116 Mobile Safari/537.36 XiaoMi/MiuiBrowser/12.13.2-gn", "Miui Browser", "12.13.2-gn", "mobile", ua.Android},

	{"Mozilla/5.0 (Linux; Android 10; MED-LX9N; HMSCore 6.6.0.311) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.105 HuaweiBrowser/12.1.0.303 Mobile Safari/537.36", "Huawei Browser", "12.1.0.303", "mobile", "Android"},

	// useragent, name, version, mobile, os
	{"Mozilla/5.0 (Linux; Android 9; ONEPLUS A6003) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.99 Mobile Safari/537.36", ua.Chrome, "71.0.3578.99", "mobile", ua.Android},
	{"Mozilla/5.0 (Android 9; Mobile; rv:64.0) Gecko/64.0 Firefox/64.0", ua.Firefox, "64.0", "mobile", ua.Android},
	{"Opera/9.80 (Android; Opera Mini/38.0.2254/128.54; U; en) Presto/2.12.423 Version/12.16", ua.OperaMini, "38.0.2254/128.54", "mobile", ua.Android},
	{"Mozilla/5.0 (Linux; Android 9; ONEPLUS A6003 Build/PKQ1.180716.001) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.110 Mobile Safari/537.36 OPR/49.2.2361.134358", ua.Opera, "49.2.2361.134358", "mobile", ua.Android},
	{"Mozilla/5.0 (Linux; Android 9; ONEPLUS A6003 Build/PKQ1.180716.001) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.86 Mobile Safari/537.36 EdgA/42.0.92.2864", ua.Edge, "42.0.92.2864", "mobile", ua.Android},
	{"Mozilla/5.0 (Linux; Android 9; ONEPLUS A6003 Build/PKQ1.180716.001) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/71.0.3578.99 Mobile Safari/537.36 OPT/1.14.51", ua.OperaTouch, "1.14.51", "mobile", ua.Android},
	{"Mozilla/5.0 (Linux; Android 7.0; Moto G (4)) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4143.7 Mobile Safari/537.36 Chrome-Lighthouse", ua.Chrome, "84.0.4143.7", "mobile", ua.Android},
	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36", ua.Chrome, "87.0.4280.88", "desktop", ua.MacOS}, // Lighthouse
	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4143.7 Safari/537.36 Chrome-Lighthouse", ua.Chrome, "84.0.4143.7", "desktop", ua.MacOS},
	{"Mozilla/5.0 (Linux; Android 7.0; Moto G (4)) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4143.7 Mobile Safari/537.36 Chrome-Lighthouse", ua.Chrome, "84.0.4143.7", "mobile", ua.Android},

	// Windows phone
	{"Mozilla/4.0 (compatible; MSIE 7.0; Windows Phone OS 7.0; Trident/3.1; IEMobile/7.0; NOKIA; Lumia 630)", ua.InternetExplorer, "7.0", "mobile", ua.WindowsPhone},

	// Bots
	{"Mozilla/5.0 (Linux; Android 6.0.1; Nexus 5X Build/MMB29P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2272.96 Mobile Safari/537.36 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)", ua.Googlebot, "2.1", "mobile", "Android"},
	{"Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)", ua.Googlebot, "2.1", "bot", ""},
	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_5) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.1.1 Safari/605.1.15 (Applebot/0.1; +http://www.apple.com/go/applebot)", "Applebot", "0.1", "bot", ""},
	{"Twitterbot/1.0", ua.Twitterbot, "1.0", ua.Applebot, ""},
	{"facebookexternalhit/1.1", ua.FacebookExternalHit, "1.1", "bot", ""},
	{"Mozilla/5.0 (compatible; SemrushBot/7~bl; +http://www.semrush.com/bot.html", "SemrushBot", "7~bl", "bot", ""},
	{"Mozilla/5.0 (compatible; YandexBot/3.0; +http://yandex.com/bots) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.268", "YandexBot", "3.0", "bot", ""},
	{"Mozilla/5.0 (compatible; Discordbot/2.0; +https://discordapp.com)", "Discordbot", "2.0", "bot", ""},
	{"Mozilla/5.0 (compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm)", "Bingbot", "2.0", "bot", ""},                                                                                                                                    // old binbot
	{"Mozilla/5.0 AppleWebKit/537.36 (KHTML, like Gecko; compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm) Chrome/100.0.0.0 Safari/537.36", "Bingbot", "2.0", "bot", ""},                                                               // new bingbot desktop
	{"Mozilla/5.0 (Linux; Android 6.0.1; Nexus 5X Build/MMB29P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.1.0.0 Mobile Safari/537.36 (compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm)", "Bingbot", "2.0", "bot", ua.Android}, // new bingbot mobile

	// Google ads bots
	{"Mozilla/5.0 (Linux; Android 4.0.0; Galaxy Nexus Build/IMM76B) AppleWebKit/537.36 (KHTML, like Gecko; Mediapartners-Google) Chrome/104.0.0.0 Mobile Safari/537.36", ua.GoogleAdsBot, "", "bot", ua.Android},
	{"Mozilla/5.0 (Linux; Android 5.0; SM-G920A) AppleWebKit (KHTML, like Gecko) Chrome Mobile Safari (compatible; AdsBot-Google-Mobile; +http://www.google.com/mobile/adsbot.html)", ua.GoogleAdsBot, "", "bot", ua.Android},
	{"Mozilla/5.0 (iPhone; CPU iPhone OS 14_7_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.1.2 Mobile/15E148 Safari/604.1 (compatible; AdsBot-Google-Mobile; +http://www.google.com/mobile/adsbot.html)", ua.GoogleAdsBot, "", "bot", ua.IOS},
	{"Mozilla/5.0 (iPhone; U; CPU iPhone OS 10_0 like Mac OS X; en-us) AppleWebKit/602.1.38 (KHTML, like Gecko) Version/10.0 Mobile/14A5297c Safari/602.1 (compatible; Mediapartners-Google/2.1; +http://www.google.com/bot.html)", ua.GoogleAdsBot, "", "bot", ua.IOS},
	// Brave
	{"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Brave Chrome/87.0.4280.101 Safari/537.36", ua.Chrome, "87.0.4280.101", "desktop", ua.Linux},
	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Safari/537.36", ua.Chrome, "87.0.4280.141", "desktop", ua.MacOS},

	// HeadlessChrome
	{"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) HeadlessChrome/98.0.4758.0 Safari/537.36", ua.HeadlessChrome, "98.0.4758.0", "desktop", ua.Linux},

	//FB App
	{"Mozilla/5.0 (iPhone; CPU iPhone OS 15_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/19E258 [FBAN/FBIOS;FBDV/iPhone8,2;FBMD/iPhone;FBSN/iOS;FBSV/15.4.1;FBSS/3;FBID/phone;FBLC/fr_FR;FBOP/5]", ua.FacebookApp, "FBIOS", "mobile", ua.IOS},
	{"Mozilla/5.0 (Linux; Android 13; SM-T220 Build/TP1A.220624.014; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/109.0.5414.117 Safari/537.36 [FB_IAB/FB4A;FBAV/400.0.0.37.76;]", ua.FacebookApp, "400.0.0.37.76", "", ua.Android},

	//Instagram
	{"Mozilla/5.0 (iPhone; CPU iPhone OS 16_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 Instagram 270.0.0.13.83 (iPhone13,2; iOS 16_3; es_ES; es-ES; scale=3.00; 1170x2532; 445843881) NW/1", ua.InstagramApp, "270.0.0.13.83", "mobile", ua.IOS},

	// Tiktok
	{"Mozilla/5.0 (iPhone; CPU iPhone OS 15_5 like Mac OS ) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 musical_ly_28.2.0 JsSdk/2.0 NetType/WIFI Channel/App Store ByteLocale/es Region/PE RevealType/Dialog isDarkMode/0 WKWebView/1 BytedanceWebview/d8a21c6 FalconTag/D6EBBF89-6D75-4BBD-9304-BF199C6B4DB1", ua.TiktokApp, "", "mobile", ua.IOS},
	{"Mozilla/5.0 (Linux; Android 10; AGS3K-W09 Build/HUAWEIAGS3K-W09; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/88.0.4324.93 Safari/537.36 trill_2022803040 JsSdk/1.0 NetType/WIFI Channel/huaweiadsglobal_int AppName/musical_ly app_version/28.3.4 ByteLocale/es ByteFullLocale/es Region/PE BytedanceWebview/d8a21c6", ua.TiktokApp, "28.3.4", ua.Android},

	// other
	{"Mozilla/5.0 (X11; CrOS x86_64 14150.74.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.114 Safari/537.36", ua.Chrome, "94.0.4606.114", "desktop", ua.ChromeOS},
	{"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.87 Safari/537.36 Google (+https://developers.google.com/+/web/snippet/)", ua.Chrome, "56.0.2924.87", "bot", ua.Linux}, // Google+ fetch

	// tools
	{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_4) AppleWebKit/537.36 (KHTML, like Gecko) QtWebEngine/5.6.0 Chrome/45.0.2454.101 Safari/537.36", "QtWebEngine", "5.6.0", "", "macOS"},
	{"Go-http-client/1.1", "Go-http-client", "1.1", "", ""},
	{"Wget/1.12 (linux-gnu)", "Wget", "1.12", "", ""},
	{"Wget/1.17.1 (darwin15.2.0)", "Wget", "1.17.1", "", ""},

	// unstandard stuff
	{"BUbiNG (+http://law.di.unimi.it/BUbiNG.html)", "BUbiNG", "", "", ""},
	//{"Aweme 8.2.0 rv:82017 (iPhone6,2; iOS 12.4; zh_CN) Cronet", "Aweme", "", "", ""},

	//GooglePlus   "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.87 Safari/537.36 Google (+https://developers.google.com/+/web/snippet/)"
	//Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_1) AppleWebKit/600.2.5 (KHTML, like Gecko) Version/8.0.2 Safari/600.2.5 (Applebot/0.1; +http://www.apple.com/go/applebot)
	//Mozilla/5.0 (Macintosh; Intel Mac OS Xt 10_11_4) AppleWebKit/537.36 (KHTML, like Gecko) QtWebEngine/5.6.0 Chrome/45.0.2454.101 Safari/537.36

}

func TestParse(t *testing.T) {
	for _, test := range testTable {
		ua := ua.Parse(test[0])
		if ua.Name != test[1] {
			t.Error("\n", test[0], "\nName should be", test[1], "not", ua.Name)
		}
		if ua.Version != test[2] {
			t.Error("\nVersion should be", test[2], "not", ua.Version)
		}

		if len(test) > 3 {
			if test[3] == "desktop" && ua.Mobile {
				t.Error("\n", ua.String, "should be desktop type not mobile")
			}

			if test[3] == "mobile" && !ua.Mobile {
				t.Error("\n", ua.String, "should be mobile")
			}
			if test[3] == "tablet" && !ua.Tablet {
				t.Error("\n", ua.String, "should be tablet")
			}
		}

		if len(test) > 4 && test[4] != ua.OS {
			t.Error("\n", test[0], "OS should", test[4], "not", ua.OS)
		}
		//fmt.Println(ua.OS, ua.OSVersion, ua.Device)

	}
}

var testUA ua.UserAgent

func BenchmarkUserAgent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range testTable {
			testUA = ua.Parse(test[0])
		}
	}
}

func ExampleParse() {
	userAgents := []string{
		// Mac
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/603.3.8 (KHTML, like Gecko) Version/10.1.2 Safari/603.3.8",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.90 Safari/537.36",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.12; rv:54.0) Gecko/20100101 Firefox/54.0",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36 OPR/46.0.2597.57",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.91 Safari/537.36 Vivaldi/1.92.917.39",

		// Windows
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36",
		"Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.1; WOW64; Trident/4.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; .NET4.0C; .NET4.0E; InfoPath.2; GWX:RED)",
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1; .NET CLR 1.1.4322) NS8/0.9.6",

		// iPhone
		"Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_2 like Mac OS X) AppleWebKit/603.2.4 (KHTML, like Gecko) Version/10.0 Mobile/14F89 Safari/602.1",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_2 like Mac OS X) AppleWebKit/603.1.30 (KHTML, like Gecko) CriOS/60.0.3112.89 Mobile/14F89 Safari/602.1",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 9_3 like Mac OS X) AppleWebKit/601.1.46 (KHTML, like Gecko) OPiOS/14.0.0.104835 Mobile/13E233 Safari/9537.53",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_2 like Mac OS X) AppleWebKit/603.2.4 (KHTML, like Gecko) FxiOS/8.1.1b4948 Mobile/14F89 Safari/603.2.4",

		// iPad
		"Mozilla/5.0 (iPad; CPU OS 10_3_2 like Mac OS X) AppleWebKit/603.2.4 (KHTML, like Gecko) Version/10.0 Mobile/14F89 Safari/602.1",
		"Mozilla/5.0 (iPad; CPU OS 10_3_2 like Mac OS X) AppleWebKit/602.1.50 (KHTML, like Gecko) CriOS/58.0.3029.113 Mobile/14F89 Safari/602.1",
		"Mozilla/5.0 (iPad; CPU OS 10_3_2 like Mac OS X) AppleWebKit/603.2.4 (KHTML, like Gecko) FxiOS/8.1.1b4948 Mobile/14F89 Safari/603.2.4",

		// Andorid
		"Mozilla/5.0 (Linux; Android 4.3; GT-I9300 Build/JSS15J) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.125 Mobile Safari/537.36",
		"Mozilla/5.0 (Android 4.3; Mobile; rv:54.0) Gecko/54.0 Firefox/54.0",
		"Mozilla/5.0 (Linux; Android 4.3; GT-I9300 Build/JSS15J) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/55.0.2883.91 Mobile Safari/537.36 OPR/42.9.2246.119956",
		"Opera/9.80 (Android; Opera Mini/28.0.2254/66.318; U; en) Presto/2.12.423 Version/12.16",
		"Mozilla/5.0 (Linux; U; Android 4.3; en-us; GT-I9300 Build/JSS15J) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",

		"Mozilla/5.0 (Linux; Android 6.0.1; SAMSUNG SM-A310F/A310FXXU2BQB1 Build/MMB29K) AppleWebKit/537.36 (KHTML, like Gecko) SamsungBrowser/5.4 Chrome/51.0.2704.106 Mobile Safari/537.36",
	}

	for _, s := range userAgents {
		ua := ua.Parse(s)
		fmt.Println()
		fmt.Println(ua.String)
		fmt.Println(strings.Repeat("=", len(ua.String)))
		fmt.Println("Name:", ua.Name, "v", ua.Version)
		fmt.Println("OS:", ua.OS, "v", ua.OSVersion)
		fmt.Println("Device:", ua.Device)
		if ua.Mobile {
			fmt.Println("(Mobile)")
		}
		if ua.Tablet {
			fmt.Println("(Tablet)")
		}
		if ua.Desktop {
			fmt.Println("(Desktop)")
		}
		if ua.Bot {
			fmt.Println("(Bot)")
		}
		if ua.URL != "" {
			fmt.Println(ua.URL)
		}

	}

}
