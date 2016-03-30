webpackJsonp([0],[
/* 0 */
/*!***********************************!*\
  !*** ./public/javascripts/app.js ***!
  \***********************************/
/***/ function(module, exports, __webpack_require__) {

	'use strict';
	
	__webpack_require__(/*! ../stylesheets/qor.scss */ 1);

/***/ },
/* 1 */
/*!*************************************!*\
  !*** ./public/stylesheets/qor.scss ***!
  \*************************************/
/***/ function(module, exports, __webpack_require__) {

	// style-loader: Adds some css to the DOM by adding a <style> tag
	
	// load the styles
	var content = __webpack_require__(/*! !./../../~/css-loader!./../../~/sass-loader?outputStyle=expanded!./qor.scss */ 2);
	if(typeof content === 'string') content = [[module.id, content, '']];
	// add the styles to the DOM
	var update = __webpack_require__(/*! ./../../~/style-loader/addStyles.js */ 4)(content, {});
	if(content.locals) module.exports = content.locals;
	// Hot Module Replacement
	if(false) {
		// When the styles change, update the <style> tags
		if(!content.locals) {
			module.hot.accept("!!./../../node_modules/css-loader/index.js!./../../node_modules/sass-loader/index.js?outputStyle=expanded!./qor.scss", function() {
				var newContent = require("!!./../../node_modules/css-loader/index.js!./../../node_modules/sass-loader/index.js?outputStyle=expanded!./qor.scss");
				if(typeof newContent === 'string') newContent = [[module.id, newContent, '']];
				update(newContent);
			});
		}
		// When the module is disposed, remove the <style> tags
		module.hot.dispose(function() { update(); });
	}

/***/ },
/* 2 */
/*!*****************************************************************************************!*\
  !*** ./~/css-loader!./~/sass-loader?outputStyle=expanded!./public/stylesheets/qor.scss ***!
  \*****************************************************************************************/
/***/ function(module, exports, __webpack_require__) {

	exports = module.exports = __webpack_require__(/*! ./../../~/css-loader/lib/css-base.js */ 3)();
	// imports
	
	
	// module
	exports.push([module.id, "/* Mixins\nhttp://bourbon.io/\n*/\n/* Base files. */\n/*! normalize.css v3.0.3 | MIT License | github.com/necolas/normalize.css */\nhtml {\n  font-family: sans-serif;\n  -ms-text-size-adjust: 100%;\n  -webkit-text-size-adjust: 100%;\n}\n\nbody {\n  margin: 0;\n}\n\narticle,\naside,\ndetails,\nfigcaption,\nfigure,\nfooter,\nheader,\nhgroup,\nmain,\nmenu,\nnav,\nsection,\nsummary {\n  display: block;\n}\n\naudio,\ncanvas,\nprogress,\nvideo {\n  display: inline-block;\n  vertical-align: baseline;\n}\n\naudio:not([controls]) {\n  display: none;\n  height: 0;\n}\n\n[hidden],\ntemplate {\n  display: none;\n}\n\na {\n  background-color: transparent;\n  text-decoration: none;\n}\n\na:active {\n  outline: 0;\n}\n\na:hover {\n  outline: 0;\n}\n\nabbr[title] {\n  border-bottom: 1px dotted;\n}\n\nb,\nstrong {\n  font-weight: bold;\n}\n\ndfn {\n  font-style: italic;\n}\n\nh1 {\n  font-size: 2em;\n  margin: 0.67em 0;\n}\n\nmark {\n  background: #ff0;\n  color: #000;\n}\n\nsmall {\n  font-size: 80%;\n}\n\nsub,\nsup {\n  font-size: 75%;\n  line-height: 0;\n  position: relative;\n  vertical-align: baseline;\n}\n\nsup {\n  top: -0.5em;\n}\n\nsub {\n  bottom: -0.25em;\n}\n\nimg {\n  border: 0;\n  max-width: 100%;\n}\n\nsvg:not(:root) {\n  overflow: hidden;\n}\n\nfigure {\n  margin: 1em 40px;\n}\n\nhr {\n  box-sizing: content-box;\n  height: 0;\n}\n\npre {\n  overflow: auto;\n}\n\ncode,\nkbd,\npre,\nsamp {\n  font-family: monospace, monospace;\n  font-size: 1em;\n}\n\nbutton,\ninput,\noptgroup,\nselect,\ntextarea {\n  color: inherit;\n  font: inherit;\n  margin: 0;\n}\n\nbutton {\n  overflow: visible;\n}\n\nbutton,\nselect {\n  text-transform: none;\n}\n\nbutton,\nhtml input[type=\"button\"],\ninput[type=\"reset\"],\ninput[type=\"submit\"] {\n  -webkit-appearance: button;\n  cursor: pointer;\n}\n\nbutton[disabled],\nhtml input[disabled] {\n  cursor: default;\n}\n\nbutton::-moz-focus-inner,\ninput::-moz-focus-inner {\n  border: 0;\n  padding: 0;\n}\n\ninput {\n  line-height: normal;\n}\n\ninput[type=\"checkbox\"],\ninput[type=\"radio\"] {\n  box-sizing: border-box;\n  padding: 0;\n}\n\ninput[type=\"number\"]::-webkit-inner-spin-button,\ninput[type=\"number\"]::-webkit-outer-spin-button {\n  height: auto;\n}\n\ninput[type=\"search\"] {\n  -webkit-appearance: textfield;\n  box-sizing: content-box;\n}\n\ninput[type=\"search\"]::-webkit-search-cancel-button,\ninput[type=\"search\"]::-webkit-search-decoration {\n  -webkit-appearance: none;\n}\n\nfieldset {\n  border: 1px solid #c0c0c0;\n  margin: 0 2px;\n  padding: 0.35em 0.625em 0.75em;\n}\n\nlegend {\n  border: 0;\n  padding: 0;\n}\n\ntextarea {\n  overflow: auto;\n}\n\noptgroup {\n  font-weight: bold;\n}\n\ntable {\n  border-collapse: collapse;\n  border-spacing: 0;\n}\n\ntd,\nth {\n  padding: 0;\n}\n\n/*\nBEM style\n\n.block {\n  @at-root __element {\n  }\n  @at-root --modifier {\n  }\n}\n */\n/*\nVariables\n*/\n/*\nBreakpoints\n*/\n/*\nColors\n*/\n/*\nTypography\nhttps://www.google.com/fonts/specimen/Titillium+Web\n*/\n/*\nGrid Variables\n*/\n/* calculates individual column width based off of # of columns */\n/* space between columns */\n/*\nMisc\n*/\nhtml {\n  font-size: 16px;\n}\n\nbody {\n  line-height: 1.6;\n  font-weight: 400;\n  font-family: \"Titillium Web\", \"Helvetica Neue\", Helvetica, Arial, sans-serif;\n  color: rgba(0,0,0,0.75);\n}\n\na {\n  color: #25A5DF;\n}\n\nhr {\n  margin-top: 3rem;\n  margin-bottom: 3.5rem;\n  border-width: 0;\n  border-top: 1px solid rgba(0,0,0,0.54);\n}\n\nul, li {\n  margin: 0;\n  padding: 0;\n  list-style: none;\n}\n\n.u-full-width {\n  width: 100%;\n  box-sizing: border-box;\n}\n\n.u-max-full-width {\n  max-width: 100%;\n  box-sizing: border-box;\n}\n\n.u-pull-right {\n  float: right;\n}\n\n.u-pull-left {\n  float: left;\n}\n\nh1, h2, h3, h4, h5, h6 {\n  margin-top: 0;\n  margin-bottom: 8px;\n  font-weight: 600;\n}\n\nh1 {\n  font-size: 32px;\n  line-height: 1.2;\n}\n\nh2 {\n  font-size: 30px;\n  line-height: 1.25;\n}\n\nh3 {\n  font-size: 26px;\n  line-height: 1.3;\n}\n\nh4 {\n  font-size: 24px;\n  line-height: 1.35;\n}\n\nh5 {\n  font-size: 20px;\n  line-height: 1.5;\n}\n\nh6 {\n  font-size: 18px;\n  line-height: 1.6;\n}\n\np {\n  margin-top: 0;\n}\n\nstrong {\n  font-weight: 600;\n}\n\n/* Modules */\n.container {\n  position: relative;\n  width: 100%;\n  max-width: 1000px;\n  margin: 0 auto;\n  padding: 0 20px;\n  box-sizing: border-box;\n}\n\n.column {\n  width: 100%;\n  float: left;\n  box-sizing: border-box;\n}\n\n@media (min-width: 400px) {\n  .container {\n    width: 85%;\n    padding: 0;\n  }\n}\n\n@media (min-width: 550px) {\n  .container {\n    width: 80%;\n  }\n  .column {\n    margin-left: 4%;\n  }\n  .column:first-child {\n    margin-left: 0;\n  }\n  .column-1 {\n    width: 4.66667%;\n  }\n  .column-2 {\n    width: 13.33333%;\n  }\n  .column-3 {\n    width: 22%;\n  }\n  .column-4 {\n    width: 30.66667%;\n  }\n  .column-5 {\n    width: 39.33333%;\n  }\n  .column-6 {\n    width: 48%;\n  }\n  .column-7 {\n    width: 56.66667%;\n  }\n  .column-8 {\n    width: 65.33333%;\n  }\n  .column-9 {\n    width: 74%;\n  }\n  .column-10 {\n    width: 82.66667%;\n  }\n  .column-11 {\n    width: 91.33333%;\n  }\n  .column-12 {\n    width: 100%;\n    margin-left: 0;\n  }\n  /* 1/3 */\n  .one-third.column {\n    width: 30.66667%;\n  }\n  /* 2/3 */\n  .two-thirds.column {\n    width: 65.33333%;\n  }\n  /* 1/2 */\n  .one-half.column {\n    width: 48%;\n  }\n  .offset-by-1.column {\n    margin-left: 8.66667%;\n  }\n  .offset-by-2.column {\n    margin-left: 17.33333%;\n  }\n  .offset-by-3.column {\n    margin-left: 26%;\n  }\n  .offset-by-4.column {\n    margin-left: 34.66667%;\n  }\n  .offset-by-5.column {\n    margin-left: 43.33333%;\n  }\n  .offset-by-6.column {\n    margin-left: 52%;\n  }\n  .offset-by-7.column {\n    margin-left: 60.66667%;\n  }\n  .offset-by-8.column {\n    margin-left: 69.33333%;\n  }\n  .offset-by-9.column {\n    margin-left: 78%;\n  }\n  .offset-by-10.column {\n    margin-left: 86.66667%;\n  }\n  .offset-by-11.column {\n    margin-left: 95.33333%;\n  }\n  .offset-by-one-third.column {\n    margin-left: 34.66667%;\n  }\n  .offset-by-two-thirds.column {\n    margin-left: 69.33333%;\n  }\n  .offset-by-one-half.column {\n    margin-left: 52%;\n  }\n}\n\n.container:after,\n.row:after,\n.u-cf {\n  content: \"\";\n  display: table;\n  clear: both;\n}\n\n.button {\n  display: inline-block;\n  height: 48px;\n  padding: 0 32px;\n  color: #33c3f0;\n  text-align: center;\n  font-size: 22px;\n  font-weight: 400;\n  line-height: 48px;\n  text-transform: uppercase;\n  text-decoration: none;\n  white-space: nowrap;\n  border-radius: 2px;\n  cursor: pointer;\n  box-sizing: border-box;\n}\n\ninput[type=\"submit\"], input[type=\"button\"] {\n  display: inline-block;\n  height: 38px;\n  padding: 0 30px;\n  color: #33c3f0;\n  text-align: center;\n  font-size: 11px;\n  font-weight: 400;\n  line-height: 38px;\n  text-transform: uppercase;\n  text-decoration: none;\n  white-space: nowrap;\n  background-color: transparent;\n  border-radius: 2px;\n  cursor: pointer;\n  box-sizing: border-box;\n}\n\n.button.button__primary {\n  color: #fff;\n  background-color: #FF9654;\n  background-image: -webkit-linear-gradient(#FF9654, #FF6D0C);\n  background-image: linear-gradient(#FF9654, #FF6D0C);\n  border: 1px solid #FF8335;\n}\n\ninput[type=\"submit\"].button__primary, input[type=\"reset\"].button__primary, input[type=\"button\"].button__primary {\n  color: #fff;\n  background-color: #33c3f0;\n  border-color: #33c3f0;\n}\n\ninput[type=\"submit\"].button__primary:hover, input[type=\"reset\"].button__primary:hover, input[type=\"button\"].button__primary:hover {\n  color: #fff;\n  background-color: #25A5DF;\n  border-color: #25A5DF;\n}\n\n.button.button__primary:focus {\n  color: #fff;\n  background-color: #25A5DF;\n  border-color: #25A5DF;\n}\n\ninput[type=\"email\"], input[type=\"number\"], input[type=\"search\"], input[type=\"text\"], input[type=\"tel\"], input[type=\"url\"], input[type=\"password\"] {\n  height: 38px;\n  padding: 6px 10px;\n  background-color: #fff;\n  border: 1px solid rgba(0,0,0,0.12);\n  border-radius: 2px;\n  box-shadow: none;\n  box-sizing: border-box;\n}\n\ntextarea,\nselect {\n  height: 38px;\n  padding: 6px 10px;\n  background-color: #fff;\n  border: 1px solid rgba(0,0,0,0.12);\n  border-radius: 2px;\n  box-shadow: none;\n  box-sizing: border-box;\n}\n\ninput[type=\"email\"], input[type=\"number\"], input[type=\"search\"], input[type=\"text\"], input[type=\"tel\"], input[type=\"url\"], input[type=\"password\"] {\n  -webkit-appearance: none;\n  -moz-appearance: none;\n  appearance: none;\n}\n\ntextarea {\n  -webkit-appearance: none;\n  -moz-appearance: none;\n  appearance: none;\n  min-height: 65px;\n  padding-top: 6px;\n  padding-bottom: 6px;\n}\n\ninput[type=\"email\"]:focus, input[type=\"number\"]:focus, input[type=\"search\"]:focus, input[type=\"text\"]:focus, input[type=\"tel\"]:focus, input[type=\"url\"]:focus, input[type=\"password\"]:focus {\n  outline: 0;\n}\n\ntextarea:focus,\nselect:focus {\n  border: 1px solid #33c3f0;\n  outline: 0;\n}\n\nlabel,\nlegend {\n  display: block;\n  margin-bottom: .5rem;\n  font-weight: 600;\n}\n\nfieldset {\n  padding: 0;\n  border-width: 0;\n}\n\ninput[type=\"checkbox\"], input[type=\"radio\"] {\n  display: inline;\n}\n\nlabel > .label-body {\n  display: inline-block;\n  margin-left: .5rem;\n  font-weight: normal;\n}\n\nul {\n  list-style: none;\n}\n\nol {\n  list-style: none;\n  padding-left: 0;\n  margin-top: 0;\n}\n\nul {\n  padding-left: 0;\n  margin-top: 0;\n}\n\nul ul, ul ol {\n  margin: 1.5rem 0 1.5rem 3rem;\n  font-size: 90%;\n}\n\nol ol, ol ul {\n  margin: 1.5rem 0 1.5rem 3rem;\n  font-size: 90%;\n}\n\nli {\n  margin-bottom: 1rem;\n}\n\ncode {\n  padding: .2rem .5rem;\n  margin: 0 .2rem;\n  font-size: 90%;\n  white-space: nowrap;\n  background: rgba(0,0,0,0.12);\n  border: 1px solid rgba(0,0,0,0.54);\n  border-radius: 2px;\n}\n\npre > code {\n  display: block;\n  padding: 1rem 1.5rem;\n  white-space: pre;\n}\n\nth,\ntd {\n  padding: 12px 15px;\n  text-align: left;\n  border-bottom: 1px solid rgba(0,0,0,0.54);\n}\n\nth:first-child,\ntd:first-child {\n  padding-left: 0;\n}\n\nth:last-child,\ntd:last-child {\n  padding-right: 0;\n}\n\nbutton,\n.button,\ninput,\ntextarea,\nselect,\nfieldset,\npre,\nblockquote,\ndl,\nfigure,\ntable,\np,\nul,\nol,\nform {\n  margin: 0;\n}\n\n/* Layout */\n@font-face {\n  font-family: 'Titillium Web';\n  src: url(\"/fonts/titilliumweb-bold-webfont.eot\");\n  src: url(\"/fonts/titilliumweb-bold-webfont.eot?#iefix\") format(\"embedded-opentype\"), url(\"/fonts/titilliumweb-bold-webfont.woff2\") format(\"woff2\"), url(\"/fonts/titilliumweb-bold-webfont.woff\") format(\"woff\"), url(\"/fonts/titilliumweb-bold-webfont.ttf\") format(\"truetype\"), url(\"/fonts/titilliumweb-bold-webfont.svg#titillium_webbold\") format(\"svg\");\n  font-weight: 700;\n  font-style: normal;\n}\n\n@font-face {\n  font-family: 'Titillium Web';\n  src: url(\"/fonts/titilliumweb-light-webfont.eot\");\n  src: url(\"/fonts/titilliumweb-light-webfont.eot?#iefix\") format(\"embedded-opentype\"), url(\"/fonts/titilliumweb-light-webfont.woff2\") format(\"woff2\"), url(\"/fonts/titilliumweb-light-webfont.woff\") format(\"woff\"), url(\"/fonts/titilliumweb-light-webfont.ttf\") format(\"truetype\"), url(\"/fonts/titilliumweb-light-webfont.svg#titillium_weblight\") format(\"svg\");\n  font-weight: 300;\n  font-style: normal;\n}\n\n@font-face {\n  font-family: 'Titillium Web';\n  src: url(\"/fonts/titilliumweb-regular-webfont.eot\");\n  src: url(\"/fonts/titilliumweb-regular-webfont.eot?#iefix\") format(\"embedded-opentype\"), url(\"/fonts/titilliumweb-regular-webfont.woff2\") format(\"woff2\"), url(\"/fonts/titilliumweb-regular-webfont.woff\") format(\"woff\"), url(\"/fonts/titilliumweb-regular-webfont.ttf\") format(\"truetype\"), url(\"/fonts/titilliumweb-regular-webfont.svg#titillium_webregular\") format(\"svg\");\n  font-weight: 400;\n  font-style: normal;\n}\n\n@font-face {\n  font-family: 'Titillium Web';\n  src: url(\"/fonts/titilliumweb-semibold-webfont.eot\");\n  src: url(\"/fonts/titilliumweb-semibold-webfont.eot?#iefix\") format(\"embedded-opentype\"), url(\"/fonts/titilliumweb-semibold-webfont.woff2\") format(\"woff2\"), url(\"/fonts/titilliumweb-semibold-webfont.woff\") format(\"woff\"), url(\"/fonts/titilliumweb-semibold-webfont.ttf\") format(\"truetype\"), url(\"/fonts/titilliumweb-semibold-webfont.svg#titillium_websemibold\") format(\"svg\");\n  font-weight: 600;\n  font-style: normal;\n}\n\n.header {\n  border-bottom: 1px solid rgba(0,0,0,0.12);\n}\n\n.header header {\n  height: 64px;\n}\n\n.header header .logo {\n  margin-left: 0;\n  margin-right: 64px;\n  padding: 0;\n}\n\n.header header .logo img {\n  width: 112px;\n  margin-top: 17px;\n}\n\n.header header .header__func {\n  text-align: right;\n}\n\n.header header .header__func a {\n  color: rgba(0,0,0,0.54);\n  font-size: 14px;\n  width: 33%;\n  text-align: center;\n  padding: 0;\n}\n\n.header header a {\n  color: rgba(0,0,0,0.75);\n  font-size: 18px;\n  height: 64px;\n  line-height: 64px;\n  padding: 0 16px;\n  float: left;\n}\n\n.footer {\n  background-color: rgba(0,0,0,0.12);\n  border-top: 1px solid rgba(0,0,0,0.23);\n  padding: 32px 0;\n  margin-top: 32px;\n}\n\n.footer h4 {\n  font-weight: 600;\n  font-size: 20px;\n  color: rgba(0,0,0,0.54);\n}\n\n.footer li {\n  width: 48%;\n  display: inline-block;\n}\n\n.footer li a {\n  color: rgba(0,0,0,0.75);\n}\n\n.footer p {\n  font-size: 14px;\n}\n\n.footer p a {\n  color: rgba(0,0,0,0.75);\n  text-decoration: underline;\n}\n\n.footer .society {\n  margin-bottom: 48px;\n}\n\n.footer .society img {\n  height: 24px;\n  margin-right: 16px;\n}\n\n.banner {\n  min-height: 425px;\n  width: 100%;\n  text-align: center;\n  background: url(\"/images/banner.png\") no-repeat center center;\n}\n\n.banner__logo {\n  padding-top: 150px;\n  margin-left: 232px;\n  width: 260px;\n  display: block;\n}\n\n.banner h1 {\n  color: #fff;\n  padding-left: 57px;\n  margin-top: -18px;\n}\n\n.section__title {\n  padding: 64px 0;\n}\n\n.section__title h3 {\n  text-align: center;\n  font-weight: 700;\n}\n\n.section__title p {\n  font-size: 18px;\n  color: rgba(0,0,0,0.54);\n  text-align: center;\n}\n\n.products {\n  /* homepage product list */\n  /* homepage product list end */\n  /* products details page featured lists */\n  /* products details page */\n}\n\n.products__list.row-three li:nth-child(3n+1), .products__list.row-four li:nth-child(4n+1) {\n  margin-left: 0;\n}\n\n.products__list .products__list--img {\n  border: 1px solid rgba(0,0,0,0.12);\n  display: block;\n  height: 321px;\n  box-sizing: border-box;\n  margin-bottom: 12px;\n  text-align: center;\n}\n\n.products__list .products__list--price {\n  float: right;\n  font-size: 18px;\n}\n\n.products__list .products__list--title {\n  float: left;\n  font-size: 20px;\n}\n\n.products__featured {\n  border-top: 1px solid rgba(0,0,0,0.12);\n  padding-top: 32px;\n}\n\n.products__featured .products__featured--swiper {\n  position: relative;\n  margin: 32px 0;\n}\n\n.products__featured .swiper-button-next {\n  right: -50px;\n}\n\n.products__featured .swiper-button-prev {\n  left: -50px;\n}\n\n.products__featured .products__list li {\n  margin-right: 16px;\n}\n\n.products__featured .products__list--img {\n  height: 234px;\n}\n\n.products__featured .products__list--img img {\n  max-height: 232px;\n}\n\n.products__details {\n  margin: 32px auto;\n}\n\n.products__details .products__details--share > a {\n  float: left;\n}\n\n.products__details .products__details--share > a img {\n  margin-right: 8px;\n}\n\n.products__details .products__details--share > a img, .products__details .products__details--share > a span {\n  display: inline-block;\n  vertical-align: middle;\n}\n\n.products__details .products__details--share > div {\n  float: right;\n}\n\n.products__details .products__details--share > div a {\n  margin-left: 8px;\n}\n\n.products__details .products__details--share > div a, .products__details .products__details--share > div span {\n  display: inline-block;\n  vertical-align: middle;\n  line-height: 1;\n}\n\n.products__info {\n  margin: 32px 0;\n  padding: 32px 0 0 0;\n  border-top: 1px solid rgba(0,0,0,0.12);\n}\n\n.products__meta .button {\n  width: 100%;\n}\n\n.products__meta > ul > li {\n  margin-bottom: 32px;\n}\n\n.products__meta a {\n  color: rgba(0,0,0,0.75);\n}\n\n.products__meta select {\n  display: block;\n  -webkit-appearance: none;\n  -moz-appearance: none;\n  -ms-appearance: none;\n  -o-appearance: none;\n  appearance: none;\n  border: 1px solid rgba(0,0,0,0.12);\n  background: rgba(0,0,0,0.12) url(\"/images/icon-arrow-down.png\") no-repeat 96% center;\n  background-size: 12px;\n  width: 100%;\n  margin-bottom: 8px;\n}\n\n.products__meta select + a {\n  text-decoration: underline;\n}\n\n.products__meta h1 {\n  font-weight: 600;\n}\n\n.products__meta strong {\n  font-size: 18px;\n  padding-bottom: 8px;\n  display: inline-block;\n}\n\n.products__meta .products__meta--qty {\n  position: relative;\n  height: 40px;\n}\n\n.products__meta .products__meta--qty input {\n  width: 100%;\n  text-align: center;\n  height: 40px;\n}\n\n.products__meta .products__meta--qty a {\n  position: absolute;\n  top: 0;\n  bottom: 0;\n  border: 1px solid rgba(0,0,0,0.12);\n  border-radius: 2px;\n  background-color: rgba(0,0,0,0.12);\n  width: 88px;\n  line-height: 34px;\n  font-size: 24px;\n  text-align: center;\n}\n\n.products__meta .products__meta--qty a.add {\n  right: 0;\n}\n\n.products__meta .products__meta--qty a.reduce {\n  left: 0;\n}\n\n.products__meta .products__meta--color span {\n  display: inline-block;\n  vertical-align: middle;\n  width: 32px;\n  height: 32px;\n  border-radius: 100%;\n  margin-right: 8px;\n  border: 1px solid rgba(0,0,0,0.12);\n  margin-top: 8px;\n}\n\n.products__meta .products__meta--color span.selected {\n  border: 0;\n  box-shadow: 0 0 0 2px #fff, 0 0 0 4px rgba(0,0,0,0.75);\n}\n\n.products__meta .products__meta--price {\n  font-size: 32px;\n  color: #33c3f0;\n  text-align: right;\n}\n\n.products__gallery--thumbs {\n  margin-top: 16px;\n}\n\n.products__gallery--thumbs .swiper-wrapper div {\n  height: 80px;\n  width: 80px;\n}\n\n.products__gallery--thumbs .swiper-slide {\n  background-size: cover;\n  background-position: center;\n  border: 1px solid #ccc;\n  box-sizing: border-box;\n}\n\n.products__gallery--thumbs .swiper-slide-active {\n  border: 1px solid #33c3f0;\n}\n\n.products__gallery--top .swiper-wrapper div {\n  height: 653px;\n  width: 653px;\n}\n\n.products__gallery--top .swiper-slide {\n  background-size: cover;\n  background-position: center;\n  box-sizing: border-box;\n}\n\n/* footer goodness section */\n.goodness {\n  background-color: rgba(0,0,0,0.12);\n  padding: 0 0 32px 0;\n  color: rgba(0,0,0,0.75);\n}\n\n.goodness .section__title {\n  padding: 32px 0;\n}\n\n.goodness .store {\n  background: url(\"/images/store.png\") no-repeat center center;\n  height: 260px;\n}\n\n.goodness .newsletter {\n  background: url(\"/images/newsletter.png\") no-repeat center center;\n  height: 260px;\n}\n\n.goodness a {\n  display: block;\n  background-color: rgb(255,255,255);\n  padding: 8px 0;\n  text-align: center;\n  font-size: 18px;\n  color: rgba(0,0,0,0.75);\n  margin-top: 125px;\n}\n\n.goodness a img {\n  width: 7px;\n}\n\n.goodness a + p {\n  font-size: 16px;\n  padding: 8px 16px;\n}\n", ""]);
	
	// exports


/***/ },
/* 3 */
/*!**************************************!*\
  !*** ./~/css-loader/lib/css-base.js ***!
  \**************************************/
/***/ function(module, exports) {

	/*
		MIT License http://www.opensource.org/licenses/mit-license.php
		Author Tobias Koppers @sokra
	*/
	// css base code, injected by the css-loader
	module.exports = function() {
		var list = [];
	
		// return the list of modules as css string
		list.toString = function toString() {
			var result = [];
			for(var i = 0; i < this.length; i++) {
				var item = this[i];
				if(item[2]) {
					result.push("@media " + item[2] + "{" + item[1] + "}");
				} else {
					result.push(item[1]);
				}
			}
			return result.join("");
		};
	
		// import a list of modules into the list
		list.i = function(modules, mediaQuery) {
			if(typeof modules === "string")
				modules = [[null, modules, ""]];
			var alreadyImportedModules = {};
			for(var i = 0; i < this.length; i++) {
				var id = this[i][0];
				if(typeof id === "number")
					alreadyImportedModules[id] = true;
			}
			for(i = 0; i < modules.length; i++) {
				var item = modules[i];
				// skip already imported module
				// this implementation is not 100% perfect for weird media query combinations
				//  when a module is imported multiple times with different media queries.
				//  I hope this will never occur (Hey this way we have smaller bundles)
				if(typeof item[0] !== "number" || !alreadyImportedModules[item[0]]) {
					if(mediaQuery && !item[2]) {
						item[2] = mediaQuery;
					} else if(mediaQuery) {
						item[2] = "(" + item[2] + ") and (" + mediaQuery + ")";
					}
					list.push(item);
				}
			}
		};
		return list;
	};


/***/ },
/* 4 */
/*!*************************************!*\
  !*** ./~/style-loader/addStyles.js ***!
  \*************************************/
/***/ function(module, exports, __webpack_require__) {

	/*
		MIT License http://www.opensource.org/licenses/mit-license.php
		Author Tobias Koppers @sokra
	*/
	var stylesInDom = {},
		memoize = function(fn) {
			var memo;
			return function () {
				if (typeof memo === "undefined") memo = fn.apply(this, arguments);
				return memo;
			};
		},
		isOldIE = memoize(function() {
			return /msie [6-9]\b/.test(window.navigator.userAgent.toLowerCase());
		}),
		getHeadElement = memoize(function () {
			return document.head || document.getElementsByTagName("head")[0];
		}),
		singletonElement = null,
		singletonCounter = 0,
		styleElementsInsertedAtTop = [];
	
	module.exports = function(list, options) {
		if(true) {
			if(typeof document !== "object") throw new Error("The style-loader cannot be used in a non-browser environment");
		}
	
		options = options || {};
		// Force single-tag solution on IE6-9, which has a hard limit on the # of <style>
		// tags it will allow on a page
		if (typeof options.singleton === "undefined") options.singleton = isOldIE();
	
		// By default, add <style> tags to the bottom of <head>.
		if (typeof options.insertAt === "undefined") options.insertAt = "bottom";
	
		var styles = listToStyles(list);
		addStylesToDom(styles, options);
	
		return function update(newList) {
			var mayRemove = [];
			for(var i = 0; i < styles.length; i++) {
				var item = styles[i];
				var domStyle = stylesInDom[item.id];
				domStyle.refs--;
				mayRemove.push(domStyle);
			}
			if(newList) {
				var newStyles = listToStyles(newList);
				addStylesToDom(newStyles, options);
			}
			for(var i = 0; i < mayRemove.length; i++) {
				var domStyle = mayRemove[i];
				if(domStyle.refs === 0) {
					for(var j = 0; j < domStyle.parts.length; j++)
						domStyle.parts[j]();
					delete stylesInDom[domStyle.id];
				}
			}
		};
	}
	
	function addStylesToDom(styles, options) {
		for(var i = 0; i < styles.length; i++) {
			var item = styles[i];
			var domStyle = stylesInDom[item.id];
			if(domStyle) {
				domStyle.refs++;
				for(var j = 0; j < domStyle.parts.length; j++) {
					domStyle.parts[j](item.parts[j]);
				}
				for(; j < item.parts.length; j++) {
					domStyle.parts.push(addStyle(item.parts[j], options));
				}
			} else {
				var parts = [];
				for(var j = 0; j < item.parts.length; j++) {
					parts.push(addStyle(item.parts[j], options));
				}
				stylesInDom[item.id] = {id: item.id, refs: 1, parts: parts};
			}
		}
	}
	
	function listToStyles(list) {
		var styles = [];
		var newStyles = {};
		for(var i = 0; i < list.length; i++) {
			var item = list[i];
			var id = item[0];
			var css = item[1];
			var media = item[2];
			var sourceMap = item[3];
			var part = {css: css, media: media, sourceMap: sourceMap};
			if(!newStyles[id])
				styles.push(newStyles[id] = {id: id, parts: [part]});
			else
				newStyles[id].parts.push(part);
		}
		return styles;
	}
	
	function insertStyleElement(options, styleElement) {
		var head = getHeadElement();
		var lastStyleElementInsertedAtTop = styleElementsInsertedAtTop[styleElementsInsertedAtTop.length - 1];
		if (options.insertAt === "top") {
			if(!lastStyleElementInsertedAtTop) {
				head.insertBefore(styleElement, head.firstChild);
			} else if(lastStyleElementInsertedAtTop.nextSibling) {
				head.insertBefore(styleElement, lastStyleElementInsertedAtTop.nextSibling);
			} else {
				head.appendChild(styleElement);
			}
			styleElementsInsertedAtTop.push(styleElement);
		} else if (options.insertAt === "bottom") {
			head.appendChild(styleElement);
		} else {
			throw new Error("Invalid value for parameter 'insertAt'. Must be 'top' or 'bottom'.");
		}
	}
	
	function removeStyleElement(styleElement) {
		styleElement.parentNode.removeChild(styleElement);
		var idx = styleElementsInsertedAtTop.indexOf(styleElement);
		if(idx >= 0) {
			styleElementsInsertedAtTop.splice(idx, 1);
		}
	}
	
	function createStyleElement(options) {
		var styleElement = document.createElement("style");
		styleElement.type = "text/css";
		insertStyleElement(options, styleElement);
		return styleElement;
	}
	
	function createLinkElement(options) {
		var linkElement = document.createElement("link");
		linkElement.rel = "stylesheet";
		insertStyleElement(options, linkElement);
		return linkElement;
	}
	
	function addStyle(obj, options) {
		var styleElement, update, remove;
	
		if (options.singleton) {
			var styleIndex = singletonCounter++;
			styleElement = singletonElement || (singletonElement = createStyleElement(options));
			update = applyToSingletonTag.bind(null, styleElement, styleIndex, false);
			remove = applyToSingletonTag.bind(null, styleElement, styleIndex, true);
		} else if(obj.sourceMap &&
			typeof URL === "function" &&
			typeof URL.createObjectURL === "function" &&
			typeof URL.revokeObjectURL === "function" &&
			typeof Blob === "function" &&
			typeof btoa === "function") {
			styleElement = createLinkElement(options);
			update = updateLink.bind(null, styleElement);
			remove = function() {
				removeStyleElement(styleElement);
				if(styleElement.href)
					URL.revokeObjectURL(styleElement.href);
			};
		} else {
			styleElement = createStyleElement(options);
			update = applyToTag.bind(null, styleElement);
			remove = function() {
				removeStyleElement(styleElement);
			};
		}
	
		update(obj);
	
		return function updateStyle(newObj) {
			if(newObj) {
				if(newObj.css === obj.css && newObj.media === obj.media && newObj.sourceMap === obj.sourceMap)
					return;
				update(obj = newObj);
			} else {
				remove();
			}
		};
	}
	
	var replaceText = (function () {
		var textStore = [];
	
		return function (index, replacement) {
			textStore[index] = replacement;
			return textStore.filter(Boolean).join('\n');
		};
	})();
	
	function applyToSingletonTag(styleElement, index, remove, obj) {
		var css = remove ? "" : obj.css;
	
		if (styleElement.styleSheet) {
			styleElement.styleSheet.cssText = replaceText(index, css);
		} else {
			var cssNode = document.createTextNode(css);
			var childNodes = styleElement.childNodes;
			if (childNodes[index]) styleElement.removeChild(childNodes[index]);
			if (childNodes.length) {
				styleElement.insertBefore(cssNode, childNodes[index]);
			} else {
				styleElement.appendChild(cssNode);
			}
		}
	}
	
	function applyToTag(styleElement, obj) {
		var css = obj.css;
		var media = obj.media;
		var sourceMap = obj.sourceMap;
	
		if(media) {
			styleElement.setAttribute("media", media)
		}
	
		if(styleElement.styleSheet) {
			styleElement.styleSheet.cssText = css;
		} else {
			while(styleElement.firstChild) {
				styleElement.removeChild(styleElement.firstChild);
			}
			styleElement.appendChild(document.createTextNode(css));
		}
	}
	
	function updateLink(linkElement, obj) {
		var css = obj.css;
		var media = obj.media;
		var sourceMap = obj.sourceMap;
	
		if(sourceMap) {
			// http://stackoverflow.com/a/26603875
			css += "\n/*# sourceMappingURL=data:application/json;base64," + btoa(unescape(encodeURIComponent(JSON.stringify(sourceMap)))) + " */";
		}
	
		var blob = new Blob([css], { type: "text/css" });
	
		var oldSrc = linkElement.href;
	
		linkElement.href = URL.createObjectURL(blob);
	
		if(oldSrc)
			URL.revokeObjectURL(oldSrc);
	}


/***/ }
]);
//# sourceMappingURL=app.js.map