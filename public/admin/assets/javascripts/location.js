!(function() {
  (function($, Export) {
    "use script";

    var Locator = function($context) {
      // jQuery object instances
      this.$context        = $context;
      this.$country        = $('[data-location-role = "country"]', $context);
      this.$city           = $('[data-location-role = "city"]', $context);
      this.$region         = $('[data-location-role = "region"]', $context);
      this.$address        = $('[data-location-role = "address"]', $context);
      this.$zip            = $('[data-location-role = "zip"]', $context);
      this.$latitude       = $('[data-location-role = "latitude"]', $context);
      this.$longitude      = $('[data-location-role = "longitude"]', $context);
      this.$geocode        = $('[data-location-role = "geocode"]', $context);
      this.$reverseGeocode = $('[data-location-role = "reverseGeocode"]', $context);
      this.$map            = $('[data-location-role = "map"]', $context);
      this.$currentAddress = $('[data-location-role = "currentAddress"]', $context);
      // Google map instances
      this.map               = null;
      this.geocoder          = null;
      this.marker            = null;

      // Main method
      this.tryToGetCurrentPosition(this.init);
    };

    Locator.prototype.init = function() {
      this.setupMap();
      this.bindEvents();
    };

    Locator.prototype.tryToGetCurrentPosition = function(alwaysCallback) {
      var _this = this,
        latitude = Number(this.$latitude.val()),
        longitude = Number(this.$longitude.val());

      if (navigator.geolocation && latitude === 0 && longitude === 0) {
        navigator.geolocation.getCurrentPosition(function(position) {
          _this.$latitude.val(position.coords.latitude);
          _this.$longitude.val(position.coords.longitude);
          alwaysCallback.call(_this);
        }, function() {
          alwaysCallback.call(_this);
        });
      } else {
          alwaysCallback.call(_this);
      };
    };

    Locator.prototype.setupMap = function() {
      var latitude = Number(this.$latitude.val()),
        longitude = Number(this.$longitude.val()),
        zoom = this.getZoom();

      this.geocoder = new google.maps.Geocoder;
      this.map = new google.maps.Map(this.$map[0], {
        zoom: zoom
      });
      this.marker = new google.maps.Marker({
        map: this.map,
        position: {lat: latitude, lng: longitude},
        draggable: true
      });
      this.map.setCenter(this.marker.getPosition());
      this.setCurrentAddress();
    };

    Locator.prototype.bindEvents = function() {
      var _this = this;

      google.maps.event.bind(this.marker, "dragend", this, function(){
        _this.setPosition();
        _this.setCurrentAddress();
      });

      this.$geocode.on('click', function(evt) {
        evt.preventDefault();
        _this.geocode();
      });
      this.$reverseGeocode.on('click', function(evt){
        evt.preventDefault();
        _this.reverseGeocode();
      });
    };

    Locator.prototype.getFormattedAddress = function() {
      var addressParts = [];
      if (this.$address.val() !== "") {
        addressParts.push(this.$address.val());
      }
      if (this.$city.val() !== "") {
        addressParts.push(this.$city.val());
      }
      if (this.$region.val() !== "") {
        addressParts.push(this.$region.val());
      }
      if (this.$country.val() !== "") {
        addressParts.push(this.$country.val());
      }
      return addressParts.join(", ");
    };

    Locator.prototype.getZoom = function() {
      var zoom;
      if(this.$address.val() != "") zoom = 17;
      else if(this.$zip.val() != "") zoom = 15;
      else if(this.$city.val() != "") zoom = 12;
      else if(this.$region.val() != "") zoom = 8;
      else if(this.$country.val() != "") zoom = 6;
      else if(this.$latitude.val() != "0" && this.$longitude.val() != "0") zoom = 10;
      else zoom = 2;

      return zoom;
    };

    Locator.prototype.geocode = function() {
      var _this = this,
        address = this.getFormattedAddress(),
        zoom = this.getZoom(),
        map = this.map,
        geocoder = this.geocoder,
        marker = this.marker;

      geocoder.geocode({'address': address}, function(results, status) {
        if (status === google.maps.GeocoderStatus.OK) {
          map.setCenter(results[0].geometry.location);
          map.setZoom(zoom);
          marker.setPosition(results[0].geometry.location);
          _this.setPosition();
          _this.renderCurrentAddress(results[0].formatted_address, results[0].geometry.location.lat(), results[0].geometry.location.lng())
        } else {
          alert('Geocode was not successful for the following reason: ' + status);
        };
      });
    };

    Locator.prototype.reverseGeocode = function() {
      var _this = this,
        geocoder = this.geocoder,
        marker = this.marker;

      geocoder.geocode({ 'latLng': marker.getPosition()}, function(results, status) {
        if (status == google.maps.GeocoderStatus.OK) {
          _this.setLocation(results[0]);
        }
      });
    };

    Locator.prototype.setCurrentAddress = function() {
      var _this = this,
        geocoder = this.geocoder,
        markerPosition = this.marker.getPosition(),
        currentAddress = '';

      geocoder.geocode({ 'latLng': markerPosition}, function(results, status) {
        if (status == google.maps.GeocoderStatus.OK) {
          _this.renderCurrentAddress(results[0].formatted_address, markerPosition.lat(), markerPosition.lng())
        }
      });
    };

    Locator.prototype.renderCurrentAddress = function(formattedAddress, lat, lng) {
        var template = '%{formatted_address} (%{lat}, %{lng})';
        currentAddress = template.replace('%{formatted_address}', formattedAddress);
        currentAddress = currentAddress.replace('%{lat}', lat);
        currentAddress = currentAddress.replace('%{lng}', lng);
        this.$currentAddress.text(currentAddress);
    };

    Locator.prototype.setPosition = function() {
      var markerPosition = this.marker.getPosition();
      this.$latitude.val(markerPosition.lat());
      this.$longitude.val(markerPosition.lng());
    };

    Locator.prototype.setLocation = function(geoResult) {
      var country = '', address = '', city = '', region = '', zip = '', latitude = '', longitude = '';

      var addressComponents = geoResult['address_components'];
      var values = { street_number : "", route : "", postal_code : "", administrative_area_level_1 : "", locality : "", country : ""};
      for (var compIndex in addressComponents){
        var types = addressComponents[compIndex].types;
        for(var typeIndex in types){
          if (types[typeIndex] == 'country') {
            country = addressComponents[compIndex].long_name;
          }
          values[types[typeIndex]] = addressComponents[compIndex].short_name;
        }
      }

      city = values["locality"] || values["sublocality"] || values["administrative_area_level_3"];
      region = values["country"] == "FR" ? values["administrative_area_level_1"] : (values["administrative_area_level_2"] || values["administrative_area_level_1"]);
      zip = values['postal_code']

      var divided_by_city = geoResult.formatted_address.split(city);
      var left_side = divided_by_city[0];
      if (zip && zip != "") {
        left_side = left_side.replace(zip, '');
      }
      if(left_side.replace(region, '') == left_side) {
        address = left_side;
      } else {
        address = divided_by_city[1];
      }
      address = address.replace(/[\,\s]+$/i, '');
      if(geoResult.formatted_address.match(/^\d*,/)) address = "";

      this.$country.val(country);
      this.$zip.val(zip);
      this.$region.val(region);
      this.$city.val(city);
      this.$address.val(address);
    };

    var Location = {
      Init : function() {
        $('[data-location="true"]').each(function() {
          var $this = $(this);
          var locator = new Locator($this);
          $this.data('locator', locator);
        });
      },

      InitAfterSlideOutOpen : function() {
        var _this = this;
        var loadedLocation = false;
        var loadingInterval = setInterval(function() {
          if (!loadedLocation && $('body').hasClass("qor-slideout-open")) {
            _this.Init();
            loadedLocation = true;
            clearInterval(loadingInterval);
          }
        }, 200)
      }
    };

    Export.QorLocation = Location;

  })(jQuery, window);
}).call(this);
