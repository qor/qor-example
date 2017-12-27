import React, {Component} from 'react';
import Link from 'next/link';
import Head from 'next/head';
import Slider from 'react-slick';

import MainLayout from '../components/MainLayout.js';

require('es6-promise').polyfill();
require('isomorphic-fetch');

// require('slick-carousel/slick/slick.css');
// require "~slick-carousel/slick/slick-theme.css";

const widgetPrefix = 'http://localhost:7000/admin/page_builder_widgets/';

const topBannerJSON = `${widgetPrefix}home%20page%20banner.json`;
const menCollectionJSON = `${widgetPrefix}men%20collection.json`;
const womenCollectionJSON = `${widgetPrefix}women%20collection.json`;

class Index extends Component {
    static async getInitialProps(context) {
        const topBannerRes = await fetch(topBannerJSON);
        const topBanner = await topBannerRes.json();

        const menCollectionRes = await fetch(menCollectionJSON);
        const menCollection = await menCollectionRes.json();

        const womenCollectionRes = await fetch(womenCollectionJSON);
        const womenCollection = await womenCollectionRes.json();

        return {
            topBanner: topBanner,
            menCollection: menCollection,
            womenCollection: womenCollection
        };
    }

    render() {
        const sliderSettings = {
            dots: true,
            infinite: false,
            arrows: false,
            lazyLoad: true
        };

        const topBanner = this.props.topBanner.SerializableMeta.SlideImages.map((banner, index) => {
            return (
                <div key={index}>
                    <div className="qor-slider__texts">
                        <h1>{banner.Title}</h1>
                        <h2>{banner.SubTitle}</h2>
                        <Link href={banner.Link}>
                            <a>{banner.Button}</a>
                        </Link>
                    </div>
                    <img src={`http://localhost:7000${banner.Image.Url}`} />
                </div>
            );
        });

        const menCollection = decodeURIComponent(this.props.menCollection.SerializableMeta.Value);
        const womenCollection = decodeURIComponent(this.props.womenCollection.SerializableMeta.Value);

        return (
            <div>
                <Head>
                    <link rel="stylesheet" href="/static/stylesheets/main.css" />
                </Head>
                <MainLayout>
                    <Slider {...sliderSettings}>{topBanner}</Slider>
                    <div className="col-2-banner">
                        <div className="fullwidth-banner" dangerouslySetInnerHTML={{__html: menCollection}} />
                        <div className="fullwidth-banner" dangerouslySetInnerHTML={{__html: womenCollection}} />
                    </div>
                    <div className="global-messages">
                        FREE RUNNING PACK <span className="light">WHEN YOU SPEND $200 OR MORE</span>
                    </div>
                </MainLayout>
            </div>
        );
    }
}

export default Index;
