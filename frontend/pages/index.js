import React, {Component} from 'react';
import Link from 'next/link';
import Head from 'next/head';
import Slider from 'react-slick';

import MainLayout from '../components/MainLayout.js';
import initialProps from '../components/index.js';

class Index extends Component {
    render() {
        const sliderSettings = {
            dots: true,
            infinite: false,
            arrows: false,
            lazyLoad: true
        };

        const topBanner = this.props.homepagebanner.SerializableMeta.SlideImages.map((banner, index) => {
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

        const menCollection = decodeURIComponent(this.props.mencollection.SerializableMeta.Value);
        const womenCollection = decodeURIComponent(this.props.womencollection.SerializableMeta.Value);
        const newArrivalsPromotion = decodeURIComponent(this.props.newarrivalspromotion.SerializableMeta.Value);
        const modelProducts = decodeURIComponent(this.props.modelproducts.SerializableMeta.Value);

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
                    <div className="widget-newarrivals">
                        <div className="fullwidth-banner" dangerouslySetInnerHTML={{__html: newArrivalsPromotion}} />
                    </div>
                    <div className="widget-modelproducts">
                        <div className="fullwidth-banner" dangerouslySetInnerHTML={{__html: modelProducts}} />
                    </div>
                </MainLayout>
            </div>
        );
    }
}

Index.getInitialProps = initialProps;

export default Index;
