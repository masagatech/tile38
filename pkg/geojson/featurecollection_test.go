package geojson

import (
	"testing"
)

func TestFeatureCollection(t *testing.T) {
	testJSON(t, `{

    "type": "FeatureCollection",
    "features": [
        {
            "type": "Feature",
            "geometry": {
                "type": "Point","coordinates": [-80.87088507656375,35.21515162500578]
            },
            "properties": {
                "name": "ABBOTT NEIGHBORHOOD PARK","address": "1300  SPRUCE ST"
            }
        },
        {
            "type": "Feature",
            "geometry": {
                "type": "Point","coordinates": [-80.83775386582222,35.24980190252168]
            },
            "properties": {
                "name": "DOUBLE OAKS CENTER","address": "1326 WOODWARD AV"
            }
        },
        {
            "type": "Feature",
            "geometry": {
                "type": "Point","coordinates": [-80.83827000459532,35.25674709224663]
            },
            "properties": {
                "name": "DOUBLE OAKS NEIGHBORHOOD PARK","address": "2605  DOUBLE OAKS RD"
            }
        },
        {
            "type": "Feature",
            "geometry": {
                "type": "Point","coordinates": [-80.83697759172735,35.25751734669229]
            },
            "properties": {
                "name": "DOUBLE OAKS POOL","address": "1200 NEWLAND RD"
            }
        },
        {
            "type": "Feature",
            "geometry": {
                "type": "Point","coordinates": [-80.81647652154736,35.40148708491418]
            },
            "properties": {
                "name": "DAVID B. WAYMER FLYING REGIONAL PARK","address": "15401 HOLBROOKS RD"
            }
        },
        {
            "type": "Feature",
            "geometry": {
                "type": "Point","coordinates": [-80.83556459443902,35.39917224760999]
            },
            "properties": {
                "name": "DAVID B. WAYMER COMMUNITY PARK","address": "302 HOLBROOKS RD"
            }
        },
        {
            "type": "Feature",
            "geometry": {
                "type": "Polygon","coordinates": [
                    [
                        [-80.72487831115721,35.26545403190955],
                        [-80.72135925292969,35.26727607954368],
                        [-80.71517944335938,35.26769654625573],
                        [-80.7125186920166,35.27035945142482],
                        [-80.70857048034668,35.268257165144064],
                        [-80.70479393005371,35.268397319259996],
                        [-80.70324897766113,35.26503355355979],
                        [-80.71088790893555,35.2553619492954],
                        [-80.71681022644043,35.2553619492954],
                        [-80.7150936126709,35.26054831539319],
                        [-80.71869850158691,35.26026797976481],
                        [-80.72032928466797,35.26061839914875],
                        [-80.72264671325684,35.26033806376283],
                        [-80.72487831115721,35.26545403190955]
                    ]
                ]
            },
            "properties": {
                "name": "Plaza Road Park"
            }
        }
    ]
}`)
}

func TestPointBounding(t *testing.T) {
	featureCollectionJSON := `{"type":"FeatureCollection","features":[{"type":"Feature","geometry":{"type":"Point","coordinates":[-75.1,56.283333]}}]}`
	polyFeatureJSON := `{"type": "FeatureCollection","features": [{"type": "Feature","properties": {},"geometry": {"type": "Polygon","coordinates": [[
		[-112.8515625,-29.535229562948444],
		[85.4296875,-2.535229562948444],
		[85.4296875,65.36683689226321],
		[-112.8515625,65.36683689226321],
		[-112.8515625,-29.535229562948444]
	]]}}]}`

	featureCollection := testJSON(t, featureCollectionJSON).(FeatureCollection)
	poly := testJSON(t, polyFeatureJSON).(FeatureCollection)

	r1 := featureCollection.Within(poly)
	r2 := featureCollection.Intersects(poly)
	if r1 != r2 || !r1 {
		t.Fatalf("expected %v/%v, got %v/%v", true, true, r1, r2)
	}
}