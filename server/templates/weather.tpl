{{define "weather"}}<iframe id="forecast_embed" type="text/html" frameborder="0" height="245" width="100%" src="http://forecast.io/embed/#lat={{.Location.Lat}}&lon={{.Location.Lon}}&name={{.Location.CityName}}&units=ca" style="background:#999;opacity:0.5;margin-top:50px;">
</iframe>{{end}}
